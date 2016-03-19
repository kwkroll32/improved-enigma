package main

import (
	"GoVending"
	"GoVending/Coins"
	"fmt"
	"os"
	"strconv"
	"testing"
	//"errors"
)

// initialize a vending machine instance for the tests
var machine = VendingMachine.NewMachine()

/* throwing all errors from the same function means that `go test` will report the same line number for all errors
   but still tells from which function the error originated */
func throwTestingErrorInt(t *testing.T, expected, received int) {
	t.Errorf("expecting " + strconv.Itoa(expected) + ", got " + strconv.Itoa(received))
}

func loadACoin(t *testing.T) func(coinIn Coins.Coin) {
	// initialize the local total and error object
	// these are modified in subsequent runs of the closure
	localTotal := 0
	var err error
	// the closure
	return func(coinIn Coins.Coin) {
		err = machine.AcceptCoins(coinIn)
		if err == nil {
			// the machine took the coin fine
			localTotal += coinIn.Value
		} else if err != nil {
			// the machine failed to take this coin
			// the coin should be returned
			fmt.Println(err)
		}
		if machine.RunningTotal != localTotal {
			// the totals dont add up
			throwTestingErrorInt(t, localTotal, machine.RunningTotal)
		}
	}
}

func TestAcceptCoins(t *testing.T) {
	testAddingThisCoin := loadACoin(t)
	coinTests := []Coins.Coin{
		Coins.NewCoin("penny"),
		Coins.NewCoin("nickel"),
		Coins.NewCoin("dime"),
		Coins.NewCoin("quarter"),
		Coins.NewCoin("11-cent")}
	for _, coin := range coinTests {
		testAddingThisCoin(coin)
	}
}

func TestIdentifyCoins(t *testing.T) {
	var coin Coins.Coin
	for name, expVal := range map[string]int{"penny": 1, "nickel": 5, "dime": 10, "quarter": 25} {
		coin = Coins.NewCoin(name)
		if VendingMachine.IdentifyCoin(coin) != expVal {
			throwTestingErrorInt(t, expVal, coin.Value)
		}
	}
}

func TestReturnAllCoinsLoop(t *testing.T) {
	for heldCoin := range machine.InputCoins {
		machine.ReturnCoin(heldCoin)
	}
	coinCount := 0
	for _, numberOfCoins := range machine.InputCoins {
		coinCount += numberOfCoins
	}
	if coinCount != 0 {
		throwTestingErrorInt(t, 0, coinCount)
	}
	if machine.RunningTotal != 0 {
		throwTestingErrorInt(t, 0, machine.RunningTotal)
	}
}

func TestReturnAllCoinsMachineFunction(t *testing.T) {
	machine.ReturnAllCoins()
	if machine.RunningTotal != 0 {
		throwTestingErrorInt(t, 0, machine.RunningTotal)
	}
	machine.AcceptCoins(Coins.NewCoin("nickel"))
	machine.AcceptCoins(Coins.NewCoin("nickel"))
	machine.AcceptCoins(Coins.NewCoin("nickel"))
	machine.AcceptCoins(Coins.NewCoin("quarter"))
	if machine.RunningTotal != 40 {
		throwTestingErrorInt(t, 40, machine.RunningTotal)
	}
	returnedCoins := machine.ReturnAllCoins()
	if len(returnedCoins) != 4 {
		throwTestingErrorInt(t, 4, len(returnedCoins))
	}
	if machine.RunningTotal != 0 {
		throwTestingErrorInt(t, 0, machine.RunningTotal)
	}
}

func TestShowSelections(t *testing.T) {
    machine.ShowSelections()
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
