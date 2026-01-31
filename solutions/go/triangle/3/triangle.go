package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	isTriangle := validateTriangle(a, b, c)
	if !isTriangle {
		return Kind(NaT)
	}

	isEquilateral := isTriangleEquilateral(a, b, c)
	if isEquilateral {
		return Kind(Equ)
	}

	isIsosceles := isTriangleIsosceles(a, b, c)
	if isIsosceles {
		return Kind(Iso)
	}

	return Kind(Sca)
}

func validateTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return a+b >= c && a+c >= b && b+c >= a
}

func isTriangleEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isTriangleIsosceles(a, b, c float64) bool {
	return a == b || a == c || b == c
}
