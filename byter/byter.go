package byter

import (
	"bytes"
	"fmt"
	"io"

	"github.com/deveox/gu/bounds"
)

// FromString returns a byte slice from a  stream.
func FromStream(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

type Bytes uint64

const (
	KB Bytes = 1024
	MB Bytes = KB * 1024
	GB Bytes = MB * 1024
	TB Bytes = GB * 1024
)

func (b Bytes) ToUint() uint64 {
	return uint64(b)
}

func (b Bytes) KB() uint64 {
	return uint64(b / KB)
}
func (b Bytes) MB() uint64 {
	return uint64(b / MB)
}
func (b Bytes) GB() uint64 {
	return uint64(b / GB)
}
func (b Bytes) TB() uint64 {
	return uint64(b / TB)
}

func (b Bytes) To(unit Unit) uint64 {
	switch unit {
	case BUnit:
		return b.ToUint()
	case KBUnit:
		return b.KB()
	case MBUnit:
		return b.MB()
	case GBUnit:
		return b.GB()
	case TBUnit:
		return b.TB()
	}
	panic(fmt.Sprintf("Can't convert from Bytes, invalid unit:%s", unit))
}

func (b Bytes) String() string {
	v := b
	units := []Unit{KBUnit, MBUnit, GBUnit, TBUnit}
	unit := BUnit
	for _, u := range units {
		if v < 1024 {
			break
		}
		v = v / 1024
		unit = u
	}
	return fmt.Sprintf("%d %s", v, unit)
}

type Unit string

const (
	BUnit  Unit = "B"
	KBUnit Unit = "KB"
	MBUnit Unit = "MB"
	GBUnit Unit = "GB"
	TBUnit Unit = "TB"
)

func From[T bounds.Integer](value T, unit Unit) Bytes {
	switch unit {
	case BUnit:
		return Bytes(value)
	case KBUnit:
		return Bytes(value) * KB
	case MBUnit:
		return Bytes(value) * MB
	case GBUnit:
		return Bytes(value) * GB
	case TBUnit:
		return Bytes(value) * TB
	}
	panic(fmt.Sprintf("Can't convert to bytes. Invalid unit:%s", unit))
}
