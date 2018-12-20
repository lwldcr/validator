package validator

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

var (
	ErrValidateFailed = fmt.Errorf("signature validate failed")
)

type Sha1Checker struct {
	Key string
}

func (e *Sha1Checker) Sign(dat []byte) (string, error) {
	s := sha1.New()

	if _, err := s.Write([]byte(e.Key)); err != nil {
		return "", err
	}

	if _, err := s.Write(dat); err != nil {
		return "", err
	}

	sha1Sum := s.Sum(nil)
	return base64.StdEncoding.EncodeToString(sha1Sum), nil
}

// Check implements checker interface
func (e *Sha1Checker) Check(dat []byte, sign string) error {
	localSign, err := e.Sign(dat)
	if err != nil {
		return err
	}
	fmt.Printf("got: %s, expected: %s", sign, localSign)
	if localSign == sign {
		return nil
	}
	return ErrValidateFailed
}