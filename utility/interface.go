//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package utility

import "net/http"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
