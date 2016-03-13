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
        fmt.Println(strconv.Itoa(localTotal))
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
    testAddingThisCoin(5)
    testAddingThisCoin(10)
    testAddingThisCoin(25)
    testAddingThisCoin(11)
    
    if machine.RunningTotal != (5 + 10 + 25) {
        t.Errorf("expecting " + strconv.Itoa((5 + 10 + 25)) + ", got " + strconv.Itoa(machine.RunningTotal))
    }
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}