package cfdi40

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/compare"
	"github.com/veyronifs/cfdi-go/types"
)

func TestNewImpuestos(t *testing.T) {
	tests := []struct {
		TestName    string
		Comprobante Comprobante
		Expected    *Impuestos
	}{
		{ // "No Impuestos (1 concepto)"
			TestName: "No impuestos (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
					},
				},
			},
			Expected: nil,
		},
		{ // "Solo Retenciones (1 concepto)"
			TestName: "Solo Retenciones (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
						},
					},
				},
			},
			Expected: &Impuestos{
				TotalImpuestosRetenidos: decimal.NewFromFloat(20.66),
				Retenciones: []*ImpuestosRetencion{
					{
						Impuesto: types.ImpuestoISR,
						Importe:  decimal.NewFromFloat(10),
					},
					{
						Impuesto: types.ImpuestoIVA,
						Importe:  decimal.NewFromFloat(10.66),
					},
				},
			},
		},
		{ // "Solo Traslados (1 concepto)"
			TestName: "Solo Traslados (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			Expected: &Impuestos{
				TotalImpuestosTrasladados: decimal.NewFromFloat(16),
				Traslados: []*ImpuestosTraslado{
					{
						Base:       decimal.NewFromFloat(100),
						Impuesto:   types.ImpuestoIVA,
						Importe:    decimal.NewFromFloat(16),
						TipoFactor: types.TipoFactorTasa,
						TasaOCuota: decimal.NewFromFloat(0.16),
					},
				},
			},
		},
		{ // "Traslados y Retenciones (1 concepto)"
			TestName: "Traslados y Retenciones (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			Expected: &Impuestos{
				TotalImpuestosRetenidos:   decimal.NewFromFloat(20.66),
				TotalImpuestosTrasladados: decimal.NewFromFloat(16),
				Retenciones: []*ImpuestosRetencion{
					{
						Impuesto: types.ImpuestoISR,
						Importe:  decimal.NewFromFloat(10),
					},
					{
						Impuesto: types.ImpuestoIVA,
						Importe:  decimal.NewFromFloat(10.66),
					},
				},
				Traslados: []*ImpuestosTraslado{
					{
						Base:       decimal.NewFromFloat(100),
						Impuesto:   types.ImpuestoIVA,
						Importe:    decimal.NewFromFloat(16),
						TipoFactor: types.TipoFactorTasa,
						TasaOCuota: decimal.NewFromFloat(0.16),
					},
				},
			},
		},
		{ // "Traslados y Retenciones (2 conceptos)"
			TestName: "Traslados y Retenciones (2 conceptos)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			Expected: &Impuestos{
				TotalImpuestosRetenidos:   decimal.NewFromFloat(41.32),
				TotalImpuestosTrasladados: decimal.NewFromFloat(32),
				Retenciones: []*ImpuestosRetencion{
					{
						Impuesto: types.ImpuestoISR,
						Importe:  decimal.NewFromFloat(20),
					},
					{
						Impuesto: types.ImpuestoIVA,
						Importe:  decimal.NewFromFloat(21.32),
					},
				},
				Traslados: []*ImpuestosTraslado{
					{
						Base:       decimal.NewFromFloat(200),
						Impuesto:   types.ImpuestoIVA,
						Importe:    decimal.NewFromFloat(32),
						TipoFactor: types.TipoFactorTasa,
						TasaOCuota: decimal.NewFromFloat(0.16),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			actual := NewImpuestos(tt.Comprobante)
			diffs := compare.NewDiffs()
			compareEqualImpuestos(diffs, tt.Expected, actual, "")
			assert.NoError(t, diffs.Err())
		})
	}
}

func TestNewEmisor(t *testing.T) {
	tests := []struct {
		TestName         string
		Rfc              string
		Nombre           string
		RegimenFiscal    types.RegimenFiscal
		FacAtrAdquirente string
		Expected         *Emisor
		ExpectedErr      error
	}{
		{
			TestName:      "OK (Sin FacAtrAdquirente)",
			Rfc:           "AAA010101AAA",
			Nombre:        "Nombre Emisor",
			RegimenFiscal: types.RegimenFiscal601,
			Expected: &Emisor{
				Rfc:              "AAA010101AAA",
				Nombre:           "Nombre Emisor",
				RegimenFiscal:    types.RegimenFiscal601,
				FacAtrAdquirente: "",
			},
		},
		{
			TestName:         "OK (Con FacAtrAdquirente)",
			Rfc:              "AAA010101AAA",
			Nombre:           "Nombre Emisor",
			RegimenFiscal:    types.RegimenFiscal608,
			FacAtrAdquirente: "00001",
			Expected: &Emisor{
				Rfc:              "AAA010101AAA",
				Nombre:           "Nombre Emisor",
				RegimenFiscal:    types.RegimenFiscal608,
				FacAtrAdquirente: "00001",
			},
		},
		{
			TestName:      "Error (Rfc vacio)",
			Rfc:           "",
			Nombre:        "Nombre Emisor",
			RegimenFiscal: types.RegimenFiscal601,
			ExpectedErr:   ErrRequired,
		},
		{
			TestName:      "Error (Nombre vacio)",
			Rfc:           "AAA010101AAA",
			Nombre:        "",
			RegimenFiscal: types.RegimenFiscal601,
			ExpectedErr:   ErrRequired,
		},
		{
			TestName:      "Error (RegimenFiscal vacio)",
			Rfc:           "AAA010101AAA",
			Nombre:        "Nombre Emisor",
			RegimenFiscal: "",
			ExpectedErr:   ErrRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			actual, err := NewEmisor(tt.Rfc, tt.Nombre, tt.RegimenFiscal, tt.FacAtrAdquirente)
			if tt.ExpectedErr != nil {
				assert.ErrorIs(t, err, tt.ExpectedErr)
			} else {
				assert.NoError(t, err)
			}
			diffs := compare.NewDiffs()
			compareEqualEmisor(diffs, tt.Expected, actual, "")
			assert.NoError(t, diffs.Err())
		})
	}
}

func TestNewReceptorPublico(t *testing.T) {
	tests := []struct {
		TestName        string
		LugarExpedicion string
		Expected        *Receptor
		ExpectedErr     error
	}{
		{ // OK
			TestName:        "OK",
			LugarExpedicion: "00000",
			Expected: &Receptor{
				Rfc:                     RFCPublico,
				Nombre:                  "PUBLICO EN GENERAL",
				DomicilioFiscalReceptor: "00000",
				ResidenciaFiscal:        "",
				NumRegIdTrib:            "",
				RegimenFiscalReceptor:   types.RegimenFiscal616,
				UsoCFDI:                 types.UsoCFDIS01,
			},
		},
		{ // Error
			TestName:        "Error",
			LugarExpedicion: "",
			ExpectedErr:     ErrRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			actual, err := NewReceptorPublico(tt.LugarExpedicion)
			if tt.ExpectedErr != nil {
				assert.ErrorIs(t, err, tt.ExpectedErr)
			} else {
				assert.NoError(t, err)
			}
			diffs := compare.NewDiffs()
			compareEqualReceptor(diffs, tt.Expected, actual, "")
			assert.NoError(t, diffs.Err())
		})
	}
}

func TestNewReceptor(t *testing.T) {
	tests := []struct {
		TestName        string
		Rfc             string
		Nombre          string
		DomicilioFiscal string
		RegimenFiscal   types.RegimenFiscal
		UsoCFDI         types.UsoCFDI
		Expected        *Receptor
		ExpectedErr     error
	}{
		{ // "OK"
			TestName:        "OK",
			Rfc:             "AAA010101AAA",
			Nombre:          "Nombre Receptor",
			DomicilioFiscal: "00000",
			RegimenFiscal:   types.RegimenFiscal601,
			UsoCFDI:         types.UsoCFDIG03,
			Expected: &Receptor{
				Rfc:                     "AAA010101AAA",
				Nombre:                  "Nombre Receptor",
				DomicilioFiscalReceptor: "00000",
				ResidenciaFiscal:        "",
				NumRegIdTrib:            "",
				RegimenFiscalReceptor:   types.RegimenFiscal601,
				UsoCFDI:                 types.UsoCFDIG03,
			},
		},
		{ // "Error (Rfc vacio)"
			TestName:        "Error (Rfc vacio)",
			Rfc:             "",
			Nombre:          "Nombre Receptor",
			DomicilioFiscal: "00000",
			RegimenFiscal:   types.RegimenFiscal601,
			UsoCFDI:         types.UsoCFDIG03,
			ExpectedErr:     ErrRequired,
		},
		{ // "Error (Rfc invalido)"
			TestName:        "Error (Rfc invalido)",
			Rfc:             RFCPublico,
			Nombre:          "Nombre Receptor",
			DomicilioFiscal: "00000",
			RegimenFiscal:   types.RegimenFiscal601,
			UsoCFDI:         types.UsoCFDIG03,
			ExpectedErr:     ErrInvalid,
		},
		{ // "Error (Nombre vacio)"
			TestName:        "Error (Nombre vacio)",
			Rfc:             "AAA010101AAA",
			Nombre:          "",
			DomicilioFiscal: "00000",
			RegimenFiscal:   types.RegimenFiscal601,
			UsoCFDI:         types.UsoCFDIG03,
			ExpectedErr:     ErrRequired,
		},
		{ // "Error (DomicilioFiscal vacio)"
			TestName:        "Error (DomicilioFiscal vacio)",
			Rfc:             "AAA010101AAA",
			Nombre:          "Nombre Receptor",
			DomicilioFiscal: "",
			RegimenFiscal:   types.RegimenFiscal601,
			UsoCFDI:         types.UsoCFDIG03,
			ExpectedErr:     ErrRequired,
		},
		{ // "Error (RegimenFiscal vacio)"
			TestName:        "Error (RegimenFiscal vacio)",
			Rfc:             "AAA010101AAA",
			Nombre:          "Nombre Receptor",
			DomicilioFiscal: "00000",
			RegimenFiscal:   "",
			UsoCFDI:         types.UsoCFDIG03,
			ExpectedErr:     ErrRequired,
		},
		{ // "Error (UsoCFDI vacio)"
			TestName:        "Error (UsoCFDI vacio)",
			Rfc:             "AAA010101AAA",
			Nombre:          "Nombre Receptor",
			DomicilioFiscal: "00000",
			RegimenFiscal:   types.RegimenFiscal601,
			UsoCFDI:         "",
			ExpectedErr:     ErrRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			actual, err := NewReceptor(tt.Rfc, tt.Nombre, tt.DomicilioFiscal, tt.RegimenFiscal, tt.UsoCFDI)
			if tt.ExpectedErr != nil {
				assert.ErrorIs(t, err, tt.ExpectedErr)
			} else {
				assert.NoError(t, err)
			}
			diffs := compare.NewDiffs()
			compareEqualReceptor(diffs, tt.Expected, actual, "")
			assert.NoError(t, diffs.Err())
		})
	}
}

func TestNewReceptorExt(t *testing.T) {
	tests := []struct {
		TestName         string
		Nombre           string
		ResidenciaFiscal types.Pais
		NumRegIdTrib     string
		UsoCFDI          types.UsoCFDI
		LugarExpedicion  string
		Expected         *Receptor
		ExpectedErr      error
	}{
		{ // "OK"
			TestName:         "OK",
			Nombre:           "Nombre Receptor",
			ResidenciaFiscal: types.PaisUSA,
			NumRegIdTrib:     "1234567890123",
			UsoCFDI:          types.UsoCFDIG03,
			LugarExpedicion:  "00000",
			Expected: &Receptor{
				Rfc:                     RFCExtranjero,
				Nombre:                  "Nombre Receptor",
				DomicilioFiscalReceptor: "00000",
				ResidenciaFiscal:        types.PaisUSA,
				NumRegIdTrib:            "1234567890123",
				RegimenFiscalReceptor:   types.RegimenFiscal616,
				UsoCFDI:                 types.UsoCFDIG03,
			},
		},
		{ // "Error (Nombre vacio)"
			TestName:         "Error (Nombre vacio)",
			Nombre:           "",
			ResidenciaFiscal: types.PaisUSA,
			NumRegIdTrib:     "1234567890123",
			UsoCFDI:          types.UsoCFDIG03,
			LugarExpedicion:  "00000",
			ExpectedErr:      ErrRequired,
		},
		{ // "Error (ResidenciaFiscal vacio)"
			TestName:         "Error (ResidenciaFiscal vacio)",
			Nombre:           "Nombre Receptor",
			ResidenciaFiscal: "",
			NumRegIdTrib:     "1234567890123",
			UsoCFDI:          types.UsoCFDIG03,
			LugarExpedicion:  "00000",
			ExpectedErr:      ErrRequired,
		},
		{ // "Error (NumRegIdTrib vacio)"
			TestName:         "Error (NumRegIdTrib vacio)",
			Nombre:           "Nombre Receptor",
			ResidenciaFiscal: types.PaisUSA,
			NumRegIdTrib:     "",
			UsoCFDI:          types.UsoCFDIG03,
			LugarExpedicion:  "00000",
			ExpectedErr:      ErrRequired,
		},
		{ // "Error (UsoCFDI vacio)"
			TestName:         "Error (UsoCFDI vacio)",
			Nombre:           "Nombre Receptor",
			ResidenciaFiscal: types.PaisUSA,
			NumRegIdTrib:     "1234567890123",
			UsoCFDI:          "",
			LugarExpedicion:  "00000",
			ExpectedErr:      ErrRequired,
		},
		{ // "Error (LugarExpedicion vacio)"
			TestName:         "Error (LugarExpedicion vacio)",
			Nombre:           "Nombre Receptor",
			ResidenciaFiscal: types.PaisUSA,
			NumRegIdTrib:     "1234567890123",
			UsoCFDI:          types.UsoCFDIG03,
			LugarExpedicion:  "",
			ExpectedErr:      ErrRequired,
		},
		{ // "Error (ResidenciaFiscal invalido)"
			TestName:         "Error (ResidenciaFiscal invalido)",
			Nombre:           "Nombre Receptor",
			ResidenciaFiscal: types.PaisMEX,
			NumRegIdTrib:     "1234567890123",
			UsoCFDI:          types.UsoCFDIG03,
			LugarExpedicion:  "00000",
			ExpectedErr:      ErrInvalid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			actual, err := NewReceptorExt(tt.Nombre, tt.ResidenciaFiscal, tt.NumRegIdTrib, tt.UsoCFDI, tt.LugarExpedicion)
			if tt.ExpectedErr != nil {
				assert.ErrorIs(t, err, tt.ExpectedErr)
			} else {
				assert.NoError(t, err)
			}
			diffs := compare.NewDiffs()
			compareEqualReceptor(diffs, tt.Expected, actual, "")
			assert.NoError(t, diffs.Err())
		})
	}
}

func TestCalcularTotales(t *testing.T) {
	tests := []struct {
		TestName          string
		Comprobante       Comprobante
		ExpectedSubTotal  decimal.Decimal
		ExpectedDescuento decimal.Decimal
		ExpectedTotal     decimal.Decimal
	}{
		{ // "No Impuestos (1 concepto)"
			TestName: "No impuestos (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(100),
			ExpectedDescuento: decimal.NewFromFloat(0),
			ExpectedTotal:     decimal.NewFromFloat(100),
		},
		{ // "No Impuestos (1 concepto) descuento"
			TestName: "No impuestos (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe:   decimal.NewFromFloat(100),
						Descuento: decimal.NewFromFloat(10),
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(100),
			ExpectedDescuento: decimal.NewFromFloat(10),
			ExpectedTotal:     decimal.NewFromFloat(90),
		},
		{ // "Solo Retenciones (1 concepto)"
			TestName: "Solo Retenciones (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
						},
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(100),
			ExpectedDescuento: decimal.NewFromFloat(0),
			ExpectedTotal:     decimal.NewFromFloat(79.34),
		},
		{ // "Solo Traslados (1 concepto)"
			TestName: "Solo Traslados (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(100),
			ExpectedDescuento: decimal.NewFromFloat(0),
			ExpectedTotal:     decimal.NewFromFloat(116),
		},
		{ // "Traslados y Retenciones (1 concepto)"
			TestName: "Traslados y Retenciones (1 concepto)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(100),
			ExpectedDescuento: decimal.NewFromFloat(0),
			ExpectedTotal:     decimal.NewFromFloat(95.34),
		},
		{ // "Traslados y Retenciones (2 conceptos)"
			TestName: "Traslados y Retenciones (2 conceptos)",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
					{
						Importe: decimal.NewFromFloat(100),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(200),
			ExpectedDescuento: decimal.NewFromFloat(0),
			ExpectedTotal:     decimal.NewFromFloat(190.68),
		},
		{ // "Traslados y Retenciones (1 concepto) Descuento"
			TestName: "Traslados y Retenciones (1 concepto) Descuento",
			Comprobante: Comprobante{
				Conceptos: []*Concepto{
					{
						Importe:   decimal.NewFromFloat(110),
						Descuento: decimal.NewFromFloat(10),
						Impuestos: &ConceptoImpuestos{
							Retenciones: []*ConceptoImpuestosRetencion{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoISR,
									Importe:    decimal.NewFromFloat(10),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.10),
								},
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(10.66),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.106666),
								},
							},
							Traslados: []*ConceptoImpuestosTraslado{
								{
									Base:       decimal.NewFromFloat(100),
									Impuesto:   types.ImpuestoIVA,
									Importe:    decimal.NewFromFloat(16),
									TipoFactor: types.TipoFactorTasa,
									TasaOCuota: decimal.NewFromFloat(0.16),
								},
							},
						},
					},
				},
			},
			ExpectedSubTotal:  decimal.NewFromFloat(110),
			ExpectedDescuento: decimal.NewFromFloat(10),
			ExpectedTotal:     decimal.NewFromFloat(95.34),
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			subtotal, descuento, total := CalcularTotales(tt.Comprobante)
			assert.True(t, subtotal.Equal(tt.ExpectedSubTotal), "Subtotal: %s != %s", subtotal, tt.ExpectedSubTotal)
			assert.True(t, descuento.Equal(tt.ExpectedDescuento), "Descuento: %s != %s", descuento, tt.ExpectedDescuento)
			assert.True(t, total.Equal(tt.ExpectedTotal), "Total: %s != %s", total, tt.ExpectedTotal)
		})
	}
}

func TestNewInformacionGlobal(t *testing.T) {
	fecha := func(s string) types.FechaH {
		t, err := types.NewFechaH(s)
		if err != nil {
			panic(err)
		}
		return t
	}
	tests := []struct {
		TestName     string
		Rfc          string
		Periodicidad types.Periodicidad
		Fecha        types.FechaH
		Expected     *InformacionGlobal
		ExpectedErr  error
	}{
		{ // "Sin Periodicidad"
			TestName:     "Sin Periodicidad",
			Rfc:          "AAA010101AAA",
			Periodicidad: types.PeriodicidadMensual,
			Fecha:        fecha("2022-01-01T01:01:01"),
			Expected:     nil,
			ExpectedErr:  nil,
		},
		{ // "Bimestre 1 (Enero)"
			TestName:     "Bimestre 1 (Enero)",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadBimestral,
			Fecha:        fecha("2022-01-01T01:01:01"),
			Expected: &InformacionGlobal{
				Periodicidad: types.PeriodicidadBimestral,
				Meses:        "13",
				Anio:         2022,
			},
			ExpectedErr: nil,
		},
		{ // "Bimestre 1 (Febrero)"
			TestName:     "Bimestre 1 (Febrero)",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadBimestral,
			Fecha:        fecha("2022-02-01T01:01:01"),
			Expected: &InformacionGlobal{
				Periodicidad: types.PeriodicidadBimestral,
				Meses:        "13",
				Anio:         2022,
			},
			ExpectedErr: nil,
		},
		{ // "Bimestre 2 (Marzo)"
			TestName:     "Bimestre 2 (Marzo)",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadBimestral,
			Fecha:        fecha("2022-03-01T01:01:01"),
			Expected: &InformacionGlobal{
				Periodicidad: types.PeriodicidadBimestral,
				Meses:        "14",
				Anio:         2022,
			},
			ExpectedErr: nil,
		},
		{ // "Bimestre 6 (Noviembre)"
			TestName:     "Bimestre 6 (Noviembre)",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadBimestral,
			Fecha:        fecha("2022-11-01T01:01:01"),
			Expected: &InformacionGlobal{
				Periodicidad: types.PeriodicidadBimestral,
				Meses:        "18",
				Anio:         2022,
			},
			ExpectedErr: nil,
		},
		{ // "Bimestre 6 (Diciembre)"
			TestName:     "Bimestre 6 (Diciembre)",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadBimestral,
			Fecha:        fecha("2022-12-01T01:01:01"),
			Expected: &InformacionGlobal{
				Periodicidad: types.PeriodicidadBimestral,
				Meses:        "18",
				Anio:         2022,
			},
			ExpectedErr: nil,
		},
		{ // "Mes 05 (Mayo)"
			TestName:     "Mes 05 (Mayo)",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadMensual,
			Fecha:        fecha("2022-05-01T01:01:01"),
			Expected: &InformacionGlobal{
				Periodicidad: types.PeriodicidadMensual,
				Meses:        "05",
				Anio:         2022,
			},
			ExpectedErr: nil,
		},
		{ // "Required Periodicidad"
			TestName:     "Required Periodicidad",
			Rfc:          RFCPublico,
			Periodicidad: "",
			Fecha:        fecha("2022-01-01T01:01:01"),
			Expected:     nil,
			ExpectedErr:  ErrRequired,
		},
		{ // "Required Fecha"
			TestName:     "Required Fecha",
			Rfc:          RFCPublico,
			Periodicidad: types.PeriodicidadMensual,
			Fecha:        types.FechaH{},
			Expected:     nil,
			ExpectedErr:  ErrRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			actual, err := NewInformacionGlobal(tt.Rfc, tt.Periodicidad, tt.Fecha)
			if tt.ExpectedErr != nil {
				assert.ErrorIs(t, err, tt.ExpectedErr)
			} else {
				assert.NoError(t, err)
			}
			diffs := compare.NewDiffs()
			compareEqualInformacionGlobal(diffs, tt.Expected, actual, "")
			assert.NoError(t, diffs.Err())
		})
	}
}

func TestConceptoTotal(t *testing.T) {
	tests := []struct {
		TestName      string
		Concepto      *Concepto
		ExpectedTotal decimal.Decimal
	}{
		{ // "No impuestos"
			TestName: "No impuestos",
			Concepto: &Concepto{
				Importe: decimal.NewFromFloat(100),
			},
			ExpectedTotal: decimal.NewFromFloat(100),
		},
		{ // "No impuestos Descuento"
			TestName: "No impuestos Descuento",
			Concepto: &Concepto{
				Importe:   decimal.NewFromFloat(100),
				Descuento: decimal.NewFromFloat(10),
			},
			ExpectedTotal: decimal.NewFromFloat(90),
		},
		{ // "Solo Traslados"
			TestName: "Solo Traslados",
			Concepto: &Concepto{
				Importe: decimal.NewFromFloat(100),
				Impuestos: &ConceptoImpuestos{
					Traslados: []*ConceptoImpuestosTraslado{
						{
							Base:       decimal.NewFromFloat(100),
							Impuesto:   types.ImpuestoIVA,
							Importe:    decimal.NewFromFloat(16),
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.16),
						},
					},
				},
			},
			ExpectedTotal: decimal.NewFromFloat(116),
		},
		{ // "Solo Retenciones"
			TestName: "Solo Retenciones",
			Concepto: &Concepto{
				Importe: decimal.NewFromFloat(100),
				Impuestos: &ConceptoImpuestos{
					Retenciones: []*ConceptoImpuestosRetencion{
						{
							Base:       decimal.NewFromFloat(100),
							Impuesto:   types.ImpuestoISR,
							Importe:    decimal.NewFromFloat(10),
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.10),
						},
					},
				},
			},
			ExpectedTotal: decimal.NewFromFloat(90),
		},
		{ // "Traslados, Retenciones y Descuento"
			TestName: "Traslados, Retenciones y Descuento",
			Concepto: &Concepto{
				Importe:   decimal.NewFromFloat(110),
				Descuento: decimal.NewFromFloat(10),
				Impuestos: &ConceptoImpuestos{
					Retenciones: []*ConceptoImpuestosRetencion{
						{
							Base:       decimal.NewFromFloat(100),
							Impuesto:   types.ImpuestoISR,
							Importe:    decimal.NewFromFloat(10),
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.10),
						},
						{
							Base:       decimal.NewFromFloat(100),
							Impuesto:   types.ImpuestoIVA,
							Importe:    decimal.NewFromFloat(10.66),
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.106666),
						},
					},
					Traslados: []*ConceptoImpuestosTraslado{
						{
							Base:       decimal.NewFromFloat(100),
							Impuesto:   types.ImpuestoIVA,
							Importe:    decimal.NewFromFloat(16),
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.16),
						},
					},
				},
			},
			ExpectedTotal: decimal.NewFromFloat(95.34),
		},
	}
	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			total := tt.Concepto.Total()
			assert.True(t, total.Equal(tt.ExpectedTotal), "Total: %s != %s", total, tt.ExpectedTotal)
		})
	}
}
