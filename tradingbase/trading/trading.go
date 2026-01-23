package trading

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	token         = "122c1435a9cfbdc38f2fc4ab4eedf2e36239d9df47025d3d523b594dbbb9e911"
	optionStock   = "NIFTY20JAN26P25550" //"NIFTY18NOV25P25800"
	id            = "FZ23969"
	niftyOneLot   = "65"
	niftyTwoLot   = "130"
	niftyThreeLot = "195"
	//For rsi 1 minute timeframe, placing last candle close - 2 rupees limit order, if the next candle current price is less than -2 rupees limit order, order executed immediately with even lesser price.
	rsiLimitlower = 7.0 //for 5 minute timeframe RSI(3)
	rsiLimitlow   = 25.0

	NiftyExchange                    = "NSE"
	OptionsExchange                  = "NFO"
	Nifty50                          = "Nifty 50"
	NiftyYesterdayCandle             Candle
	PreviousDayHigh                  float64 //previous day nifty50 high
	PreviousDayLow                   float64
	location                         *time.Location
	fileName                         = "tradinglogs.txt"
	f                                *os.File
	LastRun                          = time.Now()
	PreviousDay                      int
	year                             int
	day                              int
	month                            time.Month
	botToken                         = "8302802854:AAF_nCOnPD2KuUEsh8QP582aPZAnUP5Ye4E"
	chatID                           = "5987150863"
	RunningInterval                  int
	OptionsStockOI                   = 100000.0
	OptionsStockVolume               = 5000.0
	OptionsStockStopLossPercentage   = 4.0
	OptionsStockTargetPercentage     = 30.0
	OptionsStockLimitPricePercentage = 4.0
	optionsStockMonth                = "27JAN26" //should be last tuesday of month
	lotSize                          map[string]string
	NiftyTradingOpenFlag             = false
	diff                             float64
)

func init() {
	var err error
	location, err = time.LoadLocation("Asia/Kolkata") // IST zone
	if err != nil {
		panic(err)
	}

	now := time.Now().In(location)
	year, month, day = now.Date()

	PreviousDay = day - 1

	lotSize, err = LoadLotSizeMap("d:/GolangPractice/go/src/GolangWay/tradingbase/NiftyLotSize.csv")
	if err != nil {
		fmt.Println("error in fetching the options lot size", err)
		return
	}

	//Getting nifty 50 (or) any stock previous day's close, high, low
	//getTimestamp(4, 15) - One day before previous day - 4th day of a month
	//getTimestamp(5, 15) - previous day - 5th day of a month
	tmp1 := GetEODData(GetTimestamp(PreviousDay-1, 15), GetTimestamp(PreviousDay, 15), Nifty50)
	// fmt.Printf("EOD chart data--------%+v\n", tmp1)
	NiftyYesterdayCandle = GetCandleDataInFloat64([]Candle1{tmp1})
	// fmt.Println("NiftyYesterdayCandle--------", NiftyYesterdayCandle)
	PreviousDayHigh = NiftyYesterdayCandle.Inth[0]
	PreviousDayLow = NiftyYesterdayCandle.Intl[0]

	fmt.Println("Nifty 50 Previous day values --------", NiftyYesterdayCandle)
	f.WriteString(fmt.Sprintf("Nifty 50 Previous day values --------%+v\n", NiftyYesterdayCandle))

}

func HandlerLogic() {

	var err error
	// Append to file
	f, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	logLine := fmt.Sprintf("--------------------- %v --------------------------------------------------\n", time.Now())
	f.WriteString(logLine)

	// 	// t225 := time.Date(year, month, day, 14, 25, 0, 0, location)

	// 	// //this logic runs only before 2.25pm
	// 	// for {

	// 	// 	if time.Now().In(location).After(t225) {
	// 	// 		fmt.Println("Thats it, Trading closing for the day----------------------")
	// 	// 		break
	// 	// 	}

	// 	// 	// RSILogicTrade()
	// 	// 	// latestTime := time.Now()
	// 	// 	// // reversal logic will run every 15 minutes once
	// 	// 	// if latestTime.Sub(lastRun) >= 5*time.Minute {
	// 	// 	// 	fmt.Println("running ReversalORFallLogic()", latestTime, lastRun, latestTime.Sub(lastRun))
	// 	// 	// 	f.WriteString(fmt.Sprintf("running ReversalORFallLogic() ------------------:%+v\n", []interface{}{latestTime, lastRun, latestTime.Sub(lastRun)}))
	// 	// 	// 	stopFlag := ReversalORFallLogic()

	// 	// 	// 	if stopFlag {
	// 	// 	// 		SendTelegramMessage("Nifty or options reversal observed --" + optionStock)
	// 	// 	// 		break
	// 	// 	// 	}

	// 	// 	// 	lastRun = time.Now()
	// 	// 	// }

	// 	// 	sellOrders()

	// 	// 	intradayLogic()

	// 	// }

	// 	// ExitOrder("26010800217784")
	// 	// CancelOrder("26010800217784")

	ctx, cancel := context.WithCancel(context.Background())

	// Handle CTRL+C / SIGTERM
	go HandleShutdown(cancel)

	now := time.Now().In(location)
	interval := GetInterval(now)

	APIFetchBuffer := 4 * time.Second //ensure api fetch data for EX: 10.15am data for 10.15.04 seconds

	//EX:5 min - ensures 5 min candle starts at correct time, also works for 1 min candle also
	firstRun := nextAlignedTime(interval, APIFetchBuffer, location)

	time.Sleep(time.Until(firstRun))

	// fmt.Println("firstrun starts -------", time.Until(firstRun), "-----------", now)

	go RunTask(ctx, 1*time.Minute, NiftyCrossCheck)

	//doing one time initial run then below runtask maintains the run interval correct
	// IntradayLogic()

	// go RunTask(ctx, interval, SellOrders)

	go RunTask(ctx, interval, IntradayLogic)

	//sets default 5 minutes as of now
	go RunTask(ctx, 5*time.Minute, IntradayLogicVersion1)

	// go RunTradingScheduler(ctx, SellOrders)
	// go RunTradingScheduler(ctx, IntradayLogic)

	// fmt.Println("waiting------------")
	<-ctx.Done()
	fmt.Println("All tasks stopped safely")

}

func nextAlignedTime(interval time.Duration, buffer time.Duration, loc *time.Location) time.Time {
	now := time.Now().In(loc)

	elapsed := now.Truncate(interval)
	next := elapsed.Add(interval)

	return next.Add(buffer)
}

func Dummy() {
	fmt.Println("dummy-------")
	// GetOptionChain("IDEA-EQ", 11, 71475)
	// GetOptionChain("WIPRO-EQ", 265, 3000)
	// GetOptionChain("IDEA-EQ", 11, 1)
	// GetOptionChain("WIPRO-EQ", 265, 1)

	// GetOptionGreek()

	// "WIPRO27JAN26C260"
	// "RELIANCE27JAN26C1900"
	// "WAAREEENER27JAN26C2500"
	// tmp, _ := GetCandleData(GetTimestamp(), "WAAREEENER27JAN26C2500", 1, OptionsExchange)
	// tmp, _ := GetCandleData(GetTimestamp(), "NIFTY20JAN26C25900", 1, OptionsExchange)
	// candleData := GetCandleDataInFloat64(tmp)
	// fmt.Println(candleData)

	// PlaceOrder(optionStock, niftyOneLot, 43.55, 3.55, 43.116666666666, 43.116666666666/2)

}

func GetCurrentTimestamp() int64 {
	now := time.Now().In(location)
	return now.Unix()

}

func StockOptionsFetchAndValidation() (map[string]string, error) {

	d, err := NseGet()

	if err != nil {
		fmt.Println("error in fetching stocks-----------------:", err)
		return nil, err
	}
	PickedStocks := map[string]string{}
	// fmt.Println(t.FilterFOStocks(d))
	CEStocks, PEStocks, err := FilterFOStocks(d)
	if err != nil {
		fmt.Println("error in filtering stocks-----------------:", err)
		return nil, err
	}

	for symbol, _ := range CEStocks {

		lastPriceFloat, _ := strconv.ParseFloat(CEStocks[symbol], 64)
		lastPriceInt := int(lastPriceFloat)
		strikePrice := CalculateStrikePriceDifferences(lastPriceInt)

		maxValue := lastPriceInt + (5 * strikePrice)

		//it will round the value for strike price interval EX: 2437 value will be rounded to 2440 for CE
		if strikePrice > 1 {
			lastPriceInt = (lastPriceInt / strikePrice) * strikePrice
			lastPriceInt = lastPriceInt + strikePrice
		}

		for i := lastPriceInt; i <= maxValue; i += strikePrice {
			// fmt.Println(maxValue, "----------------", i)
			priceValue := strconv.Itoa(int(i))
			sym := symbol + optionsStockMonth + "C" + priceValue
			tmp, err := GetCandleData(GetCurrentTimestamp()-120, sym, 1, OptionsExchange)
			// fmt.Println(len(tmp), tmp, "------------------", err, "-------", sym)
			time.Sleep(10 * time.Microsecond)
			if len(tmp) > 0 && err == nil { //if we found valid ATM or slight OTM strike price
				PickedStocks[symbol] = sym
				break
			}

		}

	}

	for symbol, _ := range PEStocks {

		lastPriceFloat, _ := strconv.ParseFloat(PEStocks[symbol], 64)
		lastPriceInt := int(lastPriceFloat)
		strikePrice := CalculateStrikePriceDifferences(lastPriceInt)

		maxValue := lastPriceInt + (5 * strikePrice)

		if strikePrice > 1 {
			lastPriceInt = (lastPriceInt / strikePrice) * strikePrice
		}

		for i := lastPriceInt; i <= maxValue; i += strikePrice {
			priceValue := strconv.Itoa(int(i))
			sym := symbol + optionsStockMonth + "P" + priceValue
			//making a api call to ensure valid options strike
			tmp, err := GetCandleData(GetCurrentTimestamp()-120, sym, 1, OptionsExchange)
			// fmt.Println(len(tmp), tmp, "------------------", err, "-------", sym)
			time.Sleep(10 * time.Microsecond)
			if len(tmp) > 0 && err == nil { //if we found valid ATM or slight OTM strike price
				PickedStocks[symbol] = sym
				break
			}

		}

	}

	fmt.Println("found options stocks strikes--------------", len(CEStocks), len(PEStocks))
	fmt.Println("found options stocks strikes--------------", len(PickedStocks), PickedStocks)

	return PickedStocks, nil
}

func CalculateStrikePriceDifferences(price int) int {

	if price < 400 {
		return 1
	}

	// if price > 250 && price < 400 {
	// 	return 2.5
	// }

	if price > 400 && price < 1000 {
		return 5
	}

	if price > 1000 && price < 5000 {
		return 10
	}

	if price > 5000 {
		return 50
	}

	return 10
}

// "currently allows for nifty index options only - EX: "NIFTY20JAN26P25800","NIFTY06JAN26P25800"
func ReversalOption(symbol string) string {
	n := len(symbol)

	if !strings.HasPrefix(symbol, "NIFTY") || n != 18 {
		fmt.Println("currently allows for nifty index options only----------")
		return ""
	}

	if symbol[n-6] == 'P' {
		return symbol[:n-6] + "C" + symbol[n-5:]
	}
	if symbol[n-6] == 'C' {
		return symbol[:n-6] + "P" + symbol[n-5:]
	}
	return symbol
}

func RunTask(ctx context.Context, interval time.Duration, task func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {

		select {
		case <-ticker.C:
			task()
		case <-ctx.Done():
			return
		}
	}
}

func RunTradingScheduler(ctx context.Context, task func()) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	buffer := 5 * time.Second //EX: 10.00.00 time -- this buffer 5 second 10.00.05 will fetch that minute api

	for {
		now := time.Now().In(loc)

		interval := GetInterval(now)
		nextRun := NextAlignedRun(now, interval, buffer)

		timer := time.NewTimer(time.Until(nextRun))
		select {
		case <-timer.C:
			task()
		case <-ctx.Done():
			timer.Stop()
			return
		}
	}

	fmt.Println("completed------")
}

func GetInterval(now time.Time) time.Duration {
	h, m := now.Hour(), now.Minute()

	// 09:30 to 10:15 → 1 min
	if (h == 9 && m >= 30) || (h == 10 && m < 15) {
		RunningInterval = 1
		return 1 * time.Minute
	}

	// After 10:15 → 5 min
	RunningInterval = 5
	return 5 * time.Minute
}

func NextAlignedRun(now time.Time, interval time.Duration, buffer time.Duration) time.Time {
	base := now.Truncate(interval).Add(interval)
	return base.Add(buffer)
}

func HandleShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	cancel()
}

func main1() {

}

func SellOrders() {
	now := time.Now().In(location)
	fmt.Println("In sell order function----------", now)

	body := GetOrderBook()

	// fmt.Println("In sell order function----------", string(body))

	//no orders found
	if strings.Contains(string(body), "no data") {
		return
	}

	var orders []OrderDetails
	if err := json.Unmarshal([]byte(body), &orders); err != nil {
		panic(err)
	}

	t930 := time.Date(year, month, day, 9, 30, 0, 0, location)
	t1015 := time.Date(year, month, day, 10, 15, 0, 0, location)

	fmt.Println("In sell order----------", now, len(orders))

	for _, order := range orders {

		if order.Status == "CANCELED" || order.Status == "COMPLETE" || strings.Contains(order.Status, "REJECT") { //|| order.BuyOrSell == "S"
			continue
		}

		interval := 5

		// Between 9:30 AM and 10:15 AM strategies
		if now.After(t930) && now.Before(t1015) {
			interval = 1
		}

		tmp, _ := GetCandleData(GetTimestamp(), order.OrderSymbol, interval, OptionsExchange)
		candleData := GetCandleDataInFloat64(tmp)
		intcValues := candleData.Intc

		asma5 := SMA(intcValues, 5)
		asma9 := SMA(intcValues, 9)

		fmt.Println("sell order details", asma5, asma9, asma5 < asma9, "------------ order details ---", len(intcValues), intcValues, order)

		if asma5 < asma9 || intcValues[0] < asma9 {
			ModifyOrder(order.OrderNo, order.OrderSymbol, order.OrderQuantity, intcValues[0])
			continue
		}

		//if buy open order has not executed in 10 minutes, Am exiting the order
		if order.BuyOrSell == "B" && order.Status == "OPEN" && HasTimePassed(order.OrderTime, 10*time.Minute) {
			ExitOrder(order.OrderNo)
			// SendTelegramMessage("exit order executed, check it now" + order.OrderSymbol)
			msg := fmt.Sprintf("exit Open order - %v -- %v -- %v", order.OrderSymbol, now, intcValues[0])
			fmt.Println(msg)
			f.WriteString(msg)

		}

	}

	// time.Sleep( 1 * time.Minute)
}

func ExitOrder(orderNo string) {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/ExitSNOOrder"

	//"prd": "B" - brcket order, "H" - cover order
	data1 := `{
	    "uid": "FZ23969",
	    "prd": "B",
	    "norenordno": "%s"
	}`

	data := fmt.Sprintf(data1, orderNo)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	fmt.Println("exit order-------", string(payload))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("EXIT ORDER -- API Response in chart data:", resp.Status, string(body))

}

func CancelOrder(orderNo string) {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/CancelOrder"

	//"prd": "B" - brcket order, "H" - cover order
	data1 := `{
	    "uid": "FZ23969",
	    "norenordno": "%s"
	}`

	data := fmt.Sprintf(data1, orderNo)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("CANCEL ORDER -- API Response in chart data:", resp.Status, string(body))

}

func GetOptionChain(optionStockSymbol string, strikePrice float64, cnt int) {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/GetOptionChain"

	//"prd": "B" - brcket order, "H" - cover order
	data1 := `{
	    "uid": "FZ23969",
		"exch": "NSE", 
		"tsym": "%s",
		"strprc": "%.2f", 
		"cnt": "%d"    
	}`

	data := fmt.Sprintf(data1, optionStockSymbol, strikePrice, cnt)

	fmt.Println(data)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("GetOptionChain -- API Response in chart data:", resp.Status, string(body))

}

func GetOptionGreek() {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/GetOptionGreek"

	//not working
	data1 := `{
	"exd": "2026-01-12",    
	"strprc": "25600", 
    "sptprc": "25683",
    "int_rate": "0", 
    "volatility": "0", 
    "optt": "CE"  
	}`

	// data := fmt.Sprintf(data1, optionStockSymbol, strikePrice, cnt)

	// fmt.Println(data)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data1, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("GetOptionChain -- API Response in chart data:", resp.Status, string(body))

}

func GetOrderBook() []byte {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/OrderBook"

	data1 := `{
	    "uid": "FZ23969"
	}`

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data1, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("ORDER BOOK -- API Response in chart data:", resp.Status, string(body))

	return body

}

func GetTradeBook() {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/TradeBook"

	data1 := `{
	    "uid": "FZ23969",
	    "actid": "FZ23969"
	}`

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data1, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("TRADE BOOK --- API Response in chart data:", resp.Status, string(body))

}

type Candle2 struct {
	Intc int `json:"intc"` // closing price string
	Inth int `json:"inth"`
	Intl int `json:"intl"`
}

// searches with stocks by text - EX: WIPRO-EQ (-EQ ending - stock names only we needs for trading)
func SearchStockDetails(stockSearch string) {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/SearchScrip"

	data1 := `{
	    "uid": "FZ23969",
	    "exch": "NSE",
	    "stext": "%s"
	}`

	data := fmt.Sprintf(data1, stockSearch)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response in chart data:", string(body))

}

// gets particular stock details ohlc, last price,volume, total open interest etc details
func GetStockDetails(stockToken int) {
	url1 := "https://piconnect.flattrade.in/PiConnectTP/GetQuotes"

	data1 := `{
	    "uid": "FZ23969",
	    "exch": "NSE",
		"token":"%d"
	}`

	data := fmt.Sprintf(data1, stockToken)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response in chart data:", string(body))

}

func GetTopList() {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/TopList"

	data1 := `{
	    "uid": "FZ23969",
	    "exch": "NSE",
	    "tb": "T",
	    "bskt": "NSEEQ",
	    "crt": "LTP"
	}`

	// data := fmt.Sprintf(data1, optionStock, time930am, timeNow, interval)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data1, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response in chart data:", string(body))

}

func GetEODData(fromTime int64, toTime int64, stockOrIndex string) Candle1 {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/EODChartData"

	data1 := `{
	    "sym": "NSE:%s",
    	"from": "%d",
    	"to": "%d"
	}`

	data := fmt.Sprintf(data1, stockOrIndex, fromTime, toTime)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil || resp.StatusCode == 401 || resp.StatusCode == 400 {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("API Response in chart data:", string(body), body, resp.StatusCode)

	var arr []interface{}
	if err := json.Unmarshal(body, &arr); err != nil {
		log.Fatal("outer unmarshal error:", err)
	}

	// fmt.Println("Unmarshal EOD data in interface 000000-------", arr)

	// Step 2: extract first element as string
	s, ok := arr[0].(string)
	if !ok {
		log.Fatal("expected string inside array")
	}

	// fmt.Println("Unmarshal EOD data in interface-------", s)

	// Step 3: unmarshal actual JSON string into struct
	var candle Candle1
	if err := json.Unmarshal([]byte(s), &candle); err != nil {
		log.Fatal("inner unmarshal error:", err)
	}

	return candle

}

// I think it has 1 min delay candle data in 1 min timeframe
func GetCandleData(time930am int64, optionStock string, interval int, exchange string) ([]Candle1, error) {

	url1 := "https://piconnect.flattrade.in/PiConnectTP/TPSeries"
	data1 := `{
    "uid": "FZ23969",    
    "exch": "%s", 
    "token": "%s", 
    "st": "%d", 
    "et": "%d", 
    "intrv": "%d"
}`

	now := time.Now()
	timeNow := now.Unix()

	//testing
	// time930am = 1768189939 - 1000
	// timeNow = 1768189939

	data := fmt.Sprintf(data1, exchange, optionStock, time930am, timeNow, interval)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data, token))

	// fmt.Println("input data---------------", string(payload))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		// panic(err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request in chart data:-----", resp)
	if err != nil {
		// panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("GETCandle Data -- API Response in chart data:", resp.Status, string(body))

	var candles []Candle1
	if err := json.Unmarshal([]byte(body), &candles); err != nil {
		// panic(err)
		return nil, err
	}

	return candles, err

}

func NiftyCrossCheck() {

	//getting last 1 or 2 , 1 minutes candle data
	Ntmp, _ := GetCandleData(GetCurrentTimestamp()-120, Nifty50, 1, NiftyExchange)
	NiftyCandle := GetCandleDataInFloat64(Ntmp)

	diff = NiftyCandle.Intc[0] - NiftyYesterdayCandle.Intc[0]

	NiftyTradingOpenFlag = true

	//if points are lesser than or greater than -50 then stopping the new trades

	if diff > 0 {
		if diff < 50 {
			NiftyTradingOpenFlag = false
		}
	} else {
		// means points less than -40, EX: -30 is greater than -40, -50 is smaller than -40
		if diff > -50 {
			NiftyTradingOpenFlag = false
		}
	}

	//if points are reversed atleast 75 either sides, then switching the nifty same strike CE to PE or vice versa
	n := len(optionStock)
	if diff > 0 {
		if diff > 75 {
			if optionStock[n-6] == 'P' {
				optionStock = ReversalOption(optionStock)
			}
		}
	} else {
		// means points less than -40, EX: -30 is greater than -40, -50 is smaller than -40
		if diff < -75 {
			if optionStock[n-6] == 'C' {
				optionStock = ReversalOption(optionStock)
			}
		}
	}

}

func ReversalORFallLogic() bool {

	now := time.Now()

	tmp, _ := GetCandleData(GetTimestamp(), optionStock, 30, OptionsExchange)
	OptionsCandle := GetCandleDataInFloat64(tmp)

	Ntmp, _ := GetCandleData(GetTimestamp(), Nifty50, 30, NiftyExchange)
	NiftyCandle := GetCandleDataInFloat64(Ntmp)

	//(c.VWAP[1] > c.Intc[1]) && (c.VWAP[0] < c.Intc[0])

	diff := 0.0
	if NiftyYesterdayCandle.Intc[0] > NiftyCandle.Intc[0] {
		diff = NiftyYesterdayCandle.Intc[0] - NiftyCandle.Intc[0]
	} else {
		diff = NiftyCandle.Intc[0] - NiftyYesterdayCandle.Intc[0]
	}

	fmt.Println("reversalOrFall logic updates----------", now, diff, NiftyCandle, NiftyCandle.Intc[0], PreviousDayHigh, PreviousDayLow, OptionsCandle, OptionsCandle.Volume[0], OptionsCandle.Intc[0])

	f.WriteString(fmt.Sprintf("reversalOrFall logic updates ------------------:%+v\n", []interface{}{now, optionStock, diff, NiftyCandle, NiftyCandle.Intc[0], PreviousDayHigh, PreviousDayLow, OptionsCandle, OptionsCandle.Volume[0], OptionsCandle.Intc[0]}))

	//need atleast 2 - 30 min candles to calculate the below logic
	if len(OptionsCandle.Volume) < 2 {
		return false
	}

	//reversal (or) fall conditions
	// 0 index is the latest last candle data here
	// if today went for CE options, nifty points lesser than 35 points
	// if today went for PE options, nifty points greater than -35 points
	//current volume greater than 2 times of previous 30 minutes candle and close lesser than previous candle low
	//string(optionStock[len(optionStock)-6]) == "C" - checking calls or puts in nifty options

	// if ((OptionsCandle.Volume[0] > OptionsCandle.Volume[1]*2) && OptionsCandle.Intc[0] < OptionsCandle.Intl[1]) ||
	// 	((string(optionStock[len(optionStock)-6]) == "C" && diff < 35) || NiftyCandle.Intc[0] < PreviousDayLow-20) || ((string(optionStock[len(optionStock)-6]) == "P" && diff < 35) || NiftyCandle.Intc[0] > PreviousDayHigh+20) {
	// 	fmt.Println("reversal Or Fall condtion passed ------------------", now, optionStock, diff)
	// 	fmt.Println(OptionsCandle.Volume[0], OptionsCandle.Intc[0])
	// 	fmt.Println(NiftyCandle.Intc[0], PreviousDayHigh, PreviousDayLow)

	// 	f.WriteString(fmt.Sprintf("reversal Or Fall condtion passed ------------------:%+v\n", []interface{}{now, optionStock, diff, OptionsCandle.Volume[0], OptionsCandle.Intc[0], NiftyCandle.Intc[0], PreviousDayHigh, PreviousDayLow}))
	// 	return true

	// }

	if OptionsCandle.Volume[0] > (OptionsCandle.Volume[1]*2) && OptionsCandle.Intc[0] < OptionsCandle.Intl[1] {
		fmt.Println("reversal Or Fall condtion passed 111111------------------", now, optionStock)
		f.WriteString(fmt.Sprintf("reversal Or Fall condtion passed 11111-----------------"))
		return true
	}

	if string(optionStock[len(optionStock)-6]) == "C" && (diff < 35 || NiftyCandle.Intc[0] < PreviousDayLow-20) {
		fmt.Println("reversal Or Fall condtion passed 222222------------------", now, optionStock)
		f.WriteString(fmt.Sprintf("reversal Or Fall condtion passed 22222-----------------"))
		return true
	}

	if string(optionStock[len(optionStock)-6]) == "P" && (diff < 35 || NiftyCandle.Intc[0] > PreviousDayHigh+20) {
		fmt.Println("reversal Or Fall condtion passed 33333------------------", now, optionStock)
		f.WriteString(fmt.Sprintf("reversal Or Fall condtion passed 33333-----------------"))
		return true
	}

	return false

	// smaValue := SMA(closes, 20)                                   //SMA 20 period simple logic
	// if smaValue > closes[0] && SMA(closes, 9) > SMA(closes, 20) { //crosscheck this conditons

	// }

	//similarly using open interest, volume and close - you can add multiple combination logics EX: SMA, EMA, bollinger bands, open interest decrease, volume spike and then place order from here

}

func IntradayLogic() {

	if !NiftyTradingOpenFlag {
		fmt.Println("Nifty points are low for the day----------------------", diff, "----", optionStock)
		return
	}

	t225 := time.Date(year, month, day, 14, 25, 0, 0, location)
	if time.Now().In(location).After(t225) {
		fmt.Println("Thats it, No new buy orders not allowed for the day----------------------")
		return
	}

	now := time.Now().In(location)

	fmt.Println("intraday function executed--------", now, "-------", diff, "----", optionStock)

	var intcValues []float64

	t930 := time.Date(year, month, day, 9, 30, 0, 0, location)
	t1015 := time.Date(year, month, day, 10, 15, 0, 0, location)

	// Between 9:30 AM and 10:15 AM RSI strategies
	if now.After(t930) && now.Before(t1015) {
		fmt.Println("Now is between 9:30 AM and 10:15 AM")

		tmp, _ := GetCandleData(GetTimestamp(), optionStock, 1, OptionsExchange)
		candleData := GetCandleDataInFloat64(tmp)
		intcValues = candleData.Intc

		asma5 := SMA(intcValues, 5)
		asma9 := SMA(intcValues, 9)

		diff := asma5 - asma9 //if difference less than 2 percent of last candle close value, may be difference is low that can be false crossover
		smaCrossoverDifference := (2.0 / 100.0) * intcValues[0]

		fmt.Println(asma5, asma9, asma5 > asma9, "------------", len(intcValues), intcValues)

		if asma5 > asma9 && diff > smaCrossoverDifference && intcValues[0] > candleData.VWAP[0] && intcValues[0] > asma5 {

			target := 108.0
			//placing limit price as asma5 because this is natural pullback on trending days
			PlaceOrder(optionStock, niftyOneLot, asma5, 8, target, GetTrailingStopLoss(target))
			msg := fmt.Sprintf("before 10.15am -- SMA 5 > SMA 9 crossover happened - %v -- %v -- %v", optionStock, now, intcValues[0])
			fmt.Println(msg)
			f.WriteString(msg)

			// SendTelegramMessage(msg)

			time.Sleep(5 * time.Minute)

			return

		}

		// time.Sleep(1 * time.Minute)

		return
	}

	tmp, _ := GetCandleData(GetTimestamp(), optionStock, 5, OptionsExchange)
	candleData := GetCandleDataInFloat64(tmp)
	intcValues = candleData.Intc

	//checking SMA crossover logic after 10.15am

	bsma5 := SMA(intcValues[1:], 5)
	bsma9 := SMA(intcValues[1:], 9)
	asma5 := SMA(intcValues, 5)
	asma9 := SMA(intcValues, 9)

	fmt.Println("before SMA --------", bsma5, bsma9, bsma5 < bsma9, "------------", len(intcValues[1:]), now)
	fmt.Println("latest SMA --------", asma5, asma9, asma5 > asma9, "------------", len(intcValues), now)

	if bsma5 < bsma9 && asma5 > asma9 && intcValues[0] > asma9 {

		// if asma5 > asma9 {

		target := 40.0

		PlaceOrder(optionStock, niftyOneLot, intcValues[0]-2.0, 8, target, GetTrailingStopLoss(target))
		msg := fmt.Sprintf("SMA 5 > SMA 9 crossover happened - %v -- %v -- %v", optionStock, now, intcValues[0])
		fmt.Println(msg)
		f.WriteString(msg)
		// SendTelegramMessage(msg)
		time.Sleep(5 * time.Minute)
		return

	}

	// time.Sleep(1 * time.Minute)

	return
}

func OptionsReversalCEPE(optionsStock string) string {
	return ""
}

func Truncate(val float64) float64 {
	return math.Trunc(val*100) / 100
}

func IntradayLogicVersion1() {

	now := time.Now().In(location)

	fmt.Println("intraday function Version1 executed--------", now)

	t1015 := time.Date(year, month, day, 10, 15, 0, 0, location)
	if now.Before(t1015) {
		fmt.Println("intraday function Version1 too early to execute-------", now)
		return
	}

	// t225 := time.Date(year, month, day, 14, 25, 0, 0, location)
	// if time.Now().In(location).After(t225) {
	// 	fmt.Println("Thats it, No new buy orders not allowed for the day----------------------")
	// 	return
	// }

	var intcValues []float64

	optionsstocks, err := StockOptionsFetchAndValidation()
	if err != nil {
		fmt.Println("error in fetching the options stocks", err)
		return
	}

	// t930 := time.Date(year, month, day, 9, 30, 0, 0, location)
	// t1015 := time.Date(year, month, day, 10, 15, 0, 0, location)

	for optionSymbol, optionStock := range optionsstocks {

		// // Between 9:30 AM and 10:15 AM RSI strategies
		// if now.After(t930) && now.Before(t1015) {
		// 	fmt.Println("Now is between 9:30 AM and 10:15 AM")

		// 	tmp, _ := GetCandleData(GetTimestamp(), optionStock, 1, OptionsExchange)
		// 	candleData := GetCandleDataInFloat64(tmp)
		// 	intcValues = candleData.Intc

		// 	asma5 := SMA(intcValues, 5)
		// 	asma9 := SMA(intcValues, 9)

		// 	diff := asma5 - asma9 //if difference less than 2 percent of last candle close value, may be difference is low that can be false crossover
		// 	smaCrossoverDifference := (2.0 / 100.0) * intcValues[0]

		// 	fmt.Println(asma5, asma9, asma5 > asma9, diff > smaCrossoverDifference, "------------", len(intcValues), intcValues)

		// 	if asma5 > asma9 && diff > smaCrossoverDifference {
		// 		var target int64
		// 		target = 108
		// 		//placing limit price as asma5 because this is natural pullback on trending days
		// 		PlaceOrder(optionStock, niftyOneLot, asma5, 8, target, GetTrailingStopLoss(target))
		// 		msg := fmt.Sprintf("before 10.15am -- SMA 5 > SMA 9 crossover happened - %v -- %v -- %v", optionStock, now, intcValues[0])
		// 		fmt.Println(msg)
		// 		f.WriteString(msg)

		// 		// SendTelegramMessage(msg)

		// 		time.Sleep(5 * time.Minute)

		// 		return

		// 	}

		// 	// time.Sleep(1 * time.Minute)

		// 	return
		// }

		//as of now options stocks working for sma crossover
		tmp, _ := GetCandleData(GetTimestamp(), optionStock, 5, OptionsExchange)
		candleData := GetCandleDataInFloat64(tmp)

		intcValues = candleData.Intc

		if candleData.OpenInterest[0] < OptionsStockOI || candleData.Volume[0] < OptionsStockVolume {
			fmt.Println("less OI or volume ------------", optionStock, candleData.OpenInterest[0], OptionsStockOI, candleData.Volume[0], OptionsStockVolume)
			continue
		}

		//if price and OI is increasinbg for last n candles, n - 2, can set to 3 as well
		// candleNumbers := 2
		// if !CheckOptionsPriceWithOI(intcValues, candleData.OpenInterest, candleNumbers) {
		// 	fmt.Println("price or OI is not increasing ------------", optionStock, intcValues[:candleNumbers], candleData.OpenInterest[:candleNumbers])
		// 	continue
		// }

		//checking SMA crossover logic after 10.15am

		bsma5 := SMA(intcValues[1:], 5)
		bsma9 := SMA(intcValues[1:], 9)
		asma5 := SMA(intcValues, 5)
		asma9 := SMA(intcValues, 9)

		fmt.Println("before SMA --------", bsma5, bsma9, bsma5 < bsma9, "------------", optionStock, len(intcValues[1:]), now)
		fmt.Println("latest SMA --------", asma5, asma9, asma5 > asma9, "------------", optionStock, len(intcValues), now)

		if bsma5 < bsma9 && asma5 > asma9 && intcValues[0] > asma9 && intcValues[0] > candleData.VWAP[0] && candleData.Volume[0] > SMA(candleData.Volume, 9) {

			var target, limitPrice, TrailingSL, stoploss float64

			limitPrice = RoundTo005(Truncate(intcValues[0] - ((OptionsStockLimitPricePercentage / 100.0) * intcValues[0]))) //limitprice 2% minus last candle close
			stoploss = RoundTo005(Truncate((OptionsStockStopLossPercentage / 100.0) * intcValues[0]))                       //stop loss 5% of last candle close
			target = RoundTo005(Truncate((OptionsStockTargetPercentage / 100.0) * intcValues[0]))                           //target 50% of last candle close

			TrailingSL = RoundTo005(Truncate(target / 2))

			if limitPrice < 0.15 {
				fmt.Println("very less entry limit price---------", optionStock, limitPrice)
				continue
			}

			//flattrade require minimum value 1
			if TrailingSL < 1 {
				TrailingSL = 1
			}

			// PlaceOrder(optionStock, lotSize[optionSymbol], limitPrice, stoploss, target, TrailingSL)
			fmt.Println("-----------------------------", now, optionStock, lotSize[optionSymbol], limitPrice, stoploss, target, TrailingSL)
			msg := fmt.Sprintf("Stock OPtions SMA 5 > SMA 9 crossover happened - %v -- %v -- %v", optionStock, now, intcValues[0])
			fmt.Println(msg)
			f.WriteString(msg)
			// SendTelegramMessage(msg)
			// time.Sleep(5 * time.Minute)
			return

		}
	}

	// time.Sleep(1 * time.Minute)

}

func PercentageOfValue() {

}

func RSILogicTrade() {
	var intcValues []float64
	// var target int64

	now := time.Now().In(location)

	t930 := time.Date(year, month, day, 9, 30, 0, 0, location)
	t1015 := time.Date(year, month, day, 10, 15, 0, 0, location)

	// Between 9:30 AM and 10:15 AM RSI strategies
	if now.After(t930) && now.Before(t1015) {
		fmt.Println("Now is between 9:30 AM and 10:15 AM")

		tmp, _ := GetCandleData(GetTimestamp(), optionStock, 1, OptionsExchange)
		candleData := GetCandleDataInFloat64(tmp)
		intcValues = candleData.Intc

		slices.Reverse(intcValues)

		rsi1 := calcRSI(intcValues, 3)

		fmt.Println("RSI value", rsi1, ":", now)

		f.WriteString(fmt.Sprintf("RSI1 value - %v   - %v\n", rsi1, now))

		if rsi1 <= 15.0 {
			// target = 108
			// PlaceOrder(optionStock, niftyOneLot, intcValues[len(intcValues)-1]-4.0, 8, target, int64(GetTrailingStopLoss(target)))

			fmt.Println("---RSI000 buy order executed--------so sleeping now------------------------------", rsi1, intcValues[len(intcValues)-1], now)

			f.WriteString(fmt.Sprintf("---RSI 000 buy order executed--------so sleeping now------------------------------:%v, %v, %v\n", rsi1, intcValues[len(intcValues)-1], now))

			time.Sleep(10 * time.Minute)

			return

		}

		time.Sleep(1 * time.Minute)

		return
	}

	tmp, _ := GetCandleData(GetTimestamp(), optionStock, 5, OptionsExchange)
	candleData := GetCandleDataInFloat64(tmp)
	intcValues = candleData.Intc

	//checking SMA crossover logic

	bsma5 := SMA(intcValues[:len(intcValues)-1], 5)
	bsma9 := SMA(intcValues[:len(intcValues)-1], 9)
	asma5 := SMA(intcValues, 5)
	asma9 := SMA(intcValues, 9)

	fmt.Println(bsma5, bsma9, bsma5 < bsma9, "------------", len(intcValues[:len(intcValues)-1]), intcValues[:len(intcValues)-1])
	fmt.Println(asma5, asma9, asma5 > asma9, "------------", len(intcValues), intcValues)

	if bsma5 < bsma9 && asma5 > asma9 {
		msg := fmt.Sprintf("SMA 5 > SMA 9 crossover happened - %v -- %v -- %v", optionStock, now, intcValues[len(intcValues)-1])
		fmt.Println(msg)
		f.WriteString(msg)
		SendTelegramMessage(msg)

	}

	slices.Reverse(intcValues) //APi candle close values returns from latest to oldest, but for calculating rsi closing values should be old to latest

	// fmt.Println("candle close values", len(intcValues), "--------------------------------", intcValues)

	rsi := calcRSI(intcValues, 3)

	fmt.Println("RSI value", rsi, ":", now)

	f.WriteString(fmt.Sprintf("RSI value - %v   - %v\n", rsi, now))

	//EX: RSI > 15 && RSI <= 22.5
	if rsi > rsiLimitlower && rsi <= rsiLimitlow {

		maxValue := slices.Max(intcValues)

		maxdiff := maxValue - intcValues[len(intcValues)-1] //last candle value close

		fmt.Println("RSI conditon111 passed, checking maxvalue---", maxValue, maxdiff, GetMaxDifference(maxValue, true), now)

		//need to check for avoid placing buy order in deep reversal or fall
		// minValue := slices.Min(intcValues)

		// mindiff := intcValues[len(intcValues)-1] - minValue //last candle value close

		//if the difference price is less the Ex: 20, more convinced entry for not buying at high price range depsite less RSI
		if maxdiff > GetMaxDifference(maxValue, true) {

			// target = 12

			//for 1 min timeframe, placing last candles close - 2 as limit price
			//for 5 min timeframe, placing last candles close - 10 as limit price

			// PlaceOrder(optionStock, niftyOneLot, intcValues[len(intcValues)-1]-7.0, 12, 60, 12)

			// PlaceOrder(optionStock, niftyOneLot, intcValues[len(intcValues)-1]-8.0, 10, target, (GetTrailingStopLoss(target)))

			fmt.Println("---RSI 111 buy order executed--------so sleeping now------------------------------", rsi, intcValues[len(intcValues)-1], maxValue, now)

			f.WriteString(fmt.Sprintf("---RSI 111 buy order executed--------so sleeping now------------------------------:%v, %v, %v\n", rsi, intcValues[len(intcValues)-1], now))

			time.Sleep(10 * time.Minute) //if order executed, goes to sleep for 10 minutes to avoid buying reversal or deep fall again and make more losses

			return

		}

		time.Sleep(2 * time.Minute)

		return

	}

	if rsi <= rsiLimitlower {

		maxValue := slices.Max(intcValues)

		maxdiff := maxValue - intcValues[len(intcValues)-1] //last candle value close

		fmt.Println("RSI conditon passed222, checking maxvalue---", maxValue, maxdiff, GetMaxDifference(maxValue, false), now)

		//need to check for avoid placing buy order in deep reversal or fall
		// minValue := slices.Min(intcValues)

		// mindiff := intcValues[len(intcValues)-1] - minValue //last candle value close

		//if the difference price is less the Ex: 20, more convinced entry for not buying at high price range depsite less RSI
		if maxdiff > GetMaxDifference(maxValue, false) {

			// target = 25

			//for 1 min timeframe, placing last candles close - 2 as limit price
			//for 5 min timeframe, placing last candles close - 10 as limit price

			// PlaceOrder(optionStock, niftyTwoLot, intcValues[len(intcValues)-1]-2.0, 7, 60, 7)

			// PlaceOrder(optionStock, niftyOneLot, intcValues[len(intcValues)-1]-2.0, 8, target, GetTrailingStopLoss(target))

			fmt.Println("---RSI 222 buy order executed--------so sleeping now------------------------------", rsi, intcValues[len(intcValues)-1], maxValue, now)

			f.WriteString(fmt.Sprintf("---RSI 222 buy order executed--------so sleeping now------------------------------:%v, %v, %v\n", rsi, intcValues[len(intcValues)-1], now))

			time.Sleep(10 * time.Minute) //if order executed, goes to sleep for 10 minutes to avoid buying reversal or deep fall again and make more losses

			return

		}

	}

	time.Sleep(2 * time.Minute)

	return
}

func GetCandleDataInFloat64(candles []Candle1) Candle {

	var vwaps []float64
	var closes []float64
	var lows []float64
	var highs []float64
	var volumes []float64
	var ois []float64

	candleValues := Candle{}

	for _, c := range candles {

		val, err := strconv.ParseFloat(c.Intc, 64)
		if err != nil {
			panic(err)
		}
		closes = append(closes, val)

		val, err = strconv.ParseFloat(c.Intl, 64)
		if err != nil {
			panic(err)
		}
		lows = append(lows, val)

		val, err = strconv.ParseFloat(c.Inth, 64)
		if err != nil {
			panic(err)
		}
		highs = append(highs, val)

		val, _ = strconv.ParseFloat(c.VWAP, 64)
		vwaps = append(vwaps, val)

		val, err = strconv.ParseFloat(c.Volume, 64)
		if err != nil {
			panic(err)
		}
		volumes = append(volumes, val)

		val, _ = strconv.ParseFloat(c.OpenInterest, 64)
		ois = append(ois, val)
	}

	candleValues.Intc = closes
	candleValues.Intl = lows
	candleValues.Inth = highs
	candleValues.VWAP = vwaps
	candleValues.Volume = volumes
	candleValues.OpenInterest = ois

	return candleValues

}

func PlaceOrder(optionStock string, lot string, buyPrice float64, stopLoss float64, target float64, trailingStopLoss float64) {
	url1 := "https://piconnect.flattrade.in/PiConnectTP/PlaceOrder"

	/*
			"exch": "NFO"  --> for options, in case of normal stocks buying: "exch": "NSE"
			"tsym": "NIFTY18NOV25C26000"   ---> Example NIFTY 18NOV25 CE 26000
			prd": "B" - bracket order
			"prc": "80", -- entry limit order price
		   bpprc - target price, EX: 20 price from entry price
		   blprc - stop loss price
		   trailprc - trailing stop loss
		   "ret": "DAY"  - order validity day

	*/

	//correct working nifty options BO order with limit price, target, stop loss and trailing stop loss
	data := `{"uid": "FZ23969",
	    "actid": "FZ23969",
	    "exch": "NFO",
	    "tsym": "%s",
	    "qty": "%s",
	    "prd": "B",
	    "trantype": "B",
	    "prctyp": "LMT",
	    "ret": "DAY",
		"prc": "%v",
		"blprc":"%v",
		"bpprc":"%v",
		"trailprc":"%v"
	}`

	data1 := fmt.Sprintf(data, optionStock, lot, buyPrice, stopLoss, target, trailingStopLoss)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data1, token))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request for order placed", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response for order placed-----------:", string(body))
	f.WriteString(fmt.Sprintf("API Response for order placed-----------:%v\n", string(body)))

}

func ModifyOrder(orderNo string, symbol string, lotSize string, modifyPrice float64) {
	url1 := "https://piconnect.flattrade.in/PiConnectTP/ModifyOrder"

	/*
			"exch": "NFO"  --> for options, in case of normal stocks buying: "exch": "NSE"
			"tsym": "NIFTY18NOV25C26000"   ---> Example NIFTY 18NOV25 CE 26000
			prd": "B" - bracket order
			"prc": "80", -- entry limit order price
		   bpprc - target price, EX: 20 price from entry price
		   blprc - stop loss price
		   trailprc - trailing stop loss
		   "ret": "DAY"  - order validity day

	*/

	data := `{"uid": "FZ23969",
	    "actid": "FZ23969",
	    "exch": "NFO",
	    "tsym": "%s",
	    "qty": "%s",
		"prc": "%v",
		"prctyp": "LMT",
        "ret": "DAY", 
		"bpprc":"%v",
		"norenordno": "%s"
	}`

	//"blprc":"%v",
	// stoploss := modifyPrice + 0.05
	target := modifyPrice - 0.05
	// data1 := fmt.Sprintf(data, symbol, lotSize, modifyPrice, stoploss, target, orderNo)

	data1 := fmt.Sprintf(data, symbol, lotSize, modifyPrice, target, orderNo)

	payload := []byte(fmt.Sprintf("jData=%s&jKey=%s", data1, token))

	fmt.Println("payload-----------:", string(payload))

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	// fmt.Println("API Request for order placed", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response for order modified-----------:", string(body))
	f.WriteString(fmt.Sprintf("API Response for order modified-----------:%v\n", string(body)))

}

type OrderDetails struct {
	OrderNo       string `json:"norenordno"`
	Exchange      string `json:"exch"`
	OrderSymbol   string `json:"tsym"`
	Status        string `json:"status"`
	BuyOrSell     string `json:"trantype"`
	OrderTime     string `json:"exch_tm"`
	OrderQuantity string `json:"qty"`
	OrderToken    string `json:"token"`
}

type Candle1 struct {
	Intc         string `json:"intc"` // closing price string
	Inth         string `json:"inth"`
	Intl         string `json:"intl"`
	VWAP         string `json:"intvwap"`
	Volume       string `json:"intv"`
	OpenInterest string `json:"oi"`
}

type Candle struct {
	Intc         []float64 `json:"intc"` // closing price string
	Inth         []float64 `json:"inth"`
	Intl         []float64 `json:"intl"`
	VWAP         []float64 `json:"intvwap"`
	Volume       []float64 `json:"intv"`
	OpenInterest []float64 `json:"oi"`
}

func HasTimePassed(timeStr string, timeLimit time.Duration) bool {
	layout := "02-01-2006 15:04:05"
	inputTime, _ := time.ParseInLocation(layout, timeStr, location)
	now := time.Now().In(location)

	diff := now.Sub(inputTime)

	return diff > timeLimit
}

// type Candle struct {
// 	Intc         []float64 // closing price string
// 	Inth         []float64
// 	Intl         []float64
// 	VWAP         []float64
// 	Volume       []float64
// 	OpenInterest []float64
// }

// for calculating rsi(3), (n+1) --> last 4 candles close values should be send
func calcRSI(prices []float64, period int) float64 {

	var gainSum, lossSum float64
	for i := 1; i <= period; i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			gainSum += diff
		} else {
			lossSum += -diff
		}
	}

	avgGain := gainSum / float64(period)
	avgLoss := lossSum / float64(period)

	// Step 2: Wilder smoothing for last value
	for i := period + 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]

		var gain, loss float64
		if diff > 0 {
			gain = diff
			loss = 0
		} else {
			gain = 0
			loss = -diff
		}

		avgGain = ((avgGain * float64(period-1)) + gain) / float64(period)
		avgLoss = ((avgLoss * float64(period-1)) + loss) / float64(period)
	}

	// Step 3: Final RSI calculation
	if avgLoss == 0 {
		return 100
	}

	rs := avgGain / avgLoss
	rsi := 100 - (100 / (1 + rs))

	return rsi
}

// pass the day in number and hour in number other today 9.30 value is default
func GetTimestamp(timeValues ...int) int64 {

	if len(timeValues) == 1 {
		panic("Give no values or both day and hour values---")
	}

	// Current date in IST
	now := time.Now().In(location)

	hour := 9
	day := now.Day()

	if len(timeValues) == 2 {
		day = timeValues[0]
		hour = timeValues[1]

	}

	// Construct today's 09:30 AM in IST
	today930 := time.Date(
		now.Year(),
		now.Month(),
		day,
		hour, 20, 0, 0, //8th jan 2026, changes now 9.20am
		location,
	)

	// fmt.Println("Today's 9:30 AM IST timestamp:", today930.Unix())
	// fmt.Println("Human-readable:", today930)

	return today930.Unix()
}

func GetMaxDifference(num float64, lower bool) float64 {
	val := 0.0
	if num >= 200 {
		val = 40.0
	} else if num >= 150 && num < 200 {
		val = 30.0
	} else if num >= 100 && num < 150 {
		val = 20.0
	} else if num < 100 {
		val = 15.0
	}

	if lower {
		// return val - 10.0
	}

	return val
}

func GetTrailingStopLoss(target float64) float64 {
	if target <= 20 {
		return 0
	} else if target > 20 && target <= 60 {
		return target / 2
	} else {
		return target / 3
	}
}

func SMA(closes []float64, period int) float64 {
	if len(closes) < period || period <= 0 {
		return 0
	}
	sum := 0.0
	for _, c := range closes[:period] {
		sum += c
	}
	val := sum / float64(period)
	return RoundTo005(val)
}

func SendTelegramMessage(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Telegram Response:", string(body))

	}
	return nil
}

func GetOptionsLotSize(optionsSymbol string) string {
	var SymbolLotSizeFeb2026 = map[string]string{}

	return SymbolLotSizeFeb2026["dummy"]

}

// func CheckOptionsOIVolume(volume []float64, OI []float64) {
// 	if volume >
// }

func CheckOptionsPriceWithOI(closes []float64, OI []float64, length int) bool {
	if len(closes) < length || len(OI) < length {
		return false
	}

	for i := 0; i < length-1; i++ {
		if closes[i] < closes[i+1] || OI[i] < OI[i+1] {
			return false
		}
	}

	return true

	// if closes[0] > closes[1] && OI[0] > OI[1] {
	// 	return true
	// }

	// return false
}

func LoadLotSizeMap(csvPath string) (map[string]string, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	// skip header
	for i := 1; i < len(records); i++ {
		symbol := strings.TrimSpace(records[i][0])
		lotSize := strings.TrimSpace(records[i][1])
		m[symbol] = lotSize
	}

	return m, nil
}

// fetch top gainers ang losers stocks on options
func NseGet() ([]byte, error) {
	var defaultHeaders = map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
		"Accept":          "application/json,text/plain,*/*",
		"Accept-Language": "en-US,en;q=0.9",
		"Connection":      "keep-alive",
		"Referer":         "https://www.nseindia.com/",
	}

	url := "https://www.nseindia.com/api/equity-stockIndices?index=SECURITIES%20IN%20F%26O"

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("nse fetching options stock--", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func FilterFOStocks(body []byte) (map[string]string, map[string]string, error) {
	var Raw struct {
		Data []FoStock `json:"data"`
	}

	CEStocks := map[string]string{}
	PEStocks := map[string]string{}

	err := json.Unmarshal(body, &Raw)
	if err != err {
		return nil, nil, err
	}

	// fmt.Println("data--------------", Raw)

	for _, stock := range Raw.Data {
		if stock.PercentChange > 1.5 && stock.PercentChange < 3.0 {
			CEStocks[stock.Symbol] = stock.LastPrice
		}

		if stock.PercentChange < -1.5 && stock.PercentChange > -3.0 {
			PEStocks[stock.Symbol] = stock.LastPrice
		}
	}

	return CEStocks, PEStocks, nil
}

func RoundTo005(val float64) float64 {

	scaled := math.Round(val * 20) // because 1 / 0.05 = 20
	return scaled / 20
}

type FoStock struct {
	Symbol        string  `json:"symbol"`
	LastPrice     string  `json:"lastPrice"`
	PercentChange float64 `json:"pChange"`
	Volume        int64   `json:"totalTradedVolume"`
}
