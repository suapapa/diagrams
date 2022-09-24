package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	diagrams "github.com/suapapa/go_diagrams"
)

var (
	diagramIn = "diagram.py"
)

const (
	anchorStr = "# Diagrams Sandbox: DO NOT DELETE THIS LINE #"
	pyMine    = `import types
def imports():
	for name, val in globals().items():
		if isinstance(val, types.ModuleType):
			yield val.__name__

# this should contain [__builtin__, type] only
if len(list(imports())) != 2:
		import sys
		sys.exit("invalid input")
`
)

func main() {
	// read stdin
	buf := &bytes.Buffer{}
	io.Copy(buf, os.Stdin)
	inputStr := buf.String()

	// check for invalid input
	if err := securityCheck(inputStr); err != nil {
		ret := diagrams.Result{
			Msg: "security check failed",
			Err: err.Error(),
		}

		json.NewEncoder(os.Stdout).Encode(&ret)
		os.Exit(-1)
	}

	// insert mine for possible attack
	inputStr = strings.Replace(inputStr, anchorStr, pyMine, -1)

	// copy input to file
	w, err := os.Create(diagramIn)
	checkErr(err)
	fmt.Fprint(w, inputStr)
	w.Close()

	// deadline is 2 secs
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// run diagrams code with python (this program should run in gVisor)
	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "python", diagramIn)
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf
	err = cmd.Run()

	outStr := outBuf.String()
	errStr := errBuf.String()
	if err != nil {
		printJson(&diagrams.Result{Msg: outStr, Err: errStr})
		return
	}

	// find out diagramOut exists
	// diagrams 파이썬 파일 안에 선언된 이름으로 png가 생성됨. glob으로 찾자!
	match, err := filepath.Glob("*.png")
	checkErr(err)
	if len(match) != 1 {
		checkErr(fmt.Errorf("fail to gen diagram png"))
	}
	diagramOut := match[0]
	log.Println(diagramOut)

	_, err = os.Stat(diagramOut)
	checkErr(err)
	defer os.RemoveAll(diagramOut)

	f, err := os.Open(diagramOut)
	checkErr(err)
	defer f.Close()

	content, err := io.ReadAll(f)
	checkErr(err)
	encoded := base64.StdEncoding.EncodeToString(content)
	printJson(&diagrams.Result{Img: encoded, Name: diagramOut, Msg: outStr, Err: errStr})
}

func printJson(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}

func checkErr(err error) {
	checkErrMsg("", err)
}

func checkErrMsg(msg string, err error) {
	if err != nil {
		ret := diagrams.Result{
			Msg: msg,
			Err: err.Error(),
		}

		json.NewEncoder(os.Stdout).Encode(&ret)
		os.Exit(-1)
	}
}
