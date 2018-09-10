package module

type Mathematics interface {
	Add(a int, b int) int
	Multiply(a int, b int) int
	IsPrime(a int) bool
}

type calculator struct {
	M Mathematics
	F Filter
}

func (c *calculator) Add(a int, b int) int {
	checked := c.F.Check(a)
	if checked {
		return a + b
	} else {
		return 0
	}
}

func (c *calculator) Multiply(a int, b int) int {
	prime := c.M.IsPrime(a)
	if prime {
		return a * b
	} else {
		return 1
	}
}

func (c *calculator) IsPrime(a int) bool {
	if a == 5 {
		return true
	} else {
		return false
	}
}
