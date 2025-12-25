package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("failed to parse UUIDArray: unsupported data type")
	}

	// str = str[1 : len(str)-1] // Remove curly braces
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")

	parts := strings.Split(str, ",")

	*a = make(UUIDArray, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(strings.Trim(part, `"`)) // Remove quotes and whitespace
		if part == "" {
			continue
		}
		id, err := uuid.Parse(part)
		if err != nil {
			return fmt.Errorf("Invalid UUID in Array: %v", err)
		}
		*a = append(*a, id)
	}

	return nil
}

// Expectation => {"54145455-r24-ere12-5434","54245455-r24-ere12-5434","54345455-r24-ere12-5434"}
func (a UUIDArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	postgreFormat := make([]string, 0, len(a))
	for _, value := range a {
		postgreFormat = append(postgreFormat, fmt.Sprintf(`"%s"`, value.String()))
	}

	return "{" + strings.Join(postgreFormat, ",") + "}", nil
}

func (UUIDArray) GormDataType() string {
	return "uuid[]"
}
