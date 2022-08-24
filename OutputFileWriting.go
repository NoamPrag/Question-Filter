package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func getCellAxis(columnName string, row int) string {
	return columnName + strconv.Itoa(row)
}

func writeColumn(column []string, outputFile *excelize.File, sheet string, columnIndexInOutputFile int) {
	columnName, err := excelize.ColumnNumberToName(columnIndexInOutputFile)
	if err != nil {
		fmt.Println(err)
	}

	for rowIndex, value := range column {
		outputFile.SetCellValue(sheet, getCellAxis(columnName, rowIndex+1), value)
	}
}
