package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func writeNumbers(n int64, w io.Writer) error {
	for i := int64(0); i < n; i++ {
		_, err := fmt.Fprintf(w, "%d\n", i)
		if err != nil {
			return fmt.Errorf("failed to write %d: %v", i, err)
		}
	}

	return nil
}

func openTestFile(name string) (*os.File, error) {
	path := filepath.Join("/tmp", name)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	return f, nil
}

func benchmarkWriteFile(n int64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := fmt.Sprintf("write-%d.txt", n)
		f, err := openTestFile(filename)
		if err != nil {
			b.Error(err)
		}

		if err := writeNumbers(n, f); err != nil {
			b.Error(err)
		}
	}
}

func benchmarkWriteFileBuffered(n int64, bufSize int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		filename := fmt.Sprintf("write-%d_buffered-%d.txt", n, bufSize)
		f, err := openTestFile(filename)
		if err != nil {
			b.Error(err)
		}

		buf := bufio.NewWriterSize(f, bufSize)

		if err := writeNumbers(n, buf); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkWrite1Mil(b *testing.B) {
	benchmarkWriteFile(1000000, b)
}

func BenchmarkWrite1MilBuffered4KB(b *testing.B) {
	bufSize := 4 * 1024 // 4 KB
	benchmarkWriteFileBuffered(1000000, bufSize, b)
}
