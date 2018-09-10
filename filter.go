package module

import (
	"fmt"
)

// go:generate goautomock -template=testify Filter
type Filter interface {
	Check(a int) bool
}

type ModFilter struct {
	fl interface{}
}

func (fl *ModFilter) Check(a int) bool {
	fmt.Printf("filter default\n")
	if a > 10 {
		return false
	}
	return true
}
