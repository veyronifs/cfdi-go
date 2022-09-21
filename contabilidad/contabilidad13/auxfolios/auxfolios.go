package auxfolios

import (
	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

type RepAuxFolios struct {
	DetAuxFolios  []*DetAuxFol  `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios DetAuxFol,omitempty"`
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

type DetAuxFol struct {
	ComprNal     []*ComprNal    `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios ComprNal,omitempty"`
	ComprNalOtr  []*ComprNalOtr `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios ComprNalOtr,omitempty"`
	ComprExt     []*ComprExt    `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios ComprExt,omitempty"`
	NumUnIdenPol string         `xml:"NumUnIdenPol,attr"`
	Fecha        types.Fecha    `xml:"Fecha,attr"`
}

type ComprNal struct {
	UUIDCFDI   string          `xml:"UUID_CFDI,attr"`
	MontoTotal decimal.Decimal `xml:"MontoTotal,attr"`
	RFC        string          `xml:"RFC,attr"`
	MetPagoAux string          `xml:"MetPagoAux,attr,omitempty"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type ComprNalOtr struct {
	CFDCBBSerie  string          `xml:"CFD_CBB_Serie,attr,omitempty"`
	CFDCBBNumFol int             `xml:"CFD_CBB_NumFol,attr"`
	MontoTotal   decimal.Decimal `xml:"MontoTotal,attr"`
	RFC          string          `xml:"RFC,attr"`
	MetPagoAux   string          `xml:"MetPagoAux,attr,omitempty"`
	Moneda       types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb      decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

type ComprExt struct {
	NumFactExt string          `xml:"NumFactExt,attr"`
	TaxID      string          `xml:"TaxID,attr,omitempty"`
	MontoTotal decimal.Decimal `xml:"MontoTotal,attr"`
	MetPagoAux string          `xml:"MetPagoAux,attr,omitempty"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}
