package VendingMachine
import (
    "errors"
    "strconv"
)

// Machine is an class to represent the vending machine
type Machine struct {
    RunningTotal int
}

// NewMachine is a constructor for a new machine
func NewMachine() *Machine {
    m := new(Machine)
    m.RunningTotal = 0
    return m
}

// ToString will return a string representation of this machine
func (m *Machine) ToString() string {
    return "a machine"
}

// AcceptCoins will add the input coin to the running total
func (m *Machine) AcceptCoins(c int) error {
    var err error
    if isValidCoin(c) { 
        m.RunningTotal += c
    } else {
        err = errors.New("invalid coin value: " + strconv.Itoa(c))
    }
    return err
}

func isValidCoin(c int) bool {
    validCoins := []int{1, 5, 10, 25}
    for _, vc := range validCoins {
        if c == vc {
            return true
        }
    }
    return false
}
