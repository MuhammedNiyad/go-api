package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		Username: "alex",
		AuthToken: "1234",
	},
	"bob":{
		Username: "bob",
		AuthToken: "5678",
	},
	"charlie":{
		Username: "charlie",
		AuthToken: "9012",
	},

}

var mockCoinDetails = map[string]CoinDetails{
	"alex":{
		Username: "alex",
		Coin: 100,
	},
	"bob":{
		Username: "bob",
		Coin: 200,
	},
	"charlie":{
		Username: "charlie",
		Coin: 300,
	},

}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}