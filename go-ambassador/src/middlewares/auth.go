package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("JWT")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unautenticated",
		})
	}

	return c.Next()

}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("JWT")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*jwt.StandardClaims)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}
