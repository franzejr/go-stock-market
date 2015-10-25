# Stock Market

Go simple app using [Markit on Deman API V2](http://dev.markitondemand.com/MODApis/).


## Build

```shell
go build main.go
```

## Usage

```shell
./main <STOCK_MARKET_SYMBOL>
```

## Example

Example:
```shell
./main GOOGL
"Status":"SUCCESS"
"Name":"Alphabet Inc"
"Symbol":"GOOGL"
"LastPrice":717.94
"Change":36.8000000000001
"ChangePercent":5.40270722612093
"Timestamp":"Fri Oct 23 15:59:00 UTC-04:00 2015"
"MSDate":42300.6659722222
"MarketCap":243936627620
"Volume":301775
"ChangeYTD":530.66
"ChangePercentYTD":35.2919006520183
"High":751.99
"Low":717.94
"Open":749.07
```
