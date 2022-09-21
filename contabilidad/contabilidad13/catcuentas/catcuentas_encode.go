package catcuentas

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/veyronifs/cfdi-go/encoder"
)

// Unmarshal parses the XML-encoded data and returns the *Catalogo.
func Unmarshal(data []byte) (*Catalogo, error) {
	var catalogo Catalogo
	err := xml.Unmarshal(data, &catalogo)
	if err != nil {
		return nil, err
	}
	return &catalogo, nil
}

var catXS = encoder.NSElem{
	Prefix: "catalogocuentas",
	NS:     "http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/CatalogoCuentas",
}

func Marshal(c *Catalogo) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(catXS.ElemXS("Catalogo"))
	defer enc.EndElem("Catalogo")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	enc.WriteAttrStr("xsi:schemaLocation", catXS.NS+" http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/CatalogoCuentas/CatalogoCuentas_1_3.xsd")

	encodeHeader(enc, c)
	enc.EndAllFlush()
	return b.Bytes(), nil
}

func encodeHeader(enc *encoder.Encoder, c *Catalogo) {
	enc.WriteAttrStrZ("Version", c.Version)
	enc.WriteAttrStrZ("RFC", c.RFC)
	enc.WriteAttrStrZ("Mes", fmt.Sprintf("%02d", c.Mes))
	enc.WriteAttrIntZ("Anio", c.Anio)
	enc.WriteAttrStrZ("Sello", c.Sello)
	enc.WriteAttrStrZ("noCertificado", c.NoCertificado)
	enc.WriteAttrStrZ("Certificado", c.Certificado)

	encodeCtas(enc, c.Ctas)
}

func encodeCtas(enc *encoder.Encoder, ctas []*Cta) {
	for _, cta := range ctas {
		enc.StartElem(catXS.Elem("Ctas"))
		enc.WriteAttrStrZ("CodAgrup", string(cta.CodAgrup))
		enc.WriteAttrStrZ("NumCta", cta.NumCta)
		enc.WriteAttrStrZ("Desc", cta.Desc)
		enc.WriteAttrStrZ("SubCtaDe", cta.SubCtaDe)
		enc.WriteAttrInt("Nivel", cta.Nivel)
		enc.WriteAttrStr("Natur", string(cta.Natur))
		enc.EndElem("Ctas")
	}
}
