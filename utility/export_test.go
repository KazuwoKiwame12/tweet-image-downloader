package utility

import "net/http"

var ExportValidateEmptyParamValue func(*ParamBuilder) (string, error) = (*ParamBuilder).validateEmptyParamValue
var ExportCreateResponse func(*TwitterClient, Conditions) (*http.Request, error) = (*TwitterClient).createRequest
