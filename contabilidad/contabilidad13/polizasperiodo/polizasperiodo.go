package polizasperiodo

import (
	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/contabilidad/contabilidad13"
	"github.com/veyronifs/cfdi-go/types"
)

type Polizas struct {
	Polizas       []*Poliza     `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo Poliza"`
	Version       string        `xml:"Version,attr"`
	RFC           string        `xml:"RFC,attr"`
	Mes           int           `xml:"Mes,attr"`
	Anio          int           `xml:"Anio,attr"`
	TipoSolicitud TipoSolicitud `xml:"TipoSolicitud,attr"`
	NumOrden      string        `xml:"NumOrden,attr,omitempty"`
	NumTramite    string        `xml:"NumTramite,attr,omitempty"`
	Sello         string        `xml:"Sello,attr,omitempty"`
	NoCertificado string        `xml:"noCertificado,attr,omitempty"`
	Certificado   string        `xml:"Certificado,attr,omitempty"`
}

// Must match the pattern AF|FC|DE|CO
type TipoSolicitud string

const (
	TipoSolicitudAF TipoSolicitud = "AF"
	TipoSolicitudFC TipoSolicitud = "FC"
	TipoSolicitudDE TipoSolicitud = "DE"
	TipoSolicitudCO TipoSolicitud = "CO"
)

func (t TipoSolicitud) Desc() string {
	switch t {
	case TipoSolicitudAF:
		return "Acto de Fiscalización"
	case TipoSolicitudFC:
		return "Fiscalización por Compulsa"
	case TipoSolicitudDE:
		return "Devolución"
	case TipoSolicitudCO:
		return "Compensación"
	}
	return ""
}

type Poliza struct {
	Transaccion  []*Transaccion `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo Transaccion"`
	NumUnIdenPol string         `xml:"NumUnIdenPol,attr"`
	Fecha        types.Fecha    `xml:"Fecha,attr"`
	Concepto     string         `xml:"Concepto,attr"`
}

type Transaccion struct {
	CompNal       []*CompNal       `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo CompNal,omitempty"`
	CompNalOtr    []*CompNalOtr    `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo CompNalOtr,omitempty"`
	CompExt       []*CompExt       `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo CompExt,omitempty"`
	Cheque        []*Cheque        `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo Cheque,omitempty"`
	Transferencia []*Transferencia `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo Transferencia,omitempty"`
	OtrMetodoPago []*OtrMetodoPago `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo OtrMetodoPago,omitempty"`
	NumCta        string           `xml:"NumCta,attr"`
	DesCta        string           `xml:"DesCta,attr"`
	Concepto      string           `xml:"Concepto,attr"`
	Debe          decimal.Decimal  `xml:"Debe,attr"`
	Haber         decimal.Decimal  `xml:"Haber,attr"`
}

type CompNal struct {
	UUIDCFDI   string          `xml:"UUID_CFDI,attr"`
	RFC        string          `xml:"RFC,attr"`
	MontoTotal decimal.Decimal `xml:"MontoTotal,attr"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type CompNalOtr struct {
	CFDCBBSerie  string          `xml:"CFD_CBB_Serie,attr,omitempty"`
	CFDCBBNumFol int             `xml:"CFD_CBB_NumFol,attr"`
	RFC          string          `xml:"RFC,attr"`
	MontoTotal   decimal.Decimal `xml:"MontoTotal,attr"`
	Moneda       types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb      decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type CompExt struct {
	NumFactExt string          `xml:"NumFactExt,attr"`
	TaxID      string          `xml:"TaxID,attr,omitempty"`
	MontoTotal decimal.Decimal `xml:"MontoTotal,attr"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type Cheque struct {
	Num        string          `xml:"Num,attr"`
	BanEmisNal string          `xml:"BanEmisNal,attr"`
	BanEmisExt string          `xml:"BanEmisExt,attr,omitempty"`
	CtaOri     string          `xml:"CtaOri,attr"`
	Fecha      types.Fecha     `xml:"Fecha,attr"`
	Benef      string          `xml:"Benef,attr"`
	RFC        string          `xml:"RFC,attr"`
	Monto      decimal.Decimal `xml:"Monto,attr"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type Transferencia struct {
	CtaOri       string          `xml:"CtaOri,attr,omitempty"`
	BancoOriNal  string          `xml:"BancoOriNal,attr"`
	BancoOriExt  string          `xml:"BancoOriExt,attr,omitempty"`
	CtaDest      string          `xml:"CtaDest,attr"`
	BancoDestNal string          `xml:"BancoDestNal,attr"`
	BancoDestExt string          `xml:"BancoDestExt,attr,omitempty"`
	Fecha        types.Fecha     `xml:"Fecha,attr"`
	Benef        string          `xml:"Benef,attr"`
	RFC          string          `xml:"RFC,attr"`
	Monto        decimal.Decimal `xml:"Monto,attr"`
	Moneda       types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb      decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type OtrMetodoPago struct {
	MetPagoPol string          `xml:"MetPagoPol,attr"`
	Fecha      types.Fecha     `xml:"Fecha,attr"`
	Benef      string          `xml:"Benef,attr"`
	RFC        string          `xml:"RFC,attr"`
	Monto      decimal.Decimal `xml:"Monto,attr"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

/*
func (b Catalogo) Archivo() *contabilidad13.Archivo {
	return &contabilidad13.Archivo{
		RFC:  b.RFC,
		Anio: b.Anio,
		Mes:  b.Mes,
		Tipo: "CT",
	}
}
*/

func (b Polizas) Archivo() *contabilidad13.Archivo {
	return &contabilidad13.Archivo{
		RFC:  b.RFC,
		Mes:  b.Mes,
		Anio: b.Anio,
		Tipo: "PL",
	}
}
