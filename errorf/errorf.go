package errorf

import "fmt"

type ErrorFormatter struct {
	text string
	a    []any
}

func (e ErrorFormatter) Error() string {
	return fmt.Sprintf(e.text, e.a...)
}

func (e ErrorFormatter) Format(a ...any) ErrorFormatter {
	e.a = a
	return e
}

func New(text string, a ...any) ErrorFormatter {
	return ErrorFormatter{text, a}
}

var _ error = New("errorf: code %d", 114514)
