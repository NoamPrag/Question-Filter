package main

type Filter[T any] interface {
	Predicate(value T) bool
}

type NoFilter[T any] struct{}

func (f *NoFilter[T]) Predicate(value T) bool {
	return true
}

type SingleValueFilter[T comparable] struct {
	allowedValue T
}

func (f *SingleValueFilter[T]) Predicate(value T) bool {
	return value == f.allowedValue
}

type RangeFilter struct {
	min int
	max int
}

func (f RangeFilter) Predicate(num int) bool {
	return num >= f.min && num <= f.max
}

// Combines all its filters with logical OR operator
type FilterCombiner[T any] struct {
	Filters []Filter[T]
}

func (f *FilterCombiner[T]) Predicate(value T) bool {
	for _, filter := range f.Filters {
		if filter.Predicate(value) {
			return true
		}
	}
	return false
}

type QuestionFilter struct {
	Event         Filter[int]
	Form          Filter[int]
	QuestionIndex Filter[int]
}

func (f QuestionFilter) Predicate(question Question) bool {
	return f.Event.Predicate(question.Event) && f.Form.Predicate(question.Form) && f.QuestionIndex.Predicate(question.QuestionIndex)
}

func (f QuestionFilter) filter(questions []*Question) []*Question {
	var filteredQuestions []*Question
	for _, question := range questions {
		if f.Predicate(*question) {
			filteredQuestions = append(filteredQuestions, question)
		}
	}
	return filteredQuestions
}
