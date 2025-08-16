package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	CreateFile(f)
	AddChart(f)

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	ReadFile()
	// AddPicture()
}

func CreateFile(f *excelize.File) {
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A2", "Hello world.")

	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellValue("Sheet2", "B2", 100)
	f.SetActiveSheet(index)
}

func AddChart(f *excelize.File) {
	// Create a new sheet named "Sheet3"
	index, err := f.NewSheet("Sheet3")
	if err != nil {
		fmt.Println("failed to create Sheet3: %v", err)
	}

	// Define the data for the chart
	data := [][]interface{}{
		{nil, "Apple", "Orange", "Pear"},
		{"Small", 2, 3, 3},
		{"Normal", 5, 2, 4},
		{"Large", 6, 7, 8},
	}

	// Write data to Sheet3
	for idx, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println("failed to convert coordinates at row %d: %v", idx+1, err)
		}

		if err := f.SetSheetRow("Sheet3", cell, &row); err != nil {
			fmt.Println("failed to set row at %s: %v", cell, err)
		}
	}

	// Set Sheet3 as the active sheet
	f.SetActiveSheet(index)

	// Add a 3D clustered column chart
	err = f.AddChart("Sheet3", "E1", &excelize.Chart{
		Type: excelize.Col3DClustered,
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet3!$A$2",      // Reference Small
				Categories: "Sheet3!$B$1:$D$1", // Apple, Orange, Pear
				Values:     "Sheet3!$B$2:$D$2", // Data for Small
			},
			{
				Name:       "Sheet3!$A$3",      // Reference Normal
				Categories: "Sheet3!$B$1:$D$1", // Apple, Orange, Pear
				Values:     "Sheet3!$B$3:$D$3", // Data for Normal
			},
			{
				Name:       "Sheet3!$A$4",      // Reference Large
				Categories: "Sheet3!$B$1:$D$1", // Apple, Orange, Pear
				Values:     "Sheet3!$B$4:$D$4", // Data for Large
			},
		},
		Title: []excelize.RichTextRun{{Text: "Fruit 3D Clustered Column Chart"}},
	})
	if err != nil {
		fmt.Println("failed to add chart: %v", err)
	}
}

// func AddChart(f *excelize.File) {
// 	index, err := f.NewSheet("Sheet3")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for idx, row := range [][]interface{}{
// 		{nil, "Apple", "Orange", "Pear"}, {"Small", 2, 3, 3},
// 		{"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
// 	} {
// 		cell, err := excelize.CoordinatesToCellName(1, idx+1)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		f.SetSheetRow("Sheet3", cell, &row)
// 	}
// 	f.SetActiveSheet(index)
// 	if err := f.AddChart("Sheet3", "E1", &excelize.Chart{
// 		Type: excelize.Col3DClustered,
// 		Series: []excelize.ChartSeries{
// 			{
// 				Name:       "Sheet1!$A$2",
// 				Categories: "Sheet1!$B$1:$D$1",
// 				Values:     "Sheet1!$B$2:$D$2",
// 			},
// 			{
// 				Name:       "Sheet1!$A$3",
// 				Categories: "Sheet1!$B$1:$D$1",
// 				Values:     "Sheet1!$B$3:$D$3",
// 			},
// 			{
// 				Name:       "Sheet1!$A$4",
// 				Categories: "Sheet1!$B$1:$D$1",
// 				Values:     "Sheet1!$B$4:$D$4",
// 			}},
// 		Title: []excelize.RichTextRun{
// 			{
// 				Text: "Fruit 3D Clustered Column Chart",
// 			},
// 		},
// 	}); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if err := f.SaveAs("Book1.xlsx"); err != nil {
// 		fmt.Println(err)
// 	}
// }

func ReadFile() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet2", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cell)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

// func AddPicture() {
// 	f, err := excelize.OpenFile("Book1.xlsx")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// Insert a picture.
// 	if err := f.AddPicture("Sheet1", "A2", "image.png", nil); err != nil {
// 		fmt.Println(err)
// 	}
// 	// Insert a picture to worksheet with scaling.
// 	if err := f.AddPicture("Sheet1", "D2", "image.jpg",
// 		&excelize.GraphicOptions{ScaleX: 0.5, ScaleY: 0.5}); err != nil {
// 		fmt.Println(err)
// 	}
// 	// Insert a picture offset in the cell with printing support.
// 	enable, disable := true, false
// 	if err := f.AddPicture("Sheet1", "H2", "image.gif",
// 		&excelize.GraphicOptions{
// 			PrintObject:     &enable,
// 			LockAspectRatio: false,
// 			OffsetX:         15,
// 			OffsetY:         10,
// 			Locked:          &disable,
// 		}); err != nil {
// 		fmt.Println(err)
// 	}
// 	// Save the spreadsheet with the origin path.
// 	if err = f.Save(); err != nil {
// 		fmt.Println(err)
// 	}
// }
