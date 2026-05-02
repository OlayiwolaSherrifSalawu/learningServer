package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Error errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, val := range fields {
		value := f.Get(val)
		if value == "" {
			return
		}
		if strings.TrimSpace(value) == "" {
			f.Error.Add(value, "This Field is required it cant be empty")
		}
	}
}

func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Error.Add(value, fmt.Sprintf("The length of text is too long maximum length is %d", d))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Error.Add(field, "This field is invalid")
}

func (f *Form) Valid() bool {
	return len(f.Error) == 0
}
