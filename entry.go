package config

type Entry struct {
	exists     bool
	valueType  ValueType
	valueStr   *string
	valueInt   *int64
	valueFloat *float64
	valueBool  *bool
}

func NewEntryString(value string) Entry {
	return Entry{
		exists:    true,
		valueType: STRING,
		valueStr:  &value,
	}
}

func NewEntryInt(value int64) Entry {
	return Entry{
		exists:    true,
		valueType: INT,
		valueInt:  &value,
	}
}

func NewEntryFloat(value float64) Entry {
	return Entry{
		exists:     true,
		valueType:  FLOAT,
		valueFloat: &value,
	}
}

func NewEntryBool(value bool) Entry {
	return Entry{
		exists:    true,
		valueType: BOOL,
		valueBool: &value,
	}
}

func NewEntryEmpty() Entry {
	return Entry{
		exists: false,
	}
}

func (e Entry) Exists() bool {
	return e.exists
}

func (e Entry) ValueType() ValueType {
	return e.valueType
}

func (e Entry) String() string {
	return *e.valueStr
}

func (e Entry) Int() int64 {
	return *e.valueInt
}

func (e Entry) Float() float64 {
	return *e.valueFloat
}

func (e Entry) Bool() bool {
	return *e.valueBool
}
