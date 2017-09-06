package main

import "net/http"
import "encoding/json"
import "fmt"
import "strings"

const marketV2API = `https://www.okcoin.cn/v2/markets/market-tickers`
const okurl = "|href=https://www.okcoin.cn/"
const btc = `1AVqs4ZhfgKrspkQ3RTwnB5EM8ujaiiuhS`
const ltc = `LZ2vipxxiEHvPsf5xCuTp72M5j7zMJrnoP`

type Price struct {
	Buy              string `json:"buy"`
	Change           string `json:"change"`
	ChangePercentage string `json:"changePercentage"`
	CreatedDate      int64  `json:"createdDate"`
	DayHigh          string `json:"dayHigh"`
	DayLow           string `json:"dayLow"`
	High             string `json:"high"`
	Last             string `json:"last"`
	Low              string `json:"low"`
	Name             string `json:"name"`
	Sell             string `json:"sell"`
	Symbol           string `json:"symbol"`
	Volume           string `json:"volume"`
}

type PriceItems struct {
	Code      int     `json:"code"`
	Data      []Price `json:"data"`
	DetailMsg string  `json:"detailMsg"`
	Msg       string  `json:"msg"`
}

func main() {
	resp, err := http.Get(marketV2API)
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(resp.Body)

	var v = &PriceItems{}
	err = decoder.Decode(v)
	if err != nil {
		return
	}
	resp.Body.Close()

	var head, all string
	for _, v := range v.Data {
		name := string(v.Symbol[:3])
		if in(name, []string{"btc", "ltc", "eth"}) {
			head += fmt.Sprintf("%s%s ", name2symbol(name), v.Last)
		}

		all += fmt.Sprintf("%s ￥%s %s (%s) Vol. %s %s\n", strings.ToUpper(name), v.Last, v.Change, v.ChangePercentage, v.Volume, okurl)
	}

	fmt.Println(head, "\n---\n", "OKCoinCN Price\n", all)
	fmt.Printf("---\nDonate BTC|href=bitcoin:%s\nDonate LTC|href=litecoin:%s", btc, ltc)

}

func in(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func name2symbol(name string) (symbol string) {
	switch name {
	case "btc":
		symbol = "₿"
	case "ltc":
		symbol = "Ł"
	case "eth":
		symbol = "E"
	default:
		symbol = name
	}
	return
}