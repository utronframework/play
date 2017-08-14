package play

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gernest/utron/controller"
)

// maxSnippetSize value taken from
// https://github.com/golang/playground/blob/master/app/goplay/share.go
const maxSnippetSize = 64 * 1024

type Service struct {
	controller.BaseController
	basePath string
	Routes   []string
}

func New() controller.Controller {
	return &Service{
		basePath: "https://play.golang.org",
		Routes: []string{
			"post;/fmt;Format",
		},
	}
}

func (s *Service) Format() {
	r := s.Ctx.Request()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	o, err := s.format(b)
	if err != nil {
		log.Fatal(err)
	}
	s.Ctx.Response().Write(o)
}

func (s *Service) format(src []byte) ([]byte, error) {
	u := make(url.Values)
	u.Add("imports", "true")
	u.Add("body", string(src))
	res, err := http.PostForm(s.basePath+"/fmt", u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
