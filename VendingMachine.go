package VendingMachine

import (
	"GoVending/Coins"
	"errors"
	"strconv"
	//"fmt"
)

/*
   Machine struct and bound functions
*/

// Machine is a class to represent the vending machine
/* Monitors total input by the customer (cents),
   the number of each coin inserted,
   important messages/prompts,
   products w/ prices (cents),
   and a string to display on screen */
type Machine struct {
	RunningTotal int
	InputCoins   map[Coins.Coin]int
	Messages     map[string]string
	Products     map[string]int
    Stock        map[string]int
	Display      string
}

// NewMachine is a constructor for a new machine
/* It records the total value of input coins,
   the number of each coin type input,
   available display messages, 
   and products (prices and stock count) */
func NewMachine() *Machine {
	m := new(Machine)
	m.RunningTotal = 0
	m.InputCoins = map[Coins.Coin]int{
		Coins.NewPenny():   0,
		Coins.NewNickel():  0,
		Coins.NewDime():    0,
		Coins.NewQuarter(): 0}
	m.Messages = map[string]string{
		"invalid": "INVALID COIN",
		"coin na": "COIN NOT AVAILABLE",
		"insert":  "INSERT COIN",
		"thanks":  "THANK YOU",
        "sold out": "SOLD OUT",
	}
	m.Products = map[string]int{
		"cola":  100,
		"chips": 50,
		"candy": 65,
	}
    m.Stock = map[string]int{
		"cola":  10,
		"chips": 5,
		"candy": 20,
	}
	m.Display = m.Messages["insert"]
	return m
}

// ToString will return the current status on the display
func (m *Machine) ToString() string {
	return m.Display
}

// AcceptCoins will add the input coin to the running total
func (m *Machine) AcceptCoins(inputCoin Coins.Coin) error {
	var err error
	c := IdentifyCoin(inputCoin)
	if IsValidCoinValue(c) {
		m.RunningTotal += c
		m.InputCoins[inputCoin]++
		dollars := float64(m.RunningTotal) / 100.0
		m.Display = "$" + strconv.FormatFloat(dollars, 'f', 2, 32)
	} else {
		m.Display = m.Messages["invalid"]
		err = errors.New(m.Messages["invalid"])
	}
	return err
}

// ReturnCoin will return a customer's coin
func (m *Machine) ReturnCoin(heldCoin Coins.Coin) (Coins.Coin, error) {
	var err error
	if m.InputCoins[heldCoin] != 0 {
		// return one of these
		m.InputCoins[heldCoin]--
		m.RunningTotal -= heldCoin.Value
		return heldCoin, err
	}
	m.Display = m.Messages["coin na"]
	err = errors.New(m.Messages["coin na"])
	return heldCoin, err
}

// ReturnAllCoins will return all coins that the customer has input
func (m *Machine) ReturnAllCoins() []Coins.Coin {
	var outCoins []Coins.Coin
	for typeOfCoin, numberOfCoins := range m.InputCoins {
		for number := 0; number < numberOfCoins; number++ {
			coin, _ := m.ReturnCoin(typeOfCoin)
			outCoins = append(outCoins, coin)
		}
	}
	m.Display = m.Messages["insert"]
	return outCoins
}

// DispenseChange will make change based on the current, remaining RunningTotal
func (m *Machine) DispenseChange() map[Coins.Coin]int {
	var quarter = Coins.NewQuarter()
	var dime = Coins.NewDime()
	var nickel = Coins.NewNickel()
	var penny = Coins.NewPenny()

	change := make(map[Coins.Coin]int)
	var count int

	for _, coin := range []Coins.Coin{quarter, dime, nickel, penny} {
		count = m.RunningTotal / coin.Value
		m.RunningTotal = m.RunningTotal % coin.Value
		change[coin] = count
	}
	return change
}

// ShowSelections will print the available items and thier prices
func (m *Machine) ShowSelections() {
	selectionString := ""
	for product, price := range m.Products {
		dollars := float64(price) / 100.0
		selectionString += product + " $" + strconv.FormatFloat(dollars, 'f', 2, 32) + "\n"
	}
	m.Display = selectionString
}

// SelectProduct will allow a customer to select their purchase
func (m *Machine) SelectProduct(product string) {
	price := m.Products[product]
	if m.RunningTotal >= price && m.Stock[product]>0 {
		// dispense product
		m.Display = m.Messages["thanks"]
		m.RunningTotal -= price
	} else if m.RunningTotal < price {
		// prompt for coins 
		m.Display = m.Messages["insert"]
    } else if m.Stock[product]==0 {
        // sold out 
        m.Display = m.Messages["sold out"]
    }
}

/*
   Utility Functions
*/

// IsValidCoinValue will check if the input is of value 1, 5, 10, or 25
func IsValidCoinValue(c int) bool {
	validCoins := []int{5, 10, 25}
	for _, vc := range validCoins {
		if c == vc {
			return true
		}
	}
	return false
}

// IdentifyCoin will identify the coin and return the numerical value
func IdentifyCoin(coin Coins.Coin) int {
	for _, coinProperties := range Coins.CoinTypes {
		if coin.Thickness     == coinProperties["thickness"] && 
                coin.Weight   == coinProperties["weight"]    && 
                coin.Diameter == coinProperties["diameter"]  {
			return int(coinProperties["value"])
		}
	}
	return 0
}
