package concepts

type Calc struct {
	X, Y float64
}

func (c Calc) Example() float64 {
	return c.X + c.Y
}