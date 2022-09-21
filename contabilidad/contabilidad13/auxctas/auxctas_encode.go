package auxctas

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/veyronifs/cfdi-go/encoder"
)

// Unmarshal parses the XML-encoded data and returns the *AuxiliarCtas.
func Unmarshal(data []byte) (*AuxiliarCtas, error) {
	var auxiliarctas AuxiliarCtas
	err := xml.Unmarshal(data, &auxiliarctas)
	if err != nil {
		return nil, err
	}
	return &auxiliarctas, nil
}

var auxCtaXS = encoder.NSElem{
	Prefix: "AuxiliarCtas",
	NS:     "http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas",
}

func Marshal(c *AuxiliarCtas) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(auxCtaXS.ElemXS("AuxiliarCtas"))
	defer enc.EndElem("AuxiliarCtas")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	enc.WriteAttrStr("xsi:schemaLocation", auxCtaXS.NS+" https://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas/AuxiliarCtas_1_3.xsd")

	encodeAuxiliarCtas(enc, c)
	enc.EndAllFlush()
	return b.Bytes(), nil
}

func encodeAuxiliarCtas(enc *encoder.Encoder, c *AuxiliarCtas) {
	enc.WriteAttrStrZ("Version", c.Version)
	enc.WriteAttrStrZ("RFC", c.RFC)
	enc.WriteAttrStrZ("Mes", fmt.Sprintf("%02d", c.Mes))
	enc.WriteAttrIntZ("Anio", c.Anio)
	enc.WriteAttrStrZ("TipoSolicitud", string(c.TipoSolicitud))
	enc.WriteAttrStrZ("NumOrden", c.NumOrden)
	enc.WriteAttrStrZ("NumTramite", c.NumTramite)
	enc.WriteAttrStrZ("Sello", c.Sello)
	enc.WriteAttrStrZ("noCertificado", c.NoCertificado)
	enc.WriteAttrStrZ("Certificado", c.Certificado)
	for _, cuentas := range c.Cuentas {
		enc.StartElem(auxCtaXS.Elem("Cuenta"))
		enc.WriteAttrStrZ("NumCta", cuentas.NumCta)
		enc.WriteAttrStrZ("DesCta", cuentas.DesCta)
		enc.WriteAttrDecimal("SaldoIni", cuentas.SaldoIni, 2)
		enc.WriteAttrDecimal("SaldoFin", cuentas.SaldoFin, 2)
		encodeDetalleAux(enc, cuentas.DetallesAux)
		enc.EndElem("Cuenta")
	}
}

func encodeDetalleAux(enc *encoder.Encoder, c []*DetalleAux) {
	for _, detAux := range c {
		enc.StartElem(auxCtaXS.Elem("DetalleAux"))
		enc.WriteAttrStrZ("Fecha", detAux.Fecha.String())
		enc.WriteAttrStrZ("NumUnIdenPol", detAux.NumUnIdenPol)
		enc.WriteAttrStrZ("Concepto", detAux.Concepto)
		enc.WriteAttrDecimal("Debe", detAux.Debe, 2)
		enc.WriteAttrDecimal("Haber", detAux.Haber, 2)
		enc.EndElem("DetalleAux")
	}
}
