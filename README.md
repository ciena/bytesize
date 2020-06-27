# Byte Size Type
This repository provides a type `ByteSize` which represents a cardinal number
of bytes stored as a `uint64`, but can be parsed from a `string` of the form
for a `uint` and a `unit` where a unit can be one of the following

| UNIT | Meaning | VALUE |
| --- | --- | --- |
| K | Kilobyte | 1000 |
| KB | Kilobyte | 1000 |
| KiB | Kibibyte | 1024 |
