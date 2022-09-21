package polizasperiodo

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
		<PLZ:Polizas xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:PLZ="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo" xsi:schemaLocation="http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo https://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo/PolizasPeriodo_1_3.xsd" Version="1.3" RFC="AAA010101AAA" Mes="1" Anio="2015" TipoSolicitud="AF" NumOrden="ABC3087514/62" NumTramite="AB123456789012" Sello="TEST" noCertificado="12345678901234567890" Certificado="TEST">
			<PLZ:Poliza NumUnIdenPol="1334703" Fecha="2020-12-01" Concepto="Abono TEST">
				<PLZ:Transaccion NumCta="1110000000" DesCta="CAJA CHICA" Concepto="Abono" Debe="1400" Haber="1500">
					<PLZ:CompNal UUID_CFDI="419AA1DD-5FEB-4439-BD20-B9CC5716334A" RFC="AAA010101AAB" MontoTotal="1500" Moneda="MXN" TipCamb="1800"/>
					<PLZ:CompNalOtr CFD_CBB_Serie="AAAABBBBCC" CFD_CBB_NumFol="10010" RFC="AAA010101AAB" MontoTotal="2500" Moneda="MNX" TipCamb="2600"/>
					<PLZ:CompExt NumFactExt="A1100000000000000010" TaxID="10" MontoTotal="2500" Moneda="MXN" TipCamb="26"/>
					<PLZ:Cheque Num="123456" BanEmisNal="Banco Emi Nal" BanEmisExt="Banco Emi Ext" CtaOri="1542456356987823" Fecha="2020-02-01" Benef="Humberto Flores Garcia" RFC="AAA010101ABC" Monto="1500" Moneda="MXN" TipCamb="2520"/>
					<PLZ:Transferencia CtaOri="1542456356987824" BancoOriNal="Banco Ori Nal" BancoOriExt="Banco Ori Ext" CtaDest="1542456356987832" BancoDestNal="Banco Dest Nal" BancoDestExt="Banco Dest Ext" Fecha="2018-10-08" Benef="Jose Hernandez Torres" RFC="AAA010101CAB" Monto="27500" Moneda="MXN" TipCamb="2620"/>
					<PLZ:OtrMetodoPago MetPagoPol="Efectivo" Fecha="2019-11-20" Benef="Ricardo Maya Herrera" RFC="AAA0101011LF" Monto="3200" Moneda="MXN" TipCamb="2420"/>
				</PLZ:Transaccion>
			</PLZ:Poliza>
		</PLZ:Polizas>
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

	// fecha1, err := types.NewFecha("2020-02-01")
	// if err != nil {
	// 	t.Fatalf("Error fecha: %s", err)
	// }

	fecha2, err := types.NewFecha("2018-10-08")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	fecha3, err := types.NewFecha("2019-11-20")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	expected := &Polizas{
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
		Polizas: []*Poliza{
			{
				NumUnIdenPol: "1334703",
				Fecha:        fecha,
				Concepto:     "Abono TEST",
				Transaccion: []*Transaccion{
					{
						NumCta:   "1110000000",
						DesCta:   "CAJA CHICA",
						Concepto: "Abono",
						Debe:     decimal.NewFromFloat(1400),
						Haber:    decimal.NewFromFloat(1500),
						CompNal: []*CompNal{
							{
								UUIDCFDI:   "419AA1DD-5FEB-4439-BD20-B9CC5716334A",
								RFC:        "AAA010101AAB",
								MontoTotal: decimal.NewFromFloat(1500),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(1800),
							},
						},
						CompNalOtr: []*CompNalOtr{
							{
								CFDCBBSerie:  "AAAABBBBCC",
								CFDCBBNumFol: 10010,
								RFC:          "AAA010101AAB",
								MontoTotal:   decimal.NewFromFloat(2500),
								Moneda:       types.MonedaMXN,
								TipCamb:      decimal.NewFromFloat(2600),
							},
						},
						CompExt: []*CompExt{
							{
								NumFactExt: "A1100000000000000010",
								TaxID:      "10",
								MontoTotal: decimal.NewFromFloat(2500),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(26),
							},
						},
						Cheque: []*Cheque{
							{
								Num:        "123456",
								BanEmisNal: "Banco Emi Nal",
								BanEmisExt: "Banco Emi Ext",
								CtaOri:     "1542456356987823",
								Fecha:      fecha,
								Benef:      "Humberto Flores Garcia",
								RFC:        "AAA010101ABC",
								Monto:      decimal.NewFromFloat(1500),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(2520),
							},
						},
						Transferencia: []*Transferencia{
							{
								CtaOri:       "1542456356987824",
								BancoOriNal:  "Banco Ori Nal",
								BancoOriExt:  "Banco Ori Ext",
								CtaDest:      "1542456356987832",
								BancoDestNal: "Banco Dest Nal",
								BancoDestExt: "Banco Dest Ext",
								Fecha:        fecha2,
								Benef:        "Jose Hernandez Torres",
								RFC:          "AAA010101CAB",
								Monto:        decimal.NewFromFloat(27500),
								Moneda:       types.MonedaMXN,
								TipCamb:      decimal.NewFromFloat(2620),
							},
						},
						OtrMetodoPago: []*OtrMetodoPago{
							{
								MetPagoPol: "Efectivo",
								Fecha:      fecha3,
								Benef:      "Ricardo Maya Herrera",
								RFC:        "AAA0101011LF",
								Monto:      decimal.NewFromFloat(3200),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(2420),
							},
						},
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

	// fecha1, err := types.NewFecha("2020-02-01")
	// if err != nil {
	// 	t.Fatalf("Error fecha: %s", err)
	// }

	fecha2, err := types.NewFecha("2018-10-08")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	fecha3, err := types.NewFecha("2019-11-20")
	if err != nil {
		t.Fatalf("Error fecha: %s", err)
	}

	original := &Polizas{
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
		Polizas: []*Poliza{
			{
				NumUnIdenPol: "1334703",
				Fecha:        fecha,
				Concepto:     "Abono TEST",
				Transaccion: []*Transaccion{
					{
						NumCta:   "1110000000",
						DesCta:   "CAJA CHICA",
						Concepto: "Abono",
						Debe:     decimal.NewFromFloat(1400),
						Haber:    decimal.NewFromFloat(1500),
						CompNal: []*CompNal{
							{
								UUIDCFDI:   "419AA1DD-5FEB-4439-BD20-B9CC5716334A",
								RFC:        "AAA010101AAB",
								MontoTotal: decimal.NewFromFloat(1500),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(1800),
							},
						},
						CompNalOtr: []*CompNalOtr{
							{
								CFDCBBSerie:  "AAAABBBBCC",
								CFDCBBNumFol: 10010,
								RFC:          "AAA010101AAB",
								MontoTotal:   decimal.NewFromFloat(2500),
								Moneda:       types.MonedaMXN,
								TipCamb:      decimal.NewFromFloat(2600),
							},
						},
						CompExt: []*CompExt{
							{
								NumFactExt: "A1100000000000000010",
								TaxID:      "10",
								MontoTotal: decimal.NewFromFloat(2500),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(2600),
							},
						},
						Cheque: []*Cheque{
							{
								Num:        "123456",
								BanEmisNal: "Banco Emi Nal",
								BanEmisExt: "Banco Emi Ext",
								CtaOri:     "1542456356987823",
								Fecha:      fecha,
								Benef:      "Humberto Flores Garcia",
								RFC:        "AAA010101ABC",
								Monto:      decimal.NewFromFloat(1500),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(2520),
							},
						},
						Transferencia: []*Transferencia{
							{
								CtaOri:       "1542456356987824",
								BancoOriNal:  "Banco Ori Nal",
								BancoOriExt:  "Banco Ori Ext",
								CtaDest:      "1542456356987832",
								BancoDestNal: "Banco Dest Nal",
								BancoDestExt: "Banco Dest Ext",
								Fecha:        fecha2,
								Benef:        "Jose Hernandez Torres",
								RFC:          "AAA010101CAB",
								Monto:        decimal.NewFromFloat(27500),
								Moneda:       types.MonedaMXN,
								TipCamb:      decimal.NewFromFloat(2620),
							},
						},
						OtrMetodoPago: []*OtrMetodoPago{
							{
								MetPagoPol: "Efectivo",
								Fecha:      fecha3,
								Benef:      "Ricardo Maya Herrera",
								RFC:        "AAA0101011LF",
								Monto:      decimal.NewFromFloat(3200),
								Moneda:     types.MonedaMXN,
								TipCamb:    decimal.NewFromFloat(2420),
							},
						},
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
