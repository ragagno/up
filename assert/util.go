package assert

import (
	"reflect"
	"strings"
)

// isNil returns true if value v is nil.
func isNil(v interface{}) bool {
	if v == nil {
		return true
	} else { // a `interface{}` that returns false on an equality check against nil isn't necessarily not nil and requires further attention
		switch reflect.TypeOf(v).Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
			return reflect.ValueOf(v).IsNil()
		default:
			return false
		}
	}
}

// getLength returns the length of collection c.
// Panics if c is not a collection.
func getLength(c interface{}) int {
	if reflect.TypeOf(c).Kind() != reflect.String && reflect.TypeOf(c).Kind() != reflect.Array && reflect.TypeOf(c).Kind() != reflect.Slice && reflect.TypeOf(c).Kind() != reflect.Map {
		panic("not a collection")
	}

	return reflect.ValueOf(c).Len()
}

// contains returns true if container c contains element e.
// Panics if c is not a container.
// Panics if c does not contain values of the same type as e.
func contains(c interface{}, e interface{}) bool {
	if isNil(c) {
		panic("nil container")
	}

	switch reflect.TypeOf(c).Kind() {
	case reflect.Array, reflect.Slice:
		if reflect.TypeOf(c).Elem() != reflect.TypeOf(e) {
			panic("different types")
		}

		for i := 0; i < reflect.ValueOf(c).Len(); i++ {
			if reflect.DeepEqual(reflect.ValueOf(c).Index(i).Interface(), e) {
				return true
			}
		}
	case reflect.Map:
		if reflect.TypeOf(c).Elem() != reflect.TypeOf(e) {
			panic("different types")
		}

		for _, k := range reflect.ValueOf(c).MapKeys() {
			if reflect.DeepEqual(reflect.ValueOf(c).MapIndex(k).Interface(), e) {
				return true
			}
		}
	case reflect.String:
		switch reflect.TypeOf(e).Kind() {
		case reflect.String:
			return strings.Contains(reflect.ValueOf(c).String(), reflect.ValueOf(e).String())
		case reflect.Uint8:
			return strings.ContainsRune(reflect.ValueOf(c).String(), rune(reflect.ValueOf(e).Uint()))
		case reflect.Int32:
			return strings.ContainsRune(reflect.ValueOf(c).String(), rune(reflect.ValueOf(e).Int()))
		default:
			panic("different types")
		}
	default:
		panic("not a container")
	}

	return false
}

// hasKey returns true if map m has key k.
// Panics if the keys of m are of different type from k.
func hasKey(m interface{}, k interface{}) bool {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic("not a map")
	}

	if reflect.ValueOf(m).IsNil() {
		panic("nil map")
	}

	if reflect.TypeOf(m).Key() != reflect.TypeOf(k) {
		panic("different types")
	}

	for _, key := range reflect.ValueOf(m).MapKeys() {
		if reflect.DeepEqual(key.Interface(), k) {
			return true
		}
	}

	return false
}

// containsPair returns true if the map m contains the pair k -> v.
// Panics if the keys of m are of different type from k or if the values of m are of a different type from v.
func containsPair(m interface{}, k interface{}, v interface{}) bool {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic("not a map")
	}

	if reflect.ValueOf(m).IsNil() {
		panic("nil map")
	}

	if reflect.TypeOf(m).Key() != reflect.TypeOf(k) || reflect.TypeOf(m).Elem() != reflect.TypeOf(v) {
		panic("different types")
	}

	if !reflect.ValueOf(m).MapIndex(reflect.ValueOf(k)).IsValid() {
		return false
	}

	return reflect.DeepEqual(reflect.ValueOf(m).MapIndex(reflect.ValueOf(k)).Interface(), v)
}

// panics returns true if the function f causes a panic.
// Panics if f is nil.
func panics(f func()) (p bool) {
	if f == nil {
		panic("nil function")
	}

	p = true

	defer func() { recover() }()

	f()

	p = false

	return p
}

// panicsWithValue returns true if the function f causes a panic with value e.
// Panics if f is nil.
func panicsWithValue(f func(), e interface{}) (v bool) {
	if f == nil {
		panic("nil function")
	}

	v = true

	defer func() {
		v = v && reflect.DeepEqual(recover(), e)
	}()

	f()

	v = false

	return v
}
