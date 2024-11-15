package httx

import (
	"net/http"
	"testing"
)

func TestDo(t *testing.T) {
	url := "https://www.baidu.com"
	h := http.Header{"User-Agent": []string{"curl/1.0.0"}}
	resp, err := Do("GET", url, nil, h)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	t.Logf("Status Code: %d \r\n", resp.StatusCode)
	//b, _ := io.ReadAll(resp.Body)
	//t.Logf("Context: %s \r\n", b)
	for c, d := range resp.Header {
		t.Logf("Header: %s : %s\r\n", c, d)
	}
	for _, e := range resp.Cookies() {
		t.Logf("Cookie: %s\r\n", e)
	}

	//log.Println(jar)
}
