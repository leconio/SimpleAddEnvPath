package main

import (
	"fmt"
	"github.com/devfacet/gocmd"
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

func getScriptByType(which string) string {
	if which == "sys" {
		return "\nexport PATH=$PATH:"
	} else if which == "go" {
		return "\nexport GOPATH=$GOPATH:"
	}
	panic("error type")
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

func addPath(file *os.File, value string, which string) {
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	check(err)
	pendStr := getScriptByType(which) + value
	if strings.Contains(string(b), pendStr) {
		fmt.Println(pendStr, "\t已经存在")
	} else {
		_, err = file.Write([]byte(pendStr))
		check(err)
		fmt.Println("修改成功,请手动source")
	}
}

func Run(args []string, which string) {
	var value string
	var myPath string
	if len(args) == 1 {
		value, _ = os.Getwd()
		addPath(findDefaultPath(), value, which)
	} else if len(args) == 2 {
		value = args[1]
		addPath(findDefaultPath(), value, which)
	} else if len(args) == 3 {
		myPath = args[1]
		value = args[2]
		file, err := os.OpenFile(myPath, os.O_RDWR|os.O_APPEND, 0666)
		check(err)
		addPath(file, value, which)
	}
}

func main() {
	flags := struct {
		Help    bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version bool `short:"v" long:"version" description:"Display version"`
		Go      struct {
			ScriptPath string `short:"s" long:"scriptpath" required:"false" description:"script path"`
			Path       string `short:"p" long:"path" required:"false" description:"path"`
		} `command:"go" description:"add current path to gopath"`
		Sys struct {
			ScriptPath string `short:"s" long:"scriptpath" required:"false" description:"script path"`
			Path       string `short:"p" long:"path" required:"false" description:"path"`
		} `command:"sys" description:"add current path to path"`
	}{}

	_, _ = gocmd.HandleFlag("Go", func(cmd *gocmd.Cmd, args []string) error {
		var runArgs []string
		runArgs = append(runArgs, args[0])
		if flags.Go.ScriptPath != "" {
			runArgs = append(runArgs, flags.Go.ScriptPath)
		}
		if flags.Go.Path != "" {
			runArgs = append(runArgs, flags.Go.Path)
		}
		Run(runArgs, "go")
		return nil
	})

	_, _ = gocmd.HandleFlag("Sys", func(cmd *gocmd.Cmd, args []string) error {
		var runArgs []string
		runArgs = append(runArgs, args[0])
		if flags.Sys.ScriptPath != "" {
			runArgs = append(runArgs, flags.Sys.ScriptPath)
		}
		if flags.Sys.Path != "" {
			runArgs = append(runArgs, flags.Sys.Path)
		}
		Run(runArgs, "sys")
		return nil
	})

	// Init the app
	_, _ = gocmd.New(gocmd.Options{
		Name:        "basic",
		Version:     "1.0.0",
		Description: "快速添加环境变量和GOPATH变量",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
