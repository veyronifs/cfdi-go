package comext11

import (
	"bytes"
	"strconv"

	"github.com/shabbyrobe/xmlwriter"
	"github.com/veyronifs/cfdi-go/encoder"
)

var comext11XS = encoder.NSElem{
	Prefix: "cce11",
	NS:     "http://www.sat.gob.mx/ComercioExterior11",
}

func (ce *ComercioExterior) SchemaLocation() string {
	return comext11XS.NS + " http://www.sat.gob.mx/sitio_internet/cfd/ComercioExterior11/ComercioExterior11.xsd"
}

func (ce *ComercioExterior) XmlNSPrefix() string {
	return comext11XS.Prefix
}

func (ce *ComercioExterior) XmlNS() string {
	return comext11XS.NS
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
	enc.StartElem(comext11XS.ElemXS("ComercioExterior"))
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
			Value:  "http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd",
		},
	)

	enc.WriteAttrStr("Version", ce.Version)
	enc.WriteAttrStr("MotivoTraslado", ce.MotivoTraslado)
	enc.WriteAttrStr("TipoOperacion", ce.TipoOperacion)
	enc.WriteAttrStr("ClaveDePedimento", ce.ClaveDePedimento)
	enc.WriteAttrStr("CertificadoOrigen", strconv.Itoa(ce.CertificadoOrigen))
	enc.WriteAttrStr("NumCertificadoOrigen", ce.NumCertificadoOrigen)
	enc.WriteAttrStr("NumeroExportadorConfiable", ce.NumeroExportadorConfiable)
	enc.WriteAttrStrZ("Incoterm", ce.Incoterm)
	enc.WriteAttrStr("Subdivision", strconv.Itoa(ce.Subdivision))
	enc.WriteAttrStr("Observaciones", ce.Observaciones)
	enc.WriteAttrDecimalZ("TipoCambioUSD", ce.TipoCambioUSD, 2)
	enc.WriteAttrDecimalZ("TotalUSD", ce.TotalUSD, 2)

	ce.encodeMercancias(enc)
	ce.encodeEmisor(enc)

	for _, m := range ce.Propietarios {
		enc.WriteAttrStrZ("NumRegIdTrib", m.NumRegIdTrib)
		enc.WriteAttrStrZ("ResidenciaFiscal", string(m.ResidenciaFiscal))
	}

	for _, m := range ce.Destinatarios {
		enc.WriteAttrStrZ("NumRegIdTrib", m.NumRegIdTrib)
		enc.WriteAttrStrZ("Nombre", m.Nombre)
		for _, domicilio := range m.Domicilios {
			ce.encodeDomicilio(enc, domicilio)
		}
	}
}

func (ce *ComercioExterior) encodeMercancias(enc *encoder.Encoder) {
	enc.StartElem(comext11XS.Elem("Mercancias"))
	defer enc.EndElem("Mercancias")
	for _, m := range ce.Mercancias {
		ce.encodeMercanciasMercancia(enc, m)
	}
}

func (ce *ComercioExterior) encodeMercanciasMercancia(enc *encoder.Encoder, m *Mercancia) {
	enc.StartElem(comext11XS.Elem("Mercancia"))
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

func (ce *ComercioExterior) encodeEmisor(enc *encoder.Encoder) {
	enc.StartElem(comext11XS.Elem("Emisor"))
	defer enc.EndElem("Emisor")

	enc.WriteAttrStrZ("Curp", ce.Emisor.Curp)
	ce.encodeDomicilio(enc, ce.Emisor.Domicilio)
}

func (ce *ComercioExterior) encodeReceptor(enc *encoder.Encoder) {
	enc.StartElem(comext11XS.Elem("Receptor"))
	defer enc.EndElem("Receptor")

	enc.WriteAttrStrZ("NumRegIdTrib", ce.Receptor.NumRegIdTrib)
	ce.encodeDomicilio(enc, ce.Receptor.Domicilio)
}

func (ce *ComercioExterior) encodeDomicilio(enc *encoder.Encoder, m *Domicilio) {
	enc.StartElem(comext11XS.Elem("Domicilio"))
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
