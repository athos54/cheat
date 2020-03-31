package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/installer"
	"github.com/mitchellh/go-homedir"
)

func cmdPODownload() {
	confdir := getConfDir()
	po := path.Join(confdir, "/cheatsheets/po")
	installer.PullPO(po)
}

func cmdPOUpload() {
	confdir := getConfDir()
	po := path.Join(confdir, "/cheatsheets/po")
	installer.PushPO(po)
}

func cmdPOAdd() {
	pipeArray := strings.Split(readPipe(), "\\n")
	filename := os.Args[2]
	fullpath := path.Join(getConfDir(), "/cheatsheets/po", filename)

	if !fileExists(fullpath) {
		createEmptyFile(fullpath, filename)
	}
	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line := range pipeArray {
		file.WriteString(line + "\n")
	}
}

func createEmptyFile(path string, filename string) {
	emptyFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	txt := fmt.Sprintf("tags [ %v ]\n", filename)
	emptyFile.WriteString("---\n")
	emptyFile.WriteString(txt)
	emptyFile.WriteString("---\n")
	emptyFile.Close()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getConfDir() string {
	envvars := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envvars[pair[0]] = pair[1]
	}
	home, _ := homedir.Dir()
	confpaths, _ := config.Paths(runtime.GOOS, home, envvars)
	confpath, _ := config.Path(confpaths)
	confpath = confpaths[0]
	confdir := path.Dir(confpath)
	return confdir
}

func readPipe() string {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	var a string
	for j := 0; j < len(output); j++ {
		a = a + string(output[j])
	}
	return a
}
