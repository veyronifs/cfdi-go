package tests_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/veyronifs/cfdi-go/cfdi40"
	"github.com/veyronifs/cfdi-go/types"
	"github.com/veyronifs/pac-timbrado-go/profact"
)

var pac *profact.Profact

func TestMain(m *testing.M) {
	wsTest := os.Getenv("PAC_TEST_WS")
	if wsTest == "" {
		log.Println("PAC_TEST_WS environment variable not set")
		return
	}
	wsUserTest := os.Getenv("PAC_TEST_WS_USER")
	if wsUserTest == "" {
		log.Println("PAC_TEST_WS_USER environment variable not set")
		return
	}
	var err error
	pac, err = profact.NewProfact(wsTest, wsUserTest)
	if err != nil {
		log.Println(err)
		return
	}
	m.Run()
}

func testTimbrar(t *testing.T, c *cfdi40.Comprobante) {
	xml, err := timbrarComprobante(c)
	if err != nil {
		saveerror(t.Name(), xml, err)
		t.Fatal(err)
	}
	save(t.Name(), xml)
}

func timbrarComprobante(c *cfdi40.Comprobante) ([]byte, error) {
	if pac == nil {
		return nil, errors.New("PAC not configured")
	}

	bytez, err := cfdi40.Marshal(c)
	if err != nil {
		return nil, err
	}
	cfdiID := fmt.Sprintf("%s:%s-%s", c.TipoDeComprobante, c.Serie, c.Folio)
	xmlTimbrado, errorTimbre, err := pac.Timbrar(cfdiID, bytez)
	if err != nil {
		return bytez, err
	} else if errorTimbre != nil {
		return bytez, errorTimbre
	}
	return []byte(xmlTimbrado.XML), nil
}

// Habilitado para facturar (IVA exento, tasa 0%, 8% y 16%) Zona Fronteriza Norte y Sur
var emisor16_8_0 *cfdi40.Emisor = &cfdi40.Emisor{
	Rfc:           "KAHO641101B39",
	Nombre:        "OSCAR KALA HAAK",
	RegimenFiscal: types.RegimenFiscal612,
}

// newFechaH2 NewFechaH - 2 hours
func newFechaHNow2() types.FechaH {
	now := time.Now()
	now = now.Add(time.Hour * -2)
	return types.FechaH(now)
}

func save(name string, xmlTimbrado []byte) {
	// create outtest directory if not exists
	err := os.MkdirAll("outtests", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// write file
	err = ioutil.WriteFile("outtests/"+name+".xml", xmlTimbrado, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func saveerror(name string, xmlTimbrado []byte, err error) {
	app := []byte("<ERROR>\n" + err.Error() + "\n</ERROR>\n")
	save(name, append(app, xmlTimbrado...))
}
