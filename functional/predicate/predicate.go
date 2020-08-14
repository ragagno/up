package predicate

// Bool represents a function that accepts a bool and returns a bool if the predicate is validated.
type Bool func(bool) bool

// Int represents a function that accepts a int and returns a bool if the predicate is validated.
type Int func(int) bool

// Int8 represents a function that accepts a int8 and returns a bool if the predicate is validated.
type Int8 func(int8) bool

// Int16 represents a function that accepts a int16 and returns a bool if the predicate is validated.
type Int16 func(int16) bool

// Int32 represents a function that accepts a int32 and returns a bool if the predicate is validated.
type Int32 func(int32) bool

// Int64 represents a function that accepts a int64 and returns a bool if the predicate is validated.
type Int64 func(int64) bool

// Uint represents a function that accepts a uint and returns a bool if the predicate is validated.
type Uint func(uint) bool

// Uint8 represents a function that accepts a uint8 and returns a bool if the predicate is validated.
type Uint8 func(uint8) bool

// Uint16 represents a function that accepts a uint16 and returns a bool if the predicate is validated.
type Uint16 func(uint16) bool

// Uint32 represents a function that accepts a uint32 and returns a bool if the predicate is validated.
type Uint32 func(uint32) bool

// Uint64 represents a function that accepts a uint64 and returns a bool if the predicate is validated.
type Uint64 func(uint64) bool

// Uintptr represents a function that accepts a uintptr and returns a bool if the predicate is validated.
type Uintptr func(uintptr) bool

// Float32 represents a function that accepts a float32 and returns a bool if the predicate is validated.
type Float32 func(float32) bool

// Float64 represents a function that accepts a float64 and returns a bool if the predicate is validated.
type Float64 func(float64) bool

// Complex64 represents a function that accepts a complex64 and returns a bool if the predicate is validated.
type Complex64 func(complex64) bool

// Complex128 represents a function that accepts a complex128 and returns a bool if the predicate is validated.
type Complex128 func(complex128) bool

// Byte represents a function that accepts a byte and returns a bool if the predicate is validated.
type Byte func(byte) bool

// Rune represents a function that accepts a rune and returns a bool if the predicate is validated.
type Rune func(rune) bool

// String represents a function that accepts a string and returns a bool if the predicate is validated.
type String func(string) bool

// Error represents a function that accepts a error and returns a bool if the predicate is validated.
type Error func(error) bool

// Interface represents a function that accepts a interface{} and returns a bool if the predicate is validated.
type Interface func(interface{}) bool

// Bools represents a function that accepts a []bool and returns a bool if the predicate is validated.
type Bools func([]bool) bool

// Ints represents a function that accepts a []int and returns a bool if the predicate is validated.
type Ints func([]int) bool

// Int8s represents a function that accepts a []int8 and returns a bool if the predicate is validated.
type Int8s func([]int8) bool

// Int16s represents a function that accepts a []int16 and returns a bool if the predicate is validated.
type Int16s func([]int16) bool

// Int32s represents a function that accepts a []int32 and returns a bool if the predicate is validated.
type Int32s func([]int32) bool

// Int64s represents a function that accepts a []int64 and returns a bool if the predicate is validated.
type Int64s func([]int64) bool

// Uints represents a function that accepts a []uint and returns a bool if the predicate is validated.
type Uints func([]uint) bool

// Uint8s represents a function that accepts a []uint8 and returns a bool if the predicate is validated.
type Uint8s func([]uint8) bool

// Uint16s represents a function that accepts a []uint16 and returns a bool if the predicate is validated.
type Uint16s func([]uint16) bool

// Uint32s represents a function that accepts a []uint32 and returns a bool if the predicate is validated.
type Uint32s func([]uint32) bool

// Uint64s represents a function that accepts a []uint64 and returns a bool if the predicate is validated.
type Uint64s func([]uint64) bool

// Uintptrs represents a function that accepts a []uintptr and returns a bool if the predicate is validated.
type Uintptrs func([]uintptr) bool

// Float32s represents a function that accepts a []float32 and returns a bool if the predicate is validated.
type Float32s func([]float32) bool

// Float64s represents a function that accepts a []float64 and returns a bool if the predicate is validated.
type Float64s func([]float64) bool

// Complex64s represents a function that accepts a []complex64 and returns a bool if the predicate is validated.
type Complex64s func([]complex64) bool

// Complex128s represents a function that accepts a []complex128 and returns a bool if the predicate is validated.
type Complex128s func([]complex128) bool

// Bytes represents a function that accepts a []byte and returns a bool if the predicate is validated.
type Bytes func([]byte) bool

// Runes represents a function that accepts a []rune and returns a bool if the predicate is validated.
type Runes func([]rune) bool

// Strings represents a function that accepts a []string and returns a bool if the predicate is validated.
type Strings func([]string) bool

// Errors represents a function that accepts a []error and returns a bool if the predicate is validated.
type Errors func([]error) bool

// Interfaces represents a function that accepts a []interface{} and returns a bool if the predicate is validated.
type Interfaces func([]interface{}) bool
