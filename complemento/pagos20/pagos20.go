package pagos20

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

func Unmarshal(b []byte) (*Pagos, error) {
	pagos := &Pagos{}
	if err := xml.Unmarshal(b, pagos); err != nil {
		return nil, err
	}
	return pagos, nil
}

// Pagos Complemento para el Comprobante Fiscal Digital por Internet (CFDI) para registrar información sobre la recepción de pagos. El emisor de este complemento para recepción de pagos debe ser quien las leyes le obligue a expedir comprobantes por los actos o actividades que realicen, por los ingresos que se perciban o por las retenciones de contribuciones que efectúen.
type Pagos struct {
	// Totales Nodo requerido para especificar el monto total de los pagos y el total de los impuestos, deben ser expresados en MXN.
	Totales *Totales `xml:"Totales"`
	// Pago Elemento requerido para incorporar la información de la recepción de pagos.
	Pago []*Pago `xml:"Pago"`
	// Version Atributo requerido que indica la versión del complemento para recepción de pagos.
	Version string `xml:"Version,attr"`
}

func (pagos Pagos) TotalImpuestoTrasladoIEPSMXN() decimal.NullDecimal {
	ieps := decimal.New(0, 0)
	hasIEPS := false
	for _, pago := range pagos.Pago {
		if pago.ImpuestosP == nil || pago.ImpuestosP.TrasladosP == nil {
			continue
		}
		for _, tras := range pago.ImpuestosP.TrasladosP {
			if tras.ImpuestoP == types.ImpuestoIEPS {
				hasIEPS = true
				importeMxn := decimalMxn(tras.ImporteP, pago.MonedaP, pago.TipoCambioP)
				ieps = ieps.Add(importeMxn)
			}
		}
	}
	if hasIEPS {
		return decimal.NullDecimal{Valid: true, Decimal: ieps}
	}
	return decimal.NullDecimal{}
}

// Totales Nodo requerido para especificar el monto total de los pagos y el total de los impuestos, deben ser expresados en MXN.
type Totales struct {
	// TotalRetencionesIVA Atributo condicional para expresar el total de los impuestos retenidos de IVA que se desprenden de los pagos. No se permiten valores negativos.
	TotalRetencionesIVA decimal.NullDecimal `xml:"TotalRetencionesIVA,attr,omitempty"`
	// TotalRetencionesISR Atributo condicional para expresar el total de los impuestos retenidos de ISR que se desprenden de los pagos. No se permiten valores negativos.
	TotalRetencionesISR decimal.NullDecimal `xml:"TotalRetencionesISR,attr,omitempty"`
	// TotalRetencionesIEPS Atributo condicional para expresar el total de los impuestos retenidos de IEPS que se desprenden de los pagos. No se permiten valores negativos.
	TotalRetencionesIEPS decimal.NullDecimal `xml:"TotalRetencionesIEPS,attr,omitempty"`
	// TotalTrasladosBaseIVA16 Atributo condicional para expresar el total de la base de IVA trasladado a la tasa del 16% que se desprende de los pagos. No se permiten valores negativos.
	TotalTrasladosBaseIVA16 decimal.NullDecimal `xml:"TotalTrasladosBaseIVA16,attr,omitempty"`
	// TotalTrasladosImpuestoIVA16 Atributo condicional para expresar el total de los impuestos de IVA trasladado a la tasa del 16% que se desprenden de los pagos. No se permiten valores negativos.
	TotalTrasladosImpuestoIVA16 decimal.NullDecimal `xml:"TotalTrasladosImpuestoIVA16,attr,omitempty"`
	// TotalTrasladosBaseIVA8 Atributo condicional para expresar el total de la base de IVA trasladado a la tasa del 8% que se desprende de los pagos. No se permiten valores negativos.
	TotalTrasladosBaseIVA8 decimal.NullDecimal `xml:"TotalTrasladosBaseIVA8,attr,omitempty"`
	// TotalTrasladosImpuestoIVA8 Atributo condicional para expresar el total de los impuestos de IVA trasladado a la tasa del 8% que se desprenden de los pagos. No se permiten valores negativos.
	TotalTrasladosImpuestoIVA8 decimal.NullDecimal `xml:"TotalTrasladosImpuestoIVA8,attr,omitempty"`
	// TotalTrasladosBaseIVA0 Atributo condicional para expresar el total de la base de IVA trasladado a la tasa del 0% que se desprende de los pagos. No se permiten valores negativos.
	TotalTrasladosBaseIVA0 decimal.NullDecimal `xml:"TotalTrasladosBaseIVA0,attr,omitempty"`
	// TotalTrasladosImpuestoIVA0 Atributo condicional para expresar el total de los impuestos de IVA trasladado a la tasa del 0% que se desprenden de los pagos. No se permiten valores negativos.
	TotalTrasladosImpuestoIVA0 decimal.NullDecimal `xml:"TotalTrasladosImpuestoIVA0,attr,omitempty"`
	// TotalTrasladosBaseIVAExento Atributo condicional para expresar el total de la base de IVA trasladado exento que se desprende de los pagos. No se permiten valores negativos.
	TotalTrasladosBaseIVAExento decimal.NullDecimal `xml:"TotalTrasladosBaseIVAExento,attr,omitempty"`
	// MontoTotalPagos Atributo requerido para expresar el total de los pagos que se desprenden de los nodos Pago. No se permiten valores negativos.
	MontoTotalPagos decimal.Decimal `xml:"MontoTotalPagos,attr"`
}

// Pago Elemento requerido para incorporar la información de la recepción de pagos.
type Pago struct {
	// DoctoRelacionado Nodo requerido para expresar la lista de documentos relacionados con los pagos. Por cada documento que se relacione se debe generar un nodo DoctoRelacionado.
	DoctoRelacionado []*DoctoRelacionado `xml:"DoctoRelacionado"`
	// ImpuestosP Nodo condicional para registrar el resumen de los impuestos aplicables conforme al monto del pago recibido, expresados a la moneda de pago.
	ImpuestosP *ImpuestosP `xml:"ImpuestosP,omitempty"`
	// FechaPago Atributo requerido para expresar la fecha y hora en la que el beneficiario recibe el pago. Se expresa en la forma aaaa-mm-ddThh:mm:ss, de acuerdo con la especificación ISO 8601.En caso de no contar con la hora se debe registrar 12:00:00.
	FechaPago types.FechaH `xml:"FechaPago,attr"`
	// FormaDePagoP Atributo requerido para expresar la clave de la forma en que se realiza el pago.
	FormaDePagoP types.FormaPago `xml:"FormaDePagoP,attr"`
	// MonedaP Atributo requerido para identificar la clave de la moneda utilizada para realizar el pago conforme a la especificación ISO 4217. Cuando se usa moneda nacional se registra MXN. El atributo Pagos:Pago:Monto debe ser expresado en la moneda registrada en este atributo.
	MonedaP types.Moneda `xml:"MonedaP,attr"`
	// TipoCambioP Atributo condicional para expresar el tipo de cambio de la moneda a la fecha en que se realizó el pago. El valor debe reflejar el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo MonedaP. Es requerido cuando el atributo MonedaP es diferente a MXN.
	TipoCambioP decimal.Decimal `xml:"TipoCambioP,attr,omitempty"`
	// Monto Atributo requerido para expresar el importe del pago.
	Monto decimal.Decimal `xml:"Monto,attr"`
	// NumOperacion Atributo condicional para expresar el número de cheque, número de autorización, número de referencia, clave de rastreo en caso de ser SPEI, línea de captura o algún número de referencia análogo que identifique la operación que ampara el pago efectuado.
	NumOperacion string `xml:"NumOperacion,attr,omitempty"`
	// RfcEmisorCtaOrd Atributo condicional para expresar la clave RFC de la entidad emisora de la cuenta origen, es decir, la operadora, el banco, la institución financiera, emisor de monedero electrónico, etc., en caso de ser extranjero colocar XEXX010101000, considerar las reglas de obligatoriedad publicadas en la página del SAT para éste atributo de acuerdo con el catálogo catCFDI:c_FormaPago.
	RfcEmisorCtaOrd string `xml:"RfcEmisorCtaOrd,attr,omitempty"`
	// NomBancoOrdExt Atributo condicional para expresar el nombre del banco ordenante, es requerido en caso de ser extranjero. Considerar las reglas de obligatoriedad publicadas en la página del SAT para éste atributo de acuerdo con el catálogo catCFDI:c_FormaPago.
	NomBancoOrdExt string `xml:"NomBancoOrdExt,attr,omitempty"`
	// CtaOrdenante Atributo condicional para incorporar el número de la cuenta con la que se realizó el pago. Considerar las reglas de obligatoriedad publicadas en la página del SAT para éste atributo de acuerdo con el catálogo catCFDI:c_FormaPago.
	CtaOrdenante string `xml:"CtaOrdenante,attr,omitempty"`
	// RfcEmisorCtaBen Atributo condicional para expresar la clave RFC de la entidad operadora de la cuenta destino, es decir, la operadora, el banco, la institución financiera, emisor de monedero electrónico, etc. Considerar las reglas de obligatoriedad publicadas en la página del SAT para éste atributo de acuerdo con el catálogo catCFDI:c_FormaPago.
	RfcEmisorCtaBen string `xml:"RfcEmisorCtaBen,attr,omitempty"`
	// CtaBeneficiario Atributo condicional para incorporar el número de cuenta en donde se recibió el pago. Considerar las reglas de obligatoriedad publicadas en la página del SAT para éste atributo de acuerdo con el catálogo catCFDI:c_FormaPago.
	CtaBeneficiario string `xml:"CtaBeneficiario,attr,omitempty"`
	// TipoCadPago Atributo condicional para identificar la clave del tipo de cadena de pago que genera la entidad receptora del pago. Considerar las reglas de obligatoriedad publicadas en la página del SAT para éste atributo de acuerdo con el catálogo catCFDI:c_FormaPago.
	TipoCadPago CTipoCadenaPago `xml:"TipoCadPago,attr,omitempty"`
	// CertPago Atributo condicional que sirve para incorporar el certificado que ampara al pago, como una cadena de texto en formato base 64. Es requerido en caso de que el atributo TipoCadPago contenga información.
	CertPago string `xml:"CertPago,attr,omitempty"`
	// CadPago Atributo condicional para expresar la cadena original del comprobante de pago generado por la entidad emisora de la cuenta beneficiaria. Es requerido en caso de que el atributo TipoCadPago contenga información.
	CadPago string `xml:"CadPago,attr,omitempty"`
	// SelloPago Atributo condicional para integrar el sello digital que se asocie al pago. La entidad que emite el comprobante de pago, ingresa una cadena original y el sello digital en una sección de dicho comprobante, este sello digital es el que se debe registrar en este atributo. Debe ser expresado como una cadena de texto en formato base 64. Es requerido en caso de que el atributo TipoCadPago contenga información.
	SelloPago string `xml:"SelloPago,attr,omitempty"`
}

// May be one of 01
type CTipoCadenaPago string

const (
	CTipoCadenaPago01 CTipoCadenaPago = "01"
)

// DoctoRelacionado Nodo requerido para expresar la lista de documentos relacionados con los pagos. Por cada documento que se relacione se debe generar un nodo DoctoRelacionado.
type DoctoRelacionado struct {
	// ImpuestosDR Nodo condicional para registrar los impuestos aplicables conforme al monto del pago recibido, expresados a la moneda del documento relacionado.
	ImpuestosDR *ImpuestosDR `xml:"ImpuestosDR,omitempty"`
	// IdDocumento Atributo requerido para expresar el identificador del documento relacionado con el pago. Este dato puede ser un Folio Fiscal de la Factura Electrónica o bien el número de operación de un documento digital.
	IdDocumento string `xml:"IdDocumento,attr"`
	// Serie Atributo opcional para precisar la serie del comprobante para control interno del contribuyente, acepta una cadena de caracteres.
	Serie string `xml:"Serie,attr,omitempty"`
	// Folio Atributo opcional para precisar el folio del comprobante para control interno del contribuyente, acepta una cadena de caracteres.
	Folio string `xml:"Folio,attr,omitempty"`
	// MonedaDR Atributo requerido para identificar la clave de la moneda utilizada en los importes del documento relacionado, cuando se usa moneda nacional o el documento relacionado no especifica la moneda se registra MXN. Los importes registrados en los atributos “ImpSaldoAnt”, “ImpPagado” e “ImpSaldoInsoluto” de éste nodo, deben corresponder a esta moneda. Conforme con la especificación ISO 4217.
	MonedaDR types.Moneda `xml:"MonedaDR,attr"`
	// EquivalenciaDR Atributo condicional para expresar el tipo de cambio conforme con la moneda registrada en el documento relacionado. Es requerido cuando la moneda del documento relacionado es distinta de la moneda de pago. Se debe registrar el número de unidades de la moneda señalada en el documento relacionado que equivalen a una unidad de la moneda del pago. Por ejemplo: El documento relacionado se registra en USD. El pago se realiza por 100 EUR. Este atributo se registra como 1.114700 USD/EUR. El importe pagado equivale a 100 EUR * 1.114700 USD/EUR = 111.47 USD.
	EquivalenciaDR decimal.Decimal `xml:"EquivalenciaDR,attr,omitempty"`
	// NumParcialidad Atributo requerido para expresar el número de parcialidad que corresponde al pago.
	NumParcialidad int `xml:"NumParcialidad,attr"`
	// ImpSaldoAnt Atributo requerido para expresar el monto del saldo insoluto de la parcialidad anterior. En el caso de que sea la primer parcialidad este atributo debe contener el importe total del documento relacionado.
	ImpSaldoAnt decimal.Decimal `xml:"ImpSaldoAnt,attr"`
	// ImpPagado Atributo requerido para expresar el importe pagado para el documento relacionado.
	ImpPagado decimal.Decimal `xml:"ImpPagado,attr"`
	// ImpSaldoInsoluto Atributo requerido para expresar la diferencia entre el importe del saldo anterior y el monto del pago.
	ImpSaldoInsoluto decimal.Decimal `xml:"ImpSaldoInsoluto,attr"`
	// ObjetoImpDR Atributo requerido para expresar si el pago del documento relacionado es objeto o no de impuesto.
	ObjetoImpDR types.ObjetoImp `xml:"ObjetoImpDR,attr"`
}

// ImpuestosDR Nodo condicional para registrar los impuestos aplicables conforme al monto del pago recibido, expresados a la moneda del documento relacionado.
type ImpuestosDR struct {
	// RetencionesDR Nodo opcional para capturar los impuestos retenidos aplicables conforme al monto del pago recibido.
	RetencionesDR RetencionesDR `xml:"RetencionesDR,omitempty"`
	// TrasladosDR Nodo opcional para capturar los impuestos trasladados aplicables conforme al monto del pago recibido.
	TrasladosDR TrasladosDR `xml:"TrasladosDR,omitempty"`
}

// RetencionesDR Nodo opcional para capturar los impuestos retenidos aplicables conforme al monto del pago recibido.
type RetencionesDR []*RetencionDR

func (ret *RetencionesDR) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret2 struct {
		RetencionDR []*RetencionDR `xml:"RetencionDR"`
	}
	if err := d.DecodeElement(&ret2, &start); err != nil {
		return err
	}
	*ret = ret2.RetencionDR
	return nil
}

// RetencionDR Nodo requerido para registrar la información detallada de una retención de impuesto específico conforme al monto del pago recibido.
type RetencionDR struct {
	// BaseDR Atributo requerido para señalar la base para el cálculo de la retención conforme al monto del pago, aplicable al documento relacionado, la determinación de la base se realiza de acuerdo con las disposiciones fiscales vigentes. No se permiten valores negativos.
	BaseDR decimal.Decimal `xml:"BaseDR,attr"`
	// ImpuestoDR Atributo requerido para señalar la clave del tipo de impuesto retenido conforme al monto del pago, aplicable al documento relacionado.
	ImpuestoDR types.Impuesto `xml:"ImpuestoDR,attr"`
	// TipoFactorDR Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TipoFactorDR types.TipoFactor `xml:"TipoFactorDR,attr"`
	// TasaOCuotaDR Atributo requerido para señalar el valor de la tasa o cuota del impuesto que se retiene.
	TasaOCuotaDR decimal.Decimal `xml:"TasaOCuotaDR,attr"`
	// ImporteDR Atributo requerido para señalar el importe del impuesto retenido conforme al monto del pago, aplicable al documento relacionado. No se permiten valores negativos.
	ImporteDR decimal.Decimal `xml:"ImporteDR,attr"`
}

// TrasladosDR Nodo opcional para capturar los impuestos trasladados aplicables conforme al monto del pago recibido.
type TrasladosDR []*TrasladoDR

func (tras *TrasladosDR) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras2 struct {
		TrasladoDR []*TrasladoDR `xml:"TrasladoDR"`
	}
	if err := d.DecodeElement(&tras2, &start); err != nil {
		return err
	}
	*tras = tras2.TrasladoDR
	return nil
}

// TrasladoDR Nodo requerido para asentar la información detallada de un traslado de impuesto específico conforme al monto del pago recibido.
type TrasladoDR struct {
	// BaseDR Atributo requerido para señalar la base para el cálculo del impuesto trasladado conforme al monto del pago, aplicable al documento relacionado, la determinación de la base se realiza de acuerdo con las disposiciones fiscales vigentes. No se permiten valores negativos.
	BaseDR decimal.Decimal `xml:"BaseDR,attr"`
	// ImpuestoDR Atributo requerido para señalar la clave del tipo de impuesto trasladado conforme al monto del pago, aplicable al documento relacionado.
	ImpuestoDR types.Impuesto `xml:"ImpuestoDR,attr"`
	// TipoFactorDR Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TipoFactorDR types.TipoFactor `xml:"TipoFactorDR,attr"`
	// TasaOCuotaDR Atributo condicional para señalar el valor de la tasa o cuota del impuesto que se traslada. Es requerido cuando el atributo TipoFactorDR contenga una clave que corresponda a Tasa o Cuota.
	TasaOCuotaDR decimal.Decimal `xml:"TasaOCuotaDR,attr,omitempty"`
	// ImporteDR Atributo condicional para señalar el importe del impuesto trasladado conforme al monto del pago, aplicable al documento relacionado. No se permiten valores negativos. Es requerido cuando el tipo factor sea Tasa o Cuota.
	ImporteDR decimal.Decimal `xml:"ImporteDR,attr,omitempty"`
}

// ImpuestosP Nodo condicional para registrar el resumen de los impuestos aplicables conforme al monto del pago recibido, expresados a la moneda de pago.
type ImpuestosP struct {
	// RetencionesP Nodo condicional para señalar los impuestos retenidos aplicables conforme al monto del pago recibido. Es requerido cuando en los documentos relacionados se registre algún impuesto retenido.
	RetencionesP RetencionesP `xml:"RetencionesP,omitempty"`
	// TrasladosP Nodo condicional para capturar los impuestos trasladados aplicables conforme al monto del pago recibido. Es requerido cuando en los documentos relacionados se registre un impuesto trasladado.
	TrasladosP TrasladosP `xml:"TrasladosP,omitempty"`
}

// AddRetencion Suma el importe del retencion de acuerdo al tipo del impuesto.
func (i *ImpuestosP) AddRetencion(impuestoP types.Impuesto, importeP decimal.Decimal) {
	for _, ret := range i.RetencionesP {
		if ret.ImpuestoP == impuestoP {
			ret.ImporteP = ret.ImporteP.Add(importeP)
			return
		}
	}

	i.RetencionesP = append(i.RetencionesP, &RetencionP{
		ImpuestoP: impuestoP,
		ImporteP:  importeP,
	})
	return
}

// AddTraslado Suma el importe del traslado de acuerdo al tipo del impuesto.
func (i *ImpuestosP) AddTraslado(
	baseP decimal.Decimal,
	impuestoP types.Impuesto,
	tipoFactorP types.TipoFactor,
	tasaOCuotaP decimal.Decimal,
	importeP decimal.Decimal,
) {
	for _, tras := range i.TrasladosP {
		if tras.ImpuestoP == impuestoP && tras.TipoFactorP == tipoFactorP && tras.TasaOCuotaP.Equal(tasaOCuotaP) {
			tras.BaseP = tras.BaseP.Add(baseP)
			tras.ImporteP = tras.ImporteP.Add(importeP)
			return
		}
	}

	i.TrasladosP = append(i.TrasladosP, &TrasladoP{
		BaseP:       baseP,
		ImpuestoP:   impuestoP,
		TipoFactorP: tipoFactorP,
		TasaOCuotaP: tasaOCuotaP,
		ImporteP:    importeP,
	})
}

// RetencionesP Nodo condicional para señalar los impuestos retenidos aplicables conforme al monto del pago recibido. Es requerido cuando en los documentos relacionados se registre algún impuesto retenido.
type RetencionesP []*RetencionP

func (ret *RetencionesP) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret2 struct {
		RetencionP []*RetencionP `xml:"RetencionP"`
	}
	if err := d.DecodeElement(&ret2, &start); err != nil {
		return err
	}
	*ret = ret2.RetencionP
	return nil
}

// RetencionP Nodo requerido para señalar la información detallada de una retención de impuesto específico conforme al monto del pago recibido.
type RetencionP struct {
	// ImpuestoP Atributo requerido para señalar la clave del tipo de impuesto retenido conforme al monto del pago.
	ImpuestoP types.Impuesto `xml:"ImpuestoP,attr"`
	// ImporteP Atributo requerido para señalar el importe del impuesto retenido conforme al monto del pago. No se permiten valores negativos.
	ImporteP decimal.Decimal `xml:"ImporteP,attr"`
}

// TrasladosP Nodo condicional para capturar los impuestos trasladados aplicables conforme al monto del pago recibido. Es requerido cuando en los documentos relacionados se registre un impuesto trasladado.
type TrasladosP []*TrasladoP

func (tras *TrasladosP) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras2 struct {
		TrasladoP []*TrasladoP `xml:"TrasladoP"`
	}
	if err := d.DecodeElement(&tras2, &start); err != nil {
		return err
	}
	*tras = tras2.TrasladoP
	return nil
}

// TrasladoP Nodo requerido para señalar la información detallada de un traslado de impuesto específico conforme al monto del pago recibido.
type TrasladoP struct {
	// BaseP Atributo requerido para señalar la suma de los atributos BaseDR de los documentos relacionados del impuesto trasladado. No se permiten valores negativos.
	BaseP decimal.Decimal `xml:"BaseP,attr"`
	// ImpuestoP Atributo requerido para señalar la clave del tipo de impuesto trasladado conforme al monto del pago.
	ImpuestoP types.Impuesto `xml:"ImpuestoP,attr"`
	// TipoFactorP Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TipoFactorP types.TipoFactor `xml:"TipoFactorP,attr"`
	// TasaOCuotaP Atributo condicional para señalar el valor de la tasa o cuota del impuesto que se traslada en los documentos relacionados.
	TasaOCuotaP decimal.Decimal `xml:"TasaOCuotaP,attr,omitempty"`
	// ImporteP Atributo condicional para señalar la suma del impuesto trasladado, agrupado por ImpuestoP, TipoFactorP y TasaOCuotaP. No se permiten valores negativos.
	ImporteP decimal.Decimal `xml:"ImporteP,attr,omitempty"`
}
