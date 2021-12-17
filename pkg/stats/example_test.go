package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Munirkhuja/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 0, Amount: 10_00, Category: "phone", Status: types.StatusOk},
		{ID: 1, Amount: 10_00, Category: "phone", Status: types.StatusFail},
		{ID: 2, Amount: 1_00, Category: "phone", Status: types.StatusOk},
		{ID: 3, Amount: 5_00, Category: "phone", Status: types.StatusInProgress},
		{ID: 4, Amount: 5_000_00, Category: "phone", Status: types.StatusOk},
	}
	fmt.Println(Avg(payments))
	//Output: 125400
}

func TestCategoriesAvgEx(t *testing.T) {	
	payments := []types.Payment{
		{ID: 0, Amount: 10_00, Category: "phone", Status: types.StatusOk},
		{ID: 1, Amount: 10_00, Category: "phone", Status: types.StatusFail},
		{ID: 2, Amount: 1_00, Category: "phone", Status: types.StatusOk},
		{ID: 3, Amount: 5_00, Category: "phone", Status: types.StatusInProgress},
		{ID: 4, Amount: 5_000_00, Category: "auto", Status: types.StatusOk},
		{ID: 4, Amount: 5_00, Category: "auto", Status: types.StatusOk},
		{ID: 4, Amount: 5_000_00, Category: "auto", Status: types.StatusOk},
	}
	expected := map[types.Category]types.Money{
		"phone":5_33,
		"auto":3_335_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}
func TestPeriodsDynamicEx(t *testing.T) {	
	first := map[types.Category]types.Money{
		"auto": 10_00,
		"food": 20_00,
	}
	second := map[types.Category]types.Money{
		"auto": 5_00,
		"food": 3_00,
	}
	expected := map[types.Category]types.Money{
		"auto":-5_00,
		"food":-17_00,
	}
	result := PeriodsDynamic(first,second)
	if !reflect.DeepEqual(result,expected) {
		t.Errorf("invalid result, expected: %v actual: %v",expected,result)
	}
}
func ExampleTotalInCategory() {
	
	payments := []types.Payment{
		{ID: 0, Amount: 10_00, Category: "phone", Status: types.StatusOk},
		{ID: 1, Amount: 10_00, Category: "phone", Status: types.StatusFail},
		{ID: 2, Amount: 1_00, Category: "phone", Status: types.StatusOk},
		{ID: 3, Amount: 5_00, Category: "phone", Status: types.StatusInProgress},
		{ID: 4, Amount: 5_000_00, Category: "auto", Status: types.StatusOk},
	}
	fmt.Println(TotalInCategory(payments,"phone"))
	//Output: 1600
	
}