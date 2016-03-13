package main

import (
    "GoVending"
    "testing"
    "fmt"
)

func TestSetup(t *testing.T) {
    machine := VendingMachine.NewMachine()
    fmt.Println(machine.ToString())
    
}

func main() {}