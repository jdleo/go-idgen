package goidgen

import (
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

	t.Run("IDGen has no collisions when generating 10,000 random strings of length 20", func(t *testing.T) {
		seen := map[string]bool{}

		for i := 0; i < 10_000; i++ {
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

	t.Run("IDGen has no collisions when generating 10,000 random strings of length 20", func(t *testing.T) {
		seen := map[string]bool{}

		for i := 0; i < 10_000; i++ {
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
		_, err := idgen.Generate(-1, alphabet)

		if err == nil {
			t.Errorf("expected valid error, received nil")
		}
	})
}

func TestGenerateUnsecureThrowsErrorOnBigAlphabet(t *testing.T) {
	idgen := New()
	alphabet := ""
	for i := 0; i < 300; i++ {
		alphabet += "_"
	}

	t.Run("Generate should throw an error on big alphabet (>= 255 chars)", func(t *testing.T) {
		_, err := idgen.GenerateUnsecure(-1, alphabet)

		if err == nil {
			t.Errorf("expected valid error, received nil")
		}
	})
}

func BenchmarkSecure1(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.Generate(1)
	}
}

func BenchmarkSecure10(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.Generate(10)
	}
}

func BenchmarkSecure100(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.Generate(100)
	}
}

func BenchmarkSecure1000(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.Generate(1_000)
	}
}

func BenchmarkUnsecure1(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.GenerateUnsecure(1)
	}
}

func BenchmarkUnsecure10(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.GenerateUnsecure(10)
	}
}

func BenchmarkUnsecure100(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.GenerateUnsecure(100)
	}
}

func BenchmarkUnsecure1000(b *testing.B) {
	idgen := New()

	for n := 0; n < b.N; n++ {
		_, _ := idgen.GenerateUnsecure(1_000)
	}
}
