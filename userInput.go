package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ncruces/zenity"
)

const RANGE_SEPARATOR = "-"

func parseRangeFilter(input string) (*RangeFilter, error) {
	minAndMax := strings.Split(input, RANGE_SEPARATOR)

	if len(minAndMax) != 2 {
		return nil, errors.New("Invalid range input! input given: " + input)
	}

	min, err := strconv.Atoi(minAndMax[0])
	if err != nil {
		return nil, err
	}

	max, err := strconv.Atoi(minAndMax[1])
	if err != nil {
		return nil, err
	}

	return &RangeFilter{min, max}, nil
}

func parseFilter(filterString string) (Filter[int], error) {
	if strings.Contains(filterString, RANGE_SEPARATOR) {
		rangeFilter, err := parseRangeFilter(filterString)
		if err != nil {
			return nil, err
		}

		return rangeFilter, nil
	}

	allowedNumber, err := strconv.Atoi(filterString)
	if err != nil {
		return nil, err
	}

	return &SingleValueFilter[int]{allowedNumber}, nil
}

func getParamFilterFromUser(title string) (Filter[int], error) {
	answer, err := zenity.Entry(title)
	if err != nil {
		return nil, err
	}

	answer = strings.ReplaceAll(answer, " ", "") // removing all white space

	// No input means no filter
	if answer == "" {
		return &NoFilter[int]{}, nil
	}

	// separating by commas
	stringFilters := strings.Split(answer, ",")

	var filterCombiner FilterCombiner[int]

	for _, stringFilter := range stringFilters {
		fmt.Println(stringFilter)
		filter, err := parseFilter(stringFilter)
		if err != nil {
			return nil, err
		}
		filterCombiner.Filters = append(filterCombiner.Filters, filter)
	}

	return &filterCombiner, nil
}

func getQuestionFilterFromUser() (*QuestionFilter, error) {
	event, err := getParamFilterFromUser("Event")
	if err != nil {
		return nil, err
	}
	form, err := getParamFilterFromUser("Form")
	if err != nil {
		return nil, err
	}
	questionIndex, err := getParamFilterFromUser("Question Index")
	if err != nil {
		return nil, err
	}

	return &QuestionFilter{
		event,
		form,
		questionIndex,
	}, nil
}
