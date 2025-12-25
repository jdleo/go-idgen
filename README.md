# Go IDGen

[![Go Report](https://goreportcard.com/badge/github.com/jdleo/go-idgen)](https://goreportcard.com/report/github.com/jdleo/go-idgen)
[![Go Coverage](https://github.com/jdleo/go-idgen/wiki/coverage.svg)](https://raw.githack.com/wiki/jdleo/go-idgen/coverage.html)
[![Test Status](https://github.com/jdleo/go-idgen/workflows/Tests/badge.svg)](https://github.com/jdleo/go-idgen/actions)
[![Lint Status](https://github.com/jdleo/go-idgen/workflows/Lint/badge.svg)](https://github.com/jdleo/go-idgen/actions)
[![GitHub issues](https://img.shields.io/github/issues/jdleo/go-idgen.svg)](https://github.com/jdleo/go-idgen/issues)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/jdleo/go-idgen/actions/LICENSE)

IDGen is a cryptographically secure, URL-safe by default, ID generator. You can specify length, and get an ID quickly generated for you. You can also opt out of using crypto number generation to get quicker number generation.

## Install

```bash
go get github.com/jdleo/go-idgen
```

## Usage

Import

```go
import (
	goidgen "github.com/jdleo/go-idgen"
)
```

Generate a cryptographically, secure id.
By default, it will use `URL_SAFE` charset

```go
idgen := goidgen.New()

id, err := idgen.Generate(10) // generates random, URL-safe ID of length 10
id2, err := idgen.Generate(10, idgen.ASCII_LOWERCASE) // generates random, lowercase ID of length 10
```

Generate a random id.
By default, it will use `URL_SAFE` charset.
Not cryptographically secure, uses `math/rand`

```go
idgen := goidgen.New()

id, err := idgen.GenerateUnsecure(10) // generates random, URL-safe ID of length 10
id2, err := idgen.GenerateUnsecure(10, idgen.ASCII_LOWERCASE) // generates random, lowercase ID of length 10
```

## Built-In Charsets

```json
{
  "ASCII_LOWERCASE": "abcdefghijklmnopqrstuvwxyz",
  "ASCII_UPPERCASE": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
  "ASCII_LETTERS": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
  "DIGITS": "0123456789",
  "HEXDIGITS": "0123456789abcdefABCDEF",
  "OCTDIGITS": "01234567",
  "PUNCTUATION": "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
  "URL_SAFE": "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
  "PRINTABLE": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{\t\n\r\x0b\x0c"
}
```

## Benchmarks

| Benchmark    | Old (ns/op) | New (ns/op) | Speedup | Improvement |
|:-------------|------------:|------------:|:--------|:------------|
| Secure1      |      630.00 |      618.70 | 1.02x   | 1.79%       |
| Secure10     |      691.80 |      646.60 | 1.07x   | 6.53%       |
| Secure100    |     1052.00 |      885.70 | 1.19x   | 15.81%      |
| Secure1000   |     5377.00 |     3791.00 | 1.42x   | 29.50%      |
| Unsecure1    |       23.42 |       16.05 | 1.46x   | 31.47%      |
| Unsecure10   |       62.24 |       35.16 | 1.77x   | 43.51%      |
| Unsecure100  |      310.20 |      140.80 | 2.20x   | 54.61%      |
| Unsecure1000 |     2732.00 |     1135.00 | 2.41x   | 58.46%      |

## Notice

If you use idgen in your project, please let me know!

If you have any issues, just feel free and open it in this repository, thanks!

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
