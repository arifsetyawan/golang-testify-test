/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package module

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"
)

// MathematicsMock mock
type MathematicsMock struct {
	mock.Mock
}

func NewMathematicsMock() *MathematicsMock {
	return &MathematicsMock{}
}

// Add mocked method
func (m *MathematicsMock) Add(p0 int, p1 int) int {

	ret := m.Called(p0, p1)

	var r0 int
	switch res := ret.Get(0).(type) {
	case nil:
	case int:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// IsPrime mocked method
func (m *MathematicsMock) IsPrime(p0 int) bool {

	ret := m.Called(p0)

	var r0 bool
	switch res := ret.Get(0).(type) {
	case nil:
	case bool:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// Multiply mocked method
func (m *MathematicsMock) Multiply(p0 int, p1 int) int {

	ret := m.Called(p0, p1)

	var r0 int
	switch res := ret.Get(0).(type) {
	case nil:
	case int:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}
