package framework

import (
	"os"
)

func init() {
	if _, err := os.Stat(DataPathName); err != nil {
		projectDataPath = "../../" + DataPathName
	} else {
		projectDataPath = DataPathName
	}
}
