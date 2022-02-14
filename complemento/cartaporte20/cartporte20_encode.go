package cartaporte20

import (
	"bytes"

	"github.com/veyronifs/cfdi-go/encoder"
)

var cartaPorte20XS = encoder.NSElem{
	Prefix: "cartaporte20",
	NS:     "http://www.sat.gob.mx/CartaPorte20",
}

func (cp *CartaPorte20) SchemaLocation() string {
	return "http://www.sat.gob.mx/CartaPorte20 http://www.sat.gob.mx/sitio_internet/cfd/CartaPorte/CartaPorte20.xsd"
}

func (cp *CartaPorte20) XmlNSPrefix() string {
	return "cartaporte20"
}

func (cp *CartaPorte20) XmlNS() string {
	return "http://www.sat.gob.mx/CartaPorte20"
}

func Marshal(cp *CartaPorte20, moneda string) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	cp.MarshalComplemento(enc, moneda)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func (cp *CartaPorte20) MarshalComplemento(enc *encoder.Encoder, moneda string) {
	if cp == nil {
		return
	}
	enc.StartElem(cartaPorte20XS.Elem("CartaPorte"))
	defer enc.EndElem("CartaPorte")

	enc.WriteAttrStrZ("Version", cp.Version)
	enc.WriteAttrStrZ("TranspInternac", cp.TranspInternac)
	enc.WriteAttrStrZ("EntradaSalidaMerc", cp.EntradaSalidaMerc)
	enc.WriteAttrStrZ("PaisOrigenDestino", string(cp.PaisOrigenDestino))
	enc.WriteAttrStrZ("ViaEntradaSalida", cp.ViaEntradaSalida)
	enc.WriteAttrDecimalZ("TotalDistRec", cp.TotalDistRec, 2)

	cp.encodeUbicaciones(enc)
	cp.encodeMercancias(enc)
	cp.encodeFiguraTransporte(enc, cp.FiguraTransporte)

}

func (cp *CartaPorte20) encodeUbicaciones(enc *encoder.Encoder) {
	enc.StartElem(cartaPorte20XS.Elem("Ubicaciones"))
	defer enc.EndElem("Ubicaciones")
	for _, u := range cp.Ubicaciones {
		cp.encodeUbicacionesUbicacion(enc, u)
	}
}
func (cp *CartaPorte20) encodeUbicacionesUbicacion(enc *encoder.Encoder, u *Ubicacion) {
	enc.StartElem(cartaPorte20XS.Elem("Ubicacion"))
	defer enc.EndElem("Ubicacion")
	enc.WriteAttrStrZ("TipoUbicacion", u.TipoUbicacion)
	enc.WriteAttrStrZ("IDUbicacion", u.IDUbicacion)
	enc.WriteAttrStrZ("RFCRemitenteDestinatario", u.RFCRemitenteDestinatario)
	enc.WriteAttrStrZ("NombreRemitenteDestinatario", u.NombreRemitenteDestinatario)
	enc.WriteAttrStrZ("NumRegIdTrib", u.NumRegIdTrib)
	enc.WriteAttrStrZ("ResidenciaFiscal", string(u.ResidenciaFiscal))
	enc.WriteAttrStrZ("NumEstacion", u.NumEstacion)
	enc.WriteAttrStrZ("NombreEstacion", u.NombreEstacion)
	enc.WriteAttrStrZ("NavegacionTrafico", u.NavegacionTrafico)
	enc.WriteAttrStrZ("FechaHoraSalidaLlegada", u.FechaHoraSalidaLlegada.Encode())
	enc.WriteAttrStrZ("TipoEstacion", u.TipoEstacion)
	enc.WriteAttrDecimalZ("DistanciaRecorrida", u.DistanciaRecorrida, 2)

	if u.Domicilio == nil {
		return
	}
	enc.StartElem(cartaPorte20XS.Elem("Domicilio"))
	defer enc.EndElem("Domicilio")

	enc.WriteAttrStrZ("Calle", u.Domicilio.Calle)
	enc.WriteAttrStrZ("NumeroExterior", u.Domicilio.NumeroExterior)
	enc.WriteAttrStrZ("NumeroInterior", u.Domicilio.NumeroInterior)
	enc.WriteAttrStrZ("Colonia", u.Domicilio.Colonia)
	enc.WriteAttrStrZ("Localidad", u.Domicilio.Localidad)
	enc.WriteAttrStrZ("Referencia", u.Domicilio.Referencia)
	enc.WriteAttrStrZ("Municipio", u.Domicilio.Municipio)
	enc.WriteAttrStrZ("Estado", u.Domicilio.Estado)
	enc.WriteAttrStrZ("Pais", string(u.Domicilio.Pais))
	enc.WriteAttrStrZ("CodigoPostal", u.Domicilio.CodigoPostal)
}

func (cp *CartaPorte20) encodeMercancias(enc *encoder.Encoder) {
	enc.StartElem(cartaPorte20XS.Elem("Mercancias"))
	defer enc.EndElem("Mercancias")

	enc.WriteAttrDecimalZ("PesoBrutoTotal", cp.Mercancias.PesoBrutoTotal, 3)
	enc.WriteAttrStrZ("UnidadPeso", cp.Mercancias.UnidadPeso)
	enc.WriteAttrDecimalZ("PesoNetoTotal", cp.Mercancias.PesoNetoTotal, 3)
	enc.WriteAttrInt("NumTotalMercancias", cp.Mercancias.NumTotalMercancias)
	enc.WriteAttrDecimalZ("CargoPorTasacion", cp.Mercancias.CargoPorTasacion, 2)

	for _, m := range cp.Mercancias.Mercancia {
		cp.encodeMercancia(enc, m)
	}

	cp.encodeAutotransporte(enc, cp.Mercancias.Autotransporte)
	cp.encodeTransporteMaritimo(enc, cp.Mercancias.TransporteMaritimo)
	cp.encodeTransporteAereo(enc, cp.Mercancias.TransporteAereo)
	cp.encodeTransporteFerroviario(enc, cp.Mercancias.TransporteFerroviario)
}

func (cp *CartaPorte20) encodeMercancia(enc *encoder.Encoder, m *Mercancia) {
	enc.StartElem(cartaPorte20XS.Elem("Mercancia"))
	defer enc.EndElem("Mercancia")

	enc.WriteAttrStrZ("BienesTransp", m.BienesTransp)
	enc.WriteAttrStrZ("ClaveSTCC", m.ClaveSTCC)
	enc.WriteAttrStrZ("Descripcion", m.Descripcion)
	enc.WriteAttrDecimalZ("Cantidad", m.Cantidad, 6)
	enc.WriteAttrStrZ("ClaveUnidad", m.ClaveUnidad)
	enc.WriteAttrStrZ("Unidad", m.Unidad)
	enc.WriteAttrStrZ("Dimensiones", m.Dimensiones)
	enc.WriteAttrStrZ("MaterialPeligroso", m.MaterialPeligroso)
	enc.WriteAttrStrZ("CveMaterialPeligroso", m.CveMaterialPeligroso)
	enc.WriteAttrStrZ("Embalaje", m.Embalaje)
	enc.WriteAttrStrZ("DescripEmbalaje", m.DescripEmbalaje)
	enc.WriteAttrDecimalZ("PesoEnKg", m.PesoEnKg, 3)
	enc.WriteAttrDecimalZ("ValorMercancia", m.ValorMercancia, 2)
	enc.WriteAttrStrZ("Moneda", string(m.Moneda))
	enc.WriteAttrStrZ("FraccionArancelaria", m.FraccionArancelaria)
	enc.WriteAttrStrZ("UUIDComercioExt", m.UUIDComercioExt)

	for _, p := range m.Pedimentos {
		enc.StartElem(cartaPorte20XS.Elem("Pedimentos"))
		enc.WriteAttrStrZ("Pedimento", p.Pedimento)
		enc.EndElem("Pedimentos")
	}

	for _, g := range m.GuiasIdentificacion {
		enc.StartElem(cartaPorte20XS.Elem("GuiasIdentificacion"))
		enc.WriteAttrStrZ("NumeroGuiaIdentificacion", g.NumeroGuiaIdentificacion)
		enc.WriteAttrStrZ("DescripGuiaIdentificacion", g.DescripGuiaIdentificacion)
		enc.WriteAttrDecimalZ("PesoGuiaIdentificacion", g.PesoGuiaIdentificacion, 3)
		enc.EndElem("GuiasIdentificacion")
	}

	for _, c := range m.CantidadTransporta {
		enc.StartElem(cartaPorte20XS.Elem("CantidadTransporta"))
		enc.WriteAttrDecimalZ("Cantidad", c.Cantidad, 6)
		enc.WriteAttrStrZ("IDOrigen", c.IDOrigen)
		enc.WriteAttrStrZ("IDDestino", c.IDDestino)
		enc.WriteAttrStrZ("CvesTransporte", c.CvesTransporte)
		enc.EndElem("CantidadTransporta")
	}

	if m.DetalleMercancia != nil {
		enc.StartElem(cartaPorte20XS.Elem("DetalleMercancia"))
		enc.WriteAttrStrZ("UnidadPesoMerc", m.DetalleMercancia.UnidadPesoMerc)
		enc.WriteAttrDecimalZ("PesoBruto", m.DetalleMercancia.PesoBruto, 3)
		enc.WriteAttrDecimalZ("PesoNeto", m.DetalleMercancia.PesoNeto, 3)
		enc.WriteAttrDecimalZ("PesoTara", m.DetalleMercancia.PesoTara, 3)
		enc.WriteAttrInt("NumPiezas", m.DetalleMercancia.NumPiezas)
		enc.EndElem("DetalleMercancia")
	}
}

func (cp *CartaPorte20) encodeAutotransporte(enc *encoder.Encoder, at *Autotransporte) {
	if at == nil {
		return
	}
	enc.StartElem(cartaPorte20XS.Elem("Autotransporte"))
	defer enc.EndElem("Autotransporte")

	enc.WriteAttrStrZ("PermSCT", at.PermSCT)
	enc.WriteAttrStrZ("NumPermisoSCT", at.NumPermisoSCT)
	if idV := at.IdentificacionVehicular; idV != nil {
		enc.StartElem(cartaPorte20XS.Elem("IdentificacionVehicular"))
		enc.WriteAttrStrZ("ConfigVehicular", idV.ConfigVehicular)
		enc.WriteAttrStrZ("PlacaVM", idV.PlacaVM)
		enc.WriteAttrStrZ("AnioModeloVM", idV.AnioModeloVM)
		enc.EndElem("IdentificacionVehicular")
	}

	if seg := at.Seguros; seg != nil {
		enc.StartElem(cartaPorte20XS.Elem("Seguros"))
		enc.WriteAttrStrZ("AseguraRespCivil", seg.AseguraRespCivil)
		enc.WriteAttrStrZ("PolizaRespCivil", seg.PolizaRespCivil)
		enc.WriteAttrStrZ("AseguraMedAmbiente", seg.AseguraMedAmbiente)
		enc.WriteAttrStrZ("PolizaMedAmbiente", seg.PolizaMedAmbiente)
		enc.WriteAttrStrZ("AseguraCarga", seg.AseguraCarga)
		enc.WriteAttrStrZ("PolizaCarga", seg.PolizaCarga)
		enc.WriteAttrDecimal("PrimaSeguro", seg.PrimaSeguro, 2)
		enc.EndElem("Seguros")
	}

	if len(at.Remolques) > 0 {
		enc.StartElem(cartaPorte20XS.Elem("Remolques"))
		for _, rem := range at.Remolques {
			enc.StartElem(cartaPorte20XS.Elem("Remolque"))
			enc.WriteAttrStrZ("SubTipoRem", rem.SubTipoRem)
			enc.WriteAttrStrZ("Placa", rem.Placa)
			enc.EndElem("Remolque")
		}
		enc.EndElem("Remolques")
	}
}

func (cp *CartaPorte20) encodeTransporteMaritimo(enc *encoder.Encoder, tm *TransporteMaritimo) {
	if tm == nil {
		return
	}
	panic("not implemented")
}
func (cp *CartaPorte20) encodeTransporteAereo(enc *encoder.Encoder, ta *TransporteAereo) {
	if ta == nil {
		return
	}
	panic("not implemented")
}
func (cp *CartaPorte20) encodeTransporteFerroviario(enc *encoder.Encoder, tf *TransporteFerroviario) {
	if tf == nil {
		return
	}
	panic("not implemented")
}

func (cp *CartaPorte20) encodeFiguraTransporte(enc *encoder.Encoder, ft *FiguraTransporte) {
	if ft == nil {
		return
	}
	enc.StartElem(cartaPorte20XS.Elem("FiguraTransporte"))
	defer enc.EndElem("FiguraTransporte")
	for _, f := range ft.TiposFigura {
		cp.encodeFiguraTransporteTiposFigura(enc, f)
	}
}

func (cp *CartaPorte20) encodeFiguraTransporteTiposFigura(enc *encoder.Encoder, ft *TiposFigura) {
	if ft == nil {
		return
	}
	enc.StartElem(cartaPorte20XS.Elem("TiposFigura"))
	defer enc.EndElem("TiposFigura")

	enc.WriteAttrStrZ("TipoFigura", ft.TipoFigura)
	enc.WriteAttrStrZ("RFCFigura", ft.RFCFigura)
	enc.WriteAttrStrZ("NumLicencia", ft.NumLicencia)
	enc.WriteAttrStrZ("NombreFigura", ft.NombreFigura)
	enc.WriteAttrStrZ("NumRegIdTribFigura", ft.NumRegIdTribFigura)
	enc.WriteAttrStrZ("ResidenciaFiscalFigura", string(ft.ResidenciaFiscalFigura))
	for _, part := range ft.PartesTransporte {
		enc.StartElem(cartaPorte20XS.Elem("PartesTransporte"))
		enc.WriteAttrStrZ("ParteTransporte", part.ParteTransporte)
		enc.EndElem("PartesTransporte")
	}

	if dom := ft.Domicilio; dom != nil {
		enc.StartElem(cartaPorte20XS.Elem("Domicilio"))
		enc.WriteAttrStrZ("Calle", dom.Calle)
		enc.WriteAttrStrZ("NumeroExterior", dom.NumeroExterior)
		enc.WriteAttrStrZ("NumeroInterior", dom.NumeroInterior)
		enc.WriteAttrStrZ("Colonia", dom.Colonia)
		enc.WriteAttrStrZ("Localidad", dom.Localidad)
		enc.WriteAttrStrZ("Referencia", dom.Referencia)
		enc.WriteAttrStrZ("Municipio", dom.Municipio)
		enc.WriteAttrStrZ("Estado", dom.Estado)
		enc.WriteAttrStrZ("Pais", string(dom.Pais))
		enc.WriteAttrStrZ("CodigoPostal", dom.CodigoPostal)
		enc.EndElem("Domicilio")
	}
}
