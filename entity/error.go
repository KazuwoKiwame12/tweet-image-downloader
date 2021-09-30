package entity

import (
	"errors"
	"fmt"
)

var (
	ErrorInputMaxFlag        error = errors.New("error: you should input number between 10 and 100")
	ErrorLoadEnv             error = errors.New("error: you didn't get any tweet. check url parameters")
	ErrorNoTweet             error = errors.New("error: you didn't get any tweet. check url parameters")
	ErrorEmptyParameterValue error = errors.New("error: url parameter has empty value")
)

func ErrorIsnotIdealStatusCode(code int) error {
	return fmt.Errorf("status code is not 200. result is %d", code)
}
