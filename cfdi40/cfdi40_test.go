package cfdi40_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/complemento/cartaporte20"
	"github.com/veyronifs/cfdi-go/complemento/pagos20"
	"github.com/veyronifs/cfdi-go/complemento/tfd11"
	"github.com/veyronifs/cfdi-go/types"
)

func TestParseSimplePrecfdi(t *testing.T) {
	var xmlOriginal []byte
	{
		// A valid XML document for CFD v4.0 schema with random values for testing
		xmlOriginal = []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
		<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xsi:schemaLocation="http://www.sat.gob.mx/cfd/4 cfdv40.xsd"
			Version="4.0" LugarExpedicion="99999" MetodoPago="PPD" Confirmacion="A1234" Moneda="MXN"
			Descuento="0.00" Folio="123ABC" TipoCambio="1.0" Serie="A" Exportacion="03"
			TipoDeComprobante="P" FormaPago="99" CondicionesDePago="CONDICIONES"
			Fecha="2021-12-07T23:59:59" SubTotal="1000" Total="1500"
			NoCertificado="30001000000300023708"  Certificado="ABC"  Sello="XYZ">
			<cfdi:InformacionGlobal Meses="18" Año="2021" Periodicidad="05" />
			<cfdi:CfdiRelacionados TipoRelacion="09">
				<cfdi:CfdiRelacionado UUID="ED1752FE-E865-4FF2-BFE1-0F552E770DC9" />
			</cfdi:CfdiRelacionados>
			<cfdi:Emisor FacAtrAdquirente="0123456789" Nombre="Esta es una demostración" RegimenFiscal="630" Rfc="AAA010101AAA" />
			<cfdi:Receptor ResidenciaFiscal="MEX" DomicilioFiscalReceptor="99999" RegimenFiscalReceptor="630" Nombre="Juanito Bananas De la Sierra" NumRegIdTrib="0000000000000" Rfc="BASJ600902KL9" UsoCFDI="S01" />
			<cfdi:Conceptos>
				<cfdi:Concepto  ObjetoImp="01" ClaveProdServ="01010101" ClaveUnidad="C81" NoIdentificacion="00001" Cantidad="1.5"
				Unidad="TONELADA" Descripcion="ACERO" ValorUnitario="1500000" Importe="2250000">
					<cfdi:Impuestos>
						<cfdi:Traslados>
							<cfdi:Traslado Base="2250000" Impuesto="002" TipoFactor="Tasa" TasaOCuota="1.600000" Importe="360000"/>
						</cfdi:Traslados>
						<cfdi:Retenciones>
							<cfdi:Retencion Base="2250000" Impuesto="001" TipoFactor="Tasa" TasaOCuota="0.300000" Importe="247500"/>
						</cfdi:Retenciones>
					</cfdi:Impuestos>
					<cfdi:CuentaPredial Numero="51888"/>
				</cfdi:Concepto>
			</cfdi:Conceptos>
			<cfdi:Impuestos TotalImpuestosRetenidos="247500" TotalImpuestosTrasladados="360000">
				<cfdi:Retenciones>
					<cfdi:Retencion Impuesto="001" Importe="247000"/>
					<cfdi:Retencion Impuesto="003" Importe="500"/>
				</cfdi:Retenciones>
				<cfdi:Traslados>
					<cfdi:Traslado Base="1.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="1.600000" Importe="360000"/>
				</cfdi:Traslados>
			</cfdi:Impuestos>
			<cfdi:Complemento></cfdi:Complemento>
		</cfdi:Comprobante>`)
	}
	cfdiUnmarshalled1, err := cfdi40.Unmarshal(xmlOriginal)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlOriginal): %s", err.Error())
		return
	}
	xmlMarshalled, err := cfdi40.Marshal(cfdiUnmarshalled1)
	if err != nil {
		t.Errorf("Error Marshal(*cfdiUnmarshalled1): %s", err.Error())
		return
	}

	cfdiUnmarshalled2, err := cfdi40.Unmarshal(xmlMarshalled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshalled): %s", err.Error())
		return
	}

	err = cfdi40.CompareEqual(cfdiUnmarshalled1, cfdiUnmarshalled2)
	assert.NoError(t, err)
}

func TestMarshal(t *testing.T) {
	cFecha, _ := types.NewFechaH("2021-12-07T23:59:59")
	cfdiOriginal := cfdi40.Comprobante{
		Version:           "4.0",
		LugarExpedicion:   "99999",
		MetodoPago:        "PPD",
		Confirmacion:      "A1234",
		Moneda:            "MXN",
		Descuento:         decimal.NewFromFloat(0.00),
		Folio:             "123ABC",
		TipoCambio:        decimal.NewFromFloat(1.0),
		Serie:             "A",
		Exportacion:       "03",
		TipoDeComprobante: "P",
		FormaPago:         "99",
		CondicionesDePago: "CONDICIONES",
		Fecha:             cFecha,
		SubTotal:          decimal.NewFromFloat(1000.0),
		Total:             decimal.NewFromFloat(1500.0),
		NoCertificado:     "30001000000300023708",
		Certificado:       "ABC",
		Sello:             "XYZ",
		InformacionGlobal: &cfdi40.InformacionGlobal{
			Meses:        "18",
			Anio:         2021,
			Periodicidad: "05",
		},
		CfdiRelacionados: []*cfdi40.CfdiRelacionados{
			{
				TipoRelacion: "09",
				CfdiRelacionado: []*cfdi40.CfdiRelacionado{
					{UUID: "ED1752FE-E865-4FF2-BFE1-0F552E770DC9"},
				},
			},
		},
		Emisor: &cfdi40.Emisor{
			FacAtrAdquirente: "0123456789",
			Nombre:           "Esta es una demostración",
			RegimenFiscal:    "630",
			Rfc:              "AAA010101AAA",
		},
		Receptor: &cfdi40.Receptor{
			ResidenciaFiscal:        "MEX",
			DomicilioFiscalReceptor: "99999",
			RegimenFiscalReceptor:   "630",
			Nombre:                  "Juanito Bananas De la Sierra",
			NumRegIdTrib:            "0000000000000",
			Rfc:                     "BASJ600902KL9",
			UsoCFDI:                 "S01",
		},
		Conceptos: cfdi40.Conceptos{
			{
				ObjetoImp:        "01",
				ClaveProdServ:    "01010101",
				ClaveUnidad:      "C81",
				NoIdentificacion: "00001",
				Cantidad:         decimal.NewFromFloat(1.5),
				Unidad:           "TONELADA",
				Descripcion:      "ACERO",
				ValorUnitario:    decimal.NewFromFloat(1500000),
				Importe:          decimal.NewFromFloat(2250000),
				Impuestos: &cfdi40.ConceptoImpuestos{
					Traslados: []*cfdi40.ConceptoImpuestosTraslado{
						{
							Base:       decimal.NewFromFloat(2250000),
							Impuesto:   "002",
							TipoFactor: "Tasa",
							TasaOCuota: decimal.NewFromFloat(1.600000),
							Importe:    decimal.NewFromFloat(360000),
						},
					},
					Retenciones: []*cfdi40.ConceptoImpuestosRetencion{
						{
							Base:       decimal.NewFromFloat(2250000),
							Impuesto:   "001",
							TipoFactor: "Tasa",
							TasaOCuota: decimal.NewFromFloat(0.300000),
							Importe:    decimal.NewFromFloat(247500),
						},
					},
				},
			},
		},

		Impuestos: &cfdi40.Impuestos{
			TotalImpuestosRetenidos:   decimal.NewFromFloat(247500),
			TotalImpuestosTrasladados: decimal.NewFromFloat(360000),
			Retenciones: cfdi40.ImpuestosRetenciones{
				{Impuesto: "001", Importe: decimal.NewFromFloat(247000)},
				{Impuesto: "003", Importe: decimal.NewFromFloat(500)},
			},
			Traslados: cfdi40.ImpuestosTraslados{
				{
					Base:       decimal.NewFromFloat(1.00),
					Impuesto:   "002",
					TipoFactor: "Tasa",
					TasaOCuota: decimal.NewFromFloat(1.600000),
					Importe:    decimal.NewFromFloat(360000),
				},
			},
		},
	}

	// marshal expected
	xmlMarshalled, err := cfdi40.Marshal(&cfdiOriginal)
	if err != nil {
		t.Errorf("Error Marshal(cfdiOriginal): %s", err.Error())
		return
	}

	cfdiUnmarshalled, err := cfdi40.Unmarshal(xmlMarshalled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshalled): %s", err.Error())
		return
	}

	err = cfdi40.CompareEqual(&cfdiOriginal, cfdiUnmarshalled)
	assert.NoError(t, err)
}

func TestCartaPorte(t *testing.T) {
	cFecha, _ := types.NewFechaH("2021-12-07T23:59:59")
	cfdiOriginal := cfdi40.Comprobante{
		Version:           "4.0",
		LugarExpedicion:   "99999",
		MetodoPago:        "PPD",
		Confirmacion:      "A1234",
		Moneda:            "MXN",
		Descuento:         decimal.NewFromFloat(0.00),
		Folio:             "123ABC",
		TipoCambio:        decimal.NewFromFloat(1.0),
		Serie:             "A",
		Exportacion:       "03",
		TipoDeComprobante: "P",
		FormaPago:         "99",
		CondicionesDePago: "CONDICIONES",
		Fecha:             cFecha,
		SubTotal:          decimal.NewFromFloat(1000.0),
		Total:             decimal.NewFromFloat(1500.0),
		NoCertificado:     "30001000000300023708",
		Certificado:       "ABC",
		Sello:             "XYZ",
		InformacionGlobal: &cfdi40.InformacionGlobal{
			Meses:        "18",
			Anio:         2021,
			Periodicidad: "05",
		},
		CfdiRelacionados: []*cfdi40.CfdiRelacionados{
			{
				TipoRelacion: "09",
				CfdiRelacionado: []*cfdi40.CfdiRelacionado{
					{UUID: "ED1752FE-E865-4FF2-BFE1-0F552E770DC9"},
				},
			},
		},
		Emisor: &cfdi40.Emisor{
			FacAtrAdquirente: "0123456789",
			Nombre:           "Esta es una demostración",
			RegimenFiscal:    "630",
			Rfc:              "AAA010101AAA",
		},
		Receptor: &cfdi40.Receptor{
			ResidenciaFiscal:        "MEX",
			DomicilioFiscalReceptor: "99999",
			RegimenFiscalReceptor:   "630",
			Nombre:                  "Juanito Bananas De la Sierra",
			NumRegIdTrib:            "0000000000000",
			Rfc:                     "BASJ600902KL9",
			UsoCFDI:                 "S01",
		},
		Conceptos: cfdi40.Conceptos{
			{
				ObjetoImp:        "01",
				ClaveProdServ:    "01010101",
				ClaveUnidad:      "C81",
				NoIdentificacion: "00001",
				Cantidad:         decimal.NewFromFloat(1.5),
				Unidad:           "TONELADA",
				Descripcion:      "ACERO",
				ValorUnitario:    decimal.NewFromFloat(1500000),
				Importe:          decimal.NewFromFloat(2250000),
				Impuestos: &cfdi40.ConceptoImpuestos{
					Traslados: []*cfdi40.ConceptoImpuestosTraslado{
						{
							Base:       decimal.NewFromFloat(2250000),
							Impuesto:   "002",
							TipoFactor: "Tasa",
							TasaOCuota: decimal.NewFromFloat(1.600000),
							Importe:    decimal.NewFromFloat(360000),
						},
					},
					Retenciones: []*cfdi40.ConceptoImpuestosRetencion{
						{
							Base:       decimal.NewFromFloat(2250000),
							Impuesto:   "001",
							TipoFactor: "Tasa",
							TasaOCuota: decimal.NewFromFloat(0.300000),
							Importe:    decimal.NewFromFloat(247500),
						},
					},
				},
			},
		},

		Impuestos: &cfdi40.Impuestos{
			TotalImpuestosRetenidos:   decimal.NewFromFloat(247500),
			TotalImpuestosTrasladados: decimal.NewFromFloat(360000),
			Retenciones: cfdi40.ImpuestosRetenciones{
				{Impuesto: "001", Importe: decimal.NewFromFloat(247000)},
				{Impuesto: "003", Importe: decimal.NewFromFloat(500)},
			},
			Traslados: cfdi40.ImpuestosTraslados{
				{
					Base:       decimal.NewFromFloat(1.00),
					Impuesto:   "002",
					TipoFactor: "Tasa",
					TasaOCuota: decimal.NewFromFloat(1.600000),
					Importe:    decimal.NewFromFloat(360000),
				},
			},
		},

		Complemento: &cfdi40.Complemento{
			TFD11: &tfd11.TimbreFiscalDigital{
				Version:          "1.1",
				UUID:             "ED1752FE-E865-4FF2-BFE1-0F552E770DC9",
				FechaTimbrado:    cFecha,
				SelloCFD:         "ABC",
				NoCertificadoSAT: "30001000000300023708",
				SelloSAT:         "XYZ",
				RfcProvCertif:    "AAA010101AAA",
				Leyenda:          "CFDI de prueba",
			},
			CartaPorte20: &cartaporte20.CartaPorte20{
				//Mercancias:0xc00010e100
				Version:           "2.0",
				TranspInternac:    "No",
				EntradaSalidaMerc: "Entrada",
				PaisOrigenDestino: "USA",
				ViaEntradaSalida:  "01",
				TotalDistRec:      decimal.New(1963, 1),
				Ubicaciones: []*cartaporte20.Ubicacion{
					{
						TipoUbicacion:               "Origen",
						IDUbicacion:                 "00001",
						RFCRemitenteDestinatario:    "IXS7607092R5",
						NombreRemitenteDestinatario: "PEPE",
						NumRegIdTrib:                "123456789",
						ResidenciaFiscal:            "MEX",
						NumEstacion:                 "PM001",
						NombreEstacion:              "ESTACION PM001",
						NavegacionTrafico:           "Cabotaje",
						FechaHoraSalidaLlegada:      cFecha,
						TipoEstacion:                "PM",
						DistanciaRecorrida:          decimal.New(1963, 1),
						Domicilio: &cartaporte20.Domicilio{
							Calle:          "Calle",
							NumeroExterior: "123",
							NumeroInterior: "456",
							Colonia:        "Colonia",
							Localidad:      "Localidad",
							Referencia:     "Referencia",
							Municipio:      "Municipio",
							Estado:         "Estado",
							Pais:           "MEX",
							CodigoPostal:   "12345",
						},
					},
				},
				Mercancias: &cartaporte20.Mercancias{
					UnidadPeso:         "KGM",
					NumTotalMercancias: 1,
					PesoBrutoTotal:     decimal.NewFromFloat(199.99),
					PesoNetoTotal:      decimal.NewFromFloat(199.99),
					CargoPorTasacion:   decimal.NewFromFloat(199.99),
					Mercancia: []*cartaporte20.Mercancia{
						{
							BienesTransp:         "10101501",
							ClaveSTCC:            "10101501",
							Descripcion:          "Caja de 20 piezas",
							Cantidad:             decimal.New(10, 1),
							ClaveUnidad:          "CJ",
							Unidad:               "Caja",
							Dimensiones:          "20x20x20",
							MaterialPeligroso:    "Sí",
							CveMaterialPeligroso: "M0001",
							Embalaje:             "1A2",
							DescripEmbalaje:      "Caja de 20 piezas",
							PesoEnKg:             decimal.NewFromFloat(199.99),
							ValorMercancia:       decimal.NewFromFloat(199.99),
							Moneda:               "MXN",
							FraccionArancelaria:  "10101501",
							UUIDComercioExt:      "6713E766-DCA2-41AA-B1C0-020CFB60AC95",
							Pedimentos: []*cartaporte20.Pedimentos{
								{Pedimento: "123456789"},
							},
							GuiasIdentificacion: []*cartaporte20.GuiasIdentificacion{
								{
									NumeroGuiaIdentificacion:  "123456789",
									DescripGuiaIdentificacion: "Guia de identificacion",
									PesoGuiaIdentificacion:    decimal.NewFromFloat(199.99),
								},
							},
							DetalleMercancia: &cartaporte20.DetalleMercancia{
								UnidadPesoMerc: "KGM",
								PesoBruto:      decimal.NewFromFloat(199.99),
								PesoNeto:       decimal.NewFromFloat(199.99),
								PesoTara:       decimal.NewFromFloat(199.99),
								NumPiezas:      1,
							},
							CantidadTransporta: []*cartaporte20.CantidadTransporta{
								{
									Cantidad:       decimal.New(1, 1),
									IDOrigen:       "00001",
									IDDestino:      "00002",
									CvesTransporte: "1A2",
								},
							},
						},
					},
					Autotransporte: &cartaporte20.Autotransporte{
						PermSCT:       "TPAF02",
						NumPermisoSCT: "09381581/080002",
						IdentificacionVehicular: &cartaporte20.IdentificacionVehicular{
							ConfigVehicular: "T3S2",
							PlacaVM:         "LE57937",
							AnioModeloVM:    "2008",
						},
						Seguros: &cartaporte20.Seguros{
							AseguraRespCivil:   "Seguros Atlas",
							PolizaRespCivil:    "MS1-1-7-38294",
							AseguraMedAmbiente: "Seguros Atlas",
							PolizaMedAmbiente:  "MS1-1-7-38294",
							AseguraCarga:       "Seguros Atlas",
							PolizaCarga:        "MS1-2-2-1353",
							PrimaSeguro:        decimal.New(10, 1),
						},
						Remolques: []*cartaporte20.Remolque{
							{SubTipoRem: "CTR004", Placa: "6HU3452"},
							{SubTipoRem: "CTR001", Placa: "6HU3112"},
						},
					},
				},
				FiguraTransporte: &cartaporte20.FiguraTransporte{
					TiposFigura: []*cartaporte20.TiposFigura{
						{
							TipoFigura:             "01",
							RFCFigura:              "VAEE770919BE9",
							NumLicencia:            "N03626578",
							NombreFigura:           "MANUEL ENRIQUE VELAZQUEZ ESPINOSA",
							NumRegIdTribFigura:     "0132456789",
							ResidenciaFiscalFigura: types.Pais("MEX"),
							PartesTransporte: []*cartaporte20.PartesTransporte{
								{ParteTransporte: "ABCD"},
							},
							Domicilio: &cartaporte20.Domicilio{
								Calle:          "Calle",
								NumeroExterior: "123",
								NumeroInterior: "456",
								Colonia:        "Colonia",
								Localidad:      "Localidad",
								Referencia:     "Referencia",
								Municipio:      "Municipio",
								Estado:         "Estado",
								Pais:           "MEX",
								CodigoPostal:   "12345",
							},
						},
					},
				},
			},
		},
	}

	// marshal expected
	xmlMarshalled, err := cfdi40.Marshal(&cfdiOriginal)
	if err != nil {
		t.Errorf("Error Marshal(cfdiOriginal): %s", err.Error())
		return
	}

	cfdiUnmarshalled, err := cfdi40.Unmarshal(xmlMarshalled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshalled): %s", err.Error())
		return
	}
	err = cfdi40.CompareEqual(&cfdiOriginal, cfdiUnmarshalled)
	assert.NoError(t, err)
}

func TestComplementoPagos(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`
			<?xml version="1.0" encoding="utf-8"?>
			<cfdi:Comprobante xsi:schemaLocation="http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd http://www.sat.gob.mx/Pagos20
			http://www.sat.gob.mx/sitio_internet/cfd/Pagos/Pagos20.xsd" xmlns:pago20="http://www.sat.gob.mx/Pagos20" Version="4.0" Serie="Serie" Folio="Folio" Fecha="2021-12-16T00:18:10" Sello="e" SubTotal="0" Moneda="XXX" Total="0.00" TipoDeComprobante="P" Exportacion="01" LugarExpedicion="20000"  xmlns:cfdi="http://www.sat.gob.mx/cfd/4" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
			<cfdi:Emisor Rfc="EQA9003177C9" Nombre="Emisor de prueba" RegimenFiscal="601" />
			<cfdi:Receptor Rfc="UNI180529TM6" Nombre="UNIVERSIDAD SA DE CV" DomicilioFiscalReceptor="65000" RegimenFiscalReceptor="601" UsoCFDI="G01" />
			<cfdi:Conceptos>
				<cfdi:Concepto ClaveProdServ="84111506" Cantidad="1" ClaveUnidad="ACT" Descripcion="Pago" ValorUnitario="0.00" Importe="0.00" ObjetoImp="01" />
			</cfdi:Conceptos>
			<cfdi:Complemento>
					<pago20:Pagos Version="2.0">
						<pago20:Totales MontoTotalPagos="200.00" />
						<pago20:Pago FechaPago="2021-12-15T00:00:00" FormaDePagoP="01" MonedaP="MXN" Monto="200.00" TipoCambioP="1">
							<pago20:DoctoRelacionado IdDocumento="bfc36522-4b8e-45c4-8f14-d11b289f9eb7" MonedaDR="MXN" NumParcialidad="1" ImpSaldoAnt="200.00" ImpPagado="200.00" ImpSaldoInsoluto="0.00" ObjetoImpDR="01" EquivalenciaDR="1"/>
						</pago20:Pago>
					</pago20:Pagos>
				</cfdi:Complemento>
			</cfdi:Comprobante>
		`)
	}
	expected := cfdi40.NewComprobantePago()
	expected.LugarExpedicion = "20000"
	expected.Serie = "Serie"
	expected.Folio = "Folio"
	expected.Fecha, _ = types.NewFechaH("2021-12-16T00:18:10")
	expected.Sello = "e"
	expected.Emisor = &cfdi40.Emisor{
		Rfc:           "EQA9003177C9",
		Nombre:        "Emisor de prueba",
		RegimenFiscal: "601",
	}
	expected.Receptor = &cfdi40.Receptor{
		Rfc:                     "UNI180529TM6",
		Nombre:                  "UNIVERSIDAD SA DE CV",
		DomicilioFiscalReceptor: "65000",
		RegimenFiscalReceptor:   "601",
		UsoCFDI:                 "G01",
	}
	expected.Conceptos = []*cfdi40.Concepto{
		{
			ClaveProdServ: "84111506",
			Cantidad:      decimal.NewFromInt(1),
			ClaveUnidad:   "ACT",
			Descripcion:   "Pago",
			ValorUnitario: decimal.NewFromFloat(0),
			Importe:       decimal.NewFromFloat(0),
			ObjetoImp:     types.ObjetoImp01,
		},
	}
	fechaPago, _ := types.NewFechaH("2021-12-15T00:00:00")
	expected.Complemento = &cfdi40.Complemento{
		Pagos20: &pagos20.Pagos{
			Version: "2.0",
			Totales: &pagos20.Totales{
				MontoTotalPagos: decimal.NewFromFloat(200),
			},
			Pago: []*pagos20.Pago{
				{
					FechaPago:    fechaPago,
					FormaDePagoP: types.FormaPago01,
					MonedaP:      types.MonedaMXN,
					Monto:        decimal.NewFromFloat(200),
					TipoCambioP:  decimal.NewFromFloat(1),
					DoctoRelacionado: []*pagos20.DoctoRelacionado{
						{
							IdDocumento:      "bfc36522-4b8e-45c4-8f14-d11b289f9eb7",
							MonedaDR:         types.MonedaMXN,
							NumParcialidad:   1,
							ImpSaldoAnt:      decimal.NewFromFloat(200),
							ImpPagado:        decimal.NewFromFloat(200),
							ImpSaldoInsoluto: decimal.NewFromFloat(0),
							ObjetoImpDR:      types.ObjetoImp01,
							EquivalenciaDR:   decimal.NewFromFloat(1),
						},
					},
				},
			},
		},
	}
	cfdiUnmarshaled, err := cfdi40.Unmarshal(xmlOriginal)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlOriginal): %s", err.Error())
		return
	}

	err = cfdi40.CompareEqual(expected, cfdiUnmarshaled)
	assert.NoError(t, err)
}

func TestAcentos(t *testing.T) {
	comprobanteOriginal := cfdi40.NewComprobantePago()
	comprobanteOriginal.Serie = "áéíóúñÁÉÍÓÚÑ"

	xml1, err := cfdi40.Marshal(comprobanteOriginal)
	if err != nil {
		t.Errorf("Error Marshal(comprobanteOriginal): %s", err.Error())
		return
	}

	comprobanteUnmarshaled, err := cfdi40.Unmarshal(xml1)
	if err != nil {
		t.Errorf("Error Unmarshal(xml1): %s", err.Error())
		return
	}

	err = cfdi40.CompareEqual(comprobanteOriginal, comprobanteUnmarshaled)
	assert.NoError(t, err)
}
