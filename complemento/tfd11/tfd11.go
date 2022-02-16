package tfd11

import (
	"encoding/xml"

	"github.com/veyronifs/cfdi-go/types"
)

func Unmarshal(b []byte) (*TimbreFiscalDigital, error) {
	tfd := &TimbreFiscalDigital{}
	if err := xml.Unmarshal(b, tfd); err != nil {
		return nil, err
	}
	return tfd, nil
}

type TimbreFiscalDigital struct {
	Version          string       `xml:"Version,attr"`
	UUID             string       `xml:"UUID,attr"`
	FechaTimbrado    types.FechaH `xml:"FechaTimbrado,attr"`
	RfcProvCertif    string       `xml:"RfcProvCertif,attr"`
	Leyenda          string       `xml:"Leyenda,attr,omitempty"`
	SelloCFD         string       `xml:"SelloCFD,attr"`
	NoCertificadoSAT string       `xml:"NoCertificadoSAT,attr"`
	SelloSAT         string       `xml:"SelloSAT,attr"`
}
