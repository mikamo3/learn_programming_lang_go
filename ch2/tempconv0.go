package tempconv0

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	freezingC     Celsius = 0
	boilingC      Celsius = 100
)

func CTof(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func fToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
