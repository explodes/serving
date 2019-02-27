package expz

import "errors"

const (
	flagTypeUnknown flagType = iota
	flagTypeInt64
	flagTypeString
	flagTypeBool
	flagTypeFloat64
)

type flagType uint8

func (m *Flag) flagType() (flagType, error) {
	switch m.Flag.(type) {
	case *Flag_I64:
		return flagTypeInt64, nil
	case *Flag_Bool:
		return flagTypeBool, nil
	case *Flag_String_:
		return flagTypeString, nil
	case *Flag_F64:
		return flagTypeFloat64, nil
	default:
		return flagTypeUnknown, errors.New("unknown flag type")
	}
}

func (m *Flag) Int64Value(def int64) int64 {
	if f, ok := m.Flag.(*Flag_I64); ok {
		return f.I64
	}
	return def
}

func (m *Flag) Float64Value(def float64) float64 {
	if f, ok := m.Flag.(*Flag_F64); ok {
		return f.F64
	}
	return def
}

func (m *Flag) BoolValue(def bool) bool {
	if f, ok := m.Flag.(*Flag_Bool); ok {
		return f.Bool
	}
	return def
}

func (m *Flag) StringValue(def string) string {
	if f, ok := m.Flag.(*Flag_String_); ok {
		return f.String_
	}
	return def
}
