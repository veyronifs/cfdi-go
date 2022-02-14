package tfd11

import (
	"bytes"

	"github.com/shabbyrobe/xmlwriter"
	"github.com/veyronifs/cfdi-go/encoder"
)

var tfd11XS = encoder.NSElem{
	Prefix: "tfd",
	NS:     "http://www.sat.gob.mx/TimbreFiscalDigital",
}

func Marshal(tfd *TimbreFiscalDigital) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	tfd.MarshalComplemento(enc)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func (tfd *TimbreFiscalDigital) MarshalComplemento(enc *encoder.Encoder) {
	if tfd == nil {
		return
	}
	enc.StartElem(tfd11XS.ElemXS("TimbreFiscalDigital"))
	defer enc.EndElem("TimbreFiscalDigital")

	enc.WriteAttr(
		xmlwriter.Attr{
			Prefix: "xmlns",
			Name:   "xsi",
			Value:  "http://www.w3.org/2001/XMLSchema-instance",
		},
		xmlwriter.Attr{
			Prefix: "xsi",
			Name:   "schemaLocation",
			Value:  "http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd",
		},
	)

	enc.WriteAttrStr("Version", tfd.Version)
	enc.WriteAttrStr("UUID", tfd.UUID)
	enc.WriteAttrStr("FechaTimbrado", tfd.FechaTimbrado.Encode())
	enc.WriteAttrStr("RfcProvCertif", tfd.RfcProvCertif)
	enc.WriteAttrStr("SelloCFD", tfd.SelloCFD)
	enc.WriteAttrStr("NoCertificadoSAT", tfd.NoCertificadoSAT)
	enc.WriteAttrStr("SelloSAT", tfd.SelloSAT)
	enc.WriteAttrStrZ("Leyenda", tfd.Leyenda)
}
