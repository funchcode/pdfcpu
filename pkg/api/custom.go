package api

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/mocha"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func InjectFormDataFile(inFile string, outFile string, formFile string, resourceDir string) (err error) {
	// form 존재 확인
	if !fileExists(formFile) {
		return
	}
	var (
		f0   *os.File
		m    = map[int][]*pdfcpu.Watermark{}
		form = &mocha.Form{}
	)

	if f0, err = os.Open(formFile); err != nil {
		return err
	}
	fmt.Println(form)
	fmt.Println(&form)
	if err := mocha.FromJSON(f0, form); err != nil {
		return err
	}
	fmt.Println(&form)
	wmts, _ := mocha.ExtractWatermarkTargets(form)
	fmt.Println(len(*wmts))

	for _, wmt := range *wmts {
		wms := []*pdfcpu.Watermark{}
		// todo image인 경우 media 경로에 png 파일이 없을 경우 svg 변환 진행
		if wmt.Mode == "image" && mocha.HasSVGExtension(wmt.ModeParm) {
			pathTrees := strings.Split(wmt.ModeParm, "/")
			filename := pathTrees[len(pathTrees) - 1]
			svgFile := filepath.Join(resourceDir, filename)
			pngFilename := filename + ".png"
			pngFile := filepath.Join(resourceDir, pngFilename)
			if _, err := os.Stat(pngFile); os.IsNotExist(err) {
				mocha.ConvertSvgToPng(svgFile, pngFile)
			} else {
				fmt.Println("Is Exist")
			}
			wmt.ModeParm = pngFile
		}
		// todo checkout인 경우 준비되어 있는 png 파일로 대체 (( 추후에 진행 ** ))

		pages := wmt.SelectedPages

		switch wmt.Mode {
		case "text":
			wm, err := TextWatermark(wmt.ModeParm, wmt.WmConf, true, false, pdfcpu.POINTS)
			wms = append(wms, wm)
			if err != nil {
				return err
			}
		case "image":
			wm, err := ImageWatermark(wmt.ModeParm, wmt.WmConf, true, false, pdfcpu.POINTS)
			wms = append(wms, wm)
			if err != nil {
				return err
			}
		}

		for _, page := range pages {
			i, _ := strconv.Atoi(page)
			pm := m[i]
			pm = append(pm, wms...)
			m[i] = pm
		}
	}

	defer func() {
		if err := f0.Close(); err != nil {
			return
		}
	}()

	err = AddWatermarksSliceMapFile(inFile, outFile, m, nil)

	return err
}

func InjectFormDataFileV1(inFile string, outFile string, formFile string, resourceDir string) (err error) {
	// form 존재 확인
	if !fileExists(formFile) {
		return
	}
	var (
		in string
		out string
		f0 *os.File
		form = &mocha.Form{}
	)

	if f0, err = os.Open(formFile); err != nil {
		return err
	}
	fmt.Println(form)
	fmt.Println(&form)
	if err := mocha.FromJSON(f0, form); err != nil {
		return err
	}
	fmt.Println(&form)
	wms, _ := mocha.ExtractWatermarkTargets(form)
	fmt.Println(len(*wms))

	// todo 현재 signtext, signature 영역만 진행하도록 설정
	for i, wm := range *wms {
		if i == 0 {
			in = inFile
		} else {
			in = out
		}
		if len(*wms) == 1 || len(*wms) - 1 == i {
			out = outFile
		} else {
			out = inFile + ".tmp"
		}
		// FOR-LOOP로 실행
		fmt.Println("[1]" + in)
		fmt.Println("[2]" + out)
		fmt.Println("[3]" + wm.OutFile)
		fmt.Println("[4]" + outFile)
		fmt.Println(wm.SelectedPages)
		fmt.Println("[5]" + wm.ModeParm)
		fmt.Println("[6]" + wm.WmConf)
		// todo image인 경우 media 경로에 png 파일이 없을 경우 svg 변환 진행
		if wm.Mode == "image" && mocha.HasSVGExtension(wm.ModeParm) {
			pathTrees := strings.Split(wm.ModeParm, "/")
			filename := pathTrees[len(pathTrees) - 1]
			svgFile := filepath.Join(resourceDir, filename)
			pngFilename := filename + ".png"
			pngFile := filepath.Join(resourceDir, pngFilename)
			mocha.ConvertSvgToPng(svgFile, pngFile)
			wm.ModeParm = pngFile
		}
		// todo checkout인 경우 준비되어 있는 png 파일로 대체 (( 추후에 진행 ** ))
		switch wm.Mode {
		case "text":
			err := AddTextWatermarksFile(in, out, wm.SelectedPages, false, wm.ModeParm,  wm.WmConf, nil)
			if err != nil {
				return err
			}
		case "image":
			err := AddImageWatermarksFile(in, out, wm.SelectedPages, false, wm.ModeParm,  wm.WmConf, nil)
			if err != nil {
				return err
			}
		}
	}
	err2 := AddTextWatermarksFile("E:/default_temp.pdf", "E:/default_temp_1.pdf", []string{"1"}, false, "9999999",  "font:Courier, sc:1 abs, points:10, pos:tl, off:300 -200, fillc:#000000, rot:0", nil)
	if err2 != nil {
		return err
	}

	defer func() {
		if err := f0.Close(); err != nil {
			return
		}
	}()

	return err
}