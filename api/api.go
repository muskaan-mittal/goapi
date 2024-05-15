package api

import  (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct{
	Username string
} //parameters that API endpoint will take

type CoinBalanceResponse struct{
	Code int //Status code (eg: 200)
	Balance int //Account balance
}

type Error struct{
	Code int //Error code
	Message string
}


