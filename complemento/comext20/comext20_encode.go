package comext20

import (
	"bytes"
	"strconv"

	"github.com/shabbyrobe/xmlwriter"
	"github.com/veyronifs/cfdi-go/encoder"
)

var comextXS = encoder.NSElem{
	Prefix: "cce20",
	NS:     "http://www.sat.gob.mx/ComercioExterior20",
}

func (ce *ComercioExterior) SchemaLocation() string {
	return comextXS.NS + " http://www.sat.gob.mx/sitio_internet/cfd/ComercioExterior20/ComercioExterior20.xsd"
}

func (ce *ComercioExterior) XmlNSPrefix() string {
	return comextXS.Prefix
}

func (ce *ComercioExterior) XmlNS() string {
	return comextXS.NS
}

func Marshal(ce *ComercioExterior) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	ce.MarshalComplemento(enc)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func (ce *ComercioExterior) MarshalComplemento(enc *encoder.Encoder) {
	if ce == nil {
		return
	}
	enc.StartElem(comextXS.ElemXS("ComercioExterior"))
	defer enc.EndElem("ComercioExterior")

	enc.WriteAttr(
		xmlwriter.Attr{
			Prefix: "xmlns",
			Name:   "xsi",
			Value:  "http://www.w3.org/2001/XMLSchema-instance",
		},
		xmlwriter.Attr{
			Prefix: "xsi",
			Name:   "schemaLocation",
			Value:  ce.SchemaLocation(),
		},
	)

	enc.WriteAttrStr("Version", ce.Version)
	enc.WriteAttrStrZ("MotivoTraslado", ce.MotivoTraslado)
	enc.WriteAttrStr("ClaveDePedimento", ce.ClaveDePedimento)
	enc.WriteAttrStr("CertificadoOrigen", strconv.Itoa(ce.CertificadoOrigen))
	enc.WriteAttrStrZ("NumCertificadoOrigen", ce.NumCertificadoOrigen)
	enc.WriteAttrStrZ("NumeroExportadorConfiable", ce.NumeroExportadorConfiable)
	enc.WriteAttrStrZ("Incoterm", ce.Incoterm)
	enc.WriteAttrStrZ("Observaciones", ce.Observaciones)
	enc.WriteAttrDecimalZ("TipoCambioUSD", ce.TipoCambioUSD, 6)
	if !ce.TotalUSD.IsZero() {
		enc.WriteAttrStr("TotalUSD", ce.TotalUSD.StringFixed(2))
	}

	encodeEmisor(enc, ce.Emisor)
	encodeReceptor(enc, ce.Receptor)
	encodePropietarios(enc, ce.Propietarios)
	encodeDestinatarios(enc, ce.Destinatarios)
	encodeMercancias(enc, ce.Mercancias)

	// for _, m := range ce.Propietarios {
	// 	enc.StartElem(comextXS.Elem("Propietario"))
	// 	defer enc.EndElem("Propietario")
	// 	enc.WriteAttrStrZ("NumRegIdTrib", m.NumRegIdTrib)
	// 	enc.WriteAttrStrZ("ResidenciaFiscal", string(m.ResidenciaFiscal))
	// }

}

func encodeMercancias(enc *encoder.Encoder, mercancias []*Mercancia) {
	enc.StartElem(comextXS.Elem("Mercancias"))
	defer enc.EndElem("Mercancias")
	for _, m := range mercancias {
		encodeMercanciasMercancia(enc, m)
	}
}

func encodeMercanciasMercancia(enc *encoder.Encoder, m *Mercancia) {
	enc.StartElem(comextXS.Elem("Mercancia"))
	defer enc.EndElem("Mercancia")

	enc.WriteAttrStrZ("NoIdentificacion", m.NoIdentificacion)
	enc.WriteAttrStrZ("FraccionArancelaria", m.FraccionArancelaria)
	enc.WriteAttrDecimalZ("CantidadAduana", m.CantidadAduana, 2)
	enc.WriteAttrStrZ("UnidadAduana", m.UnidadAduana)
	enc.WriteAttrDecimalZ("ValorUnitarioAduana", m.ValorUnitarioAduana, 2)
	enc.WriteAttrDecimalZ("ValorDolares", m.ValorDolares, 2)

	for _, m := range m.DescripcionesEspecificas {
		enc.WriteAttrStrZ("Marca", m.Marca)
		enc.WriteAttrStrZ("Modelo", m.Modelo)
		enc.WriteAttrStrZ("SubModelo", m.SubModelo)
		enc.WriteAttrStrZ("NumeroSerie", m.NumeroSerie)
	}

}

func encodeEmisor(enc *encoder.Encoder, emisor *Emisor) {
	if emisor == nil {
		return
	}
	enc.StartElem(comextXS.Elem("Emisor"))
	defer enc.EndElem("Emisor")

	enc.WriteAttrStrZ("Curp", emisor.Curp)
	encodeDomicilio(enc, emisor.Domicilio)
}

func encodeReceptor(enc *encoder.Encoder, receptor *Receptor) {
	if receptor == nil {
		return
	}
	enc.StartElem(comextXS.Elem("Receptor"))
	defer enc.EndElem("Receptor")

	enc.WriteAttrStrZ("NumRegIdTrib", receptor.NumRegIdTrib)
	encodeDomicilio(enc, receptor.Domicilio)
}

func encodePropietarios(enc *encoder.Encoder, Propietarios []*Propietario) {
	for _, propietario := range Propietarios {
		enc.StartElem(comextXS.Elem("Propietario"))
		enc.WriteAttrStrZ("NumRegIdTrib", propietario.NumRegIdTrib)
		enc.WriteAttrStrZ("ResidenciaFiscal", string(propietario.ResidenciaFiscal))
		enc.EndElem("Propietario")
	}
}

func encodeDestinatarios(enc *encoder.Encoder, Destinatarios []*Destinatario) {
	for _, m := range Destinatarios {
		enc.StartElem(comextXS.Elem("Destinatario"))
		enc.WriteAttrStrZ("NumRegIdTrib", m.NumRegIdTrib)
		enc.WriteAttrStrZ("Nombre", m.Nombre)
		for _, domicilio := range m.Domicilios {
			encodeDomicilio(enc, domicilio)
		}
		enc.EndElem("Destinatario")
	}
}

func encodeDomicilio(enc *encoder.Encoder, m *Domicilio) {
	enc.StartElem(comextXS.Elem("Domicilio"))
	defer enc.EndElem("Domicilio")

	enc.WriteAttrStrZ("Calle", m.Calle)
	enc.WriteAttrStrZ("NumeroExterior", m.NumeroExterior)
	enc.WriteAttrStrZ("NumeroInterior", m.NumeroInterior)
	enc.WriteAttrStrZ("Colonia", m.Colonia)
	enc.WriteAttrStrZ("Localidad", m.Localidad)
	enc.WriteAttrStrZ("Referencia", m.Referencia)
	enc.WriteAttrStrZ("Municipio", m.Municipio)
	enc.WriteAttrStrZ("Estado", m.Estado)
	enc.WriteAttrStrZ("Pais", string(m.Pais))
	enc.WriteAttrStrZ("CodigoPostal", m.CodigoPostal)
}
