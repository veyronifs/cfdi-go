package cartaporte30

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/types"
)

func TestUnmarshal(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`
			<cartaporte30:CartaPorte Version="2.0" TranspInternac="No" TotalDistRec="1963">
				<cartaporte30:Ubicaciones>
					<cartaporte30:Ubicacion TipoUbicacion="Origen" RFCRemitenteDestinatario="IXS7607092R5" FechaHoraSalidaLlegada="2021-12-15T12:24:21">
						<cartaporte30:Domicilio Calle="Filomeno Mata 922" NumeroExterior="35" Referencia="Filomeno Mata 922, Ejidos de Santa María Aztahuacán, DIF" Estado="DIF" Pais="MEX" CodigoPostal="02900" />
					</cartaporte30:Ubicacion>
					<cartaporte30:Ubicacion TipoUbicacion="Destino" RFCRemitenteDestinatario="AMX120216J89" FechaHoraSalidaLlegada="2021-12-15T03:00:00" DistanciaRecorrida="1963">
						<cartaporte30:Domicilio Calle="TAMAULIPAS 7" NumeroExterior="7" Referencia="TAMAULIPAS 7, CENTRO, SON" Estado="SON" Pais="MEX" CodigoPostal="83001" />
					</cartaporte30:Ubicacion>
				</cartaporte30:Ubicaciones>
				<cartaporte30:Mercancias PesoBrutoTotal="1" UnidadPeso="KGM" NumTotalMercancias="1">
					<cartaporte30:Mercancia BienesTransp="50425400" Descripcion="CACAHUATE RUNNER 38/42 B R" Cantidad="1" ClaveUnidad="KGM" PesoEnKg="1" />
					<cartaporte30:Autotransporte PermSCT="TPAF02" NumPermisoSCT="09381581/080002">
						<cartaporte30:IdentificacionVehicular ConfigVehicular="T3S2" PlacaVM="LE57937" AnioModeloVM="2008" />
						<cartaporte30:Seguros AseguraRespCivil="Seguros Atlas" PolizaRespCivil="MS1-1-7-38294" AseguraMedAmbiente="Seguros Atlas" PolizaMedAmbiente="MS1-1-7-38294" AseguraCarga="Seguros Atlas" PolizaCarga="MS1-2-2-1353" />
						<cartaporte30:Remolques>
							<cartaporte30:Remolque SubTipoRem="CTR004" Placa="6HU3452" />
						</cartaporte30:Remolques>
					</cartaporte30:Autotransporte>
				</cartaporte30:Mercancias>
				<cartaporte30:FiguraTransporte>
					<cartaporte30:TiposFigura TipoFigura="01" RFCFigura="VAEE770919BE9" NombreFigura="MANUEL ENRIQUE VELAZQUEZ ESPINOSA" NumLicencia="N03626578" />
				</cartaporte30:FiguraTransporte>
			</cartaporte30:CartaPorte>
		`)
	}

	cartaPorteUnmarshaled, err := Unmarshal(xmlOriginal)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlOriginal): %s", err)
	}

	xmlMarshaled, err := Marshal(cartaPorteUnmarshaled, "MXN")
	if err != nil {
		t.Errorf("Error Marshal(cartaPorteUnmarshal): %s", err)
	}

	cartaPorteUnmarshaled2, err := Unmarshal(xmlMarshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshaled): %s", err)
	}
	err = CompareEqual(cartaPorteUnmarshaled, cartaPorteUnmarshaled2)
	assert.NoError(t, err)

}

func TestMarshal(t *testing.T) {
	var cartaPorte *CartaPorte30
	{
		fechaHoraSalidaLlegada, _ := types.NewFechaH("2021-12-15T03:00:00")
		cartaPorte = &CartaPorte30{
			//Mercancias:0xc00010e100
			Version:           "3.0",
			TranspInternac:    "No",
			EntradaSalidaMerc: "Entrada",
			PaisOrigenDestino: "USA",
			ViaEntradaSalida:  "01",
			TotalDistRec:      decimal.New(1963, 0),
			Ubicaciones: []*Ubicacion{
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
					FechaHoraSalidaLlegada:      fechaHoraSalidaLlegada,
					TipoEstacion:                "PM",
					DistanciaRecorrida:          decimal.New(1963, 0),
					Domicilio: &Domicilio{
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
			Mercancias: &Mercancias{
				UnidadPeso:         "KGM",
				NumTotalMercancias: 1,
				PesoBrutoTotal:     decimal.New(199, 99),
				PesoNetoTotal:      decimal.New(199, 99),
				CargoPorTasacion:   decimal.New(199, 99),
				Mercancia: []*Mercancia{
					{
						BienesTransp:         "10101501",
						ClaveSTCC:            "10101501",
						Descripcion:          "Caja de 20 piezas",
						Cantidad:             decimal.New(10, 0),
						ClaveUnidad:          "CJ",
						Unidad:               "Caja",
						Dimensiones:          "20x20x20",
						MaterialPeligroso:    "Sí",
						CveMaterialPeligroso: "M0001",
						Embalaje:             "1A2",
						DescripEmbalaje:      "Caja de 20 piezas",
						PesoEnKg:             decimal.New(199, 99),
						ValorMercancia:       decimal.New(199, 99),
						Moneda:               "MXN",
						FraccionArancelaria:  "10101501",
						UUIDComercioExt:      "6713E766-DCA2-41AA-B1C0-020CFB60AC95",
						DocumentacionAduanera: []*DocumentacionAduanera{
							{NumPedimento: "123456789"},
						},
						GuiasIdentificacion: []*GuiasIdentificacion{
							{
								NumeroGuiaIdentificacion:  "123456789",
								DescripGuiaIdentificacion: "Guia de identificacion",
								PesoGuiaIdentificacion:    decimal.New(199, 99),
							},
						},
						DetalleMercancia: &DetalleMercancia{
							UnidadPesoMerc: "KGM",
							PesoBruto:      decimal.New(199, 99),
							PesoNeto:       decimal.New(199, 99),
							PesoTara:       decimal.New(199, 99),
							NumPiezas:      1,
						},
						CantidadTransporta: []*CantidadTransporta{
							{
								Cantidad:       decimal.New(1, 0),
								IDOrigen:       "00001",
								IDDestino:      "00002",
								CvesTransporte: "1A2",
							},
						},
					},
				},
				Autotransporte: &Autotransporte{
					PermSCT:       "TPAF02",
					NumPermisoSCT: "09381581/080002",
					IdentificacionVehicular: &IdentificacionVehicular{
						ConfigVehicular: "T3S2",
						PlacaVM:         "LE57937",
						AnioModeloVM:    "2008",
					},
					Seguros: &Seguros{
						AseguraRespCivil:   "Seguros Atlas",
						PolizaRespCivil:    "MS1-1-7-38294",
						AseguraMedAmbiente: "Seguros Atlas",
						PolizaMedAmbiente:  "MS1-1-7-38294",
						AseguraCarga:       "Seguros Atlas",
						PolizaCarga:        "MS1-2-2-1353",
						PrimaSeguro:        decimal.New(10, 0),
					},
					Remolques: []*Remolque{
						{SubTipoRem: "CTR004", Placa: "6HU3452"},
						{SubTipoRem: "CTR001", Placa: "6HU3112"},
					},
				},
			},
			FiguraTransporte: &FiguraTransporte{
				TiposFigura: []*TiposFigura{
					{
						TipoFigura:             "01",
						RFCFigura:              "VAEE770919BE9",
						NumLicencia:            "N03626578",
						NombreFigura:           "MANUEL ENRIQUE VELAZQUEZ ESPINOSA",
						NumRegIdTribFigura:     "0132456789",
						ResidenciaFiscalFigura: types.Pais("MEX"),
						PartesTransporte: []*PartesTransporte{
							{ParteTransporte: "ABCD"},
						},
						Domicilio: &Domicilio{
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
		}
	}
	xmlMarshaled, err := Marshal(cartaPorte, "MXN")
	if err != nil {
		t.Errorf("Error Marshal(cartaPorte): %s", err)
	}

	cartaPorteUnmarshaled, err := Unmarshal(xmlMarshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshaled): %s", err)
	}
	err = CompareEqual(cartaPorte, cartaPorteUnmarshaled)
	assert.NoError(t, err)
}
