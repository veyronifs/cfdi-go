package cfdi40

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

var (
	ErrInvalid  = errors.New("dato invalido")
	ErrRequired = errors.New("dato requerido")
)

// NewComprobantePago Crea la base de un comprobante de pago.
func NewComprobantePago() *Comprobante {
	return &Comprobante{
		Version:           Version,
		Fecha:             types.NewFechaHNow(),
		Moneda:            types.MonedaXXX,
		Total:             decimal.Zero,
		SubTotal:          decimal.Zero,
		TipoDeComprobante: types.ComprobanteP,
		Exportacion:       types.Exportacion01,
		Conceptos: Conceptos{
			{
				ClaveProdServ: "84111506",
				Cantidad:      decimal.NewFromFloat(1),
				ClaveUnidad:   "ACT",
				Descripcion:   "Pago",
				ValorUnitario: decimal.NewFromFloat(0),
				Importe:       decimal.NewFromFloat(0),
				ObjetoImp:     types.ObjetoImp01,
			},
		},
	}
}

// NewComprobante crea la base de un comprobante 4.0.
func NewComprobante() *Comprobante {
	return &Comprobante{
		Version: Version,
	}
}

// CalcularTotales calcula el total, subtotal y descuento del comprobante de acuerdo a los conceptos y sus impuestos.
func CalcularTotales(c Comprobante) (subTotal, descuento, total decimal.Decimal) {
	subTotal, descuento, total = decimal.Zero, decimal.Zero, decimal.Zero

	for _, concept := range c.Conceptos {
		subTotal = subTotal.Add(concept.Importe)
		descuento = descuento.Add(concept.Descuento)
		total = total.Add(concept.Importe).Sub(concept.Descuento)
		if concept.Impuestos != nil {
			for _, tras := range concept.Impuestos.Traslados {
				total = total.Add(tras.Importe)
			}
			for _, ret := range concept.Impuestos.Retenciones {
				total = total.Sub(ret.Importe)
			}
		}
	}

	return subTotal, descuento, total
}

// NewImpuestos regresa el resumen de impuestos de acuerdo a la suma de los impuestos de los conceptos.
func NewImpuestos(c Comprobante) *Impuestos {
	ok := false
	impuesto := &Impuestos{}
	for _, concept := range c.Conceptos {
		if concept.Impuestos == nil {
			continue
		}
		if concept.Impuestos.Retenciones != nil {
			for _, ret := range concept.Impuestos.Retenciones {
				ok = true
				impuesto.AddRetencion(ret.Impuesto, ret.Importe)
			}
		}
		if concept.Impuestos.Traslados != nil {
			for _, tras := range concept.Impuestos.Traslados {
				ok = true
				impuesto.AddTraslado(tras.Base, tras.Impuesto, tras.TipoFactor, tras.TasaOCuota, tras.Importe)
			}
		}
	}

	if ok {
		return impuesto
	}
	return nil
}

func NewEmisor(
	Rfc string,
	Nombre string,
	RegimenFiscal types.RegimenFiscal,
	FacAtrAdquirente string,
) (*Emisor, error) {
	if Rfc == "" {
		return nil, fmt.Errorf("Emisor.Rfc %w", ErrRequired)
	}
	if Nombre == "" {
		return nil, fmt.Errorf("Emisor.Nombre %w", ErrRequired)
	}
	if RegimenFiscal == "" {
		return nil, fmt.Errorf("Emisor.RegimenFiscal %w", ErrRequired)
	}

	return &Emisor{
		Rfc:              Rfc,
		Nombre:           Nombre,
		RegimenFiscal:    RegimenFiscal,
		FacAtrAdquirente: FacAtrAdquirente,
	}, nil
}

func NewReceptor(
	Rfc string,
	Nombre string,
	DomicilioFiscalReceptor string,
	RegimenFiscalReceptor types.RegimenFiscal,
	UsoCFDI types.UsoCFDI,
) (*Receptor, error) {
	switch Rfc {
	case "":
		return nil, fmt.Errorf("Receptor.Rfc %w", ErrRequired)
	case RFCPublico, RFCExtranjero:
		return nil, fmt.Errorf("Receptor.Rfc %w", ErrInvalid)
	}

	if Nombre == "" {
		return nil, fmt.Errorf("Receptor.Nombre %w", ErrRequired)
	}
	if DomicilioFiscalReceptor == "" {
		return nil, fmt.Errorf("Receptor.DomicilioFiscalReceptor %w", ErrRequired)
	}
	if RegimenFiscalReceptor == "" {
		return nil, fmt.Errorf("Receptor.RegimenFiscalReceptor %w", ErrRequired)
	}
	if UsoCFDI == "" {
		return nil, fmt.Errorf("Receptor.UsoCFDI %w", ErrRequired)
	}
	return &Receptor{
		Rfc:                     Rfc,
		Nombre:                  Nombre,
		DomicilioFiscalReceptor: DomicilioFiscalReceptor,
		RegimenFiscalReceptor:   RegimenFiscalReceptor,
		UsoCFDI:                 UsoCFDI,
	}, nil
}

// NewReceptorPublico genera el receptor para publico en general.
func NewReceptorPublico(LugarExpedicion string) (*Receptor, error) {
	if LugarExpedicion == "" {
		return nil, fmt.Errorf("Receptor.DomicilioFiscalReceptor (LugarExpedicion) %w", ErrRequired)
	}
	return &Receptor{
		Rfc:                     RFCPublico,
		Nombre:                  "PUBLICO EN GENERAL",
		DomicilioFiscalReceptor: LugarExpedicion,
		ResidenciaFiscal:        "",
		NumRegIdTrib:            "",
		RegimenFiscalReceptor:   types.RegimenFiscal616,
		UsoCFDI:                 types.UsoCFDIG03,
	}, nil
}

// NewReceptorExt genera el receptor extrajero.
func NewReceptorExt(
	Nombre string,
	ResidenciaFiscal types.Pais,
	NumRegIdTrib string,
	UsoCFDI types.UsoCFDI,
	LugarExpedicion string,
) (*Receptor, error) {
	if Nombre == "" {
		return nil, fmt.Errorf("Receptor.Nombre %w", ErrRequired)
	}

	if ResidenciaFiscal == "" {
		return nil, fmt.Errorf("Receptor.ResidenciaFiscal %w", ErrRequired)
	} else if ResidenciaFiscal == types.PaisMEX {
		return nil, fmt.Errorf("Receptor.ResidenciaFiscal %w", ErrInvalid)
	}

	if NumRegIdTrib == "" {
		return nil, fmt.Errorf("Receptor.NumRegIdTrib %w", ErrRequired)
	}

	if UsoCFDI == "" {
		return nil, fmt.Errorf("Receptor.UsoCFDI %w", ErrRequired)
	}

	if LugarExpedicion == "" {
		return nil, fmt.Errorf("Receptor.DomicilioFiscalReceptor (LugarExpedicion) %w", ErrRequired)
	}

	return &Receptor{
		Rfc:                     "XEXX010101000",
		Nombre:                  Nombre,
		DomicilioFiscalReceptor: LugarExpedicion,
		ResidenciaFiscal:        ResidenciaFiscal,
		NumRegIdTrib:            NumRegIdTrib,
		RegimenFiscalReceptor:   types.RegimenFiscal616,
		UsoCFDI:                 UsoCFDI,
	}, nil
}

// NewInformacionGlobal genera la informacion global cuando el rfc es de publico en general.
func NewInformacionGlobal(rfc string, periodicidad types.Periodicidad, fecha types.FechaH) (*InformacionGlobal, error) {
	if rfc != RFCPublico {
		return nil, nil
	}
	if periodicidad == "" {
		return nil, fmt.Errorf("InformacionGlobal.Peridiocidad %w", ErrRequired)
	}
	time := time.Time(fecha)
	if time.IsZero() {
		return nil, fmt.Errorf("InformacionGlobal.Fecha %w", ErrRequired)
	}

	// Si la Periodicidad es distinto a "05" Bimestral, Meses debe contener 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11 o 12.
	if periodicidad != types.PeriodicidadBimestral {
		return &InformacionGlobal{
			Periodicidad: periodicidad,
			Meses:        time.Format("01"),
			Anio:         time.Year(),
		}, nil
	}

	// Si la Periodicidad es "05" Bimestral, Meses debe contener 13, 14, 15, 16, 17 o 18.
	info := InformacionGlobal{
		Periodicidad: periodicidad,
		Meses:        strconv.Itoa(int((time.Month()-1)/2 + 13)),
		Anio:         time.Year(),
	}

	return &info, nil
}
