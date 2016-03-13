package VendingMachine

// Machine is an class to represent the vending machine
type Machine struct {}

// NewMachine is a constructor for a new machine
func NewMachine() *Machine {
    return new(Machine)
}

// ToString will return a string representation of this machine
func (m Machine) ToString() string {
    return "a machine"
}