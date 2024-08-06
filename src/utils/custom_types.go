package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (ua *UUIDArray) Scan(src interface{}) error {
	srcStr, ok := src.(string)

	if !ok {
		return errors.New("type assertion to string failed")
	}

	srcStr = strings.Trim(srcStr, "{}")

	uuidStrs := strings.Split(srcStr, ",")

	for _, uuidStr := range uuidStrs {
		u, err := uuid.Parse(uuidStr)
		if err != nil {
			return err
		}
		*ua = append(*ua, u)
	}
	return nil
}

func (ua UUIDArray) Value() (driver.Value, error) {
	var uuidStrs []string
	for _, u := range ua {
		uuidStrs = append(uuidStrs, u.String())
	}

	return fmt.Sprintf("{%s}", strings.Join(uuidStrs, ",")), nil
}
