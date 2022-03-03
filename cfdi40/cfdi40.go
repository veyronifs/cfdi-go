package cfdi40

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/complemento/cartaporte20"
	"github.com/veyronifs/cfdi-go/complemento/comext11"
	"github.com/veyronifs/cfdi-go/complemento/pagos20"
	"github.com/veyronifs/cfdi-go/complemento/tfd11"
	"github.com/veyronifs/cfdi-go/types"
)

const (
	RFCPublico    = "XAXX010101000"
	RFCExtranjero = "XEXX010101000"

	Version = "4.0"
)

// Comprobante Estándar de Comprobante Fiscal Digital por Internet.
type Comprobante struct {
	// InformacionGlobal Nodo condicional para precisar la información relacionada con el comprobante global.
	InformacionGlobal *InformacionGlobal `xml:"InformacionGlobal,omitempty"`
	// CfdiRelacionados Nodo opcional para precisar la información de los comprobantes relacionados.
	CfdiRelacionados []*CfdiRelacionados `xml:"CfdiRelacionados,omitempty"`
	// Emisor Nodo requerido para expresar la información del contribuyente emisor del comprobante.
	Emisor *Emisor `xml:"Emisor"`
	// Receptor Nodo requerido para precisar la información del contribuyente receptor del comprobante.
	Receptor *Receptor `xml:"Receptor"`
	// Conceptos Nodo requerido para listar los conceptos cubiertos por el comprobante.
	Conceptos Conceptos `xml:"Conceptos"`
	// Impuestos Nodo condicional para expresar el resumen de los impuestos aplicables.
	Impuestos *Impuestos `xml:"Impuestos,omitempty"`
	// Complemento Nodo opcional donde se incluye el complemento Timbre Fiscal Digital de manera obligatoria y los nodos complementarios determinados por el SAT, de acuerdo con las disposiciones particulares para un sector o actividad específica.
	Complemento *Complemento `xml:"Complemento,omitempty"`
	// Addenda Nodo opcional para recibir las extensiones al presente formato que sean de utilidad al contribuyente. Para las reglas de uso del mismo, referirse al formato origen.
	Addenda *Addenda `xml:"Addenda,omitempty"`
	// Version Atributo requerido con valor prefijado a 4.0 que indica la versión del estándar bajo el que se encuentra expresado el comprobante.
	Version string `xml:"Version,attr"`
	// Serie Atributo opcional para precisar la serie para control interno del contribuyente. Este atributo acepta una cadena de caracteres.
	Serie string `xml:"Serie,attr,omitempty"`
	// Folio Atributo opcional para control interno del contribuyente que expresa el folio del comprobante, acepta una cadena de caracteres.
	Folio string `xml:"Folio,attr,omitempty"`
	// Fecha Atributo requerido para la expresión de la fecha y hora de expedición del Comprobante Fiscal Digital por Internet. Se expresa en la forma AAAA-MM-DDThh:mm:ss y debe corresponder con la hora local donde se expide el comprobante.
	Fecha types.FechaH `xml:"Fecha,attr"`
	// Sello Atributo requerido para contener el sello digital del comprobante fiscal, al que hacen referencia las reglas de resolución miscelánea vigente. El sello debe ser expresado como una cadena de texto en formato Base 64.
	Sello string `xml:"Sello,attr"`
	// FormaPago Atributo condicional para expresar la clave de la forma de pago de los bienes o servicios amparados por el comprobante.
	FormaPago types.FormaPago `xml:"FormaPago,attr,omitempty"`
	// NoCertificado Atributo requerido para expresar el número de serie del certificado de sello digital que ampara al comprobante, de acuerdo con el acuse correspondiente a 20 posiciones otorgado por el sistema del SAT.
	NoCertificado string `xml:"NoCertificado,attr"`
	// Certificado Atributo requerido que sirve para incorporar el certificado de sello digital que ampara al comprobante, como texto en formato base 64.
	Certificado string `xml:"Certificado,attr"`
	// CondicionesDePago Atributo condicional para expresar las condiciones comerciales aplicables para el pago del comprobante fiscal digital por Internet. Este atributo puede ser condicionado mediante atributos o complementos.
	CondicionesDePago string `xml:"CondicionesDePago,attr,omitempty"`
	// SubTotal Atributo requerido para representar la suma de los importes de los conceptos antes de descuentos e impuesto. No se permiten valores negativos.
	SubTotal decimal.Decimal `xml:"SubTotal,attr"`
	// Descuento Atributo condicional para representar el importe total de los descuentos aplicables antes de impuestos. No se permiten valores negativos. Se debe registrar cuando existan conceptos con descuento.
	Descuento decimal.Decimal `xml:"Descuento,attr,omitempty"`
	// Moneda Atributo requerido para identificar la clave de la moneda utilizada para expresar los montos, cuando se usa moneda nacional se registra MXN. Conforme con la especificación ISO 4217.
	Moneda types.Moneda `xml:"Moneda,attr"`
	// TipoCambio Atributo condicional para representar el tipo de cambio FIX conforme con la moneda usada. Es requerido cuando la clave de moneda es distinta de MXN y de XXX. El valor debe reflejar el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo moneda. Si el valor está fuera del porcentaje aplicable a la moneda tomado del catálogo c_Moneda, el emisor debe obtener del PAC que vaya a timbrar el CFDI, de manera no automática, una clave de confirmación para ratificar que el valor es correcto e integrar dicha clave en el atributo Confirmacion.
	TipoCambio decimal.Decimal `xml:"TipoCambio,attr,omitempty"`
	// Total Atributo requerido para representar la suma del subtotal, menos los descuentos aplicables, más las contribuciones recibidas (impuestos trasladados - federales y/o locales, derechos, productos, aprovechamientos, aportaciones de seguridad social, contribuciones de mejoras) menos los impuestos retenidos federales y/o locales. Si el valor es superior al límite que establezca el SAT en la Resolución Miscelánea Fiscal vigente, el emisor debe obtener del PAC que vaya a timbrar el CFDI, de manera no automática, una clave de confirmación para ratificar que el valor es correcto e integrar dicha clave en el atributo Confirmacion. No se permiten valores negativos.
	Total decimal.Decimal `xml:"Total,attr"`
	// TipoDeComprobante Atributo requerido para expresar la clave del efecto del comprobante fiscal para el contribuyente emisor.
	TipoDeComprobante types.TipoDeComprobante `xml:"TipoDeComprobante,attr"`
	// Exportacion Atributo requerido para expresar si el comprobante ampara una operación de exportación.
	Exportacion types.Exportacion `xml:"Exportacion,attr"`
	// MetodoPago Atributo condicional para precisar la clave del método de pago que aplica para este comprobante fiscal digital por Internet, conforme al Artículo 29-A fracción VII incisos a y b del CFF.
	MetodoPago types.MetodoPago `xml:"MetodoPago,attr,omitempty"`
	// LugarExpedicion Atributo requerido para incorporar el código postal del lugar de expedición del comprobante (domicilio de la matriz o de la sucursal).
	LugarExpedicion string `xml:"LugarExpedicion,attr"`
	// Confirmacion Atributo condicional para registrar la clave de confirmación que entregue el PAC para expedir el comprobante con importes grandes, con un tipo de cambio fuera del rango establecido o con ambos casos. Es requerido cuando se registra un tipo de cambio o un total fuera del rango establecido.
	Confirmacion string `xml:"Confirmacion,attr,omitempty"`
}

var ErrNoTimbrado = errors.New("comprobante no timbrado")

func (c Comprobante) QRText() (string, error) {
	if c.Complemento.TFD11 == nil {
		return "", ErrNoTimbrado
	}

	idxSubstr := len(c.Complemento.TFD11.SelloCFD) - 8
	if idxSubstr < 0 {
		return "", fmt.Errorf("%w selloCFD invalido", ErrNoTimbrado)
	}
	return "https://verificacfdi.facturaelectronica.sat.gob.mx/default.aspx?" +
		"&id=" + c.Complemento.TFD11.UUID +
		"&re=" + c.Emisor.Rfc +
		"&rr=" + c.Receptor.Rfc +
		"&tt=" + c.Total.String() +
		"&fe=" + c.Complemento.TFD11.SelloCFD[idxSubstr:], nil
}

// InformacionGlobal Nodo condicional para precisar la información relacionada con el comprobante global.
type InformacionGlobal struct {
	// Periodicidad Atributo requerido para expresar el período al que corresponde la información del comprobante global.
	Periodicidad types.Periodicidad `xml:"Periodicidad,attr"`
	// Meses Atributo requerido para expresar el mes o los meses al que corresponde la información del comprobante global.
	Meses string `xml:"Meses,attr"`
	// Anio Atributo requerido para expresar el año al que corresponde la información del comprobante global.
	Anio int `xml:"Año,attr"`
}

// CfdiRelacionados Nodo opcional para precisar la información de los comprobantes relacionados.
type CfdiRelacionados struct {
	// CfdiRelacionado Nodo requerido para precisar la información de los comprobantes relacionados.
	CfdiRelacionado []*CfdiRelacionado `xml:"http://www.sat.gob.mx/cfd/4 CfdiRelacionado"`
	// TipoRelacion Atributo requerido para indicar la clave de la relación que existe entre éste que se está generando y el o los CFDI previos.
	TipoRelacion types.TipoRelacion `xml:"TipoRelacion,attr"`
}

// CfdiRelacionado Nodo requerido para precisar la información de los comprobantes relacionados.
type CfdiRelacionado struct {
	// UUID Atributo requerido para registrar el folio fiscal (UUID) de un CFDI relacionado con el presente comprobante, por ejemplo: Si el CFDI relacionado es un comprobante de traslado que sirve para registrar el movimiento de la mercancía. Si este comprobante se usa como nota de crédito o nota de débito del comprobante relacionado. Si este comprobante es una devolución sobre el comprobante relacionado. Si éste sustituye a una factura cancelada.
	UUID string `xml:"UUID,attr"`
}

// Emisor Nodo requerido para expresar la información del contribuyente emisor del comprobante.
type Emisor struct {
	// Rfc Atributo requerido para registrar la Clave del Registro Federal de Contribuyentes correspondiente al contribuyente emisor del comprobante.
	Rfc string `xml:"Rfc,attr"`
	// Nombre Atributo requerido para registrar el nombre, denominación o razón social del contribuyente inscrito en el RFC, del emisor del comprobante.
	Nombre string `xml:"Nombre,attr"`
	// RegimenFiscal Atributo requerido para incorporar la clave del régimen del contribuyente emisor al que aplicará el efecto fiscal de este comprobante.
	RegimenFiscal types.RegimenFiscal `xml:"RegimenFiscal,attr"`
	// FacAtrAdquirente Atributo condicional para expresar el número de operación proporcionado por el SAT cuando se trate de un comprobante a través de un PCECFDI o un PCGCFDISP.
	FacAtrAdquirente string `xml:"FacAtrAdquirente,attr,omitempty"`
}

// Receptor Nodo requerido para precisar la información del contribuyente receptor del comprobante.
type Receptor struct {
	// Rfc Atributo requerido para registrar la Clave del Registro Federal de Contribuyentes correspondiente al contribuyente receptor del comprobante.
	Rfc string `xml:"Rfc,attr"`
	// Nombre Atributo requerido para registrar el nombre(s), primer apellido, segundo apellido, según corresponda, denominación o razón social del contribuyente, inscrito en el RFC, del receptor del comprobante.
	Nombre string `xml:"Nombre,attr"`
	// DomicilioFiscalReceptor Atributo requerido para registrar el código postal del domicilio fiscal del receptor del comprobante.
	DomicilioFiscalReceptor string `xml:"DomicilioFiscalReceptor,attr"`
	// ResidenciaFiscal Atributo condicional para registrar la clave del país de residencia para efectos fiscales del receptor del comprobante, cuando se trate de un extranjero, y que es conforme con la especificación ISO 3166-1 alpha-3. Es requerido cuando se incluya el complemento de comercio exterior o se registre el atributo NumRegIdTrib.
	ResidenciaFiscal types.Pais `xml:"ResidenciaFiscal,attr,omitempty"`
	// NumRegIdTrib Atributo condicional para expresar el número de registro de identidad fiscal del receptor cuando sea residente en el extranjero. Es requerido cuando se incluya el complemento de comercio exterior.
	NumRegIdTrib string `xml:"NumRegIdTrib,attr,omitempty"`
	// RegimenFiscalReceptor Atributo requerido para incorporar la clave del régimen fiscal del contribuyente receptor al que aplicará el efecto fiscal de este comprobante.
	RegimenFiscalReceptor types.RegimenFiscal `xml:"RegimenFiscalReceptor,attr"`
	// UsoCFDI Atributo requerido para expresar la clave del uso que dará a esta factura el receptor del CFDI.
	UsoCFDI types.UsoCFDI `xml:"UsoCFDI,attr"`
}

// Nodo requerido para listar los conceptos cubiertos por el comprobante.
type Conceptos []*Concepto

func (c *Conceptos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var c1 struct {
		Conceptos []*Concepto `xml:"http://www.sat.gob.mx/cfd/4 Concepto"`
	}
	if err := d.DecodeElement(&c1, &start); err != nil {
		return err
	}
	*c = c1.Conceptos
	return nil
}

// Concepto Nodo requerido para registrar la información detallada de un bien o servicio amparado en el comprobante.
type Concepto struct {
	// Impuestos Nodo condicional para capturar los impuestos aplicables al presente concepto.
	Impuestos *ConceptoImpuestos `xml:"http://www.sat.gob.mx/cfd/4 Impuestos,omitempty"`
	// ACuentaTerceros Nodo opcional para registrar información del contribuyente Tercero, a cuenta del que se realiza la operación.
	ACuentaTerceros *ConceptoACuentaTerceros `xml:"http://www.sat.gob.mx/cfd/4 ACuentaTerceros,omitempty"`
	// InformacionAduanera Nodo opcional para introducir la información aduanera aplicable cuando se trate de ventas de primera mano de mercancías importadas o se trate de operaciones de comercio exterior con bienes o servicios.
	InformacionAduanera []*ConceptoInformacionAduanera `xml:"http://www.sat.gob.mx/cfd/4 InformacionAduanera,omitempty"`
	// CuentaPredial Nodo opcional para asentar el número de cuenta predial con el que fue registrado el inmueble, en el sistema catastral de la entidad federativa de que trate, o bien para incorporar los datos de identificación del certificado de participación inmobiliaria no amortizable.
	CuentaPredial []*ConceptoCuentaPredial `xml:"http://www.sat.gob.mx/cfd/4 CuentaPredial,omitempty"`
	// ComplementoConcepto Nodo opcional donde se incluyen los nodos complementarios de extensión al concepto definidos por el SAT, de acuerdo con las disposiciones particulares para un sector o actividad específica.
	ComplementoConcepto Addenda `xml:"http://www.sat.gob.mx/cfd/4 ComplementoConcepto,omitempty"`
	// Parte Nodo opcional para expresar las partes o componentes que integran la totalidad del concepto expresado en el comprobante fiscal digital por Internet.
	Parte []*Parte `xml:"http://www.sat.gob.mx/cfd/4 Parte,omitempty"`
	// ClaveProdServ Atributo requerido para expresar la clave del producto o del servicio amparado por el presente concepto. Es requerido y deben utilizar las claves del catálogo de productos y servicios, cuando los conceptos que registren por sus actividades correspondan con dichos conceptos.
	ClaveProdServ string `xml:"ClaveProdServ,attr"`
	// NoIdentificacion Atributo opcional para expresar el número de parte, identificador del producto o del servicio, la clave de producto o servicio, SKU o equivalente, propia de la operación del emisor, amparado por el presente concepto. Opcionalmente se puede utilizar claves del estándar GTIN.
	NoIdentificacion string `xml:"NoIdentificacion,attr,omitempty"`
	// Cantidad Atributo requerido para precisar la cantidad de bienes o servicios del tipo particular definido por el presente concepto.
	Cantidad decimal.Decimal `xml:"Cantidad,attr"`
	// ClaveUnidad Atributo requerido para precisar la clave de unidad de medida estandarizada aplicable para la cantidad expresada en el concepto. La unidad debe corresponder con la descripción del concepto.
	ClaveUnidad string `xml:"ClaveUnidad,attr"`
	// Unidad Atributo opcional para precisar la unidad de medida propia de la operación del emisor, aplicable para la cantidad expresada en el concepto. La unidad debe corresponder con la descripción del concepto.
	Unidad string `xml:"Unidad,attr,omitempty"`
	// Descripcion Atributo requerido para precisar la descripción del bien o servicio cubierto por el presente concepto.
	Descripcion string `xml:"Descripcion,attr"`
	// ValorUnitario Atributo requerido para precisar el valor o precio unitario del bien o servicio cubierto por el presente concepto.
	ValorUnitario decimal.Decimal `xml:"ValorUnitario,attr"`
	// Importe Atributo requerido para precisar el importe total de los bienes o servicios del presente concepto. Debe ser equivalente al resultado de multiplicar la cantidad por el valor unitario expresado en el concepto. No se permiten valores negativos.
	Importe decimal.Decimal `xml:"Importe,attr"`
	// Descuento Atributo opcional para representar el importe de los descuentos aplicables al concepto. No se permiten valores negativos.
	Descuento decimal.Decimal `xml:"Descuento,attr,omitempty"`
	// ObjetoImp Atributo requerido para expresar si la operación comercial es objeto o no de impuesto.
	ObjetoImp types.ObjetoImp `xml:"ObjetoImp,attr"`
}

// Total calcula el total de los conceptos.
//
//	  Subtotal
//	- Descuentos
//	- ImpuestosTrasladados
//	+ ImpuestosTetenidos
func (c Concepto) Total() decimal.Decimal {
	total := c.Importe.Sub(c.Descuento)
	if c.Impuestos != nil {
		for _, tras := range c.Impuestos.Traslados {
			total = total.Add(tras.Importe)
		}
		for _, ret := range c.Impuestos.Retenciones {
			total = total.Sub(ret.Importe)
		}
	}
	return total
}

// ConceptoImpuestos Nodo condicional para capturar los impuestos aplicables al presente concepto.
type ConceptoImpuestos struct {
	// Traslados Nodo opcional para asentar los impuestos trasladados aplicables al presente concepto.
	Traslados ConceptoImpuestosTraslados `xml:"http://www.sat.gob.mx/cfd/4 Traslados,omitempty"`
	// Retenciones Nodo opcional para asentar los impuestos retenidos aplicables al presente concepto.
	Retenciones ConceptoImpuestosRetenciones `xml:"http://www.sat.gob.mx/cfd/4 Retenciones,omitempty"`
}

// ConceptoImpuestosTraslados Nodo opcional para asentar los impuestos trasladados aplicables al presente concepto.
type ConceptoImpuestosTraslados []*ConceptoImpuestosTraslado

func (tras *ConceptoImpuestosTraslados) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras1 struct {
		Traslados []*ConceptoImpuestosTraslado `xml:"http://www.sat.gob.mx/cfd/4 Traslado"`
	}

	if err := d.DecodeElement(&tras1, &start); err != nil {
		return err
	}
	*tras = tras1.Traslados
	return nil
}

// ConceptoImpuestosTraslado Nodo requerido para asentar la información detallada de un traslado de impuestos aplicable al presente concepto.
type ConceptoImpuestosTraslado struct {
	// Atributo requerido para señalar la base para el cálculo del impuesto, la determinación de la base se realiza de acuerdo con las disposiciones fiscales vigentes. No se permiten valores negativos.
	Base decimal.Decimal `xml:"Base,attr"`
	// Atributo requerido para señalar la clave del tipo de impuesto trasladado aplicable al concepto.
	Impuesto types.Impuesto `xml:"Impuesto,attr"`
	// Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TipoFactor types.TipoFactor `xml:"TipoFactor,attr"`
	// Atributo condicional para señalar el valor de la tasa o cuota del impuesto que se traslada para el presente concepto. Es requerido cuando el atributo TipoFactor tenga una clave que corresponda a Tasa o Cuota.
	TasaOCuota decimal.Decimal `xml:"TasaOCuota,attr,omitempty"`
	// Atributo condicional para señalar el importe del impuesto trasladado que aplica al concepto. No se permiten valores negativos. Es requerido cuando TipoFactor sea Tasa o Cuota.
	Importe decimal.Decimal `xml:"Importe,attr,omitempty"`
}

// ConceptoImpuestosRetenciones Nodo opcional para asentar los impuestos retenidos aplicables al presente concepto.
type ConceptoImpuestosRetenciones []*ConceptoImpuestosRetencion

func (ret *ConceptoImpuestosRetenciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret1 struct {
		Retenciones []*ConceptoImpuestosRetencion `xml:"http://www.sat.gob.mx/cfd/4 Retencion"`
	}

	if err := d.DecodeElement(&ret1, &start); err != nil {
		return err
	}
	*ret = ret1.Retenciones
	return nil
}

// ConceptoImpuestosRetencion Nodo requerido para asentar la información detallada de una retención de impuestos aplicable al presente concepto.
type ConceptoImpuestosRetencion struct {
	// Base Atributo requerido para señalar la base para el cálculo de la retención, la determinación de la base se realiza de acuerdo con las disposiciones fiscales vigentes. No se permiten valores negativos.
	Base decimal.Decimal `xml:"Base,attr"`
	// Impuesto Atributo requerido para señalar la clave del tipo de impuesto retenido aplicable al concepto.
	Impuesto types.Impuesto `xml:"Impuesto,attr"`
	// TipoFactor Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TipoFactor types.TipoFactor `xml:"TipoFactor,attr"`
	// TasaOCuota Atributo requerido para señalar la tasa o cuota del impuesto que se retiene para el presente concepto.
	TasaOCuota decimal.Decimal `xml:"TasaOCuota,attr"`
	// Importe Atributo requerido para señalar el importe del impuesto retenido que aplica al concepto. No se permiten valores negativos.
	Importe decimal.Decimal `xml:"Importe,attr"`
}

// ConceptoACuentaTerceros Nodo opcional para registrar información del contribuyente Tercero, a cuenta del que se realiza la operación.
type ConceptoACuentaTerceros struct {
	// RfcACuentaTerceros Atributo requerido para registrar la Clave del Registro Federal de Contribuyentes del contribuyente Tercero, a cuenta del que se realiza la operación.
	RfcACuentaTerceros string `xml:"RfcACuentaTerceros,attr"`
	// NombreACuentaTerceros Atributo requerido para registrar el nombre, denominación o razón social del contribuyente Tercero correspondiente con el Rfc, a cuenta del que se realiza la operación.
	NombreACuentaTerceros string `xml:"NombreACuentaTerceros,attr"`
	// RegimenFiscalACuentaTerceros Atributo requerido para incorporar la clave del régimen del contribuyente Tercero, a cuenta del que se realiza la operación.
	RegimenFiscalACuentaTerceros types.RegimenFiscal `xml:"RegimenFiscalACuentaTerceros,attr"`
	// DomicilioFiscalACuentaTerceros Atributo requerido para incorporar el código postal del domicilio fiscal del Tercero, a cuenta del que se realiza la operación.
	DomicilioFiscalACuentaTerceros string `xml:"DomicilioFiscalACuentaTerceros,attr"`
}

// ConceptoInformacionAduanera Nodo opcional para introducir la información aduanera aplicable cuando se trate de ventas de primera mano de mercancías importadas o se trate de operaciones de comercio exterior con bienes o servicios.
type ConceptoInformacionAduanera struct {
	// NumeroPedimento Atributo requerido para expresar el número del pedimento que ampara la importación del bien que se expresa en el siguiente formato: últimos 2 dígitos del año de validación seguidos por dos espacios, 2 dígitos de la aduana de despacho seguidos por dos espacios, 4 dígitos del número de la patente seguidos por dos espacios, 1 dígito que corresponde al último dígito del año en curso, salvo que se trate de un pedimento consolidado iniciado en el año inmediato anterior o del pedimento original de una rectificación, seguido de 6 dígitos de la numeración progresiva por aduana.
	NumeroPedimento string `xml:"NumeroPedimento,attr"`
}

// ConceptoCuentaPredial Nodo opcional para asentar el número de cuenta predial con el que fue registrado el inmueble, en el sistema catastral de la entidad federativa de que trate, o bien para incorporar los datos de identificación del certificado de participación inmobiliaria no amortizable.
type ConceptoCuentaPredial struct {
	// Numero Atributo requerido para precisar el número de la cuenta predial del inmueble cubierto por el presente concepto, o bien para incorporar los datos de identificación del certificado de participación inmobiliaria no amortizable, tratándose de arrendamiento.
	Numero string `xml:"Numero,attr"`
}

// Parte Nodo opcional para expresar las partes o componentes que integran la totalidad del concepto expresado en el comprobante fiscal digital por Internet.
type Parte struct {
	// InformacionAduanera Nodo opcional para introducir la información aduanera aplicable cuando se trate de ventas de primera mano de mercancías importadas o se trate de operaciones de comercio exterior con bienes o servicios.
	InformacionAduanera []*ConceptoInformacionAduanera `xml:"http://www.sat.gob.mx/cfd/4 InformacionAduanera,omitempty"`
	// ClaveProdServ Atributo requerido para expresar la clave del producto o del servicio amparado por la presente parte. Es requerido y deben utilizar las claves del catálogo de productos y servicios, cuando los conceptos que registren por sus actividades correspondan con dichos conceptos.
	ClaveProdServ string `xml:"ClaveProdServ,attr"`
	// NoIdentificacion Atributo opcional para expresar el número de serie, número de parte del bien o identificador del producto o del servicio amparado por la presente parte. Opcionalmente se puede utilizar claves del estándar GTIN.
	NoIdentificacion string `xml:"NoIdentificacion,attr,omitempty"`
	// Cantidad Atributo requerido para precisar la cantidad de bienes o servicios del tipo particular definido por la presente parte.
	Cantidad decimal.Decimal `xml:"Cantidad,attr"`
	// Unidad Atributo opcional para precisar la unidad de medida propia de la operación del emisor, aplicable para la cantidad expresada en la parte. La unidad debe corresponder con la descripción de la parte.
	Unidad string `xml:"Unidad,attr,omitempty"`
	// Descripcion Atributo requerido para precisar la descripción del bien o servicio cubierto por la presente parte.
	Descripcion string `xml:"Descripcion,attr"`
	// ValorUnitario Atributo opcional para precisar el valor o precio unitario del bien o servicio cubierto por la presente parte. No se permiten valores negativos.
	ValorUnitario decimal.Decimal `xml:"ValorUnitario,attr,omitempty"`
	// Importe Atributo opcional para precisar el importe total de los bienes o servicios de la presente parte. Debe ser equivalente al resultado de multiplicar la cantidad por el valor unitario expresado en la parte. No se permiten valores negativos.
	Importe decimal.Decimal `xml:"Importe,attr,omitempty"`
}

// Impuestos Nodo condicional para expresar el resumen de los impuestos aplicables.
type Impuestos struct {
	// Retenciones Nodo condicional para capturar los impuestos retenidos aplicables. Es requerido cuando en los conceptos se registre algún impuesto retenido.
	Retenciones ImpuestosRetenciones `xml:"http://www.sat.gob.mx/cfd/4 Retenciones,omitempty"`
	// Traslados Nodo condicional para capturar los impuestos trasladados aplicables. Es requerido cuando en los conceptos se registre un impuesto trasladado.
	Traslados ImpuestosTraslados `xml:"http://www.sat.gob.mx/cfd/4 Traslados,omitempty"`
	// TotalImpuestosRetenidos Atributo condicional para expresar el total de los impuestos retenidos que se desprenden de los conceptos expresados en el comprobante fiscal digital por Internet. No se permiten valores negativos. Es requerido cuando en los conceptos se registren impuestos retenidos.
	TotalImpuestosRetenidos decimal.Decimal `xml:"TotalImpuestosRetenidos,attr,omitempty"`
	// TotalImpuestosTrasladados Atributo condicional para expresar el total de los impuestos trasladados que se desprenden de los conceptos expresados en el comprobante fiscal digital por Internet. No se permiten valores negativos. Es requerido cuando en los conceptos se registren impuestos trasladados.
	TotalImpuestosTrasladados decimal.Decimal `xml:"TotalImpuestosTrasladados,attr,omitempty"`
}

// GetRetencion regresa la primera retención que encuentre de acuerdo a la clave del impuesto.
func (i *Impuestos) GetRetencion(tipo types.Impuesto) *ImpuestosRetencion {
	for _, ret := range i.Retenciones {
		if ret.Impuesto == tipo {
			return ret
		}
	}
	return nil
}

// AddRetencion suma el impuesto de retención al total de impuestos retenidos y al arreglo de retenciones.
func (i *Impuestos) AddRetencion(tipo types.Impuesto, importe decimal.Decimal) {
	ret := i.GetRetencion(tipo)
	i.TotalImpuestosRetenidos = i.TotalImpuestosRetenidos.Add(importe)
	if ret == nil {
		i.Retenciones = append(i.Retenciones, &ImpuestosRetencion{
			Impuesto: tipo,
			Importe:  importe,
		})
		return
	}
	ret.Importe = ret.Importe.Add(importe)
}

// GetTraslado regresa el primer traslado que encuentre de acuerdo a la clave del impuesto, tipo de factor y tasa del impuesto.
func (i *Impuestos) GetTraslado(tipo types.Impuesto, factor types.TipoFactor, tasa decimal.Decimal) *ImpuestosTraslado {
	for _, tras := range i.Traslados {
		if tras.Impuesto == tipo && tras.TipoFactor == factor && tras.TasaOCuota.Equal(tasa) {
			return tras
		}
	}
	return nil
}

// AddTraslado suma el impuesto de traslado al total de impuestos trasladados y al arreglo de traslados.
func (i *Impuestos) AddTraslado(
	base decimal.Decimal,
	tipo types.Impuesto,
	factor types.TipoFactor,
	tasa decimal.Decimal,
	importe decimal.Decimal,
) {
	tras := i.GetTraslado(tipo, factor, tasa)
	i.TotalImpuestosTrasladados = i.TotalImpuestosTrasladados.Add(importe)
	if tras == nil {
		i.Traslados = append(i.Traslados, &ImpuestosTraslado{
			Base:       base,
			Impuesto:   tipo,
			TipoFactor: factor,
			TasaOCuota: tasa,
			Importe:    importe,
		})
		return
	}
	tras.Importe = tras.Importe.Add(importe)
	tras.Base = tras.Base.Add(base)
}

// ImpuestosRetenciones Nodo condicional para capturar los impuestos retenidos aplicables. Es requerido cuando en los conceptos se registre algún impuesto retenido.
type ImpuestosRetenciones []*ImpuestosRetencion

func (ret *ImpuestosRetenciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret1 struct {
		Retenciones []*ImpuestosRetencion `xml:"http://www.sat.gob.mx/cfd/4 Retencion"`
	}

	if err := d.DecodeElement(&ret1, &start); err != nil {
		return err
	}
	*ret = ret1.Retenciones
	return nil
}

// ImpuestosRetencion Nodo requerido para la información detallada de una retención de impuesto específico.
type ImpuestosRetencion struct {
	// Impuesto Atributo requerido para señalar la clave del tipo de impuesto retenido.
	Impuesto types.Impuesto `xml:"Impuesto,attr"`
	// Importe Atributo requerido para señalar el monto del impuesto retenido. No se permiten valores negativos.
	Importe decimal.Decimal `xml:"Importe,attr"`
}

// ImpuestosTraslados Nodo condicional para capturar los impuestos trasladados aplicables. Es requerido cuando en los conceptos se registre un impuesto trasladado.
type ImpuestosTraslados []*ImpuestosTraslado

func (tras *ImpuestosTraslados) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras1 struct {
		Traslados []*ImpuestosTraslado `xml:"http://www.sat.gob.mx/cfd/4 Traslado"`
	}

	if err := d.DecodeElement(&tras1, &start); err != nil {
		return err
	}
	*tras = tras1.Traslados
	return nil
}

// ImpuestosTraslado Nodo requerido para la información detallada de un traslado de impuesto específico.
type ImpuestosTraslado struct {
	// Base Atributo requerido para señalar la suma de los atributos Base de los conceptos del impuesto trasladado. No se permiten valores negativos.
	Base decimal.Decimal `xml:"Base,attr"`
	// Impuesto Atributo requerido para señalar la clave del tipo de impuesto trasladado.
	Impuesto types.Impuesto `xml:"Impuesto,attr"`
	// TipoFactor Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TipoFactor types.TipoFactor `xml:"TipoFactor,attr"`
	// TasaOCuota Atributo condicional para señalar el valor de la tasa o cuota del impuesto que se traslada por los conceptos amparados en el comprobante.
	TasaOCuota decimal.Decimal `xml:"TasaOCuota,attr,omitempty"`
	// Importe Atributo condicional para señalar la suma del importe del impuesto trasladado, agrupado por impuesto, TipoFactor y TasaOCuota. No se permiten valores negativos.
	Importe decimal.Decimal `xml:"Importe,attr,omitempty"`
}

type Addenda []byte

func (a *Addenda) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var add struct {
		XMLName xml.Name
		Any     []byte `xml:",innerxml"`
	}
	if err := d.DecodeElement(&add, &start); err != nil {
		return err
	}
	*a = add.Any
	return nil
}

// Complemento Nodo opcional donde se incluye el complemento Timbre Fiscal Digital de manera obligatoria y los nodos complementarios determinados por el SAT, de acuerdo con las disposiciones particulares para un sector o actividad específica.
type Complemento struct {
	// CartaPorte20 Complemento para incorporar al Comprobante Fiscal Digital por Internet (CFDI), la información relacionada a los bienes y/o mercancías, ubicaciones de origen, puntos intermedios y destinos, así como lo referente al medio por el que se transportan; ya sea por vía terrestre (autotransporte y férrea), marítima y/o aérea; además de incluir el traslado de hidrocarburos y petrolíferos.
	CartaPorte20 *cartaporte20.CartaPorte20 `xml:"CartaPorte,omitempty"`
	// Pagos20 Complemento para el Comprobante Fiscal Digital por Internet (CFDI) para registrar información sobre la recepción de pagos. El emisor de este complemento para recepción de pagos debe ser quien las leyes le obligue a expedir comprobantes por los actos o actividades que realicen, por los ingresos que se perciban o por las retenciones de contribuciones que efectúen.
	Pagos20 *pagos20.Pagos `xml:"Pagos,omitempty"`
	// CCE11 Complemento para incorporar la información en el caso de Exportación de Mercancías en definitiva.
	CCE11 *comext11.ComercioExterior `xml:"ComercioExterior,omitempty"`
	// TFD11 Complemento requerido para el Timbrado Fiscal Digital que da validez al Comprobante fiscal digital por Internet.
	TFD11 *tfd11.TimbreFiscalDigital `xml:"TimbreFiscalDigital,omitempty"`
}
