package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	//cm "github.com/swinsome/skynet/common"
)

const url string = "http://web.ifzq.gtimg.cn/appstock/app/fqkline/get?_var=kline_dayhfq2007&param=sh600016,day,%s,%s,640,qfq&r=0.14717337880617885"

type (
	ifzqData struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			StockData struct {
				Qfqday [][]string `json:"qfqday"`
				Qt     struct {
					StockInfo []string `json:"stockinfo"`
				} `json:"qt"`
			} `json:"stockcode"`
		} `json:"data"`
	}
	IfzqRequest struct {
		code     string
		start    string
		end      string
		ktype    string
		autotype string
		retry    int
		pause    float64
		Name     string
	}
)

func (r IfzqRequest) genUrl() ([]string, error) {
	var result []string
	splitDate, err := _split_date(r.start, r.end, "2006-01-02", 640)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	fmt.Printf("%v\n", splitDate)
	for _, v := range splitDate {
		result = append(result, fmt.Sprintf(url, v.Start, v.End))
	}
	return result, nil
}
func (r IfzqRequest) preprocess(buffer string, head string) (string, error) {
	n := strings.Index(buffer, "=")
	if n == -1 {
		n = 0
	}
	out := string([]byte(buffer)[n+1:])
	pat := []string{",{\"nd\":[^}]+}", "\"mx price\":{[^}]+}"}
	for _, v := range pat {
		re, _ := regexp.Compile(v)
		out = re.ReplaceAllString(out, "")
	}
	out = strings.Replace(out, head, "stockcode", 1)
	out = strings.Replace(out, head, "stockinfo", 1)
	return out, nil
}
func (r *IfzqRequest) unmarshal(buffer string) (Stock, error) {
	var out ifzqData
	var result Stock
	//fmt.Println(buffer)
	err := json.Unmarshal([]byte(buffer), &out)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	//fmt.Println(out.Data.StockData.Qt.StockInfo[1])
	result.Name = out.Data.StockData.Qt.StockInfo[1]
	result.Raw = out.Data.StockData.Qfqday
	result.Data, err = _Convert_2_real(out.Data.StockData.Qfqday)
	if err != nil {
		return result, err
	}
	return result, nil
}

func _Get_k_data(u UrlGet) (Stock, error) {
	urlStrs, err := u.genUrl()
	var buffer string
	var result Stock
	var tmp Stock
	//fmt.Printf("%v\n", urlStr)
	for _, v := range urlStrs {
		resp, err := http.Get(v)
		if err != nil {
			return result, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return result, err
		}
		buffer = string(body)
		buffer, err = u.preprocess(buffer, "sh600016")
		tmp, err = u.unmarshal(buffer)
		(&result).append(tmp)
	}
	return result, err
}
