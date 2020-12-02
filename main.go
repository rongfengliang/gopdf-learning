package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
	"github.com/jung-kurt/gofpdf"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/valyala/fasttemplate"
)

func main() {
	box := packr.New("pdf", "./font")
	pdf := gofpdf.New("P", "mm", "A4", "")
	arialBytes, err := box.Find("arial.ttf")
	if err != nil {
		log.Fatal(err)
	}
	arialItalicBytes, err := box.Find("arial_italic.ttf")
	if err != nil {
		log.Fatal(err)
	}
	arialBoldBytes, err := box.Find("arial_bold.ttf")
	if err != nil {
		log.Fatal(err)
	}
	notoSansBytes, err := box.Find("NotoSansSC-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	templateBytes, err := box.Find("app.html")
	if err != nil {
		log.Fatal(err)
	}
	pdf.AddUTF8FontFromBytes("ArialTrue", "", arialBytes)
	pdf.AddUTF8FontFromBytes("ArialTrue", "I", arialItalicBytes)
	pdf.AddUTF8FontFromBytes("ArialTrue", "B", arialBoldBytes)
	pdf.AddUTF8FontFromBytes("NotoSansSC-Regular", "", notoSansBytes)
	pdf.SetFont("NotoSansSC-Regular", "", 16)
	pdf.AddPage()
	pt := pdf.PointConvert(6)
	v, _ := mem.VirtualMemory()
	t := fasttemplate.New(string(templateBytes), "{{", "}}")
	s := t.ExecuteString(map[string]interface{}{
		"total":       fmt.Sprintf("%d", v.Total),
		"free":        fmt.Sprintf("%d", v.Free),
		"usedPercent": fmt.Sprintf("%f", v.UsedPercent),
	})
	_, lineHt := pdf.GetFontSize()
	html := pdf.HTMLBasicNew()
	pdf.Ln(lineHt + pt)
	html.Write(lineHt, s)
	err = pdf.OutputFileAndClose("mem-report.pdf")
	if err != nil {
		log.Fatalln("some wrong: ", err.Error())
	}
}
