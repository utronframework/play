package play

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gernest/utron/controller"
)

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
			"post;/compile;Compile",
			"post;/share;Share",
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

func (s *Service) compile(src []byte) ([]byte, error) {
	u := make(url.Values)
	u.Add("body", string(src))
	u.Add("version", "2")
	res, err := http.PostForm(s.basePath+"/compile", u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (s *Service) Compile() {
	r := s.Ctx.Request()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	nb, err := s.format(b)
	if err != nil {
		log.Fatal(err)
	}
	fres := make(map[string]interface{})
	json.Unmarshal(nb, &fres)
	o, err := s.compile([]byte(fres["Body"].(string)))
	if err != nil {
		log.Fatal(err)
	}
	s.Ctx.Response().Write(o)
}

func (s *Service) Share() {
	r := s.Ctx.Request()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	nb, err := s.share(b)
	if err != nil {
		log.Fatal(err)
	}
	s.Ctx.Response().Write(nb)
}

func (s *Service) share(src []byte) ([]byte, error) {
	res, err := http.Post(s.basePath+"/share", "text/plain", bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
