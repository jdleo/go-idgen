package goidgen

import (
	"crypto/rand"
	"errors"
	rand2 "math/rand/v2"
)

// Character set constants for ID generation
const (
	ASCIILowercase = "abcdefghijklmnopqrstuvwxyz"
	ASCIIUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ASCIILetters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits         = "0123456789"
	HexDigits      = "0123456789abcdefABCDEF"
	OctDigits      = "01234567"
	Punctuation    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	URLSafe        = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Printable      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
)

// IDGen provides methods for generating random IDs
type IDGen struct {
	// Deprecated: Use package constants instead (ASCIILowercase, ASCIIUppercase, etc.)
	// These fields are kept for backward compatibility
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

// New returns a new IDGen instance
// Note: This creates an instance for backward compatibility, but the package-level
// constants (ASCIILowercase, Digits, etc.) are preferred for better performance
func New() IDGen {
	return IDGen{
		ASCII_LOWERCASE: ASCIILowercase,
		ASCII_UPPERCASE: ASCIIUppercase,
		ASCII_LETTERS:   ASCIILetters,
		DIGITS:          Digits,
		HEXDIGITS:       HexDigits,
		OCTDIGITS:       OctDigits,
		PUNCTUATION:     Punctuation,
		URL_SAFE:        URLSafe,
		PRINTABLE:       Printable,
	}
}

// Generate generates cryptographically secure, random IDs using crypto/rand
// Accepts optional parameter - alphabet to use for ID generation. If omitted, it will default to URL-safe characters
func (g *IDGen) Generate(length int, alphabet ...string) (string, error) {
	// error checking
	if length <= 0 {
		return "", errors.New("length must be > 0")
	}

	// establish char set to be used
	var chars string

	// check if an alphabet was provided
	if len(alphabet) > 0 {
		chars = alphabet[0]
		if len(chars) == 0 {
			return "", errors.New("alphabet cannot be empty")
		}
		if len(chars) > 255 {
			return "", errors.New("alphabet size must be <= 255 characters")
		}
	} else {
		// use url_safe characters (backward compat with struct field)
		if g.URL_SAFE != "" {
			chars = g.URL_SAFE
		} else {
			chars = URLSafe
		}
	}

	charsLen := len(chars)

	// Fast path for power-of-2 alphabets (no rejection needed)
	if charsLen&(charsLen-1) == 0 {
		return generatePowerOfTwo(chars, charsLen, length)
	}

	// Calculate the mask for rejection sampling
	mask := 1
	for mask < charsLen {
		mask <<= 1
	}
	mask--

	// Pre-allocate result
	result := make([]byte, length)

	// Batch read random bytes - allocate generously to minimize rejection overhead
	bufSize := length + (length >> 1) // 1.5x length
	if bufSize < 32 {
		bufSize = 32
	}
	randomBytes := make([]byte, bufSize)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	resultIdx := 0
	byteIdx := 0

	// Rejection sampling loop
	for resultIdx < length {
		// Refill buffer if needed
		if byteIdx >= len(randomBytes) {
			if _, err := rand.Read(randomBytes); err != nil {
				return "", err
			}
			byteIdx = 0
		}

		// Process byte
		b := int(randomBytes[byteIdx]) & mask
		byteIdx++

		if b < charsLen {
			result[resultIdx] = chars[b]
			resultIdx++
		}
	}

	return string(result), nil
}

// generatePowerOfTwo is optimized for alphabets with power-of-2 lengths (no rejection needed)
func generatePowerOfTwo(chars string, charsLen, length int) (string, error) {
	result := make([]byte, length)

	// Read random bytes directly
	if _, err := rand.Read(result); err != nil {
		return "", err
	}

	// Simple mask, no rejection needed
	mask := byte(charsLen - 1)
	for i := 0; i < length; i++ {
		result[i] = chars[result[i]&mask]
	}

	return string(result), nil
}

// GenerateUnsecure generates unsecure, random IDs using math/rand
// "Unsecure" refers to math/rand being used for RNG rather than a crypto-safe solution
// This is faster but should not be used for security-sensitive applications
// Accepts optional parameter - alphabet to use for ID generation. If omitted, it will default to URL-safe characters
func (g *IDGen) GenerateUnsecure(length int, alphabet ...string) (string, error) {
	// error checking
	if length <= 0 {
		return "", errors.New("length must be > 0")
	}

	// establish char set to be used
	var chars string

	// check if an alphabet was provided
	if len(alphabet) > 0 {
		chars = alphabet[0]
		if len(chars) == 0 {
			return "", errors.New("alphabet cannot be empty")
		}
		if len(chars) > 255 {
			return "", errors.New("alphabet size must be <= 255 characters")
		}
	} else {
		// use url_safe characters (backward compat with struct field)
		if g.URL_SAFE != "" {
			chars = g.URL_SAFE
		} else {
			chars = URLSafe
		}
	}

	// result byte buffer
	result := make([]byte, length)
	charsLen := len(chars)

	// Use math/rand/v2 (automatically seeded, no bias with IntN)
	for i := 0; i < length; i++ {
		result[i] = chars[rand2.IntN(charsLen)]
	}

	return string(result), nil
}
