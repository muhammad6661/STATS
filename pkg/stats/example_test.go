package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/muhammad6661/bank/v2/pkg/types"
)

func TestPeriodDynamic(t *testing.T) {
	 first :=map[types.Category]types.Money{
		"auto":10,
		"food":20,
	 } 
	 second :=map[types.Category]types.Money{
		"auto":5,
         "food":3,
	 }
	 expected:=map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}
   result:=PeriodsDynamic(first,second)

   for key:=range result{
	  fmt.Println(key,":",)
   }
   if !reflect.DeepEqual(expected,result){
	t.Errorf("Invalid match result map, expected: %v,actual: %v",expected,result)
}
}