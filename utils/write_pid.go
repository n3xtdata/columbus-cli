package utils

import (
	"github.com/n3xtdata/columbus-cli/cmd/config"
	"os"
)

func WritePid(pid string) {

	pidFile, err := os.Create(config.ColumbusPidFile)

	if err != nil {
		panic(err)
	}

	defer pidFile.Close()

	pidFile.WriteString(pid)

	pidFile.Sync()

}
