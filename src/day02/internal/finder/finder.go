package finder

import (
	"day02/internal/options"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Find(path string, findFlags options.FindOptions) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("Skipping file or directory due to permission error: %v\n", err)
				return filepath.SkipDir
			}

			fmt.Printf("Error processing %q: %v\n", path, err)
			return err
		}

		if findFlags.File && isRegularFile(info) {
			handleFile(path, findFlags)
		}

		if findFlags.Dir && isDirectory(info, path) {
			handleDirectory(path)
		}

		if findFlags.SymbolicLinks && isSymbolicLink(info) {
			handleSymbolicLink(path)
		}

		return nil
	})

	return fmt.Errorf("Error while searching in directory %q: %v", path, err)
}

func isRegularFile(info os.FileInfo) bool {
	return info.Mode()&os.ModeType == 0
}

func isDirectory(info os.FileInfo, path string) bool {
	return info.Mode()&os.ModeDir != 0 && filepath.Base(path) != filepath.Base(flag.Arg(0))
}

func isSymbolicLink(info os.FileInfo) bool {
	return info.Mode()&os.ModeSymlink != 0
}

func handleFile(path string, findFlags options.FindOptions) {
	if findFlags.Extension == "" || (len(filepath.Ext(path)) > 0 && findFlags.Extension == filepath.Ext(path)[1:]) {
		fmt.Println(path)
	}
}

func handleDirectory(path string) {
	fmt.Println(path)
}

func handleSymbolicLink(path string) {
	realPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		fmt.Printf("%s -> [broken]\n", path)
	} else {
		fmt.Printf("%s -> %s\n", path, realPath)
	}
}
