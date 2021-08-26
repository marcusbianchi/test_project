package algoritms_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/marcusbianchi/test_project/internal/algoritms"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := algoritms.Hello("Marcus")
	assert.Equal(t, "Hello Marcus", actual)
}

func TestNewGiftCard(t *testing.T) {
	actual := algoritms.NewGiftCard("amazon", 500, 500)
	assert.Equal(t, "amazon", actual.Brand)
	assert.Equal(t, 0, actual.Id)
	assert.Equal(t, 475, actual.Discounted_value)
	assert.Equal(t, 500, actual.Face_value)
}

func TestBasket(t *testing.T) {
	actual := algoritms.NewBasket(1, 0)
	assert.Equal(t, 1, actual.Customer_id)
	assert.Equal(t, 0, actual.Gift_card_ids[0])
	assert.Equal(t, 1, len(actual.Gift_card_ids))
}

func TestAddGiftToBasket(t *testing.T) {
	basket := algoritms.NewBasket(1, 0)
	actual := algoritms.AddGiftToBasket(basket, 1)
	assert.Equal(t, 1, actual.Customer_id)
	assert.Equal(t, 0, actual.Gift_card_ids[0])
	assert.Equal(t, 1, actual.Gift_card_ids[1])
	assert.Equal(t, 2, len(actual.Gift_card_ids))
}

func TestUpdateGiftCardValue(t *testing.T) {
	giftCard := algoritms.NewGiftCard("amazon", 500, 500)
	actual := algoritms.UpdateGiftCardValue(giftCard, 1000, 1000)
	assert.Equal(t, "amazon", actual.Brand)
	assert.Equal(t, 0, actual.Id)
	assert.Equal(t, 900, actual.Discounted_value)
	assert.Equal(t, 1000, actual.Face_value)
}
