package tests_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/complemento/pagos20"
	"github.com/veyronifs/cfdi-go/types"
)

func TestPagoHonorarios(t *testing.T) {
	cfdiPago := cfdi40.NewComprobantePago()
	cfdiPago.LugarExpedicion = "20000"
	cfdiPago.Serie = "Serie"
	cfdiPago.Folio = "Folio"
	cfdiPago.Fecha = newFechaHNow2()
	cfdiPago.Emisor = emisor16_8_0
	cfdiPago.Receptor = &cfdi40.Receptor{
		Rfc:                     "BAR011108CC6",
		Nombre:                  "BARCEL",
		DomicilioFiscalReceptor: "52000",
		RegimenFiscalReceptor:   "601",
		UsoCFDI:                 types.UsoCFDICP01,
	}
	cfdiPago.Complemento = &cfdi40.Complemento{
		Pagos20: &pagos20.Pagos{
			Version: "2.0",
			Totales: &pagos20.Totales{
				MontoTotalPagos:             decimal.NewFromFloat(953.33),
				TotalRetencionesIVA:         decimal.NewNullDecimal(decimal.NewFromFloat(106.67)),
				TotalRetencionesISR:         decimal.NewNullDecimal(decimal.NewFromFloat(100)),
				TotalTrasladosBaseIVA16:     decimal.NewNullDecimal(decimal.NewFromFloat(1000)),
				TotalTrasladosImpuestoIVA16: decimal.NewNullDecimal(decimal.NewFromFloat(160)),
			},
			Pago: []*pagos20.Pago{
				{
					FechaPago:    newFechaHNow2(),
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaMXN,
					Monto:        decimal.NewFromFloat(953.33),
					TipoCambioP:  decimal.NewFromFloat(1),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb7",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(953.33),
							ImpPagado:        decimal.NewFromFloat(953.33),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								RetencionesDR: pagos20.RetencionesDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoISR,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0.100000),
										ImporteDR:    decimal.NewFromFloat(100),
									},
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0.106666),
										ImporteDR:    decimal.NewFromFloat(106.67),
									},
								},
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0.160000),
										ImporteDR:    decimal.NewFromFloat(160),
									},
								},
							},
						},
					},
					ImpuestosP: &pagos20.ImpuestosP{
						RetencionesP: pagos20.RetencionesP{
							{
								ImpuestoP: types.ImpuestoISR,
								ImporteP:  decimal.NewFromFloat(100),
							},
							{
								ImpuestoP: types.ImpuestoIVA,
								ImporteP:  decimal.NewFromFloat(106.67),
							},
						},
						TrasladosP: pagos20.TrasladosP{
							{
								BaseP:       decimal.NewFromFloat(1000),
								ImpuestoP:   types.ImpuestoIVA,
								TipoFactorP: types.TipoFactorTasa,
								TasaOCuotaP: decimal.NewFromFloat(0.160000),
								ImporteP:    decimal.NewFromFloat(160),
							},
						},
					},
				},
			},
		},
	}

	testTimbrar(t, cfdiPago)
}

func TestPagoSinImpuestos(t *testing.T) {
	cfdiPago := cfdi40.NewComprobantePago()
	cfdiPago.LugarExpedicion = "20000"
	cfdiPago.Serie = "Serie"
	cfdiPago.Folio = "Folio"
	cfdiPago.Fecha = newFechaHNow2()
	cfdiPago.Emisor = emisor16_8_0
	cfdiPago.Receptor = &cfdi40.Receptor{
		Rfc:                     "BAR011108CC6",
		Nombre:                  "BARCEL",
		DomicilioFiscalReceptor: "52000",
		RegimenFiscalReceptor:   "601",
		UsoCFDI:                 types.UsoCFDICP01,
	}
	cfdiPago.Complemento = &cfdi40.Complemento{
		Pagos20: &pagos20.Pagos{
			Version: "2.0",
			Totales: &pagos20.Totales{
				MontoTotalPagos: decimal.NewFromFloat(953.33),
			},
			Pago: []*pagos20.Pago{
				{
					FechaPago:    newFechaHNow2(),
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaMXN,
					Monto:        decimal.NewFromFloat(953.33),
					TipoCambioP:  decimal.NewFromFloat(1),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb7",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(953.33),
							ImpPagado:        decimal.NewFromFloat(953.33),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp01,
							EquivalenciaDR:   decimal.NewFromFloat(1),
						},
					},
				},
			},
		},
	}

	testTimbrar(t, cfdiPago)
}
