package util

const MINFLOAT = 0.000001
const ROUGHTFLOAT = 0.05

func IsEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return f1-f2 < MINFLOAT
	} else {
		return f2-f1 < MINFLOAT
	}
}

func IsLTEqual(f1, f2 float64) bool {
	if f1 > f2 {
		if f1-f2 < MINFLOAT {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func IsGTEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return true
	} else {
		if f2-f1 < MINFLOAT {
			return true
		} else {
			return false
		}
	}
}

func IsLess(f1, f2 float64) bool {
	if IsEqual(f1, f2) {
		return false
	}

	return f1 < f2
}

func IsGreater(f1, f2 float64) bool {
	if IsEqual(f1, f2) {
		return false
	}

	return f1 > f2
}

func IsRoughlyEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return f1-f2 < ROUGHTFLOAT
	} else {
		return f2-f1 < ROUGHTFLOAT
	}
}

func IsRoughlyLTEqual(f1, f2 float64) bool {
	if f1 > f2 {
		if f1-f2 < ROUGHTFLOAT {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
