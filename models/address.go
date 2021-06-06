package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Address struct {
	Name        string `json:"name"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	Town        string `json:"town"`
}

// Make the Address struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
// Ref: https://www.alexedwards.net/blog/using-postgresql-jsonb
func (a Address) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Address struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
// Ref: https://www.alexedwards.net/blog/using-postgresql-jsonb
func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (a *Address) IsValid() bool {
	return a.Name != "" && a.Street != "" && a.HouseNumber != "" && a.Town != ""
}
