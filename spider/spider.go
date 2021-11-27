package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func GetVerb() {
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/")
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=sb&age=19")
	/*resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=周林&age=19")
	if err != nil {
		fmt.Println("get url failed, err:%v\n", err) // 服务端响应的body
		return
	}*/
	data := url.Values{} // url encode
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "周林")
	data.Set("age", "9000")
	queryStr := data.Encode() // URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// 请求不是特别频繁,用完就关闭该连接
	// 禁用 KeepAlive的client
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	// 走默认的cient
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}
	defer resp.Body.Close() // 一定记得要关闭resp.Body
	// 发请求
	// 从resp中把服务返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}
