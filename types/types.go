package types

import (
    "database/sql/driver"
    "encoding/json"
    "errors"
)

// StringArray is a custom type for handling JSON arrays in MySQL
type StringArray []string

// Scan implements the sql.Scanner interface for StringArray
func (a *StringArray) Scan(value interface{}) error {
    // Convert []uint8 (byte slice) to StringArray
    if value == nil {
        *a = nil
        return nil
    }

    switch v := value.(type) {
    case []uint8:
        return json.Unmarshal(v, a)
    case string:
        return json.Unmarshal([]byte(v), a)
    default:
        return errors.New("unsupported data type")
    }
}

// Value implements the driver.Valuer interface for StringArray
func (a StringArray) Value() (driver.Value, error) {
    if len(a) == 0 {
        return "[]", nil
    }
    return json.Marshal(a)
}
