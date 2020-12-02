package main

import (
	"log"

	"github.com/gobuffalo/packr/v2"
	"github.com/jung-kurt/gofpdf"
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
	pdf.AddUTF8FontFromBytes("ArialTrue", "", arialBytes)
	pdf.AddUTF8FontFromBytes("ArialTrue", "I", arialItalicBytes)
	pdf.AddUTF8FontFromBytes("ArialTrue", "B", arialBoldBytes)
	pdf.AddUTF8FontFromBytes("NotoSansSC-Regular", "", notoSansBytes)
	pdf.SetFont("NotoSansSC-Regular", "", 16)
	pdf.AddPage()
	pdf.Cell(40, 10, "Hello, world, 荣锋亮")
	err = pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatalln("some wrong: ", err.Error())
	}
}
