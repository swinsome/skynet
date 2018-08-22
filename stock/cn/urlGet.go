package stock

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

type (
	Kdata struct {
		Date   string  `json:"date"`
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		High   float64 `json:"high"`
		Lower  float64 `json:"lower"`
		Volume float64 `json:"volume"`
	}
	Stock struct {
		Name string     `json:"name"`
		Data []Kdata    `json:"data"`
		Raw  [][]string `json:"raw"`
	}
	datePair struct {
		Start string
		End   string
	}
	UrlGet interface {
		genUrl() ([]string, error)
		preprocess(buffer string, head string) (string, error)
		unmarshal(buffer string) (Stock, error)
	}
)

func (r *Stock) append(l Stock) {
	if r.Name == "" {
		r.Name = l.Name
	}
	r.Data = append(r.Data, l.Data...)
	r.Raw = append(r.Raw, l.Raw...)
}
func (d Stock) Print() {
	fmt.Printf("%s\n", d.Name)
	fmt.Printf("%8s%10s%15s%12s%11s%12s%10s\n", "序号", "日期", "开盘价", "收盘价", "最高价", "最低价", "成交手数")
	for i, v := range d.Raw {
		fmt.Printf("%8d:", i)
		for _, d := range v {
			fmt.Printf("%15s", d)
		}
		fmt.Printf("\n")
	}
}

// func _convert_each_kdata_from_string(data []string) (Kdata,error){
// 	var tmp Kdata
// 	for i, v := range data {

// 	}
// }
func _split_date(start string, end string, layout string, cycle int) ([]datePair, error) {
	var result []datePair
	// var tmp datePair
	t_start, err := time.Parse(layout, start)
	if err != nil {
		return result, err
	}
	t_end, err := time.Parse(layout, end)
	if err != nil {
		return result, err
	}
	days := t_end.Sub(t_start).Hours() / 24
	times := int(math.Floor(days / float64(cycle)))
	if times == 0 {
		result = append(result, datePair{start, end})
		return result, nil
	}
	s_start := t_start
	s_end := t_start.Add(time.Duration(cycle) * 24 * time.Hour)
	result = append(result, datePair{s_start.Format(layout), s_end.Format(layout)})
	for i := 2; i <= times; i++ {
		s_start = s_end.Add(time.Hour * 24)
		s_end = s_end.Add(time.Hour * 24 * time.Duration(cycle))
		result = append(result, datePair{s_start.Format(layout), s_end.Format(layout)})
	}
	if s_end.Before(t_end) {
		s_start = s_end.Add(time.Hour * 24)
		result = append(result, datePair{s_start.Format(layout), end})
	}
	return result, nil
}
func _Convert_2_real(data [][]string) ([]Kdata, error) {
	var tmp Kdata
	var result []Kdata
	records := len(data)
	var err error
	if records == 0 {
		return result, errors.New("input [][]string is empty")
	}
	for _, v := range data {
		length := len(v)
		if length < 5 {
			return result, errors.New("each recode has not enough length")
		}
		tmp.Date = v[0]
		tmp.Open, err = strconv.ParseFloat(v[1], 32)
		if err != nil {
			return result, err
		}
		tmp.Close, err = strconv.ParseFloat(v[2], 32)
		if err != nil {
			return result, err
		}
		tmp.High, err = strconv.ParseFloat(v[3], 32)
		if err != nil {
			return result, err
		}
		tmp.Lower, err = strconv.ParseFloat(v[4], 32)
		if err != nil {
			return result, err
		}
		tmp.Volume, err = strconv.ParseFloat(v[5], 32)
		if err != nil {
			return result, err
		}
		result = append(result, tmp)
	}
	return result, nil
}
func Get_k_data(code string, start string, end string, ktype string, autotype string, retry int, pause float64) (Stock, error) {
	var request UrlGet
	r := IfzqRequest{code, start, end, ktype, autotype, retry, pause, ""}
	request = &r
	result, err := _Get_k_data(request)
	return result, err
}
