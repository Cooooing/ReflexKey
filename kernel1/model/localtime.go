package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// LocalTime 是一个自定义的时间类型
type LocalTime time.Time

// MarshalJSON 实现 json.Marshaler 接口
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	if time.Time(*t).IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		*t = LocalTime(time.Time{})
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str[1:len(str)-1])
	if err != nil {
		return err
	}
	*t = LocalTime(parsedTime)
	return nil
}

// Value 实现 driver.Valuer 接口
func (t *LocalTime) Value() (driver.Value, error) {
	if time.Time(*t).IsZero() {
		return nil, nil
	}
	return time.Time(*t).Format("2006-01-02 15:04:05"), nil
}

// Scan 实现 sql.Scanner 接口
func (t *LocalTime) Scan(value any) error {
	switch v := value.(type) {
	case time.Time:
		*t = LocalTime(v)
	case string:
		if v == "" {
			*t = LocalTime(time.Time{})
			return nil
		}
		parsedTime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = LocalTime(parsedTime)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}
