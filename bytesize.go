package bytesize

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ByteSize represents the number of bytes as converted to
// or from a size specification
type ByteSize uint64

const (
	BiByte   ByteSize = 1
	Kibibyte          = BiByte * 1024
	Mebibyte          = Kibibyte * 1024
	Gibibyte          = Mebibyte * 1024
	Tebibyte          = Gibibyte * 1024
	Pebibyte          = Tebibyte * 1024
	Exbibyte          = Pebibyte * 1024
)

// Binary or base 2 size definitions

const (
	Byte     ByteSize = 1
	Kilobyte          = Byte * 1000
	Megabyte          = Kilobyte * 1000
	Gigabyte          = Megabyte * 1000
	Terabyte          = Gigabyte * 1000
	Petabyte          = Terabyte * 1000
	Exabyte           = Petabyte * 1000
)

// Decimal or base 10 size definitions

var sizeParser = regexp.MustCompile(`^\s*([0-9]+)\s*([EPTGMK]?)(IB|B)?\s*$`)

// BinaryString convert the binary (base 2) byte size value to a string
// containing both the value and the unit designation. values which
// cannot be represented as a single value and unit are returned as
// bytes.
func (b ByteSize) BinaryString() string {

	if uint64(b)%uint64(Exbibyte) == 0 {
		return fmt.Sprintf("%dEiB", uint64(b)/uint64(Exbibyte))
	}
	if uint64(b)%uint64(Pebibyte) == 0 {
		return fmt.Sprintf("%dPiB", uint64(b)/uint64(Pebibyte))
	}
	if uint64(b)%uint64(Tebibyte) == 0 {
		return fmt.Sprintf("%dTiB", uint64(b)/uint64(Tebibyte))
	}
	if uint64(b)%uint64(Gibibyte) == 0 {
		return fmt.Sprintf("%dGiB", uint64(b)/uint64(Gibibyte))
	}
	if uint64(b)%uint64(Mebibyte) == 0 {
		return fmt.Sprintf("%dMiB", uint64(b)/uint64(Mebibyte))
	}
	if uint64(b)%uint64(Kibibyte) == 0 {
		return fmt.Sprintf("%dKiB", uint64(b)/uint64(Kibibyte))
	}
	return fmt.Sprintf("%dB", uint64(b))
}

// DecimalString convert the decimal (base 10) byte size value to a string
// containing both the value and the unit designation. values which
// cannot be represented as a single value and unit are returned as
// bytes.
func (b ByteSize) DecimalString() string {

	if uint64(b)%uint64(Exabyte) == 0 {
		return fmt.Sprintf("%dEB", uint64(b)/uint64(Exabyte))
	}
	if uint64(b)%uint64(Petabyte) == 0 {
		return fmt.Sprintf("%dPB", uint64(b)/uint64(Petabyte))
	}
	if uint64(b)%uint64(Terabyte) == 0 {
		return fmt.Sprintf("%dTB", uint64(b)/uint64(Terabyte))
	}
	if uint64(b)%uint64(Gigabyte) == 0 {
		return fmt.Sprintf("%dGB", uint64(b)/uint64(Gigabyte))
	}
	if uint64(b)%uint64(Megabyte) == 0 {
		return fmt.Sprintf("%dMB", uint64(b)/uint64(Megabyte))
	}
	if uint64(b)%uint64(Kilobyte) == 0 {
		return fmt.Sprintf("%dKB", uint64(b)/uint64(Kilobyte))
	}
	return fmt.Sprintf("%dB", uint64(b))
}

func (b ByteSize) String() string {

	if uint64(b)%uint64(Exbibyte) == 0 {
		return fmt.Sprintf("%dEiB", uint64(b)/uint64(Exbibyte))
	}
	if uint64(b)%uint64(Exabyte) == 0 {
		return fmt.Sprintf("%dEB", uint64(b)/uint64(Exabyte))
	}
	if uint64(b)%uint64(Pebibyte) == 0 {
		return fmt.Sprintf("%dPiB", uint64(b)/uint64(Pebibyte))
	}
	if uint64(b)%uint64(Petabyte) == 0 {
		return fmt.Sprintf("%dPB", uint64(b)/uint64(Petabyte))
	}
	if uint64(b)%uint64(Tebibyte) == 0 {
		return fmt.Sprintf("%dTiB", uint64(b)/uint64(Tebibyte))
	}
	if uint64(b)%uint64(Terabyte) == 0 {
		return fmt.Sprintf("%dTB", uint64(b)/uint64(Terabyte))
	}
	if uint64(b)%uint64(Gibibyte) == 0 {
		return fmt.Sprintf("%dGiB", uint64(b)/uint64(Gibibyte))
	}
	if uint64(b)%uint64(Gigabyte) == 0 {
		return fmt.Sprintf("%dGB", uint64(b)/uint64(Gigabyte))
	}
	if uint64(b)%uint64(Mebibyte) == 0 {
		return fmt.Sprintf("%dMiB", uint64(b)/uint64(Mebibyte))
	}
	if uint64(b)%uint64(Megabyte) == 0 {
		return fmt.Sprintf("%dMB", uint64(b)/uint64(Megabyte))
	}
	if uint64(b)%uint64(Kibibyte) == 0 {
		return fmt.Sprintf("%dKiB", uint64(b)/uint64(Kibibyte))
	}
	if uint64(b)%uint64(Kilobyte) == 0 {
		return fmt.Sprintf("%dKB", uint64(b)/uint64(Kilobyte))
	}
	return fmt.Sprintf("%dB", uint64(b))
}

// ParseBinarySize parse a byte size specification using binary
// (base 2) unit constants. A specification is a single value
// and unit combination.
func ParseBinarySize(size string) (ByteSize, error) {
	parts := sizeParser.FindAllStringSubmatch(strings.ToUpper(size), -1)
	if len(parts) == 0 {
		return 0, fmt.Errorf("size: invalid size '%s'", size)
	}

	value, err := strconv.ParseUint(parts[0][1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("size: invalid size '%s'", size)
	}

	unit := BiByte
	switch parts[0][2] {
	case "E":
		unit = Exbibyte
	case "P":
		unit = Pebibyte
	case "T":
		unit = Tebibyte
	case "G":
		unit = Gibibyte
	case "M":
		unit = Mebibyte
	case "K":
		unit = Kibibyte
	default:
	}
	return ByteSize(value) * unit, nil
}

// ParseDecimalSize parse a byte size specification using decimal
// (base 10) unit constants. A specification is a single value
// and unit combination.
func ParseDecimalSize(size string) (ByteSize, error) {
	parts := sizeParser.FindAllStringSubmatch(strings.ToUpper(size), -1)
	if len(parts) == 0 {
		return 0, fmt.Errorf("size: invalid size '%s'", size)
	}

	value, err := strconv.ParseUint(parts[0][1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("size: invalid size '%s'", size)
	}

	unit := Byte
	switch parts[0][2] {
	case "E":
		unit = Exabyte
	case "P":
		unit = Petabyte
	case "T":
		unit = Terabyte
	case "G":
		unit = Gigabyte
	case "M":
		unit = Megabyte
	case "K":
		unit = Kilobyte
	default:
	}
	return ByteSize(value) * unit, nil
}

func pow(base, power uint64) uint64 {
	val := uint64(1)
	for i := uint64(0); i < power; i++ {
		val *= base
	}
	return val
}

// ParseSize parse a size specification as either a binary (base 2)
// or decimnal (base 10) value using strict adherance to the unit
// specification to determine the type (binary or decimal) of the
// value.
//
// Format: X or XB is considered decimal
// Format: XiB is considered binary
//
// Capitialization is not considered
func ParseSize(size string) (ByteSize, error) {

	parts := sizeParser.FindAllStringSubmatch(strings.ToUpper(size), -1)
	if len(parts) == 0 {
		return 0, fmt.Errorf("size: invalid size '%s'", size)
	}

	value, err := strconv.ParseUint(parts[0][1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("size: invalid size '%s'", size)
	}

	factor := uint64(1000)
	if parts[0][3] == "IB" {
		factor = 1024
	}
	switch parts[0][2] {
	case "E":
		factor = pow(factor, 6)
	case "P":
		factor = pow(factor, 5)
	case "T":
		factor = pow(factor, 4)
	case "G":
		factor = pow(factor, 3)
	case "M":
		factor = pow(factor, 2)
	case "K":
		factor = pow(factor, 1)
	case "B":
		fallthrough
	default:
		factor = pow(factor, 0)
	}
	return ByteSize(value * factor), nil
}
