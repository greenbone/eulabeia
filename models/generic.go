package models

import (
	"errors"
	"reflect"
)

var errInvalidTargetKind = errors.New("unexpected kind")
var errInvalidField = errors.New("invalid field")
var errInvalidValueKind = errors.New("kind of value does not match kind of field")

func SetValueOf(n interface{}, field string, value interface{}) error {
	ps := reflect.ValueOf(n)
	s := ps.Elem()
	if s.Kind() != reflect.Struct {
		return errInvalidTargetKind
	}
	f := s.FieldByName(field)
	if !f.IsValid() || !f.CanSet() {
		return errInvalidField
	}
	vv := reflect.ValueOf(value)
	if vv.Kind() != f.Kind() {
		return errInvalidValueKind
	}
	f.Set(vv)
	return nil
}
