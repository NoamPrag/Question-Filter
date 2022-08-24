package main

import (
	"fmt"
	"strings"

	"github.com/ncruces/zenity"
	"github.com/xuri/excelize/v2"
)

const subjectIDColumnIndex = 2

type Question struct {
	Event              int
	Form               int
	QuestionIndex      int
	IndexInSpreadsheet int
}

func main() {
	inputFileName, err := zenity.SelectFile(
		zenity.FileFilters{
			{Name: "Excel files", Patterns: []string{"*.xlsx"}},
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	inputFile, err := excelize.OpenFile(inputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := inputFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	questionFilter, err := getQuestionFilterFromUser()
	if err != nil {
		fmt.Println(err)
		return // Stopping the program in case of an invalid input
	}

	sheet := inputFile.WorkBook.Sheets.Sheet[0].Name

	inputFileRows, err := inputFile.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
	}
	inputFileColumns := transpose(inputFileRows)

	// Eliminating white space at top
	startWhiteSpace := findWhiteSpace(inputFileColumns[0])
	inputFileRows = inputFileRows[startWhiteSpace:]

	inputFileColumns = transpose(inputFileRows)

	columnsWhichAreNotQuestions := findFirstQuestionColumn(inputFileRows[0])

	questionTitles := inputFileRows[0][columnsWhichAreNotQuestions:]

	var questions []*Question
	for i, questionTitle := range questionTitles {
		question, err := parseQuestion(questionTitle)
		question.IndexInSpreadsheet = i
		if err != nil {
			fmt.Println(err)
			continue
		}
		questions = append(questions, question)
	}

	filteredQuestions := questionFilter.filter(questions)

	outputFile := excelize.NewFile()
	outputFile.SetSheetName("Sheet1", sheet)

	writeColumn(inputFileColumns[subjectIDColumnIndex], outputFile, sheet, 1)

	for i, filteredQuestion := range filteredQuestions {
		column := inputFileColumns[filteredQuestion.IndexInSpreadsheet+columnsWhichAreNotQuestions]
		// plus 2, because first column is subject ID
		writeColumn(column, outputFile, sheet, i+2)
	}

	outputFileName, err := zenity.SelectFileSave(
		zenity.FileFilters{
			{Name: "Excel files", Patterns: []string{"*.xlsx"}},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	if !strings.HasSuffix(outputFileName, ".xlsx") {
		outputFileName += ".xlsx"
	}

	outputFile.SaveAs(outputFileName)
}
