package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Vale-sail/maroto/pkg/consts"
	"github.com/Vale-sail/maroto/pkg/pdf"
	"github.com/Vale-sail/maroto/pkg/props"
)

func main() {
	begin := time.Now()
	m := pdf.NewMarotoCustomSize(consts.Landscape, "C6", "mm", 114.0, 162.0)
	m.SetPageMargins(5, 5, 5)
	m.AddUTF8Font("NotoSansSC", "", "./font/NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "I", "./font/NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "B", "./font/NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "BI", "./font/NotoSansSC-Regular.ttf")
	// m.SetBorder(true)
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
			m.Text("João Sant'Ana 100 Main Street", props.Text{
				Size:  10,
				Align: consts.Right,
			})
			m.Text("荣锋亮 TN 39021", props.Text{
				Size:   10,
				Align:  consts.Right,
				Family: "NotoSansSC",
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

	err := m.OutputFileAndClose("customsize.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
