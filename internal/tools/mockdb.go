package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
}

var mockBonusDetails = map[string]BonusDetails {
	"alex": {
		Bonuses: 100,
		Username: "alex",
	},
	"jason": {
		Bonuses: 200,
		Username: "jason",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserBonuses(username string) *BonusDetails {
	time.Sleep(time.Second * 1)

	var clientData = BonusDetails{}
	clientData, ok := mockBonusDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}