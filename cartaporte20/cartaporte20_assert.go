package cartaporte20

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func assertDecimal(t *testing.T, ex, act decimal.Decimal, path ...any) {
	msg := fmt.Sprintf(path[0].(string), path[1:]...)
	assert.True(t, ex.Equal(act), msg+" %s != %s", ex.String(), act.String())
}

func AssertEqual(t *testing.T, cp1, cp2 *CartaPorte20) {
	if cp1 == nil {
		assert.Nil(t, cp2, "CartaPorte20")
		return
	}
	assert.NotNil(t, cp2, "CartaPorte20")
	if cp2 == nil {
		return
	}
	assert.Equal(t, cp1.Version, cp2.Version, ".CartaPorte20.Version")
	assert.Equal(t, cp1.TranspInternac, cp2.TranspInternac, ".CartaPorte20.TranspInternac")
	assert.Equal(t, cp1.EntradaSalidaMerc, cp2.EntradaSalidaMerc, ".CartaPorte20.EntradaSalidaMerc")
	assert.Equal(t, cp1.PaisOrigenDestino, cp2.PaisOrigenDestino, ".CartaPorte20.PaisOrigenDestino")
	assert.Equal(t, cp1.ViaEntradaSalida, cp2.ViaEntradaSalida, ".CartaPorte20.ViaEntradaSalida")
	assertDecimal(t, cp1.TotalDistRec, cp2.TotalDistRec, ".CartaPorte20.TotalDistRec")

	l1, l2 := len(cp1.Ubicaciones), len(cp2.Ubicaciones)
	assert.Equal(t, l1, l2, ".CartaPorte20.Ubicaciones len %d != %d", l1, l2)
	if l1 == l2 {
		for i, ub1 := range cp1.Ubicaciones {
			ub2 := cp2.Ubicaciones[i]
			assertUbicaciones(t, ub1, ub2, fmt.Sprintf(".CartaPorte20.Ubicaciones[%d]", i))
		}
	}

	assertMercancias(t, cp1.Mercancias, cp2.Mercancias, ".CartaPorte20.Mercancias")
	assertFiguraTransporte(t, cp1.FiguraTransporte, cp2.FiguraTransporte, ".CartaPorte20.FiguraTransporte")
}

func assertUbicaciones(t *testing.T, ub1, ub2 *Ubicacion, path string) {
	if ub1 == nil {
		assert.Nil(t, ub2, path)
		return
	}
	assert.NotNil(t, ub2, path)
	if ub2 == nil {
		return
	}
	assert.Equal(t, ub1.TipoUbicacion, ub2.TipoUbicacion, path+"TipoUbicacion")
	assert.Equal(t, ub1.IDUbicacion, ub2.IDUbicacion, path+"IDUbicacion")
	assert.Equal(t, ub1.RFCRemitenteDestinatario, ub2.RFCRemitenteDestinatario, path+"RFCRemitenteDestinatario")
	assert.Equal(t, ub1.NombreRemitenteDestinatario, ub2.NombreRemitenteDestinatario, path+"NombreRemitenteDestinatario")
	assert.Equal(t, ub1.NumRegIdTrib, ub2.NumRegIdTrib, path+"NumRegIdTrib")
	assert.Equal(t, ub1.ResidenciaFiscal, ub2.ResidenciaFiscal, path+"ResidenciaFiscal")
	assert.Equal(t, ub1.NumEstacion, ub2.NumEstacion, path+"NumEstacion")
	assert.Equal(t, ub1.NombreEstacion, ub2.NombreEstacion, path+"NombreEstacion")
	assert.Equal(t, ub1.NavegacionTrafico, ub2.NavegacionTrafico, path+"NavegacionTrafico")
	assert.Equal(t, ub1.FechaHoraSalidaLlegada, ub2.FechaHoraSalidaLlegada, path+"FechaHoraSalidaLlegada")
	assert.Equal(t, ub1.TipoEstacion, ub2.TipoEstacion, path+"TipoEstacion")
	if ub1.Domicilio == nil {
		assert.Nil(t, ub2.Domicilio, path+".Domicilio")
		return
	}
	assert.NotNil(t, ub2.Domicilio, path+".Domicilio")
	if ub2.Domicilio == nil {
		return
	}

	assert.Equal(t, ub1.Domicilio.Calle, ub2.Domicilio.Calle, path+".Domicilio.Calle")
	assert.Equal(t, ub1.Domicilio.NumeroExterior, ub2.Domicilio.NumeroExterior, path+".Domicilio.NumeroExterior")
	assert.Equal(t, ub1.Domicilio.NumeroInterior, ub2.Domicilio.NumeroInterior, path+".Domicilio.NumeroInterior")
	assert.Equal(t, ub1.Domicilio.Colonia, ub2.Domicilio.Colonia, path+".Domicilio.Colonia")
	assert.Equal(t, ub1.Domicilio.Localidad, ub2.Domicilio.Localidad, path+".Domicilio.Localidad")
	assert.Equal(t, ub1.Domicilio.Referencia, ub2.Domicilio.Referencia, path+".Domicilio.Referencia")
	assert.Equal(t, ub1.Domicilio.Municipio, ub2.Domicilio.Municipio, path+".Domicilio.Municipio")
	assert.Equal(t, ub1.Domicilio.Estado, ub2.Domicilio.Estado, path+".Domicilio.Estado")
	assert.Equal(t, ub1.Domicilio.Pais, ub2.Domicilio.Pais, path+".Domicilio.Pais")
	assert.Equal(t, ub1.Domicilio.CodigoPostal, ub2.Domicilio.CodigoPostal, path+".Domicilio.CodigoPostal")
}

func assertMercancias(t *testing.T, v1, v2 *Mercancias, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	l1, l2 := len(v1.Mercancia), len(v2.Mercancia)
	assert.Equal(t, l1, l2, path+".Mercancia len %d != %d", l1, l2)
	if l1 == l2 {
		for i, m1 := range v1.Mercancia {
			m2 := v2.Mercancia[i]
			assertMercancia(t, m1, m2, fmt.Sprintf(".Mercancia[%d]", i))
		}
	}
	assertAutotransporte(t, v1.Autotransporte, v2.Autotransporte, path+".Autotransporte")
	assertTransporteMaritimo(t, v1.TransporteMaritimo, v2.TransporteMaritimo, path+".TransporteMaritimo")
	assertTransporteAereo(t, v1.TransporteAereo, v2.TransporteAereo, path+".TransporteAereo")
	assertTransporteFerroviario(t, v1.TransporteFerroviario, v2.TransporteFerroviario, path+".TransporteFerroviario")
}

func assertMercancia(t *testing.T, v1, v2 *Mercancia, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.BienesTransp, v2.BienesTransp, path+".BienesTransp")
	assert.Equal(t, v1.ClaveSTCC, v2.ClaveSTCC, path+".ClaveSTCC")
	assert.Equal(t, v1.Descripcion, v2.Descripcion, path+".Descripcion")
	assertDecimal(t, v1.Cantidad, v2.Cantidad, path+".Cantidad")
	assert.Equal(t, v1.ClaveUnidad, v2.ClaveUnidad, path+".ClaveUnidad")
	assert.Equal(t, v1.Unidad, v2.Unidad, path+".Unidad")
	assert.Equal(t, v1.Dimensiones, v2.Dimensiones, path+".Dimensiones")
	assert.Equal(t, v1.MaterialPeligroso, v2.MaterialPeligroso, path+".MaterialPeligroso")
	assert.Equal(t, v1.CveMaterialPeligroso, v2.CveMaterialPeligroso, path+".CveMaterialPeligroso")
	assert.Equal(t, v1.Embalaje, v2.Embalaje, path+".Embalaje")
	assert.Equal(t, v1.DescripEmbalaje, v2.DescripEmbalaje, path+".DescripEmbalaje")
	assertDecimal(t, v1.PesoEnKg, v2.PesoEnKg, path+".PesoEnKg")
	assertDecimal(t, v1.ValorMercancia, v2.ValorMercancia, path+".ValorMercancia")
	assert.Equal(t, v1.Moneda, v2.Moneda, path+".Moneda")
	assert.Equal(t, v1.FraccionArancelaria, v2.FraccionArancelaria, path+".FraccionArancelaria")
	assert.Equal(t, v1.UUIDComercioExt, v2.UUIDComercioExt, path+".UUIDComercioExt")

	l1, l2 := len(v1.Pedimentos), len(v2.Pedimentos)
	assert.Equal(t, l1, l2, path+".Pedimentos len %d != %d", l1, l2)
	if l1 == l2 {
		for i, p1 := range v1.Pedimentos {
			p2 := v2.Pedimentos[i]
			if p1 == nil || p2 == nil {
				assert.Nil(t, p2, path+".Pedimentos[%d]", i)
				assert.Nil(t, p2, path+".Pedimentos[%d]", i)
				continue
			}
			assert.Equal(t, p1.Pedimento, p2.Pedimento, path+".Pedimentos[%d].Pedimento", i)
		}
	}

	l1, l2 = len(v1.GuiasIdentificacion), len(v2.GuiasIdentificacion)
	assert.Equal(t, l1, l2, path+".GuiasIdentificacion len %d != %d", l1, l2)
	if l1 == l2 {
		for i, g1 := range v1.GuiasIdentificacion {
			g2 := v2.GuiasIdentificacion[i]
			if g1 == nil || g2 == nil {
				assert.Nil(t, g2, path+".GuiasIdentificacion[%d]", i)
				assert.Nil(t, g2, path+".GuiasIdentificacion[%d]", i)
				continue
			}
			assert.Equal(t, g1.NumeroGuiaIdentificacion, g2.NumeroGuiaIdentificacion, path+".GuiasIdentificacion[%d].NumeroGuiaIdentificacion", i)
			assert.Equal(t, g1.DescripGuiaIdentificacion, g2.DescripGuiaIdentificacion, path+".GuiasIdentificacion[%d].DescripGuiaIdentificacion", i)
		}
	}
	l1, l2 = len(v1.CantidadTransporta), len(v2.CantidadTransporta)
	assert.Equal(t, l1, l2, path+".CantidadTransporta len %d != %d", l1, l2)
	if l1 == l2 {
		for i, c1 := range v1.CantidadTransporta {
			c2 := v2.CantidadTransporta[i]
			if c1 == nil || c2 == nil {
				assert.Nil(t, c2, path+".CantidadTransporta[%d]", i)
				assert.Nil(t, c2, path+".CantidadTransporta[%d]", i)
				continue
			}
			assertDecimal(t, c1.Cantidad, c2.Cantidad, path+".CantidadTransporta[%d].Cantidad", i)
			assert.Equal(t, c1.IDOrigen, c2.IDOrigen, path+".CantidadTransporta[%d].IDOrigen", i)
			assert.Equal(t, c1.IDDestino, c2.IDDestino, path+".CantidadTransporta[%d].IDDestino", i)
			assert.Equal(t, c1.CvesTransporte, c2.CvesTransporte, path+".CantidadTransporta[%d].CvesTransporte", i)
		}
	}

	if v1.DetalleMercancia == nil {
		assert.Nil(t, v2.DetalleMercancia, path+".DetalleMercancia")
		return
	}
	assert.NotNil(t, v2.DetalleMercancia, path+".DetalleMercancia")
	if v2.DetalleMercancia == nil {
		return
	}
	assert.Equal(t, v1.DetalleMercancia.UnidadPesoMerc, v2.DetalleMercancia.UnidadPesoMerc, path+".DetalleMercancia.UnidadPesoMerc")
	assertDecimal(t, v1.DetalleMercancia.PesoBruto, v2.DetalleMercancia.PesoBruto, path+".DetalleMercancia.PesoBruto")
	assertDecimal(t, v1.DetalleMercancia.PesoNeto, v2.DetalleMercancia.PesoNeto, path+".DetalleMercancia.PesoNeto")
	assertDecimal(t, v1.DetalleMercancia.PesoTara, v2.DetalleMercancia.PesoTara, path+".DetalleMercancia.PesoTara")
	assert.Equal(t, v1.DetalleMercancia.NumPiezas, v2.DetalleMercancia.NumPiezas, path+".DetalleMercancia.NumPiezas")

}

func assertAutotransporte(t *testing.T, v1, v2 *Autotransporte, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assertAutotransporteIdentificacionVehicular(t, v1.IdentificacionVehicular, v2.IdentificacionVehicular, path+".IdentificacionVehicular")
	assertAutotransporteSeguros(t, v1.Seguros, v2.Seguros, path+".Seguros")
	assertAutotransporteRemolques(t, v1.Remolques, v2.Remolques, path+".Remolques")
	assert.Equal(t, v1.PermSCT, v2.PermSCT, path+".PermSCT")
	assert.Equal(t, v1.NumPermisoSCT, v2.NumPermisoSCT, path+".NumPermisoSCT")
}

func assertAutotransporteIdentificacionVehicular(t *testing.T, v1, v2 *IdentificacionVehicular, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.ConfigVehicular, v2.ConfigVehicular, path+".ConfigVehicular")
	assert.Equal(t, v1.PlacaVM, v2.PlacaVM, path+".PlacaVM")
	assert.Equal(t, v1.AnioModeloVM, v2.AnioModeloVM, path+".AnioModeloVM")
}

func assertAutotransporteSeguros(t *testing.T, v1, v2 *Seguros, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.AseguraRespCivil, v2.AseguraRespCivil, path+".AseguraRespCivil")
	assert.Equal(t, v1.PolizaRespCivil, v2.PolizaRespCivil, path+".PolizaRespCivil")
	assert.Equal(t, v1.AseguraMedAmbiente, v2.AseguraMedAmbiente, path+".AseguraMedAmbiente")
	assert.Equal(t, v1.PolizaMedAmbiente, v2.PolizaMedAmbiente, path+".PolizaMedAmbiente")
	assert.Equal(t, v1.AseguraCarga, v2.AseguraCarga, path+".AseguraCarga")
	assert.Equal(t, v1.PolizaCarga, v2.PolizaCarga, path+".PolizaCarga")
	assert.True(t, v1.PrimaSeguro.Equal(v2.PrimaSeguro), path+".PrimaSeguro")
}

func assertAutotransporteRemolques(t *testing.T, v1, v2 Remolques, path string) {
	l1, l2 := len(v1), len(v2)
	assert.Equal(t, l1, l2, path+".Len()")
	if l1 != l2 {
		return
	}
	for i := 0; i < l1; i++ {
		r1, r2 := v1[i], v2[i]
		assert.Equal(t, r1.SubTipoRem, r2.SubTipoRem, path+".[%d].SubTipoRem", i)
		assert.Equal(t, r1.Placa, r2.Placa, path+".[%d].Placa", i)
	}
}

func assertTransporteMaritimo(t *testing.T, v1, v2 *TransporteMaritimo, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	panic("not implemented")
}

func assertTransporteAereo(t *testing.T, v1, v2 *TransporteAereo, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	panic("not implemented")
}

func assertTransporteFerroviario(t *testing.T, v1, v2 *TransporteFerroviario, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	panic("not implemented")
}

func assertFiguraTransporte(t *testing.T, v1, v2 *FiguraTransporte, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	l1, l2 := len(v1.TiposFigura), len(v2.TiposFigura)
	assert.Equal(t, l1, l2, path+".TiposFigura")
	if l1 != l2 {
		return
	}
	for i, fig1 := range v1.TiposFigura {
		fig2 := v2.TiposFigura[i]
		assertFiguraTransporteTiposFigura(t, fig1, fig2, fmt.Sprintf("%s.TiposFigura[%d]", path, i))
	}
}

func assertFiguraTransporteTiposFigura(t *testing.T, v1, v2 *TiposFigura, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.TipoFigura, v2.TipoFigura, path+".TipoFigura")
	assert.Equal(t, v1.RFCFigura, v2.RFCFigura, path+".RFCFigura")
	assert.Equal(t, v1.NumLicencia, v2.NumLicencia, path+".NumLicencia")
	assert.Equal(t, v1.NombreFigura, v2.NombreFigura, path+".NombreFigura")
	assert.Equal(t, v1.NumRegIdTribFigura, v2.NumRegIdTribFigura, path+".NumRegIdTribFigura")
	assert.Equal(t, v1.ResidenciaFiscalFigura, v2.ResidenciaFiscalFigura, path+".ResidenciaFiscalFigura")

	for i, parte1 := range v1.PartesTransporte {
		parte2 := v2.PartesTransporte[i]
		if parte1 == nil || parte2 == nil {
			assert.Nil(t, parte1, fmt.Sprintf("%s.PartesTransporte[%d]", path, i))
			assert.Nil(t, parte2, fmt.Sprintf("%s.PartesTransporte[%d]", path, i))
			continue
		}
		assert.Equal(t, parte1.ParteTransporte, parte2.ParteTransporte, fmt.Sprintf("%s.PartesTransporte[%d].ParteTransporte", path, i))
	}

	if v1.Domicilio == nil || v2.Domicilio == nil {
		assert.Nil(t, v1.Domicilio, path+".Domicilio")
		assert.Nil(t, v2.Domicilio, path+".Domicilio")
		return
	}
	assert.Equal(t, v1.Domicilio.Calle, v2.Domicilio.Calle, path+".Domicilio.Calle")
	assert.Equal(t, v1.Domicilio.NumeroExterior, v2.Domicilio.NumeroExterior, path+".Domicilio.NumeroExterior")
	assert.Equal(t, v1.Domicilio.NumeroInterior, v2.Domicilio.NumeroInterior, path+".Domicilio.NumeroInterior")
	assert.Equal(t, v1.Domicilio.Colonia, v2.Domicilio.Colonia, path+".Domicilio.Colonia")
	assert.Equal(t, v1.Domicilio.Localidad, v2.Domicilio.Localidad, path+".Domicilio.Localidad")
	assert.Equal(t, v1.Domicilio.Referencia, v2.Domicilio.Referencia, path+".Domicilio.Referencia")
	assert.Equal(t, v1.Domicilio.Municipio, v2.Domicilio.Municipio, path+".Domicilio.Municipio")
	assert.Equal(t, v1.Domicilio.Estado, v2.Domicilio.Estado, path+".Domicilio.Estado")
	assert.Equal(t, v1.Domicilio.Pais, v2.Domicilio.Pais, path+".Domicilio.Pais")
	assert.Equal(t, v1.Domicilio.CodigoPostal, v2.Domicilio.CodigoPostal, path+".Domicilio.CodigoPostal")
}
