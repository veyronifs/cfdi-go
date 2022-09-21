package auxctas

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/types"
)

func TestUnmarshal(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
        <AuxiliarCtas:AuxiliarCtas xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:AuxiliarCtas="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas" xsi:schemaLocation="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas https://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarCtas/AuxiliarCtas_1_3.xsd" Version="1.3" RFC="AAA010101AAA" Mes="01" Anio="2015" TipoSolicitud="AF" NumOrden="ABC3087514/62" NumTramite="AB123456789012" Sello="TEST" noCertificado="12345678901234567890" Certificado="TEST">
            <AuxiliarCtas:Cuenta NumCta="1542456356987832" DesCta="Cuenta TEST" SaldoIni="0.00" SaldoFin="8523.45">
				<AuxiliarCtas:DetalleAux Fecha="2020-05-10" NumUnIdenPol="9334703" Concepto="Concepto Ctas TEST" Debe="1230.00" Haber="1236.05"/>
            </AuxiliarCtas:Cuenta>
			<AuxiliarCtas:Cuenta NumCta="1542456356988965" DesCta="Descrip TEST" SaldoIni="1520.36" SaldoFin="20145.20">
				<AuxiliarCtas:DetalleAux Fecha="2020-05-10" NumUnIdenPol="9334725" Concepto="Ctas Concepto TEST" Debe="0.00" Haber="456.25"/>
            </AuxiliarCtas:Cuenta>
        </AuxiliarCtas:AuxiliarCtas>
		`)
	}

	auxctasUnmarshaled, err := Unmarshal(xmlOriginal)
	if err != nil {
		t.Fatalf("Error Unmarshal(xmlOriginal): %s", err)
	}

	fecha, err := types.NewFecha("2020-05-10")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	expected := &AuxiliarCtas{
		Version:       "1.3",
		RFC:           "AAA010101AAA",
		Mes:           1,
		Anio:          2015,
		TipoSolicitud: "AF",
		NumOrden:      "ABC3087514/62",
		NumTramite:    "AB123456789012",
		Sello:         "TEST",
		NoCertificado: "12345678901234567890",
		Certificado:   "TEST",
		Cuentas: []*Cuenta{
			{
				NumCta:   "1542456356987832",
				DesCta:   "Cuenta TEST",
				SaldoIni: decimal.NewFromFloat(0.00),
				SaldoFin: decimal.NewFromFloat(8523.45),
				DetallesAux: []*DetalleAux{
					{
						Fecha:        fecha,
						NumUnIdenPol: "9334703",
						Concepto:     "Concepto Ctas TEST",
						Debe:         decimal.NewFromFloat(1230.00),
						Haber:        decimal.NewFromFloat(1236.05),
					},
				},
			},
			{
				NumCta:   "1542456356988965",
				DesCta:   "Descrip TEST",
				SaldoIni: decimal.NewFromFloat(1520.36),
				SaldoFin: decimal.NewFromFloat(20145.20),
				DetallesAux: []*DetalleAux{
					{
						Fecha:        fecha,
						NumUnIdenPol: "9334725",
						Concepto:     "Ctas Concepto TEST",
						Debe:         decimal.NewFromFloat(0.00),
						Haber:        decimal.NewFromFloat(456.25),
					},
				},
			},
		},
	}

	err = CompareEqual(auxctasUnmarshaled, expected)
	assert.NoError(t, err)
}

func TestMarshal(t *testing.T) {

	fecha, err := types.NewFecha("2020-05-10")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	original := &AuxiliarCtas{
		Version:       "1.3",
		RFC:           "AAA010101AAA",
		Mes:           1,
		Anio:          2015,
		TipoSolicitud: "AF",
		NumOrden:      "ABC3087514/62",
		NumTramite:    "AB123456789012",
		Sello:         "TEST",
		NoCertificado: "12345678901234567890",
		Certificado:   "TEST",
		Cuentas: []*Cuenta{
			{
				NumCta:   "1542456356987832",
				DesCta:   "Cuenta TEST",
				SaldoIni: decimal.NewFromFloat(0.00),
				SaldoFin: decimal.NewFromFloat(8523.45),
				DetallesAux: []*DetalleAux{
					{
						Fecha:        fecha,
						NumUnIdenPol: "9334703",
						Concepto:     "Concepto Ctas TEST",
						Debe:         decimal.NewFromFloat(1230.00),
						Haber:        decimal.NewFromFloat(1236.05),
					},
				},
			},
			{
				NumCta:   "1542456356988965",
				DesCta:   "Descrip TEST",
				SaldoIni: decimal.NewFromFloat(1520.36),
				SaldoFin: decimal.NewFromFloat(20145.20),
				DetallesAux: []*DetalleAux{
					{
						Fecha:        fecha,
						NumUnIdenPol: "9334725",
						Concepto:     "Ctas Concepto TEST",
						Debe:         decimal.NewFromFloat(0.00),
						Haber:        decimal.NewFromFloat(456.25),
					},
				},
			},
		},
	}

	marshaled, err := Marshal(original)
	if err != nil {
		t.Errorf("Error Marshal(original): %s", err)
		return
	}
	fmt.Println(string(marshaled))
	unmarshaled, err := Unmarshal(marshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(marshaled): %s", err)
		return
	}
	fmt.Println(unmarshaled)
	fmt.Println(original)
	err = CompareEqual(unmarshaled, original)
	assert.NoError(t, err)
}
