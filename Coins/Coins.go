package Coins

// Coin is a class to represent the inserted coins
type Coin struct {
    Value int
    Weight float32
    Diameter float32
    Thickness float32 
}

// CoinTypes is a map of valid coins and their respective properties
var CoinTypes = map[string][]float32{
        "penny":  {2.5, 19.05, 1.52}  ,
        "nickel": {5, 21.21, 1.95}    , 
        "dime":   {2.268, 17.91, 1.35}, 
        "quarter":{5.670, 24.26, 1.75}}
        
// NewCoin is a constructor for a new coin
func NewCoin(inCoinKind string) *Coin {
    coin := new(Coin)
    coin.Weight = CoinTypes[inCoinKind][0]
    coin.Diameter = CoinTypes[inCoinKind][1]
    coin.Thickness = CoinTypes[inCoinKind][2]
    return coin
}