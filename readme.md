# cfdi-go

Implementacion del Estándar de Comprobante Fiscal Digital por Internet 4.0 (CFDI).

## CFDI 4.0

```go
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/types"
)

func main() {
	c := &cfdi40.Comprobante{
		Version:           cfdi40.Version,
		Moneda:            types.MonedaMXN,
		TipoDeComprobante: types.ComprobanteI,
		Exportacion:       types.Exportacion01,
		LugarExpedicion:   "20000",
		Serie:             "Serie",
		Folio:             "Folio",
		Fecha:             types.NewFechaHNow(),
		Emisor: &cfdi40.Emisor{
			Rfc:           "KAHO641101B39",
			Nombre:        "OSCAR KALA HAAK",
			RegimenFiscal: types.RegimenFiscal612,
		},
		Receptor: &cfdi40.Receptor{
			Rfc:                     "BAR011108CC6",
			Nombre:                  "BARCEL",
			DomicilioFiscalReceptor: "52000",
			RegimenFiscalReceptor:   "601",
			UsoCFDI:                 types.UsoCFDICP01,
		},
		Conceptos: cfdi40.Conceptos{
			{
				ObjetoImp:        types.ObjetoImp02,
				Cantidad:         decimal.NewFromFloat(1),
				ClaveProdServ:    "50192100",
				ClaveUnidad:      "XBX",
				Descripcion:      "Cacahuate",
				Importe:          decimal.NewFromFloat(1000),
				NoIdentificacion: "1",
				ValorUnitario:    decimal.NewFromFloat(1000),
				Impuestos: &cfdi40.ConceptoImpuestos{
					Traslados: cfdi40.ConceptoImpuestosTraslados{
						{
							Base:       decimal.NewFromFloat(1000),
							Impuesto:   types.ImpuestoIVA,
							TipoFactor: types.TipoFactorTasa,
							TasaOCuota: decimal.NewFromFloat(0.16),
							Importe:    decimal.NewFromFloat(160),
						},
					},
				},
			},
		},
	}
	c.Impuestos = cfdi40.NewImpuestos(*c)
	c.SubTotal, c.Descuento, c.Total = cfdi40.CalcularTotales(*c)

	// Marshal CFDI 4.0
	xml, err := cfdi40.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xml))

	// Unmarshal CFDI 4.0
	c2, err := cfdi40.Unmarshal(xml)
	if err != nil {
		panic(err)
	}
	fmt.Println(c2)
}

```

*XML Generado:*


```xml
<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd" Version="4.0" Serie="Serie" Folio="Folio" Fecha="2022-03-07T15:12:39" Moneda="MXN" TipoDeComprobante="I" LugarExpedicion="20000" Exportacion="01" SubTotal="1000" Total="1160"><cfdi:Emisor Rfc="KAHO641101B39" Nombre="OSCAR KALA HAAK" RegimenFiscal="612"/><cfdi:Receptor Rfc="BAR011108CC6" Nombre="BARCEL" DomicilioFiscalReceptor="52000" RegimenFiscalReceptor="601" UsoCFDI="CP01"/><cfdi:Conceptos><cfdi:Concepto ClaveProdServ="50192100" NoIdentificacion="1" ClaveUnidad="XBX" Descripcion="Cacahuate" ObjetoImp="02" ValorUnitario="1000" Cantidad="1" Importe="1000"><cfdi:Impuestos><cfdi:Traslados><cfdi:Traslado Base="1000" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.16" Importe="160"/></cfdi:Traslados></cfdi:Impuestos></cfdi:Concepto></cfdi:Conceptos><cfdi:Impuestos TotalImpuestosTrasladados="160"><cfdi:Traslados><cfdi:Traslado Base="1000" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.16" Importe="160"/></cfdi:Traslados></cfdi:Impuestos></cfdi:Comprobante>
```