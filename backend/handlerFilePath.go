package backend

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/samber/lo"
)

func init() {
	bindHandlers = append(bindHandlers, &FilePathHandler{})
}

type FilePathHandler struct{}

func (fp *FilePathHandler) OpenDir(path string) {
	absPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	exec.Command(`cmd`, `/c`, `explorer`, filepath.Join(absPath, lo.If(path != "", path).Else("./out"))).Start()
}
