func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func httpPost(url string, params []byte) ([]byte, error) {
	reader := bytes.NewReader(params)
	resp, err := http.Post(url, "application/json", reader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func httpGetWithHeader(url string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "baidu.com")
	req.Header.Set("Cookie", "CMTOKEN-test="+CMCookie)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
func httpPostWithHeader(url string, param string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", "baidu.com")
	req.Header.Set("Content-Type", "application/json;charset=utf-8;")
	req.Header.Set("Cookie", "CMTOKEN-test=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjUsImNpdHkiOiIiLCJmYWNlMjAwIjoiIiwiZmFjZUltZyI6Ii8vdGVzdGFwaS5wbGF5Y29kZW1vbmtleS5jbi9pbWFnZXMvYXZhdGFyL2F2YXRhcjAwLnBuZyIsImdlbmRlciI6IueUtyIsImdyYWRlIjoi5pyq5YWl5a2mIiwiaWF0IjoxNTcxMTM1MDcyLCJqdGkiOiIxNTcxMTM1MDcyNTc2MzczNTEzIiwibmlja05hbWUiOiIxMzAqKioqMDY1MyIsInByb3ZpbmNlIjoiIiwicmVnaXN0ZXJlZCI6ZmFsc2UsInJlbW90ZVVpZCI6IiIsInJvbGUiOiJzdHVkZW50IiwidWlkIjoxNjkzNTUyLCJ1cmwiOiIiLCJ1c2VybmFtZSI6IjEzMDIwMDkwNjUzIn0.hHSqogUzUuzOGjON2EZdAhcwfGR0NnHbtHCvvTglJuo")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
