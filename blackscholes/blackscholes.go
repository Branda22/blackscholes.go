package blackscholes

import (
	math "math"

	"github.com/chobie/go-gaussian"
)

type Option struct {
	StrikePrice      float64
	TimeToExpiration float64
	Type             string
}

type Underlying struct {
	Symbol     string
	Price      float64
	Volatility float64
}

type BS struct {
	StrikePrice          float64
	UnderlyingPrice      float64
	RiskFreeInterestRate float64
	Volatility           float64
	TimeToExpiration     float64
	Type                 string
	D1                   float64
	D2                   float64
	Delta                float64
	Theta                float64
	ImpliedVolatility    float64
	TheoPrice            float64
	Norm                 *gaussian.Gaussian
}

func NewBlackScholes(option *Option, underlying *Underlying, riskFreeInterestRate float64) *BS {
	bs := &BS{
		StrikePrice:          option.StrikePrice,
		UnderlyingPrice:      underlying.Price,
		RiskFreeInterestRate: riskFreeInterestRate,
		Volatility:           underlying.Volatility,
		TimeToExpiration:     float64(option.TimeToExpiration / 365),
		Type:                 option.Type,
	}

	bs.Initialize()

	return bs
}

func (bs *BS) Initialize() {
	bs.Norm = gaussian.NewGaussian(0, 1)
	bs.D1 = bs.calcD1(bs.UnderlyingPrice, bs.StrikePrice, bs.RiskFreeInterestRate, bs.TimeToExpiration, bs.Volatility)
	bs.D2 = bs.calcD2(bs.D1, bs.Volatility, bs.TimeToExpiration)
	bs.Delta = bs.calcDelta()
	bs.TheoPrice = bs.calcTheoreticalPrice()
	bs.ImpliedVolatility = bs.calcIv()
	bs.Theta = bs.calcTheta()
}

func (bs *BS) HistoricalVolatility() {

}

func (bs *BS) StandardDeviation(days int, dataPoints []float64) float64 {
	data := dataPoints[len(dataPoints)-days:]

	var total float64

	for _, d := range data {
		total += d
	}

	mean := total / float64(days)

	var temp float64

	for _, d := range data {
		temp += math.Pow(d-mean, 2)
	}

	return math.Sqrt(temp / float64(days))
}

func (bs *BS) calcD1(underlyingPrice float64, strikePrice float64, riskFreeInterestRate float64, timeToExpiration float64, volatility float64) float64 {
	return (math.Log(underlyingPrice/strikePrice) + (riskFreeInterestRate+math.Pow(volatility, 2)/2)*timeToExpiration) / (volatility * math.Sqrt(timeToExpiration))
}

func (bs *BS) calcD2(d1 float64, volatility float64, timeToExpiration float64) float64 {
	return d1 - (volatility * math.Sqrt(timeToExpiration))
}

func (bs *BS) calcDelta() float64 {
	delta := bs.Norm.Cdf(bs.D1)
	if bs.Type == "CALL" {
		return delta
	}

	return delta - 1.0
}

func (bs *BS) calcTheta() float64 {
	return -((bs.UnderlyingPrice * bs.Volatility * bs.Norm.Cdf(bs.D1)) / (2 * math.Sqrt(bs.TimeToExpiration)) - bs.RiskFreeInterestRate * bs.StrikePrice * math.Exp(-bs.RiskFreeInterestRate * (bs.TimeToExpiration)) * bs.Norm.Cdf(bs.D2)) / 365
}

// func (bs *BS) calcVega() float64 {
	
// }

func (bs *BS) calcIv() float64 {
	vol := math.Sqrt(2*math.Pi/bs.TimeToExpiration) * bs.TheoPrice / bs.UnderlyingPrice

	for i := 0; i < 100; i++ {
		// d1 := (math.Log(bs.UnderlyingPrice/bs.StrikePrice) + (bs.RiskFreeInterestRate+0.5*math.Pow(vol, 2))*bs.TimeToExpiration) / (vol * math.Sqrt(bs.TimeToExpiration))
		// d2 := d1 - vol*math.Sqrt(bs.TimeToExpiration)

		d1 := bs.calcD1(bs.UnderlyingPrice, bs.StrikePrice, bs.RiskFreeInterestRate, bs.TimeToExpiration, vol)
		d2 := bs.calcD2(d1, vol, bs.TimeToExpiration)
		vega := bs.UnderlyingPrice * bs.Norm.Cdf(d1) * math.Sqrt(bs.TimeToExpiration)

		cp := 1.0
		if bs.Type == "PUT" {
			cp = -1
		}

		price0 := cp*bs.UnderlyingPrice*bs.Norm.Cdf(cp*d1) - cp*bs.StrikePrice*math.Exp(bs.RiskFreeInterestRate*bs.TimeToExpiration)*bs.Norm.Cdf(cp*d2)
		vol = vol - (price0-bs.TheoPrice)/vega

		if math.Abs(price0-bs.TheoPrice) < math.Pow(10, -25) {
			break
		}
	}
	return vol
}

func (bs *BS) calcTheoreticalPrice() float64 {
	normD1 := bs.Norm.Cdf(bs.D1)
	normD2 := bs.Norm.Cdf(bs.D2)

	return bs.UnderlyingPrice*normD1 - bs.StrikePrice*math.Pow(math.E, -bs.RiskFreeInterestRate*bs.TimeToExpiration)*normD2
}
