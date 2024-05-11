package length

// Unit is a base type for measures
type Unit float64

func (u Unit) Add(v Unit) Unit {
	u += v
	return u
}

func (u Unit) Sub(v Unit) Unit {
	u -= v
	return u
}

func (u Unit) Micrometers() float64 {
	return float64(u) * 1e6
}

func (u Unit) Millimeters() float64 {
	return float64(u) * 1e3
}

func (u Unit) Cantimeters() float64 {
	return float64(u) * 1e2
}

func (u Unit) Decimeters() float64 {
	return float64(u) * 10
}

func (u Unit) Meters() float64 {
	return float64(u)
}

func (u Unit) Kilometers() float64 {
	return float64(u) / 1e3
}

var (
	Micrometer Unit = 1e-6
	Millimeter Unit = 1e-3
	Cantimeter Unit = 1e-2
	Decimeter  Unit = 1e-1
	Meter      Unit = 1
	Kilometer  Unit = 1e3
)
