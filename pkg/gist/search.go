package gist

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/google/go-github/github"
)

func client() *http.Client {
	tr := &http.Transport{
		MaxIdleConns: 10,
		DialContext: (&net.Dialer{
			// tcp keepalive interval
			KeepAlive: 10 * time.Second,
			// connect timeout
			Timeout: 2 * time.Second,
		}).DialContext,
		// max idle time
		IdleConnTimeout: 5 * time.Minute,
	}
	ret := &http.Client{
		// max time before quitting a request
		Timeout:   time.Minute,
		Transport: tr,
	}
	return ret
}

func search(query string) (io.ReadCloser, error) {
	addr := "https://gist.github.com/search?q=" + url.QueryEscape("user:roccowong95 ") + query
	fmt.Println(addr)
	c := client()
	req, e := http.NewRequest("GET", addr, nil)
	if nil != e {
		return nil, e
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	resp, e := c.Do(req)
	if nil != e {
		return nil, e
	}
	return resp.Body, nil
}

func Search(query string) {
	r, e := search(query)
	if nil != e {
		panic(e)
	}
	defer r.Close()
	root, err := htmlquery.Parse(r)
	if nil != err {
		panic(err)
	}
	nodes := htmlquery.Find(root, `//*[@id="gist-pjax-container"]/div[2]/div/div/div[2]/div[*]`)
	if len(nodes) == 0 {
		fmt.Println(htmlquery.InnerText(root))
	}
	for _, node := range nodes {
		fmt.Println(node.Type)
		// v := htmlquery.FindOne(node, `./div[2]/div/div/table/tbody`)
		v := htmlquery.FindOne(node, `./div`)
		fmt.Printf("%p\n", v)
		fmt.Println(htmlquery.InnerText(v))
		// fmt.Println(htmlquery.InnerText(node))
	}
}

func T() {
	c := github.NewClient(http.DefaultClient)
}
