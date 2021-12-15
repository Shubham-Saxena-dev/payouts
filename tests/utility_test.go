package tests

/*
This is a unit test class where utility methods are tested
*/

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"takeHomeTest/models"
	"takeHomeTest/utility"
	"testing"
)

var requestBody = `[
    {
        "name": "Top",
        "price_amount":10000,
        "price_currency":"USD",
        "seller_reference": 1
    },
	{
        "name": "Bottoms",
        "price_amount":50000,
        "price_currency":"USD",
        "seller_reference": 1
    }]`

func TestCreatePayout_whenPriceExceedsMaximumAmount_thenSplitPayouts(t *testing.T) {
	//given
	var items []models.Item
	json.Unmarshal([]byte(requestBody), &items)

	//when
	payout, err := utility.CreatePayout(items)

	//then
	assert.NoError(t, err)
	assert.Equal(t, len(payout), 2)
	var totalAmount int64
	for index := range payout {
		totalAmount += payout[index].Amount
		assert.Equal(t, int(payout[index].SellerReference), 1)
		assert.Equal(t, payout[index].Currency, "USD")
	}
	assert.Equal(t, int(totalAmount), 60000)
}
