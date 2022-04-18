package pagos20

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/compare"
	"github.com/veyronifs/cfdi-go/types"
)

func TestCalcTotales(t *testing.T) {
	tests := []struct {
		name  string
		pagos *Pagos
		want  *Totales
	}{
		{ // PagosSimple
			name: "PagosSimple",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(100.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(100.00),
								ImpPagado:        decimal.NewFromFloat(100.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos: decimal.NewFromFloat(100.00),
			},
		},
		{ // Iva16Simple
			name: "Iva16Simple",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(116.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(116.00),
								ImpPagado:        decimal.NewFromFloat(116.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									TrasladosDR: TrasladosDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIVA,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.16),
											ImporteDR:    decimal.NewFromFloat(16.00),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos:             decimal.NewFromFloat(116.00),
				TotalTrasladosBaseIVA16:     decimal.NewNullDecimal(decimal.NewFromFloat(100.00)),
				TotalTrasladosImpuestoIVA16: decimal.NewNullDecimal(decimal.NewFromFloat(16.00)),
			},
		},
		{ // Iva8Simple
			name: "Iva8Simple",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(108.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(108.00),
								ImpPagado:        decimal.NewFromFloat(108.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									TrasladosDR: TrasladosDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIVA,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.08),
											ImporteDR:    decimal.NewFromFloat(8.00),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos:            decimal.NewFromFloat(108.00),
				TotalTrasladosBaseIVA8:     decimal.NewNullDecimal(decimal.NewFromFloat(100.00)),
				TotalTrasladosImpuestoIVA8: decimal.NewNullDecimal(decimal.NewFromFloat(8.00)),
			},
		},
		{ // Iva0Simple
			name: "Iva0Simple",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(100.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(100.00),
								ImpPagado:        decimal.NewFromFloat(100.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									TrasladosDR: TrasladosDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIVA,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.00),
											ImporteDR:    decimal.NewFromFloat(0.00),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos:            decimal.NewFromFloat(100.00),
				TotalTrasladosBaseIVA0:     decimal.NewNullDecimal(decimal.NewFromFloat(100.00)),
				TotalTrasladosImpuestoIVA0: decimal.NewNullDecimal(decimal.NewFromFloat(0.00)),
			},
		},
		{ // IvaExento
			name: "IvaExento",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(100.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(100.00),
								ImpPagado:        decimal.NewFromFloat(100.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									TrasladosDR: TrasladosDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIVA,
											TipoFactorDR: types.TipoFactorExento,
											TasaOCuotaDR: decimal.NewFromFloat(0.00),
											ImporteDR:    decimal.NewFromFloat(0.00),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos:             decimal.NewFromFloat(100.00),
				TotalTrasladosBaseIVAExento: decimal.NewNullDecimal(decimal.NewFromFloat(100.00)),
			},
		},
		{ // IepsSimple
			name: "IepsSimple",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(108.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(108.00),
								ImpPagado:        decimal.NewFromFloat(108.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									TrasladosDR: TrasladosDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIEPS,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.08),
											ImporteDR:    decimal.NewFromFloat(8.00),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos: decimal.NewFromFloat(108.00),
			},
		},
		{ // Honorarios
			name: "Honorarios",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(95.33),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(95.33),
								ImpPagado:        decimal.NewFromFloat(95.33),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									TrasladosDR: TrasladosDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIVA,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.16),
											ImporteDR:    decimal.NewFromFloat(16.00),
										},
									},
									RetencionesDR: RetencionesDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoISR,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.10),
											ImporteDR:    decimal.NewFromFloat(10.00),
										},
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIVA,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.10),
											ImporteDR:    decimal.NewFromFloat(10.67),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos:             decimal.NewFromFloat(95.33),
				TotalTrasladosBaseIVA16:     decimal.NewNullDecimal(decimal.NewFromFloat(100.00)),
				TotalTrasladosImpuestoIVA16: decimal.NewNullDecimal(decimal.NewFromFloat(16.00)),
				TotalRetencionesISR:         decimal.NewNullDecimal(decimal.NewFromFloat(10.00)),
				TotalRetencionesIVA:         decimal.NewNullDecimal(decimal.NewFromFloat(10.67)),
			},
		},
		{ // RetencionIepsSimple
			name: "RetencionIepsSimple",
			pagos: &Pagos{
				Pago: []*Pago{
					{
						MonedaP: types.MonedaMXN,
						Monto:   decimal.NewFromFloat(92.00),
						DoctoRelacionado: []*DoctoRelacionado{
							{
								ImpSaldoAnt:      decimal.NewFromFloat(92.00),
								ImpPagado:        decimal.NewFromFloat(92.00),
								ImpSaldoInsoluto: decimal.Zero,
								MonedaDR:         types.MonedaMXN,
								ImpuestosDR: &ImpuestosDR{
									RetencionesDR: RetencionesDR{
										{
											BaseDR:       decimal.NewFromFloat(100.00),
											ImpuestoDR:   types.ImpuestoIEPS,
											TipoFactorDR: types.TipoFactorTasa,
											TasaOCuotaDR: decimal.NewFromFloat(0.08),
											ImporteDR:    decimal.NewFromFloat(8.00),
										},
									},
								},
							},
						},
					},
				},
			},
			want: &Totales{
				MontoTotalPagos:      decimal.NewFromFloat(92.00),
				TotalRetencionesIEPS: decimal.NewNullDecimal(decimal.NewFromFloat(8.00)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diffs := compare.NewDiffs()
			got := CalcTotales(tt.pagos)
			compareEqualTotales(diffs, got, tt.want, "TestCalcTotales")
			if err := diffs.Err(); err != nil {
				t.Errorf("CalcTotales()\n%s", err)
			}
		})
	}
}

func Test_decimalMxn(t *testing.T) {
	tests := []struct {
		name       string
		importe    decimal.Decimal
		moneda     types.Moneda
		tipoCambio decimal.Decimal
		want       decimal.Decimal
	}{
		{
			name:       "MXN",
			importe:    decimal.NewFromFloat(100.00),
			moneda:     types.MonedaMXN,
			tipoCambio: decimal.NewFromFloat(1.00),
			want:       decimal.NewFromFloat(100.00),
		},
		{
			name:       "USD",
			importe:    decimal.NewFromFloat(100.00),
			moneda:     types.MonedaUSD,
			tipoCambio: decimal.NewFromFloat(20.00),
			want:       decimal.NewFromFloat(2000.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decimalMxn(tt.importe, tt.moneda, tt.tipoCambio); !got.Equal(tt.want) {
				t.Errorf("decimalMxn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImpuestoPAddRetencionTraslado(t *testing.T) {
	impP := &ImpuestosP{}

	decimalF := decimal.NewFromFloat
	impP.AddRetencion(types.ImpuestoIVA, decimalF(106.67))
	impP.AddRetencion(types.ImpuestoISR, decimalF(100))
	impP.AddRetencion(types.ImpuestoIVA, decimalF(106.67))
	impP.AddRetencion(types.ImpuestoISR, decimalF(100))
	impP.AddRetencion(types.ImpuestoISR, decimalF(400))

	impP.AddTraslado(decimalF(1000), types.ImpuestoIVA, types.TipoFactorTasa, decimalF(0.16), decimalF(160))
	impP.AddTraslado(decimalF(1000), types.ImpuestoIVA, types.TipoFactorTasa, decimalF(0.16), decimalF(160))
	impP.AddTraslado(decimalF(1000), types.ImpuestoIVA, types.TipoFactorTasa, decimalF(0.08), decimalF(80))
	impP.AddTraslado(decimalF(1000), types.ImpuestoIVA, types.TipoFactorTasa, decimalF(0.00), decimalF(0))

	diffs := compare.NewDiffs()
	compareEqualImpuestosP(diffs, impP, &ImpuestosP{
		RetencionesP: RetencionesP{
			{
				ImpuestoP: types.ImpuestoIVA,
				ImporteP:  decimalF(213.34),
			},
			{
				ImpuestoP: types.ImpuestoISR,
				ImporteP:  decimalF(600),
			},
		},
		TrasladosP: TrasladosP{
			{
				BaseP:       decimalF(2000),
				ImpuestoP:   types.ImpuestoIVA,
				TipoFactorP: types.TipoFactorTasa,
				TasaOCuotaP: decimalF(0.16),
				ImporteP:    decimalF(320),
			},
			{
				BaseP:       decimalF(1000),
				ImpuestoP:   types.ImpuestoIVA,
				TipoFactorP: types.TipoFactorTasa,
				TasaOCuotaP: decimalF(0.08),
				ImporteP:    decimalF(80),
			},
			{
				BaseP:       decimalF(1000),
				ImpuestoP:   types.ImpuestoIVA,
				TipoFactorP: types.TipoFactorTasa,
				TasaOCuotaP: decimalF(0.00),
				ImporteP:    decimalF(0),
			},
		},
	}, t.Name())

	if err := diffs.Err(); err != nil {
		t.Errorf("%s\n%s", t.Name(), err)
	}

}
