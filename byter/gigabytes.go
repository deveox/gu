package byter

import (
	"fmt"
	"math"

	"github.com/deveox/gu/bounds"
)

type Gigabytes uint64

func (b Gigabytes) ToUint() uint64 {
	return uint64(b)
}

func (b Gigabytes) ToBytes() Bytes {
	return Bytes(b.B())
}

func (b Gigabytes) B() uint64 {
	return uint64(b * 1024 * 1024 * 1024)
}
func (b Gigabytes) KB() uint64 {
	return uint64(b * 1024 * 1024)
}
func (b Gigabytes) MB() uint64 {
	return uint64(b * 1024)
}

func (b Gigabytes) TB() uint64 {
	return uint64(b / 1024)
}

func (b Gigabytes) To(unit Unit) uint64 {
	switch unit {
	case BUnit:
		return b.B()
	case KBUnit:
		return b.KB()
	case MBUnit:
		return b.MB()
	case GBUnit:
		return b.ToUint()
	case TBUnit:
		return b.TB()
	}
	panic(fmt.Sprintf("Can't convert from Gigabytes, invalid unit:%s", unit))
}

func (b Gigabytes) String() string {
	return fmt.Sprintf("%d GB", b)
}

func GigabytesFrom[T bounds.Integer](value T, unit Unit) Gigabytes {
	switch unit {
	case BUnit:
		return Gigabytes(value) / 1024 / 1024 / 1024
	case KBUnit:
		return Gigabytes(value) / 1024 / 1024
	case MBUnit:
		return Gigabytes(value) / 1024
	case GBUnit:
		return Gigabytes(value)
	case TBUnit:
		return Gigabytes(value) * 1024
	}
	panic(fmt.Sprintf("Can't convert from gigabytes. Invalid unit:%s", unit))
}
func GigabytesFromFloat[T bounds.Float](value T, unit Unit) Gigabytes {
	i := math.Round(float64(value))

	switch unit {
	case BUnit:
		return Gigabytes(i) / 1024 / 1024 / 1024
	case KBUnit:
		return Gigabytes(i) / 1024 / 1024
	case MBUnit:
		return Gigabytes(i) / 1024
	case GBUnit:
		return Gigabytes(i)
	case TBUnit:
		return Gigabytes(i) * 1024
	}
	panic(fmt.Sprintf("Can't convert to gigabytes. Invalid unit:%s", unit))
}
