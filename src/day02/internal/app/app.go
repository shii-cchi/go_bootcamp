package app

import (
	"day02/internal/archiver"
	"day02/internal/counter"
	"day02/internal/finder"
	"day02/internal/input"
	"day02/internal/options"
	"day02/internal/runner"
	"fmt"
)

func RunFinder() {
	var findFlags options.FindOptions
	options.SetupFindOptions(&findFlags)

	if findFlags.Extension != "" && !findFlags.File {
		fmt.Println("Error: Flag -ext works only with flag -f")
		return
	}

	path, err := input.GetPathForFind()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = finder.Find(path, findFlags)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func RunCounter() {
	var countFlags options.CountOptions
	options.SetupCountOptions(&countFlags)

	if !input.IsValidCombination(countFlags) {
		fmt.Println("Error: Invalid combination of flags. Only one of -l, -m, or -w can be true.")
		return
	}

	paths, err := input.GetPathsForCount()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	counter.Count(paths, countFlags)
}

func RunRunner() {
	command, err := input.GetCommandForRunner()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = runner.Run(command)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func RunArchiver() {
	dir := options.SetupArchiverOptions()

	err := input.CheckDir(dir)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	paths, err := input.GetPathsForArchive()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = archiver.Archive(dir, paths)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
