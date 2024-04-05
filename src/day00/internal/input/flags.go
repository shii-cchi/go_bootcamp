package input

import "flag"

type OutputOptions struct {
	Mean              bool
	Median            bool
	Mode              bool
	StandardDeviation bool
}

func SetupOutputOptions(flags *OutputOptions) {
	defineFlags(flags)

	flag.Parse()

	if allFlagsAreFalse(*flags) {
		setAllFlagsTrue(flags)
	}
}

func defineFlags(flags *OutputOptions) {
	flag.BoolVar(&flags.Mean, "mean", false, "Print mean")
	flag.BoolVar(&flags.Median, "median", false, "Print median")
	flag.BoolVar(&flags.Mode, "mode", false, "Print mode")
	flag.BoolVar(&flags.StandardDeviation, "sd", false, "Print standard deviation")
}

func allFlagsAreFalse(flags OutputOptions) bool {
	return !flags.Mean && !flags.Median && !flags.Mode && !flags.StandardDeviation
}

func setAllFlagsTrue(flags *OutputOptions) {
	flags.Mean = true
	flags.Median = true
	flags.Mode = true
	flags.StandardDeviation = true
}
