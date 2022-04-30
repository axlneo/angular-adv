package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		var orderItems []models.OrderItem
		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(98) + 10)
			qty := uint(rand.Intn(5))

			orderItems = append(orderItems, models.OrderItem{
				ProductTitle:      faker.Word(),
				Price:             price,
				Quantity:          qty,
				AdminRevenue:      0.9 * float64(qty),
				AmbassadorRevenue: 0.1 * float64(qty),
			})

			database.DB.Create(&models.Order{
				UserId:          uint(rand.Intn(98) + 10),
				Code:            faker.Username(),
				AmbassadorEmail: faker.Email(),
				FirstName:       faker.FirstName(),
				Lastname:        faker.LastName(),
				Email:           faker.Email(),
				Complete:        true,
				OrderItems:      orderItems,
			})
		}

	}

}
