package system

import "os"

func RecordHistoryMode() bool {
	return os.Getenv("HISTORY") != ""
}
