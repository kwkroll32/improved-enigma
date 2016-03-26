package Coins

// Coin is a class to represent the inserted coins
type Coin struct {
	Value     int
	Weight    float32
	Diameter  float32
	Thickness float32
}

// CoinTypes is a map of valid coins and their respective properties
var CoinTypes = map[string]map[string]float32{
	"penny":   {"weight": 2.5, "diameter": 19.05, "thickness": 1.52, "value": 1},
	"nickel":  {"weight": 5, "diameter": 21.21, "thickness": 1.95, "value": 5},
	"dime":    {"weight": 2.268, "diameter": 17.91, "thickness": 1.35, "value": 10},
	"quarter": {"weight": 5.670, "diameter": 24.26, "thickness": 1.75, "value": 25}}

// NewPenny is a constructor for a penny
func NewPenny() Coin {
    penny := Coin{Weight: 2.5, Diameter: 19.05, Thickness: 1.52, Value: 1}
    return penny
}
// NewNickel is a constructor for a nickel
func NewNickel() Coin {
    nickel := Coin{Weight: 5, Diameter: 21.21, Thickness: 1.95, Value: 5}
    return nickel
}
// NewDime is a constructor for a dime
func NewDime() Coin {
    dime := Coin{Weight: 2.268, Diameter: 17.91, Thickness: 1.35, Value: 10}
    return dime
}
// NewQuarter is a constructor for a quarter
func NewQuarter() Coin {
    quarter := Coin{Weight: 5.670, Diameter: 24.26, Thickness: 1.75, Value: 25}
    return quarter
}