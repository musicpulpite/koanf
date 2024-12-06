package koanf_test

import (
	"strconv"
	"testing"

	"github.com/knadh/koanf/v2"
	koanfRefactored "github.com/musicpulpite/koanf/v2"
)

const ITERATIONS = 1000

func BenchmarkSerialReads(b *testing.B) {
	k := koanf.New("_")
	entries := generateEntrys(ITERATIONS)

	for _, entry := range entries {
		k.Set(entry.k, entry.v)
	}

	for _, entry := range entries {
		k.Get(entry.k)
	}
}

func BenchmarkSerialReadsWithLockingOverhead(b *testing.B) {
	k := koanfRefactored.NewWithConf(koanfRefactored.Conf{
		Delim:          "_",
		StrictMerge:    false,
		UseConcurrency: true,
	})
	entries := generateEntrys(ITERATIONS)

	for _, entry := range entries {
		k.Set(entry.k, entry.v)
	}

	for _, entry := range entries {
		k.Get(entry.k)
	}
}

type Entry struct {
	k string
	v string
}

func generateEntrys(iterations int) []Entry {
	entries := make([]Entry, iterations)
	for i := range entries {
		iStr := strconv.Itoa(i)
		entries[i] = Entry{k: iStr, v: iStr}
	}

	return entries
}
