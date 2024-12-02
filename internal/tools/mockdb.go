package tools

import (
	"time"
)

type mockDB struct{}

var mockUsers = map[string]Users{
	"alex": {
		Username: "alex",
		Password: "123ABC",
		Email:    "alex@gmail.com",
	},
	"jason": {
		Username: "jason",
		Password: "456DEF",
		Email:    "jason@gmail.com",
	},
	"marie": {
		Username: "marie",
		Password: "789GHI",
		Email:    "marie@gmail.com",
	},
}

var mockLogins = map[string]Logins{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockAlerts = map[string][]Alerts{
	"alex": {
		{
			IdAlert:   "abc",
			Username:  "alex",
			Price:     100.01,
			Condition: "over",
			StartDate: "30-11-24",
			EndDate:   "07-12-24",
			IsActive:  "true",
		},
		{
			IdAlert:   "hola",
			Username:  "alex",
			Price:     100.01,
			Condition: "over",
			StartDate: "30-11-24",
			EndDate:   "07-12-24",
			IsActive:  "true",
		},
	},
	"jason": {
		{
			IdAlert:   "def",
			Username:  "jason",
			Price:     100.01,
			Condition: "over",
			StartDate: "30-11-24",
			EndDate:   "07-12-24",
			IsActive:  "true",
		},
	},
	"marie": {
		{
			IdAlert:   "ghi",
			Username:  "marie",
			Price:     100.01,
			Condition: "over",
			StartDate: "30-11-24",
			EndDate:   "07-12-24",
			IsActive:  "true",
		},
	},
}

func (d *mockDB) GetUser(username string, password string) *Users {
	time.Sleep(time.Second * 1)

	var clientData = Users{}
	clientData, ok := mockUsers[username]
	if ok {
		ok = (clientData.Password == password)
	}
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserLoginDetails(username string) *Logins {
	time.Sleep(time.Second * 1)

	var clientData = Logins{}
	clientData, ok := mockLogins[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserAlerts(username string) *[]Alerts {
	time.Sleep(time.Second * 1)

	var clientData = []Alerts{}
	clientData, ok := mockAlerts[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
