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

type Pagos struct {
	Totales *Totales `xml:"Totales"`
	Pago    []*Pago  `xml:"Pago"`
	Version string   `xml:"Version,attr"`
}

type Totales struct {
	TotalRetencionesIVA         decimal.Decimal `xml:"TotalRetencionesIVA,attr,omitempty"`
	TotalRetencionesISR         decimal.Decimal `xml:"TotalRetencionesISR,attr,omitempty"`
	TotalRetencionesIEPS        decimal.Decimal `xml:"TotalRetencionesIEPS,attr,omitempty"`
	TotalTrasladosBaseIVA16     decimal.Decimal `xml:"TotalTrasladosBaseIVA16,attr,omitempty"`
	TotalTrasladosImpuestoIVA16 decimal.Decimal `xml:"TotalTrasladosImpuestoIVA16,attr,omitempty"`
	TotalTrasladosBaseIVA8      decimal.Decimal `xml:"TotalTrasladosBaseIVA8,attr,omitempty"`
	TotalTrasladosImpuestoIVA8  decimal.Decimal `xml:"TotalTrasladosImpuestoIVA8,attr,omitempty"`
	TotalTrasladosBaseIVA0      decimal.Decimal `xml:"TotalTrasladosBaseIVA0,attr,omitempty"`
	TotalTrasladosImpuestoIVA0  decimal.Decimal `xml:"TotalTrasladosImpuestoIVA0,attr,omitempty"`
	TotalTrasladosBaseIVAExento decimal.Decimal `xml:"TotalTrasladosBaseIVAExento,attr,omitempty"`
	MontoTotalPagos             decimal.Decimal `xml:"MontoTotalPagos,attr"`
}

type Pago struct {
	DoctoRelacionado []*DoctoRelacionado `xml:"DoctoRelacionado"`
	ImpuestosP       *ImpuestosP         `xml:"ImpuestosP,omitempty"`
	FechaPago        types.FechaH        `xml:"FechaPago,attr"`
	FormaDePagoP     types.FormaPago     `xml:"FormaDePagoP,attr"`
	MonedaP          types.Moneda        `xml:"MonedaP,attr"`
	TipoCambioP      decimal.Decimal     `xml:"TipoCambioP,attr,omitempty"`
	Monto            decimal.Decimal     `xml:"Monto,attr"`
	NumOperacion     string              `xml:"NumOperacion,attr,omitempty"`
	RfcEmisorCtaOrd  string              `xml:"RfcEmisorCtaOrd,attr,omitempty"`
	NomBancoOrdExt   string              `xml:"NomBancoOrdExt,attr,omitempty"`
	CtaOrdenante     string              `xml:"CtaOrdenante,attr,omitempty"`
	RfcEmisorCtaBen  string              `xml:"RfcEmisorCtaBen,attr,omitempty"`
	CtaBeneficiario  string              `xml:"CtaBeneficiario,attr,omitempty"`
	TipoCadPago      CTipoCadenaPago     `xml:"TipoCadPago,attr,omitempty"`
	CertPago         string              `xml:"CertPago,attr,omitempty"`
	CadPago          string              `xml:"CadPago,attr,omitempty"`
	SelloPago        string              `xml:"SelloPago,attr,omitempty"`
}

// May be one of 01
type CTipoCadenaPago string

const (
	CTipoCadenaPago01 CTipoCadenaPago = "01"
)

type DoctoRelacionado struct {
	ImpuestosDR      *ImpuestosDR    `xml:"ImpuestosDR,omitempty"`
	IdDocumento      string          `xml:"IdDocumento,attr"`
	Serie            string          `xml:"Serie,attr,omitempty"`
	Folio            string          `xml:"Folio,attr,omitempty"`
	MonedaDR         types.Moneda    `xml:"MonedaDR,attr"`
	EquivalenciaDR   decimal.Decimal `xml:"EquivalenciaDR,attr,omitempty"`
	NumParcialidad   int             `xml:"NumParcialidad,attr"`
	ImpSaldoAnt      decimal.Decimal `xml:"ImpSaldoAnt,attr"`
	ImpPagado        decimal.Decimal `xml:"ImpPagado,attr"`
	ImpSaldoInsoluto decimal.Decimal `xml:"ImpSaldoInsoluto,attr"`
	ObjetoImpDR      types.ObjetoImp `xml:"ObjetoImpDR,attr"`
}

type ImpuestosDR struct {
	RetencionesDR RetencionesDR `xml:"RetencionesDR,omitempty"`
	TrasladosDR   TrasladosDR   `xml:"TrasladosDR,omitempty"`
}

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

type RetencionDR struct {
	BaseDR       decimal.Decimal  `xml:"BaseDR,attr"`
	ImpuestoDR   types.Impuesto   `xml:"ImpuestoDR,attr"`
	TipoFactorDR types.TipoFactor `xml:"TipoFactorDR,attr"`
	TasaOCuotaDR decimal.Decimal  `xml:"TasaOCuotaDR,attr"`
	ImporteDR    decimal.Decimal  `xml:"ImporteDR,attr"`
}

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

type TrasladoDR struct {
	BaseDR       decimal.Decimal  `xml:"BaseDR,attr"`
	ImpuestoDR   types.Impuesto   `xml:"ImpuestoDR,attr"`
	TipoFactorDR types.TipoFactor `xml:"TipoFactorDR,attr"`
	TasaOCuotaDR decimal.Decimal  `xml:"TasaOCuotaDR,attr,omitempty"`
	ImporteDR    decimal.Decimal  `xml:"ImporteDR,attr,omitempty"`
}

type ImpuestosP struct {
	RetencionesP RetencionesP `xml:"RetencionesP,omitempty"`
	TrasladosP   TrasladosP   `xml:"TrasladosP,omitempty"`
}

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

type RetencionP struct {
	ImpuestoP types.Impuesto  `xml:"ImpuestoP,attr"`
	ImporteP  decimal.Decimal `xml:"ImporteP,attr"`
}

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

type TrasladoP struct {
	BaseP       decimal.Decimal  `xml:"BaseP,attr"`
	ImpuestoP   types.Impuesto   `xml:"ImpuestoP,attr"`
	TipoFactorP types.TipoFactor `xml:"TipoFactorP,attr"`
	TasaOCuotaP decimal.Decimal  `xml:"TasaOCuotaP,attr,omitempty"`
	ImporteP    decimal.Decimal  `xml:"ImporteP,attr,omitempty"`
}
