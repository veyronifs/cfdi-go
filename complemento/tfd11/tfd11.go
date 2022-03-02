package tfd11

import (
	"encoding/xml"

	"github.com/veyronifs/cfdi-go/types"
)

func Unmarshal(b []byte) (*TimbreFiscalDigital, error) {
	tfd := &TimbreFiscalDigital{}
	if err := xml.Unmarshal(b, tfd); err != nil {
		return nil, err
	}
	return tfd, nil
}

// TimbreFiscalDigital Complemento requerido para el Timbrado Fiscal Digital que da validez al Comprobante fiscal digital por Internet.
type TimbreFiscalDigital struct {
	// Version Atributo requerido para la expresión de la versión del estándar del Timbre Fiscal Digital
	Version string `xml:"Version,attr"`
	// UUID Atributo requerido para expresar los 36 caracteres del folio fiscal (UUID) de la transacción de timbrado conforme al estándar RFC 4122
	UUID string `xml:"UUID,attr"`
	// FechaTimbrado Atributo requerido para expresar la fecha y hora, de la generación del timbre por la certificación digital del SAT. Se expresa en la forma AAAA-MM-DDThh:mm:ss y debe corresponder con la hora de la Zona Centro del Sistema de Horario en México.
	FechaTimbrado types.FechaH `xml:"FechaTimbrado,attr"`
	// RfcProvCertif Atributo requerido para expresar el RFC del proveedor de certificación de comprobantes fiscales digitales que genera el timbre fiscal digital.
	RfcProvCertif string `xml:"RfcProvCertif,attr"`
	// Leyenda Atributo opcional para registrar información que el SAT comunique a los usuarios del CFDI.
	Leyenda string `xml:"Leyenda,attr,omitempty"`
	// SelloCFD Atributo requerido para contener el sello digital del comprobante fiscal o del comprobante de retenciones, que se ha timbrado. El sello debe ser expresado como una cadena de texto en formato Base 64.
	SelloCFD string `xml:"SelloCFD,attr"`
	// NoCertificadoSAT Atributo requerido para expresar el número de serie del certificado del SAT usado para generar el sello digital del Timbre Fiscal Digital.
	NoCertificadoSAT string `xml:"NoCertificadoSAT,attr"`
	// SelloSAT Atributo requerido para contener el sello digital del Timbre Fiscal Digital, al que hacen referencia las reglas de la Resolución Miscelánea vigente. El sello debe ser expresado como una cadena de texto en formato Base 64.
	SelloSAT string `xml:"SelloSAT,attr"`
}
