package main

import (
	"flag"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/cli"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"os"
)

var (
	inFile, outFile, formFile, resourceDir string
)

func customFlag() {
	mochaUsage := "mocha"
	flag.StringVar(&inFile, "i", "", mochaUsage)
	flag.StringVar(&outFile, "o", "", mochaUsage)
	flag.StringVar(&formFile, "f", "", mochaUsage)
	flag.StringVar(&resourceDir, "r", "", mochaUsage)
}

func customCommandMap() {
	cmdMap.register("mocha", command{processMochaCommand, nil, `설명 부분입니다.`, `설명 부분입니다.`})
}

func processMochaCommand(conf *pdfcpu.Configuration) {
	fmt.Println(inFile)
	fmt.Println(outFile)
	fmt.Println(formFile)
	fmt.Println(resourceDir)

	// 원본 파일 체크
	ensurePDFExtension(inFile)

	// todo inFile Password 설정 확인
	// todo FORM 파일 존재 여부 확인

	customProcess(cli.MochaCommand(inFile, outFile, formFile, resourceDir, conf))
}

func customProcess(cmd *cli.Command) {
	out, err := cli.CustomProcess(cmd)
	if err != nil {
		if needStackTrace {
			fmt.Fprintf(os.Stderr, "Fatal: %+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		os.Exit(1)
	}

	if out != nil && !quiet {
		for _, s := range out {
			fmt.Fprintln(os.Stdout, s)
		}
	}

	os.Exit(0)
}