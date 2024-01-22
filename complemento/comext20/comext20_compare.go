package comext20

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *ComercioExterior) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *ComercioExterior) {
	path := "TimbreFiscalDigital11"
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Version, v2.Version, path+".Version")
	compare.Comparable(diffs, v1.MotivoTraslado, v2.MotivoTraslado, path+".MotivoTraslado")
	compare.Comparable(diffs, v1.TipoOperacion, v2.TipoOperacion, path+".TipoOperacion")
	compare.Comparable(diffs, v1.ClaveDePedimento, v2.ClaveDePedimento, path+".ClaveDePedimento")
	compare.Comparable(diffs, v1.CertificadoOrigen, v2.CertificadoOrigen, path+".CertificadoOrigen")
	compare.Comparable(diffs, v1.NumCertificadoOrigen, v2.NumCertificadoOrigen, path+".NumCertificadoOrigen")
	compare.Comparable(diffs, v1.NumeroExportadorConfiable, v2.NumeroExportadorConfiable, path+".NumeroExportadorConfiable")
	compare.Comparable(diffs, v1.Incoterm, v2.Incoterm, path+".Incoterm")
	compare.Comparable(diffs, v1.Subdivision, v2.Subdivision, path+".Subdivision")
	compare.Comparable(diffs, v1.Observaciones, v2.Observaciones, path+".Observaciones")
	compare.Decimal(diffs, v1.TipoCambioUSD, v2.TipoCambioUSD, path+".TipoCambioUSD")
	compare.Decimal(diffs, v1.TotalUSD, v2.TotalUSD, path+".TotalUSD")

	compareMercancias(diffs, v1.Mercancias, v2.Mercancias, path+".Mercancias")
	compareEmisor(diffs, v1.Emisor, v2.Emisor, path+".Emisor")
	compareReceptor(diffs, v1.Receptor, v2.Receptor, path+".Receptor")

	l1, l2 := len(v1.Propietarios), len(v2.Propietarios)
	compare.Comparable(diffs, l1, l2, path+".Propietarios.len()")
	if l1 == l2 {
		for i, m1 := range v1.Propietarios {
			m2 := v2.Propietarios[i]
			if compare.Nil(diffs, m1, m2, path+".Propietario[%d]", i) {
				continue
			} else if m1 == nil || m2 == nil {
				continue
			}
			compare.Comparable(diffs, m1.NumRegIdTrib, m2.NumRegIdTrib, path+".Propietario[%d].NumRegIdTrib", i)
			compare.Comparable(diffs, m1.ResidenciaFiscal, m2.ResidenciaFiscal, path+".Propietario[%d].ResidenciaFiscal", i)
		}
	}

	l1, l2 = len(v1.Destinatarios), len(v2.Destinatarios)
	compare.Comparable(diffs, l1, l2, path+".Destinatario.len()")
	if l1 == l2 {
		for i, m1 := range v1.Destinatarios {
			m2 := v2.Destinatarios[i]
			if compare.Nil(diffs, m1, m2, path+".Destinatario[%d]", i) {
				continue
			} else if m1 == nil || m2 == nil {
				continue
			}
			compare.Comparable(diffs, m1.NumRegIdTrib, m2.NumRegIdTrib, path+".Destinatario[%d].NumRegIdTrib", i)
			compare.Comparable(diffs, m1.Nombre, m2.Nombre, path+".Destinatario[%d].Nombre", i)

			l3, l4 := len(m1.Domicilios), len(m2.Domicilios)
			compare.Comparable(diffs, l3, l4, path+".Domicilio.len()")
			if l3 == l4 {
				for j, m3 := range m1.Domicilios {
					m4 := m2.Domicilios[j]
					if compare.Nil(diffs, m3, m4, path+".Destinatario[%d]", j) {
						continue
					} else if m3 == nil || m4 == nil {
						continue
					}
					compareDomicilio(diffs, m3, m4, path+".Domicilio[%d]")
				}
			}
		}
	}
}

func compareMercancias(diffs *compare.Diffs, v1, v2 Mercancias, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	l1, l2 := len(v1), len(v2)
	compare.Comparable(diffs, l1, l2, path+".Mercancia.len()")
	if l1 == l2 {
		for i, m1 := range v1 {
			m2 := v2[i]
			compareMercancia(diffs, m1, m2, fmt.Sprintf(".Mercancia[%d]", i))
		}
	}
}
func compareMercancia(diffs *compare.Diffs, v1, v2 *Mercancia, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.NoIdentificacion, v2.NoIdentificacion, path+".NoIdentificacion")
	compare.Comparable(diffs, v1.FraccionArancelaria, v2.FraccionArancelaria, path+".FraccionArancelaria")
	compare.Decimal(diffs, v1.CantidadAduana, v2.CantidadAduana, path+".CantidadAduana")
	compare.Comparable(diffs, v1.UnidadAduana, v2.UnidadAduana, path+".UnidadAduana")
	compare.Decimal(diffs, v1.ValorUnitarioAduana, v2.ValorUnitarioAduana, path+".ValorUnitarioAduana")
	compare.Decimal(diffs, v1.ValorDolares, v2.ValorDolares, path+".ValorDolares")

	l1, l2 := len(v1.DescripcionesEspecificas), len(v2.DescripcionesEspecificas)
	compare.Comparable(diffs, l1, l2, path+".DescripcionesEspecificas.len()")
	if l1 == l2 {
		for i, m1 := range v1.DescripcionesEspecificas {
			m2 := v2.DescripcionesEspecificas[i]
			compareDescripcionesEspecificas(diffs, m1, m2, fmt.Sprintf(".Mercancia[%d]", i))
		}
	}
}
func compareDescripcionesEspecificas(diffs *compare.Diffs, v1, v2 *DescripcionesEspecificas, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Marca, v2.Marca, path+".Marca")
	compare.Comparable(diffs, v1.Modelo, v2.Modelo, path+".Modelo")
	compare.Comparable(diffs, v1.SubModelo, v2.SubModelo, path+".SubModelo")
	compare.Comparable(diffs, v1.NumeroSerie, v2.NumeroSerie, path+".NumeroSerie")
}

func compareEmisor(diffs *compare.Diffs, v1, v2 *Emisor, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.Curp, v2.Curp, path+".Curp")
	compareDomicilio(diffs, v1.Domicilio, v2.Domicilio, path+".Domicilio")
}

func compareReceptor(diffs *compare.Diffs, v1, v2 *Receptor, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.NumRegIdTrib, v2.NumRegIdTrib, path+".NumRegIdTrib")
	compareDomicilio(diffs, v1.Domicilio, v2.Domicilio, path+".Domicilio")
}

func compareDomicilio(diffs *compare.Diffs, v1, v2 *Domicilio, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Calle, v2.Calle, path+".Calle")
	compare.Comparable(diffs, v1.NumeroExterior, v2.NumeroExterior, path+".NumeroExterior")
	compare.Comparable(diffs, v1.NumeroInterior, v2.NumeroInterior, path+".NumeroInterior")
	compare.Comparable(diffs, v1.Colonia, v2.Colonia, path+".Colonia")
	compare.Comparable(diffs, v1.Localidad, v2.Localidad, path+".Localidad")
	compare.Comparable(diffs, v1.Referencia, v2.Referencia, path+".Referencia")
	compare.Comparable(diffs, v1.Municipio, v2.Municipio, path+".Municipio")
	compare.Comparable(diffs, v1.Estado, v2.Estado, path+".Estado")
	compare.Comparable(diffs, v1.Pais, v2.Pais, path+".Pais")
	compare.Comparable(diffs, v1.CodigoPostal, v2.CodigoPostal, path+".CodigoPostal")
}

// func compareDestinartarios(diffs *compare.Diffs, v1, v2 []*Destinatario, path string) {
// 	if compare.Nil(diffs, v1, v2, path) {
// 		return
// 	} else if v1 == nil || v2 == nil {
// 		return
// 	}

// 	compare.Comparable(diffs, v1.NumRegIdTrib, v2.NumRegIdTrib, path+".NumRegIdTrib")
// 	compare.Comparable(diffs, v1.Nombre, v2.Nombre, path+".Nombre")

// 	l1, l2 := len(v1.Domicilios), len(v2.Domicilios)
// 	compare.Comparable(diffs, l1, l2, path+".Domicilio.len()")
// 	if l1 == l2 {
// 		for i, m1 := range v1.Domicilios {
// 			m2 := v2.Domicilios[i]
// 			compareDomicilio(diffs, m1, m2, fmt.Sprintf(".Domicilio[%d]", i))
// 		}
// 	}
// }
