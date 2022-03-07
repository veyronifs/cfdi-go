package comext11

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`
			<cce11:ComercioExterior Version="1.1" TipoOperacion="2" ClaveDePedimento="A1" CertificadoOrigen="0" Incoterm="CFR" Subdivision="0" TipoCambioUSD="19.0894" TotalUSD="65.00">
			<cce11:Emisor Curp="MAHJ280603MSPRRV09">
				<cce11:Domicilio Calle="CUAJIMALPA" NumeroExterior="276" NumeroInterior="64" Colonia="Cualquier Colonia" Localidad="10" Referencia="Cualquier referencia" Municipio="120" Estado="JAL" Pais="MEX" CodigoPostal="45199" />
			</cce11:Emisor>
			<cce11:Propietario NumRegIdTrib="alkaf141" ResidenciaFiscal="USA"/>
			<cce11:Receptor>
				<cce11:Domicilio Calle="Calle 2nda #3456 Zona centro" NumeroExterior="5970" Colonia="Zona centro" Referencia="A la vuelta de la esquina" Municipio="Tijuana" Estado="NY" Pais="USA" CodigoPostal="22000" />
			</cce11:Receptor>
			<cce11:Mercancias>
				<cce11:Mercancia NoIdentificacion="1" FraccionArancelaria="68101101" CantidadAduana="13.88" UnidadAduana="01" ValorUnitarioAduana="4.68" ValorDolares="65.00"/>
			</cce11:Mercancias>
		</cce11:ComercioExterior>`)
	}

	_ = xmlOriginal

	// comextUnmarshaled, err := Unmarshal(xmlOriginal)
	// if err != nil {
	// 	t.Errorf("Error Unmarshal(xmlOriginal): %s", err)
	// 	return
	// }

	expected := &ComercioExterior{
		Emisor: &Emisor{
			Domicilio: &Domicilio{
				Calle:          "CUAJIMALPA",
				NumeroExterior: "276",
				NumeroInterior: "64",
				Colonia:        "Cualquier Colonia",
				Localidad:      "10",
				Referencia:     "Cualquier referencia",
				Municipio:      "120",
				Estado:         "JAL",
				Pais:           "MEX",
				CodigoPostal:   "45199",
			},
			Curp: "MAHJ280603MSPRRV09",
		},
		Propietarios: []*Propietario{
			{
				NumRegIdTrib:     "alkaf141",
				ResidenciaFiscal: "USA",
			},
			{
				NumRegIdTrib:     "wiwieuyer",
				ResidenciaFiscal: "MXN",
			},
			{
				NumRegIdTrib:     "wiwieuyeQWERTYr",
				ResidenciaFiscal: "USA",
			},
		},
		Receptor: &Receptor{
			Domicilio: &Domicilio{
				Calle:          "Calle 2nda #3456 Zona centro",
				NumeroExterior: "5970",
				Colonia:        "Zona centro",
				Referencia:     "A la vuelta de la esquina",
				Municipio:      "Tijuana",
				Estado:         "NY",
				Pais:           "USA",
				CodigoPostal:   "22000",
			},
		},
		Destinatarios: []*Destinatario{
			{
				Domicilios: []*Domicilio{
					{
						Calle:          "Calle 2nda #3456 Zona centro",
						NumeroExterior: "5970",
						Colonia:        "Zona centro",
						Referencia:     "A la vuelta de la esquina",
						Municipio:      "Tijuana",
						Estado:         "NY",
						Pais:           "USA",
						CodigoPostal:   "22000",
					},
					{
						Calle:          "Calle 2nda #3456 Zona centro",
						NumeroExterior: "5970",
						Colonia:        "Zona centro",
						Referencia:     "A la vuelta de la esquina",
						Municipio:      "Tijuana",
						Estado:         "NY",
						Pais:           "USA",
						CodigoPostal:   "22000",
					},
				},
				NumRegIdTrib: "alkaf141",
				Nombre:       "Nombre 0",
			},
			{
				Domicilios: []*Domicilio{
					{
						Calle:          "Calle 2nda #3456 Zona centro",
						NumeroExterior: "5970",
						Colonia:        "Zona centro",
						Referencia:     "A la vuelta de la esquina",
						Municipio:      "Tijuana",
						Estado:         "NY",
						Pais:           "USA",
						CodigoPostal:   "22000",
					},
				},
				NumRegIdTrib: "wiwieuyer",
				Nombre:       "Nombre 1",
			},
			{
				NumRegIdTrib: "wiwieuyeQWERTYr",
				Nombre:       "Nombre 2",
			},
		},
		Mercancias: Mercancias{
			{
				NoIdentificacion:    "1",
				FraccionArancelaria: "68101101",
				CantidadAduana:      decimal.NewFromFloat(13.88),
				UnidadAduana:        "01",
				ValorUnitarioAduana: decimal.NewFromFloat(4.68),
				ValorDolares:        decimal.NewFromFloat(65.00),
			},
			{
				NoIdentificacion:    "3",
				FraccionArancelaria: "654168156",
				CantidadAduana:      decimal.NewFromFloat(13.88),
				UnidadAduana:        "02",
				ValorUnitarioAduana: decimal.NewFromFloat(4.68),
				ValorDolares:        decimal.NewFromFloat(65.00),
			},
		},
		Version:          "1.1",
		TipoOperacion:    "2",
		ClaveDePedimento: "A1",
		Incoterm:         "CFR",
		TipoCambioUSD:    decimal.NewFromFloat(19.08),
		TotalUSD:         decimal.NewFromFloat(65.00),
	}

	// err = CompareEqual(expected, comextUnmarshaled)
	// assert.NoError(t, err)

	xmlMarshaled, err := Marshal(expected)
	if err != nil {
		t.Errorf("Error Marshal(expected): %s", err)
	}

	comextUnmarshaled, err := Unmarshal(xmlMarshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshaled): %s", err)
	}

	err = CompareEqual(expected, comextUnmarshaled)
	assert.NoError(t, err)
}