package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units = make(map[string]int)
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728

    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, ok := units[unit]
    if !ok {
        return false
    }
    bill[item] += val

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitVal, ok := units[unit]
	if !ok {
		return false
	}

	current, ok := bill[item]
	if !ok || current < unitVal {
		return false
	}

	if current == unitVal {
		delete(bill, item)
	} else {
		bill[item] -= unitVal
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemInBill, ok := bill[item]

    if !ok {
        return 0,false
    }
    return itemInBill, true
}
