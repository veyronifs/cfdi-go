package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	var d Diffs = nil
	d.Append("value1", "value2", "string")
	d.Append(1, 2, "int")
	d.Append(1.0, 2.0, "float")

	expected := Diffs{
		&Diff{
			Msg: "string",
			V1:  "value1",
			V2:  "value2",
		},
		&Diff{
			Msg: "int",
			V1:  1,
			V2:  2,
		},
		&Diff{
			Msg: "float",
			V1:  1.0,
			V2:  2.0,
		},
	}

	assert.Equal(t, expected, d)

	err := d.Err()
	expectedErr := "string\n\t\"value1\"!=\"value2\"\nint\n\t1!=2\nfloat\n\t1!=2\n"
	if expectedErr != err.Error() {
		t.Errorf("Expected error: \n%s\nBut got:\n%s\n", expectedErr, d.Error())
	}
}

func TestNoError(t *testing.T) {
	d := Diffs{}
	assert.Nil(t, d.Err())

	var d2 Diffs = nil
	assert.Nil(t, d2.Err())
}

/*

func (d *Diffs) CompareNil(value1 any, value2 any, msg string, args ...interface{}) bool {
	if (value1 == nil && value2 != nil) || (value1 != nil && value2 == nil) {
		d.Append(value1, value2, msg, args...)
		return true
	}

	return false
}

*/
func TestCompareNil(t *testing.T) {
	tests := []struct {
		TestName string
		Value1   interface{}
		Value2   interface{}
		Expected bool
	}{
		{
			TestName: "nil value1",
			Value1:   nil,
			Value2:   "value2",
			Expected: true,
		},
		{
			TestName: "nil value2",
			Value1:   "value1",
			Value2:   nil,
			Expected: true,
		},
		{
			TestName: "no nil",
			Value1:   "value1",
			Value2:   "value2",
			Expected: false,
		},
		{
			TestName: "nil value1 and value2",
			Value1:   nil,
			Value2:   nil,
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			var d Diffs = nil
			assert.Equal(t, test.Expected, Nil(&d, test.Value1, test.Value2, "msg"))
			err := d.Err()
			if test.Expected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCompareString(t *testing.T) {
	var diffs *Diffs = &Diffs{}
	actual := Comparable(diffs, "value1", "value2", "msg")
	expected := true
	if actual != expected {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
	assert.Error(t, diffs.Err())

	type Str string
	diffs = &Diffs{}
	actual = Comparable(diffs, Str("value1"), Str("value1"), "msg")
	expected = false
	if actual != expected {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
	assert.NoError(t, diffs.Err())
}
