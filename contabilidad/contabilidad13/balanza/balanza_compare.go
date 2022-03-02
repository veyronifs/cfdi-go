package balanza

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *Balanza) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *Balanza) {
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
	compare.Comparable(diffs, v1.TipoEnvio, v2.TipoEnvio, path+".TipoEnvio")
	compare.Comparable(diffs, v1.FechaModBal, v2.FechaModBal, path+".FechaModBal")
	compare.Comparable(diffs, v1.Sello, v2.Sello, path+".Sello")
	compare.Comparable(diffs, v1.NoCertificado, v2.NoCertificado, path+".NoCertificado")
	compare.Comparable(diffs, v1.Certificado, v2.Certificado, path+".Certificado")

	compareCtas(diffs, v1.Ctas, v2.Ctas, path+".Ctas")
}

func compareCtas(diffs *compare.Diffs, v1, v2 []*Cta, path string) {
	l1, l2 := len(v1), len(v2)
	compare.Comparable(diffs, l1, l2, path+".len()")
	if l1 != l2 {
		return
	}
	for i := range v1 {
		path := path + fmt.Sprintf(".Ctas[%d]", i)
		compare.Comparable(diffs, v1[i].NumCta, v2[i].NumCta, path+".NumCta")
		compare.Decimal(diffs, v1[i].SaldoIni, v2[i].SaldoIni, path+".SaldoIni")
		compare.Decimal(diffs, v1[i].Debe, v2[i].Debe, path+".Debe")
		compare.Decimal(diffs, v1[i].Haber, v2[i].Haber, path+".Haber")
		compare.Decimal(diffs, v1[i].SaldoFin, v2[i].SaldoFin, path+".SaldoFin")
	}
}
