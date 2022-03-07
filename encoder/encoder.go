package encoder

import (
	"io"
	"strconv"

	"github.com/shabbyrobe/xmlwriter"
	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/curconv"
)

type NSElem struct {
	Prefix string
	NS     string
}

func (xs *NSElem) Elem(s string) xmlwriter.Elem {
	return xmlwriter.Elem{Prefix: xs.Prefix, Name: s}
}

func (xs *NSElem) ElemXS(s string) xmlwriter.Elem {
	return xmlwriter.Elem{Prefix: xs.Prefix, URI: xs.NS, Name: s}
}

type Encoder struct {
	xml *xmlwriter.Writer
	err *xmlwriter.ErrCollector
	w   io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	e := Encoder{
		w:   w,
		xml: xmlwriter.Open(w),
		err: &xmlwriter.ErrCollector{},
	}
	e.err.Do(e.xml.StartDoc(xmlwriter.Doc{}))
	return &e
}

func (e *Encoder) StartElem(elem xmlwriter.Elem) {
	e.err.Do(e.xml.StartElem(elem))
}

func (e *Encoder) EndElem(name ...string) {
	e.err.Do(e.xml.EndElem(name...))
}

func (e *Encoder) EndAllFlush() {
	e.err.Do(e.xml.EndAllFlush())
}

func (e *Encoder) GetError() error {
	if e.err.Err != nil {
		return e.err
	}
	return nil
}

// WriteAttr writes one or more XML element attributes to the output.
func (e *Encoder) WriteAttr(attr ...xmlwriter.Attr) {
	e.err.Do(e.xml.WriteAttr(attr...))
}

// WriteAttrStr writes the named attribute to the output.
func (e *Encoder) WriteAttrStr(attr string, value string) {
	e.err.Do(e.xml.WriteAttr(xmlwriter.Attr{Name: attr, Value: value}))
}

// WriteAttrStrZ writes the named attribute to the output if the value is not a zero value.
func (e *Encoder) WriteAttrStrZ(attr string, value string) {
	if value != "" {
		e.WriteAttrStr(attr, value)
	}
}

// WriteAttrStrMax writes the named attribute to the output with the maximun number of characters.
func (e *Encoder) WriteAttrStrMax(attr string, value string, n int) {
	e.WriteAttrStr(attr, getMaxStr(value, n))
}

// WriteAttrStrMaxZ writes the named attribute to the output with the maximun number of characters, if the value is not a zero value.
func (e *Encoder) WriteAttrStrMaxZ(attr string, value string, n int) {
	if value != "" {
		e.WriteAttrStrMax(attr, value, n)
	}
}

// WriteAttrStrMaxEllipsis writes the named attribute to the output with the maximun number of characters.
//
// If the value is longer than n, it is truncated and an ellipsis is added.
func (e *Encoder) WriteAttrStrMaxEllipsis(attr string, value string, n int) {
	e.WriteAttrStr(attr, getMaxStrEllipsis(value, n))
}

// WriteAttrStrMaxEllipsisZ writes the named attribute to the output with the maximun number of characters, if the value is not a zero value.
//
// If the value is longer than n, it is truncated and an ellipsis is added.
func (e *Encoder) WriteAttrStrMaxEllipsisZ(attr string, value string, n int) {
	if value != "" {
		e.WriteAttrStrMaxEllipsis(attr, value, n)
	}
}

// WriteAttrDecimalCurr writes the named attribute with the value rounded to the MAX currency decimal.
func (e *Encoder) WriteAttrDecimalCurr(attr string, value decimal.Decimal, curr string) {
	e.WriteAttrStr(attr, curconv.RoundToMaxStr(value, curr))
}

// WriteAttrNullDecimalCurr writes the named attribute with the value rounded to the MAX currency decimal.
func (e *Encoder) WriteAttrNullDecimalCurr(attr string, value decimal.NullDecimal, curr string) {
	if value.Valid {
		e.WriteAttrDecimalCurr(attr, value.Decimal, curr)
	}
}

// WriteAttrDecimalCurrZ writes the named attribute with the value rounded to the MAX currency decimal,
// if the value is not a zero value.
func (e *Encoder) WriteAttrDecimalCurrZ(attr string, value decimal.Decimal, moneda string) {
	if !value.IsZero() {
		e.WriteAttrDecimalCurr(attr, value, moneda)
	}
}

// WriteAttrDecimal writes the named attribute with the value rounded to nDec decimals.
func (e *Encoder) WriteAttrDecimal(attr string, value decimal.Decimal, nDec int) {
	e.WriteAttrStr(attr, curconv.RoundToDecStr(value, nDec))
}

// WriteAttrNullDecimal writes the named attribute with the value rounded to nDec decimals.
func (e *Encoder) WriteAttrNullDecimal(attr string, value decimal.NullDecimal, nDec int) {
	if value.Valid {
		e.WriteAttrDecimal(attr, value.Decimal, nDec)
	}
}

// WriteAttrDecimal writes the named attribute with the value rounded to nDec decimals,
// if the value is not a zero value.
func (e *Encoder) WriteAttrDecimalZ(attr string, value decimal.Decimal, nDec int) {
	if !value.IsZero() {
		e.WriteAttrDecimal(attr, value, nDec)
	}
}

// WriteAttrInt writes the named attribute to the output.
func (e *Encoder) WriteAttrInt(attr string, value int) {
	e.err.Do(e.xml.WriteAttr(xmlwriter.Attr{Name: attr, Value: strconv.Itoa(value)}))
}

// WriteAttrIntZ writes the named attribute to the output if the value is not a zero value.
func (e *Encoder) WriteAttrIntZ(attr string, value int) {
	if value != 0 {
		e.WriteAttrInt(attr, value)
	}
}

func getMaxStr(value string, n int) string {
	r := []rune(value)
	if len(r) <= n {
		return value
	}
	return string(r[:n])
}

func getMaxStrEllipsis(value string, n int) string {
	r := []rune(value)
	if len(r) <= n {
		return value
	}
	if n > 3 {
		return string(r[:n-3]) + "..."
	}
	return string(r[:n])
}
