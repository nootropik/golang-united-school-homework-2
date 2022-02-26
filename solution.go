package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type side int

const (
	SidesTriangle side = 3
	SidesSquare   side = 4
	SidesCircle   side = 0
	Pi                 = 3.14
)

func CalcSquare(sideLen float64, sidesNum side) float64 {
	switch sidesNum {
	case 3:
		CalcSquare(3, SidesTriangle)
	case 4:
		CalcSquare(4, SidesSquare)
	case 0:
		CalcSquare(2, SidesCircle)
	default
		return 0
	}
}
