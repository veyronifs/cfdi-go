package pagos20

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/types"
)

func fecha(s string) types.FechaH {
	t, err := types.NewFechaH(s)
	if err != nil {
		panic(err)
	}
	return t
}
func TestUnmarshal(t *testing.T) {
	var originalXML []byte
	{
		originalXML = []byte(`
		<pago20:Pagos Version="2.0">
            <pago20:Totales TotalRetencionesIVA="1" TotalRetencionesISR="1" TotalTrasladosBaseIVA16="100" TotalTrasladosImpuestoIVA16="16" TotalTrasladosBaseIVA8="100" TotalTrasladosImpuestoIVA8="8" MontoTotalPagos="1000" />
            <pago20:Pago FechaPago="2022-01-12T12:00:00" FormaDePagoP="03" MonedaP="MXN" TipoCambioP="1" Monto="1000" NumOperacion="123456789" RfcEmisorCtaBen="BBA940707IE1" CtaBeneficiario="0006985412">
                <pago20:DoctoRelacionado IdDocumento="3CACDB80-D1C0-4CB9-B312-FDC2193ECB32" Serie="A" Folio="88" MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1" ImpSaldoAnt="500" ImpPagado="500" ImpSaldoInsoluto="0" ObjetoImpDR="02">
                    <pago20:ImpuestosDR>
                        <pago20:RetencionesDR>
                            <pago20:RetencionDR BaseDR="10" ImpuestoDR="001" TipoFactorDR="Tasa" TasaOCuotaDR="0.1" ImporteDR="1" />
                            <pago20:RetencionDR BaseDR="10" ImpuestoDR="002" TipoFactorDR="Tasa" TasaOCuotaDR="0.1" ImporteDR="1" />
                        </pago20:RetencionesDR>
                        <pago20:TrasladosDR>
                            <pago20:TrasladoDR BaseDR="100" ImpuestoDR="002" TipoFactorDR="Tasa" TasaOCuotaDR="0.16" ImporteDR="16" />
                            <pago20:TrasladoDR BaseDR="100" ImpuestoDR="002" TipoFactorDR="Tasa" TasaOCuotaDR="0.08" ImporteDR="8" />
                        </pago20:TrasladosDR>
                    </pago20:ImpuestosDR>
                </pago20:DoctoRelacionado>
                <pago20:DoctoRelacionado IdDocumento="07e28c28-9328-43da-b17e-02a1edc684d9" Serie="F" Folio="99" MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1" ImpSaldoAnt="500" ImpPagado="500" ImpSaldoInsoluto="0" ObjetoImpDR="01"></pago20:DoctoRelacionado>
                <pago20:ImpuestosP>
                    <pago20:RetencionesP>
                        <pago20:RetencionP ImpuestoP="001" ImporteP="1" />
                        <pago20:RetencionP ImpuestoP="002" ImporteP="1" />
                    </pago20:RetencionesP>
                    <pago20:TrasladosP>
                        <pago20:TrasladoP BaseP="100" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.16" ImporteP="16" />
                        <pago20:TrasladoP BaseP="100" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.08" ImporteP="8" />
                    </pago20:TrasladosP>
                </pago20:ImpuestosP>
            </pago20:Pago>
        </pago20:Pagos>
		`)
	}
	pagosUnmarshaled, err := Unmarshal(originalXML)
	if err != nil {
		t.Errorf("Error Unmarshal(originalXML): %s", err)
		return
	}
	expectedPagos := &Pagos{
		Version: "2.0",
		Totales: &Totales{
			TotalRetencionesIVA:         decimal.NewNullDecimal(decimal.NewFromFloat(1)),
			TotalRetencionesISR:         decimal.NewNullDecimal(decimal.NewFromFloat(1)),
			TotalTrasladosBaseIVA16:     decimal.NewNullDecimal(decimal.NewFromFloat(100)),
			TotalTrasladosImpuestoIVA16: decimal.NewNullDecimal(decimal.NewFromFloat(16)),
			TotalTrasladosBaseIVA8:      decimal.NewNullDecimal(decimal.NewFromFloat(100)),
			TotalTrasladosImpuestoIVA8:  decimal.NewNullDecimal(decimal.NewFromFloat(8)),
			MontoTotalPagos:             decimal.NewFromFloat(1000),
		},
		Pago: []*Pago{
			{
				FechaPago:       fecha("2022-01-12T12:00:00"),
				FormaDePagoP:    types.FormaPago03,
				MonedaP:         types.MonedaMXN,
				TipoCambioP:     decimal.NewFromFloat(1),
				Monto:           decimal.NewFromFloat(1000),
				NumOperacion:    "123456789",
				RfcEmisorCtaBen: "BBA940707IE1",
				CtaBeneficiario: "0006985412",
				DoctoRelacionado: []*DoctoRelacionado{
					{
						IdDocumento:      "3CACDB80-D1C0-4CB9-B312-FDC2193ECB32",
						Serie:            "A",
						Folio:            "88",
						MonedaDR:         types.MonedaMXN,
						EquivalenciaDR:   decimal.NewFromFloat(1),
						NumParcialidad:   1,
						ImpSaldoAnt:      decimal.NewFromFloat(500),
						ImpPagado:        decimal.NewFromFloat(500),
						ImpSaldoInsoluto: decimal.NewFromFloat(0),
						ObjetoImpDR:      types.ObjetoImp02,
						ImpuestosDR: &ImpuestosDR{
							RetencionesDR: []*RetencionDR{
								{
									BaseDR:       decimal.NewFromFloat(10),
									ImpuestoDR:   types.ImpuestoISR,
									TipoFactorDR: types.TipoFactorTasa,
									TasaOCuotaDR: decimal.NewFromFloat(0.1),
									ImporteDR:    decimal.NewFromFloat(1),
								},
								{
									BaseDR:       decimal.NewFromFloat(10),
									ImpuestoDR:   types.ImpuestoIVA,
									TipoFactorDR: types.TipoFactorTasa,
									TasaOCuotaDR: decimal.NewFromFloat(0.1),
									ImporteDR:    decimal.NewFromFloat(1),
								},
							},
							TrasladosDR: []*TrasladoDR{
								{
									BaseDR:       decimal.NewFromFloat(100),
									ImpuestoDR:   types.ImpuestoIVA,
									TipoFactorDR: types.TipoFactorTasa,
									TasaOCuotaDR: decimal.NewFromFloat(0.16),
									ImporteDR:    decimal.NewFromFloat(16),
								},
								{
									BaseDR:       decimal.NewFromFloat(100),
									ImpuestoDR:   types.ImpuestoIVA,
									TipoFactorDR: types.TipoFactorTasa,
									TasaOCuotaDR: decimal.NewFromFloat(0.08),
									ImporteDR:    decimal.NewFromFloat(8),
								},
							},
						},
					},
					{
						IdDocumento:      "07e28c28-9328-43da-b17e-02a1edc684d9",
						Serie:            "F",
						Folio:            "99",
						MonedaDR:         types.MonedaMXN,
						EquivalenciaDR:   decimal.NewFromFloat(1),
						NumParcialidad:   1,
						ImpSaldoAnt:      decimal.NewFromFloat(500),
						ImpPagado:        decimal.NewFromFloat(500),
						ImpSaldoInsoluto: decimal.NewFromFloat(0),
						ObjetoImpDR:      types.ObjetoImp01,
					},
				},
				ImpuestosP: &ImpuestosP{
					RetencionesP: []*RetencionP{
						{ImpuestoP: types.ImpuestoISR, ImporteP: decimal.NewFromFloat(1)},
						{ImpuestoP: types.ImpuestoIVA, ImporteP: decimal.NewFromFloat(1)},
					},
					TrasladosP: []*TrasladoP{
						{
							BaseP:       decimal.NewFromFloat(100),
							ImpuestoP:   types.ImpuestoIVA,
							TipoFactorP: types.TipoFactorTasa,
							TasaOCuotaP: decimal.NewFromFloat(0.16),
							ImporteP:    decimal.NewFromFloat(16),
						},
						{
							BaseP:       decimal.NewFromFloat(100),
							ImpuestoP:   types.ImpuestoIVA,
							TipoFactorP: types.TipoFactorTasa,
							TasaOCuotaP: decimal.NewFromFloat(0.08),
							ImporteP:    decimal.NewFromFloat(8),
						},
					},
				},
			},
		},
	}
	err = CompareEqual(expectedPagos, pagosUnmarshaled)
	assert.NoError(t, err)

	t.Run("PagosMarshal", func(t *testing.T) {
		xmlMarshaled, err := Marshal(expectedPagos)
		if err != nil {
			t.Errorf("Error marshaling: %v", err)
			return
		}
		pagosUnmarshaled, err := Unmarshal(xmlMarshaled)
		if err != nil {
			t.Errorf("Error Unmarshal(xmlMarshaled): %s", err)
			return
		}
		err = CompareEqual(expectedPagos, pagosUnmarshaled)
		assert.NoError(t, err)
	})
}

func TestCountDecimals(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"9", 0},
		{"3.14", 2},
		{"123.456", 3},
		{"0.0001", 4},
		{"-5.789", 3},
		{"-.789", 3},
		{"12.34.56", -1}, // Invalid input, decimal point occurs more than once
		{"abc", -1},      // Invalid input, no numeric characters
		{"", -1},         // Invalid input, empty string
		{"123..456", -1}, // Invalid input, multiple consecutive decimal points
		{"1.23a", -1},    // Invalid input, alphanumeric character after the decimal point
	}

	for _, test := range tests {
		result := countDecimals(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", test.input, test.expected, result)
		}
	}
}

func TestLimitesImporteDR(t *testing.T) {
	newDecimal := func(s string) decimal.Decimal {
		d, err := decimal.NewFromString(s)
		if err != nil {
			panic(err)
		}
		return d
	}
	tests := []struct {
		baseDR         decimal.Decimal
		tasaOCuotaDR   decimal.Decimal
		monedaDR       types.Moneda
		limiteInferior decimal.Decimal
		limiteSuperior decimal.Decimal
	}{
		{
			baseDR:         newDecimal("78855.7"),
			tasaOCuotaDR:   newDecimal("0.08"),
			monedaDR:       "MXN",
			limiteInferior: newDecimal("6308.45"),
			limiteSuperior: newDecimal("6308.47"),
		},
		{
			baseDR:         newDecimal("78194.92"),
			tasaOCuotaDR:   newDecimal("0.08"),
			monedaDR:       "MXN",
			limiteInferior: newDecimal("6255.59"),
			limiteSuperior: newDecimal("6255.60"),
		},
	}

	for _, test := range tests {
		//fmt.Println("*********************")
		limiteInferior, limiteSuperior := LimitesImporteDR(test.baseDR, test.tasaOCuotaDR, test.monedaDR)
		assert.True(t, test.limiteInferior.Equal(limiteInferior), "expected limiteInferior: %s, got: %s", test.limiteInferior.String(), limiteInferior.String())
		assert.True(t, test.limiteSuperior.Equal(limiteSuperior), "expected limiteSuperior: %s, got: %s", test.limiteSuperior.String(), limiteSuperior.String())
		//fmt.Println("expected limiteInferior: ", test.limiteInferior.String(), "got: ", limiteInferior.String())
		//fmt.Println("expected limiteSuperior: ", test.limiteSuperior.String(), "got: ", limiteSuperior.String())
		//fmt.Println("*********************")
	}
}
