package main

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"errors"
)

// Json struct for Yahoo finance API

type StockInfo struct {
    Query struct {
	    Count   int    `json:"count"`
	    Results struct {
	        Quote struct {
		        LastTradePriceOnly    string `json:"LastTradePriceOnly"`
				Symbol                string `json:"symbol"`
	        } `json:"quote"`
		} `json:"results"`
	} `json:"query"`
}


func QueryInfo(Symbol string) (OneStock, error){

	//create Yahoo YQL to query database
	queryStr := "select symbol, LastTradePriceOnly from yahoo.finance.quote where symbol in ('"+Symbol+"')"
	urlPath :=  "http://query.yahooapis.com/v1/public/yql?q="
	//convert space in url to UTF8
	urlPath += url.QueryEscape(queryStr)
	urlPath += "&format=json&env=store://datatables.org/alltableswithkeys"

	//fmt.Println(urlPath)
	var s StockInfo
	var stock OneStock

	//sent the http request to yahoo server
	res, err := http.Get(urlPath)
	if err!=nil {
		fmt.Println("QueryInfo: http.Get",err)
		return stock,err
	}
	defer res.Body.Close()

	body,err := ioutil.ReadAll(res.Body)
	if err!=nil {
		fmt.Println("QueryInfo: ioutil.ReadAll",err)
		return stock,err
	}

	//parser the json that yahoo server responses, and store in the q

	err = json.Unmarshal(body, &s)
	if err!=nil {
		fmt.Println("QueryInfo: json.Unmarshal",err)
		return stock,err
	}

	stock.Name = s.Query.Results.Quote.Symbol

	//cannot find the stock input from command line
	if s.Query.Results.Quote.LastTradePriceOnly == "" {

		err = errors.New("Yahoo API Error: Cannot find the stock: "+stock.Name)

		return stock,err
	}

	stock.Price,_ = strconv.ParseFloat(s.Query.Results.Quote.LastTradePriceOnly,64)


	return stock,nil
}
