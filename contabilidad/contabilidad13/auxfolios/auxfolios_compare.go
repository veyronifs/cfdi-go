package auxfolios

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *RepAuxFolios) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *RepAuxFolios) {
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

	l1, l2 := len(v1.DetAuxFolios), len(v2.DetAuxFolios)
	compare.Comparable(diffs, l1, l2, path+".RepAuxFolios.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareDetAuxFolios(diffs, v1.DetAuxFolios[i], v2.DetAuxFolios[i], fmt.Sprintf("%s.RepAuxFolios[%d]", path, i))
		}
		return
	}
}

func compareDetAuxFolios(diffs *compare.Diffs, v1, v2 *DetAuxFol, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.NumUnIdenPol, v2.NumUnIdenPol, path+".NumUnIdenPol")
	compare.Comparable(diffs, v1.Fecha.String(), v2.Fecha.String(), path+".Fecha")

	l1, l2 := len(v1.ComprNal), len(v2.ComprNal)
	compare.Comparable(diffs, l1, l2, path+".ComprNal.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareComprNals(diffs, v1.ComprNal[i], v2.ComprNal[i], fmt.Sprintf("%s.ComprNal[%d]", path, i))
		}
		return
	}

	l3, l4 := len(v1.ComprNalOtr), len(v2.ComprNalOtr)
	compare.Comparable(diffs, l3, l4, path+".ComprNalOtr.Len()")
	if l3 == l4 {
		for i := 0; i < l3; i++ {
			compareComprNalOtrs(diffs, v1.ComprNalOtr[i], v2.ComprNalOtr[i], fmt.Sprintf("%s.ComprNalOtr[%d]", path, i))
		}
		return
	}

	l5, l6 := len(v1.ComprExt), len(v2.ComprExt)
	compare.Comparable(diffs, l5, l6, path+".CompExt.Len()")
	if l5 == l6 {
		for i := 0; i < l5; i++ {
			compareComprExts(diffs, v1.ComprExt[i], v2.ComprExt[i], fmt.Sprintf("%s.ComprExt[%d]", path, i))

		}
		return
	}
}

func compareComprNals(diffs *compare.Diffs, v1, v2 *ComprNal, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.UUIDCFDI, v2.UUIDCFDI, path+".UUIDCFDI")
	compare.Decimal(diffs, v1.MontoTotal, v2.MontoTotal, path+".MontoTotal")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Comparable(diffs, v1.MetPagoAux, v2.MetPagoAux, path+".MetPagoAux")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareComprNalOtrs(diffs *compare.Diffs, v1, v2 *ComprNalOtr, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.CFDCBBSerie, v2.CFDCBBSerie, path+".CFDCBBSerie")
	compare.Comparable(diffs, v1.CFDCBBNumFol, v2.CFDCBBNumFol, path+".CFDCBBNumFol")
	compare.Decimal(diffs, v1.MontoTotal, v2.MontoTotal, path+".MontoTotal")
	compare.Comparable(diffs, v1.RFC, v2.RFC, path+".RFC")
	compare.Comparable(diffs, v1.MetPagoAux, v2.MetPagoAux, path+".MetPagoAux")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}

func compareComprExts(diffs *compare.Diffs, v1, v2 *ComprExt, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}
	compare.Comparable(diffs, v1.NumFactExt, v2.NumFactExt, path+".NumFactExt")
	compare.Comparable(diffs, v1.TaxID, v2.TaxID, path+".TaxID")
	compare.Decimal(diffs, v1.MontoTotal, v2.MontoTotal, path+".MontoTotal")
	compare.Comparable(diffs, v1.MetPagoAux, v2.MetPagoAux, path+".MetPagoAux")
	compare.Comparable(diffs, v1.Moneda, v2.Moneda, path+".Moneda")
	compare.Decimal(diffs, v1.TipCamb, v2.TipCamb, path+".TipCamb")
}
