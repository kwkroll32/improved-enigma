package main

import (
    "GoVending"
    "testing"
    "fmt"
    "strconv"
    "os"
)

var machine = VendingMachine.NewMachine()

func TestSetup(t *testing.T) {
    fmt.Println(machine.ToString())
    
}

func TestAcceptCoins(t *testing.T) {
    machine.AcceptCoins(5)
    if machine.RunningTotal != 5 {
        t.Errorf("expecting 5, got " + strconv.Itoa(machine.RunningTotal) )
    }
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}