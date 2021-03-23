package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type List []string

func (l *List) Scan(value interface{}) error {
	val, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("wrong type for List", value))
	}
	*l = List(strings.Split(string(val), ","))
	return nil
}

func (l List) Value() (driver.Value, error) {
	if len(l) == 0 {
		return nil, nil
	}
	return []byte(strings.Join(l, ",")), nil
}
