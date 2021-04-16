package goidgen

import (
	"crypto/rand"
	"errors"
	rand2 "math/rand"
	"strings"
	"time"
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

// constructor for a new goidgen instance
func New() goidgen {
	// seed random
	rand2.Seed(time.Now().UTC().UnixNano())
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
// Accepts optional parameter - alphabet to use for ID generation. If ommitted, it will default to URL-safe characters
func (g *goidgen) Generate(length int, alphabet ...string) (string, error) {
	// error checking
	if length <= 0 {
		return "", errors.New("length must be >= 0")
	} else if len(alphabet) > 0 && len(alphabet[0]) > 255 {
		return "", errors.New("alphabet size must be <= 255 characters")
	}

	// establish char set to be used
	var chars string

	// check if an alphabet was provided
	if len(alphabet) > 0 {
		// use provided alphabet
		chars = alphabet[0]
	} else {
		// use url_safe characters
		chars = g.URL_SAFE
	}

	// randomly generate random bytes
	b := make([]byte, length)
	x, _ := rand.Read(b)
	_ = x

	// len of chars as byte
	len := byte(len(chars))

	// result string builder with preallocated buffer size
	var builder strings.Builder
	// iterate length times
	for i := 0; i < length; i++ {
		// write randomly-drawn byte to builder
		builder.WriteByte(chars[(b[i]/(255/len))%len])
	}

	// return builder's string
	return builder.String(), nil
}

// Generate generates unsecure, random ID's
// "Unsecure" refers to math/rand being used for RNG rather than a crypto-safe solution
// Accepts optional parameter - alphabet to use for ID generation. If ommitted, it will default to URL-safe characters
func (g *goidgen) GenerateUnsecure(length int, alphabet ...string) (string, error) {
	// error checking
	if length <= 0 {
		return "", errors.New("length must be >= 0")
	} else if len(alphabet) > 0 && len(alphabet[0]) > 255 {
		return "", errors.New("alphabet size must be <= 255 characters")
	}

	// establish char set to be used
	var chars string

	// check if an alphabet was provided
	if len(alphabet) > 0 {
		// use provided alphabet
		chars = alphabet[0]
	} else {
		// use url_safe characters
		chars = g.URL_SAFE
	}

	// randomly generate random bytes
	b := make([]byte, length)
	x, _ := rand2.Read(b)
	_ = x

	// len of chars as byte
	len := byte(len(chars))

	// result string builder
	var builder strings.Builder

	// iterate length times
	for i := 0; i < length; i++ {
		// write randomly-drawn byte to builder
		builder.WriteByte(chars[(b[i]/(255/len))%len])
	}

	// return builder's string
	return builder.String(), nil
}
