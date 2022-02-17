package tfd11

import (
	"github.com/veyronifs/cfdi-go/compare"
)

func CompareEqual(v1, v2 *TimbreFiscalDigital) error {
	diffs := compare.NewDiffs()
	Compare(diffs, v1, v2)
	return diffs.Err()
}

func Compare(diffs *compare.Diffs, v1, v2 *TimbreFiscalDigital) {
	path := "TimbreFiscalDigital11"
	if compare.Nil(diffs, v1, v2, path) {
		return
	} else if v1 == nil || v2 == nil {
		return
	}

	compare.Comparable(diffs, v1.Version, v2.Version, path+".Version")
	compare.Comparable(diffs, v1.UUID, v2.UUID, path+".UUID")
	compare.Comparable(diffs, v1.FechaTimbrado, v2.FechaTimbrado, path+".FechaTimbrado")
	compare.Comparable(diffs, v1.RfcProvCertif, v2.RfcProvCertif, path+".RfcProvCertif")
	compare.Comparable(diffs, v1.Leyenda, v2.Leyenda, path+".Leyenda")
	compare.Comparable(diffs, v1.SelloCFD, v2.SelloCFD, path+".SelloCFD")
	compare.Comparable(diffs, v1.NoCertificadoSAT, v2.NoCertificadoSAT, path+".NoCertificadoSAT")
	compare.Comparable(diffs, v1.SelloSAT, v2.SelloSAT, path+".SelloSAT")
}
