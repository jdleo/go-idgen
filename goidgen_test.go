package goidgen

import (
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	idgen := New()

	t.Run("Constructor initializes property, and fields are non-nil", func(t *testing.T) {
		if idgen.ASCII_LETTERS == "" {
			t.Errorf("Expected igen instance to have valid ASCII_LETTERS property, had empty string")
		}
	})
}

func TestGenerate(t *testing.T) {
	idgen := New()

	t.Run("IDGen generates a string of valid, supplied length", func(t *testing.T) {

		for i := 1; i < 20; i++ {
			id, _ := idgen.Generate(i)

			if len(id) != i {
				t.Errorf("expected {%d}, received {%d}", i, len(id))
			}
		}
	})
}

func TestGenerateCollisionsDefault(t *testing.T) {
	idgen := New()

	t.Run("IDGen has no collisions when generating 10,000 random strings of length 20", func(t *testing.T) {
		seen := map[string]bool{}

		for i := 0; i < 10_000; i++ {
			id, _ := idgen.Generate(20)

			if _, ok := seen[id]; ok {
				t.Errorf("Collision! {%s} was randomly generated twice", id)
				return
			}

			seen[id] = true
		}
	})
}

func TestGenerateCollisions(t *testing.T) {
	idgen := New()

	t.Run("IDGen has no collisions when generating 10,000 random strings of length 20", func(t *testing.T) {
		seen := map[string]bool{}

		for i := 0; i < 10_000; i++ {
			id, _ := idgen.Generate(20, idgen.ASCII_LETTERS)

			if _, ok := seen[id]; ok {
				t.Errorf("Collision! {%s} was randomly generated twice", id)
				return
			}

			seen[id] = true
		}
	})
}

func TestGenerateUnsecureCollisionsDefault(t *testing.T) {
	idgen := New()

	t.Run("IDGen has no collisions when generating 100,000 random strings of length 20", func(t *testing.T) {
		seen := map[string]bool{}

		for i := 0; i < 100_000; i++ {
			id, _ := idgen.GenerateUnsecure(20)

			if _, ok := seen[id]; ok {
				t.Errorf("Collision! {%s} was randomly generated twice", id)
				return
			}

			seen[id] = true
		}
	})
}

func TestGenerateUnsecureCollisions(t *testing.T) {
	idgen := New()

	t.Run("IDGen has no collisions when generating 100,000 random strings of length 20", func(t *testing.T) {
		seen := map[string]bool{}

		for i := 0; i < 100_000; i++ {
			id, _ := idgen.GenerateUnsecure(20, idgen.ASCII_LETTERS)

			if _, ok := seen[id]; ok {
				t.Errorf("Collision! {%s} was randomly generated twice", id)
				return
			}

			seen[id] = true
		}
	})
}

func TestGenerateThrowsErrorOnBadLength(t *testing.T) {
	idgen := New()

	t.Run("Generate should throw an error on invalid lengths", func(t *testing.T) {
		_, err := idgen.Generate(-1)

		if err == nil {
			t.Errorf("expected valid error, received nil")
		}
	})
}

func TestGenerateUnsecureThrowsErrorOnBadLength(t *testing.T) {
	idgen := New()

	t.Run("GenerateUnsecure should throw an error on invalid lengths", func(t *testing.T) {
		_, err := idgen.GenerateUnsecure(-1)

		if err == nil {
			t.Errorf("expected valid error, received nil")
		}
	})
}

func TestGenerateThrowsErrorOnBigAlphabet(t *testing.T) {
	idgen := New()
	alphabet := ""
	for i := 0; i < 300; i++ {
		alphabet += "_"
	}

	t.Run("Generate should throw an error on big alphabet (>= 255 chars)", func(t *testing.T) {
		_, err := idgen.Generate(10, strings.Repeat("_", 300))

		if err == nil {
			t.Errorf("expected valid error, received nil")
		}
	})
}

func TestGenerateUnsecureThrowsErrorOnBigAlphabet(t *testing.T) {
	idgen := New()

	t.Run("Generate should throw an error on big alphabet (>= 255 chars)", func(t *testing.T) {
		_, err := idgen.GenerateUnsecure(10, strings.Repeat("_", 300))

		if err == nil {
			t.Errorf("expected valid error, received nil")
		}
	})
}

func BenchmarkSecure1(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.Generate(1)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkSecure10(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.Generate(10)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkSecure100(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.Generate(100)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkSecure1000(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.Generate(1_000)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkUnsecure1(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.GenerateUnsecure(1)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkUnsecure10(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.GenerateUnsecure(10)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkUnsecure100(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.GenerateUnsecure(100)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func BenchmarkUnsecure1000(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, err := idgen.GenerateUnsecure(1_000)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}
