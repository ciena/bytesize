# Byte Size Type
This repository provides a type `ByteSize` which represents a cardinal number
of bytes stored as a `uint64`, but can be parsed from a `string` of the form of
a `uint64` followed by a `unit` where a unit can be one of the following.

| UNIT | NAME | VALUE | | UNIT | NAME | VALUE |
| --- | --- | --- | --- | --- | --- | --- |
| B | Byte | 1 | | B | BiByte | 1 |
| K, KB | Kilobyte | 1000<sup>1</sup> | | KiB | Kibibyte | 1024<sup>1</sup> |
| M, MB | Megabyte | 1000<sup>2</sup> | | MiB | Mebibyte | 1024<sup>2</sup> |
| G, GB | Gigabyte | 1000<sup>3</sup> | | GiB | Gibibyte | 1024<sup>3</sup> |
| T, TB | Terabyte | 1000<sup>4</sup> | | TiB | Tebibyte | 1024<sup>4</sup> |
| P, PB | Petabyte | 1000<sup>5</sup> | | TiB | Pebibyte | 1024<sup>5</sup> |
| E, EB | Exabyte  | 1000<sup>6</sup> | | EiB | Exbibyte | 1024<sup>6</sup> |

## String To ByteSize
This library supports functions to parse a string explicitly as base 2 (binary)
or explicitly as base 10 (decimal) units of data. Additionally, a method to
parse the input string taking into account the unit specifier as specified in
the above table is provided.

## ByteSize To String
This library supports functions to convert a `ByteSize` value to its string
equivalent. There are two functions to explicitly convert as either base 2
(binary) or base 10 (decimal). Additinally, a `String()` function is provided
that attempt to output a string that is the closest match, i.e. if the `ByteSize`
value is `1024` bytes it will output the string `1KiB` where as if the `ByteSize`
value is `1000` bytes it will output the string `1KB`. If the `ByteSize` value
cannot be converted to a whole unit multiplier then it will be output as Bytes,
i.e. `1053` bytes will be output as the string `1053B`.
