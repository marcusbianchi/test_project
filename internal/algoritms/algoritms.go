package algoritms

import "math"

var counter = 0

type GiftCard struct {
	// customer string //Nil
	Id               int
	Brand            string
	Face_value       int
	Discounted_value int //Customer payed value
}

type Basket struct {
	Customer_id   int
	Gift_card_ids []int //Customer payed value
}

func (b Basket) AddGiftToBasket(gift_card_id int) Basket {
	//validations
	b.Gift_card_ids = append(b.Gift_card_ids, gift_card_id)
	return b
}

//The system should be demo-able via unit tests (preferably) or the console.
//The system does not require persistence.
//The system should support more than one customer.
//There should be some default gift cards which you can create using static data.
//Multiple customers should be able to add multiple types of gift cards to their basket.
//Gift cards have a face value, and a price which is discounted by the brand. The customer pays the discounted price.
//The discount, and therefore the price of a gift card, can be changed at any time.
//Provide a mechanism to calculate a customer’s basket total price. Customers’ basket values should reflect the updated prices.

func Hello(name string) string {
	return "Hello " + name
}

func NewBasket(customer_id int, gift_card_id int) Basket {
	//validations
	return Basket{
		Customer_id:   customer_id,
		Gift_card_ids: []int{gift_card_id},
	}
}

func NewGiftCard(brand string, face_value int, discount_percentage int) GiftCard {
	id := counter
	counter += 1
	discounted := CalculateDiscountValue(discount_percentage, face_value)
	return GiftCard{
		Id:               id,
		Brand:            brand,
		Face_value:       face_value,
		Discounted_value: discounted,
	}
}

func CalculateDiscountValue(discount_percentage int, face_value int) int {
	float_value := float64(discount_percentage) / 10000
	discounted := float64(face_value) * (1 - float_value)
	return int(math.Floor(discounted))
}

func AddGiftToBasket(basket Basket, gift_card_id int) Basket {
	//validations
	basket.Gift_card_ids = append(basket.Gift_card_ids, gift_card_id)
	return basket
}

func UpdateGiftCardValue(giftCard GiftCard, face_value int, discount_percentage int) GiftCard {
	discounted := CalculateDiscountValue(discount_percentage, face_value)
	giftCard.Discounted_value = discounted
	giftCard.Face_value = face_value
	return giftCard
}

//func CalculateBasketValue(giftCards map[int]GiftCard, basket Basket) int {
//	for i := range collection {
//
//	}
//}

//func NewGiftCardAmazon5Pound() GiftCard {
//	id := counter
//	counter += 1
//	return GiftCard{
//		id:               id,
//		brand:            "amazon",
//		face_value:       1000,
//		discounted_value: 950,
//	}
//}
//
//func NewGiftCardAmazon10Pound() GiftCard {
//	id := counter
//	counter += 1
//	return GiftCard{
//		id:               id,
//		brand:            "amazon",
//		face_value:       500,
//		discounted_value: 475,
//	}
//}
//
//func NewGiftCardStarbucks5Pound() GiftCard {
//	id := counter
//	counter += 1
//	return GiftCard{
//		id:               id,
//		brand:            "starbucks",
//		face_value:       1000,
//		discounted_value: 900,
//	}
//}
//
//func NewGiftCardStarbucks10Pound() GiftCard {
//	id := counter
//	counter += 1
//	return GiftCard{
//		id:               id,
//		brand:            "starbucks",
//		face_value:       500,
//		discounted_value: 450,
//	}
//}

//employeeSalary := make(map[string]int)
//p := []string{"a", "b", "c"}
//p = append(p, "d")
//s := make([]string, 2, 3)
//for k, v := range myMap {
//	keys = append(keys, k)
//	values = append(values, v)
//}

// GitftCards [1: GiftCard]
