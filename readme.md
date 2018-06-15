# BlackScholes.go

### Example

```go
package main

import (
	"fmt"

	"github.com/branda22/blackscholes/blackscholes"
)

func main() {
	option := &blackscholes.Option{
		StrikePrice:      3500,
		TimeToExpiration: 87,
		Type:             "CALL",
	}

	underlying := &blackscholes.Underlying{
		Symbol:     "BTC_USD",
		Price:      4410.00,
		Volatility: .8915,
	}

	bs := blackscholes.NewBlackScholes(option, underlying, .0102)

	fmt.Println("delta", bs.Delta)
	fmt.Println("IV", bs.ImpliedVolatility)
	fmt.Println("Theo price", bs.TheoPrice)
	fmt.Println("Theta", bs.Theta)
}

```

