package stats

import "github.com/Munirkhuja/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	count := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			count++			
		}
	}
	avg := sum / types.Money(count)
	return avg
}

func  TotalInCategory(payments []types.Payment, category types.Category) types.Money {	
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount			
		}
	}
	return sum
}