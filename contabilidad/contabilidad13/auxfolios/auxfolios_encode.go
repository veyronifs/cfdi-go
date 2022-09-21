package auxfolios

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/veyronifs/cfdi-go/encoder"
)

func Unmarshal(data []byte) (*RepAuxFolios, error) {
	var RepAuxFolios RepAuxFolios
	err := xml.Unmarshal(data, &RepAuxFolios)
	if err != nil {
		return nil, err
	}
	return &RepAuxFolios, nil
}

var auxfolXS = encoder.NSElem{
	Prefix: "RepAux",
	NS:     "http://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios",
}

func Marshal(c *RepAuxFolios) ([]byte, error) {
	b := bytes.Buffer{}
	enc := encoder.NewEncoder(&b)
	enc.StartElem(auxfolXS.ElemXS("RepAuxFol"))
	defer enc.EndElem("RepAuxFol")

	enc.WriteAttrStr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	enc.WriteAttrStr("xsi:schemaLocation", auxfolXS.NS+" https://www.sat.gob.mx/esquemas/ContabilidadE/1_3/AuxiliarFolios/AuxiliarFolios_1_3.xsd")

	encodeRepAuxFolios(enc, c)
	enc.EndAllFlush()
	return b.Bytes(), nil
}

func encodeRepAuxFolios(enc *encoder.Encoder, c *RepAuxFolios) {
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
	for _, detauxfol := range c.DetAuxFolios {
		enc.StartElem(auxfolXS.Elem("DetAuxFol"))
		enc.WriteAttrStrZ("NumUnIdenPol", detauxfol.NumUnIdenPol)
		enc.WriteAttrStrZ("Fecha", detauxfol.Fecha.String())
		encondeComprNals(enc, detauxfol.ComprNal)
		encondeComprNalOtrs(enc, detauxfol.ComprNalOtr)
		encondeComprExts(enc, detauxfol.ComprExt)
		enc.EndElem("DetAuxFol")
	}
}

func encondeComprNals(enc *encoder.Encoder, ctas []*ComprNal) {
	for _, cnal := range ctas {
		enc.StartElem(auxfolXS.Elem("ComprNal"))
		enc.WriteAttrStrZ("UUID_CFDI", cnal.UUIDCFDI)
		enc.WriteAttrDecimal("MontoTotal", cnal.MontoTotal, 2)
		enc.WriteAttrStrZ("RFC", cnal.RFC)
		enc.WriteAttrStrZ("MetPagoAux", cnal.MetPagoAux)
		enc.WriteAttrStrZ("Moneda", string(cnal.Moneda))
		enc.WriteAttrDecimal("TipCamb", cnal.TipCamb, 5)
		enc.EndElem("ComprNal")
	}
}

func encondeComprNalOtrs(enc *encoder.Encoder, ctas []*ComprNalOtr) {
	for _, cnalotr := range ctas {
		enc.StartElem(auxfolXS.Elem("ComprNalOtr"))
		enc.WriteAttrStrZ("CFD_CBB_Serie", cnalotr.CFDCBBSerie)
		enc.WriteAttrStrZ("CFD_CBB_NumFol", fmt.Sprintf("%d", cnalotr.CFDCBBNumFol))
		enc.WriteAttrDecimal("MontoTotal", cnalotr.MontoTotal, 2)
		enc.WriteAttrStrZ("RFC", cnalotr.RFC)
		enc.WriteAttrStrZ("MetPagoAux", cnalotr.MetPagoAux)
		enc.WriteAttrStrZ("Moneda", string(cnalotr.Moneda))
		enc.WriteAttrDecimal("TipCamb", cnalotr.TipCamb, 5)
		enc.EndElem("ComprNalOtr")
	}
}

func encondeComprExts(enc *encoder.Encoder, ctas []*ComprExt) {
	for _, cext := range ctas {
		enc.StartElem(auxfolXS.Elem("ComrpExt"))
		enc.WriteAttrStrZ("NumFactExt", cext.NumFactExt)
		enc.WriteAttrStrZ("TaxID", cext.TaxID)
		enc.WriteAttrDecimal("MontoTotal", cext.MontoTotal, 2)
		enc.WriteAttrStrZ("MetPagoAux", cext.MetPagoAux)
		enc.WriteAttrStrZ("Moneda", string(cext.Moneda))
		enc.WriteAttrDecimal("TipCamb", cext.TipCamb, 5)
		enc.EndElem("ComprExt")
	}
}

/*
type ComprExt struct {
	NumFactExt string          `xml:"NumFactExt,attr"`
	TaxID      string          `xml:"TaxID,attr,omitempty"`
	MontoTotal decimal.Decimal `xml:"MontoTotal,attr"`
	MetPagoAux string          `xml:"MetPagoAux,attr,omitempty"`
	Moneda     types.Moneda    `xml:"Moneda,attr,omitempty"`
	TipCamb    decimal.Decimal `xml:"TipCamb,attr,omitempty"`
}

*/
