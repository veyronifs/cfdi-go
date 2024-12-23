package compare

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Diffs []*Diff

type Diff struct {
	Msg string
	V1  any
	V2  any
}

func NewDiffs() *Diffs {
	return &Diffs{}
}

func (d Diffs) Err() error {
	if len(d) == 0 {
		return nil
	}
	return fmt.Errorf("%s", d.Error())
}

func (d Diffs) Error() string {
	var s string
	for _, diff := range d {
		s += diff.Msg + "\n" +
			"\t" + fmt.Sprintf("%#v", diff.V1) + "!=" + fmt.Sprintf("%#v", diff.V2) + "\n"
	}
	return s
}

func (d *Diffs) Append(v1 any, v2 any, msg string, args ...interface{}) {
	*d = append(*d, &Diff{
		Msg: fmt.Sprintf(msg, args...),
		V1:  v1,
		V2:  v2,
	})
}

func Nil(d *Diffs, v1, v2 any, msg string, args ...interface{}) bool {
	if (v1 == nil && v2 != nil) || (v1 != nil && v2 == nil) {
		d.Append(v1, v2, msg, args...)
		return true
	}

	return false
}

func Comparable[S comparable](d *Diffs, v1, v2 S, msg string, args ...interface{}) bool {
	if v1 != v2 {
		d.Append(v1, v2, msg, args...)
		return true
	}

	return false
}

func Decimal(d *Diffs, v1, v2 decimal.Decimal, msg string, args ...interface{}) {
	if !v1.Equal(v2) {
		d.Append(v1.String(), v2.String(), msg, args...)
	}
}

func NullDecimal(d *Diffs, v1, v2 decimal.NullDecimal, msg string, args ...interface{}) {
	if !v1.Valid && !v2.Valid { // both are invalid
		return
	} else if (!v1.Valid && v2.Valid) || (v1.Valid && !v2.Valid) { // xor
		d.Append(v1.Valid, v2.Valid, msg+" (validity): ", args...)
		return
	}

	Decimal(d, v1.Decimal, v2.Decimal, msg, args...)
}
