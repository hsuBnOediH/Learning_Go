package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "this filed can not be blank ")
		return false
	}
	return true
}

func (f *Form) Valid() bool {

	return len(f.Errors) == 0
}

func (f *Form) Requied(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "this field can not be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("this field must be at least %d char long", length))
		return false
	}
	return true
}


func (f *Form) IsEmail(field string){
	if !govalidator.IsEmail(f.Get(field)){
		f.Errors.Add(field,"Invalid Email address")
	}

}