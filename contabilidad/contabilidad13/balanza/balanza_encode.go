package balanza

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/veyronifs/cfdi-go/encoder"
)

// Unmarshal parses the XML-encoded data and returns the *Catalogo.
func Unmarshal(data []byte) (*Balanza, error) {
	var catalogo Balanza
	err := xml.Unmarshal(data, &catalogo)
	if err != nil {
		return nil, err
	}
	return &catalogo, nil
}

var catXS = encoder.NSElem{
	Prefix: "BCE",
	NS:     "http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/BalanzaComprobacion",
}

func Marshal(bal *Balanza) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(catXS.ElemXS("Balanza"))
	defer enc.EndElem("Balanza")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	enc.WriteAttrStr("xsi:schemaLocation", catXS.NS+" http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/BalanzaComprobacion/BalanzaComprobacion_1_3.xsd")

	encodeHeader(enc, bal)
	enc.EndAllFlush()
	return b.Bytes(), nil
}

func encodeHeader(enc *encoder.Encoder, bal *Balanza) {
	enc.WriteAttrStrZ("Version", bal.Version)
	enc.WriteAttrStrZ("RFC", bal.RFC)
	enc.WriteAttrStrZ("Mes", fmt.Sprintf("%02d", bal.Mes))
	enc.WriteAttrIntZ("Anio", bal.Anio)
	enc.WriteAttrStrZ("TipoEnvio", string(bal.TipoEnvio))
	if !bal.FechaModBal.IsZero() {
		enc.WriteAttrStrZ("FechaModBal", bal.FechaModBal.String())
	}
	enc.WriteAttrStrZ("Sello", bal.Sello)
	enc.WriteAttrStrZ("noCertificado", bal.NoCertificado)
	enc.WriteAttrStrZ("Certificado", bal.Certificado)

	encodeCtas(enc, bal.Ctas)
}

func encodeCtas(enc *encoder.Encoder, ctas []*Cta) {
	for _, cta := range ctas {
		enc.StartElem(catXS.Elem("Ctas"))
		enc.WriteAttrStrMaxEllipsisZ("NumCta", cta.NumCta, 100)
		enc.WriteAttrDecimal("SaldoIni", cta.SaldoIni, 2)
		enc.WriteAttrDecimal("Debe", cta.Debe, 2)
		enc.WriteAttrDecimal("Haber", cta.Haber, 2)
		enc.WriteAttrDecimal("SaldoFin", cta.SaldoFin, 2)
		enc.EndElem("Ctas")
	}
}
