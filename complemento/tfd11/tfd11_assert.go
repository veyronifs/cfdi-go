package tfd11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertEqual(t *testing.T, v1, v2 *TimbreFiscalDigital) {
	path := "TimbreFiscalDigital11"
	if v1 == nil || v2 == nil {
		assert.Nil(t, v1, path)
		assert.Nil(t, v2, path)
		return
	}
	assert.Equal(t, v1.Version, v2.Version, path+".Version")
	assert.Equal(t, v1.UUID, v2.UUID, path+".UUID")
	assert.Equal(t, v1.FechaTimbrado, v2.FechaTimbrado, path+".FechaTimbrado")
	assert.Equal(t, v1.RfcProvCertif, v2.RfcProvCertif, path+".RfcProvCertif")
	assert.Equal(t, v1.Leyenda, v2.Leyenda, path+".Leyenda")
	assert.Equal(t, v1.SelloCFD, v2.SelloCFD, path+".SelloCFD")
	assert.Equal(t, v1.NoCertificadoSAT, v2.NoCertificadoSAT, path+".NoCertificadoSAT")
	assert.Equal(t, v1.SelloSAT, v2.SelloSAT, path+".SelloSAT")
}
