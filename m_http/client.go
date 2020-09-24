package m_http

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func httpDemo() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com/", nil)

	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	client := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:", req)
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", s)
}
