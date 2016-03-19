package VendingMachine

import (
	"GoVending/Coins"
	"errors"
	"strconv"
    "fmt"
)

/*
   Machine struct and bound functions
*/

// Machine is a class to represent the vending machine
type Machine struct {
	RunningTotal int
	InputCoins   map[Coins.Coin]int
	messages     map[string]string
}

// NewMachine is a constructor for a new machine
func NewMachine() *Machine {
	m := new(Machine)
	m.RunningTotal = 0
	m.InputCoins = map[Coins.Coin]int{
		Coins.NewCoin("penny"):   0,
		Coins.NewCoin("nickel"):  0,
		Coins.NewCoin("dime"):    0,
		Coins.NewCoin("quarter"): 0}
	m.messages = map[string]string{
		"invalid": "INVALID COIN",
		"coin na": "COIN NOT AVAILABLE",
		"insert":  "INSERT COIN",
	}
	return m
}

// ToString will return a string representation of this machine
func (m *Machine) ToString() string {
	outstr := "\n == A Vending Machine == \n"
	outstr += " Running Total:\t" + strconv.Itoa(m.RunningTotal) + "\n"
	outstr += " ======================= \n"
	return outstr
}

// AcceptCoins will add the input coin to the running total
func (m *Machine) AcceptCoins(inputCoin Coins.Coin) error {
	var err error
	c := IdentifyCoin(inputCoin)
	if IsValidCoinValue(c) {
		m.RunningTotal += c
		m.InputCoins[inputCoin]++
        dollars := float64(m.RunningTotal)/100.0
        fmt.Println( "$ " + strconv.FormatFloat(dollars,'f',2,32)  )
	} else {
		err = errors.New(m.messages["invalid"])
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
	err = errors.New(m.messages["coin na"])
	return heldCoin, err
}

// ReturnAllCoins will return all coins that the customer has input
func (m *Machine) ReturnAllCoins() ([]Coins.Coin) {
	var outCoins []Coins.Coin
	for typeOfCoin, numberOfCoins := range m.InputCoins {
		for number := 0; number < numberOfCoins; number++ {
			coin, _ := m.ReturnCoin(typeOfCoin)
			outCoins = append(outCoins, coin)
		}
	}
    fmt.Println( m.messages["insert"])
	return outCoins
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
		if coin.Thickness == coinProperties["thickness"] && coin.Weight == coinProperties["weight"] && coin.Diameter == coinProperties["diameter"] {
			return int(coinProperties["value"])
		}
	}
	return 0
}
