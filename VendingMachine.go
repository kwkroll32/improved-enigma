package VendingMachine
import (
    "errors"
    "strconv"
)

/*
    Machine struct and bound functions 
*/

// Machine is a class to represent the vending machine
type Machine struct {
    RunningTotal int
    InputCoins map[int]int 
}

// NewMachine is a constructor for a new machine
func NewMachine() *Machine {
    m := new(Machine)
    m.RunningTotal = 0
    m.InputCoins = map[int]int{1: 0, 5: 0, 10: 0, 25: 0}
    return m
}

// ToString will return a string representation of this machine
func (m *Machine) ToString() string {
    outstr := "\n == A Vending Machine == \n"
    outstr += " Running Total:\t" + strconv.Itoa(m.RunningTotal) + "\n"
    outstr += " ======================= \n"
    return outstr
}

// AcceptCoins will add the input coin to the running total
func (m *Machine) AcceptCoins(c int) error {
    var err error
    if IsValidCoin(c) { 
        m.RunningTotal += c
        m.InputCoins[c]++
    } else {
        err = errors.New("invalid coin value: " + strconv.Itoa(c))
    }
    return err
}


/*
    Utility Functions 
*/

// IsValidCoin will check if the input is of value 1, 5, 10, or 25
func IsValidCoin(c int) bool {
    validCoins := []int{1, 5, 10, 25}
    for _, vc := range validCoins {
        if c == vc {
            return true
        }
    }
    return false
}
