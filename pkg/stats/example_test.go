package stats

import (
	"fmt"
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