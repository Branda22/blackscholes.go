package main

import (
	"fmt"

	"github.com/branda22/blackscholes/blackscholes"
)

func main() {
	option := &blackscholes.Option{
		StrikePrice:      955,
		TimeToExpiration: 48,
		Type:             "CALL",
	}

	underlying := &blackscholes.Underlying{
		Symbol:     "GOOG",
		Price:      959.11,
		Volatility: .13,
	}

	bs := blackscholes.NewBlackScholes(option, underlying, .0102)

	fmt.Println("delta", bs.Delta)
	fmt.Println("IV", bs.ImpliedVolatility)
	fmt.Println("Theo price", bs.TheoPrice)
	fmt.Println("Theta", bs.Theta)
}
