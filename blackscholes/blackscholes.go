package blackscholes

import (
	math "math"

	"github.com/chobie/go-gaussian"
)

type BS struct {
	StrikePrice          float64
	UnderlyingPrice      float64
	RiskFreeInterestRate float64
	Volatility           float64
	TimeToExpiration     float64
	Norm                 *gaussian.Gaussian
}

func NewBlackScholes(strikePrice float64, underlyingPrice float64, riskFreeInterestRate float64, volatility float64, timeToExpiration float64) BS {
	bs := BS{
		strikePrice,
		underlyingPrice,
		riskFreeInterestRate,
		volatility,
		float64(timeToExpiration / 365),
		gaussian.NewGaussian(0, 1),
	}

	return bs
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

func (bs *BS) D1() float64 {
	return (math.Log(bs.UnderlyingPrice/bs.StrikePrice) + (bs.RiskFreeInterestRate+math.Pow(bs.Volatility, 2)/2)*bs.TimeToExpiration) / (bs.Volatility * math.Sqrt(bs.TimeToExpiration))
}

func (bs *BS) D2() float64 {
	return bs.D1() - (bs.Volatility * math.Sqrt(bs.TimeToExpiration))
}

func (bs *BS) Delta() float64 {
	return bs.Norm.Cdf(bs.D1())
}

func (bs *BS) TheoreticalPrice() {

}
