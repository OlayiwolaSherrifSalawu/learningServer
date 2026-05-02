package forms

import "net/url"

type Form struct {
	url   url.Values
	Error errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		url:   data,
		Error: errors(map[string][]string{}),
	}
}
