package tfd11

import (
	"github.com/veyronifs/cfdi-go/types"
)

type TimbreFiscalDigital struct {
	Version          string        `xml:"Version,attr"`
	UUID             string        `xml:"UUID,attr"`
	FechaTimbrado    types.TFechaH `xml:"FechaTimbrado,attr"`
	RfcProvCertif    string        `xml:"RfcProvCertif,attr"`
	Leyenda          string        `xml:"Leyenda,attr,omitempty"`
	SelloCFD         string        `xml:"SelloCFD,attr"`
	NoCertificadoSAT string        `xml:"NoCertificadoSAT,attr"`
	SelloSAT         string        `xml:"SelloSAT,attr"`
}
