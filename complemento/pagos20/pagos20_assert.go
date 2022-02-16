package pagos20

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

/*
package pagos20

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

func Unmarshal(b []byte) (*Pagos, error) {
	pagos := &Pagos{}
	if err := xml.Unmarshal(b, pagos); err != nil {
		return nil, err
	}
	return pagos, nil
}

type Pagos struct {
	Totales *Totales `xml:"Totales"`
	Pago    []*Pago  `xml:"Pago"`
	Version string   `xml:"Version,attr"`
}

type Totales struct {
	TotalRetencionesIVA         decimal.Decimal `xml:"TotalRetencionesIVA,attr,omitempty"`
	TotalRetencionesISR         decimal.Decimal `xml:"TotalRetencionesISR,attr,omitempty"`
	TotalRetencionesIEPS        decimal.Decimal `xml:"TotalRetencionesIEPS,attr,omitempty"`
	TotalTrasladosBaseIVA16     decimal.Decimal `xml:"TotalTrasladosBaseIVA16,attr,omitempty"`
	TotalTrasladosImpuestoIVA16 decimal.Decimal `xml:"TotalTrasladosImpuestoIVA16,attr,omitempty"`
	TotalTrasladosBaseIVA8      decimal.Decimal `xml:"TotalTrasladosBaseIVA8,attr,omitempty"`
	TotalTrasladosImpuestoIVA8  decimal.Decimal `xml:"TotalTrasladosImpuestoIVA8,attr,omitempty"`
	TotalTrasladosBaseIVA0      decimal.Decimal `xml:"TotalTrasladosBaseIVA0,attr,omitempty"`
	TotalTrasladosImpuestoIVA0  decimal.Decimal `xml:"TotalTrasladosImpuestoIVA0,attr,omitempty"`
	TotalTrasladosBaseIVAExento decimal.Decimal `xml:"TotalTrasladosBaseIVAExento,attr,omitempty"`
	MontoTotalPagos             decimal.Decimal `xml:"MontoTotalPagos,attr"`
}

type Pago struct {
	DoctoRelacionado []*DoctoRelacionado `xml:"DoctoRelacionado"`
	ImpuestosP       *ImpuestosP         `xml:"ImpuestosP,omitempty"`
	FechaPago        types.Fecha       `xml:"FechaPago,attr"`
	FormaDePagoP     types.FormaPago     `xml:"FormaDePagoP,attr"`
	MonedaP          types.Moneda        `xml:"MonedaP,attr"`
	TipoCambioP      decimal.Decimal     `xml:"TipoCambioP,attr,omitempty"`
	Monto            decimal.Decimal     `xml:"Monto,attr"`
	NumOperacion     string              `xml:"NumOperacion,attr,omitempty"`
	RfcEmisorCtaOrd  string              `xml:"RfcEmisorCtaOrd,attr,omitempty"`
	NomBancoOrdExt   string              `xml:"NomBancoOrdExt,attr,omitempty"`
	CtaOrdenante     string              `xml:"CtaOrdenante,attr,omitempty"`
	RfcEmisorCtaBen  string              `xml:"RfcEmisorCtaBen,attr,omitempty"`
	CtaBeneficiario  string              `xml:"CtaBeneficiario,attr,omitempty"`
	TipoCadPago      CTipoCadenaPago     `xml:"TipoCadPago,attr,omitempty"`
	CertPago         string              `xml:"CertPago,attr,omitempty"`
	CadPago          string              `xml:"CadPago,attr,omitempty"`
	SelloPago        string              `xml:"SelloPago,attr,omitempty"`
}

// May be one of 01
type CTipoCadenaPago string

const (
	CTipoCadenaPago01 CTipoCadenaPago = "01"
)

type DoctoRelacionado struct {
	ImpuestosDR      *ImpuestosDR    `xml:"ImpuestosDR,omitempty"`
	IdDocumento      string          `xml:"IdDocumento,attr"`
	Serie            string          `xml:"Serie,attr,omitempty"`
	Folio            string          `xml:"Folio,attr,omitempty"`
	MonedaDR         types.Moneda    `xml:"MonedaDR,attr"`
	EquivalenciaDR   float64         `xml:"EquivalenciaDR,attr,omitempty"`
	NumParcialidad   int             `xml:"NumParcialidad,attr"`
	ImpSaldoAnt      decimal.Decimal `xml:"ImpSaldoAnt,attr"`
	ImpPagado        decimal.Decimal `xml:"ImpPagado,attr"`
	ImpSaldoInsoluto decimal.Decimal `xml:"ImpSaldoInsoluto,attr"`
	ObjetoImpDR      types.ObjetoImp `xml:"ObjetoImpDR,attr"`
}

type ImpuestosDR struct {
	RetencionesDR RetencionesDR `xml:"RetencionesDR,omitempty"`
	TrasladosDR   TrasladosDR   `xml:"TrasladosDR,omitempty"`
}

type RetencionesDR []*RetencionDR

func (ret *RetencionesDR) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret2 struct {
		RetencionDR []*RetencionDR `xml:"RetencionDR"`
	}
	if err := d.DecodeElement(&ret2, &start); err != nil {
		return err
	}
	*ret = ret2.RetencionDR
	return nil
}

type RetencionDR struct {
	BaseDR       decimal.Decimal  `xml:"BaseDR,attr"`
	ImpuestoDR   types.Impuesto   `xml:"ImpuestoDR,attr"`
	TipoFactorDR types.TipoFactor `xml:"TipoFactorDR,attr"`
	TasaOCuotaDR decimal.Decimal  `xml:"TasaOCuotaDR,attr"`
	ImporteDR    decimal.Decimal  `xml:"ImporteDR,attr"`
}

type TrasladosDR []*TrasladoDR

func (tras *TrasladosDR) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras2 struct {
		TrasladoDR []*TrasladoDR `xml:"TrasladoDR"`
	}
	if err := d.DecodeElement(&tras2, &start); err != nil {
		return err
	}
	*tras = tras2.TrasladoDR
	return nil
}

type TrasladoDR struct {
	BaseDR       decimal.Decimal  `xml:"BaseDR,attr"`
	ImpuestoDR   decimal.Decimal  `xml:"ImpuestoDR,attr"`
	TipoFactorDR types.TipoFactor `xml:"TipoFactorDR,attr"`
	TasaOCuotaDR decimal.Decimal  `xml:"TasaOCuotaDR,attr,omitempty"`
	ImporteDR    decimal.Decimal  `xml:"ImporteDR,attr,omitempty"`
}

type ImpuestosP struct {
	RetencionesP *RetencionesP `xml:"RetencionesP,omitempty"`
	TrasladosP   *TrasladosP   `xml:"TrasladosP,omitempty"`
}

type RetencionesP []*RetencionP

func (ret *RetencionesP) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ret2 struct {
		RetencionP []*RetencionP `xml:"RetencionP"`
	}
	if err := d.DecodeElement(&ret2, &start); err != nil {
		return err
	}
	*ret = ret2.RetencionP
	return nil
}

type RetencionP struct {
	ImpuestoP types.Impuesto  `xml:"ImpuestoP,attr"`
	ImporteP  decimal.Decimal `xml:"ImporteP,attr"`
}

type TrasladosP []*TrasladoP

func (tras *TrasladosP) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tras2 struct {
		TrasladoP []*TrasladoP `xml:"TrasladoP"`
	}
	if err := d.DecodeElement(&tras2, &start); err != nil {
		return err
	}
	*tras = tras2.TrasladoP
	return nil
}

type TrasladoP struct {
	BaseP       decimal.Decimal  `xml:"BaseP,attr"`
	ImpuestoP   types.Impuesto   `xml:"ImpuestoP,attr"`
	TipoFactorP types.TipoFactor `xml:"TipoFactorP,attr"`
	TasaOCuotaP decimal.Decimal  `xml:"TasaOCuotaP,attr,omitempty"`
	ImporteP    decimal.Decimal  `xml:"ImporteP,attr,omitempty"`
}

*/

func assertEqualDecimal(t *testing.T, ex, act decimal.Decimal, path ...any) {
	msg := fmt.Sprintf(path[0].(string), path[1:]...)
	assert.True(t, ex.Equal(act), msg+" %s != %s", ex.String(), act.String())
}

func AssertEqual(t *testing.T, v1, v2 *Pagos) {
	path := "Pagos20"
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.Version, v2.Version, path+".Version")

	assertEqualTotales(t, v1.Totales, v2.Totales, path+".Totales")

	l1, l2 := len(v1.Pago), len(v2.Pago)
	assert.Equal(t, l1, l2, path+".Pago.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			assertEqualPago(t, v1.Pago[i], v2.Pago[i], fmt.Sprintf("%s.Pago[%d]", path, i))
		}
	}
}
func assertEqualTotales(t *testing.T, v1, v2 *Totales, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assertEqualDecimal(t, v1.TotalRetencionesIVA, v2.TotalRetencionesIVA, path+".TotalRetencionesIVA")
	assertEqualDecimal(t, v1.TotalRetencionesIVA, v2.TotalRetencionesIVA, path+".TotalRetencionesIVA")
	assertEqualDecimal(t, v1.TotalRetencionesISR, v2.TotalRetencionesISR, path+".TotalRetencionesISR")
	assertEqualDecimal(t, v1.TotalRetencionesIEPS, v2.TotalRetencionesIEPS, path+".TotalRetencionesIEPS")
	assertEqualDecimal(t, v1.TotalTrasladosBaseIVA16, v2.TotalTrasladosBaseIVA16, path+".TotalTrasladosBaseIVA16")
	assertEqualDecimal(t, v1.TotalTrasladosImpuestoIVA16, v2.TotalTrasladosImpuestoIVA16, path+".TotalTrasladosImpuestoIVA16")
	assertEqualDecimal(t, v1.TotalTrasladosBaseIVA8, v2.TotalTrasladosBaseIVA8, path+".TotalTrasladosBaseIVA8")
	assertEqualDecimal(t, v1.TotalTrasladosImpuestoIVA8, v2.TotalTrasladosImpuestoIVA8, path+".TotalTrasladosImpuestoIVA8")
	assertEqualDecimal(t, v1.TotalTrasladosBaseIVA0, v2.TotalTrasladosBaseIVA0, path+".TotalTrasladosBaseIVA0")
	assertEqualDecimal(t, v1.TotalTrasladosImpuestoIVA0, v2.TotalTrasladosImpuestoIVA0, path+".TotalTrasladosImpuestoIVA0")
	assertEqualDecimal(t, v1.TotalTrasladosBaseIVAExento, v2.TotalTrasladosBaseIVAExento, path+".TotalTrasladosBaseIVAExento")
	assertEqualDecimal(t, v1.MontoTotalPagos, v2.MontoTotalPagos, path+".MontoTotalPagos")
}

func assertEqualPago(t *testing.T, v1, v2 *Pago, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.FechaPago.String(), v2.FechaPago.String(), path+".FechaPago")
	assert.Equal(t, v1.FormaDePagoP, v2.FormaDePagoP, path+".FormaDePagoP")
	assert.Equal(t, v1.MonedaP, v2.MonedaP, path+".MonedaP")
	assertEqualDecimal(t, v1.TipoCambioP, v2.TipoCambioP, path+".TipoCambioP")
	assertEqualDecimal(t, v1.Monto, v2.Monto, path+".Monto")
	assert.Equal(t, v1.NumOperacion, v2.NumOperacion, path+".NumOperacion")
	assert.Equal(t, v1.RfcEmisorCtaOrd, v2.RfcEmisorCtaOrd, path+".RfcEmisorCtaOrd")
	assert.Equal(t, v1.NomBancoOrdExt, v2.NomBancoOrdExt, path+".NomBancoOrdExt")
	assert.Equal(t, v1.CtaOrdenante, v2.CtaOrdenante, path+".CtaOrdenante")
	assert.Equal(t, v1.RfcEmisorCtaBen, v2.RfcEmisorCtaBen, path+".RfcEmisorCtaBen")
	assert.Equal(t, v1.CtaBeneficiario, v2.CtaBeneficiario, path+".CtaBeneficiario")
	assert.Equal(t, v1.TipoCadPago, v2.TipoCadPago, path+".TipoCadPago")
	assert.Equal(t, v1.CertPago, v2.CertPago, path+".CertPago")
	assert.Equal(t, v1.CadPago, v2.CadPago, path+".CadPago")
	assert.Equal(t, v1.SelloPago, v2.SelloPago, path+".SelloPago")

	l1, l2 := len(v1.DoctoRelacionado), len(v2.DoctoRelacionado)
	assert.Equal(t, l1, l2, path+".DoctoRelacionado.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			assertEqualDoctoRelacionado(t, v1.DoctoRelacionado[i], v2.DoctoRelacionado[i], fmt.Sprintf("%s.DoctoRelacionado[%d]", path, i))
		}
	}
	assertEqualImpuestosP(t, v1.ImpuestosP, v2.ImpuestosP, path+".ImpuestosP")
}

func assertEqualDoctoRelacionado(t *testing.T, v1, v2 *DoctoRelacionado, path string) {
	assert.Equal(t, v1.IdDocumento, v2.IdDocumento, path+".IdDocumento")
	assert.Equal(t, v1.Serie, v2.Serie, path+".Serie")
	assert.Equal(t, v1.Folio, v2.Folio, path+".Folio")
	assert.Equal(t, v1.MonedaDR, v2.MonedaDR, path+".MonedaDR")
	assertEqualDecimal(t, v1.EquivalenciaDR, v2.EquivalenciaDR, path+".EquivalenciaDR")
	assert.Equal(t, v1.NumParcialidad, v2.NumParcialidad, path+".NumParcialidad")
	assertEqualDecimal(t, v1.ImpSaldoAnt, v2.ImpSaldoAnt, path+".ImpSaldoAnt")
	assertEqualDecimal(t, v1.ImpPagado, v2.ImpPagado, path+".ImpPagado")
	assertEqualDecimal(t, v1.ImpSaldoInsoluto, v2.ImpSaldoInsoluto, path+".ImpSaldoInsoluto")

	assertEqualImpuestosDR(t, v1.ImpuestosDR, v2.ImpuestosDR, path+".ImpuestosDR")
}

func assertEqualImpuestosDR(t *testing.T, v1, v2 *ImpuestosDR, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	l1, l2 := len(v1.TrasladosDR), len(v2.TrasladosDR)
	assert.Equal(t, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			assertEqualTrasladoDR(t, v1.TrasladosDR[i], v2.TrasladosDR[i], fmt.Sprintf("%s.TrasladosDR[%d]", path, i))
		}
	}

	l1, l2 = len(v1.RetencionesDR), len(v2.RetencionesDR)
	assert.Equal(t, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			assertEqualRetencionDR(t, v1.RetencionesDR[i], v2.RetencionesDR[i], fmt.Sprintf("%s.RetencionesDR[%d]", path, i))
		}
	}
}

func assertEqualTrasladoDR(t *testing.T, v1, v2 *TrasladoDR, path string) {
	assertEqualDecimal(t, v1.BaseDR, v2.BaseDR, path+".BaseDR")
	assert.Equal(t, v1.ImpuestoDR, v2.ImpuestoDR, path+".ImpuestoDR")
	assert.Equal(t, v1.TipoFactorDR, v2.TipoFactorDR, path+".TipoFactorDR")
	assertEqualDecimal(t, v1.TasaOCuotaDR, v2.TasaOCuotaDR, path+".TasaOCuotaDR")
	assertEqualDecimal(t, v1.ImporteDR, v2.ImporteDR, path+".ImporteDR")
}

func assertEqualRetencionDR(t *testing.T, v1, v2 *RetencionDR, path string) {
	assertEqualDecimal(t, v1.BaseDR, v2.BaseDR, path+".BaseDR")
	assert.Equal(t, v1.ImpuestoDR, v2.ImpuestoDR, path+".ImpuestoDR")
	assert.Equal(t, v1.TipoFactorDR, v2.TipoFactorDR, path+".TipoFactorDR")
	assertEqualDecimal(t, v1.TasaOCuotaDR, v2.TasaOCuotaDR, path+".TasaOCuotaDR")
	assertEqualDecimal(t, v1.ImporteDR, v2.ImporteDR, path+".ImporteDR")
}

func assertEqualImpuestosP(t *testing.T, v1, v2 *ImpuestosP, path string) {
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	l1, l2 := len(v1.TrasladosP), len(v2.TrasladosP)
	assert.Equal(t, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			assertEqualTrasladoP(t, v1.TrasladosP[i], v2.TrasladosP[i], fmt.Sprintf("%s.TrasladosP[%d]", path, i))
		}
	}

	l1, l2 = len(v1.RetencionesP), len(v2.RetencionesP)
	assert.Equal(t, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			assertEqualRetencionP(t, v1.RetencionesP[i], v2.RetencionesP[i], fmt.Sprintf("%s.RetencionesP[%d]", path, i))
		}
	}
}

func assertEqualTrasladoP(t *testing.T, v1, v2 *TrasladoP, path string) {
	assertEqualDecimal(t, v1.BaseP, v2.BaseP, path+".BaseP")
	assert.Equal(t, v1.ImpuestoP, v2.ImpuestoP, path+".ImpuestoP")
	assert.Equal(t, v1.TipoFactorP, v2.TipoFactorP, path+".TipoFactorP")
	assertEqualDecimal(t, v1.TasaOCuotaP, v2.TasaOCuotaP, path+".TasaOCuotaP")
	assertEqualDecimal(t, v1.ImporteP, v2.ImporteP, path+".ImporteP")
}

func assertEqualRetencionP(t *testing.T, v1, v2 *RetencionP, path string) {
	assert.Equal(t, v1.ImpuestoP, v2.ImpuestoP, path+".ImpuestoP")
	assertEqualDecimal(t, v1.ImporteP, v2.ImporteP, path+".ImporteP")
}
