package tfd11

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/veyronifs/cfdi-go/types"
)

func TestUnmarshal(t *testing.T) {
	var xmlOriginal []byte
	{
		xmlOriginal = []byte(`<tfd:TimbreFiscalDigital xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sat.gob.mx/TimbreFiscalDigital http://www.sat.gob.mx/sitio_internet/cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd" xmlns:tfd="http://www.sat.gob.mx/TimbreFiscalDigital" FechaTimbrado="2022-02-03T16:34:31" NoCertificadoSAT="30001000000400002495" RfcProvCertif="SPR190613I52" SelloCFD="CVq5UXzahvFIv0C17RUE4TELUfavBYNHTjUczcN6HSrByR7WrMzHM1HHtaqDAfhAxRPQ3PI5rGZvryM5S1cMnhO/ifl1N51/WqUsjcH30CNENRhIXDgwoeL4mrMKb3tsi+d1JnRBwrEwQCo+ClHpfedLH9r5PhzISbFRpOHy6Dw4D9tEZuApLm7AyaDo0d0o6vzVmZ66ZakbVWqbCZL8/sZ/wRXa8XPuR/9QBZab7ivITjIgFn9gx5jTOt2i+bhH0HMR118x8lCpiNnj4BAjb6oPC0n96TptOJOjv9Ilco1q1XyBnOdNwC8RPFbRav9VRbYumlwNRFgf5aT90bU8/g==" SelloSAT="Supuj6zzoj5xrJUdb4OvExbT49jCBccjxrnMkuCpng1/LSh2f7GvBneqGmnVuqXTKbDUGwqiZGOT2qIkrguoetcSsnLhhG4/iw/0FR0DDAfYTYC+iJwtTorcW8q8XnKmiruG9rSpjwlxhKS94w/mJW+tUYLUrRbEjEZuOlvixVshElNUZXXeE9NChJT06iV/wmzp5Cgd4mbLQb3DubwrQQeHsZLSFIsc4qRRusnpR6YNOPBBzOinkuMTs3IS+4HodN/mZypNumh0JMNQfd3YO2dUCP85m8ZJ1ECk66iJWBZmu+NCuGDqkT5H3m1e54vIY0Uacja07m2WohHL6PAIEQ==" UUID="86ae7b6e-27e8-4714-b11a-535d623e1469" Version="1.1"/>`)
	}

	var tfdUnmarshaled TimbreFiscalDigital
	err := xml.Unmarshal(xmlOriginal, &tfdUnmarshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlOriginal): %s", err)
		return
	}

	xmlMarshaled, err := Marshal(&tfdUnmarshaled)
	if err != nil {
		t.Errorf("Error Encode(tfdUnmarshal): %s", err)
		return
	}

	var tfdUnmarshaled2 TimbreFiscalDigital
	err = xml.Unmarshal(xmlMarshaled, &tfdUnmarshaled2)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshaled): %s", err)
	}
	AssertEqual(t, &tfdUnmarshaled, &tfdUnmarshaled2)
}

func TestMarshal(t *testing.T) {
	fecha, _ := types.TFechaHParse("2022-02-03T16:34:31")
	tfd := &TimbreFiscalDigital{
		Version:          "1.1",
		UUID:             "86ae7b6e-27e8-4714-b11a-535d623e1469",
		FechaTimbrado:    fecha,
		RfcProvCertif:    "SPR190613I52",
		Leyenda:          "",
		SelloCFD:         "SELLO",
		NoCertificadoSAT: "30001000000400002495",
		SelloSAT:         "SELLOSAT",
	}

	xmlMarshaled, err := Marshal(tfd)
	if err != nil {
		t.Errorf("Error Marshal(TimbreFiscalDigital): %s", err)
		return
	}

	var tfdUnmarshaled TimbreFiscalDigital
	err = xml.Unmarshal(xmlMarshaled, &tfdUnmarshaled)
	if err != nil {
		t.Errorf("Error Unmarshal(xmlMarshaled): %s", err)
	}
	AssertEqual(t, tfd, &tfdUnmarshaled)
	fmt.Println(tfdUnmarshaled)
}
