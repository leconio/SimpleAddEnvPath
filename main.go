package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}

}

func help() {
	fmt.Println("Usage for addPath")
	fmt.Println("=================")
	fmt.Println("pathAdd [value]")
	fmt.Println("pathAdd [file] [value]")
}

// 寻找默认Path
func findDefaultPath() *os.File {
	userCur, err := user.Current()
	check(err)
	file, err := os.Open(path.Join(userCur.HomeDir, ".zshrc"))
	if err != nil {
		file, err = os.Open(path.Join(userCur.HomeDir, ".bashrc"))
		if err != nil {
			file, err = os.Create(path.Join(userCur.HomeDir, ".bashrc"))
			check(err)
		}
	}
	file, err = os.OpenFile(file.Name(), os.O_RDWR|os.O_APPEND, 0666)
	check(err)
	return file
}

func addPath(file *os.File, value string) {
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	check(err)
	pendStr := "\nexport PATH=$PATH:" + value
	if strings.Contains(string(b), pendStr) {
		fmt.Println(pendStr, "\t已经存在")
	} else {
		_, err = file.Write([]byte(pendStr))
		check(err)
		fmt.Println("修改成功,请手动source")
	}
}

func main() {
	args := os.Args

	var value string
	var myPath string

	if len(args) == 1 {
		value, _ = os.Getwd()
		addPath(findDefaultPath(), value)
	} else if len(args) == 2 {
		value = args[1]
		addPath(findDefaultPath(), value)
	} else if len(args) == 3 {
		myPath = args[1]
		value = args[2]
		file, err := os.OpenFile(myPath, os.O_RDWR|os.O_APPEND, 0666)
		check(err)
		addPath(file, value)
	} else {
		help()
	}
}
