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

// NewCoin is a constructor for a new coin
func NewCoin(inCoinKind string) Coin {
	var coin Coin
	coin.Weight = CoinTypes[inCoinKind]["weight"]
	coin.Diameter = CoinTypes[inCoinKind]["diameter"]
	coin.Thickness = CoinTypes[inCoinKind]["thickness"]
	coin.Value = int(CoinTypes[inCoinKind]["value"])
	return coin
}
