package p

import (
	"testing"
	)

//go:proto B=Builtin L=B.lower
func TestUint(t *testing.T) {
	var x uint
	ptr := Uint(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestUint8(t *testing.T) {
	var x uint8
	ptr := Uint8(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestUint16(t *testing.T) {
	var x uint16
	ptr := Uint16(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestUint32(t *testing.T) {
	var x uint32
	ptr := Uint32(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestUint64(t *testing.T) {
	var x uint64
	ptr := Uint64(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestInt(t *testing.T) {
	var x int
	ptr := Int(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestInt8(t *testing.T) {
	var x int8
	ptr := Int8(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestInt16(t *testing.T) {
	var x int16
	ptr := Int16(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestInt32(t *testing.T) {
	var x int32
	ptr := Int32(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestInt64(t *testing.T) {
	var x int64
	ptr := Int64(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestFloat32(t *testing.T) {
	var x float32
	ptr := Float32(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestFloat64(t *testing.T) {
	var x float64
	ptr := Float64(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
func TestBool(t *testing.T) {
	var x bool
	ptr := Bool(x)

	if x != *ptr {
		t.Fatalf("unexpected value %v, expected %v", *ptr, x)
	}
}
