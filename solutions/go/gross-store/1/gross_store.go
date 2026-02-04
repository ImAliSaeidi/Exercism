package gross

func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, ok := units[unit]
	if !ok {
		return false
	}

	billValue, ok := bill[item]
	if !ok {
		bill[item] = unitValue
		return true
	}

	bill[item] = billValue + unitValue
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billValue, isItemExist := bill[item]
	if !isItemExist {
		return false
	}

	unitValue, isUnitExist := units[unit]
	if !isUnitExist {
		return false
	}

	newQuantity := billValue - unitValue

	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, ok := bill[item]
	if !ok {
		return 0, false
	}
	return value, true
}
