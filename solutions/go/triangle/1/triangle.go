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
	result := true

	if a+b+c == 0 {
		result = false
	} else if a+b < c {
		result = false
	} else if a+c < b {
		result = false
	} else if b+c < a {
		result = false
	}

	return result
}

func isTriangleEquilateral(a, b, c float64) bool {
	result := true

	if a-b != 0 {
		result = false
	} else if a-c != 0 {
		result = false
	} else if b-c != 0 {
		result = false
	}

	return result
}

func isTriangleIsosceles(a, b, c float64) bool {

	result := false

	if a-b == 0 {
		result = true
	} else if a-c == 0 {
		result = true
	} else if b-c == 0 {
		result = true
	}

	return result
}
