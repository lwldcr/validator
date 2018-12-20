package validator

import (
	"io/ioutil"
	"net/http"
)

const (
	SignParam = "sign"
)

type Validator struct {
	Checker checker
}

// checker interface, user need to implement this according to their own signature algorithms
type checker interface {
	Check([]byte, string) error
}

func NewValidator(c checker) *Validator {
	return &Validator{
		Checker:c,
	}
}

// Validate checks given data with internal checker
func (v *Validator) Validate(dat []byte, sign string) error {
	if err := v.Checker.Check(dat, sign); err != nil {
		return err
	}
	return nil
}

// ServeHTTP implements http.Handler interface
func (v *Validator) ServeHTTP(w http.ResponseWriter, r *http.Request, h http.HandlerFunc) {
	r.ParseForm()
	sign := r.FormValue(SignParam)
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil { // handle error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	if err := v.Validate(bs, sign); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	h(w, r)
}