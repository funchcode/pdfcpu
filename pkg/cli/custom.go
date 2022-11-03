package cli

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pkg/errors"
)

var customCmdMap = map[pdfcpu.CommandMode]func(cmd *Command) ([]string, error){
	pdfcpu.MOCHA: 		Mocha,
}

func CustomProcess(cmd *Command) (out []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("unexpected panic attack: %v\n", r)
		}
	}()

	cmd.Conf.Cmd = cmd.Mode

	if f, ok := customCmdMap[cmd.Mode]; ok {
		return f(cmd)
	}

	return nil, errors.Errorf("pdfcpu: process: Unknown command mode %d\n", cmd.Mode)
}

func Mocha(cmd *Command) ([]string, error) {
	fmt.Println("Mocha Func")
	return nil, api.InjectFormDataFile(*cmd.InFile, *cmd.OutFile, *cmd.inFileJSON, *cmd.InDir)
}

func MochaCommand(inFile string, outFile string, formFile string, resourceDir string, conf *pdfcpu.Configuration) *Command {
	if conf == nil {
		conf = pdfcpu.NewDefaultConfiguration()
	}
	conf.Cmd = pdfcpu.MOCHA
	return &Command{
		Mode:    pdfcpu.MOCHA,
		InFile: &inFile,
		OutFile: &outFile,
		inFileJSON: &formFile,
		InDir: &resourceDir,
		Conf:    conf}
}