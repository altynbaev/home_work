package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	sourceFile, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("open source file error: %w", err)
	}
	defer func() {
		err := sourceFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("source file get info error: %w", err)
	}
	sourceFileSize := sourceFileInfo.Size()
	if sourceFileSize <= offset {
		return ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		_, err = sourceFile.Seek(offset, 0)
		if err != nil {
			return fmt.Errorf("source file seek error: %w", err)
		}
	}

	if (limit == 0) || (limit+offset > sourceFileSize) {
		limit = sourceFileSize - offset
	}

	bar := pb.Full.Start64(limit)
	defer bar.Finish()
	barReader := bar.NewProxyReader(sourceFile)

	targetFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("target file create error: %w", err)
	}
	defer func() {
		err := targetFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = io.CopyN(targetFile, barReader, limit)
	if err != nil {
		return fmt.Errorf("copy file error: %w", err)
	}

	return nil
}
