package main

import (
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx, _ := excelize.OpenFile("Book1.xlsx")

	options := excelize.GraphicOptions{
		AutoFit: true,
	}
	xlsx.AddPicture("Sheet1", "B2", "image.png", &options)

	xlsx.SaveAs("Book1_output.xlsx")
}
