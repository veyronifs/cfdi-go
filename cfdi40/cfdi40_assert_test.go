package cfdi40

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/complemento/cartaporte20"
	"github.com/veyronifs/cfdi-go/complemento/tfd11"
)

func AssertEqualComprobante(t *testing.T, c1, c2 *Comprobante) {
	assert.Equal(t, c1.Version, c2.Version, ".Version")
	assert.Equal(t, c1.Serie, c2.Serie, ".Serie")
	assert.Equal(t, c1.Folio, c2.Folio, ".Folio")
	assert.Equal(t, c1.Fecha.Encode(), c2.Fecha.Encode(), ".Fecha")
	assert.Equal(t, c1.Sello, c2.Sello, ".Sello")
	assert.Equal(t, c1.FormaPago, c2.FormaPago, ".FormaPago")
	assert.Equal(t, c1.NoCertificado, c2.NoCertificado, ".NoCertificado")
	assert.Equal(t, c1.Certificado, c2.Certificado, ".Certificado")
	assert.Equal(t, c1.CondicionesDePago, c2.CondicionesDePago, ".CondicionesDePago")
	assert.Equal(t, c1.Moneda, c2.Moneda, ".Moneda")
	assert.Equal(t, c1.TipoDeComprobante, c2.TipoDeComprobante, ".TipoDeComprobante")
	assert.Equal(t, c1.Exportacion, c2.Exportacion, ".Exportacion")
	assert.Equal(t, c1.MetodoPago, c2.MetodoPago, ".MetodoPago")
	assert.Equal(t, c1.LugarExpedicion, c2.LugarExpedicion, ".LugarExpedicion")
	assert.Equal(t, c1.Confirmacion, c2.Confirmacion, ".Confirmacion")
	assert.True(t, c1.SubTotal.Equal(c2.SubTotal), ".SubTotal %s != %s", c1.SubTotal.String(), c2.SubTotal.String())
	assert.True(t, c1.Descuento.Equal(c2.Descuento), ".Descuento %s != %s", c1.Descuento.String(), c2.Descuento.String())
	assert.True(t, c1.TipoCambio.Equal(c2.TipoCambio), ".TipoCambio %s != %s", c1.TipoCambio.String(), c2.TipoCambio.String())
	assert.True(t, c1.Total.Equal(c2.Total), ".Total %s != %s", c1.Total.String(), c2.Total.String())

	assertEqualInformacionGlobal(t, c1.InformacionGlobal, c2.InformacionGlobal)
	assertEqualEmisor(t, &c1.Emisor, &c2.Emisor)
	assertEqualReceptor(t, &c1.Receptor, &c2.Receptor)

	assertEqualCfdiRelacionados(t, c1.CfdiRelacionados, c2.CfdiRelacionados)

	l1, l2 := len(c1.Conceptos), len(c2.Conceptos)
	assert.Equal(t, l1, l2, ".Conceptos len %d != %d", l1, l2)
	for i, conc1 := range c1.Conceptos {
		assertEqualComprobanteConceptos(t, conc1, c2.Conceptos[i], fmt.Sprintf(".Conceptos[%d]", i))
	}

	assertEqualImpuestos(t, c1.Impuestos, c2.Impuestos)
	assertComplemento(t, c1.Complemento, c2.Complemento)
}

func assertEqualInformacionGlobal(t *testing.T, v1, v2 *InformacionGlobal) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, ".PeriodicidadPago")
		assert.Nil(t, v2, ".PeriodicidadPago")
		return
	}

	assert.Equal(t, v1.Periodicidad, v2.Periodicidad, ".InformacionGlobal.Periodicidad")
	assert.Equal(t, v1.Meses, v2.Meses, ".InformacionGlobal.Meses")
	assert.Equal(t, v1.Anio, v2.Anio, ".InformacionGlobal.Anio")
}

func assertEqualEmisor(t *testing.T, v1, v2 *Emisor) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, "Emisor is nil")
		assert.Nil(t, v2, "Emisor is nil")
		return
	}

	assert.Equal(t, v1.Rfc, v2.Rfc, ".Emisor.Rfc")
	assert.Equal(t, v1.Nombre, v2.Nombre, ".Emisor.Nombre")
	assert.Equal(t, v1.RegimenFiscal, v2.RegimenFiscal, ".Emisor.RegimenFiscal")
	assert.Equal(t, v1.FacAtrAdquirente, v2.FacAtrAdquirente, ".Emisor.FacAtrAdquirente")
}

func assertEqualReceptor(t *testing.T, v1, v2 *Receptor) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, "Receptor is nil")
		assert.Nil(t, v2, "Receptor is nil")
		return
	}

	assert.Equal(t, v1.Rfc, v2.Rfc, ".Receptor.Rfc")
	assert.Equal(t, v1.Nombre, v2.Nombre, ".Receptor.Nombre")
	assert.Equal(t, v1.DomicilioFiscalReceptor, v2.DomicilioFiscalReceptor, ".Receptor.DomicilioFiscalReceptor")
	assert.Equal(t, v1.ResidenciaFiscal, v2.ResidenciaFiscal, ".Receptor.ResidenciaFiscal")
	assert.Equal(t, v1.NumRegIdTrib, v2.NumRegIdTrib, ".Receptor.NumRegIdTrib")
	assert.Equal(t, v1.RegimenFiscalReceptor, v2.RegimenFiscalReceptor, ".Receptor.RegimenFiscalReceptor")
	assert.Equal(t, v1.UsoCFDI, v2.UsoCFDI, ".Receptor.UsoCFDI")
}

func assertComplemento(t *testing.T, c1, c2 *Complemento) {
	if c1 == nil || c2 == nil {
		assert.Nil(t, c1, "c1")
		assert.Nil(t, c2, "c2")
		return
	}
	cartaporte20.AssertEqual(t, c1.CartaPorte20, c2.CartaPorte20)
	tfd11.AssertEqual(t, c1.TFD11, c2.TFD11)
}

func assertEqualCfdiRelacionados(t *testing.T, r1, r2 []*CfdiRelacionados) {
	l1, l2 := len(r1), len(r2)
	assert.Equal(t, l1, l2, ".CfdiRelacionados len %d != %d", l1, l2)
	if l1 != l2 {
		return
	}

	for i, rel1 := range r1 {
		rel2 := r2[i]
		assert.Equal(t, rel1.TipoRelacion, rel2.TipoRelacion, ".CfdiRelacionados[%d].TipoRelacion", i)
		l1, l2 = len(rel1.CfdiRelacionado), len(rel2.CfdiRelacionado)
		assert.Equal(t, l1, l2, ".CfdiRelacionados[%d].CfdiRelacionado len %d != %d", i, l1, l2)
		if l1 != l2 {
			continue
		}
		for j, cfdi1 := range rel1.CfdiRelacionado {
			cfdi2 := rel2.CfdiRelacionado[j]
			assert.Equal(t, cfdi1.UUID, cfdi2.UUID, ".CfdiRelacionados[%d].CfdiRelacionado[%d].UUID", i, j)
		}
	}
}

func assertEqualImpuestos(t *testing.T, v1 *Impuestos, v2 *Impuestos) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, ".Impuestos")
		assert.Nil(t, v2, ".Impuestos")
		return
	}

	assert.True(t, v1.TotalImpuestosTrasladados.Equal(v2.TotalImpuestosTrasladados), ".Impuestos.TotalImpuestosTrasladados %s != %s", v1.TotalImpuestosTrasladados, v2.TotalImpuestosTrasladados)
	assert.True(t, v1.TotalImpuestosRetenidos.Equal(v2.TotalImpuestosRetenidos), ".Impuestos.TotalImpuestosRetenidos %s != %s", v1.TotalImpuestosRetenidos, v2.TotalImpuestosRetenidos)

	l1, l2 := len(v1.Traslados), len(v2.Traslados)
	assert.Equal(t, l1, l2, ".Impuestos.Traslados len %d != %d", l1, l2)
	if len(v1.Traslados) == len(v2.Traslados) {
		for i, tras1 := range v1.Traslados {
			tras2 := v2.Traslados[i]
			assert.Equal(t, tras1.Impuesto, tras2.Impuesto, ".Impuestos.Traslados[%d].Impuesto", i)
			assert.Equal(t, tras1.TipoFactor, tras2.TipoFactor, ".Impuestos.Traslados[%d].TipoFactor", i)
			assert.True(t, tras1.TasaOCuota.Equal(tras2.TasaOCuota), ".Impuestos.Traslados[%d].TasaOCuota %s != %s", i, tras1.TasaOCuota, tras2.TasaOCuota)
			assert.True(t, tras1.Importe.Equal(tras2.Importe), ".Impuestos.Traslados[%d].Importe %s != %s", i, tras1.Importe, tras2.Importe)
		}
	}
	l1, l2 = len(v1.Retenciones), len(v2.Retenciones)
	assert.Equal(t, l1, l1, ".Impuestos.Retenciones len %d != %d", l1, l2)
	if len(v1.Retenciones) == len(v2.Retenciones) {
		for i, ret1 := range v1.Retenciones {
			ret2 := v2.Retenciones[i]
			assert.Equal(t, ret1.Impuesto, ret2.Impuesto, ".Impuestos.Retenciones[%d].Impuesto", i)
			assert.True(t, ret1.Importe.Equal(ret2.Importe), ".Impuestos.Retenciones[%d].Importe %s != %s", i, ret1.Importe, ret2.Importe)
		}
	}
}

func assertEqualComprobanteConceptos(t *testing.T, c1, c2 *Concepto, path string) {
	assert.Equal(t, c1.ClaveProdServ, c2.ClaveProdServ, path+".ClaveProdServ")
	assert.Equal(t, c1.NoIdentificacion, c2.NoIdentificacion, path+".NoIdentificacion")
	assert.Equal(t, c1.ClaveUnidad, c2.ClaveUnidad, path+".ClaveUnidad")
	assert.Equal(t, c1.Unidad, c2.Unidad, path+".Unidad")
	assert.Equal(t, c1.Descripcion, c2.Descripcion, path+".Descripcion")
	assert.Equal(t, c1.ObjetoImp, c2.ObjetoImp, path+".ObjetoImp")
	assert.True(t, c1.Cantidad.Equal(c2.Cantidad), path+".Cantidad %s != %s", c1.Cantidad, c2.Cantidad)
	assert.True(t, c1.ValorUnitario.Equal(c2.ValorUnitario), path+".ValorUnitario %s != %s", c1.ValorUnitario, c2.ValorUnitario)
	assert.True(t, c1.Importe.Equal(c2.Importe), path+".Importe %s != %s", c1.Importe, c2.Importe)
	assert.True(t, c1.Descuento.Equal(c2.Descuento), path+".Descuento %s != %s", c1.Descuento, c2.Descuento)

	assertEqualComprobanteConceptosImpuestos(t, c1.Impuestos, c2.Impuestos, path+".Impuestos")
	assertConceptoACuentaTerceros(t, c1.ACuentaTerceros, c2.ACuentaTerceros, path+".ACuentaTerceros")

	l1, l2 := len(c1.InformacionAduanera), len(c2.InformacionAduanera)
	assert.Equal(t, l1, l2, path+".InformacionAduanera len %d != %d", l1, l2)
	if len(c1.InformacionAduanera) == len(c2.InformacionAduanera) {
		for i, ia1 := range c1.InformacionAduanera {
			ia2 := c2.InformacionAduanera[i]
			assert.Equal(t, ia1.NumeroPedimento, ia2.NumeroPedimento, path+".InformacionAduanera[%d].NumeroPedimento", i)
		}
	}

	l1, l2 = len(c1.CuentaPredial), len(c2.CuentaPredial)
	assert.Equal(t, l1, l2, path+".CuentaPredial len %d != %d", l1, l2)
	if len(c1.CuentaPredial) == len(c2.CuentaPredial) {
		for i, cp1 := range c1.CuentaPredial {
			cp2 := c2.CuentaPredial[i]
			assert.Equal(t, cp1.Numero, cp2.Numero, path+".CuentaPredial[%d].Numero", i)
		}
	}

	l1, l2 = len(c1.Parte), len(c2.Parte)
	assert.Equal(t, l1, l2, path+".Parte len %d != %d", l1, l2)
	if len(c1.Parte) == len(c2.Parte) {
		for i, p1 := range c1.Parte {
			p2 := c2.Parte[i]
			assert.Equal(t, p1.Cantidad, p2.Cantidad, path+".Parte[%d].Cantidad", i)
			assert.Equal(t, p1.NoIdentificacion, p2.NoIdentificacion, path+".Parte[%d].NoIdentificacion", i)
			assert.Equal(t, p1.Unidad, p2.Unidad, path+".Parte[%d].Unidad", i)
			assert.Equal(t, p1.Descripcion, p2.Descripcion, path+".Parte[%d].Descripcion", i)
			assert.Equal(t, p1.ValorUnitario, p2.ValorUnitario, path+".Parte[%d].ValorUnitario", i)
			assert.Equal(t, p1.Importe, p2.Importe, path+".Parte[%d].Importe", i)
		}
	}

	l1, l2 = len(c1.InformacionAduanera), len(c2.InformacionAduanera)
	assert.Equal(t, l1, l2, path+".InformacionAduanera")
	if len(c1.InformacionAduanera) == len(c2.InformacionAduanera) {
		for i, ia1 := range c1.InformacionAduanera {
			ia2 := c2.InformacionAduanera[i]
			assert.Equal(t, ia1.NumeroPedimento, ia2.NumeroPedimento, path+".InformacionAduanera[%d].NumeroPedimento", i)
		}
	}

	l1, l2 = len(c1.CuentaPredial), len(c2.CuentaPredial)
	assert.Equal(t, l1, l2, path+".CuentaPredial")
	if len(c1.CuentaPredial) == len(c2.CuentaPredial) {
		for i, cp1 := range c1.CuentaPredial {
			cp2 := c2.CuentaPredial[i]
			assert.Equal(t, cp1.Numero, cp2.Numero, path+".CuentaPredial[%d].Numero", i)
		}
	}
}

func assertConceptoACuentaTerceros(t *testing.T, at1, at2 *ConceptoACuentaTerceros, path string) {
	if at1 == nil {
		assert.Nil(t, at2, path+".ACuentaTerceros")
		return
	}
	assert.NotNil(t, at2, path+".ACuentaTerceros")
	if at2 == nil {
		return
	}

	assert.Equal(t, at1.RfcACuentaTerceros, at2.RfcACuentaTerceros, path+".RfcACuentaTerceros")
	assert.Equal(t, at1.NombreACuentaTerceros, at2.NombreACuentaTerceros, path+".NombreACuentaTerceros")
	assert.Equal(t, at1.RegimenFiscalACuentaTerceros, at2.RegimenFiscalACuentaTerceros, path+".RegimenFiscalACuentaTerceros")
	assert.Equal(t, at1.DomicilioFiscalACuentaTerceros, at2.DomicilioFiscalACuentaTerceros, path+".DomicilioFiscalACuentaTerceros")
}

func assertEqualComprobanteConceptosImpuestos(t *testing.T, imp1, imp2 *ConceptoImpuestos, path string) {
	if imp1 == nil {
		assert.Nil(t, imp2, path+".Impuestos")
		return
	}
	assert.NotNil(t, imp2, path)
	if imp2 == nil {
		return
	}
	l1, l2 := len(imp1.Traslados), len(imp2.Traslados)
	assert.Equal(t, l1, l2, path+".Traslados")
	if len(imp1.Traslados) == len(imp2.Traslados) {
		for i, tras1 := range imp1.Traslados {
			tras2 := imp2.Traslados[i]
			assert.True(t, tras1.Base.Equal(tras2.Base), path+".Traslados[%d].Base %s != %s", i, tras1.Base, tras2.Base)
			assert.Equal(t, tras1.Impuesto, tras2.Impuesto, path+".Traslados[%d].Impuesto", i)
			assert.Equal(t, tras1.TipoFactor, tras2.TipoFactor, path+".Traslados[%d].TipoFactor", i)
			assert.True(t, tras1.TasaOCuota.Equal(tras2.TasaOCuota), path+".Traslados[%d].TasaOCuota %s != %s", i, tras1.TasaOCuota, tras2.TasaOCuota)
			assert.True(t, tras1.Importe.Equal(tras2.Importe), path+".Traslados[%d].Importe %s != %s", i, tras1.Importe, tras2.Importe)
		}
	}

	l1, l2 = len(imp1.Retenciones), len(imp2.Retenciones)
	assert.Equal(t, l1, l2, path+".Retenciones")
	if len(imp1.Retenciones) == len(imp2.Retenciones) {
		for i, ret1 := range imp1.Retenciones {
			ret2 := imp2.Retenciones[i]
			assert.True(t, ret1.Base.Equal(ret2.Base), path+".Retenciones[%d].Base %s != %s", i, ret1.Base, ret2.Base)
			assert.Equal(t, ret1.Impuesto, ret2.Impuesto, path+".Retenciones[%d].Impuesto", i)
			assert.Equal(t, ret1.TipoFactor, ret2.TipoFactor, path+".Retenciones[%d].TipoFactor", i)
			assert.True(t, ret1.TasaOCuota.Equal(ret2.TasaOCuota), path+".Retenciones[%d].TasaOCuota %s != %s", i, ret1.TasaOCuota, ret2.TasaOCuota)
			assert.True(t, ret1.Importe.Equal(ret2.Importe), path+".Retenciones[%d].Importe %s != %s", i, ret1.Importe, ret2.Importe)
		}
	}
}
