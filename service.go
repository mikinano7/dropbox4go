package dropbox4go

import (
	"net/http"
)

func New(c *http.Client, token string) *Service {
	return &Service{c: c, token: token}
}

//TODO: OAUTH 2.0
type Service struct {
	c      *http.Client
	token  string
}
