package piscine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) int {
	menu := food{
		burger:  15,
		chips:   10,
		nuggets: 12,
	}
	if order == "burger" {
		return menu.burger
	}
	if order == "chips" {
		return menu.chips
	}
	if order == "nuggets" {
		return menu.nuggets
	}
	return 404
}
