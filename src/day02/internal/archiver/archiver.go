package archiver

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func Archive(dir string, paths []string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(paths))

	for _, path := range paths {
		wg.Add(1)
		go func(path string, dir string) {
			defer wg.Done()

			err := archiveFile(path, dir)

			if err != nil {
				errCh <- err
				return
			}

		}(path, dir)
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to archive files: %v", errs)
	}

	return nil
}

func archiveFile(path string, dir string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()
	archiveName := fmt.Sprintf("%s_%d.tar.gz", filepath.Base(path), timestamp)
	archivePath := filepath.Join(dir, archiveName)

	archive, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer archive.Close()

	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	header := &tar.Header{
		Name:    filepath.Base(path),
		Size:    fileStat.Size(),
		Mode:    int64(fileStat.Mode()),
		ModTime: fileStat.ModTime(),
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}
