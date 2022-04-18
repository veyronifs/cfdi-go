package pagos20

import (
	"bytes"

	"github.com/veyronifs/cfdi-go/encoder"
	"github.com/veyronifs/cfdi-go/types"
)

var pagos20XS = encoder.NSElem{
	Prefix: "pago20",
	NS:     "http://www.sat.gob.mx/Pagos20",
}

func (pagos *Pagos) SchemaLocation() string {
	return pagos20XS.NS + " http://www.sat.gob.mx/sitio_internet/cfd/Pagos/Pagos20.xsd"
}

func (pagos *Pagos) XmlNSPrefix() string {
	return pagos20XS.Prefix
}

func (pagos *Pagos) XmlNS() string {
	return pagos20XS.NS
}

func Marshal(cp *Pagos) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	cp.MarshalComplemento(enc)
	enc.EndAllFlush()
	return b.Bytes(), enc.GetError()
}

func (pagos *Pagos) MarshalComplemento(enc *encoder.Encoder) {
	encodePagos(enc, pagos)
}

func encodePagos(enc *encoder.Encoder, pagos *Pagos) {
	if pagos == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("Pagos"))
	defer enc.EndElem("Pagos")

	enc.WriteAttrStr("Version", pagos.Version)

	encodeTotales(enc, pagos.Totales)
	for _, pago := range pagos.Pago {
		encodePago(enc, pago)
	}
}

func encodeTotales(enc *encoder.Encoder, v *Totales) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("Totales"))
	defer enc.EndElem("Totales")
	enc.WriteAttrNullDecimal("TotalRetencionesIVA", v.TotalRetencionesIVA, 2)
	enc.WriteAttrNullDecimal("TotalRetencionesISR", v.TotalRetencionesISR, 2)
	enc.WriteAttrNullDecimal("TotalRetencionesIEPS", v.TotalRetencionesIEPS, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosBaseIVA16", v.TotalTrasladosBaseIVA16, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosImpuestoIVA16", v.TotalTrasladosImpuestoIVA16, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosBaseIVA8", v.TotalTrasladosBaseIVA8, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosImpuestoIVA8", v.TotalTrasladosImpuestoIVA8, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosBaseIVA0", v.TotalTrasladosBaseIVA0, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosImpuestoIVA0", v.TotalTrasladosImpuestoIVA0, 2)
	enc.WriteAttrNullDecimal("TotalTrasladosBaseIVAExento", v.TotalTrasladosBaseIVAExento, 2)
	enc.WriteAttrDecimal("MontoTotalPagos", v.MontoTotalPagos, 2)

}

func encodePago(enc *encoder.Encoder, v *Pago) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("Pago"))
	defer enc.EndElem("Pago")

	enc.WriteAttrStrZ("FechaPago", v.FechaPago.String())
	enc.WriteAttrStrZ("FormaDePagoP", string(v.FormaDePagoP))
	enc.WriteAttrStrZ("MonedaP", string(v.MonedaP))
	enc.WriteAttrDecimalCurr("TipoCambioP", v.TipoCambioP, string(v.MonedaP))
	enc.WriteAttrDecimalCurr("Monto", v.Monto, string(v.MonedaP))
	enc.WriteAttrStrZ("NumOperacion", v.NumOperacion)
	enc.WriteAttrStrZ("RfcEmisorCtaOrd", v.RfcEmisorCtaOrd)
	enc.WriteAttrStrZ("NomBancoOrdExt", v.NomBancoOrdExt)
	enc.WriteAttrStrZ("CtaOrdenante", v.CtaOrdenante)
	enc.WriteAttrStrZ("RfcEmisorCtaBen", v.RfcEmisorCtaBen)
	enc.WriteAttrStrZ("CtaBeneficiario", v.CtaBeneficiario)
	enc.WriteAttrStrZ("TipoCadPago", string(v.TipoCadPago))
	enc.WriteAttrStrZ("CertPago", v.CertPago)
	enc.WriteAttrStrZ("CadPago", v.CadPago)
	enc.WriteAttrStrZ("SelloPago", v.SelloPago)

	for _, rel := range v.DoctoRelacionado {
		encodeDoctoRelacionado(enc, rel, string(v.MonedaP))
	}

	encodeImpuestosP(enc, v.ImpuestosP, string(v.MonedaP))
}

func encodeDoctoRelacionado(enc *encoder.Encoder, v *DoctoRelacionado, monedaPago string) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("DoctoRelacionado"))
	defer enc.EndElem("DoctoRelacionado")

	monedaDR := string(v.MonedaDR)
	if monedaDR == "" {
		monedaDR = monedaPago
	}

	enc.WriteAttrStrZ("IdDocumento", v.IdDocumento)
	enc.WriteAttrStrZ("Serie", v.Serie)
	enc.WriteAttrStrZ("Folio", v.Folio)
	enc.WriteAttrStrZ("MonedaDR", monedaDR)
	enc.WriteAttrDecimalZ("EquivalenciaDR", v.EquivalenciaDR, 6)
	enc.WriteAttrIntZ("NumParcialidad", v.NumParcialidad)
	enc.WriteAttrDecimalCurr("ImpSaldoAnt", v.ImpSaldoAnt, monedaDR)
	enc.WriteAttrDecimalCurr("ImpPagado", v.ImpPagado, monedaDR)
	enc.WriteAttrDecimalCurr("ImpSaldoInsoluto", v.ImpSaldoInsoluto, monedaDR)
	enc.WriteAttrStrZ("ObjetoImpDR", string(v.ObjetoImpDR))

	encodeImpuestosDR(enc, v.ImpuestosDR, monedaDR)
}

func encodeImpuestosDR(enc *encoder.Encoder, v *ImpuestosDR, monedaDR string) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("ImpuestosDR"))
	defer enc.EndElem("ImpuestosDR")
	if len(v.RetencionesDR) > 0 {
		enc.StartElem(pagos20XS.Elem("RetencionesDR"))
		for _, imp := range v.RetencionesDR {
			encodeRetencionDR(enc, imp, monedaDR)
		}
		enc.EndElem("RetencionesDR")
	}

	if len(v.TrasladosDR) > 0 {
		enc.StartElem(pagos20XS.Elem("TrasladosDR"))
		for _, imp := range v.TrasladosDR {
			encodeTrasladoDR(enc, imp, monedaDR)
		}
		enc.EndElem("TrasladosDR")
	}
}

func encodeRetencionDR(enc *encoder.Encoder, v *RetencionDR, monedaDR string) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("RetencionDR"))
	defer enc.EndElem("RetencionDR")

	enc.WriteAttrDecimalCurr("BaseDR", v.BaseDR, monedaDR)
	enc.WriteAttrStrZ("ImpuestoDR", string(v.ImpuestoDR))
	enc.WriteAttrStrZ("TipoFactorDR", string(v.TipoFactorDR))
	enc.WriteAttrStr("TasaOCuotaDR", v.TasaOCuotaDR.StringFixed(6))
	enc.WriteAttrDecimalCurr("ImporteDR", v.ImporteDR, monedaDR)
}

func encodeTrasladoDR(enc *encoder.Encoder, v *TrasladoDR, monedaDR string) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("TrasladoDR"))
	defer enc.EndElem("TrasladoDR")

	enc.WriteAttrDecimalCurr("BaseDR", v.BaseDR, monedaDR)
	enc.WriteAttrStrZ("ImpuestoDR", string(v.ImpuestoDR))
	enc.WriteAttrStrZ("TipoFactorDR", string(v.TipoFactorDR))
	if v.TipoFactorDR != types.TipoFactorExento {
		enc.WriteAttrStr("TasaOCuotaDR", v.TasaOCuotaDR.StringFixed(6))
		enc.WriteAttrDecimalCurr("ImporteDR", v.ImporteDR, monedaDR)
	}
}

func encodeImpuestosP(enc *encoder.Encoder, v *ImpuestosP, monedaP string) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("ImpuestosP"))
	defer enc.EndElem("ImpuestosP")

	if len(v.RetencionesP) > 0 {
		enc.StartElem(pagos20XS.Elem("RetencionesP"))
		for _, imp := range v.RetencionesP {
			encodeRetencionP(enc, imp, monedaP)
		}
		enc.EndElem("RetencionesP")
	}

	if len(v.TrasladosP) > 0 {
		enc.StartElem(pagos20XS.Elem("TrasladosP"))
		for _, imp := range v.TrasladosP {
			encodeTrasladoP(enc, imp, monedaP)
		}
		enc.EndElem("TrasladosP")
	}
}

func encodeRetencionP(enc *encoder.Encoder, v *RetencionP, monedaP string) {
	if v == nil {
		return
	}

	enc.StartElem(pagos20XS.Elem("RetencionP"))
	defer enc.EndElem("RetencionP")

	enc.WriteAttrStrZ("ImpuestoP", string(v.ImpuestoP))
	enc.WriteAttrDecimalCurr("ImporteP", v.ImporteP, monedaP)

}

func encodeTrasladoP(enc *encoder.Encoder, v *TrasladoP, monedaP string) {
	if v == nil {
		return
	}
	enc.StartElem(pagos20XS.Elem("TrasladoP"))
	defer enc.EndElem("TrasladoP")

	enc.WriteAttrDecimalCurr("BaseP", v.BaseP, monedaP)
	enc.WriteAttrStrZ("ImpuestoP", string(v.ImpuestoP))
	enc.WriteAttrStrZ("TipoFactorP", string(v.TipoFactorP))
	if v.TipoFactorP != types.TipoFactorExento {
		enc.WriteAttrStr("TasaOCuotaP", v.TasaOCuotaP.StringFixed(6))
		enc.WriteAttrDecimalCurr("ImporteP", v.ImporteP, monedaP)
	}
}
