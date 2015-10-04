package main

/*
 * RPC - process all methods callback by rpc service site
 */

import (
	"net/http"
	"fmt"
	"strconv"
	"strings"
	"errors"
)



type InputArgs struct {
	Budget      float64
	Input       string
	TradeId     int
}

type InputStockInfo struct {
	Symbol      string
	Percentage  float64
}

type ReplyMessage struct {
	Message     string
}

type Transaction struct {
	Tradeid     int
	Stocks      []OneStock
	Budget      float64
	Unvested    float64
}

type OneStock struct {
	Name        string
	Price       float64
	Amount      int
	percentage  float64
}

var transactions []Transaction;


type Service struct{}

func (h *Service) BuyRequest(r *http.Request, args *InputArgs, reply *ReplyMessage) (error) {

	if(args.Budget<=0) {
		e := errors.New("Input Error: invalid Budget")
		reply.Message = e.Error()
		return e
	}

	stockinfo,err := ParseInput(args.Input)
	if err != nil {
		reply.Message = err.Error()
		fmt.Printf(reply.Message)
		return err
	}

	var tran Transaction

	tran.Budget = args.Budget

	for _,v := range stockinfo {
		s,err := QueryInfo(v.Symbol)
		if err != nil {
			return err
		}
		s.percentage = v.Percentage
		tran.Stocks = append(tran.Stocks,s)
	}

	tran.Tradeid = len(transactions)+1

	tran.BuyStock();

	transactions = append(transactions,tran)

	reply.Message = tran.BuyResponse()

	return nil
}


func (tran *Transaction)BuyResponse() string {

	str:= "tradeId: "+ strconv.Itoa(tran.Tradeid)
	str += "\nstocks: "

	for _,v := range tran.Stocks {
		str += v.Name+":"+strconv.Itoa(v.Amount)+":$"+strconv.FormatFloat(v.Price, 'f', 3, 64) + ", "
	}
	str += "\nunvestedAmount: " + strconv.FormatFloat(tran.Unvested, 'f', 3, 64)

	return str
}

func (tran *Transaction)BuyStock(){

	tran.Unvested = tran.Budget

	for i,v := range tran.Stocks {

		n := (v.percentage/100)*tran.Budget
		if tran.Unvested>=n {
			tran.Stocks[i].Amount = int(n/v.Price)
		} else {
			tran.Stocks[i].Amount = int(tran.Unvested/v.Price)
		}
		tran.Unvested = tran.Unvested - float64(tran.Stocks[i].Amount)*v.Price

	}
}

func ParseInput(input string) ([]InputStockInfo,error) {

	str := strings.Split(input, ",")

	var stocks []InputStockInfo
	var temp InputStockInfo

	for _,oneStock := range str {

		s := strings.Split(oneStock,":")
		//s[0] is the symbol of stock , s[1] is the precentage of stock
		if len(s) == 1 {
			err := errors.New("Format Error: stock Symbol And Percentage")
			return stocks,err
		}
		s[1] = strings.TrimSuffix(s[1], "%")
		percentage, err := strconv.ParseFloat(s[1], 64);
		if err != nil {
			err := errors.New("Format Error: Percentage")
			return stocks,err
		}

		temp.Percentage = percentage
		temp.Symbol = s[0]
		stocks = append(stocks,temp)
	}

	return stocks,nil;
}

func (h *Service) QueryRequest(r *http.Request, args *InputArgs, reply *ReplyMessage) error {

	str :="stocks: "

	if (args.TradeId > len(transactions)) || (args.TradeId <= 0) {

		err := errors.New("Input Error: invalid TradeId")
		reply.Message = "Input Error: invalid TradeId"
		return err;
	}

	tran := transactions[args.TradeId-1]
	var currentValue float64
	for _,v := range tran.Stocks {

		currentStock,err := QueryInfo(v.Name)
		if err != nil {
			fmt.Println(err)
			return err
		}

		str += v.Name+":"+strconv.Itoa(v.Amount)+":"

		if currentStock.Price > v.Price {
			str += "+"
		}
		if currentStock.Price == v.Price {
			str += "="
		}
		if currentStock.Price < v.Price {
			str += "-"
		}
		str += "$"+strconv.FormatFloat(v.Price, 'f', 3, 64) + ", "
		currentValue += float64(v.Amount)*currentStock.Price
	}

	str += "\ncurrentMarketValue: "+ strconv.FormatFloat(currentValue, 'f', 3, 64)
	str += "\nunvestedAmount: " + strconv.FormatFloat(tran.Unvested, 'f', 3, 64)

	reply.Message = str

	//fmt.Println(reply.Message)

	return nil
}

