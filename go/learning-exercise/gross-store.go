// https://exercism.org/tracks/go/exercises/gross-store
package learningexercise

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	pieces, exists := units[unit]

	if !exists {
		return false
	}

	bill[item] += pieces

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItemGrossStore(bill, units map[string]int, item, unit string) bool {
	pieces, exists := units[unit]
	bpieces, bexists := bill[item]

	if !exists || !bexists || pieces > bpieces {
		return false
	} else if bpieces == pieces {
		delete(bill, item)
	} else {
		bill[item] -= pieces
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItemGrossStore(bill map[string]int, item string) (int, bool) {
	pieces, exists := bill[item]

	return pieces, exists
}
