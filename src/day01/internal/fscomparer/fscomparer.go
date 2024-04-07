package fscomparer

import (
	"bufio"
	"fmt"
	"os"
)

func Compare(oldDump, newDump *os.File) {
	oldDumpLines := make(map[string]bool)

	scannerOldDump := bufio.NewScanner(oldDump)
	for scannerOldDump.Scan() {
		oldDumpLines[scannerOldDump.Text()] = true
	}

	if err := scannerOldDump.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	scannerNewDump := bufio.NewScanner(newDump)
	for scannerNewDump.Scan() {
		line := scannerNewDump.Text()

		if _, ok := oldDumpLines[line]; !ok {
			fmt.Println("ADDED", line)
		} else {
			delete(oldDumpLines, line)
		}
	}

	if err := scannerNewDump.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	for line := range oldDumpLines {
		fmt.Println("REMOVED", line)
	}

}
