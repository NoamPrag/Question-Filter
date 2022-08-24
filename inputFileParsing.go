package main

import (
	"strconv"
	"strings"
)

func findWhiteSpace(column []string) int {
	for i, text := range column {
		// Finding the start of the "id" list.
		if text == "id" {
			return i
		}
	}
	return -1
}

func isQuestion(title string) bool {
	return strings.Contains(title, "E") && strings.Contains(title, "_F") && strings.Contains(title, "_Q")
}

func findFirstQuestionColumn(titles []string) int {
	for i, title := range titles {
		if isQuestion(title) {
			return i
		}
	}
	return -1
}

func parseQuestion(question string) (*Question, error) {
	params := strings.Split(question, "_")

	event, err := strconv.Atoi(params[0][1:])
	if err != nil {
		return nil, err
	}

	form, err := strconv.Atoi(params[1][1:])
	if err != nil {
		return nil, err
	}

	questionIndex, err := strconv.Atoi(params[2][1:])
	if err != nil {
		return nil, err
	}

	return &Question{
		Event:         event,
		Form:          form,
		QuestionIndex: questionIndex,
	}, nil
}
