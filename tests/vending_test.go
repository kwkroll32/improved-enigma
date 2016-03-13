package main

import (
    "GoVending"
    "testing"
    "fmt"
    "strconv"
    "os"
    //"errors"
)

var machine = VendingMachine.NewMachine()

func loadACoin(t *testing.T) func(coinIn int) {
    localTotal := 0
    var err error
    return func(coinIn int)  {
        err = machine.AcceptCoins(coinIn)
        if err == nil {
            // the machine took the coin fine 
            localTotal += coinIn
            if machine.RunningTotal != localTotal {
                // although coin was accepted, the totals dont add up 
                t.Errorf("expecting " + strconv.Itoa(localTotal) +", got " + strconv.Itoa(machine.RunningTotal) )
            }
        } else if err != nil {
            // the machine failed to take this coin 
            fmt.Println(err)
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

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}