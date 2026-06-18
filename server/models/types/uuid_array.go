package types

import (
	"database/sql/driver"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value any) error {

	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("invalid uuid array")
	}

	str = strings.TrimPrefix(str,"{")
	str = strings.TrimSuffix(str,"}")
	parts := strings.Split(str,",")

	*a = make(UUIDArray, 0, len(parts))

	for _,p := range parts {
		p = strings.TrimSpace(strings.Trim(p,`"`))
		if p == "" { continue }
		uuid,err := uuid.Parse(p)
		if err != nil {
			return err
		}
		*a = append(*a, uuid)
	}
	return nil
}

func (a UUIDArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	postgresFormat := make([]string, 0, len(a))
	for _,value := range a {
		postgresFormat = append(postgresFormat, `"` + value.String() + `"`)
	}

	return `{` + strings.Join(postgresFormat,",") + `}` , nil
}

func (UUIDArray) GormDataType() string {
	return "uuid[]"
}