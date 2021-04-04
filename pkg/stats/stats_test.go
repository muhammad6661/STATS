package stats

import (
	"reflect"
	"testing"

	"github.com/muhammad6661/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments :=  []types.Payment{
        {
			ID:  1,
			Amount:  10,
			Category: "auto",
		},
		{
			ID :  2,
			Amount : 50,
			Category: "auto",
		},
		{
			ID:  3,
			Amount: 10,
			Category: "food",
		},
			{
			ID:  4,
			Amount: 30,
			Category: "food",
		},
			{
			ID:  4,
			Amount: 10,
			Category: "fun",
		},

	}

    expected:=map[types.Category]types.Money{
		"auto": 30,
		"food": 20,
		"fun" : 10,
	}

	result:=CategoriesAvg(payments)

	if !reflect.DeepEqual(expected,result){
		t.Errorf("Invalid match result map, expected: %v,actual: %v",expected,result)
	}
}