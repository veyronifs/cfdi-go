package cfdi40

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/complemento/cartaporte20"
	"github.com/veyronifs/cfdi-go/complemento/tfd11"
	"github.com/veyronifs/cfdi-go/types"
)

type Comprobante struct {
	InformacionGlobal InformacionGlobal        `xml:"InformacionGlobal,omitempty"`
	CfdiRelacionados  []CfdiRelacionados       `xml:"CfdiRelacionados,omitempty"`
	Emisor            Emisor                   `xml:"Emisor"`
	Receptor          Receptor                 `xml:"Receptor"`
	Conceptos         Conceptos                `xml:"Conceptos"`
	Impuestos         *Impuestos               `xml:"Impuestos,omitempty"`
	Addenda           *Addenda                 `xml:"Addenda,omitempty"`
	Complemento       *Complemento             `xml:"Complemento,omitempty"`
	Version           string                   `xml:"Version,attr"`
	Serie             string                   `xml:"Serie,attr,omitempty"`
	Folio             string                   `xml:"Folio,attr,omitempty"`
	Fecha             types.TFechaH            `xml:"Fecha,attr"`
	Sello             string                   `xml:"Sello,attr"`
	FormaPago         types.CFormaPago         `xml:"FormaPago,attr,omitempty"`
	NoCertificado     string                   `xml:"NoCertificado,attr"`
	Certificado       string                   `xml:"Certificado,attr"`
	CondicionesDePago string                   `xml:"CondicionesDePago,attr,omitempty"`
	SubTotal          decimal.Decimal          `xml:"SubTotal,attr"`
	Descuento         decimal.Decimal          `xml:"Descuento,attr,omitempty"`
	Moneda            types.CMoneda            `xml:"Moneda,attr"`
	TipoCambio        decimal.Decimal          `xml:"TipoCambio,attr,omitempty"`
	Total             decimal.Decimal          `xml:"Total,attr"`
	TipoDeComprobante types.CTipoDeComprobante `xml:"TipoDeComprobante,attr"`
	Exportacion       types.CExportacion       `xml:"Exportacion,attr"`
	MetodoPago        types.CMetodoPago        `xml:"MetodoPago,attr,omitempty"`
	LugarExpedicion   string                   `xml:"LugarExpedicion,attr"`
	Confirmacion      string                   `xml:"Confirmacion,attr,omitempty"`
}

type InformacionGlobal struct {
	Periodicidad string `xml:"Periodicidad,attr"`
	Meses        string `xml:"Meses,attr"`
	Anio         int    `xml:"Año,attr"`
}

type CfdiRelacionados struct {
	CfdiRelacionado []CfdiRelacionado   `xml:"http://www.sat.gob.mx/cfd/4 CfdiRelacionado"`
	TipoRelacion    types.CTipoRelacion `xml:"TipoRelacion,attr"`
}

type CfdiRelacionado struct {
	UUID string `xml:"UUID,attr"`
}

type Emisor struct {
	Rfc              string               `xml:"Rfc,attr"`
	Nombre           string               `xml:"Nombre,attr"`
	RegimenFiscal    types.CRegimenFiscal `xml:"RegimenFiscal,attr"`
	FacAtrAdquirente string               `xml:"FacAtrAdquirente,attr,omitempty"`
}

type Receptor struct {
	Rfc                     string               `xml:"Rfc,attr"`
	Nombre                  string               `xml:"Nombre,attr"`
	DomicilioFiscalReceptor string               `xml:"DomicilioFiscalReceptor,attr"`
	ResidenciaFiscal        types.CPais          `xml:"ResidenciaFiscal,attr,omitempty"`
	NumRegIdTrib            string               `xml:"NumRegIdTrib,attr,omitempty"`
	RegimenFiscalReceptor   types.CRegimenFiscal `xml:"RegimenFiscalReceptor,attr"`
	UsoCFDI                 types.CUsoCFDI       `xml:"UsoCFDI,attr"`
}

type Conceptos []Concepto

func (c *Conceptos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var c1 struct {
		Conceptos []Concepto `xml:"http://www.sat.gob.mx/cfd/4 Concepto"`
	}
	if err := d.DecodeElement(&c1, &start); err != nil {
		return err
	}
	*c = c1.Conceptos
	return nil
}

type Concepto struct {
	Impuestos           *ConceptoImpuestos            `xml:"http://www.sat.gob.mx/cfd/4 Impuestos,omitempty"`
	ACuentaTerceros     *ConceptoACuentaTerceros      `xml:"http://www.sat.gob.mx/cfd/4 ACuentaTerceros,omitempty"`
	InformacionAduanera []ConceptoInformacionAduanera `xml:"http://www.sat.gob.mx/cfd/4 InformacionAduanera,omitempty"`
	CuentaPredial       []ConceptoCuentaPredial       `xml:"http://www.sat.gob.mx/cfd/4 CuentaPredial,omitempty"`
	//ComplementoConcepto ComplementoConcepto   `xml:"http://www.sat.gob.mx/cfd/4 ComplementoConcepto,omitempty"`
	Parte            []Parte          `xml:"http://www.sat.gob.mx/cfd/4 Parte,omitempty"`
	ClaveProdServ    string           `xml:"ClaveProdServ,attr"`
	NoIdentificacion string           `xml:"NoIdentificacion,attr,omitempty"`
	Cantidad         decimal.Decimal  `xml:"Cantidad,attr"`
	ClaveUnidad      string           `xml:"ClaveUnidad,attr"`
	Unidad           string           `xml:"Unidad,attr,omitempty"`
	Descripcion      string           `xml:"Descripcion,attr"`
	ValorUnitario    decimal.Decimal  `xml:"ValorUnitario,attr"`
	Importe          decimal.Decimal  `xml:"Importe,attr"`
	Descuento        decimal.Decimal  `xml:"Descuento,attr,omitempty"`
	ObjetoImp        types.CObjetoImp `xml:"ObjetoImp,attr"`
}

type ConceptoImpuestos struct {
	Traslados   ConceptoImpuestosTraslados   `xml:"http://www.sat.gob.mx/cfd/4 Traslados,omitempty"`
	Retenciones ConceptoImpuestosRetenciones `xml:"http://www.sat.gob.mx/cfd/4 Retenciones,omitempty"`
}

type ConceptoImpuestosTraslados []ConceptoImpuestosTraslado

func (tras *ConceptoImpuestosTraslados) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras1 struct {
		Traslados []ConceptoImpuestosTraslado `xml:"http://www.sat.gob.mx/cfd/4 Traslado"`
	}

	if err := d.DecodeElement(&tras1, &start); err != nil {
		return err
	}
	*tras = tras1.Traslados
	return nil
}

type ConceptoImpuestosTraslado struct {
	Base       decimal.Decimal   `xml:"Base,attr"`
	Impuesto   types.CImpuesto   `xml:"Impuesto,attr"`
	TipoFactor types.CTipoFactor `xml:"TipoFactor,attr"`
	TasaOCuota decimal.Decimal   `xml:"TasaOCuota,attr,omitempty"`
	Importe    decimal.Decimal   `xml:"Importe,attr,omitempty"`
}

type ConceptoImpuestosRetenciones []ConceptoImpuestosRetencion

func (ret *ConceptoImpuestosRetenciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret1 struct {
		Retenciones []ConceptoImpuestosRetencion `xml:"http://www.sat.gob.mx/cfd/4 Retencion"`
	}

	if err := d.DecodeElement(&ret1, &start); err != nil {
		return err
	}
	*ret = ret1.Retenciones
	return nil
}

type ConceptoImpuestosRetencion struct {
	Base       decimal.Decimal   `xml:"Base,attr"`
	Impuesto   types.CImpuesto   `xml:"Impuesto,attr"`
	TipoFactor types.CTipoFactor `xml:"TipoFactor,attr"`
	TasaOCuota decimal.Decimal   `xml:"TasaOCuota,attr"`
	Importe    decimal.Decimal   `xml:"Importe,attr"`
}

type ConceptoACuentaTerceros struct {
	RfcACuentaTerceros             string               `xml:"RfcACuentaTerceros,attr"`
	NombreACuentaTerceros          string               `xml:"NombreACuentaTerceros,attr"`
	RegimenFiscalACuentaTerceros   types.CRegimenFiscal `xml:"RegimenFiscalACuentaTerceros,attr"`
	DomicilioFiscalACuentaTerceros string               `xml:"DomicilioFiscalACuentaTerceros,attr"`
}

type ConceptoInformacionAduanera struct {
	NumeroPedimento string `xml:"NumeroPedimento,attr"`
}

type ConceptoCuentaPredial struct {
	Numero string `xml:"Numero,attr"`
}

type Parte struct {
	InformacionAduanera []ConceptoInformacionAduanera `xml:"http://www.sat.gob.mx/cfd/4 InformacionAduanera,omitempty"`
	ClaveProdServ       types.CClaveProdServ          `xml:"ClaveProdServ,attr"`
	NoIdentificacion    string                        `xml:"NoIdentificacion,attr,omitempty"`
	Cantidad            decimal.Decimal               `xml:"Cantidad,attr"`
	Unidad              string                        `xml:"Unidad,attr,omitempty"`
	Descripcion         string                        `xml:"Descripcion,attr"`
	ValorUnitario       decimal.Decimal               `xml:"ValorUnitario,attr,omitempty"`
	Importe             decimal.Decimal               `xml:"Importe,attr,omitempty"`
}

type Impuestos struct {
	Retenciones               ImpuestosRetenciones `xml:"http://www.sat.gob.mx/cfd/4 Retenciones,omitempty"`
	Traslados                 ImpuestosTraslados   `xml:"http://www.sat.gob.mx/cfd/4 Traslados,omitempty"`
	TotalImpuestosRetenidos   decimal.Decimal      `xml:"TotalImpuestosRetenidos,attr,omitempty"`
	TotalImpuestosTrasladados decimal.Decimal      `xml:"TotalImpuestosTrasladados,attr,omitempty"`
}

type ImpuestosRetenciones []ImpuestosRetencion

func (ret *ImpuestosRetenciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret1 struct {
		Retenciones []ImpuestosRetencion `xml:"http://www.sat.gob.mx/cfd/4 Retencion"`
	}

	if err := d.DecodeElement(&ret1, &start); err != nil {
		return err
	}
	*ret = ret1.Retenciones
	return nil
}

type ImpuestosRetencion struct {
	Impuesto types.CImpuesto `xml:"Impuesto,attr"`
	Importe  decimal.Decimal `xml:"Importe,attr"`
}

type ImpuestosTraslados []ImpuestosTraslado

func (tras *ImpuestosTraslados) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras1 struct {
		Traslados []ImpuestosTraslado `xml:"http://www.sat.gob.mx/cfd/4 Traslado"`
	}

	if err := d.DecodeElement(&tras1, &start); err != nil {
		return err
	}
	*tras = tras1.Traslados
	return nil
}

type ImpuestosTraslado struct {
	Base       decimal.Decimal   `xml:"Base,attr"`
	Impuesto   types.CImpuesto   `xml:"Impuesto,attr"`
	TipoFactor types.CTipoFactor `xml:"TipoFactor,attr"`
	TasaOCuota decimal.Decimal   `xml:"TasaOCuota,attr,omitempty"`
	Importe    decimal.Decimal   `xml:"Importe,attr,omitempty"`
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

type Complemento struct {
	CartaPorte20 *cartaporte20.CartaPorte20 `xml:"CartaPorte,omitempty"`
	TFD11        *tfd11.TimbreFiscalDigital `xml:"TimbreFiscalDigital,omitempty"`
}
