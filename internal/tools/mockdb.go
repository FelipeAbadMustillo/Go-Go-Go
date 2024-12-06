package tools

import (
	"errors"
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

var mockAlerts = map[string][]Alerts{
	"alex": {
		{
			ID:        1,
			UserID:    "alex",
			Price:     100.01,
			Condition: "over",
			StartDate: time.Now().AddDate(0, 0, -7),
			Status:    "pending",
			CoinID:    "bitcoin",
			CoinName:  "BitCoin",
		},
		{
			ID:        2,
			UserID:    "alex",
			Price:     101.01,
			Condition: "under",
			StartDate: time.Now().AddDate(0, 0, -2),
			Status:    "closed",
			CoinID:    "bitcoin",
			CoinName:  "BitCoin",
		},
	},
	"jason": {
		{
			ID:        3,
			UserID:    "jason",
			Price:     102.01,
			Condition: "over",
			StartDate: time.Now().AddDate(0, 0, -1),
			Status:    "pending",
			CoinID:    "bitcoin",
			CoinName:  "BitCoin",
		},
	},
	"marie": {
		{
			ID:        4,
			UserID:    "marie",
			Price:     103.01,
			Condition: "over",
			StartDate: time.Now(),
			Status:    "pending",
			CoinID:    "bitcoin",
			CoinName:  "BitCoin",
		},
	},
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

func (d *mockDB) CreateUser(username string, email string, password string) (*Users, error) {
	time.Sleep(time.Second * 1)

	_, ok := mockUsers[username]
	if ok {
		return nil, errors.New("Username already picked.")
	}

	newUser := Users{
		Username: username,
		Password: password,
		Email:    email,
	}
	mockUsers[username] = newUser

	return &newUser, nil
}

func (d *mockDB) GetUser(username string, password string) (*Users, error) {
	time.Sleep(time.Second * 1)

	var clientData = Users{}
	clientData, ok := mockUsers[username]
	if ok {
		ok = (clientData.Password == password)
	}
	if !ok {
		return nil, errors.New("invalid username or password")
	}

	return &clientData, nil
}

func (d *mockDB) CreateAlert(newAlert *Alerts) error {
	time.Sleep(time.Second * 1)

	//Aca van validaciones
	mockAlerts[newAlert.UserID] = append(mockAlerts[newAlert.UserID], *newAlert)

	return nil
}

func (d *mockDB) GetUserAlerts(username string) *[]Alerts {
	time.Sleep(time.Second * 1)

	var clientData = []Alerts{}
	clientData, ok := mockAlerts[username]
	if !ok {
		return &[]Alerts{}
	}

	return &clientData
}
