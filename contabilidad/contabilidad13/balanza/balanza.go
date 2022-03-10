package balanza

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

// Balanza Estándar de balanza de comprobación que se entrega como parte de la contabilidad electrónica.
type Balanza struct {
	// Ctas Nodo obligatorio para expresar el detalle de cada cuenta o subcuenta de la balanza de comprobación.
	Ctas []*Cta `xml:"Ctas"`
	// Version Atributo requerido para expresar la versión del formato.
	Version string `xml:"Version,attr"`
	// RFC Atributo requerido para expresar el RFC del contribuyente que envía los datos.
	RFC string `xml:"RFC,attr"`
	// Mes Atributo requerido para expresar el mes al que corresponde la balanza de comprobación.
	Mes int `xml:"Mes,attr"`
	// Anio Atributo requerido para expresar el año al que corresponde la balanza.
	Anio int `xml:"Anio,attr"`
	// TipoEnvio Atributo requerido para expresar el tipo de envío de la balanza (N - Normal; C - Complementaria).
	TipoEnvio TipoEnvio `xml:"TipoEnvio,attr"`
	// FechaModBal Atributo opcional para expresar la fecha de la última modificación contable de la balanza de comprobación. Es requerido cuando el tipo de Envío es complementario.
	FechaModBal types.Fecha `xml:"FechaModBal,attr,omitempty"`
	// Sello Atributo opcional para contener el sello digital del archivo de contabilidad electrónica. El sello deberá ser expresado cómo una cadena de texto en formato Base 64.
	Sello string `xml:"Sello,attr,omitempty"`
	// NoCertificado Atributo opcional para expresar el número de serie del certificado de sello digital que ampara el archivo de contabilidad electrónica, de acuerdo al acuse correspondiente a 20 posiciones otorgado por el sistema del SAT.
	NoCertificado string `xml:"noCertificado,attr,omitempty"`
	// Certificado Atributo opcional que sirve para expresar el certificado de sello digital que ampara al archivo de contabilidad electrónica como texto, en formato base 64.
	Certificado string `xml:"Certificado,attr,omitempty"`
}

var ErrRequerido = fmt.Errorf("requerido")

func (b Balanza) FileName() (string, error) {
	if b.RFC == "" {
		return "", fmt.Errorf("RFC %w", ErrRequerido)
	}
	if b.Mes == 0 {
		return "", fmt.Errorf("mes %w", ErrRequerido)
	}
	if b.Anio == 0 {
		return "", fmt.Errorf("año %w", ErrRequerido)
	}
	if b.TipoEnvio == "" {
		return "", fmt.Errorf("tipo envio %w", ErrRequerido)
	}
	return fmt.Sprintf("%s%d%02dB%s", b.RFC, b.Anio, b.Mes, b.TipoEnvio), nil
}

// Cta Nodo obligatorio para expresar el detalle de cada cuenta o subcuenta de la balanza de comprobación.
type Cta struct {
	// NumCta Atributo requerido para expresar la clave asignada con que se distingue la cuenta o subcuenta en el catálogo de cuentas del  contribuyente.
	NumCta string `xml:"NumCta,attr"`
	// SaldoIni Atributo requerido para expresar el monto del saldo inicial de la cuenta o subcuenta en el periodo. De acuerdo a la naturaleza de la cuenta o subcuenta, deberá de corresponder el saldo inicial, de lo contrario se entenderá que es un saldo inicial de naturaleza inversa. En caso de no existir dato, colocar cero (0).
	SaldoIni decimal.Decimal `xml:"SaldoIni,attr"`
	// Debe Atributo requerido para expresar el monto de los movimientos deudores de la cuenta o subcuenta. En caso de no existir dato, colocar cero (0).
	Debe decimal.Decimal `xml:"Debe,attr"`
	// Haber Atributo requerido para expresar el monto de los movimientos acreedores de la cuenta o subcuenta. En caso de no existir dato, colocar cero (0).
	Haber decimal.Decimal `xml:"Haber,attr"`
	// SaldoFin Atributo requerido para expresar el monto del saldo final de la cuenta o subcuenta en el periodo. De acuerdo a la naturaleza de la cuenta o subcuenta, deberá de corresponder el saldo final, de lo contrario se entenderá que es un saldo final de naturaleza inversa. En caso de no existir dato, colocar cero (0)
	SaldoFin decimal.Decimal `xml:"SaldoFin,attr"`
}

// Must match the pattern [NC]
type TipoEnvio string

const (
	// TipoEnvioN Normal.
	TipoEnvioN TipoEnvio = "N"
	// TipoEnvioC Complementaria.
	TipoEnvioC TipoEnvio = "C"
)
