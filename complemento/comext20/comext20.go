package comext20

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

func Unmarshal(b []byte) (*ComercioExterior, error) {
	comext := &ComercioExterior{}
	if err := xml.Unmarshal(b, comext); err != nil {
		return nil, err
	}
	return comext, nil
}

type ComercioExterior struct {
	Emisor                    *Emisor         `xml:"Emisor,omitempty"`
	Propietarios              []*Propietario  `xml:"Propietario,omitempty"`
	Receptor                  *Receptor       `xml:"Receptor,omitempty"`
	Destinatarios             []*Destinatario `xml:"Destinatario,omitempty"`
	Mercancias                Mercancias      `xml:"Mercancias,omitempty"`
	Version                   string          `xml:"Version,attr"`
	MotivoTraslado            string          `xml:"MotivoTraslado,attr,omitempty"`
	TipoOperacion             string          `xml:"TipoOperacion,attr"`
	ClaveDePedimento          string          `xml:"ClaveDePedimento,attr,omitempty"`
	CertificadoOrigen         int             `xml:"CertificadoOrigen,attr,omitempty"`
	NumCertificadoOrigen      string          `xml:"NumCertificadoOrigen,attr,omitempty"`
	NumeroExportadorConfiable string          `xml:"NumeroExportadorConfiable,attr,omitempty"`
	Incoterm                  string          `xml:"Incoterm,attr,omitempty"`
	Subdivision               int             `xml:"Subdivision,attr,omitempty"`
	Observaciones             string          `xml:"Observaciones,attr,omitempty"`
	TipoCambioUSD             decimal.Decimal `xml:"TipoCambioUSD,attr,omitempty"`
	TotalUSD                  decimal.Decimal `xml:"TotalUSD,attr,omitempty"`
}

type Domicilio struct {
	Calle          string     `xml:"Calle,attr"`
	NumeroExterior string     `xml:"NumeroExterior,attr,omitempty"`
	NumeroInterior string     `xml:"NumeroInterior,attr,omitempty"`
	Colonia        string     `xml:"Colonia,attr,omitempty"`
	Localidad      string     `xml:"Localidad,attr,omitempty"`
	Referencia     string     `xml:"Referencia,attr,omitempty"`
	Municipio      string     `xml:"Municipio,attr,omitempty"`
	Estado         string     `xml:"Estado,attr"`
	Pais           types.Pais `xml:"Pais,attr"`
	CodigoPostal   string     `xml:"CodigoPostal,attr"`
}

type DescripcionesEspecificas struct {
	Marca       string `xml:"Marca,attr"`
	Modelo      string `xml:"Modelo,attr,omitempty"`
	SubModelo   string `xml:"SubModelo,attr,omitempty"`
	NumeroSerie string `xml:"NumeroSerie,attr,omitempty"`
}

type Destinatario struct {
	Domicilios   []*Domicilio `xml:"Domicilio"`
	NumRegIdTrib string       `xml:"NumRegIdTrib,attr,omitempty"`
	Nombre       string       `xml:"Nombre,attr,omitempty"`
}

type Emisor struct {
	Domicilio *Domicilio `xml:"Domicilio,omitempty"`
	Curp      string     `xml:"Curp,attr,omitempty"`
}

type Mercancias []*Mercancia

func (u *Mercancias) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ubics struct {
		Slice []*Mercancia `xml:"Mercancia"`
	}

	if err := d.DecodeElement(&ubics, &start); err != nil {
		return err
	}
	*u = ubics.Slice
	return nil
}

type Mercancia struct {
	DescripcionesEspecificas []*DescripcionesEspecificas `xml:"DescripcionesEspecificas,omitempty"`
	NoIdentificacion         string                      `xml:"NoIdentificacion,attr"`
	FraccionArancelaria      string                      `xml:"FraccionArancelaria,attr,omitempty"`
	CantidadAduana           decimal.Decimal             `xml:"CantidadAduana,attr,omitempty"`
	UnidadAduana             string                      `xml:"UnidadAduana,attr,omitempty"`
	ValorUnitarioAduana      decimal.Decimal             `xml:"ValorUnitarioAduana,attr,omitempty"`
	ValorDolares             decimal.Decimal             `xml:"ValorDolares,attr"`
}

type Propietario struct {
	NumRegIdTrib     string     `xml:"NumRegIdTrib,attr"`
	ResidenciaFiscal types.Pais `xml:"ResidenciaFiscal,attr"`
}

type Receptor struct {
	Domicilio    *Domicilio `xml:"Domicilio,omitempty"`
	NumRegIdTrib string     `xml:"NumRegIdTrib,attr,omitempty"`
}
