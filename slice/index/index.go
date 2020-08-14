package index

// Bool returns the index of value in slice, or -1 if it doesn't exist.
func Bool(slice []bool, value bool) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Int returns the index of value in slice, or -1 if it doesn't exist.
func Int(slice []int, value int) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Int8 returns the index of value in slice, or -1 if it doesn't exist.
func Int8(slice []int8, value int8) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Int16 returns the index of value in slice, or -1 if it doesn't exist.
func Int16(slice []int16, value int16) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Int32 returns the index of value in slice, or -1 if it doesn't exist.
func Int32(slice []int32, value int32) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Int64 returns the index of value in slice, or -1 if it doesn't exist.
func Int64(slice []int64, value int64) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Uint returns the index of value in slice, or -1 if it doesn't exist.
func Uint(slice []uint, value uint) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Uint8 returns the index of value in slice, or -1 if it doesn't exist.
func Uint8(slice []uint8, value uint8) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Uint16 returns the index of value in slice, or -1 if it doesn't exist.
func Uint16(slice []uint16, value uint16) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Uint32 returns the index of value in slice, or -1 if it doesn't exist.
func Uint32(slice []uint32, value uint32) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Uint64 returns the index of value in slice, or -1 if it doesn't exist.
func Uint64(slice []uint64, value uint64) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Uintptr returns the index of value in slice, or -1 if it doesn't exist.
func Uintptr(slice []uintptr, value uintptr) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Float32 returns the index of value in slice, or -1 if it doesn't exist.
func Float32(slice []float32, value float32) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Float64 returns the index of value in slice, or -1 if it doesn't exist.
func Float64(slice []float64, value float64) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Complex64 returns the index of value in slice, or -1 if it doesn't exist.
func Complex64(slice []complex64, value complex64) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Complex128 returns the index of value in slice, or -1 if it doesn't exist.
func Complex128(slice []complex128, value complex128) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Byte returns the index of value in slice, or -1 if it doesn't exist.
func Byte(slice []byte, value byte) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Rune returns the index of value in slice, or -1 if it doesn't exist.
func Rune(slice []rune, value rune) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// String returns the index of value in slice, or -1 if it doesn't exist.
func String(slice []string, value string) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}

// Error returns the index of value in slice, or -1 if it doesn't exist.
func Error(slice []error, value error) int {
	for i, e := range slice {
		if value == e {
			return i
		}
	}

	return -1
}
