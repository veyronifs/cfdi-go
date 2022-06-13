package cfdi40

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"

	"github.com/veyronifs/cfdi-go/curconv"
	"github.com/veyronifs/cfdi-go/encoder"
	"github.com/veyronifs/cfdi-go/types"
)

// Unmarshal parses the XML-encoded data and creates a Comprobante valuenc.
func Unmarshal(b []byte) (*Comprobante, error) {
	comprobante := Comprobante{}
	err := xml.Unmarshal(b, &comprobante)
	if err != nil {
		return nil, err
	}
	return &comprobante, nil
}

var cfdiXS = encoder.NSElem{
	Prefix: "cfdi",
	NS:     "http://www.sat.gob.mx/cfd/4",
}

type ComplementoNS interface {
	SchemaLocation() string
	XmlNSPrefix() string
	XmlNS() string
}

func Marshal(c *Comprobante) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(cfdiXS.ElemXS("Comprobante"))
	defer enc.EndElem("Comprobante")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	eschema := "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd"

	if c.Complemento != nil {
		/*
			if cp := c.Complemento.CartaPorte20; cp != nil {
				eschema += cp.SchemaLocation()
				enc.WriteAttrStr(cp.XmlNSPrefix(), cp.XmlNS())
			}
		*/

		// Usings reflection to get the namespaces instead of hardcoding them.
		val := reflect.ValueOf(c.Complemento).Elem()
		typeNS := reflect.TypeOf((*ComplementoNS)(nil)).Elem()

		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i)
			if f.IsNil() {
				continue
			}
			// check if implements ComplementoNS
			if f.Type().Implements(typeNS) {
				v := f.Interface().(ComplementoNS)
				eschema += " " + v.SchemaLocation()
				enc.WriteAttrStr("xmlns:"+v.XmlNSPrefix(), v.XmlNS())
			} else if f.Type().String() != "*tfd11.TimbreFiscalDigital" {
				return nil, fmt.Errorf("complemento %T does not implement ComplementoNS", f.Interface())
			}
		}
	}
	enc.WriteAttrStr("xsi:schemaLocation", eschema)

	encodeHeader(enc, c)
	encodeInformacionGlobal(enc, c)
	encodeCfdiRelacionadosAll(enc, c)
	encodeEmisor(enc, c.Emisor)
	encodeReceptor(enc, c.Receptor)
	encodeConceptos(enc, c)
	encodeImpuestos(enc, c)
	encodeComplemento(enc, c)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func encodeHeader(enc *encoder.Encoder, c *Comprobante) {
	moneda := string(c.Moneda)
	enc.WriteAttrStrZ("Version", c.Version)
	enc.WriteAttrStrZ("Serie", c.Serie)
	enc.WriteAttrStrZ("Folio", c.Folio)
	enc.WriteAttrStrZ("Fecha", c.Fecha.String())
	enc.WriteAttrStrZ("Sello", c.Sello)
	enc.WriteAttrStrZ("NoCertificado", c.NoCertificado)
	enc.WriteAttrStrZ("Certificado", c.Certificado)
	enc.WriteAttrStrZ("Moneda", moneda)
	enc.WriteAttrStrZ("TipoDeComprobante", string(c.TipoDeComprobante))
	enc.WriteAttrStrZ("LugarExpedicion", c.LugarExpedicion)
	enc.WriteAttrStrZ("FormaPago", string(c.FormaPago))
	enc.WriteAttrStrMaxEllipsisZ("CondicionesDePago", c.CondicionesDePago, 1000)
	enc.WriteAttrStrZ("MetodoPago", string(c.MetodoPago))
	enc.WriteAttrStrZ("Exportacion", string(c.Exportacion))
	enc.WriteAttrStrZ("Confirmacion", c.Confirmacion)
	if c.TipoDeComprobante == types.ComprobanteP {
		enc.WriteAttrStr("SubTotal", "0")
		enc.WriteAttrStr("Total", "0")
	} else {
		enc.WriteAttrStr("SubTotal", curconv.RoundFixed(c.SubTotal, c.Moneda))
		enc.WriteAttrStr("Total", curconv.RoundFixed(c.Total, c.Moneda))
		enc.WriteAttrDecimalCurrZ("Descuento", c.Descuento, moneda)
	}
	enc.WriteAttrDecimalZ("TipoCambio", c.TipoCambio, 6)
}

func encodeInformacionGlobal(enc *encoder.Encoder, c *Comprobante) {
	if c.InformacionGlobal == nil {
		return
	}

	enc.StartElem(cfdiXS.Elem("InformacionGlobal"))
	defer enc.EndElem("InformacionGlobal")

	enc.WriteAttrStr("Periodicidad", string(c.InformacionGlobal.Periodicidad))
	enc.WriteAttrStr("Meses", c.InformacionGlobal.Meses)
	enc.WriteAttrInt("Año", c.InformacionGlobal.Anio)
}

func encodeCfdiRelacionadosAll(enc *encoder.Encoder, c *Comprobante) {
	for _, rel := range c.CfdiRelacionados {
		encodeCfdiRelacionados(enc, rel)
	}
}

func encodeCfdiRelacionados(enc *encoder.Encoder, rel *CfdiRelacionados) {
	enc.StartElem(cfdiXS.Elem("CfdiRelacionados"))
	defer enc.EndElem("CfdiRelacionados")

	enc.WriteAttrStr("TipoRelacion", string(rel.TipoRelacion))

	for _, cfdiRel := range rel.CfdiRelacionado {
		enc.StartElem(cfdiXS.Elem("CfdiRelacionado"))
		enc.WriteAttrStr("UUID", cfdiRel.UUID)
		enc.EndElem("CfdiRelacionado")
	}
}

func encodeEmisor(enc *encoder.Encoder, c *Emisor) {
	if c == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Emisor"))
	defer enc.EndElem("Emisor")

	enc.WriteAttrStrZ("Rfc", c.Rfc)
	enc.WriteAttrStrMaxZ("Nombre", c.Nombre, 254)
	enc.WriteAttrStrZ("RegimenFiscal", string(c.RegimenFiscal))
	enc.WriteAttrStrZ("FacAtrAdquirente", c.FacAtrAdquirente)
}

func encodeReceptor(enc *encoder.Encoder, c *Receptor) {
	if c == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Receptor"))
	defer enc.EndElem("Receptor")

	enc.WriteAttrStrZ("Rfc", c.Rfc)
	enc.WriteAttrStrMaxZ("Nombre", c.Nombre, 254)
	enc.WriteAttrStrZ("DomicilioFiscalReceptor", c.DomicilioFiscalReceptor)
	enc.WriteAttrStrZ("ResidenciaFiscal", string(c.ResidenciaFiscal))
	enc.WriteAttrStrZ("NumRegIdTrib", c.NumRegIdTrib)
	enc.WriteAttrStrZ("RegimenFiscalReceptor", string(c.RegimenFiscalReceptor))
	enc.WriteAttrStrZ("UsoCFDI", string(c.UsoCFDI))
}
func encodeConceptos(enc *encoder.Encoder, c *Comprobante) {
	enc.StartElem(cfdiXS.Elem("Conceptos"))
	defer enc.EndElem("Conceptos")

	for _, concepto := range c.Conceptos {
		encodeConcepto(enc, concepto, string(c.Moneda))
	}
}

func encodeConcepto(enc *encoder.Encoder, concepto *Concepto, moneda string) {
	enc.StartElem(cfdiXS.Elem("Concepto"))
	defer enc.EndElem("Concepto")

	enc.WriteAttrStrZ("ClaveProdServ", concepto.ClaveProdServ)
	enc.WriteAttrStrMaxZ("NoIdentificacion", concepto.NoIdentificacion, 100)
	enc.WriteAttrStrZ("ClaveUnidad", concepto.ClaveUnidad)
	enc.WriteAttrStrMaxZ("Unidad", concepto.Unidad, 20)
	enc.WriteAttrStrMaxEllipsisZ("Descripcion", concepto.Descripcion, 1000)
	enc.WriteAttrStrZ("ObjetoImp", string(concepto.ObjetoImp))
	enc.WriteAttrDecimalCurr("ValorUnitario", concepto.ValorUnitario, moneda)
	enc.WriteAttrDecimal("Cantidad", concepto.Cantidad, 6)
	enc.WriteAttrDecimalCurr("Importe", concepto.Importe, moneda)
	enc.WriteAttrDecimalCurrZ("Descuento", concepto.Descuento, moneda)

	encodeConceptoImpuestos(enc, concepto.Impuestos, moneda)
	encodeConceptoACuentaTerceros(enc, concepto.ACuentaTerceros)
	encodeConceptoInformacionAduanera(enc, concepto.InformacionAduanera)
	for _, cPred := range concepto.CuentaPredial {
		enc.StartElem(cfdiXS.Elem("CuentaPredial"))
		enc.WriteAttrStrMax("Numero", cPred.Numero, 150)
		enc.EndElem("CuentaPredial")
	}
	encodeConceptoPartes(enc, concepto.Parte, moneda)
}
func encodeConceptoACuentaTerceros(enc *encoder.Encoder, at *ConceptoACuentaTerceros) {
	if at == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("ACuentaTerceros"))
	defer enc.EndElem("ACuentaTerceros")
	enc.WriteAttrStrZ("RfcACuentaTerceros", at.RfcACuentaTerceros)
	enc.WriteAttrStrMaxZ("NombreACuentaTerceros", at.NombreACuentaTerceros, 254)
	enc.WriteAttrStrZ("RegimenFiscalACuentaTerceros", string(at.RegimenFiscalACuentaTerceros))
	enc.WriteAttrStrZ("DomicilioFiscalACuentaTerceros", at.DomicilioFiscalACuentaTerceros)
}
func encodeConceptoImpuestos(enc *encoder.Encoder, impuestos *ConceptoImpuestos, moneda string) {
	if impuestos == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Impuestos"))
	defer enc.EndElem("Impuestos")
	if len(impuestos.Traslados) > 0 {
		enc.StartElem(cfdiXS.Elem("Traslados"))
		for _, impuesto := range impuestos.Traslados {
			encodeConceptoImpuestosTraslados(enc, impuesto, moneda)
		}
		enc.EndElem("Traslados")
	}
	if len(impuestos.Retenciones) > 0 {
		enc.StartElem(cfdiXS.Elem("Retenciones"))
		for _, impuesto := range impuestos.Retenciones {
			encodeConceptoImpuestosRetenciones(enc, impuesto, moneda)
		}
		enc.EndElem("Retenciones")
	}
}
func encodeConceptoImpuestosTraslados(enc *encoder.Encoder, impuestos *ConceptoImpuestosTraslado, moneda string) {
	enc.StartElem(cfdiXS.Elem("Traslado"))
	defer enc.EndElem("Traslado")

	enc.WriteAttrDecimalCurr("Base", impuestos.Base, moneda)
	enc.WriteAttrStrZ("Impuesto", string(impuestos.Impuesto))
	enc.WriteAttrStrZ("TipoFactor", string(impuestos.TipoFactor))
	if impuestos.TipoFactor != types.TipoFactorExento {
		enc.WriteAttrStr("TasaOCuota", impuestos.TasaOCuota.StringFixed(6))
		enc.WriteAttrDecimalCurr("Importe", impuestos.Importe, moneda)
	}
}
func encodeConceptoImpuestosRetenciones(enc *encoder.Encoder, impuestos *ConceptoImpuestosRetencion, moneda string) {
	enc.StartElem(cfdiXS.Elem("Retencion"))
	defer enc.EndElem("Retencion")

	enc.WriteAttrDecimalCurr("Base", impuestos.Base, moneda)
	enc.WriteAttrStrZ("Impuesto", string(impuestos.Impuesto))
	enc.WriteAttrStrZ("TipoFactor", string(impuestos.TipoFactor))
	enc.WriteAttrStr("TasaOCuota", impuestos.TasaOCuota.StringFixed(6))
	enc.WriteAttrDecimalCurr("Importe", impuestos.Importe, moneda)
}

func encodeConceptoInformacionAduanera(enc *encoder.Encoder, ci []*ConceptoInformacionAduanera) {
	for _, ia := range ci {
		enc.StartElem(cfdiXS.Elem("InformacionAduanera"))
		enc.WriteAttrStr("NumeroPedimento", ia.NumeroPedimento)
		enc.EndElem("InformacionAduanera")
	}
}
func encodeConceptoPartes(enc *encoder.Encoder, parte []*Parte, moneda string) {
	if len(parte) == 0 {
		return
	}
	for _, p := range parte {
		encodeConceptoParte(enc, p, moneda)
	}
}

func encodeConceptoParte(enc *encoder.Encoder, parte *Parte, moneda string) {
	enc.StartElem(cfdiXS.Elem("Parte"))
	defer enc.EndElem("Parte")
	enc.WriteAttrStrZ("ClaveProdServ", parte.ClaveProdServ)
	enc.WriteAttrStrMaxZ("NoIdentificacion", parte.NoIdentificacion, 100)
	enc.WriteAttrDecimalCurr("Cantidad", parte.Cantidad, moneda)
	enc.WriteAttrStrMaxZ("Unidad", parte.Unidad, 20)
	enc.WriteAttrStrMaxEllipsisZ("Descripcion", parte.Descripcion, 1000)
	enc.WriteAttrDecimalCurr("ValorUnitario", parte.ValorUnitario, moneda)
	enc.WriteAttrDecimalCurr("Importe", parte.Importe, moneda)
	encodeConceptoInformacionAduanera(enc, parte.InformacionAduanera)
}
func encodeImpuestos(enc *encoder.Encoder, c *Comprobante) {
	if c.Impuestos == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Impuestos"))
	defer enc.EndElem("Impuestos")
	if len(c.Impuestos.Retenciones) > 0 {
		enc.WriteAttrStr("TotalImpuestosRetenidos", curconv.RoundFixed(c.Impuestos.TotalImpuestosRetenidos, c.Moneda))
	}
	if len(c.Impuestos.Traslados) > 0 && c.Impuestos.WriteTotalImpuestosTrasladados() {
		enc.WriteAttrStr("TotalImpuestosTrasladados", curconv.RoundFixed(c.Impuestos.TotalImpuestosTrasladados, c.Moneda))
	}
	encodeImpuestosRetenciones(enc, c.Impuestos.Retenciones, string(c.Moneda))
	encodeImpuestosTraslados(enc, c.Impuestos.Traslados, string(c.Moneda))
}
func encodeImpuestosRetenciones(enc *encoder.Encoder, ret ImpuestosRetenciones, moneda string) {
	if len(ret) == 0 {
		return
	}
	enc.StartElem(cfdiXS.Elem("Retenciones"))
	defer enc.EndElem("Retenciones")
	for _, r := range ret {
		enc.StartElem(cfdiXS.Elem("Retencion"))
		enc.WriteAttrStr("Impuesto", string(r.Impuesto))
		enc.WriteAttrStr("Importe", curconv.RoundFixed(r.Importe, moneda))
		enc.EndElem("Retencion")
	}
}
func encodeImpuestosTraslados(enc *encoder.Encoder, tras ImpuestosTraslados, moneda string) {
	if len(tras) == 0 {
		return
	}
	enc.StartElem(cfdiXS.Elem("Traslados"))
	defer enc.EndElem("Traslados")
	for _, r := range tras {
		enc.StartElem(cfdiXS.Elem("Traslado"))
		enc.WriteAttrStr("Base", curconv.RoundFixed(r.Base, moneda))
		enc.WriteAttrStr("Impuesto", string(r.Impuesto))
		enc.WriteAttrStr("TipoFactor", string(r.TipoFactor))
		if r.TipoFactor != types.TipoFactorExento {
			enc.WriteAttrStr("TasaOCuota", r.TasaOCuota.StringFixed(6))
			enc.WriteAttrStr("Importe", curconv.RoundFixed(r.Importe, moneda))
		}

		enc.EndElem("Traslado")
	}
}

func encodeComplemento(enc *encoder.Encoder, c *Comprobante) {
	if c.Complemento == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Complemento"))
	defer enc.EndElem("Complemento")
	if c.Complemento.CartaPorte20 != nil {
		c.Complemento.CartaPorte20.MarshalComplemento(enc, string(c.Moneda))
	}
	if c.Complemento.Pagos20 != nil {
		c.Complemento.Pagos20.MarshalComplemento(enc)
	}

	if c.Complemento.CCE11 != nil {
		c.Complemento.CCE11.MarshalComplemento(enc)
	}

	if c.Complemento.TFD11 != nil {
		c.Complemento.TFD11.MarshalComplemento(enc)
	}

}
