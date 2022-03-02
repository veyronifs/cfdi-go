package types

import (
	"bytes"
	"encoding/xml"
	"time"
)

func NewFechaH(value string) (FechaH, error) {
	t, err := time.Parse("2006-01-02T15:04:05", value)
	return FechaH(t), err
}

func NewFechaHNow() FechaH {
	return FechaH(time.Now().Truncate(time.Second))
}

// Tipo definido para la expresión de la fecha y hora. Se expresa en la forma AAAA-MM-DDThh:mm:ss
type FechaH time.Time

func (t *FechaH) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t FechaH) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}
func (t FechaH) String() string {
	return time.Time(t).Format("2006-01-02T15:04:05")
}
func (t FechaH) Format(s string) string {
	return time.Time(t).Format(s)
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

func NewFecha(value string) (Fecha, error) {
	t, err := time.Parse("2006-01-02", value)
	return Fecha(t), err
}

func NewFechaNow() Fecha {
	return Fecha(time.Now().Truncate(time.Second))
}

// Tipo definido para la expresión de la fecha. Se expresa en la forma AAAA-MM-DD.
type Fecha time.Time

func (t *Fecha) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t Fecha) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}
func (t Fecha) String() string {
	return time.Time(t).Format("2006-01-02")
}
func (t Fecha) Format(s string) string {
	return time.Time(t).Format(s)
}
