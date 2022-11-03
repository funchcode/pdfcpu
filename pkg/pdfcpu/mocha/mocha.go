package mocha

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type WatermarkTarget struct {
	InFile, OutFile, Mode, ModeParm, WmConf string
	SelectedPages []string
}

func FromJSON(rd io.Reader, form *Form) error {
	fmt.Println("FromJSON")
	bb, err := io.ReadAll(rd)
	if err != nil {
		return err
	}

	inform, err := parseFormJSON(bb)
	// todo nil 처리
	form.Common = inform.Common
	form.Users = inform.Users
	if err != nil {
		return err
	}
	if err := form.Validate(); err != nil {
		return err
	}

	return nil
}

func parseFormJSON(bb []byte) (*Form, error) {
	// todo Form 검증
	if !json.Valid(bb) {
		return nil, errors.Errorf("mocha: invalid JSON encoding detected.")
	}

	form := &Form{}
	if err := json.Unmarshal(bb, form); err != nil {
		return nil, err
	}
	fmt.Println(&form)
	return form, nil
}

func ExtractWatermarkTargets(form *Form) (*[]WatermarkTarget, error) {
	fmt.Println("0a0a0a0a0a0a0a0a0a0a")
	fmt.Println(form)
	var wms []WatermarkTarget

	//for _, common := range form.Common {
	//	cwm, _ := extractWatermarkFromCommon(&common)
	//	wms = append(wms, *cwm)
	//}

	for _, users := range form.Users {
		data := users.Data
		signTextSlice := data.SignTextGroup
		for _, signtext := range signTextSlice {
			stwm, _ := extractWatermarkFromSignText(&signtext)
			wms = append(wms, *stwm)
		}
		//signCheckboxSlice := data.SignCheckboxGroup
		//for _, signcheckbox := range signCheckboxSlice {
		//	scbwm, _ := extractWatermarkFromSignCheckbox(&signcheckbox)
		//	wms = append(wms, *scbwm)
		//}
		signatureSlice := data.SignatureGroup
		for _, signature := range signatureSlice {
			swm, _ := extractWatermarkFromSignature(&signature)
			wms = append(wms, *swm)
		}
	}

	fmt.Println("0a0a0a0a0a0a0a0a0a0a")
	return &wms, nil
}

func extractWatermarkFromCommon(common *Common) (*WatermarkTarget, error) {
	wm := &WatermarkTarget{
		"",
		"",
		"",
		"",
		"",
		[]string{strconv.Itoa(common.Page)},
	}
	return wm, nil
}

func extractWatermarkFromSignText(group *SignTextGroup) (*WatermarkTarget, error) {
	y := group.Top - float64(group.FontSize) / 2
	wm := &WatermarkTarget{
		"E:/default.pdf",
		"E:/default_temp.pdf",
		"text",
		group.Value,
		"font:NotoSansCJKkr-Regular, sc:1 abs, points:" + strconv.Itoa(group.FontSize) + ", pos:tl, off:" + fmt.Sprintf("%f", group.Left) + " -" + fmt.Sprintf("%f", y) + ", fillc:#000000, rot:0",
		[]string{strconv.Itoa(group.Page)},
	}
	return wm, nil
}

func extractWatermarkFromSignature(group *SignatureGroup) (*WatermarkTarget, error) {
	wm := &WatermarkTarget{
		"",
		"",
		"image",
		group.Value,
		"sc:.33 abs, pos:tl, offset:" + fmt.Sprintf("%f", group.Left) + " -" + fmt.Sprintf("%f", group.Top) + ", rot:0",
		[]string{strconv.Itoa(group.Page)},
	}
	return wm, nil
}

func extractWatermarkFromSignCheckbox(group *SignCheckboxGroup) (*WatermarkTarget, error) {
	return nil, nil
}

func ConvertSvgToPng(svgFile string, outFile string) {
	inFile := filepath.Join(svgFile)
	fmt.Println(inFile)

	in, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	icon, _ := oksvg.ReadIconStream(in)
	w := icon.ViewBox.W
	h := icon.ViewBox.H
	icon.SetTarget(0, 0, w, h)
	rgba := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	icon.Draw(rasterx.NewDasher(int(w), int(h), rasterx.NewScannerGV(int(w), int(h), rgba, rgba.Bounds())), 1)

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	err = png.Encode(out, rgba)
	if err != nil {
		panic(err)
	}
}

func HasSVGExtension(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".svg")
}