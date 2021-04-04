package stats

import "github.com/muhammad6661/bank/v2/pkg/types"

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {

	Avg := make(map[types.Category]types.Money)
	index := make(map[types.Category]types.Money)

	for _, i := range payments {

		_, key := Avg[i.Category]
		if key == true {
			Avg[i.Category] += i.Amount
			index[i.Category]++
		} else {
			Avg[i.Category] = i.Amount
			index[i.Category] = 1
		}
	}
	for key, value := range Avg {
		Avg[key] = value / index[key]
	}
	return Avg
}