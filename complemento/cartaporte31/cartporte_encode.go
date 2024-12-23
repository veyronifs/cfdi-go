package cartaporte31

import (
	"bytes"

	"github.com/veyronifs/cfdi-go/encoder"
)

var cartaporte31XS = encoder.NSElem{
	Prefix: "cartaporte31",
	NS:     "http://www.sat.gob.mx/CartaPorte31",
}

func (cp *CartaPorte31) SchemaLocation() string {
	return cartaporte31XS.NS + " http://www.sat.gob.mx/sitio_internet/cfd/CartaPorte/CartaPorte31.xsd"
}

func (cp *CartaPorte31) XmlNSPrefix() string {
	return cartaporte31XS.Prefix
}

func (cp *CartaPorte31) XmlNS() string {
	return cartaporte31XS.NS
}

func Marshal(cp *CartaPorte31, moneda string) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	cp.MarshalComplemento(enc, moneda)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func (cp *CartaPorte31) MarshalComplemento(enc *encoder.Encoder, moneda string) {
	if cp == nil {
		return
	}
	enc.StartElem(cartaporte31XS.Elem("CartaPorte"))
	defer enc.EndElem("CartaPorte")

	enc.WriteAttrStrZ("Version", cp.Version)
	enc.WriteAttrStrZ("IdCCP", cp.IdCCP)
	enc.WriteAttrStrZ("TranspInternac", cp.TranspInternac)
	enc.WriteAttrStrZ("RegimenAduanero", string(cp.RegimenAduanero))
	enc.WriteAttrStrZ("EntradaSalidaMerc", cp.EntradaSalidaMerc)
	enc.WriteAttrStrZ("PaisOrigenDestino", string(cp.PaisOrigenDestino))
	enc.WriteAttrStrZ("ViaEntradaSalida", cp.ViaEntradaSalida)
	enc.WriteAttrDecimalZ("TotalDistRec", cp.TotalDistRec, 2)
	enc.WriteAttrStrZ("RegistroISTMO", cp.RegistroISTMO)
	enc.WriteAttrStrZ("UbicacionPoloOrigen", cp.UbicacionPoloOrigen)
	enc.WriteAttrStrZ("UbicacionPoloDestino", cp.UbicacionPoloDestino)

	encodeUbicaciones(enc, cp.Ubicaciones)
	encodeMercancias(enc, cp.Mercancias)
	encodeFiguraTransporte(enc, cp.FiguraTransporte)

}

func encodeUbicaciones(enc *encoder.Encoder, ubicaciones []*Ubicacion) {
	enc.StartElem(cartaporte31XS.Elem("Ubicaciones"))
	defer enc.EndElem("Ubicaciones")
	for _, u := range ubicaciones {
		encodeUbicacionesUbicacion(enc, u)
	}
}
func encodeUbicacionesUbicacion(enc *encoder.Encoder, u *Ubicacion) {
	enc.StartElem(cartaporte31XS.Elem("Ubicacion"))
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
	enc.WriteAttrStrZ("FechaHoraSalidaLlegada", u.FechaHoraSalidaLlegada.String())
	enc.WriteAttrStrZ("TipoEstacion", u.TipoEstacion)
	enc.WriteAttrDecimalZ("DistanciaRecorrida", u.DistanciaRecorrida, 2)

	if u.Domicilio == nil {
		return
	}
	enc.StartElem(cartaporte31XS.Elem("Domicilio"))
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

func encodeMercancias(enc *encoder.Encoder, mercancias *Mercancias) {
	if mercancias == nil {
		return
	}
	enc.StartElem(cartaporte31XS.Elem("Mercancias"))
	defer enc.EndElem("Mercancias")

	enc.WriteAttrDecimalZ("PesoBrutoTotal", mercancias.PesoBrutoTotal, 3)
	enc.WriteAttrStrZ("UnidadPeso", mercancias.UnidadPeso)
	enc.WriteAttrDecimalZ("PesoNetoTotal", mercancias.PesoNetoTotal, 3)
	enc.WriteAttrInt("NumTotalMercancias", mercancias.NumTotalMercancias)
	enc.WriteAttrDecimalZ("CargoPorTasacion", mercancias.CargoPorTasacion, 2)
	enc.WriteAttrStrZ("LogisticaInversaRecoleccionDevolucion", mercancias.LogisticaInversaRecoleccionDevolucion)

	for _, m := range mercancias.Mercancia {
		encodeMercancia(enc, m)
	}

	encodeAutotransporte(enc, mercancias.Autotransporte)
	encodeTransporteMaritimo(enc, mercancias.TransporteMaritimo)
	encodeTransporteAereo(enc, mercancias.TransporteAereo)
	encodeTransporteFerroviario(enc, mercancias.TransporteFerroviario)
}

func encodeMercancia(enc *encoder.Encoder, m *Mercancia) {
	enc.StartElem(cartaporte31XS.Elem("Mercancia"))
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
	enc.WriteAttrStrZ("SectorCOFEPRIS", string(m.SectorCOFEPRIS))
	enc.WriteAttrStrZ("NombreIngredienteActivo", m.NombreIngredienteActivo)
	enc.WriteAttrStrZ("NomQuimico", m.NomQuimico)
	enc.WriteAttrStrZ("DenominacionGenericaProd", m.DenominacionGenericaProd)
	enc.WriteAttrStrZ("DenominacionDistintivaProd", m.DenominacionDistintivaProd)
	enc.WriteAttrStrZ("Fabricante", m.Fabricante)
	enc.WriteAttrStrZ("FechaCaducidad", m.FechaCaducidad)
	enc.WriteAttrStrZ("LoteMedicamento", m.LoteMedicamento)
	enc.WriteAttrStrZ("FormaFarmaceutica", string(m.FormaFarmaceutica))
	enc.WriteAttrStrZ("CondicionesEspTransp", string(m.CondicionesEspTransp))
	enc.WriteAttrStrZ("RegistroSanitarioFolioAutorizacion", m.RegistroSanitarioFolioAutorizacion)
	enc.WriteAttrStrZ("PermisoImportacion", m.PermisoImportacion)
	enc.WriteAttrStrZ("FolioImpoVUCEM", m.FolioImpoVUCEM)
	enc.WriteAttrStrZ("NumCAS", m.NumCAS)
	enc.WriteAttrStrZ("RazonSocialEmpImp", m.RazonSocialEmpImp)
	enc.WriteAttrStrZ("NumRegSanPlagCOFEPRIS", m.NumRegSanPlagCOFEPRIS)
	enc.WriteAttrStrZ("DatosFabricante", m.DatosFabricante)
	enc.WriteAttrStrZ("DatosFormulador", m.DatosFormulador)
	enc.WriteAttrStrZ("DatosMaquilador", m.DatosMaquilador)
	enc.WriteAttrStrZ("UsoAutorizado", m.UsoAutorizado)
	enc.WriteAttrDecimal("PesoEnKg", m.PesoEnKg, 0)
	enc.WriteAttrDecimalZ("ValorMercancia", m.ValorMercancia, 2)
	enc.WriteAttrStrZ("Moneda", string(m.Moneda))
	enc.WriteAttrStrZ("FraccionArancelaria", m.FraccionArancelaria)
	enc.WriteAttrStrZ("UUIDComercioExt", m.UUIDComercioExt)
	enc.WriteAttrStrZ("TipoMateria", m.TipoMateria)
	enc.WriteAttrStrZ("DescripcionMateria", m.DescripcionMateria)

	for _, p := range m.DocumentacionAduanera {
		enc.StartElem(cartaporte31XS.Elem("DocumentacionAduanera"))
		enc.WriteAttrStrZ("TipoDocumento", string(p.TipoDocumento))
		enc.WriteAttrStrZ("NumPedimento", p.NumPedimento)
		enc.WriteAttrStrZ("IdentDocAduanero", p.IdentDocAduanero)
		enc.WriteAttrStrZ("RFCImpo", p.RFCImpo)
		enc.EndElem("DocumentacionAduanera")
	}

	for _, g := range m.GuiasIdentificacion {
		enc.StartElem(cartaporte31XS.Elem("GuiasIdentificacion"))
		enc.WriteAttrStrZ("NumeroGuiaIdentificacion", g.NumeroGuiaIdentificacion)
		enc.WriteAttrStrZ("DescripGuiaIdentificacion", g.DescripGuiaIdentificacion)
		enc.WriteAttrDecimalZ("PesoGuiaIdentificacion", g.PesoGuiaIdentificacion, 3)
		enc.EndElem("GuiasIdentificacion")
	}

	for _, c := range m.CantidadTransporta {
		enc.StartElem(cartaporte31XS.Elem("CantidadTransporta"))
		enc.WriteAttrDecimalZ("Cantidad", c.Cantidad, 6)
		enc.WriteAttrStrZ("IDOrigen", c.IDOrigen)
		enc.WriteAttrStrZ("IDDestino", c.IDDestino)
		enc.WriteAttrStrZ("CvesTransporte", c.CvesTransporte)
		enc.EndElem("CantidadTransporta")
	}

	if m.DetalleMercancia != nil {
		enc.StartElem(cartaporte31XS.Elem("DetalleMercancia"))
		enc.WriteAttrStrZ("UnidadPesoMerc", m.DetalleMercancia.UnidadPesoMerc)
		enc.WriteAttrDecimalZ("PesoBruto", m.DetalleMercancia.PesoBruto, 3)
		enc.WriteAttrDecimalZ("PesoNeto", m.DetalleMercancia.PesoNeto, 3)
		enc.WriteAttrDecimalZ("PesoTara", m.DetalleMercancia.PesoTara, 3)
		enc.WriteAttrInt("NumPiezas", m.DetalleMercancia.NumPiezas)
		enc.EndElem("DetalleMercancia")
	}
}

func encodeAutotransporte(enc *encoder.Encoder, at *Autotransporte) {
	if at == nil {
		return
	}
	enc.StartElem(cartaporte31XS.Elem("Autotransporte"))
	defer enc.EndElem("Autotransporte")

	enc.WriteAttrStrZ("PermSCT", at.PermSCT)
	enc.WriteAttrStrZ("NumPermisoSCT", at.NumPermisoSCT)
	if idV := at.IdentificacionVehicular; idV != nil {
		enc.StartElem(cartaporte31XS.Elem("IdentificacionVehicular"))
		enc.WriteAttrStrZ("ConfigVehicular", idV.ConfigVehicular)
		enc.WriteAttrDecimal("PesoBrutoVehicular", idV.PesoBrutoVehicular, 2)
		enc.WriteAttrStrZ("PlacaVM", idV.PlacaVM)
		enc.WriteAttrStrZ("AnioModeloVM", idV.AnioModeloVM)
		enc.EndElem("IdentificacionVehicular")
	}

	if seg := at.Seguros; seg != nil {
		enc.StartElem(cartaporte31XS.Elem("Seguros"))
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
		enc.StartElem(cartaporte31XS.Elem("Remolques"))
		for _, rem := range at.Remolques {
			enc.StartElem(cartaporte31XS.Elem("Remolque"))
			enc.WriteAttrStrZ("SubTipoRem", rem.SubTipoRem)
			enc.WriteAttrStrZ("Placa", rem.Placa)
			enc.EndElem("Remolque")
		}
		enc.EndElem("Remolques")
	}
}

func encodeTransporteMaritimo(enc *encoder.Encoder, tm *TransporteMaritimo) {
	if tm == nil {
		return
	}
	panic("not implemented")
}
func encodeTransporteAereo(enc *encoder.Encoder, ta *TransporteAereo) {
	if ta == nil {
		return
	}
	panic("not implemented")
}
func encodeTransporteFerroviario(enc *encoder.Encoder, tf *TransporteFerroviario) {
	if tf == nil {
		return
	}
	panic("not implemented")
}

func encodeFiguraTransporte(enc *encoder.Encoder, ft *FiguraTransporte) {
	if ft == nil {
		return
	}
	enc.StartElem(cartaporte31XS.Elem("FiguraTransporte"))
	defer enc.EndElem("FiguraTransporte")
	for _, f := range ft.TiposFigura {
		encodeFiguraTransporteTiposFigura(enc, f)
	}
}

func encodeFiguraTransporteTiposFigura(enc *encoder.Encoder, ft *TiposFigura) {
	if ft == nil {
		return
	}
	enc.StartElem(cartaporte31XS.Elem("TiposFigura"))
	defer enc.EndElem("TiposFigura")

	enc.WriteAttrStrZ("TipoFigura", ft.TipoFigura)
	enc.WriteAttrStrZ("RFCFigura", ft.RFCFigura)
	enc.WriteAttrStrZ("NumLicencia", ft.NumLicencia)
	enc.WriteAttrStrZ("NombreFigura", ft.NombreFigura)
	enc.WriteAttrStrZ("NumRegIdTribFigura", ft.NumRegIdTribFigura)
	enc.WriteAttrStrZ("ResidenciaFiscalFigura", string(ft.ResidenciaFiscalFigura))
	for _, part := range ft.PartesTransporte {
		enc.StartElem(cartaporte31XS.Elem("PartesTransporte"))
		enc.WriteAttrStrZ("ParteTransporte", part.ParteTransporte)
		enc.EndElem("PartesTransporte")
	}

	if dom := ft.Domicilio; dom != nil {
		enc.StartElem(cartaporte31XS.Elem("Domicilio"))
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
