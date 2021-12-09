package utility

/*
This utility class create the payout slice and is used by the service.
If the amount exceeds a certain criteria then we split the payout.
For example: if maximumAmount=30000 and totalAmount >=maximumAmount the we will split the payout into 2 transactions
*/

import (
	"takeHomeTest/models"
)

const MaximumAmount = 30000

var payoutJson []models.Payout

func CreatePayout(items []models.Item) ([]models.Payout, error) {
	payoutJson = make([]models.Payout, 0)
	payoutMap := make(map[models.PayoutKey]models.Payout)
	var key models.PayoutKey
	for _, item := range items {
		key = models.PayoutKey{SellerReference: item.SellerReference, Currency: item.PriceCurrency}
		if value, ok := payoutMap[key]; ok {
			UpdatePayoutJson(value, item)
			payoutMap[key] = value
		} else {
			newPayout := models.Payout{
				SellerReference: item.SellerReference,
				Amount:          item.PriceAmount,
				Currency:        item.PriceCurrency,
			}

			UpdatePayoutJson(newPayout, item)
			payoutMap[key] = newPayout
		}
	}
	return splitPayout(), nil
}

func splitPayout() []models.Payout {
	for index := range payoutJson {
		if payoutJson[index].Amount > MaximumAmount {
			newPayouts := createNewPayouts(&payoutJson[index])
			payoutJson = append(payoutJson[:index], payoutJson[index+1:]...)
			payoutJson = append(payoutJson, newPayouts...)
		}
	}
	return payoutJson
}

func createNewPayouts(payout *models.Payout) []models.Payout {
	var newPayouts []models.Payout
	for payout.Amount >= MaximumAmount {
		po := models.Payout{
			SellerReference: payout.SellerReference,
			Amount:          payout.Amount / 2,
			Currency:        payout.Currency}
		createNewPayouts(&po)
		newPayouts = append(newPayouts, po)
		createNewPayouts(&po)
		newPayouts = append(newPayouts, po)
		payout.Amount /= 2
	}

	return newPayouts
}

func UpdatePayoutJson(value models.Payout, item models.Item) {
	for index := range payoutJson {
		if payoutJson[index].SellerReference == value.SellerReference && payoutJson[index].Currency == value.Currency {
			payoutJson[index].Amount += item.PriceAmount
			return
		}
	}
	payoutJson = append(payoutJson, value)
}
