package forms

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form is a type struct that holds the data from forms and errors
type Form struct {
	url.Values
	Errors errors
}

// New Initializes an empty Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks all the required fields passed to it
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, fmt.Sprintf("This field %s cannot be blank", field))
		}
	}
}

// MinLength checks the minimum length of characters in the field
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field %s must be at least %d characters long.", field, length))
		return false
	}
	return true
}

// MaxLength checks the maximum length of characters in the field
func (f *Form) MaxLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) > length {
		f.Errors.Add(field, fmt.Sprintf("This field %s must be at least %d characters long.", field, length))
		return false
	}
	return true
}

// MinFloatValue checks the minimum value of characters in the field
func (f *Form) MinFloatValue(field string, value float64) bool {
	actualValue, err := strconv.ParseFloat(f.Get(field), 64)
	if err != nil {
		f.Errors.Add("This field %s must be floating point number", field)
		return false
	}
	if actualValue < value {
		f.Errors.Add(field, fmt.Sprintf("This field %s must be greater and equal to %f", field, value))
		return false
	}
	return true
}

// MaxFloatValue checks the maximum floating value of characters in the field
func (f *Form) MaxFloatValue(field string, value float64) bool {
	actualValue, err := strconv.ParseFloat(f.Get(field), 64)
	if err != nil {
		f.Errors.Add("This field %s must be floating point number", field)
		return false
	}
	if actualValue > value {
		f.Errors.Add(field, fmt.Sprintf("This field %s must be less and equal to %f", field, value))
		return false
	}
	return true
}

// Has checks if the form field is empty or not
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		f.Errors.Add(field, fmt.Sprintf("This field %s cannot be black", field))
		return false
	}
	return true
}

// ValidatePhone validate phone number using regex
func (f *Form) ValidatePhone(field string) bool {
	// defining the pattern for validating phone
	pattern := `^\+?[1-9]\d{1,14}$`

	// create a regex object
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(f.Get(field)) {
		f.Errors.Add(field, "Enter a valid phone number")
		return false
	}
	return true
}

// HasUpperCase checks if the value of field consist of one upper case
func (f *Form) HasUpperCase(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		exp, err := regexp.Compile("([A-Z])")
		if err != nil {
			log.Println(err)
		}
		u := exp.FindAllString(x, 1)
		if len(u) == 0 {
			f.Errors.Add(field, fmt.Sprintf("This field %s must have at least one upper case character", field))
		}
	}
}

// HasLowerCase checks if the value of the field has at least one lower case character
func (f *Form) HasLowerCase(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		exp, err := regexp.Compile("([a-z])")
		if err != nil {
			log.Println(err)
		}
		u := exp.FindAllString(x, 1)
		if len(u) == 0 {
			f.Errors.Add(field, fmt.Sprintf("This field %s must have at least one lower case character", field))
		}
	}
}

// HasNumber checks if the value of the field has at least one number
func (f *Form) HasNumber(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		exp, err := regexp.Compile("([0-9])")
		if err != nil {
			log.Println(err)
		}
		u := exp.FindAllString(x, 1)
		if len(u) == 0 {
			f.Errors.Add(field, fmt.Sprintf("This field %s must have at least one number", field))
		}
	}
}

func (f *Form) HasSpecialCharacter(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		exp, err := regexp.Compile("([!@#$%^&*.?-])+")
		if err != nil {
			log.Println(err)
		}
		u := exp.FindAllString(x, 1)
		if len(u) == 0 {
			f.Errors.Add(field, fmt.Sprintf("This field %s must have at least special characters", field))
		}
	}
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

// Valid reutrns true if there are no errors else returns false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
