package module

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockFilter struct {
	mock.Mock
}

func (m *MockFilter) Check(a int) bool {
	args := m.Called(a)
	return args.Bool(0)
}

type MockCalculator struct {
	mock.Mock
}

func (m *MockCalculator) Add(a int, b int) int {
	args := m.Called(a)
	return args.Int(0)
}

func (m *MockCalculator) Multiply(a int, b int) int {
	args := m.Called(a)
	return args.Int(0)
}

func (m *MockCalculator) IsPrime(a int) bool {
	args := m.Called(a)
	return args.Bool(0)
}

func TestCalcMethod_Add(t *testing.T) {
	mf := new(MockFilter)
	mf.On("Check", 5).Return(false)

	clc := calculator{
		F: mf,
	}
	result := clc.Add(5, 9)

	t.Log(result)
	mf.AssertExpectations(t)
}

func TestCalcMethod_Multiply(t *testing.T) {
	mc := new(MockCalculator)
	clc := calculator{
		M: mc,
	}

	mc.On("IsPrime", 5).Return(false)

	result := clc.Multiply(5, 2)

	t.Log(result)
	mc.AssertExpectations(t)

}
