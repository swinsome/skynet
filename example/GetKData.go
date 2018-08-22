package test

import (
	"fmt"

	iq "github.com/swinsome/skynet/stock/cn"
)

func test1() {
	_, err := iq.Get_k_data("sh600016", "2008-06-01", "2017-06-07", "D", "qfq", 3, 0.01)
	if err != nil {
		fmt.Println(err)
	} else {
		//data.Print()
	}

}
