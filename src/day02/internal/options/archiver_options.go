package options

import "flag"

func SetupArchiverOptions() string {
	var dir string
	flag.StringVar(&dir, "a", ".", "Directory for saving archives")

	flag.Parse()

	return dir
}
