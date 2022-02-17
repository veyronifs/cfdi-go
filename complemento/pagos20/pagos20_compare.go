package pagos20

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *Pagos) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *Pagos) {
	path := "Pagos20"
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Version, v2.Version, path+".Version")

	compareEqualTotales(diffs, v1.Totales, v2.Totales, path+".Totales")

	l1, l2 := len(v1.Pago), len(v2.Pago)
	compare.Comparable(diffs, l1, l2, path+".Pago.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareEqualPago(diffs, v1.Pago[i], v2.Pago[i], fmt.Sprintf("%s.Pago[%d]", path, i))
		}
	}
	return
}

func compareEqualTotales(diffs *compare.Diffs, v1, v2 *Totales, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Decimal(diffs, v1.TotalRetencionesIVA, v2.TotalRetencionesIVA, path+".TotalRetencionesIVA")
	compare.Decimal(diffs, v1.TotalRetencionesISR, v2.TotalRetencionesISR, path+".TotalRetencionesISR")
	compare.Decimal(diffs, v1.TotalRetencionesIEPS, v2.TotalRetencionesIEPS, path+".TotalRetencionesIEPS")
	compare.Decimal(diffs, v1.TotalTrasladosBaseIVA16, v2.TotalTrasladosBaseIVA16, path+".TotalTrasladosBaseIVA16")
	compare.Decimal(diffs, v1.TotalTrasladosImpuestoIVA16, v2.TotalTrasladosImpuestoIVA16, path+".TotalTrasladosImpuestoIVA16")
	compare.Decimal(diffs, v1.TotalTrasladosBaseIVA8, v2.TotalTrasladosBaseIVA8, path+".TotalTrasladosBaseIVA8")
	compare.Decimal(diffs, v1.TotalTrasladosImpuestoIVA8, v2.TotalTrasladosImpuestoIVA8, path+".TotalTrasladosImpuestoIVA8")
	compare.Decimal(diffs, v1.TotalTrasladosBaseIVA0, v2.TotalTrasladosBaseIVA0, path+".TotalTrasladosBaseIVA0")
	compare.Decimal(diffs, v1.TotalTrasladosImpuestoIVA0, v2.TotalTrasladosImpuestoIVA0, path+".TotalTrasladosImpuestoIVA0")
	compare.Decimal(diffs, v1.TotalTrasladosBaseIVAExento, v2.TotalTrasladosBaseIVAExento, path+".TotalTrasladosBaseIVAExento")
	compare.Decimal(diffs, v1.MontoTotalPagos, v2.MontoTotalPagos, path+".MontoTotalPagos")
}

func compareEqualPago(diffs *compare.Diffs, v1, v2 *Pago, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.FechaPago, v2.FechaPago, path+".FechaPago")
	compare.Comparable(diffs, v1.FormaDePagoP, v2.FormaDePagoP, path+".FormaDePagoP")
	compare.Comparable(diffs, v1.MonedaP, v2.MonedaP, path+".MonedaP")
	compare.Decimal(diffs, v1.TipoCambioP, v2.TipoCambioP, path+".TipoCambioP")
	compare.Decimal(diffs, v1.Monto, v2.Monto, path+".Monto")
	compare.Comparable(diffs, v1.NumOperacion, v2.NumOperacion, path+".NumOperacion")
	compare.Comparable(diffs, v1.RfcEmisorCtaOrd, v2.RfcEmisorCtaOrd, path+".RfcEmisorCtaOrd")
	compare.Comparable(diffs, v1.NomBancoOrdExt, v2.NomBancoOrdExt, path+".NomBancoOrdExt")
	compare.Comparable(diffs, v1.CtaOrdenante, v2.CtaOrdenante, path+".CtaOrdenante")
	compare.Comparable(diffs, v1.RfcEmisorCtaBen, v2.RfcEmisorCtaBen, path+".RfcEmisorCtaBen")
	compare.Comparable(diffs, v1.CtaBeneficiario, v2.CtaBeneficiario, path+".CtaBeneficiario")
	compare.Comparable(diffs, v1.TipoCadPago, v2.TipoCadPago, path+".TipoCadPago")
	compare.Comparable(diffs, v1.CertPago, v2.CertPago, path+".CertPago")
	compare.Comparable(diffs, v1.CadPago, v2.CadPago, path+".CadPago")
	compare.Comparable(diffs, v1.SelloPago, v2.SelloPago, path+".SelloPago")

	l1, l2 := len(v1.DoctoRelacionado), len(v2.DoctoRelacionado)
	compare.Comparable(diffs, l1, l2, path+".DoctoRelacionado.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareEqualDoctoRelacionado(diffs, v1.DoctoRelacionado[i], v2.DoctoRelacionado[i], fmt.Sprintf("%s.DoctoRelacionado[%d]", path, i))
		}
	}
	compareEqualImpuestosP(diffs, v1.ImpuestosP, v2.ImpuestosP, path+".ImpuestosP")
}

func compareEqualDoctoRelacionado(diffs *compare.Diffs, v1, v2 *DoctoRelacionado, path string) {
	compare.Comparable(diffs, v1.IdDocumento, v2.IdDocumento, path+".IdDocumento")
	compare.Comparable(diffs, v1.Serie, v2.Serie, path+".Serie")
	compare.Comparable(diffs, v1.Folio, v2.Folio, path+".Folio")
	compare.Comparable(diffs, v1.MonedaDR, v2.MonedaDR, path+".MonedaDR")
	compare.Decimal(diffs, v1.EquivalenciaDR, v2.EquivalenciaDR, path+".EquivalenciaDR")
	compare.Comparable(diffs, v1.NumParcialidad, v2.NumParcialidad, path+".NumParcialidad")
	compare.Decimal(diffs, v1.ImpSaldoAnt, v2.ImpSaldoAnt, path+".ImpSaldoAnt")
	compare.Decimal(diffs, v1.ImpPagado, v2.ImpPagado, path+".ImpPagado")
	compare.Decimal(diffs, v1.ImpSaldoInsoluto, v2.ImpSaldoInsoluto, path+".ImpSaldoInsoluto")

	compareEqualImpuestosDR(diffs, v1.ImpuestosDR, v2.ImpuestosDR, path+".ImpuestosDR")
}

func compareEqualImpuestosDR(diffs *compare.Diffs, v1, v2 *ImpuestosDR, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	l1, l2 := len(v1.TrasladosDR), len(v2.TrasladosDR)
	compare.Comparable(diffs, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareEqualTrasladoDR(diffs, v1.TrasladosDR[i], v2.TrasladosDR[i], fmt.Sprintf("%s.TrasladosDR[%d]", path, i))
		}
	}

	l1, l2 = len(v1.RetencionesDR), len(v2.RetencionesDR)
	compare.Comparable(diffs, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareEqualRetencionDR(diffs, v1.RetencionesDR[i], v2.RetencionesDR[i], fmt.Sprintf("%s.RetencionesDR[%d]", path, i))
		}
	}
}

func compareEqualTrasladoDR(diffs *compare.Diffs, v1, v2 *TrasladoDR, path string) {
	compare.Decimal(diffs, v1.BaseDR, v2.BaseDR, path+".BaseDR")
	compare.Comparable(diffs, v1.ImpuestoDR, v2.ImpuestoDR, path+".ImpuestoDR")
	compare.Comparable(diffs, v1.TipoFactorDR, v2.TipoFactorDR, path+".TipoFactorDR")
	compare.Decimal(diffs, v1.TasaOCuotaDR, v2.TasaOCuotaDR, path+".TasaOCuotaDR")
	compare.Decimal(diffs, v1.ImporteDR, v2.ImporteDR, path+".ImporteDR")
}

func compareEqualRetencionDR(diffs *compare.Diffs, v1, v2 *RetencionDR, path string) {
	compare.Decimal(diffs, v1.BaseDR, v2.BaseDR, path+".BaseDR")
	compare.Comparable(diffs, v1.ImpuestoDR, v2.ImpuestoDR, path+".ImpuestoDR")
	compare.Comparable(diffs, v1.TipoFactorDR, v2.TipoFactorDR, path+".TipoFactorDR")
	compare.Decimal(diffs, v1.TasaOCuotaDR, v2.TasaOCuotaDR, path+".TasaOCuotaDR")
	compare.Decimal(diffs, v1.ImporteDR, v2.ImporteDR, path+".ImporteDR")
}

func compareEqualImpuestosP(diffs *compare.Diffs, v1, v2 *ImpuestosP, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	l1, l2 := len(v1.TrasladosP), len(v2.TrasladosP)
	compare.Comparable(diffs, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareEqualTrasladoP(diffs, v1.TrasladosP[i], v2.TrasladosP[i], fmt.Sprintf("%s.TrasladosP[%d]", path, i))
		}
	}

	l1, l2 = len(v1.RetencionesP), len(v2.RetencionesP)
	compare.Comparable(diffs, l1, l2, path+".Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareEqualRetencionP(diffs, v1.RetencionesP[i], v2.RetencionesP[i], fmt.Sprintf("%s.RetencionesP[%d]", path, i))
		}
	}
}

func compareEqualTrasladoP(diffs *compare.Diffs, v1, v2 *TrasladoP, path string) {
	compare.Decimal(diffs, v1.BaseP, v2.BaseP, path+".BaseP")
	compare.Comparable(diffs, v1.ImpuestoP, v2.ImpuestoP, path+".ImpuestoP")
	compare.Comparable(diffs, v1.TipoFactorP, v2.TipoFactorP, path+".TipoFactorP")
	compare.Decimal(diffs, v1.TasaOCuotaP, v2.TasaOCuotaP, path+".TasaOCuotaP")
	compare.Decimal(diffs, v1.ImporteP, v2.ImporteP, path+".ImporteP")
}

func compareEqualRetencionP(diffs *compare.Diffs, v1, v2 *RetencionP, path string) {
	compare.Comparable(diffs, v1.ImpuestoP, v2.ImpuestoP, path+".ImpuestoP")
	compare.Decimal(diffs, v1.ImporteP, v2.ImporteP, path+".ImporteP")
}
