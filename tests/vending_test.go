package main

import (
    "GoVending"
    "testing"
    "fmt"
    "strconv"
    "os"
)

var machine = VendingMachine.NewMachine()

func TestAcceptCoins(t *testing.T) {
    localTotal := 0
    localTotal +=5
    err := machine.AcceptCoins(5)
    if err != nil {
        fmt.Println(err)
    }
    if err != nil && machine.RunningTotal != localTotal {
        t.Errorf("expecting " + strconv.Itoa(localTotal) +", got " + strconv.Itoa(machine.RunningTotal) )
    }
    localTotal+=10
    err = machine.AcceptCoins(10)
    if err != nil {
        fmt.Println(err)
    }
    if err != nil && machine.RunningTotal != localTotal{
        t.Errorf("expecting " + strconv.Itoa(localTotal) +", got " + strconv.Itoa(machine.RunningTotal))
    }
    localTotal+=25
    err = machine.AcceptCoins(25)
    if err != nil {
        fmt.Println(err)
    }
    if err != nil && machine.RunningTotal != localTotal{
        t.Errorf("expecting " + strconv.Itoa(localTotal) +", got " + strconv.Itoa(machine.RunningTotal))
    }
    // do not add 11 to running total; no 11 value coins 
    err = machine.AcceptCoins(11) // shouldn't accept invalid coin value
    if err != nil {
        fmt.Println(err)
    } 
    if err != nil && machine.RunningTotal != localTotal {
        t.Errorf("expecting " + strconv.Itoa(localTotal) +", got " + strconv.Itoa(machine.RunningTotal))
    }
    
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}