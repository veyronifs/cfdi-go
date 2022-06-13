package tests_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/types"
)

func TestEgresoIva16(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteE,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		Fecha:             types.NewFechaHNow(),
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Emisor:            emisor16_8_0,
		Receptor: &cfdi40.Receptor{
			Rfc:                     "BAR011108CC6",
			Nombre:                  "BARCEL",
			DomicilioFiscalReceptor: "52000",
			RegimenFiscalReceptor:   "601",
			UsoCFDI:                 types.UsoCFDICP01,
		},
		Conceptos: cfdi40.Conceptos{
			{
				ObjetoImp:        types.ObjetoImp02,
				Cantidad:         decimal.NewFromFloat(1),
				ClaveProdServ:    "50192100",
				ClaveUnidad:      "XBX",
				Descripcion:      "Cacahuate",
				Importe:          decimal.NewFromFloat(1000),
				NoIdentificacion: "1",
				ValorUnitario:    decimal.NewFromFloat(1000),
				Impuestos: &cfdi40.ConceptoImpuestos{
					Traslados: cfdi40.ConceptoImpuestosTraslados{
						{
							Base:       decimal.NewFromFloat(1000),
							Impuesto:   types.ImpuestoIVA,
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.16),
							Importe:    decimal.NewFromFloat(160),
						},
					},
				},
			},
		},
	}
	c.Impuestos = cfdi40.NewImpuestos(*c)
	c.SubTotal, c.Descuento, c.Total = cfdi40.CalcularTotales(*c)

	testTimbrar(t, c)
}
