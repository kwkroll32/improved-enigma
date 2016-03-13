package main

import (
    "GoVending"
    "testing"
    "fmt"
    "strconv"
)

func TestSetup(t *testing.T) {
    machine := VendingMachine.NewMachine()
    fmt.Println(machine.ToString())
    
}

func TestAcceptCoins(t *testing.T) {
    machine := VendingMachine.NewMachine()
    machine.AcceptCoins(5)
    if machine.RunningTotal != 5 {
        t.Errorf("expecting 5, got " + strconv.Itoa(machine.RunningTotal) )
    }
}

func main() {}