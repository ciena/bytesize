package bytesize

import (
	"fmt"
	"testing"
)

type tSB struct {
	SVal        string
	BVal        ByteSize
	ExpectError bool
}

func checkConv(t *testing.T, f func(string) (ByteSize, error), test tSB) {
	val, err := f(test.SVal)
	if test.ExpectError && err == nil {
		t.Errorf("From '%s', expected 'error', recieved 'no error'", test.SVal)
	}
	if !test.ExpectError {
		if err != nil {
			t.Errorf("Unexpected fail: %s", err.Error())
		} else {
			if val != test.BVal {
				t.Errorf("From '%s', expected '%d' received '%d'", test.SVal, uint64(test.BVal), val)
			}
		}
	}
}

func checkString(t *testing.T, f func(ByteSize) string, test tSB) {
	val := f(test.BVal)
	if val != test.SVal {
		t.Errorf("From '%d', expected '%s' received '%s'", test.BVal, test.SVal, val)
	}
}

func TestForceBinary(t *testing.T) {
	data := []tSB{
		tSB{"10", BiByte * 10, false},
		tSB{"10b", BiByte * 10, false},
		tSB{"10B", BiByte * 10, false},
		tSB{"10g", Gibibyte * 10, false},
		tSB{"10G", Gibibyte * 10, false},
		tSB{"10Mb", Mebibyte * 10, false},
		tSB{"10KB", Kibibyte * 10, false},
		tSB{"10KiB", Kibibyte * 10, false},
		tSB{"10kIB", Kibibyte * 10, false},
		tSB{"10Tib", Tebibyte * 10, false},
		tSB{"10pIb", Pebibyte * 10, false},
		tSB{"10E", Exbibyte * 10, false},
		tSB{" 10KiB", Kibibyte * 10, false},
		tSB{"10 G", Gibibyte * 10, false},
		tSB{"\t 10TiB \t", Tebibyte * 10, false},
		tSB{"a10E", Byte, true},
		tSB{"10Gok", Byte, true},
		tSB{"10GiBs", Byte, true},
		tSB{"  a 10GiBs", Byte, true},
		tSB{"10 l GiBs", Byte, true},
	}

	for _, test := range data {
		t.Run(fmt.Sprintf("ParseBinarySize('%s')", test.SVal), func(t *testing.T) {
			checkConv(t, ParseBinarySize, test)
		})
	}
}

func TestForceDecimal(t *testing.T) {
	data := []tSB{
		tSB{"10", Byte * 10, false},
		tSB{"10b", Byte * 10, false},
		tSB{"10B", Byte * 10, false},
		tSB{"10g", Gigabyte * 10, false},
		tSB{"10G", Gigabyte * 10, false},
		tSB{"10Mb", Megabyte * 10, false},
		tSB{"10KB", Kilobyte * 10, false},
		tSB{"10KiB", Kilobyte * 10, false},
		tSB{"10kIB", Kilobyte * 10, false},
		tSB{"10Tib", Terabyte * 10, false},
		tSB{"10pIb", Petabyte * 10, false},
		tSB{"10E", Exabyte * 10, false},
		tSB{" 10KiB", Kilobyte * 10, false},
		tSB{"10 G", Gigabyte * 10, false},
		tSB{"\t 10TiB \t", Terabyte * 10, false},
		tSB{"a10E", Byte, true},
		tSB{"10Gok", Byte, true},
		tSB{"10GiBs", Byte, true},
		tSB{"  a 10GiBs", Byte, true},
		tSB{"10 l GiBs", Byte, true},
	}

	for _, test := range data {
		t.Run(fmt.Sprintf("ParseDecinalSize('%s')", test.SVal), func(t *testing.T) {
			checkConv(t, ParseDecimalSize, test)
		})
	}
}

func TestParse(t *testing.T) {
	data := []tSB{
		tSB{"10", Byte * 10, false},
		tSB{"10b", Byte * 10, false},
		tSB{"10B", Byte * 10, false},
		tSB{"10g", Gigabyte * 10, false},
		tSB{"10G", Gigabyte * 10, false},
		tSB{"10Mb", Megabyte * 10, false},
		tSB{"10KB", Kilobyte * 10, false},
		tSB{"10KiB", Kibibyte * 10, false},
		tSB{"10kIB", Kibibyte * 10, false},
		tSB{"10Tib", Tebibyte * 10, false},
		tSB{"10pIb", Pebibyte * 10, false},
		tSB{"10E", Exabyte * 10, false},
		tSB{" 10KiB", Kibibyte * 10, false},
		tSB{"10 G", Gigabyte * 10, false},
		tSB{"\t 10TiB \t", Tebibyte * 10, false},
		tSB{"a10E", Byte, true},
		tSB{"10Gok", Byte, true},
		tSB{"10GiBs", Byte, true},
		tSB{"  a 10GiBs", Byte, true},
		tSB{"10 l GiBs", Byte, true},
	}

	for _, test := range data {
		t.Run(fmt.Sprintf("ParseSize('%s')", test.SVal), func(t *testing.T) {
			checkConv(t, ParseSize, test)
		})
	}
}

func TestForceBinaryString(t *testing.T) {
	data := []tSB{
		tSB{"10B", BiByte * 10, false},
		tSB{"10GiB", Gibibyte * 10, false},
		tSB{"10MiB", Mebibyte * 10, false},
		tSB{"10KiB", Kibibyte * 10, false},
		tSB{"10TiB", Tebibyte * 10, false},
		tSB{"10PiB", Pebibyte * 10, false},
		tSB{"10EiB", Exbibyte * 10, false},
		tSB{"1000B", Kilobyte, false},
		tSB{"1000000000B", Gigabyte, false},
		tSB{"42B", ByteSize(42), false},
	}
	for _, test := range data {
		t.Run(fmt.Sprintf("BinaryString('%d')", test.BVal), func(t *testing.T) {
			checkString(t, ByteSize.BinaryString, test)
		})
	}
}

func TestForceDecimalString(t *testing.T) {
	data := []tSB{
		tSB{"10B", Byte * 10, false},
		tSB{"10GB", Gigabyte * 10, false},
		tSB{"10MB", Megabyte * 10, false},
		tSB{"10KB", Kilobyte * 10, false},
		tSB{"10TB", Terabyte * 10, false},
		tSB{"10PB", Petabyte * 10, false},
		tSB{"10EB", Exabyte * 10, false},
		tSB{"1024B", Kibibyte, false},
		tSB{"1073741824B", Gibibyte, false},
		tSB{"42B", ByteSize(42), false},
	}
	for _, test := range data {
		t.Run(fmt.Sprintf("DecimalString('%d')", test.BVal), func(t *testing.T) {
			checkString(t, ByteSize.DecimalString, test)
		})
	}
}

func TestString(t *testing.T) {
	data := []tSB{
		tSB{"10B", Byte * 10, false},
		tSB{"10GB", Gigabyte * 10, false},
		tSB{"10MB", Megabyte * 10, false},
		tSB{"10KB", Kilobyte * 10, false},
		tSB{"10TB", Terabyte * 10, false},
		tSB{"10PB", Petabyte * 10, false},
		tSB{"10EB", Exabyte * 10, false},
		tSB{"1KiB", Kibibyte, false},
		tSB{"1GiB", Gibibyte, false},
		tSB{"42B", ByteSize(42), false},
		tSB{"10B", BiByte * 10, false},
		tSB{"10GiB", Gibibyte * 10, false},
		tSB{"10MiB", Mebibyte * 10, false},
		tSB{"10KiB", Kibibyte * 10, false},
		tSB{"10TiB", Tebibyte * 10, false},
		tSB{"10PiB", Pebibyte * 10, false},
		tSB{"10EiB", Exbibyte * 10, false},
	}
	for _, test := range data {
		t.Run(fmt.Sprintf("String('%d')", test.BVal), func(t *testing.T) {
			checkString(t, ByteSize.String, test)
		})
	}
}
