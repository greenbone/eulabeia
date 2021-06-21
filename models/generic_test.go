package models

import (
	"testing"
)

type TestStruct struct {
	A string
	B int
}

func TestSetValueOf(t *testing.T) {
	var tests = []struct {
		name     string
		val      interface{}
		err      error
		expected interface{}
	}{
		{
			"A",
			"b",
			nil,
			TestStruct{"b", 0},
		},
		{
			"B",
			"b",
			errInvalidValueKind,
			TestStruct{"b", 0},
		},
		{
			"C",
			"b",
			errInvalidField,
			TestStruct{"b", 0},
		},
	}
	for i, test := range tests {
		target := &TestStruct{"a", 0}
		err := SetValueOf(target, test.name, test.val)

		if err != test.err {
			t.Errorf("[%d] returned %v but expected %v", i, err, test.err)
		}
		if test.err == nil {
			if test.expected != *target {
				t.Errorf("[%d] returned %v but expected %v", i, target, test.expected)

			}
		}
	}
}
