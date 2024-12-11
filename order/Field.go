package order

import (
	"fmt"
)

type Field struct {
	Type   string // Y / M / etc...
	Number int
}

func (f *Field) String() string {
	fieldStr := fmt.Sprintf("%s:%d,12:00-13:00;", f.Type, f.Number)

	return fieldStr
}
