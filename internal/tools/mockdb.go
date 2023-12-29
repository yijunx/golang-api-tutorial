package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"tom": {
		AuthToken: "tomtom",
		Username:  "tom",
	},
	"bar": {
		AuthToken: "barbar",
		Username:  "bar",
	},
	"foo": {
		AuthToken: "foofoo",
		Username:  "foo",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"tom": {
		Coins:    100,
		Username: "tom",
	},
	"bar": {
		Coins:    101,
		Username: "bar",
	},
	"foo": {
		Coins:    102,
		Username: "foo",
	},
}

// this is a class function of the mockDB!
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate db call which takes a second
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate db call which takes a second
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
