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
			fmt.Println(fmt.Errorf("close source file error: %w", err))
		}
	}()

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	sourceFileSize := sourceFileInfo.Size()
	if sourceFileSize <= offset {
		return ErrOffsetExceedsFileSize
	}
	if offset > 0 {
		_, err = sourceFile.Seek(offset, 0)
		if err != nil {
			return err
		}
	}

	bar := pb.Full.Start64(limit)
	defer bar.Finish()
	barReader := bar.NewProxyReader(sourceFile)

	targetFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func() {
		err := targetFile.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("close target file error: %w", err))
		}
	}()

	if limit != 0 {
		if limit+offset > sourceFileSize {
			limit = sourceFileSize - offset
		}
		_, err = io.CopyN(targetFile, barReader, limit)
		if err != nil {
			return fmt.Errorf("copy file error: %w", err)
		}
		return nil
	}
	_, err = io.Copy(targetFile, barReader)
	if err != nil {
		return fmt.Errorf("copy file error: %w", err)
	}

	return nil
}
