# assignment1
1. The client main() function: client.go(/src/client/client.go)
2. The server main() function: server.go(/src/server/server.go)

3. When client is running, the following menu is dsiplayed on the shell windows:

    Please input the number of function that you want to choose:
    1. Buying stocks
    2. Checking your portfolio
    3. Exit
    
4. If you want to buy stocks, please choose "1". Three parameters(budget, symbol and percentage) are needed in this operation:

    Buying stocks:
    Please enter the budget:
    10000
    stockSymbolAndPercentage(E.g. “GOOG:50%,YHOO:50%”):
    goog:20%,yhoo:50%
    
    ***************************************
    tradeId: 1
    stocks: goog:3:$626.910, yhoo:162:$30.710,
    unvestedAmount: 3144.250
    ***************************************
    
5. If you need to query transactions that you have already bought, please choose "2" and input the trade id to specify which trade you are going to query.
    Checking portfolio:
    Please enter the tradeIde:1
    
    ***************************************
    stocks: goog:3:=$626.910, yhoo:162:=$30.710,
    currentMarketValue: 6855.750
    unvestedAmount: 3144.250
    ***************************************
*The meaning of the signs which are included in the result of query: 
"=": the current price equals the price that you bought.
"+": the current price is greater than the price that you bought.
"-": the current price is less than the price that you bought

6. The program will exit if "3" is chosen.

7. Error handle
  a. If the server is down and you want to check the protfolios or buy stocks, the client will show information of Error. For example,
      Http.Post Error
      Post http://127.0.0.1:8080/rpc: dial tcp 127.0.0.1:8080: getsockopt: connection refused
  
  b. If you input the wrong number (e.g. 6) at the main menu, the system will notice the wrong input:
      Please input the number of function that you want to choose:
      1. Buying stocks
      2. Checking your portfolio
      3. Exit
      6
      Input Error: please enter 1, 2, 3
      
  c. If you input the budget using characters, the shell will display the information to notice this error. For example:
      Please enter the budget:
      sfsdfasdf
      Format Error: Do not input the character
      
  d. If you input the incorrect format value of symbol and percentage during the operation of "1. Buying stocks" , the system will display as below:
      stockSymbolAndPercentage(E.g. “GOOG:50%,YHOO:50%”):
      gg50%.yhoo100%
      Format Error: stock Symbol And 
    *The correct format of stockSymbolAndPercentage is "StockName"+":"+"number"+"%"
    
  e. If you input the stock name that doesn't exsit during the operation of "1. Buying stocks", the system wil display the following notice:
      Buying stocks:
      Please enter the budget:
      100
      stockSymbolAndPercentage(E.g. “GOOG:50%,YHOO:50%”):
      goo:50%
      Yahoo API Error: Cannot find the stock: goo
      
  f. If you input the wrong tradeId during the operation of "2. Checking your portfolio", the system will display the message as below:
      Checking portfolio:
      Please enter the tradeIde:1
      Input Error: invalid TradeId
