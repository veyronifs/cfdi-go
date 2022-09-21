package balanza

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/types"
)

func TestUnmarshal(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<catalogocuentas:Catalogo xmlns:catalogocuentas="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/BalanzaComprobacion"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/BalanzaComprobacion http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/BalanzaComprobacion/BalanzaComprobacion_1_3.xsd" Version="1.3" RFC="AAA010101AAA" Mes="12" Anio="2020" TipoEnvio="N" FechaModBal="2020-12-31" Sello="SELLO" noCertificado="12345678901234567890" Certificado="CERTIFICADO">
			<catalogocuentas:Ctas NumCta="AAAAA" SaldoIni="1000.00" Debe="500.00" Haber="300.00" SaldoFin="700.00"/>
			<catalogocuentas:Ctas NumCta="BBBBB" SaldoIni="2000.00" Debe="0" Haber="300.00" SaldoFin="1700.00"/>
		</catalogocuentas:Catalogo>
		`)
	}

	catalogocuentasUnmarshaled, err := Unmarshal(xmlOriginal)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlOriginal): %s", err)
	}

	fecha, _ := types.NewFecha("2020-12-31")
	expected := &Balanza{
		Version:       "1.3",
		RFC:           "AAA010101AAA",
		Mes:           12,
		Anio:          2020,
		TipoEnvio:     TipoEnvioN,
		FechaModBal:   fecha,
		Sello:         "SELLO",
		NoCertificado: "12345678901234567890",
		Certificado:   "CERTIFICADO",
		Ctas: []*Cta{
			{
				NumCta:   "AAAAA",
				SaldoIni: decimal.NewFromFloat(1000),
				Debe:     decimal.NewFromFloat(500),
				Haber:    decimal.NewFromFloat(300),
				SaldoFin: decimal.NewFromFloat(700),
			},
			{
				NumCta:   "BBBBB",
				SaldoIni: decimal.NewFromFloat(2000),
				Debe:     decimal.NewFromFloat(0),
				Haber:    decimal.NewFromFloat(300),
				SaldoFin: decimal.NewFromFloat(1700),
			},
		},
	}

	err = CompareEqual(catalogocuentasUnmarshaled, expected)
	assert.NoError(t, err)

}

func TestMarshal(t *testing.T) {
	fecha, _ := types.NewFecha("2020-12-31")
	original := &Balanza{
		Version:       "1.3",
		RFC:           "AAA010101AAA",
		Mes:           12,
		Anio:          2020,
		TipoEnvio:     TipoEnvioN,
		FechaModBal:   fecha,
		Sello:         "SELLO",
		NoCertificado: "12345678901234567890",
		Certificado:   "CERTIFICADO",
		Ctas: []*Cta{
			{
				NumCta:   "AAAAA",
				SaldoIni: decimal.NewFromFloat(1000),
				Debe:     decimal.NewFromFloat(500),
				Haber:    decimal.NewFromFloat(300),
				SaldoFin: decimal.NewFromFloat(700),
			},
			{
				NumCta:   "BBBBB",
				SaldoIni: decimal.NewFromFloat(2000),
				Debe:     decimal.NewFromFloat(0),
				Haber:    decimal.NewFromFloat(300),
				SaldoFin: decimal.NewFromFloat(1700),
			},
		},
	}
	marshaled, err := Marshal(original)
	if err != nil {
		t.Errorf("Error Marshal(original): %s", err)
		return
	}
	unmarshaled, err := Unmarshal(marshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(marshaled): %s", err)
		return
	}

	err = CompareEqual(unmarshaled, original)
	assert.NoError(t, err)
}

func TestBalanzaFileName(t *testing.T) {
	tests := []struct {
		name     string
		balanza  *Balanza
		expected string
	}{
		{
			name: "Mes 12",
			balanza: &Balanza{
				RFC:       "AAA010101AAA",
				Anio:      2020,
				Mes:       12,
				TipoEnvio: TipoEnvioN,
			},
			expected: "AAA010101AAA202012BN",
		},
		{
			name: "Mes 01",
			balanza: &Balanza{
				RFC:       "AAA010101AAA",
				Anio:      2020,
				Mes:       1,
				TipoEnvio: TipoEnvioC,
			},
			expected: "AAA010101AAA202001BC",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			archivo := test.balanza.Archivo()
			actual := archivo.FileName()
			assert.Equal(t, test.expected, actual)
		})
	}

}
