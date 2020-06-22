package hash

import "fmt"

type Orders struct {
	TableNum int
	Dish     string
}

func TableOrdersMain(orders []Orders) [][]string {
	hashmap, dishes := tableOrders(orders)
	fmt.Println(hashmap)
	fmt.Println(dishes)

	result := [][]string{}
	row0 := []string{}
	for dish := range dishes {
		row0 = append(row0, dish)
	}
	result = append(result, row0)
	for tableNum := range hashmap {
		for dish := range dishes {
			fmt.Println(hashmap[tableNum][dish])
		}
	}

	return nil
}

// return map[table] map[dish] orderCount
func tableOrders(orders []Orders) (map[int]map[string]int, map[string]bool) {
	dishes := map[string]bool{}
	result := map[int]map[string]int{}

	for _, order := range orders {
		if result[order.TableNum] == nil {
			result[order.TableNum] = map[string]int{}
		}
		result[order.TableNum][order.Dish]++
		dishes[order.Dish] = true
	}

	return result, dishes
}

/*
	result := hash.TableOrdersMain([]hash.Orders{
		{TableNum: 1, Dish: "A"},
		{TableNum: 1, Dish: "B"},
		{TableNum: 1, Dish: "C"},
		{TableNum: 2, Dish: "B"},
		{TableNum: 2, Dish: "B"},
		{TableNum: 3, Dish: "C"},
		{TableNum: 3, Dish: "D"},
	})
	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

*/
