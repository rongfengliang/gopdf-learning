package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rongfengliang/maroto/pkg/color"

	"github.com/gobuffalo/packr/v2"
	"github.com/rongfengliang/maroto/pkg/consts"
	"github.com/rongfengliang/maroto/pkg/pdf"
	"github.com/rongfengliang/maroto/pkg/props"
)

func main() {
	box := packr.New("pdf", "../font")
	begin := time.Now()
	m := pdf.NewMarotoCustomSize(consts.Landscape, "C6", "mm", 114.0, 162.0)
	m.SetPageMargins(5, 5, 5)
	notoSansBytes, err := box.Find("NotoSansSC-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	m.AddUTF8FontFromBytes("notosanssc", "", notoSansBytes)
	// m.AddUTF8FontFromBytes("notosanssc", "I", notoSansBytes)
	m.AddUTF8FontFromBytes("notosanssc", "B", notoSansBytes)
	// m.SetBorder(true)
	headers := []string{"姓名", "年龄"}
	contents := [][]string{
		{"大龙", "11"},
		{"rong", "233"},
	}
	m.TableList(headers, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Family:    consts.NotoSansSC,
			GridSizes: []uint{3, 9},
		},
		ContentProp: props.TableListContent{
			Family:    consts.NotoSansSC,
			GridSizes: []uint{3, 9},
		},
		Align: consts.Center,
		AlternatedBackground: &color.Color{
			Red:   100,
			Green: 20,
			Blue:  255,
		},
		HeaderContentSpace: 10.0,
		Line:               false,
	})
	m.Row(40, func() {
		m.Col(4, func() {
			_ = m.FileImage("biplane.jpg", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.Text("Gopher International Shipping, Inc.", props.Text{
				Top:         12,
				Size:        12,
				Extrapolate: true,
			})
		})
		m.ColSpace(4)
	})

	m.Line(10)

	m.Row(30, func() {
		m.Col(12, func() {
			m.Text("北京市海淀区", props.Text{
				Size:   10,
				Align:  consts.Right,
				Family: consts.NotoSansSC,
			})
			m.Text("荣锋亮 TN 39021", props.Text{
				Size:   10,
				Align:  consts.Right,
				Family: consts.NotoSansSC,
				Top:    10,
			})
			m.Text("United States (USA)", props.Text{
				Size:  10,
				Align: consts.Right,
				Top:   20,
			})
		})
	})

	m.Row(30, func() {
		m.Col(12, func() {
			m.QrCode("https://cnblogs.com/rongfengliang")
		})
	})

	err = m.OutputFileAndClose("customsize.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
