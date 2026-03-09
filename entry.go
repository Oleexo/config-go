package config

import "fmt"

type Entry struct {
	value  any
	kind   ValueKind
	exists bool
}

func NewString(value string) Entry {
	return Entry{
		exists: true,
		kind:   KindString,
		value:  value,
	}
}

func NewInt(value int64) Entry {
	return Entry{
		exists: true,
		kind:   KindInt,
		value:  value,
	}
}

func NewFloat(value float64) Entry {
	return Entry{
		exists: true,
		kind:   KindFloat,
		value:  value,
	}
}

func NewBool(value bool) Entry {
	return Entry{
		exists: true,
		kind:   KindBool,
		value:  value,
	}
}

func Empty() Entry {
	return Entry{
		exists: false,
	}
}

func (e Entry) Exists() bool {
	return e.exists
}

func (e Entry) Kind() ValueKind {
	return e.kind
}

func (e Entry) String() (string, error) {
	if !e.exists {
		return "", ErrKeyNotFound
	}
	v, ok := e.value.(string)
	if !ok {
		return "", fmt.Errorf("%w: expected %s got %s", ErrTypeMismatch, KindString, e.kind)
	}
	return v, nil
}

func (e Entry) Int() (int64, error) {
	if !e.exists {
		return 0, ErrKeyNotFound
	}
	v, ok := e.value.(int64)
	if !ok {
		return 0, fmt.Errorf("%w: expected %s got %s", ErrTypeMismatch, KindInt, e.kind)
	}
	return v, nil
}

func (e Entry) Float() (float64, error) {
	if !e.exists {
		return 0, ErrKeyNotFound
	}
	v, ok := e.value.(float64)
	if !ok {
		return 0, fmt.Errorf("%w: expected %s got %s", ErrTypeMismatch, KindFloat, e.kind)
	}
	return v, nil
}

func (e Entry) Bool() (bool, error) {
	if !e.exists {
		return false, ErrKeyNotFound
	}
	v, ok := e.value.(bool)
	if !ok {
		return false, fmt.Errorf("%w: expected %s got %s", ErrTypeMismatch, KindBool, e.kind)
	}
	return v, nil
}
