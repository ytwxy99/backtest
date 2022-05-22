package market

import (
	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/utils"
)

type Average struct {
	CurrencyPair string
	Level        string
	MA           int
	Markets      interface{}
}

// Average
// * @Description: five's average market index
// * @param: the level of markets which support:
//           10s, 1m, 5m, 15m, 30m, 1h, 4h, 8h, 1d, 7d
//
func (average *Average) Average(backOff bool, index int) float64 {
	var sum float64

	if average.MA == utils.MA21 && average.Level == utils.Level4Hour {
		markets := average.Markets.([]*database.HistoryFourHour)
		if len(markets) == 0 {
			return 0
		}

		if index < 22 {
			return 0
		}

		if !backOff {
			for i, market := range markets {
				if i > index-utils.MA21 && i <= index {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA21)
		} else {
			for i, market := range markets {
				if i > index-utils.MA21-1 && i <= index-1 {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA21)
		}

	} else if average.MA == utils.MA10 && average.Level == utils.Level4Hour {
		markets := average.Markets.([]*database.HistoryFourHour)
		if len(markets) == 0 {
			return 0
		}

		if index == 0 {
			return 0
		}

		if !backOff {
			for i, market := range markets {
				if i > index-utils.MA10 && i <= index {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA10)
		} else {
			for i, market := range markets {
				if i > index-utils.MA10-1 && i <= index-1 {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA10)
		}

	} else if average.MA == utils.MA5 && average.Level == utils.Level4Hour {
		markets := average.Markets.([]*database.HistoryFourHour)
		if len(markets) == 0 {
			return 0
		}

		if index == 0 {
			return 0
		}

		if !backOff {
			for i, market := range markets {
				if i > index-utils.MA5 && i <= index {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA5)
		} else {
			for i, market := range markets {
				if i > index-utils.MA5-1 && i <= index-1 {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA5)
		}

	} else if average.MA == utils.MA21 && average.Level == utils.Level30Min {
		markets := average.Markets.([]*database.HistoryThirtyMin)
		if len(markets) == 0 {
			return 0
		}

		if index < 22 {
			return 0
		}

		if !backOff {
			for i, market := range markets {
				if i > index-utils.MA21 && i <= index {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA21)
		} else {
			for i, market := range markets {
				if i > index-utils.MA21-1 && i <= index-1 {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA21)
		}

	} else if average.MA == utils.MA10 && average.Level == utils.Level30Min {
		markets := average.Markets.([]*database.HistoryThirtyMin)
		if len(markets) == 0 {
			return 0
		}

		if index == 0 {
			return 0
		}

		if !backOff {
			for i, market := range markets {
				if i > index-utils.MA10 && i <= index {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA10)
		} else {
			for i, market := range markets {
				if i > index-utils.MA10-1 && i <= index-1 {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA10)
		}

	} else if average.MA == utils.MA5 && average.Level == utils.Level30Min {
		markets := average.Markets.([]*database.HistoryThirtyMin)
		if len(markets) == 0 {
			return 0
		}

		if index == 0 {
			return 0
		}

		if !backOff {
			for i, market := range markets {
				if i > index-utils.MA5 && i <= index {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA5)
		} else {
			for i, market := range markets {
				if i > index-utils.MA5-1 && i <= index-1 {
					sum = sum + utils.StringToFloat64(market.Price)
				}
			}

			return sum / float64(utils.MA5)
		}
	} else {
		return 0
	}
}
