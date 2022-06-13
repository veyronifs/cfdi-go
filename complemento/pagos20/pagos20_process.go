package pagos20

import (
	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

func CalcTotales(pagos *Pagos) *Totales {
	totales := Totales{}
	for _, pago := range pagos.Pago {
		totales.MontoTotalPagos = decimalMxn(pago.Monto, pago.MonedaP, pago.TipoCambioP).Add(totales.MontoTotalPagos)

		for _, docRel := range pago.DoctoRelacionado {
			if docRel.ImpuestosDR == nil {
				continue
			}
			for _, tras := range docRel.ImpuestosDR.TrasladosDR {
				calcTotalesSumTras(&totales, pago, docRel, tras)
			}
			for _, ret := range docRel.ImpuestosDR.RetencionesDR {
				calcTotalesSumRet(&totales, pago, docRel, ret)
			}
		}
	}
	return &totales
}

func calcTotalesSumTras(totales *Totales, pago *Pago, docRel *DoctoRelacionado, tras *TrasladoDR) {
	if tras.ImpuestoDR != types.ImpuestoIVA {
		return
	}
	importeMxn := decimalMxn(tras.ImporteDR.Div(docRel.EquivalenciaDR), pago.MonedaP, pago.TipoCambioP)
	baseMxn := decimalMxn(tras.BaseDR.Div(docRel.EquivalenciaDR), pago.MonedaP, pago.TipoCambioP)
	switch tras.TipoFactorDR {
	case types.TipoFactorExento:
		totales.TotalTrasladosBaseIVAExento.Valid = true
		totales.TotalTrasladosBaseIVAExento.Decimal = totales.TotalTrasladosBaseIVAExento.Decimal.Add(baseMxn)
	case types.TipoFactorTasa:
		tasaCuota := tras.TasaOCuotaDR.StringFixed(6)
		switch tasaCuota {
		case "0.160000":
			totales.TotalTrasladosBaseIVA16.Valid = true
			totales.TotalTrasladosBaseIVA16.Decimal = totales.TotalTrasladosBaseIVA16.Decimal.Add(baseMxn)

			totales.TotalTrasladosImpuestoIVA16.Valid = true
			totales.TotalTrasladosImpuestoIVA16.Decimal = totales.TotalTrasladosImpuestoIVA16.Decimal.Add(importeMxn)
		case "0.080000":
			totales.TotalTrasladosBaseIVA8.Valid = true
			totales.TotalTrasladosBaseIVA8.Decimal = totales.TotalTrasladosBaseIVA8.Decimal.Add(baseMxn)

			totales.TotalTrasladosImpuestoIVA8.Valid = true
			totales.TotalTrasladosImpuestoIVA8.Decimal = totales.TotalTrasladosImpuestoIVA8.Decimal.Add(importeMxn)
		case "0.000000":
			totales.TotalTrasladosBaseIVA0.Valid = true
			totales.TotalTrasladosBaseIVA0.Decimal = totales.TotalTrasladosBaseIVA0.Decimal.Add(baseMxn)

			totales.TotalTrasladosImpuestoIVA0.Valid = true
			totales.TotalTrasladosImpuestoIVA0.Decimal = totales.TotalTrasladosImpuestoIVA0.Decimal.Add(importeMxn)
		}
	}
}

func calcTotalesSumRet(totales *Totales, pago *Pago, docRel *DoctoRelacionado, ret *RetencionDR) {
	importeMxn := decimalMxn(ret.ImporteDR.Div(docRel.EquivalenciaDR), pago.MonedaP, pago.TipoCambioP)
	switch ret.ImpuestoDR {
	case types.ImpuestoIVA:
		totales.TotalRetencionesIVA.Valid = true
		totales.TotalRetencionesIVA.Decimal = totales.TotalRetencionesIVA.Decimal.Add(importeMxn)
	case types.ImpuestoISR:
		totales.TotalRetencionesISR.Valid = true
		totales.TotalRetencionesISR.Decimal = totales.TotalRetencionesISR.Decimal.Add(importeMxn)
	case types.ImpuestoIEPS:
		totales.TotalRetencionesIEPS.Valid = true
		totales.TotalRetencionesIEPS.Decimal = totales.TotalRetencionesIEPS.Decimal.Add(importeMxn)
	}
}
func decimalMxn(d decimal.Decimal, moneda types.Moneda, tipoCambio decimal.Decimal) decimal.Decimal {
	if moneda == types.MonedaMXN {
		return d
	}
	return d.Mul(tipoCambio)
}

func CalcImpuestosP(pago *Pago) *ImpuestosP {
	impP := &ImpuestosP{}
	for _, docRel := range pago.DoctoRelacionado {
		if docRel.ImpuestosDR == nil {
			continue
		}
		for _, tras := range docRel.ImpuestosDR.TrasladosDR {
			impP.AddTraslado(tras.BaseDR.Div(docRel.EquivalenciaDR), tras.ImpuestoDR, tras.TipoFactorDR, tras.TasaOCuotaDR, tras.ImporteDR.Div(docRel.EquivalenciaDR))
		}
		for _, ret := range docRel.ImpuestosDR.RetencionesDR {
			impP.AddRetencion(ret.ImpuestoDR, ret.ImporteDR.Div(docRel.EquivalenciaDR))
		}
	}
	return impP
}
