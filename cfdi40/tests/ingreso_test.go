package tests_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/types"
)

func TestIngresoIva16(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             types.NewFechaHNow(),
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

func TestIngresoIva0(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             types.NewFechaHNow(),
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
							TasaOCuota: decimal.NewFromFloat(0.0),
							Importe:    decimal.NewFromFloat(0.0),
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

func TestIngresoIva16_Exento(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             types.NewFechaHNow(),
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
							TipoFactor: types.TipoFactorExento,
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

func TestIngresoIvaExento(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             types.NewFechaHNow(),
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
							TipoFactor: types.TipoFactorExento,
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

func TestIngresoIva16_0_exento(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             newFechaHNow2(),
		Total:             decimal.NewFromFloat(1160),
		SubTotal:          decimal.NewFromFloat(1000),
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
							TasaOCuota: decimal.NewFromFloat(0.0),
							Importe:    decimal.NewFromFloat(0),
						},
					},
				},
			},
			{
				ObjetoImp:        types.ObjetoImp02,
				Cantidad:         decimal.NewFromFloat(1),
				ClaveProdServ:    "50192100",
				ClaveUnidad:      "XBX",
				Descripcion:      "Cacahuate",
				Importe:          decimal.NewFromFloat(2000),
				NoIdentificacion: "1",
				ValorUnitario:    decimal.NewFromFloat(2000),
				Impuestos: &cfdi40.ConceptoImpuestos{
					Traslados: cfdi40.ConceptoImpuestosTraslados{
						{
							Base:       decimal.NewFromFloat(2000),
							Impuesto:   types.ImpuestoIVA,
							TipoFactor: types.TipoFactorExento,
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

func TestIngresoIva8(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "32697",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             newFechaHNow2(),
		Emisor:            emisor16_8_0,
		Receptor: &cfdi40.Receptor{
			Rfc:                     "BAR011108CC6",
			Nombre:                  "BARCEL",
			DomicilioFiscalReceptor: "52000",
			RegimenFiscalReceptor:   types.RegimenFiscal601,
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
							TasaOCuota: decimal.NewFromFloat(0.08),
							Importe:    decimal.NewFromFloat(80),
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

func TestIngresoIva16_8_0(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "32697",
		Serie:             "Serie",
		Folio:             "Folio",
		FormaPago:         types.FormaPago01,
		MetodoPago:        types.MetodoPagoPUE,
		Fecha:             newFechaHNow2(),
		Total:             decimal.NewFromFloat(1160),
		SubTotal:          decimal.NewFromFloat(1000),
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
							TasaOCuota: decimal.NewFromFloat(0.08),
							Importe:    decimal.NewFromFloat(80),
						},
					},
				},
			},
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
							TasaOCuota: decimal.NewFromFloat(0.0),
							Importe:    decimal.NewFromFloat(0),
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
