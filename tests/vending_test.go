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
            t.Errorf("expecting " + strconv.Itoa(localTotal) +", got " + strconv.Itoa(machine.RunningTotal) )
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
        t.Errorf("expecting " + strconv.Itoa(coinTestSum) + ", got " + strconv.Itoa(machine.RunningTotal))
    }
}

func TestIdentifyCoins(t *testing.T) {
    coin := Coins.NewCoin("penny")
    if coin.Value != 1 {
        t.Errorf("expecting " + strconv.Itoa(1) + ", got " + strconv.Itoa(coin.Value))
    }
    
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}