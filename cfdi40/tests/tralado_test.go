package tests_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/complemento/cartaporte20"
	"github.com/veyronifs/cfdi-go/types"
)

func TestTrasladoSimple(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteT,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
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
				ObjetoImp:        types.ObjetoImp01,
				Cantidad:         decimal.NewFromFloat(10),
				ClaveProdServ:    "50192100",
				ClaveUnidad:      "XBX",
				Descripcion:      "Cacahuate",
				Importe:          decimal.NewFromFloat(0),
				NoIdentificacion: "1",
				ValorUnitario:    decimal.NewFromFloat(0),
			},
		},
	}

	testTimbrar(t, c)
}

func TestTrasladoCartaPorte20Autotransporte(t *testing.T) {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaXXX,
		TipoDeComprobante: types.ComprobanteT,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "14210",
		Serie:             "Serie",
		Folio:             "Folio",
		Fecha:             newFechaHNow2(),
		Emisor:            emisor16_8_0,
		Receptor: &cfdi40.Receptor{
			Rfc:                     "KAHO641101B39",
			Nombre:                  "OSCAR KALA HAAK",
			DomicilioFiscalReceptor: "29950",
			RegimenFiscalReceptor:   "612",
			UsoCFDI:                 types.UsoCFDIS01,
		},
		Conceptos: cfdi40.Conceptos{
			{
				ObjetoImp:        types.ObjetoImp01,
				Cantidad:         decimal.NewFromFloat(10),
				ClaveProdServ:    "50425400",
				ClaveUnidad:      "KGM",
				Descripcion:      "Cacahuate",
				Importe:          decimal.NewFromFloat(0),
				NoIdentificacion: "1",
				ValorUnitario:    decimal.NewFromFloat(0),
			},
		},
		Complemento: &cfdi40.Complemento{
			CartaPorte20: &cartaporte20.CartaPorte20{
				Version:        "2.0",
				TranspInternac: "No",
				TotalDistRec:   decimal.NewFromFloat(10),
				Ubicaciones: cartaporte20.Ubicaciones{
					{
						TipoUbicacion:               "Origen",
						RFCRemitenteDestinatario:    emisor16_8_0.Rfc,
						NombreRemitenteDestinatario: emisor16_8_0.Nombre,
						FechaHoraSalidaLlegada:      types.NewFechaHNow(),
						Domicilio: &cartaporte20.Domicilio{
							Calle:          "Filomeno Mata 922",
							NumeroExterior: "922",
							Referencia:     "Filomeno Mata 922, Ejidos de Santa María Aztahuacán, DIF",
							Estado:         "DIF",
							Pais:           "MEX",
							CodigoPostal:   "09570",
						},
					},
					{
						TipoUbicacion:               "Destino",
						RFCRemitenteDestinatario:    "BAR011108CC6",
						NombreRemitenteDestinatario: "BARCEL",
						FechaHoraSalidaLlegada:      types.NewFechaHTime(time.Now().Add(time.Hour * 2)),
						DistanciaRecorrida:          decimal.NewFromFloat(10),
						Domicilio: &cartaporte20.Domicilio{
							Calle:          "KM 54 Carretera Mexico-Toluca",
							NumeroExterior: "SN",
							Referencia:     "KM 54 Carretera Mexico-Toluca, Rancho El Jazmin Lerma, Mexico, MEX",
							Estado:         "MEX",
							Pais:           "MEX",
							CodigoPostal:   "52000",
						},
					},
				},
				Mercancias: &cartaporte20.Mercancias{
					PesoBrutoTotal:     decimal.NewFromFloat(10),
					UnidadPeso:         "KGM",
					NumTotalMercancias: 1,
					Mercancia: []*cartaporte20.Mercancia{
						{
							BienesTransp: "50425400",
							Descripcion:  "Cacahuate",
							Cantidad:     decimal.NewFromFloat(10),
							ClaveUnidad:  "KGM",
							PesoEnKg:     decimal.NewFromFloat(10),
						},
					},
					Autotransporte: &cartaporte20.Autotransporte{
						PermSCT:       "TPAF01",
						NumPermisoSCT: "1234567890123456789012345678901234567890",
						IdentificacionVehicular: &cartaporte20.IdentificacionVehicular{
							ConfigVehicular: "VL",
							PlacaVM:         "plac892",
							AnioModeloVM:    "2020",
						},
						Seguros: &cartaporte20.Seguros{
							AseguraCarga:     "SW Seguros",
							AseguraRespCivil: "SW Seguros",
							PolizaRespCivil:  "123456789",
						},
						Remolques: cartaporte20.Remolques{
							{
								SubTipoRem: "CTR004",
								Placa:      "VL45K98",
							},
						},
					},
				},
				FiguraTransporte: &cartaporte20.FiguraTransporte{
					TiposFigura: []*cartaporte20.TiposFigura{
						{
							TipoFigura:  "01",
							RFCFigura:   "VAAM130719H60",
							NumLicencia: "1234567890",
							Domicilio: &cartaporte20.Domicilio{
								Calle:          "calle",
								NumeroExterior: "211",
								Colonia:        "0347",
								Localidad:      "23",
								Referencia:     "casa blanca 1",
								Municipio:      "004",
								Estado:         "COA",
								Pais:           "MEX",
								CodigoPostal:   "25350",
							},
						},
					},
				},
			},
		},
	}

	testTimbrar(t, c)
}
