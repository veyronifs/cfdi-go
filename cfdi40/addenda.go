package cfdi40

import (
	"bytes"
	"fmt"
)

var (
	ErrAddendaExists = fmt.Errorf("addenda ya existente")
)

func AddAddenda(cfdi, addenda []byte) ([]byte, error) {
	if len(addenda) == 0 {
		return cfdi, nil
	}
	if len(cfdi) == 0 {
		return nil, fmt.Errorf("%w cfdi", ErrRequired)
	}
	if !bytes.Contains(cfdi, []byte(`</cfdi:Comprobante>`)) {
		return nil, fmt.Errorf("%w cfdi", ErrInvalid)
	}
	if bytes.Contains(cfdi, []byte(`<cfdi:Addenda`)) {
		return nil, ErrAddendaExists
	}

	add := "<cfdi:Addenda>" + string(addenda) + "</cfdi:Addenda></cfdi:Comprobante>"
	return bytes.ReplaceAll(cfdi, []byte(`</cfdi:Comprobante>`), []byte(add)), nil
}
