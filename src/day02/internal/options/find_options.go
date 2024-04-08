package options

import (
	"flag"
)

type FindOptions struct {
	File          bool
	Dir           bool
	SymbolicLinks bool
	Extension     string
}

func SetupFindOptions(flags *FindOptions) {
	flag.BoolVar(&flags.File, "f", false, "Find files")
	flag.BoolVar(&flags.Dir, "d", false, "Find directories")
	flag.BoolVar(&flags.SymbolicLinks, "sl", false, "Find symbolic links")
	flag.StringVar(&flags.Extension, "ext", "", "Find files with a specific extension")

	flag.Parse()

	if allFindFlagsAreFalse(*flags) {
		setAllFindFlagsTrue(flags)
	}
}

func allFindFlagsAreFalse(flags FindOptions) bool {
	return !flags.File && !flags.Dir && !flags.SymbolicLinks && flags.Extension == ""
}

func setAllFindFlagsTrue(flags *FindOptions) {
	flags.File = true
	flags.Dir = true
	flags.SymbolicLinks = true
}
