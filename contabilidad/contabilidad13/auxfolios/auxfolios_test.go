package auxfolios

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
        <RepAux:RepAuxFol xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:RepAux="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios" xsi:schemaLocation="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios https://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios/AuxiliarFolios_1_3.xsd" Version="1.3" RFC="AAA010101AAA" Mes="01" Anio="2015" TipoSolicitud="AF" NumOrden="ABC3087514/62" NumTramite="AB123456789012" Sello="TEST" noCertificado="12345678901234567890" Certificado="TEST">
            <RepAux:DetAuxFol NumUnIdenPol="1334703" Fecha="2020-12-01">
				<RepAux:ComprNal UUID_CFDI="419AA1DD-5FEB-4439-BD20-B9CC5716334A" MontoTotal="1500" RFC="AAA010101AAB" MetPagoAux="Metodo1" Moneda="MXN" TipCamb="52"/>
				<RepAux:ComprNal UUID_CFDI="529AA1DD-7MAR-4439-BD20-B9CC5716334A" MontoTotal="1850.23" RFC="BBB010101AAB" MetPagoAux="Metodo4" Moneda="MXN" TipCamb="18"/>
                <RepAux:ComprNalOtr CFD_CBB_Serie="AAAABBBBCC" CFD_CBB_NumFol="10010" MontoTotal="2500" RFC="AAA010101AAB" MetPagoAux="Metodo2" Moneda="USD" TipCamb="45"/>
                <RepAux:ComprExt NumFactExt="A1100000000000000010" TaxID="10" MontoTotal="2500" MetPagoAux="Metodo3" Moneda="MXN" TipCamb="26"/>
            </RepAux:DetAuxFol>
        </RepAux:RepAuxFol>
		`)
	}

	polizasperiodoUnmarshaled, err := Unmarshal(xmlOriginal)
	if err != nil {
		t.Fatalf("Error Unmarshal(xmlOriginal): %s", err)
	}

	fecha, err := types.NewFecha("2020-12-01")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	expected := &RepAuxFolios{
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
		DetAuxFolios: []*DetAuxFol{
			{
				NumUnIdenPol: "1334703",
				Fecha:        fecha,
				ComprNal: []*ComprNal{
					{
						UUIDCFDI:   "419AA1DD-5FEB-4439-BD20-B9CC5716334A",
						MontoTotal: decimal.NewFromFloat(1500),
						RFC:        "AAA010101AAB",
						MetPagoAux: "Metodo1",
						Moneda:     types.MonedaMXN,
						TipCamb:    decimal.NewFromFloat(52),
					},
					{
						UUIDCFDI:   "529AA1DD-7MAR-4439-BD20-B9CC5716334A",
						MontoTotal: decimal.NewFromFloat(1850.23),
						RFC:        "BBB010101AAB",
						MetPagoAux: "Metodo4",
						Moneda:     types.MonedaMXN,
						TipCamb:    decimal.NewFromFloat(18),
					},
				},
				ComprNalOtr: []*ComprNalOtr{
					{
						CFDCBBSerie:  "AAAABBBBCC",
						CFDCBBNumFol: 10010,
						MontoTotal:   decimal.NewFromFloat(2500),
						RFC:          "AAA010101AAB",
						MetPagoAux:   "Metodo2",
						Moneda:       types.MonedaUSD,
						TipCamb:      decimal.NewFromFloat(45),
					},
				},
				ComprExt: []*ComprExt{
					{
						NumFactExt: "A1100000000000000010",
						TaxID:      "10",
						MontoTotal: decimal.NewFromFloat(2500),
						MetPagoAux: "Metodo3",
						Moneda:     types.MonedaMXN,
						TipCamb:    decimal.NewFromFloat(26),
					},
				},
			},
		},
	}

	err = CompareEqual(polizasperiodoUnmarshaled, expected)
	assert.NoError(t, err)
}

func TestMarshal(t *testing.T) {

	fecha, err := types.NewFecha("2020-12-01")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	original := &RepAuxFolios{
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
		DetAuxFolios: []*DetAuxFol{
			{
				NumUnIdenPol: "1334703",
				Fecha:        fecha,
				ComprNal: []*ComprNal{
					{
						UUIDCFDI:   "419AA1DD-5FEB-4439-BD20-B9CC5716334A",
						MontoTotal: decimal.NewFromFloat(1500),
						RFC:        "AAA010101AAB",
						MetPagoAux: "Metodo1",
						Moneda:     types.MonedaMXN,
						TipCamb:    decimal.NewFromFloat(52),
					},
					{
						UUIDCFDI:   "529AA1DD-7MAR-4439-BD20-B9CC5716334A",
						MontoTotal: decimal.NewFromFloat(1850.23),
						RFC:        "BBB010101AAB",
						MetPagoAux: "Metodo4",
						Moneda:     types.MonedaMXN,
						TipCamb:    decimal.NewFromFloat(18),
					},
				},
				ComprNalOtr: []*ComprNalOtr{
					{
						CFDCBBSerie:  "AAAABBBBCC",
						CFDCBBNumFol: 10010,
						MontoTotal:   decimal.NewFromFloat(2500),
						RFC:          "AAA010101AAB",
						MetPagoAux:   "Metodo2",
						Moneda:       types.MonedaUSD,
						TipCamb:      decimal.NewFromFloat(45),
					},
				},
				ComprExt: []*ComprExt{
					{
						NumFactExt: "A1100000000000000010",
						TaxID:      "10",
						MontoTotal: decimal.NewFromFloat(2500),
						MetPagoAux: "Metodo3",
						Moneda:     types.MonedaMXN,
						TipCamb:    decimal.NewFromFloat(26),
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
	// fmt.Println(string(marshaled))
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
