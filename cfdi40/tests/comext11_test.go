package tests_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/complemento/comext11"
	"github.com/veyronifs/cfdi-go/types"
)

func TestComext11TipoI(t *testing.T) {
	comext := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Exportacion:       types.Exportacion02,
		Serie:             "B",
		Folio:             "2921",
		Fecha:             types.NewFechaHNow(),
		FormaPago:         "99",
		CondicionesDePago: "A 45 dias",
		SubTotal:          decimal.NewFromFloat(3000.00),
		Descuento:         decimal.NewFromFloat(0.00),
		Moneda:            types.MonedaUSD,
		TipoCambio:        decimal.NewFromFloat(20.300200),
		Total:             decimal.NewFromFloat(3000.00),
		TipoDeComprobante: types.ComprobanteI,
		MetodoPago:        types.MetodoPagoPPD,
		LugarExpedicion:   "14210",
		Emisor: &cfdi40.Emisor{
			Rfc:           "IXS7607092R5",
			Nombre:        "INTERNACIONAL XIMBO Y SABORES",
			RegimenFiscal: "601",
		},
		Receptor: &cfdi40.Receptor{
			Rfc:                     "XEXX010101000",
			Nombre:                  "GALDISA, USA INC. TAX ID: 36-4812096",
			DomicilioFiscalReceptor: "14210",
			RegimenFiscalReceptor:   "616",
			ResidenciaFiscal:        "USA",
			NumRegIdTrib:            "364812096",
			UsoCFDI:                 types.UsoCFDIS01,
		},
		Conceptos: cfdi40.Conceptos{
			{
				ClaveProdServ:    "50405500",
				NoIdentificacion: "1A2RNHG2A",
				ObjetoImp:        types.ObjetoImp02,
				Cantidad:         decimal.NewFromFloat(1.00),
				ClaveUnidad:      types.UnidadKilogramo,
				ValorUnitario:    decimal.NewFromFloat(3000.00),
				Unidad:           "kg",
				Descripcion:      "CACAHUATE RUNNER HIGH OLEIC 38/42",
				Importe:          decimal.NewFromFloat(3000.00),
				Descuento:        decimal.NewFromFloat(0.00),
				Impuestos: &cfdi40.ConceptoImpuestos{
					Traslados: cfdi40.ConceptoImpuestosTraslados{
						{
							Importe:    decimal.NewFromFloat(0.00),
							TasaOCuota: decimal.NewFromFloat(0.000000),
							TipoFactor: types.TipoFactorTasa,
							Impuesto:   types.ImpuestoIVA,
							Base:       decimal.NewFromFloat(3000.00),
						},
					},
				},
			},
		},
		Impuestos: &cfdi40.Impuestos{
			TotalImpuestosTrasladados: decimal.NewFromFloat(0.00),
			Traslados: cfdi40.ImpuestosTraslados{
				{
					TasaOCuota: decimal.NewFromFloat(0.000000),
					Importe:    decimal.NewFromFloat(0.000000),
					TipoFactor: types.TipoFactorTasa,
					Impuesto:   types.ImpuestoIVA,
					Base:       decimal.NewFromFloat(3000.00),
				},
			},
		},
		Complemento: &cfdi40.Complemento{
			CCE11: &comext11.ComercioExterior{
				Version:           "1.1",
				TotalUSD:          decimal.NewFromFloat(3000.00),
				TipoCambioUSD:     decimal.NewFromFloat(20.300200),
				Subdivision:       0,
				Incoterm:          "FOB",
				CertificadoOrigen: 0,
				ClaveDePedimento:  "A1",
				TipoOperacion:     "2",
				Emisor: &comext11.Emisor{
					Domicilio: &comext11.Domicilio{
						Calle:          "Pico de Verapaz",
						NumeroExterior: "449-A",
						Colonia:        "2085",
						Municipio:      "012",
						Estado:         "DIF",
						Pais:           "MEX",
						CodigoPostal:   "14210",
					},
				},
				Receptor: &comext11.Receptor{
					NumRegIdTrib: "364812096",
					Domicilio: &comext11.Domicilio{
						Calle:          "3455 POLLOK DR",
						NumeroExterior: "3455",
						Colonia:        ".",
						Municipio:      ".",
						Estado:         "TX",
						Pais:           "USA",
						CodigoPostal:   "14210",
						Localidad:      "LAREDO, TX",
					},
				},
				Mercancias: comext11.Mercancias{
					{
						NoIdentificacion:    "1A2RNHG2A",
						ValorDolares:        decimal.NewFromFloat(3000.00),
						CantidadAduana:      decimal.NewFromFloat(1.0),
						ValorUnitarioAduana: decimal.NewFromFloat(3000.0),
						UnidadAduana:        "01",
						FraccionArancelaria: "1202420100",
					},
				},
			},
		},
	}
	testTimbrar(t, comext)
}
