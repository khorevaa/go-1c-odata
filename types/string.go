package types

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"reflect"
	"strings"
)

// Type for Edm.String
type String string

// Maps String to the graphql scalar type in the schema.
func (String) ImplementsGraphQLType(name string) bool {
	return name == "String"
}

// A custom unmarshaler for String type
func (t *String) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = String(input)
		return nil
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
}

// A custom json/graphql marshaller for String type
func (t String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

// A custom json unmarshaller for String type
func (t *String) UnmarshalJSON(b []byte) error {
	s := ""
	err := json.Unmarshal(b, &s)
	*t = String(s)
	return err
}

// Escape quotes
func (t String) Escape() string {
	return strings.Replace(string(t), "'", "''", -1)
}

// A custom marshaller to uri query format for String type
func (t *String) Query() string {
	if t == nil {
		return "''"
	}
	return fmt.Sprintf("'%s'", t.Escape())
}
