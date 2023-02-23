package polizasperiodo

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/veyronifs/cfdi-go/encoder"
)

// Unmarshal parses the XML-encoded data and returns the *Pólizas.
func Unmarshal(data []byte) (*Polizas, error) {
	var polizas Polizas
	err := xml.Unmarshal(data, &polizas)
	if err != nil {
		return nil, err
	}
	return &polizas, nil
}

var polXS = encoder.NSElem{
	Prefix: "PLZ",
	NS:     "http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo",
}

func Marshal(c *Polizas) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(polXS.ElemXS("Polizas"))
	defer enc.EndElem("Polizas")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	enc.WriteAttrStr("xsi:schemaLocation", polXS.NS+" http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/PolizasPeriodo/PolizasPeriodo_1_3.xsd")

	encodePolizas(enc, c)
	enc.EndAllFlush()
	return b.Bytes(), nil
}

func encodePolizas(enc *encoder.Encoder, c *Polizas) {
	enc.WriteAttrStrZ("Version", c.Version)
	enc.WriteAttrStrZ("RFC", c.RFC)
	enc.WriteAttrStrZ("Mes", fmt.Sprintf("%02d", c.Mes))
	enc.WriteAttrIntZ("Anio", c.Anio)
	enc.WriteAttrStrZ("TipoSolicitud", string(c.TipoSolicitud))
	enc.WriteAttrStrZ("NumOrden", c.NumOrden)
	enc.WriteAttrStrZ("NumTramite", c.NumTramite)
	enc.WriteAttrStrZ("Sello", c.Sello)
	enc.WriteAttrStrZ("noCertificado", c.NoCertificado)
	enc.WriteAttrStrZ("Certificado", c.Certificado)
	for _, poliza := range c.Polizas {
		enc.StartElem(polXS.Elem("Poliza"))
		enc.WriteAttrStrZ("NumUnIdenPol", poliza.NumUnIdenPol)
		enc.WriteAttrStrZ("Fecha", poliza.Fecha.String())
		enc.WriteAttrStrZ("Concepto", poliza.Concepto)
		encodeTransacciones(enc, poliza.Transaccion)
		enc.EndElem("Poliza")
	}
}

func encodeTransacciones(enc *encoder.Encoder, t []*Transaccion) {
	for _, transaccion := range t {
		enc.StartElem(polXS.Elem("Transaccion"))
		enc.WriteAttrStrZ("NumCta", transaccion.NumCta)
		enc.WriteAttrStrZ("DesCta", transaccion.DesCta)
		enc.WriteAttrStrZ("Concepto", transaccion.Concepto)
		enc.WriteAttrDecimal("Debe", transaccion.Debe, 2)
		enc.WriteAttrDecimal("Haber", transaccion.Haber, 2)
		encondeCompNals(enc, transaccion.CompNal)
		encondeCompNalOtrs(enc, transaccion.CompNalOtr)
		encondeCompExts(enc, transaccion.CompExt)
		encondeCheques(enc, transaccion.Cheque)
		encodeTransferencias(enc, transaccion.Transferencia)
		encodeOtrMetodoPagos(enc, transaccion.OtrMetodoPago)
		enc.EndElem("Transaccion")
	}
}

func encondeCompNals(enc *encoder.Encoder, ctas []*CompNal) {
	for _, cnal := range ctas {
		enc.StartElem(polXS.Elem("CompNal"))
		enc.WriteAttrStrZ("UUID_CFDI", cnal.UUIDCFDI)
		enc.WriteAttrStrZ("RFC", cnal.RFC)
		enc.WriteAttrDecimal("MontoTotal", cnal.MontoTotal, 2)
		enc.WriteAttrStrZ("Moneda", string(cnal.Moneda))
		enc.WriteAttrDecimal("TipCamb", cnal.TipCamb, 5)
		enc.EndElem("CompNal")
	}
}

func encondeCompNalOtrs(enc *encoder.Encoder, ctas []*CompNalOtr) {
	for _, cnalotr := range ctas {
		enc.StartElem(polXS.Elem("CompNalOtr"))
		enc.WriteAttrStrZ("CFD_CBB_Serie", cnalotr.CFDCBBSerie)
		enc.WriteAttrStrZ("CFD_CBB_NumFol", fmt.Sprintf("%d", cnalotr.CFDCBBNumFol))
		enc.WriteAttrStrZ("RFC", cnalotr.RFC)
		enc.WriteAttrDecimal("MontoTotal", cnalotr.MontoTotal, 2)
		enc.WriteAttrStrZ("Moneda", string(cnalotr.Moneda))
		enc.WriteAttrDecimal("TipCamb", cnalotr.TipCamb, 5)
		enc.EndElem("CompNalOtr")
	}
}

func encondeCompExts(enc *encoder.Encoder, ctas []*CompExt) {
	for _, cext := range ctas {
		enc.StartElem(polXS.Elem("CompExt"))
		enc.WriteAttrStrZ("NumFactExt", cext.NumFactExt)
		enc.WriteAttrStrZ("TaxID", cext.TaxID)
		enc.WriteAttrDecimal("MontoTotal", cext.MontoTotal, 2)
		enc.WriteAttrStrZ("Moneda", string(cext.Moneda))
		enc.WriteAttrDecimal("TipCamb", cext.TipCamb, 5)
		enc.EndElem("CompExt")
	}
}

func encondeCheques(enc *encoder.Encoder, ctas []*Cheque) {
	for _, cheque := range ctas {
		enc.StartElem(polXS.Elem("Cheque"))
		enc.WriteAttrStrZ("Num", cheque.Num)
		enc.WriteAttrStrZ("BanEmisNal", cheque.BanEmisNal)
		enc.WriteAttrStrZ("BanEmisExt", cheque.BanEmisExt)
		enc.WriteAttrStrZ("CtaOri", cheque.CtaOri)
		enc.WriteAttrStrZ("Fecha", cheque.Fecha.String())
		enc.WriteAttrStrZ("Benef", cheque.Benef)
		enc.WriteAttrStrZ("RFC", cheque.RFC)
		enc.WriteAttrDecimal("Monto", cheque.Monto, 2)
		enc.WriteAttrStrZ("Moneda", string(cheque.Moneda))
		enc.WriteAttrDecimal("TipCamb", cheque.TipCamb, 5)
		enc.EndElem("Cheque")
	}
}

func encodeTransferencias(enc *encoder.Encoder, ctas []*Transferencia) {
	for _, transf := range ctas {
		enc.StartElem(polXS.Elem("Transferencia"))
		enc.WriteAttrStrZ("CtaOri", transf.CtaOri)
		enc.WriteAttrStrZ("BancoOriNal", transf.BancoOriNal)
		enc.WriteAttrStrZ("BancoOriExt", transf.BancoOriExt)
		enc.WriteAttrStrZ("CtaDest", transf.CtaDest)
		enc.WriteAttrStrZ("BancoDestNal", transf.BancoDestNal)
		enc.WriteAttrStrZ("BancoDestExt", transf.BancoDestExt)
		enc.WriteAttrStrZ("Fecha", transf.Fecha.String())
		enc.WriteAttrStrZ("Benef", transf.Benef)
		enc.WriteAttrStrZ("RFC", transf.RFC)
		enc.WriteAttrDecimal("Monto", transf.Monto, 2)
		enc.WriteAttrStrZ("Moneda", string(transf.Moneda))
		enc.WriteAttrDecimal("TipCamb", transf.TipCamb, 5)
		enc.EndElem("Transferencia")
	}
}

func encodeOtrMetodoPagos(enc *encoder.Encoder, ctas []*OtrMetodoPago) {
	for _, otrmetpag := range ctas {
		enc.StartElem(polXS.Elem("OtrMetodoPago"))
		enc.WriteAttrStrZ("MetPagoPol", otrmetpag.MetPagoPol)
		enc.WriteAttrStrZ("Fecha", otrmetpag.Fecha.String())
		enc.WriteAttrStrZ("Benef", otrmetpag.Benef)
		enc.WriteAttrStrZ("RFC", otrmetpag.RFC)
		enc.WriteAttrDecimal("Monto", otrmetpag.Monto, 2)
		enc.WriteAttrStrZ("Moneda", string(otrmetpag.Moneda))
		enc.WriteAttrDecimal("TipCamb", otrmetpag.TipCamb, 5)
		enc.EndElem("OtrMetodoPago")
	}
}
