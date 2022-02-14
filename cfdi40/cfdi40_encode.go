package cfdi40

import (
	"bytes"
	"encoding/xml"

	"github.com/veyronifs/cfdi-go/encoder"
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

func Marshal(c Comprobante) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(cfdiXS.ElemXS("Comprobante"))
	defer enc.EndElem("Comprobante")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	eschema := "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd"

	if c.Complemento != nil {
		if cp := c.Complemento.CartaPorte20; cp != nil {
			eschema += cp.SchemaLocation()
			enc.WriteAttrStr(cp.XmlNSPrefix(), cp.XmlNS())
		}
	}
	enc.WriteAttrStr("xsi:schemaLocation", eschema)

	encodeHeader(enc, c)
	encodeInformacionGlobal(enc, c)
	encodeCfdiRelacionadosAll(enc, c)
	encodeEmisor(enc, c)
	encodeReceptor(enc, c)
	encodeConceptos(enc, c)
	encodeImpuestos(enc, c)
	encodeComplemento(enc, c)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func encodeHeader(enc *encoder.Encoder, c Comprobante) {
	moneda := string(c.Moneda)
	enc.WriteAttrStrZ("Version", c.Version)
	enc.WriteAttrStrZ("Serie", c.Serie)
	enc.WriteAttrStrZ("Folio", c.Folio)
	enc.WriteAttrStrZ("Fecha", c.Fecha.Encode())
	enc.WriteAttrStrZ("Sello", c.Sello)
	enc.WriteAttrStrZ("NoCertificado", c.NoCertificado)
	enc.WriteAttrStrZ("Certificado", c.Certificado)
	enc.WriteAttrStrZ("Moneda", moneda)
	enc.WriteAttrStrZ("TipoDeComprobante", string(c.TipoDeComprobante))
	enc.WriteAttrStrZ("LugarExpedicion", c.LugarExpedicion)
	enc.WriteAttrStrZ("FormaPago", string(c.FormaPago))
	enc.WriteAttrStrZ("CondicionesDePago", c.CondicionesDePago)
	enc.WriteAttrStrZ("MetodoPago", string(c.MetodoPago))
	enc.WriteAttrStrZ("Exportacion", string(c.Exportacion))
	enc.WriteAttrStrZ("Confirmacion", c.Confirmacion)
	enc.WriteAttrDecimalCurr("SubTotal", c.SubTotal, moneda)
	enc.WriteAttrDecimalCurr("Total", c.Total, moneda)
	enc.WriteAttrDecimalCurr("Descuento", c.Descuento, moneda)
	enc.WriteAttrDecimalZ("TipoCambio", c.TipoCambio, 6)
}

func encodeInformacionGlobal(enc *encoder.Encoder, c Comprobante) {
	enc.StartElem(cfdiXS.Elem("InformacionGlobal"))
	defer enc.EndElem("InformacionGlobal")

	enc.WriteAttrStr("Periodicidad", c.InformacionGlobal.Periodicidad)
	enc.WriteAttrStr("Meses", c.InformacionGlobal.Meses)
	enc.WriteAttrInt("Año", c.InformacionGlobal.Anio)
}

func encodeCfdiRelacionadosAll(enc *encoder.Encoder, c Comprobante) {
	for _, rel := range c.CfdiRelacionados {
		encodeCfdiRelacionados(enc, rel)
	}
}

func encodeCfdiRelacionados(enc *encoder.Encoder, rel CfdiRelacionados) {
	enc.StartElem(cfdiXS.Elem("CfdiRelacionados"))
	defer enc.EndElem("CfdiRelacionados")

	enc.WriteAttrStr("TipoRelacion", string(rel.TipoRelacion))

	for _, cfdiRel := range rel.CfdiRelacionado {
		enc.StartElem(cfdiXS.Elem("CfdiRelacionado"))
		enc.WriteAttrStr("UUID", cfdiRel.UUID)
		enc.EndElem("CfdiRelacionado")
	}
}

func encodeEmisor(enc *encoder.Encoder, c Comprobante) {
	enc.StartElem(cfdiXS.Elem("Emisor"))
	defer enc.EndElem("Emisor")

	enc.WriteAttrStrZ("Rfc", c.Emisor.Rfc)
	enc.WriteAttrStrZ("Nombre", c.Emisor.Nombre)
	enc.WriteAttrStrZ("RegimenFiscal", string(c.Emisor.RegimenFiscal))
	enc.WriteAttrStrZ("FacAtrAdquirente", c.Emisor.FacAtrAdquirente)
}

func encodeReceptor(enc *encoder.Encoder, c Comprobante) {
	enc.StartElem(cfdiXS.Elem("Receptor"))
	defer enc.EndElem("Receptor")

	enc.WriteAttrStrZ("Rfc", c.Receptor.Rfc)
	enc.WriteAttrStrZ("Nombre", c.Receptor.Nombre)
	enc.WriteAttrStrZ("DomicilioFiscalReceptor", c.Receptor.DomicilioFiscalReceptor)
	enc.WriteAttrStrZ("ResidenciaFiscal", string(c.Receptor.ResidenciaFiscal))
	enc.WriteAttrStrZ("NumRegIdTrib", c.Receptor.NumRegIdTrib)
	enc.WriteAttrStrZ("RegimenFiscalReceptor", string(c.Receptor.RegimenFiscalReceptor))
	enc.WriteAttrStrZ("UsoCFDI", string(c.Receptor.UsoCFDI))
}
func encodeConceptos(enc *encoder.Encoder, c Comprobante) {
	enc.StartElem(cfdiXS.Elem("Conceptos"))
	defer enc.EndElem("Conceptos")

	for _, concepto := range c.Conceptos {
		encodeConcepto(enc, concepto, string(c.Moneda))
	}
}

func encodeConcepto(enc *encoder.Encoder, concepto Concepto, moneda string) {
	enc.StartElem(cfdiXS.Elem("Concepto"))
	defer enc.EndElem("Concepto")

	enc.WriteAttrStrZ("ClaveProdServ", concepto.ClaveProdServ)
	enc.WriteAttrStrZ("NoIdentificacion", concepto.NoIdentificacion)
	enc.WriteAttrStrZ("ClaveUnidad", concepto.ClaveUnidad)
	enc.WriteAttrStrZ("Unidad", concepto.Unidad)
	enc.WriteAttrStrZ("Descripcion", concepto.Descripcion)
	enc.WriteAttrStrZ("ObjetoImp", string(concepto.ObjetoImp))
	enc.WriteAttrDecimalCurr("ValorUnitario", concepto.ValorUnitario, moneda)
	enc.WriteAttrDecimal("Cantidad", concepto.Cantidad, 6)
	enc.WriteAttrDecimalCurr("Importe", concepto.Importe, moneda)
	enc.WriteAttrDecimalCurrZ("Descuento", concepto.Descuento, moneda)

	encodeConceptoImpuestos(enc, concepto.Impuestos, moneda)
	encodeConceptoACuentaTerceros(enc, concepto.ACuentaTerceros)
	for _, ia := range concepto.InformacionAduanera {
		enc.StartElem(cfdiXS.Elem("InformacionAduanera"))
		enc.WriteAttrStr("NumeroPedimento", ia.NumeroPedimento)
		enc.EndElem("InformacionAduanera")
	}
	for _, cPred := range concepto.CuentaPredial {
		enc.StartElem(cfdiXS.Elem("CuentaPredial"))
		enc.WriteAttrStr("Numero", cPred.Numero)
		enc.EndElem("CuentaPredial")
	}
}
func encodeConceptoACuentaTerceros(enc *encoder.Encoder, at *ConceptoACuentaTerceros) {
	if at == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("ACuentaTerceros"))
	defer enc.EndElem("ACuentaTerceros")
	enc.WriteAttrStrZ("RfcACuentaTerceros", at.RfcACuentaTerceros)
	enc.WriteAttrStrZ("NombreACuentaTerceros", at.NombreACuentaTerceros)
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
	if len(impuestos.Traslados) > 0 {
		enc.StartElem(cfdiXS.Elem("Retenciones"))
		for _, impuesto := range impuestos.Retenciones {
			encodeConceptoImpuestosRetenciones(enc, impuesto, moneda)
		}
		enc.EndElem("Retenciones")
	}
}
func encodeConceptoImpuestosTraslados(enc *encoder.Encoder, impuestos ConceptoImpuestosTraslado, moneda string) {
	enc.StartElem(cfdiXS.Elem("Traslado"))
	defer enc.EndElem("Traslado")

	enc.WriteAttrDecimalCurr("Base", impuestos.Base, moneda)
	enc.WriteAttrStrZ("Impuesto", string(impuestos.Impuesto))
	enc.WriteAttrStrZ("TipoFactor", string(impuestos.TipoFactor))
	enc.WriteAttrDecimal("TasaOCuota", impuestos.TasaOCuota, 6)
	enc.WriteAttrDecimalCurr("Importe", impuestos.Importe, moneda)
}
func encodeConceptoImpuestosRetenciones(enc *encoder.Encoder, impuestos ConceptoImpuestosRetencion, moneda string) {
	enc.StartElem(cfdiXS.Elem("Retencion"))
	defer enc.EndElem("Retencion")

	enc.WriteAttrDecimalCurr("Base", impuestos.Base, moneda)
	enc.WriteAttrStrZ("Impuesto", string(impuestos.Impuesto))
	enc.WriteAttrStrZ("TipoFactor", string(impuestos.TipoFactor))
	enc.WriteAttrDecimal("TasaOCuota", impuestos.TasaOCuota, 6)
	enc.WriteAttrDecimalCurr("Importe", impuestos.Importe, moneda)
}

func encodeImpuestos(enc *encoder.Encoder, c Comprobante) {
	if c.Impuestos == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Impuestos"))
	defer enc.EndElem("Impuestos")
	if len(c.Impuestos.Retenciones) > 0 {
		enc.WriteAttrDecimalCurr("TotalImpuestosRetenidos", c.Impuestos.TotalImpuestosRetenidos, string(c.Moneda))
	}
	if len(c.Impuestos.Traslados) > 0 {
		enc.WriteAttrDecimalCurr("TotalImpuestosTrasladados", c.Impuestos.TotalImpuestosTrasladados, string(c.Moneda))
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
		enc.WriteAttrDecimalCurr("Importe", r.Importe, moneda)
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
		enc.WriteAttrDecimalCurr("Base", r.Importe, moneda)
		enc.WriteAttrStr("Impuesto", string(r.Impuesto))
		enc.WriteAttrStr("TipoFactor", string(r.TipoFactor))
		enc.WriteAttrDecimal("TasaOCuota", r.TasaOCuota, 6)
		enc.WriteAttrDecimalCurr("Importe", r.Importe, moneda)
		enc.EndElem("Traslado")
	}
}

func encodeComplemento(enc *encoder.Encoder, c Comprobante) {
	if c.Complemento == nil {
		return
	}
	enc.StartElem(cfdiXS.Elem("Complemento"))
	defer enc.EndElem("Complemento")
	if c.Complemento.CartaPorte20 != nil {
		c.Complemento.CartaPorte20.EncodeComplemento(enc, string(c.Moneda))
	}
}
