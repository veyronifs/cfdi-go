package tests_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/complemento/pagos20"
	"github.com/veyronifs/cfdi-go/types"
)

func TestPagoIva16_Parcialidad1(t *testing.T) {
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
			Pago: []*pagos20.Pago{
				{
					FechaPago:    newFechaHNow2(),
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaMXN,
					Monto:        decimal.NewFromFloat(580),
					TipoCambioP:  decimal.NewFromFloat(1),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb7",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1160),
							ImpPagado:        decimal.NewFromFloat(580),
							ImpSaldoInsoluto: decimal.NewFromFloat(580),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(500),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0.160000),
										ImporteDR:    decimal.NewFromFloat(80),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	cfdiPago.Complemento.Pagos20.Totales = pagos20.CalcTotales(cfdiPago.Complemento.Pagos20)
	pago := cfdiPago.Complemento.Pagos20.Pago[0]
	pago.ImpuestosP = pagos20.CalcImpuestosP(pago)
	testTimbrar(t, cfdiPago)
}

func TestPagoIva16_Parcialidad2(t *testing.T) {
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
			Pago: []*pagos20.Pago{
				{
					FechaPago:    newFechaHNow2(),
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaMXN,
					Monto:        decimal.NewFromFloat(580),
					TipoCambioP:  decimal.NewFromFloat(1),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb7",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   2,
							ImpSaldoAnt:      decimal.NewFromFloat(580),
							ImpPagado:        decimal.NewFromFloat(580),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(500),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0.160000),
										ImporteDR:    decimal.NewFromFloat(80),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	cfdiPago.Complemento.Pagos20.Totales = pagos20.CalcTotales(cfdiPago.Complemento.Pagos20)
	pago := cfdiPago.Complemento.Pagos20.Pago[0]
	pago.ImpuestosP = pagos20.CalcImpuestosP(pago)
	testTimbrar(t, cfdiPago)
}

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

func TestPago2Iva_2Ret_2SinImp(t *testing.T) {
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
			Pago: []*pagos20.Pago{
				{
					FechaPago:    newFechaHNow2(),
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaMXN,
					Monto:        decimal.NewFromFloat(7226.67),
					TipoCambioP:  decimal.NewFromFloat(1),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb1",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1160.00),
							ImpPagado:        decimal.NewFromFloat(1160.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb2",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1160.00),
							ImpPagado:        decimal.NewFromFloat(1160.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb3",
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb4",
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb5",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1000.00),
							ImpPagado:        decimal.NewFromFloat(1000.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0),
										ImporteDR:    decimal.NewFromFloat(0),
									},
								},
							},
						},
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb5",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1000.00),
							ImpPagado:        decimal.NewFromFloat(1000.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0),
										ImporteDR:    decimal.NewFromFloat(0),
									},
								},
							},
						},
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb5",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1000.00),
							ImpPagado:        decimal.NewFromFloat(1000.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorExento,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	cfdiPago.Complemento.Pagos20.Totales = pagos20.CalcTotales(cfdiPago.Complemento.Pagos20)
	pago := cfdiPago.Complemento.Pagos20.Pago[0]
	pago.ImpuestosP = pagos20.CalcImpuestosP(pago)
	testTimbrar(t, cfdiPago)
}

func TestPago2Iva_2Ret_2SinImpUSD(t *testing.T) {
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
			Pago: []*pagos20.Pago{
				{
					FechaPago:    newFechaHNow2(),
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaUSD,
					Monto:        decimal.NewFromFloat(6226.67),
					TipoCambioP:  decimal.NewFromFloat(20),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb1",
							MonedaDR:         types.MonedaUSD,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1160.00),
							ImpPagado:        decimal.NewFromFloat(1160.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb2",
							MonedaDR:         types.MonedaUSD,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1160.00),
							ImpPagado:        decimal.NewFromFloat(1160.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb3",
							MonedaDR:         types.MonedaUSD,
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb4",
							MonedaDR:         types.MonedaUSD,
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
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb5",
							MonedaDR:         types.MonedaUSD,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1000.00),
							ImpPagado:        decimal.NewFromFloat(1000.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0),
										ImporteDR:    decimal.NewFromFloat(0),
									},
								},
							},
						},
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb5",
							MonedaDR:         types.MonedaUSD,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(1000.00),
							ImpPagado:        decimal.NewFromFloat(1000.00),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp02,
							EquivalenciaDR:   decimal.NewFromFloat(1),
							ImpuestosDR: &pagos20.ImpuestosDR{
								TrasladosDR: pagos20.TrasladosDR{
									{
										BaseDR:       decimal.NewFromFloat(1000),
										ImpuestoDR:   types.ImpuestoIVA,
										TipoFactorDR: types.TipoFactorTasa,
										TasaOCuotaDR: decimal.NewFromFloat(0),
										ImporteDR:    decimal.NewFromFloat(0),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	cfdiPago.Complemento.Pagos20.Totales = pagos20.CalcTotales(cfdiPago.Complemento.Pagos20)
	pago := cfdiPago.Complemento.Pagos20.Pago[0]
	pago.ImpuestosP = pagos20.CalcImpuestosP(pago)
	testTimbrar(t, cfdiPago)
}

// TODO: FACTURA MXN y USD
