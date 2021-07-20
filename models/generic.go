// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package models

import (
	"reflect"
)

type InvalidTargetError struct {
	Type reflect.Type
}

func (e *InvalidTargetError) Error() string {
	if e.Type == nil {
		return "target (nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "target (non-pointer " + e.Type.String() + ")"
	}
	return "target (nil " + e.Type.String() + ")"
}

type InvalidFieldError struct {
	Type  reflect.Type
	Field string
}

func (e *InvalidFieldError) Error() string {
	return "field (" + e.Field + ") not found on target (" + e.Type.String() + ")"
}

type InvalidValueError struct {
	FieldType reflect.Type
	ValueType reflect.Type
}

func (e *InvalidValueError) Error() string {
	return "field type (" + e.FieldType.String() + ") does not match value type (" + e.ValueType.String() + ")"
}

// SetValueOf set a field of a pointer to a struct to given value
//
// If t is nil or not a pointer, SetValueOf returns an InvalidTargetError.
// If field of t is unknown or cannot be set SetValueOf returns an InvalidFieldError
// If value is not matching the type of the field of t SetValueOf returns an InvalidValueError
func SetValueOf(t interface{}, field string, value interface{}) error {
	tv := reflect.ValueOf(t)
	if tv.Kind() != reflect.Ptr || tv.IsNil() {
		return &InvalidTargetError{Type: reflect.TypeOf(t)}
	}
	tve := tv.Elem()
	if tve.Kind() != reflect.Struct {
		return &InvalidTargetError{Type: tve.Type()}
	}
	f := tve.FieldByName(field)
	if !f.IsValid() || !f.CanSet() {
		return &InvalidFieldError{Type: tve.Type(), Field: field}
	}
	vv := reflect.ValueOf(value)
	if vv.Kind() != f.Kind() {
		return &InvalidValueError{FieldType: f.Type(), ValueType: vv.Type()}
	}
	f.Set(vv)
	return nil
}
