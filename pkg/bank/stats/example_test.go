package stats

import (
	"fmt"

	"github.com/Munirkhuja/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 1, Amount: 10_00, Category: "phone"},
		{ID: 2, Amount: 1_00, Category: "phone"},
		{ID: 3, Amount: 5_00, Category: "phone"},
		{ID: 4, Amount: 5_000_00, Category: "phone"},
	}
	fmt.Println(Avg(payments))
	//Output: 125400
}

func ExampleTotalInCategory() {
	
	payments := []types.Payment{
		{ID: 1, Amount: 10_00, Category: "phone"},
		{ID: 2, Amount: 1_00, Category: "phone"},
		{ID: 3, Amount: 5_00, Category: "phone"},
		{ID: 4, Amount: 5_000_00, Category: "auto"},
	}
	fmt.Println(TotalInCategory(payments,"phone"))
	//Output: 1600
	
}