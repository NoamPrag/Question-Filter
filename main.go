package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ncruces/zenity"
	"github.com/xuri/excelize/v2"
)

const sheet string = "Scheduled Events"
const columnsWhichAreNotQuestions = 19
const subjectIDColumnIndex = 2

type Question struct {
	Event         int
	Form          int
	QuestionIndex int
}

func getFilterParamFromUser(title string) int {
	answer, err := zenity.Entry(title)
	if err != nil {
		fmt.Println(err)
	}

	param, err := strconv.Atoi(answer)
	if err != nil {
		return -1
	}
	return param
}

func getQuestionFilterFromUser() *Question {
	event := getFilterParamFromUser("Event")
	form := getFilterParamFromUser("Form")
	questionIndex := getFilterParamFromUser("Question Index")

	return &Question{
		Event:         event,
		Form:          form,
		QuestionIndex: questionIndex,
	}
}

func parseQuestion(question string) *Question {
	params := strings.Split(question, "_")

	event, err := strconv.Atoi(params[0][1:])
	if err != nil {
		fmt.Println(err)
	}

	form, err := strconv.Atoi(params[1][1:])
	if err != nil {
		fmt.Println(err)
	}

	questionIndex, err := strconv.Atoi(params[2][1:])
	if err != nil {
		fmt.Println(err)
	}

	return &Question{
		Event:         event,
		Form:          form,
		QuestionIndex: questionIndex,
	}
}

func filterQuestions(filter *Question, questions []*Question) []int {
	var filteredQuestionsIndices []int

	for questionIndex, question := range questions {
		if (filter.Event == -1 || question.Event == filter.Event) && (filter.Form == -1 || question.Form == filter.Form) && (filter.QuestionIndex == -1 || question.QuestionIndex == filter.QuestionIndex) {
			filteredQuestionsIndices = append(filteredQuestionsIndices, questionIndex)
		}
	}
	return filteredQuestionsIndices
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			if len(slice[j]) < i+1 {
				result[i][j] = ""
			} else {
				result[i][j] = slice[j][i]
			}
		}
	}
	return result
}

func getCellAxis(columnName string, row int) string {
	return columnName + strconv.Itoa(row)
}

func writeColumn(column []string, outputFile *excelize.File, columnIndexInOutputFile int) {
	columnName, err := excelize.ColumnNumberToName(columnIndexInOutputFile)
	if err != nil {
		fmt.Println(err)
	}

	for rowIndex, value := range column {
		outputFile.SetCellValue(sheet, getCellAxis(columnName, rowIndex+1), value)
	}
}

func main() {
	inputFileName, err := zenity.SelectFile(
		zenity.FileFilters{
			{Name: "Excel files", Patterns: []string{"*.xlsx"}},
		})
	if err != nil {
		fmt.Println(err)
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

	questionFilter := getQuestionFilterFromUser()

	rows, err := inputFile.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
	}
	questionTitles := rows[0][columnsWhichAreNotQuestions:]

	var questions []*Question
	for _, questionTitle := range questionTitles {
		questions = append(questions, parseQuestion(questionTitle))
	}

	filteredQuestionIndices := filterQuestions(questionFilter, questions)

	inputFileColumns := transpose(rows)

	outputFile := excelize.NewFile()
	outputFile.SetSheetName("Sheet1", sheet)

	writeColumn(inputFileColumns[subjectIDColumnIndex], outputFile, 1)

	for i, filteredQuestionIndex := range filteredQuestionIndices {
		column := inputFileColumns[filteredQuestionIndex+columnsWhichAreNotQuestions]
		// plus 2, because first column is subject ID
		writeColumn(column, outputFile, i+2)
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
