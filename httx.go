package httx

import (
	"io"
	"log"
	"net/http"
	"time"
)

func Do(method string, url string, body io.Reader, header http.Header) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	defer cancel()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Jar:     jar,
		Timeout: time.Second * 5,
	}
	if header != nil {
		for k, v := range header {
			req.Header[k] = v
		}
	}
	resp, err = client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			log.Printf("请求超时: %v\n", ctx.Err())
		default:
			log.Printf("请求失败: %v\n", err)
		}
		return nil, err
	}
	return resp, nil
}
