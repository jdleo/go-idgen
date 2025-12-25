package goidgen

import (
	"crypto/rand"
	"errors"
	rand2 "math/rand"
)

// properties for a goidgen instance
type goidgen struct {
	ASCII_LOWERCASE string
	ASCII_UPPERCASE string
	ASCII_LETTERS   string
	DIGITS          string
	HEXDIGITS       string
	OCTDIGITS       string
	PUNCTUATION     string
	URL_SAFE        string
	PRINTABLE       string
}

// New returns a new goidgen instance
func New() goidgen {
	// fill fields with predefined character sets
	return goidgen{
		ASCII_LOWERCASE: "abcdefghijklmnopqrstuvwxyz",
		ASCII_UPPERCASE: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		ASCII_LETTERS:   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		DIGITS:          "0123456789",
		HEXDIGITS:       "0123456789abcdefABCDEF",
		OCTDIGITS:       "01234567",
		PUNCTUATION:     "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~",
		URL_SAFE:        "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		PRINTABLE:       "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c",
	}
}

// Generate generates secure, random ID's
// Accepts optional parameter - alphabet to use for ID generation. If omitted, it will default to URL-safe characters
func (g *goidgen) Generate(length int, alphabet ...string) (string, error) {
	// error checking
	if length <= 0 {
		return "", errors.New("length must be >= 0")
	} else if len(alphabet) > 0 && len(alphabet[0]) > 255 {
		return "", errors.New("alphabet size must be <= 255 characters")
	}

	// establish char set to be used
	chars := g.URL_SAFE
	if len(alphabet) > 0 {
		chars = alphabet[0]
	}

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	l := len(chars)
	if l == 64 {
		// Optimization: If length is 64, we can use bitmasking (mostly for URL_SAFE)
		// This is significantly faster than modulo.
		for i := 0; i < length; i++ {
			b[i] = chars[b[i]&63]
		}
	} else {
		// Generic path
		for i := 0; i < length; i++ {
			b[i] = chars[int(b[i])%l]
		}
	}

	return string(b), nil
}

// Generate generates unsecure, random ID's
// "Unsecure" refers to math/rand being used for RNG rather than a crypto-safe solution
// Accepts optional parameter - alphabet to use for ID generation. If omitted, it will default to URL-safe characters
func (g *goidgen) GenerateUnsecure(length int, alphabet ...string) (string, error) {
	// error checking
	if length <= 0 {
		return "", errors.New("length must be >= 0")
	} else if len(alphabet) > 0 && len(alphabet[0]) > 255 {
		return "", errors.New("alphabet size must be <= 255 characters")
	}

	// establish char set to be used
	chars := g.URL_SAFE
	if len(alphabet) > 0 {
		chars = alphabet[0]
	}

	b := make([]byte, length)
	l := len(chars)

	if l == 64 {
		// Optimized path for 64-char alphabet (default)
		// Uses 63-bit random integers to process 10 characters per call
		// avoiding overhead of rand.Read or byte-by-byte calls.
		for i, cache, remain := 0, rand2.Int63(), 10; i < length; {
			if remain == 0 {
				cache, remain = rand2.Int63(), 10
			}
			b[i] = chars[int(cache&63)]
			cache >>= 6
			remain--
			i++
		}
	} else {
		// Generic path using rand.Read for simplicity on non-standard lengths
		rand2.Read(b)
		for i := 0; i < length; i++ {
			b[i] = chars[int(b[i])%l]
		}
	}

	return string(b), nil
}
