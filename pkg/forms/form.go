package forms

import (
	"net/url"
	"strings"
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
	if len(value) > d {
		f.Error.Add(value, "content is too long.")
	}
}
