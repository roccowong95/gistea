package gist

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/roccowong95/gistea/pkg"
)

var (
	_defaultCTimeoutSec = 5
	_defaultTimeoutSec  = 10
	_defaultDir         = "~/.config/gistea/files"
)

type Gist struct {
	*GistOpt
	client *github.Client
}

type GistOpt struct {
	Bg          bool
	Token       string
	Dir         string
	CTimeoutSec int
	TimeoutSec  int
}

func (o *GistOpt) validate() {
	if o.CTimeoutSec <= 0 {
		o.CTimeoutSec = _defaultCTimeoutSec
	}
	if o.TimeoutSec <= 0 {
		o.TimeoutSec = _defaultTimeoutSec
	}
	if len(o.Dir) <= 0 {
		o.Dir = _defaultDir
	}
}

func NewGist(opt *GistOpt) *Gist {
	opt.validate()
	httpClient := pkg.TimeoutClient(time.Duration(opt.CTimeoutSec)*time.Second, time.Duration(opt.TimeoutSec)*time.Second)
	if len(opt.Token) > 0 {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: opt.Token})
		ctx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)
		httpClient = oauth2.NewClient(ctx, ts)
	}
	ret := &Gist{client: github.NewClient(httpClient), GistOpt: opt}
	if opt.Bg {
		ret.start()
	}
	return ret
}

func (g *Gist) start() {
}
