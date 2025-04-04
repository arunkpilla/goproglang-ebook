package conv

type Kilos float64
type Pounds float64
type Meters float64
type Feet float64

func KtoP(k Kilos) Pounds { return Pounds(2.20462 * k) }
func PtoK(p Pounds) Kilos { return Kilos(p / 2.20462) }
func MtoF(m Meters) Feet  { return Feet(m * 3.28084) }
func FtoM(f Feet) Meters  { return Meters(f / 3.28084) }
