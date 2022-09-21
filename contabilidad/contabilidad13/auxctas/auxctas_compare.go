package auxctas

import (
	"fmt"

	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *AuxiliarCtas) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *AuxiliarCtas) {
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

	l1, l2 := len(v1.Cuentas), len(v2.Cuentas)
	compare.Comparable(diffs, l1, l2, path+".AuxiliarCtas.Len()")
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			compareCuentas(diffs, v1.Cuentas[i], v2.Cuentas[i], fmt.Sprintf("%s.Cuenta[%d]", path, i))
		}
		return
	}
}

func compareCuentas(diffs *compare.Diffs, v1, v2 *Cuenta, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.NumCta, v2.NumCta, path+".NumCta")
	compare.Comparable(diffs, v1.DesCta, v2.DesCta, path+".DesCta")
	compare.Decimal(diffs, v1.SaldoIni, v2.SaldoIni, path+".SaldoIni")
	compare.Decimal(diffs, v1.SaldoFin, v2.SaldoFin, path+".SaldoFin")

	l1, l2 := len(v1.DetallesAux), len(v2.DetallesAux)
	compare.Comparable(diffs, l1, l2, path+".DetallesAux.Len()")
	if l1 == l2 {
		for i := 0; i < l2; i++ {
			compareDetallesAux(diffs, v1.DetallesAux[i], v2.DetallesAux[i], fmt.Sprintf("%s.DetallesAux[%d]", path, i))
		}
		return
	}
}

func compareDetallesAux(diffs *compare.Diffs, v1, v2 *DetalleAux, path string) {
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Fecha, v2.Fecha, path+".Fecha")
	compare.Comparable(diffs, v1.NumUnIdenPol, v2.NumUnIdenPol, path+".NumUnIdenPol")
	compare.Comparable(diffs, v1.Concepto, v2.Concepto, path+".Concepto")
	compare.Decimal(diffs, v1.Debe, v2.Debe, path+".Debe")
	compare.Decimal(diffs, v1.Haber, v2.Haber, path+".Haber")
}
