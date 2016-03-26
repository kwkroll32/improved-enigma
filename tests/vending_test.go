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

// initialize some coins for the tests
var quarter = Coins.NewQuarter()
var dime = Coins.NewDime()
var nickel = Coins.NewNickel()
var penny = Coins.NewPenny()

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
			if dollars := float64(machine.RunningTotal) / 100.0; machine.Display != "$"+strconv.FormatFloat(dollars, 'f', 2, 32) {
				throwTestingErrorDisplayString(t, machine.Display, "$"+strconv.FormatFloat(dollars, 'f', 2, 32))
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
		penny,
		nickel,
		dime,
		quarter}
	for _, coin := range coinTests {
		testAddingThisCoin(coin)
	}
}

func TestIdentifyCoins(t *testing.T) {
	for coin, expVal := range map[Coins.Coin]int{penny: 1, nickel: 5, dime: 10, quarter: 25} {
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
	machine.AcceptCoins(nickel)
	if dollars := float64(machine.RunningTotal) / 100.0; machine.Display != "$"+strconv.FormatFloat(dollars, 'f', 2, 32) {
		throwTestingErrorDisplayString(t, machine.Display, "$"+strconv.FormatFloat(dollars, 'f', 2, 32))
	}
	machine.AcceptCoins(nickel)
	machine.AcceptCoins(nickel)
	machine.AcceptCoins(quarter)
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
		throwTestingErrorDisplayString(t, machine.Display, machine.Messages["insert"])
	}
}

func TestCustomerSelectsAtExactChange(t *testing.T) {
	machine.ShowSelections()
	machine.RunningTotal = machine.Products["chips"]
	machine.SelectProduct("chips")
	if machine.Display != machine.Messages["thanks"] {
		throwTestingErrorDisplayString(t, machine.Messages["thanks"], machine.Display)
	}
	if machine.RunningTotal != 0 {
		throwTestingErrorInt(t, 0, machine.RunningTotal)
	}
}

func TestMachineMakeChange(t *testing.T) {
	var change map[Coins.Coin]int
	// return a quarter
	machine.RunningTotal = 25
	change = machine.DispenseChange()
	if change[quarter] != 1 {
		t.Errorf("expected 1 quarter, got " + strconv.Itoa(change[quarter]))
	}
	// return two dimes
	machine.RunningTotal = 20
	change = machine.DispenseChange()
	if change[dime] != 2 {
		t.Errorf("expected 2 dimes, got " + strconv.Itoa(change[dime]))
	}
}

func TestCustomerSelectsWithMoreThanEnoughMoney(t *testing.T) {
	machine.ShowSelections()
	machine.RunningTotal = machine.Products["chips"] + 66
	machine.SelectProduct("chips")
	if machine.Display != machine.Messages["thanks"] {
		throwTestingErrorDisplayString(t, machine.Messages["thanks"], machine.Display)
	}
	change := machine.DispenseChange()
	if change[quarter] != 2 {
		throwTestingErrorInt(t, 2, change[quarter])
	}
	for _, coin := range []Coins.Coin{dime, nickel, penny} {
		if change[coin] != 1 {
			throwTestingErrorInt(t, 1, change[coin])
		}
	}
	if machine.RunningTotal != 0 {
		throwTestingErrorInt(t, 0, machine.RunningTotal)
	}
}

func TestCustomerSelectsWithInsufficientMoney(t *testing.T) {
    machine.RunningTotal = machine.Products["chips"]/2
    machine.SelectProduct("chips")
    if machine.Display != machine.Messages["insert"] {
		throwTestingErrorDisplayString(t, machine.Messages["insert"],machine.Display)
	}
    if machine.RunningTotal != machine.Products["chips"]/2 {
        throwTestingErrorInt(t, machine.Products["chips"]/2, machine.RunningTotal)
    }
}

func TestSelectSomethingSoldOut(t *testing.T) {
    machine.RunningTotal = machine.Products["chips"]
    machine.Stock["chips"] = 0
    machine.SelectProduct("chips")
    if machine.Display == machine.Messages["thanks"] {
		throwTestingErrorDisplayString(t,machine.Messages["sold out"], machine.Display )
	}
    if machine.RunningTotal != machine.Products["chips"] {
		throwTestingErrorInt(t, machine.Products["chips"], machine.RunningTotal)
	}
}

func TestExactChangeOnly(t *testing.T) {
    machine.Bank = map[Coins.Coin]int{Coins.NewPenny():0, Coins.NewNickel():0, Coins.NewDime():0, Coins.NewQuarter():0}
    if !machine.NeedsExactChange() {
        t.Errorf("machine needs exact change, but doesn't know it")
    }
    if machine.Display != machine.Messages["exact change"] {
        throwTestingErrorDisplayString(t, machine.Messages["exact change"], machine.Display)
    }
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
