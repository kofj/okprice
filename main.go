package main

import "net/http"
import "encoding/json"
import "fmt"
import "strings"

const marketV2API = `https://www.okcoin.cn/v2/markets/market-tickers`
const okurl = "https://www.okcoin.cn/"
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
		fmt.Println(err)
		return
	}
	resp.Body.Close()

	var head, all, link string

	for _, v := range v.Data {
		var color string
		name := string(v.Symbol[:3])
		switch v.Change[:1] {
		case "+":
			color = "green"

		case "-":
			color = "red"

		default:
			color = "gray"
		}

		if in(name, []string{"btc", "ltc", "eth"}) {
			head += fmt.Sprintf("%s%s ", name2symbol(name), v.Last)
		}

		switch name {
		case "btc":
			link = okurl + "/trade/btc.do"
		case "ltc":
			link = okurl + "/trade/ltc.do"
		case "eth":
			link = okurl + "/trade/eth.do"
		case "etc":
			link = okurl + "/spot/trade.do#etc"
		case "bcc":
			link = okurl + "/spot/trade.do#bcc"
		default:
			link = okurl
		}
		all += fmt.Sprintf("%s ￥%s %s (%s) Vol. %s |href=%s size=9 color=%s\n", strings.ToUpper(name), v.Last, v.Change, v.ChangePercentage, v.Volume, link, color)

	}

	fmt.Println(head, " | size=12", "\n---\n", "OKCoinCN Price|color=#ccc href=", okurl, "\n", all)
	fmt.Printf("---\n ☕️ Donate\n--BTC|href=bitcoin:%s\n--LTC|href=litecoin:%s", btc, ltc)
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
