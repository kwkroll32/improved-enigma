package main

import (
	"GoVending"
	"GoVending/Coins"
	//"fmt"
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

/* throwing all errors from the same function means that `go test` will report the same line number for all errors
   but still tells from which function the error originated */
func throwTestingErrorDisplayString(t *testing.T, expected, received string) {
	t.Errorf("incorrect display. should be '" + expected + "', but it says '" + received + "'")
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
            if dollars := float64(machine.RunningTotal)/100.0; machine.Display != "$" + strconv.FormatFloat(dollars,'f',2,32) {
                throwTestingErrorDisplayString(t, machine.Display, "$" + strconv.FormatFloat(dollars,'f',2,32))
            }
			localTotal += coinIn.Value
		} else if err != nil {
			// the machine failed to take this coin
			// the coin should be returned
            if machine.Display != machine.Messages["invalid"] {
                throwTestingErrorDisplayString(t, machine.Messages["invalid"], machine.Display)
            }
			//fmt.Println(err)
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
    if machine.Display != machine.Messages["insert"] {
        throwTestingErrorDisplayString(t, machine.Display, machine.Messages["insert"])
    }
	if machine.RunningTotal != 0 {
		throwTestingErrorInt(t, 0, machine.RunningTotal)
	}
	machine.AcceptCoins(Coins.NewCoin("nickel"))
    if dollars := float64(machine.RunningTotal)/100.0; machine.Display != "$" + strconv.FormatFloat(dollars,'f',2,32) {
        throwTestingErrorDisplayString(t, machine.Display, "$" + strconv.FormatFloat(dollars,'f',2,32))
    }
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
    machine.Display = "force test fail"
    if machine.Display == machine.Messages["insert"] {
        throwTestingErrorDisplayString(t, machine.Display, machine.Messages["insert"] )
    }
}

func TestCustomerSelectsAtExactChange(t *testing.T) {
    machine.ShowSelections()
    machine.RunningTotal = machine.Products["chips"]
    machine.SelectProduct("chips")
    if machine.Display != machine.Messages["thanks"] {
        throwTestingErrorDisplayString(t, machine.Display, machine.Messages["thanks"] )
    }
    if machine.RunningTotal != 0 {
        throwTestingErrorInt(t, 0, machine.RunningTotal)
    }
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
