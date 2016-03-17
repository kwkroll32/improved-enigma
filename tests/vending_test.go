package main

import (
    "GoVending"
    "GoVending/Coins"
    "testing"
    "fmt"
    "strconv"
    "os"
    //"errors"
)

// initialize a vending machine instance for the tests
var machine = VendingMachine.NewMachine()

func throwTestingErrorInt(t *testing.T, expected, received int) {
    t.Errorf("expecting " + strconv.Itoa(expected) +", got " + strconv.Itoa(received) )
}

func loadACoin(t *testing.T) func(coinIn int) {
    // initialize the local total and error object
    // these are modified in subsequent runs of the closure
    localTotal := 0
    var err error
    // the closure 
    return func(coinIn int)  {
        err = machine.AcceptCoins(coinIn)
        if err == nil {
            // the machine took the coin fine 
            localTotal += coinIn
        } else if err != nil {
            // the machine failed to take this coin 
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
    coinTests := []int{1,5,10,25,11}
    coinTestSum := 0
    for _,coin := range(coinTests) {
        testAddingThisCoin(coin)
        if VendingMachine.IsValidCoin(coin) {
            coinTestSum += coin 
        }
    }
    if machine.RunningTotal != coinTestSum {
        throwTestingErrorInt(t, coinTestSum, machine.RunningTotal)
    }
}

func TestIdentifyCoins(t *testing.T) {
    var coin Coins.Coin
    for name, expVal := range(map[string]int{"penny":1, "nickel":5, "dime":10, "quarter":25}) {
        coin = Coins.NewCoin(name)
        if VendingMachine.IdentifyCoin(coin) != expVal {
            throwTestingErrorInt(t, expVal, coin.Value)
        }
    }
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}