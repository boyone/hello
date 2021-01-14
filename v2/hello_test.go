package hello

import "testing"

func TestHelloV2(t *testing.T) {
    want := "Hello, world. v2"
    if got := HelloV2(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestProverbV2(t *testing.T) {
    want := "Concurrency is not parallelism. v2"
    if got := ProverbV2(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}