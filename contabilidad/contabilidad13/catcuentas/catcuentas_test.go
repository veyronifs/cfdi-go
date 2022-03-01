package catcuentas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
		<catalogocuentas:Catalogo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:catalogocuentas="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/CatalogoCuentas" xsi:schemaLocation="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/CatalogoCuentas https://www.sat.gob.mx/esquemas/ContabilidadE/1_3/CatalogoCuentas/CatalogoCuentas_1_3.xsd" Anio="2015" Mes="1" RFC="AAA010101AAA" Version="1.3" Sello="TEST" noCertificado="12345678901234567890" Certificado="TEST">
			<catalogocuentas:Ctas Natur="D" Nivel="1" Desc="FONDO FIJO DE CAJA" NumCta="1110000000" CodAgrup="101"/>
			<catalogocuentas:Ctas Natur="D" Nivel="2" SubCtaDe="1110000000" Desc="CAJA CHICA" NumCta="1110001000" CodAgrup="101.01"/>
		</catalogocuentas:Catalogo>
		`)
	}

	catalogocuentasUnmarshaled, err := Unmarshal(xmlOriginal)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlOriginal): %s", err)
	}

	expected := &Catalogo{
		Anio:          2015,
		Mes:           1,
		RFC:           "AAA010101AAA",
		Version:       "1.3",
		Sello:         "TEST",
		NoCertificado: "12345678901234567890",
		Certificado:   "TEST",
		Ctas: []*Cta{
			{
				Natur:    "D",
				Nivel:    1,
				Desc:     "FONDO FIJO DE CAJA",
				NumCta:   "1110000000",
				CodAgrup: "101",
			},
			{
				Natur:    "D",
				Nivel:    2,
				SubCtaDe: "1110000000",
				Desc:     "CAJA CHICA",
				NumCta:   "1110001000",
				CodAgrup: "101.01",
			},
		},
	}

	err = CompareEqual(catalogocuentasUnmarshaled, expected)
	assert.NoError(t, err)

}

func TestMarshal(t *testing.T) {
	original := &Catalogo{
		Anio:          2015,
		Mes:           1,
		RFC:           "AAA010101AAA",
		Version:       "1.3",
		Sello:         "TEST",
		NoCertificado: "12345678901234567890",
		Certificado:   "TEST",
		Ctas: []*Cta{
			{
				Natur:    "D",
				Nivel:    1,
				Desc:     "FONDO FIJO DE CAJA",
				NumCta:   "1110000000",
				CodAgrup: "101",
			},
			{
				Natur:    "D",
				Nivel:    2,
				SubCtaDe: "1110000000",
				Desc:     "CAJA CHICA",
				NumCta:   "1110001000",
				CodAgrup: "101.01",
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
