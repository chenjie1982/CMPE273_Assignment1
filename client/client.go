package main

/*
 * Client - implement Command line and send RPC to server
 */


import (
	"fmt"
	"github.com/gorilla/rpc/json"
	"net/http"
	"bytes"
	"bufio"
	"os"
)


type InputArgs struct {
	Budget  float64
	Input   string
	TradeId int
}

type ReplyMessage struct {
	Message string
}

func main() {

    var rec InputArgs
	var n int
	stdin := bufio.NewReader(os.Stdin)
	n = 0

	for n!=3 {

		switch n {
			case 1:{
				fmt.Println("Buying stocks:")
				fmt.Println("Please enter the budget:" )
				_,err := fmt.Fscanln(stdin,&rec.Budget)
				if err != nil {
					fmt.Println("Format Error: Do not input the character");
					stdin.ReadString('\n')
					break;
				}
				fmt.Println("stockSymbolAndPercentage(E.g. “GOOG:50%,YHOO:50%”):")
				_,err =fmt.Fscanln(stdin,&rec.Input)
				if err != nil {
					fmt.Println("Format Error: Do not input the space");
					stdin.ReadString('\n')
					break;
				}
				mes,err:= RPC("Service.BuyRequest",rec)
				if err != nil {
					fmt.Println(err);
					break
				}
				fmt.Println("\n***************************************")
				fmt.Println(mes.Message)
				fmt.Println("***************************************")
				break;
			}
			case 2:{
				fmt.Println("Checking portfolio:")
				fmt.Print("Please enter the tradeIde:")
				_,err := fmt.Fscanln(stdin,&rec.TradeId)
				if err != nil {
					fmt.Println("Format Error: Do not input the character");
					stdin.ReadString('\n')
					break;
				}
				mes,err:= RPC("Service.QueryRequest", rec)
				if err != nil {
					fmt.Println(err);
					break
				}
				fmt.Println("\n***************************************")
				fmt.Println(mes.Message)
				fmt.Println("***************************************")
				break;
			}
			default:{

			}
		}

		fmt.Println("\nPlease input the number of function that you want to choose:" )
		fmt.Println("1. Buying stocks ")
		fmt.Println("2. Checking your portfolio ")
		fmt.Println("3. Exit ")
		n=0;
		rec.Budget = 0;
		rec.Input=""
		rec.TradeId = 0;
		_,err := fmt.Fscanln(stdin,&n)
		if (err != nil) || (n>3) || (n<1) {
			fmt.Println("Input Error: please enter 1, 2, 3");
			stdin.ReadString('\n')
			n = 0;
		}
	}

}


func RPC(method string, args InputArgs) (reply ReplyMessage, err error) {

	buf, err := json.EncodeClientRequest(method,args)
	if err != nil {
		return reply, err
	}
	resp, err := http.Post("http://127.0.0.1:8080/rpc", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Println("Http.Post Error")
		return reply, err
	}

	defer resp.Body.Close()

	err = json.DecodeClientResponse(resp.Body, &reply)

	return reply, err
}


