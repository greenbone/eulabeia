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
		err      string
		expected interface{}
	}{
		{
			"A",
			"b",
			"",
			TestStruct{"b", 0},
		},
		{
			"B",
			"b",
			"field type (int) does not match value type (string)",
			TestStruct{"b", 0},
		},
		{
			"C",
			"b",
			"field (C) not found on target (models.TestStruct)",
			TestStruct{"b", 0},
		},
	}
	for i, test := range tests {
		target := &TestStruct{"a", 0}
		err := SetValueOf(target, test.name, test.val)

		if err != nil {
			if test.err == "" || test.err != err.Error() {
				t.Errorf("[%d] returned '%s' but expected '%v'", i, err, test.err)
			}
		}
		if test.err == "" {
			if test.expected != *target {
				t.Errorf("[%d] returned %v but expected %v", i, target, test.expected)

			}
		}
	}
}
