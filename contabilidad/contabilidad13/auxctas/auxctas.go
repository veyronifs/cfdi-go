package auxctas

import (
	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

type AuxiliarCtas struct {
	Cuentas       []*Cuenta     `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas Cuenta"`
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

type Cuenta struct {
	DetallesAux []*DetalleAux   `xml:"http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas DetalleAux"`
	NumCta      string          `xml:"NumCta,attr"`
	DesCta      string          `xml:"DesCta,attr"`
	SaldoIni    decimal.Decimal `xml:"SaldoIni,attr"`
	SaldoFin    decimal.Decimal `xml:"SaldoFin,attr"`
}

type DetalleAux struct {
	Fecha        types.Fecha     `xml:"Fecha,attr"`
	NumUnIdenPol string          `xml:"NumUnIdenPol,attr"`
	Concepto     string          `xml:"Concepto,attr"`
	Debe         decimal.Decimal `xml:"Debe,attr"`
	Haber        decimal.Decimal `xml:"Haber,attr"`
}
