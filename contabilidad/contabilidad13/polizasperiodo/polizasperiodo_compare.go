package polizasperiodo

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *Polizas) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *Polizas) {
	path := ""
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.Version, v2.Version, path+".Version")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Comparable(diffs, v1.Mes, v2.Mes, path+".Mes")
	compare.Comparable(diffs, v1.Anio, v2.Anio, path+".Anio")
	compare.Comparable(diffs, v1.TipoSolicitud, v2.TipoSolicitud, path+".TipoSolicitud")
	compare.Comparable(diffs, v1.NumOrden, v2.NumOrden, path+".NumOrden")
	compare.Comparable(diffs, v1.NumTramite, v2.NumTramite, path+".NumTramite")
	compare.Comparable(diffs, v1.Sello, v2.Sello, path+".Sello")
	compare.Comparable(diffs, v1.NoCertificado, v2.NoCertificado, path+".NoCertificado")
	compare.Comparable(diffs, v1.Certificado, v2.Certificado, path+".Certificado")

	l1, l2 := len(v1.Polizas), len(v2.Polizas)
	compare.Comparable(diffs, l1, l2, path+".Poliza.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			comparePolizas(diffs, v1.Polizas[i], v2.Polizas[i], fmt.Sprintf("%s.Poliza[%d]", path, i))
		}
		return
	}

	// comparePolizas(diffs, v1.Polizas, v2.Polizas, path+".Polizas")
}

func comparePolizas(diffs *compare.Diffs, v1, v2 *Poliza, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.NumUnIdenPol, v2.NumUnIdenPol, path+".NumUnIdenPol")
	compare.Comparable(diffs, v1.Fecha.String(), v2.Fecha.String(), path+".Fecha")
	compare.Comparable(diffs, v1.Concepto, v2.Concepto, path+".Concepto")
	// compareTransacciones(diffs, []*Transaccion{}, []*Transaccion{}, path+".Poliza")
	// Revisar llamada a función
	// compareTransacciones(diffs, []*Transaccion{}, []*Transaccion{}, path+".Poliza")
	l1, l2 := len(v1.Transaccion), len(v2.Transaccion)
	compare.Comparable(diffs, l1, l2, path+".Polizas.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareTransacciones(diffs, v1.Transaccion[i], v2.Transaccion[i], fmt.Sprintf("%s.Transaccion[%d]", path, i))
		}
		return
	}

}

func compareTransacciones(diffs *compare.Diffs, v1, v2 *Transaccion, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.NumCta, v2.NumCta, path+".NumCta")
	compare.Comparable(diffs, v1.DesCta, v2.DesCta, path+".DesCta")
	compare.Comparable(diffs, v1.Concepto, v2.Concepto, path+".Concepto")
	compare.Decimal(diffs, v1.Debe, v2.Debe, path+".Debe")
	compare.Decimal(diffs, v1.Haber, v2.Haber, path+".Haber")

	l1, l2 := len(v1.CompNal), len(v2.CompNal)
	compare.Comparable(diffs, l1, l2, path+".CompNal.Len()")
	if l1 == l2 {
		for i := 0; i < l2; i++ {
			compareCompNals(diffs, v1.CompNal[i], v2.CompNal[i], fmt.Sprintf("%s.CompNal[%d]", path, i))
		}
		return
	}

	l3, l4 := len(v1.CompNalOtr), len(v2.CompNalOtr)
	compare.Comparable(diffs, l3, l4, path+".CompNalOtr.Len()")
	if l3 == l4 {
		for i := 0; i < l3; i++ {
			compareCompNalOtrs(diffs, v1.CompNalOtr[i], v2.CompNalOtr[i], fmt.Sprintf("%s.CompNalOtr[%d]", path, i))
		}
		return
	}

	l5, l6 := len(v1.CompExt), len(v2.CompExt)
	compare.Comparable(diffs, l5, l6, path+".CompExt.Len()")
	if l5 == l6 {
		for i := 0; i < l5; i++ {
			compareCompExts(diffs, v1.CompExt[i], v2.CompExt[i], fmt.Sprintf("%s.CompExt[%d]", path, i))

		}
		return
	}

	l7, l8 := len(v1.Cheque), len(v2.Cheque)
	compare.Comparable(diffs, l7, l8, path+".Cheque.Len()")
	if l7 == l8 {
		for i := 0; i < l7; i++ {
			compareCheques(diffs, v1.Cheque[i], v2.Cheque[i], fmt.Sprintf("%s.Cheque[%d]", path, i))
		}
		return
	}

	l9, l10 := len(v1.Transferencia), len(v2.Transferencia)
	compare.Comparable(diffs, l9, l10, path+".Transferencia.Len()")
	if l9 == l10 {
		for i := 0; i < l9; i++ {
			compareTransferencias(diffs, v1.Transferencia[i], v2.Transferencia[i], fmt.Sprintf("%s.Transferencia[%d]", path, i))
		}
		return
	}

	l11, l12 := len(v1.OtrMetodoPago), len(v2.OtrMetodoPago)
	compare.Comparable(diffs, l11, l12, path+".OtrMetodoPago.Len()")
	if l10 == l11 {
		for i := 0; i < l10; i++ {
			compareOtrMetodoPagos(diffs, v1.OtrMetodoPago[i], v2.OtrMetodoPago[i], fmt.Sprintf("%s.Cheque[%d]", path, i))
		}
		return
	}
}

func compareCompNals(diffs *compare.Diffs, v1, v2 *CompNal, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.UUIDCFDI, v2.UUIDCFDI, path+".UUIDCFDI")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Decimal(diffs, v1.MontoTotal, v2.MontoTotal, path+".MontoTotal")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareCompNalOtrs(diffs *compare.Diffs, v1, v2 *CompNalOtr, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.CFDCBBSerie, v2.CFDCBBSerie, path+".CFDCBBSerie")
	compare.Comparable(diffs, v1.CFDCBBNumFol, v2.CFDCBBNumFol, path+".CFDCBBNumFol")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Decimal(diffs, v1.MontoTotal, v2.MontoTotal, path+".MontoTotal")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareCompExts(diffs *compare.Diffs, v1, v2 *CompExt, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.NumFactExt, v2.NumFactExt, path+".NumFactExt")
	compare.Comparable(diffs, v1.TaxID, v2.TaxID, path+".TaxID")
	compare.Decimal(diffs, v1.MontoTotal, v2.MontoTotal, path+".MontoTotal")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareCheques(diffs *compare.Diffs, v1, v2 *Cheque, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.Num, v2.Num, path+".Num")
	compare.Comparable(diffs, v1.BanEmisNal, v2.BanEmisNal, path+".BanEmisNal")
	compare.Comparable(diffs, v1.BanEmisExt, v2.BanEmisExt, path+".BanEmisExt")
	compare.Comparable(diffs, v1.CtaOri, v2.CtaOri, path+".CtaOri")
	compare.Comparable(diffs, v1.Fecha, v2.Fecha, path+".Fecha")
	compare.Comparable(diffs, v1.Benef, v2.Benef, path+".Benef")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Decimal(diffs, v1.Monto, v2.Monto, path+".Monto")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareTransferencias(diffs *compare.Diffs, v1, v2 *Transferencia, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.CtaOri, v2.CtaOri, path+".CtaOri")
	compare.Comparable(diffs, v1.BancoDestNal, v2.BancoDestNal, path+".BancoOriNal")
	compare.Comparable(diffs, v1.BancoDestExt, v2.BancoDestExt, path+".BancoOriExt")
	compare.Comparable(diffs, v1.CtaDest, v2.CtaDest, path+".CtaDest")
	compare.Comparable(diffs, v1.BancoDestNal, v2.BancoDestNal, path+".BancoDestNal")
	compare.Comparable(diffs, v1.BancoDestExt, v2.BancoDestExt, path+".BancoDestExt")
	compare.Comparable(diffs, v1.Fecha, v2.Fecha, path+".Fecha")
	compare.Comparable(diffs, v1.Benef, v2.Benef, path+".Benef")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Decimal(diffs, v1.Monto, v2.Monto, path+".Monto")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareOtrMetodoPagos(diffs *compare.Diffs, v1, v2 *OtrMetodoPago, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.MetPagoPol, v2.MetPagoPol, path+".MetPagoPol")
	compare.Comparable(diffs, v1.Fecha, v2.Fecha, path+".Fecha")
	compare.Comparable(diffs, v1.Benef, v2.Benef, path+".Benef")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Decimal(diffs, v1.Monto, v2.Monto, path+".Monto")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}
