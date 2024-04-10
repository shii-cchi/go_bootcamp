package input

import "flag"

func GetDbFilenameForRead() string {
	var dbFilename string

	flag.StringVar(&dbFilename, "f", "", "Specifies the filename of the database")
	flag.Parse()

	return dbFilename
}

func GetDbFilenamesForCompare() (string, string) {
	var oldDBFilename string
	var newDBFilename string

	flag.StringVar(&oldDBFilename, "old", "", "Specifies the filename of the old database")
	flag.StringVar(&newDBFilename, "new", "", "Specifies the filename of the new database")
	flag.Parse()

	return oldDBFilename, newDBFilename
}

func GetDumpsFilenames() (string, string) {
	var oldDumpFilename string
	var newDumpFilename string

	flag.StringVar(&oldDumpFilename, "old", "", "Specifies the filename of the old filesystem dump")
	flag.StringVar(&newDumpFilename, "new", "", "Specifies the filename of the new filesystem dump")
	flag.Parse()

	return oldDumpFilename, newDumpFilename
}
