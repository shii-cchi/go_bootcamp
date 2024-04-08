package options

import "flag"

type CountOptions struct {
	Lines      bool
	Characters bool
	Words      bool
}

func SetupCountOptions(flags *CountOptions) {
	flag.BoolVar(&flags.Lines, "l", false, "Count lines")
	flag.BoolVar(&flags.Characters, "m", false, "Count characters")
	flag.BoolVar(&flags.Words, "w", false, "Count words")

	flag.Parse()

	if allCountFlagsAreFalse(*flags) {
		flags.Words = true
	}
}

func allCountFlagsAreFalse(flags CountOptions) bool {
	return !flags.Lines && !flags.Characters && !flags.Words
}
