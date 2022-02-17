package cfdi40

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
	"github.com/veyronifs/cfdi-go/complemento/cartaporte20"
	"github.com/veyronifs/cfdi-go/complemento/pagos20"
	"github.com/veyronifs/cfdi-go/complemento/tfd11"
)

func CompareEqual(v1, v2 *Comprobante) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *Comprobante) {
	path := ""
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Version, v2.Version, path+".Version")
	compare.Comparable(diffs, v1.Serie, v2.Serie, path+".Serie")
	compare.Comparable(diffs, v1.Folio, v2.Folio, path+".Folio")
	compare.Comparable(diffs, v1.Fecha.String(), v2.Fecha.String(), path+".Fecha")
	compare.Comparable(diffs, v1.Sello, v2.Sello, path+".Sello")
	compare.Comparable(diffs, v1.FormaPago, v2.FormaPago, path+".FormaPago")
	compare.Comparable(diffs, v1.NoCertificado, v2.NoCertificado, path+".NoCertificado")
	compare.Comparable(diffs, v1.Certificado, v2.Certificado, path+".Certificado")
	compare.Comparable(diffs, v1.CondicionesDePago, v2.CondicionesDePago, path+".CondicionesDePago")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Comparable(diffs, v1.TipoDeComprobante, v2.TipoDeComprobante, path+".TipoDeComprobante")
	compare.Comparable(diffs, v1.Exportacion, v2.Exportacion, path+".Exportacion")
	compare.Comparable(diffs, v1.MetodoPago, v2.MetodoPago, path+".MetodoPago")
	compare.Comparable(diffs, v1.LugarExpedicion, v2.LugarExpedicion, path+".LugarExpedicion")
	compare.Comparable(diffs, v1.Confirmacion, v2.Confirmacion, path+".Confirmacion")
	compare.Decimal(diffs, v1.SubTotal, v2.SubTotal, path+".SubTotal")
	compare.Decimal(diffs, v1.Descuento, v2.Descuento, path+".Descuento")
	compare.Decimal(diffs, v1.TipoCambio, v2.TipoCambio, path+".TipoCambio")
	compare.Decimal(diffs, v1.Total, v2.Total, path+".Total")

	compareEqualInformacionGlobal(diffs, v1.InformacionGlobal, v2.InformacionGlobal, path+".InformacionGlobal")
	compareEqualEmisor(diffs, &v1.Emisor, &v2.Emisor, path+".Emisor")
	compareEqualReceptor(diffs, &v1.Receptor, &v2.Receptor, path+".Receptor")

	compareEqualCfdiRelacionados(diffs, v1.CfdiRelacionados, v2.CfdiRelacionados, path+".CfdiRelacionados")

	l1, l2 := len(v1.Conceptos), len(v2.Conceptos)
	compare.Comparable(diffs, l1, l2, path+".Conceptos.len()")
	for i, conc1 := range v1.Conceptos {
		compareEqualComprobanteConceptos(diffs, conc1, v2.Conceptos[i], fmt.Sprintf(".Conceptos[%d]", i))
	}

	compareEqualImpuestos(diffs, v1.Impuestos, v2.Impuestos, path+".Impuestos")
	compareComplemento(diffs, v1.Complemento, v2.Complemento, path+".Complemento")
}

func compareEqualInformacionGlobal(diffs *compare.Diffs, v1, v2 *InformacionGlobal, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Periodicidad, v2.Periodicidad, path+".Periodicidad")
	compare.Comparable(diffs, v1.Meses, v2.Meses, path+".Meses")
	compare.Comparable(diffs, v1.Anio, v2.Anio, path+".Anio")

}

func compareEqualEmisor(diffs *compare.Diffs, v1, v2 *Emisor, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Rfc, v2.Rfc, path+".Rfc")
	compare.Comparable(diffs, v1.Nombre, v2.Nombre, path+".Nombre")
	compare.Comparable(diffs, v1.RegimenFiscal, v2.RegimenFiscal, path+".RegimenFiscal")
	compare.Comparable(diffs, v1.FacAtrAdquirente, v2.FacAtrAdquirente, path+".FacAtrAdquirente")
}

func compareEqualReceptor(diffs *compare.Diffs, v1, v2 *Receptor, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Rfc, v2.Rfc, path+".Rfc")
	compare.Comparable(diffs, v1.Nombre, v2.Nombre, path+".Nombre")
	compare.Comparable(diffs, v1.DomicilioFiscalReceptor, v2.DomicilioFiscalReceptor, path+".DomicilioFiscalReceptor")
	compare.Comparable(diffs, v1.ResidenciaFiscal, v2.ResidenciaFiscal, path+".ResidenciaFiscal")
	compare.Comparable(diffs, v1.NumRegIdTrib, v2.NumRegIdTrib, path+".NumRegIdTrib")
	compare.Comparable(diffs, v1.RegimenFiscalReceptor, v2.RegimenFiscalReceptor, path+".RegimenFiscalReceptor")
	compare.Comparable(diffs, v1.UsoCFDI, v2.UsoCFDI, path+".UsoCFDI")
}

func compareComplemento(diffs *compare.Diffs, v1, v2 *Complemento, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	cartaporte20.Compare(diffs, v1.CartaPorte20, v2.CartaPorte20)
	tfd11.Compare(diffs, v1.TFD11, v2.TFD11)
	pagos20.Compare(diffs, v1.Pagos20, v2.Pagos20)
}

func compareEqualCfdiRelacionados(diffs *compare.Diffs, r1, r2 []*CfdiRelacionados, path string) {
	l1, l2 := len(r1), len(r2)
	compare.Comparable(diffs, l1, l2, path+".len()")
	if l1 != l2 {
		return
	}

	for i, rel1 := range r1 {
		rel2 := r2[i]
		if compare.Nil(diffs, rel1, rel2, path+fmt.Sprintf(".CfdiRelacionados[%d]", i)) {
			continue
		} else if rel1 == nil || rel2 == nil {
			continue
		}

		compare.Comparable(diffs, rel1.TipoRelacion, rel2.TipoRelacion, path+fmt.Sprintf(".CfdiRelacionados[%d].TipoRelacion", i))

		l1, l2 = len(rel1.CfdiRelacionado), len(rel2.CfdiRelacionado)
		compare.Comparable(diffs, l1, l2, path+fmt.Sprintf(".CfdiRelacionados[%d].CfdiRelacionado.len()", i))
		if l1 != l2 {
			continue
		}
		for j, cfdi1 := range rel1.CfdiRelacionado {
			cfdi2 := rel2.CfdiRelacionado[j]
			compare.Comparable(diffs, cfdi1.UUID, cfdi2.UUID, path+fmt.Sprintf(".CfdiRelacionados[%d].CfdiRelacionado[%d].UUID", i, j))
		}
	}
}

func compareEqualImpuestos(diffs *compare.Diffs, v1 *Impuestos, v2 *Impuestos, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Decimal(diffs, v1.TotalImpuestosTrasladados, v2.TotalImpuestosTrasladados, path+".TotalImpuestosTrasladados")
	compare.Decimal(diffs, v1.TotalImpuestosRetenidos, v2.TotalImpuestosRetenidos, path+".TotalImpuestosRetenidos")

	l1, l2 := len(v1.Traslados), len(v2.Traslados)
	compare.Comparable(diffs, l1, l2, path+".Traslados.len()")
	if len(v1.Traslados) == len(v2.Traslados) {
		for i, tras1 := range v1.Traslados {
			tras2 := v2.Traslados[i]
			compare.Comparable(diffs, tras1.Impuesto, tras2.Impuesto, path+fmt.Sprintf(".Traslados[%d].Impuesto", i))
			compare.Comparable(diffs, tras1.TipoFactor, tras2.TipoFactor, path+fmt.Sprintf(".Traslados[%d].TipoFactor", i))
			compare.Decimal(diffs, tras1.TasaOCuota, tras2.TasaOCuota, path+fmt.Sprintf(".Traslados[%d].TasaOCuota", i))
			compare.Decimal(diffs, tras1.Importe, tras2.Importe, path+fmt.Sprintf(".Traslados[%d].Importe", i))
		}
	}
	l1, l2 = len(v1.Retenciones), len(v2.Retenciones)
	compare.Comparable(diffs, l1, l2, path+".Retenciones.len()")
	if len(v1.Retenciones) == len(v2.Retenciones) {
		for i, ret1 := range v1.Retenciones {
			ret2 := v2.Retenciones[i]
			compare.Comparable(diffs, ret1.Impuesto, ret2.Impuesto, path+fmt.Sprintf(".Retenciones[%d].Impuesto", i))
			compare.Decimal(diffs, ret1.Importe, ret2.Importe, path+fmt.Sprintf(".Retenciones[%d].Importe", i))
		}
	}
}

func compareEqualComprobanteConceptos(diffs *compare.Diffs, v1, v2 *Concepto, path string) {
	compare.Comparable(diffs, v1.ClaveProdServ, v2.ClaveProdServ, path+".ClaveProdServ")
	compare.Comparable(diffs, v1.NoIdentificacion, v2.NoIdentificacion, path+".NoIdentificacion")
	compare.Comparable(diffs, v1.ClaveUnidad, v2.ClaveUnidad, path+".ClaveUnidad")
	compare.Comparable(diffs, v1.Unidad, v2.Unidad, path+".Unidad")
	compare.Comparable(diffs, v1.Descripcion, v2.Descripcion, path+".Descripcion")
	compare.Comparable(diffs, v1.ObjetoImp, v2.ObjetoImp, path+".ObjetoImp")
	compare.Decimal(diffs, v1.Cantidad, v2.Cantidad, path+".Cantidad")
	compare.Decimal(diffs, v1.ValorUnitario, v2.ValorUnitario, path+".ValorUnitario")
	compare.Decimal(diffs, v1.Importe, v2.Importe, path+".Importe")
	compare.Decimal(diffs, v1.Descuento, v2.Descuento, path+".Descuento")

	compareEqualComprobanteConceptosImpuestos(diffs, v1.Impuestos, v2.Impuestos, path+".Impuestos")
	compareConceptoACuentaTerceros(diffs, v1.ACuentaTerceros, v2.ACuentaTerceros, path+".ACuentaTerceros")

	l1, l2 := len(v1.InformacionAduanera), len(v2.InformacionAduanera)
	compare.Comparable(diffs, l1, l2, path+".InformacionAduanera.len()")
	if len(v1.InformacionAduanera) == len(v2.InformacionAduanera) {
		for i, ia1 := range v1.InformacionAduanera {
			ia2 := v2.InformacionAduanera[i]
			compare.Comparable(diffs, ia1.NumeroPedimento, ia2.NumeroPedimento, path+fmt.Sprintf(".InformacionAduanera[%d].NumeroPedimento", i))
		}
	}

	l1, l2 = len(v1.CuentaPredial), len(v2.CuentaPredial)
	compare.Comparable(diffs, l1, l2, path+".CuentaPredial.len()")
	if len(v1.CuentaPredial) == len(v2.CuentaPredial) {
		for i, cp1 := range v1.CuentaPredial {
			cp2 := v2.CuentaPredial[i]
			compare.Comparable(diffs, cp1.Numero, cp2.Numero, path+fmt.Sprintf(".CuentaPredial[%d].Numero", i))
		}
	}

	l1, l2 = len(v1.Parte), len(v2.Parte)
	compare.Comparable(diffs, l1, l2, path+".Parte.len()")
	if l1 == l2 {
		for i, p1 := range v1.Parte {
			p2 := v2.Parte[i]
			compare.Decimal(diffs, p1.Cantidad, p2.Cantidad, path+fmt.Sprintf(".Parte[%d].Cantidad", i))
			compare.Comparable(diffs, p1.NoIdentificacion, p2.NoIdentificacion, path+fmt.Sprintf(".Parte[%d].NoIdentificacion", i))
			compare.Comparable(diffs, p1.Unidad, p2.Unidad, path+fmt.Sprintf(".Parte[%d].Unidad", i))
			compare.Comparable(diffs, p1.Descripcion, p2.Descripcion, path+fmt.Sprintf(".Parte[%d].Descripcion", i))
			compare.Decimal(diffs, p1.ValorUnitario, p2.ValorUnitario, path+fmt.Sprintf(".Parte[%d].ValorUnitario", i))
			compare.Decimal(diffs, p1.Importe, p2.Importe, path+fmt.Sprintf(".Parte[%d].Importe", i))
		}
	}

	l1, l2 = len(v1.InformacionAduanera), len(v2.InformacionAduanera)
	compare.Comparable(diffs, l1, l2, path+".InformacionAduanera.len()")
	if len(v1.InformacionAduanera) == len(v2.InformacionAduanera) {
		for i, ia1 := range v1.InformacionAduanera {
			ia2 := v2.InformacionAduanera[i]
			compare.Comparable(diffs, ia1.NumeroPedimento, ia2.NumeroPedimento, path+fmt.Sprintf(".InformacionAduanera[%d].NumeroPedimento", i))
		}
	}

	l1, l2 = len(v1.CuentaPredial), len(v2.CuentaPredial)
	compare.Comparable(diffs, l1, l2, path+".CuentaPredial.len()")
	if l1 == l2 {
		for i, cp1 := range v1.CuentaPredial {
			cp2 := v2.CuentaPredial[i]
			compare.Comparable(diffs, cp1.Numero, cp2.Numero, path+fmt.Sprintf(".CuentaPredial[%d].Numero", i))
		}
	}
}

func compareConceptoACuentaTerceros(diffs *compare.Diffs, v1, v2 *ConceptoACuentaTerceros, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.RfcACuentaTerceros, v2.RfcACuentaTerceros, path+".RfcACuentaTerceros")
	compare.Comparable(diffs, v1.NombreACuentaTerceros, v2.NombreACuentaTerceros, path+".NombreACuentaTerceros")
	compare.Comparable(diffs, v1.RegimenFiscalACuentaTerceros, v2.RegimenFiscalACuentaTerceros, path+".RegimenFiscalACuentaTerceros")
	compare.Comparable(diffs, v1.DomicilioFiscalACuentaTerceros, v2.DomicilioFiscalACuentaTerceros, path+".DomicilioFiscalACuentaTerceros")

}

func compareEqualComprobanteConceptosImpuestos(diffs *compare.Diffs, v1, v2 *ConceptoImpuestos, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	l1, l2 := len(v1.Traslados), len(v2.Traslados)
	compare.Comparable(diffs, l1, l2, path+".Traslados.len()")
	if l1 == l2 {
		for i, tras1 := range v1.Traslados {
			tras2 := v2.Traslados[i]
			compare.Decimal(diffs, tras1.Base, tras2.Base, path+fmt.Sprintf(".Traslados[%d].Base", i))
			compare.Comparable(diffs, tras1.Impuesto, tras2.Impuesto, path+fmt.Sprintf(".Traslados[%d].Impuesto", i))
			compare.Comparable(diffs, tras1.TipoFactor, tras2.TipoFactor, path+fmt.Sprintf(".Traslados[%d].TipoFactor", i))
			compare.Decimal(diffs, tras1.TasaOCuota, tras2.TasaOCuota, path+fmt.Sprintf(".Traslados[%d].TasaOCuota", i))
			compare.Decimal(diffs, tras1.Importe, tras2.Importe, path+fmt.Sprintf(".Traslados[%d].Importe", i))

		}
	}

	l1, l2 = len(v1.Retenciones), len(v2.Retenciones)
	compare.Comparable(diffs, l1, l2, path+".Retenciones.len()")
	if l1 == l2 {
		for i, ret1 := range v1.Retenciones {
			ret2 := v2.Retenciones[i]
			compare.Decimal(diffs, ret1.Base, ret2.Base, path+fmt.Sprintf(".Retenciones[%d].Base", i))
			compare.Comparable(diffs, ret1.Impuesto, ret2.Impuesto, path+fmt.Sprintf(".Retenciones[%d].Impuesto", i))
			compare.Comparable(diffs, ret1.TipoFactor, ret2.TipoFactor, path+fmt.Sprintf(".Retenciones[%d].TipoFactor", i))
			compare.Decimal(diffs, ret1.TasaOCuota, ret2.TasaOCuota, path+fmt.Sprintf(".Retenciones[%d].TasaOCuota", i))
			compare.Decimal(diffs, ret1.Importe, ret2.Importe, path+fmt.Sprintf(".Retenciones[%d].Importe", i))
		}
	}
}
