package cfdi40

import (
	"encoding/xml"
	"io"
	"testing"

	"github.com/veyronifs/cfdi-go/complemento/cartaporte31"
	"github.com/veyronifs/cfdi-go/complemento/comext20"
	"github.com/veyronifs/cfdi-go/complemento/pagos20"
	"github.com/veyronifs/cfdi-go/complemento/tfd11"
)

// xmlAttributes decodes the root attributes of an XML document.
type xmlAttributes map[string]string

func (attrs *xmlAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *attrs == nil {
		*attrs = make(xmlAttributes)
	}
	for _, attr := range start.Attr {
		if attr.Name.Space == "" {
			(*attrs)[attr.Name.Local] = attr.Value
		} else {
			(*attrs)[attr.Name.Space+":"+attr.Name.Local] = attr.Value
		}
	}
	for {
		_, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func TestNamespaces(t *testing.T) {
	tests := []struct {
		name              string
		c                 *Comprobante
		expectedNS        map[string]string
		expectedLocations string
	}{
		{
			name: "All",
			c: &Comprobante{
				Complemento: &Complemento{
					CartaPorte: &cartaporte31.CartaPorte31{},
					Pagos20:    &pagos20.Pagos{},
					CCE20:      &comext20.ComercioExterior{},
					TFD11:      &tfd11.TimbreFiscalDigital{},
				},
			},
			expectedNS: map[string]string{
				"xmlns:cfdi":         "http://www.sat.gob.mx/cfd/4",
				"xmlns:xsi":          "http://www.w3.org/2001/XMLSchema-instance",
				"xmlns:cartaporte31": "http://www.sat.gob.mx/CartaPorte31",
				"xmlns:pago20":       "http://www.sat.gob.mx/Pagos20",
				"xmlns:cce20":        "http://www.sat.gob.mx/ComercioExterior20",
			},
			expectedLocations: "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd http://www.sat.gob.mx/CartaPorte31 http://www.sat.gob.mx/sitio_internet/cfd/CartaPorte/CartaPorte31.xsd http://www.sat.gob.mx/Pagos20 http://www.sat.gob.mx/sitio_internet/cfd/Pagos/Pagos20.xsd http://www.sat.gob.mx/ComercioExterior20 http://www.sat.gob.mx/sitio_internet/cfd/ComercioExterior20/ComercioExterior20.xsd",
		},
		{
			name: "NoComplemento",
			c: &Comprobante{
				Complemento: &Complemento{},
			},
			expectedNS: map[string]string{
				"xmlns:cfdi":         "http://www.sat.gob.mx/cfd/4",
				"xmlns:xsi":          "http://www.w3.org/2001/XMLSchema-instance",
				"xmlns:cartaporte31": "",
				"xmlns:pago20":       "",
				"xmlns:cce20":        "",
			},
			expectedLocations: "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd",
		},
		{
			name: "cartaporte31",
			c: &Comprobante{
				Complemento: &Complemento{
					CartaPorte: &cartaporte31.CartaPorte31{},
				},
			},
			expectedNS: map[string]string{
				"xmlns:cfdi":         "http://www.sat.gob.mx/cfd/4",
				"xmlns:xsi":          "http://www.w3.org/2001/XMLSchema-instance",
				"xmlns:cartaporte31": "http://www.sat.gob.mx/CartaPorte31",
				"xmlns:pago20":       "",
				"xmlns:cce20":        "",
			},
			expectedLocations: "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd http://www.sat.gob.mx/CartaPorte31 http://www.sat.gob.mx/sitio_internet/cfd/CartaPorte/CartaPorte31.xsd",
		},
		{
			name: "pago20",
			c: &Comprobante{
				Complemento: &Complemento{
					Pagos20: &pagos20.Pagos{},
				},
			},
			expectedNS: map[string]string{
				"xmlns:cfdi":         "http://www.sat.gob.mx/cfd/4",
				"xmlns:xsi":          "http://www.w3.org/2001/XMLSchema-instance",
				"xmlns:cartaporte31": "",
				"xmlns:pago20":       "http://www.sat.gob.mx/Pagos20",
				"xmlns:cce20":        "",
			},
			expectedLocations: "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd http://www.sat.gob.mx/Pagos20 http://www.sat.gob.mx/sitio_internet/cfd/Pagos/Pagos20.xsd",
		},
		{
			name: "cce20",
			c: &Comprobante{
				Complemento: &Complemento{
					CCE20: &comext20.ComercioExterior{},
				},
			},
			expectedNS: map[string]string{
				"xmlns:cfdi":         "http://www.sat.gob.mx/cfd/4",
				"xmlns:xsi":          "http://www.w3.org/2001/XMLSchema-instance",
				"xmlns:cartaporte31": "",
				"xmlns:pago20":       "",
				"xmlns:cce20":        "http://www.sat.gob.mx/ComercioExterior20",
			},
			expectedLocations: "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd http://www.sat.gob.mx/ComercioExterior20 http://www.sat.gob.mx/sitio_internet/cfd/ComercioExterior20/ComercioExterior20.xsd",
		},
		{
			name: "cartaporte31, cce20",
			c: &Comprobante{
				Complemento: &Complemento{
					CartaPorte: &cartaporte31.CartaPorte31{},
					CCE20:      &comext20.ComercioExterior{},
				},
			},
			expectedNS: map[string]string{
				"xmlns:cfdi":         "http://www.sat.gob.mx/cfd/4",
				"xmlns:xsi":          "http://www.w3.org/2001/XMLSchema-instance",
				"xmlns:cartaporte31": "http://www.sat.gob.mx/CartaPorte31",
				"xmlns:pago20":       "",
				"xmlns:cce20":        "http://www.sat.gob.mx/ComercioExterior20",
			},
			expectedLocations: "http://www.sat.gob.mx/cfd/4 http://www.sat.gob.mx/sitio_internet/cfd/4/cfdv40.xsd http://www.sat.gob.mx/CartaPorte31 http://www.sat.gob.mx/sitio_internet/cfd/CartaPorte/CartaPorte31.xsd http://www.sat.gob.mx/ComercioExterior20 http://www.sat.gob.mx/sitio_internet/cfd/ComercioExterior20/ComercioExterior20.xsd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytez, err := Marshal(tt.c)
			if err != nil {
				t.Errorf("Marshal() error = %v", err)
				return
			}

			m := xmlAttributes{}
			err = xml.Unmarshal(bytez, &m)
			if err != nil {
				t.Errorf("Unmarshal() error = %v", err)
				return
			}

			/*fmt.Println(string(bytez))
			fmt.Println("************************************************")
			j, _ := json.MarshalIndent(m, "", "  ")
			fmt.Println(string(j))*/

			for k, v := range tt.expectedNS {
				if v == "" {
					if _, ok := m[k]; ok {
						t.Errorf("%s should not be present", k)
					}
				} else {
					if m[k] != v {
						t.Errorf("expected %s (%s), got (%s)", k, v, m[k])
					}
				}
			}
			if actual := m["http://www.w3.org/2001/XMLSchema-instance:schemaLocation"]; actual != tt.expectedLocations {
				t.Errorf("expected schemaLocation (%s), got (%s)", tt.expectedLocations, actual)
			}
		})
	}
}
