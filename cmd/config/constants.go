package config

import (
	"os"
)

var (
	ColumbusHome         string
	ColumbusLogFile      string
	ColumbusPidFile      string
	ColumbusJar          string
	ColumbusJarGithubUrl string;
)

func init() {

	ColumbusHome = os.Getenv("COLUMBUS_HOME")

	var columbusLogDir = ColumbusHome + "/logs"
	var columbusTmpDir = ColumbusHome + "/tmp"
	var columbusLibsDir = ColumbusHome + "/libs"

	os.MkdirAll(columbusLogDir, os.ModePerm)
	os.MkdirAll(columbusTmpDir, os.ModePerm)
	os.MkdirAll(columbusLibsDir, os.ModePerm)

	ColumbusLogFile = columbusLogDir + "/columbus.log"
	ColumbusPidFile = columbusTmpDir + "/pid.txt"
	ColumbusJar = columbusLibsDir + "/columbus.jar"

	ColumbusJarGithubUrl = "https://github.com/n3xtdata/columbus/raw/artifacts/columbus-0.0.1-SNAPSHOT.jar"

}
