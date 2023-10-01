package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkResults(b *testing.B) {
	r, _ := zip.OpenReader("testdata/users.dat.zip")
	defer r.Close()

	data, _ := r.File[0].Open()

	for i := 0; i < b.N; i++ {
		GetDomainStat(data, "biz")
	}
}
