package cfdi40_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veyronifs/cfdi-go/cfdi40"
)

/*
var (
	ErrAddendaExists = fmt.Errorf("addenda ya existente")
)

func AddAddenda(cfdi, addenda []byte) ([]byte, error) {
	cfdiCopy := make([]byte, 0, len(cfdi)+len(addenda)+40)
	copy(cfdiCopy, cfdi)

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

	add := "<cfdi:Addenda></cfdi:Addenda>" + string(addenda) + "</cfdi:Comprobante>"
	bytes.ReplaceAll(cfdi, []byte(`</cfdi:Comprobante>`), []byte(add))
	return cfdiCopy, nil
}
*/

func TestAddAddenda(t *testing.T) {
	tests := []struct {
		name        string
		cfdi        string
		addenda     string
		expected    string
		expectedErr error
	}{
		{
			name:        "empty",
			cfdi:        "",
			addenda:     `<someaddenda></someaddenda>`,
			expected:    "",
			expectedErr: cfdi40.ErrRequired,
		},
		{
			name:        "invalid cfdi:Comprobante",
			cfdi:        `<cfdi:Comprobante>`,
			addenda:     `<someaddenda></someaddenda>`,
			expected:    "",
			expectedErr: cfdi40.ErrInvalid,
		},
		{
			name:        "addenda already exists",
			cfdi:        `<cfdi:Comprobante><cfdi:Addenda></cfdi:Addenda></cfdi:Comprobante>`,
			addenda:     `<someaddenda></someaddenda>`,
			expected:    "",
			expectedErr: cfdi40.ErrAddendaExists,
		},
		{
			name:        "addenda ok",
			cfdi:        `<cfdi:Comprobante></cfdi:Comprobante>`,
			addenda:     `<someaddenda></someaddenda>`,
			expected:    `<cfdi:Comprobante><cfdi:Addenda><someaddenda></someaddenda></cfdi:Addenda></cfdi:Comprobante>`,
			expectedErr: nil,
		},
		{
			name:        "empty addenda",
			cfdi:        `<cfdi:Comprobante></cfdi:Comprobante>`,
			addenda:     "",
			expected:    `<cfdi:Comprobante></cfdi:Comprobante>`,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfdi, err := cfdi40.AddAddenda([]byte(tt.cfdi), []byte(tt.addenda))
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			if string(cfdi) != tt.expected {
				assert.Equal(t, tt.expected, string(cfdi))
			}
		})
	}
}
