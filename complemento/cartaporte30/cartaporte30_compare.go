package cartaporte30

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *CartaPorte30) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *CartaPorte30) {
	path := "Cartaporte30"
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Version, v2.Version, path+".Version")
	compare.Comparable(diffs, v1.TranspInternac, v2.TranspInternac, path+".TranspInternac")
	compare.Comparable(diffs, v1.EntradaSalidaMerc, v2.EntradaSalidaMerc, path+".EntradaSalidaMerc")
	compare.Comparable(diffs, v1.PaisOrigenDestino, v2.PaisOrigenDestino, path+".PaisOrigenDestino")
	compare.Comparable(diffs, v1.ViaEntradaSalida, v2.ViaEntradaSalida, path+".ViaEntradaSalida")
	compare.Decimal(diffs, v1.TotalDistRec, v2.TotalDistRec, path+".TotalDistRec")

	l1, l2 := len(v1.Ubicaciones), len(v2.Ubicaciones)
	compare.Comparable(diffs, l1, l2, path+".Ubicaciones.len()")
	if l1 == l2 {
		for i, ub1 := range v1.Ubicaciones {
			ub2 := v2.Ubicaciones[i]
			compareUbicaciones(diffs, ub1, ub2, fmt.Sprintf(path+".Ubicaciones[%d]", i))
		}
	}

	compareMercancias(diffs, v1.Mercancias, v2.Mercancias, path+".Mercancias")
	compareFiguraTransporte(diffs, v1.FiguraTransporte, v2.FiguraTransporte, path+".FiguraTransporte")
}

func compareUbicaciones(diffs *compare.Diffs, v1, v2 *Ubicacion, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.TipoUbicacion, v2.TipoUbicacion, path+".TipoUbicacion")
	compare.Comparable(diffs, v1.IDUbicacion, v2.IDUbicacion, path+".IDUbicacion")
	compare.Comparable(diffs, v1.RFCRemitenteDestinatario, v2.RFCRemitenteDestinatario, path+".RFCRemitenteDestinatario")
	compare.Comparable(diffs, v1.NombreRemitenteDestinatario, v2.NombreRemitenteDestinatario, path+".NombreRemitenteDestinatario")
	compare.Comparable(diffs, v1.NumRegIdTrib, v2.NumRegIdTrib, path+".NumRegIdTrib")
	compare.Comparable(diffs, v1.ResidenciaFiscal, v2.ResidenciaFiscal, path+".ResidenciaFiscal")
	compare.Comparable(diffs, v1.NumEstacion, v2.NumEstacion, path+".NumEstacion")
	compare.Comparable(diffs, v1.NombreEstacion, v2.NombreEstacion, path+".NombreEstacion")
	compare.Comparable(diffs, v1.NavegacionTrafico, v2.NavegacionTrafico, path+".NavegacionTrafico")
	compare.Comparable(diffs, v1.FechaHoraSalidaLlegada, v2.FechaHoraSalidaLlegada, path+".FechaHoraSalidaLlegada")
	compare.Comparable(diffs, v1.TipoEstacion, v2.TipoEstacion, path+".TipoEstacion")

	if compare.Nil(diffs, v1.Domicilio, v2.Domicilio, path) {
		return
	} else if v1.Domicilio == nil || v2.Domicilio == nil {
		return
	}

	compare.Comparable(diffs, v1.Domicilio.Calle, v2.Domicilio.Calle, path+".Domicilio.Calle")
	compare.Comparable(diffs, v1.Domicilio.NumeroExterior, v2.Domicilio.NumeroExterior, path+".Domicilio.NumeroExterior")
	compare.Comparable(diffs, v1.Domicilio.NumeroInterior, v2.Domicilio.NumeroInterior, path+".Domicilio.NumeroInterior")
	compare.Comparable(diffs, v1.Domicilio.Colonia, v2.Domicilio.Colonia, path+".Domicilio.Colonia")
	compare.Comparable(diffs, v1.Domicilio.Localidad, v2.Domicilio.Localidad, path+".Domicilio.Localidad")
	compare.Comparable(diffs, v1.Domicilio.Referencia, v2.Domicilio.Referencia, path+".Domicilio.Referencia")
	compare.Comparable(diffs, v1.Domicilio.Municipio, v2.Domicilio.Municipio, path+".Domicilio.Municipio")
	compare.Comparable(diffs, v1.Domicilio.Estado, v2.Domicilio.Estado, path+".Domicilio.Estado")
	compare.Comparable(diffs, v1.Domicilio.Pais, v2.Domicilio.Pais, path+".Domicilio.Pais")
	compare.Comparable(diffs, v1.Domicilio.CodigoPostal, v2.Domicilio.CodigoPostal, path+".Domicilio.CodigoPostal")

}

func compareMercancias(diffs *compare.Diffs, v1, v2 *Mercancias, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	l1, l2 := len(v1.Mercancia), len(v2.Mercancia)
	compare.Comparable(diffs, l1, l2, path+".Mercancia.len()")
	if l1 == l2 {
		for i, m1 := range v1.Mercancia {
			m2 := v2.Mercancia[i]
			compareMercancia(diffs, m1, m2, fmt.Sprintf(".Mercancia[%d]", i))
		}
	}
	compareAutotransporte(diffs, v1.Autotransporte, v2.Autotransporte, path+".Autotransporte")
	compareTransporteMaritimo(diffs, v1.TransporteMaritimo, v2.TransporteMaritimo, path+".TransporteMaritimo")
	compareTransporteAereo(diffs, v1.TransporteAereo, v2.TransporteAereo, path+".TransporteAereo")
	compareTransporteFerroviario(diffs, v1.TransporteFerroviario, v2.TransporteFerroviario, path+".TransporteFerroviario")
}

func compareMercancia(diffs *compare.Diffs, v1, v2 *Mercancia, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.BienesTransp, v2.BienesTransp, path+".BienesTransp")
	compare.Comparable(diffs, v1.ClaveSTCC, v2.ClaveSTCC, path+".ClaveSTCC")
	compare.Comparable(diffs, v1.Descripcion, v2.Descripcion, path+".Descripcion")
	compare.Decimal(diffs, v1.Cantidad, v2.Cantidad, path+".Cantidad")
	compare.Comparable(diffs, v1.ClaveUnidad, v2.ClaveUnidad, path+".ClaveUnidad")
	compare.Comparable(diffs, v1.Unidad, v2.Unidad, path+".Unidad")
	compare.Comparable(diffs, v1.Dimensiones, v2.Dimensiones, path+".Dimensiones")
	compare.Comparable(diffs, v1.MaterialPeligroso, v2.MaterialPeligroso, path+".MaterialPeligroso")
	compare.Comparable(diffs, v1.CveMaterialPeligroso, v2.CveMaterialPeligroso, path+".CveMaterialPeligroso")
	compare.Comparable(diffs, v1.Embalaje, v2.Embalaje, path+".Embalaje")
	compare.Comparable(diffs, v1.DescripEmbalaje, v2.DescripEmbalaje, path+".DescripEmbalaje")
	compare.Comparable(diffs, v1.SectorCOFEPRIS, v2.SectorCOFEPRIS, path+".SectorCOFEPRIS")
	compare.Comparable(diffs, v1.NombreIngredienteActivo, v2.NombreIngredienteActivo, path+".NombreIngredienteActivo")
	compare.Comparable(diffs, v1.NomQuimico, v2.NomQuimico, path+".NomQuimico")
	compare.Comparable(diffs, v1.DenominacionGenericaProd, v2.DenominacionGenericaProd, path+".DenominacionGenericaProd")
	compare.Comparable(diffs, v1.DenominacionDistintivaProd, v2.DenominacionDistintivaProd, path+".DenominacionDistintivaProd")
	compare.Comparable(diffs, v1.Fabricante, v2.Fabricante, path+".Fabricante")
	compare.Comparable(diffs, v1.FechaCaducidad, v2.FechaCaducidad, path+".FechaCaducidad")
	compare.Comparable(diffs, v1.LoteMedicamento, v2.LoteMedicamento, path+".LoteMedicamento")
	compare.Comparable(diffs, v1.FormaFarmaceutica, v2.FormaFarmaceutica, path+".FormaFarmaceutica")
	compare.Comparable(diffs, v1.CondicionesEspTransp, v2.CondicionesEspTransp, path+".CondicionesEspTransp")
	compare.Comparable(diffs, v1.RegistroSanitarioFolioAutorizacion, v2.RegistroSanitarioFolioAutorizacion, path+".RegistroSanitarioFolioAutorizacion")
	compare.Comparable(diffs, v1.PermisoImportacion, v2.PermisoImportacion, path+".PermisoImportacion")
	compare.Comparable(diffs, v1.FolioImpoVUCEM, v2.FolioImpoVUCEM, path+".FolioImpoVUCEM")
	compare.Comparable(diffs, v1.NumCAS, v2.NumCAS, path+".NumCAS")
	compare.Comparable(diffs, v1.RazonSocialEmpImp, v2.RazonSocialEmpImp, path+".RazonSocialEmpImp")
	compare.Comparable(diffs, v1.NumRegSanPlagCOFEPRIS, v2.NumRegSanPlagCOFEPRIS, path+".NumRegSanPlagCOFEPRIS")
	compare.Comparable(diffs, v1.DatosFabricante, v2.DatosFabricante, path+".DatosFabricante")
	compare.Comparable(diffs, v1.DatosFormulador, v2.DatosFormulador, path+".DatosFormulador")
	compare.Comparable(diffs, v1.DatosMaquilador, v2.DatosMaquilador, path+".DatosMaquilador")
	compare.Comparable(diffs, v1.UsoAutorizado, v2.UsoAutorizado, path+".UsoAutorizado")
	compare.Decimal(diffs, v1.PesoEnKg, v2.PesoEnKg, path+".PesoEnKg")
	compare.Decimal(diffs, v1.ValorMercancia, v2.ValorMercancia, path+".ValorMercancia")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Comparable(diffs, v1.FraccionArancelaria, v2.FraccionArancelaria, path+".FraccionArancelaria")
	compare.Comparable(diffs, v1.UUIDComercioExt, v2.UUIDComercioExt, path+".UUIDComercioExt")

	l1, l2 := len(v1.DocumentacionAduanera), len(v2.DocumentacionAduanera)
	compare.Comparable(diffs, l1, l2, path+".DocumentacionAduanera.len()")
	if l1 == l2 {
		for i, p1 := range v1.DocumentacionAduanera {
			p2 := v2.DocumentacionAduanera[i]
			if compare.Nil(diffs, p1, p2, path+".DocumentacionAduanera[%d]", i) {
				continue
			} else if p1 == nil || p2 == nil {
				continue
			}
			compare.Comparable(diffs, p1.TipoDocumento, p2.TipoDocumento, path+".DocumentacionAduanera[%d].Pedimento", i)
			compare.Comparable(diffs, p1.NumPedimento, p2.NumPedimento, path+".DocumentacionAduanera[%d].NumPedimento", i)
			compare.Comparable(diffs, p1.IdentDocAduanero, p2.IdentDocAduanero, path+".DocumentacionAduanera[%d].IdentDocAduanero", i)
			compare.Comparable(diffs, p1.RFCImpo, p2.RFCImpo, path+".DocumentacionAduanera[%d].RFCImpo", i)
		}
	}

	l1, l2 = len(v1.GuiasIdentificacion), len(v2.GuiasIdentificacion)
	compare.Comparable(diffs, l1, l2, path+".GuiasIdentificacion.len()")
	if l1 == l2 {
		for i, g1 := range v1.GuiasIdentificacion {
			g2 := v2.GuiasIdentificacion[i]
			if compare.Nil(diffs, g1, g2, path+".GuiasIdentificacion[%d]", i) {
				continue
			} else if g1 == nil || g2 == nil {
				continue
			}

			compare.Comparable(diffs, g1.NumeroGuiaIdentificacion, g2.NumeroGuiaIdentificacion, path+".GuiasIdentificacion[%d].NumeroGuiaIdentificacion", i)
			compare.Comparable(diffs, g1.DescripGuiaIdentificacion, g2.DescripGuiaIdentificacion, path+".GuiasIdentificacion[%d].DescripGuiaIdentificacion", i)
		}
	}
	l1, l2 = len(v1.CantidadTransporta), len(v2.CantidadTransporta)
	compare.Comparable(diffs, l1, l2, path+".CantidadTransporta.len()")
	if l1 == l2 {
		for i, c1 := range v1.CantidadTransporta {
			c2 := v2.CantidadTransporta[i]
			if compare.Nil(diffs, c1, c2, path+".CantidadTransporta[%d]", i) {
				continue
			} else if c1 == nil || c2 == nil {
				continue
			}

			compare.Decimal(diffs, c1.Cantidad, c2.Cantidad, path+".CantidadTransporta[%d].Cantidad", i)
			compare.Comparable(diffs, c1.IDOrigen, c2.IDOrigen, path+".CantidadTransporta[%d].IDOrigen", i)
			compare.Comparable(diffs, c1.IDDestino, c2.IDDestino, path+".CantidadTransporta[%d].IDDestino", i)
			compare.Comparable(diffs, c1.CvesTransporte, c2.CvesTransporte, path+".CantidadTransporta[%d].CvesTransporte", i)
		}
	}

	if compare.Nil(diffs, v1.DetalleMercancia, v2.DetalleMercancia, path) {
		return
	} else if v1.DetalleMercancia == nil || v2.DetalleMercancia == nil {
		return
	}

	compare.Comparable(diffs, v1.DetalleMercancia.UnidadPesoMerc, v2.DetalleMercancia.UnidadPesoMerc, path+".DetalleMercancia.UnidadPesoMerc")
	compare.Decimal(diffs, v1.DetalleMercancia.PesoBruto, v2.DetalleMercancia.PesoBruto, path+".DetalleMercancia.PesoBruto")
	compare.Decimal(diffs, v1.DetalleMercancia.PesoNeto, v2.DetalleMercancia.PesoNeto, path+".DetalleMercancia.PesoNeto")
	compare.Decimal(diffs, v1.DetalleMercancia.PesoTara, v2.DetalleMercancia.PesoTara, path+".DetalleMercancia.PesoTara")
	compare.Comparable(diffs, v1.DetalleMercancia.NumPiezas, v2.DetalleMercancia.NumPiezas, path+".DetalleMercancia.NumPiezas")
}

func compareAutotransporte(diffs *compare.Diffs, v1, v2 *Autotransporte, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compareAutotransporteIdentificacionVehicular(diffs, v1.IdentificacionVehicular, v2.IdentificacionVehicular, path+".IdentificacionVehicular")
	compareAutotransporteSeguros(diffs, v1.Seguros, v2.Seguros, path+".Seguros")
	compareAutotransporteRemolques(diffs, v1.Remolques, v2.Remolques, path+".Remolques")

	compare.Comparable(diffs, v1.PermSCT, v2.PermSCT, path+".PermSCT")
	compare.Comparable(diffs, v1.NumPermisoSCT, v2.NumPermisoSCT, path+".NumPermisoSCT")
}

func compareAutotransporteIdentificacionVehicular(diffs *compare.Diffs, v1, v2 *IdentificacionVehicular, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.ConfigVehicular, v2.ConfigVehicular, path+".ConfigVehicular")
	compare.Comparable(diffs, v1.PlacaVM, v2.PlacaVM, path+".PlacaVM")
	compare.Comparable(diffs, v1.AnioModeloVM, v2.AnioModeloVM, path+".AnioModeloVM")
}

func compareAutotransporteSeguros(diffs *compare.Diffs, v1, v2 *Seguros, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.AseguraRespCivil, v2.AseguraRespCivil, path+".AseguraRespCivil")
	compare.Comparable(diffs, v1.PolizaRespCivil, v2.PolizaRespCivil, path+".PolizaRespCivil")
	compare.Comparable(diffs, v1.AseguraMedAmbiente, v2.AseguraMedAmbiente, path+".AseguraMedAmbiente")
	compare.Comparable(diffs, v1.PolizaMedAmbiente, v2.PolizaMedAmbiente, path+".PolizaMedAmbiente")
	compare.Comparable(diffs, v1.AseguraCarga, v2.AseguraCarga, path+".AseguraCarga")
	compare.Comparable(diffs, v1.PolizaCarga, v2.PolizaCarga, path+".PolizaCarga")
	compare.Decimal(diffs, v1.PrimaSeguro, v2.PrimaSeguro, path+".PrimaSeguro")
}

func compareAutotransporteRemolques(diffs *compare.Diffs, v1, v2 Remolques, path string) {
	l1, l2 := len(v1), len(v2)
	compare.Comparable(diffs, l1, l2, path+".Len()")
	if l1 != l2 {
		return
	}
	for i := 0; i < l1; i++ {
		r1, r2 := v1[i], v2[i]
		compare.Comparable(diffs, r1.SubTipoRem, r2.SubTipoRem, path+".[%d].SubTipoRem", i)
		compare.Comparable(diffs, r1.Placa, r2.Placa, path+".[%d].Placa", i)
	}
}

func compareTransporteMaritimo(diffs *compare.Diffs, v1, v2 *TransporteMaritimo, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	panic("not implemented")
}

func compareTransporteAereo(diffs *compare.Diffs, v1, v2 *TransporteAereo, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	panic("not implemented")
}

func compareTransporteFerroviario(diffs *compare.Diffs, v1, v2 *TransporteFerroviario, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	panic("not implemented")
}

func compareFiguraTransporte(diffs *compare.Diffs, v1, v2 *FiguraTransporte, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	l1, l2 := len(v1.TiposFigura), len(v2.TiposFigura)
	if l1 != l2 {
		return
	}
	for i, fig1 := range v1.TiposFigura {
		fig2 := v2.TiposFigura[i]
		compareFiguraTransporteTiposFigura(diffs, fig1, fig2, fmt.Sprintf("%s.TiposFigura[%d]", path, i))
	}
}

func compareFiguraTransporteTiposFigura(diffs *compare.Diffs, v1, v2 *TiposFigura, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.TipoFigura, v2.TipoFigura, path+".TipoFigura")
	compare.Comparable(diffs, v1.RFCFigura, v2.RFCFigura, path+".RFCFigura")
	compare.Comparable(diffs, v1.NumLicencia, v2.NumLicencia, path+".NumLicencia")
	compare.Comparable(diffs, v1.NombreFigura, v2.NombreFigura, path+".NombreFigura")
	compare.Comparable(diffs, v1.NumRegIdTribFigura, v2.NumRegIdTribFigura, path+".NumRegIdTribFigura")
	compare.Comparable(diffs, v1.ResidenciaFiscalFigura, v2.ResidenciaFiscalFigura, path+".ResidenciaFiscalFigura")

	for i, parte1 := range v1.PartesTransporte {
		parte2 := v2.PartesTransporte[i]
		if compare.Nil(diffs, parte1, parte2, path+".PartesTransporte[%d]", i) {
			continue
		} else if parte1 == nil || parte2 == nil {
			continue
		}

		compare.Comparable(diffs, parte1.ParteTransporte, parte2.ParteTransporte, fmt.Sprintf("%s.PartesTransporte[%d].ParteTransporte", path, i))
	}

	if compare.Nil(diffs, v1.Domicilio, v2.Domicilio, path+".Domicilio") {
		return
	} else if v1.Domicilio == nil || v2.Domicilio == nil {
		return
	}

	compare.Comparable(diffs, v1.Domicilio.Calle, v2.Domicilio.Calle, path+".Domicilio.Calle")
	compare.Comparable(diffs, v1.Domicilio.NumeroExterior, v2.Domicilio.NumeroExterior, path+".Domicilio.NumeroExterior")
	compare.Comparable(diffs, v1.Domicilio.NumeroInterior, v2.Domicilio.NumeroInterior, path+".Domicilio.NumeroInterior")
	compare.Comparable(diffs, v1.Domicilio.Colonia, v2.Domicilio.Colonia, path+".Domicilio.Colonia")
	compare.Comparable(diffs, v1.Domicilio.Localidad, v2.Domicilio.Localidad, path+".Domicilio.Localidad")
	compare.Comparable(diffs, v1.Domicilio.Referencia, v2.Domicilio.Referencia, path+".Domicilio.Referencia")
	compare.Comparable(diffs, v1.Domicilio.Municipio, v2.Domicilio.Municipio, path+".Domicilio.Municipio")
	compare.Comparable(diffs, v1.Domicilio.Estado, v2.Domicilio.Estado, path+".Domicilio.Estado")
	compare.Comparable(diffs, v1.Domicilio.Pais, v2.Domicilio.Pais, path+".Domicilio.Pais")
	compare.Comparable(diffs, v1.Domicilio.CodigoPostal, v2.Domicilio.CodigoPostal, path+".Domicilio.CodigoPostal")
}
