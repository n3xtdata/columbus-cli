package utils

import (
	"github.com/n3xtdata/columbus-cli/cmd/config"
	"io/ioutil"
	"strconv"
)

func GetPid() (string, int) {

	b, _ := ioutil.ReadFile(config.ColumbusPidFile) // just pass the file name

	s := string(b)

	i, err := strconv.Atoi(s)

	if err != nil {
		return s, 0
	}

	return s, i

}
