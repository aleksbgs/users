package middlewares

import (
	"github.com/aleksbgs/ambassador/src/database"
	"github.com/aleksbgs/users/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

const SecretKey = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	var userToken models.UserToken

	database.DB.Where("user_id = ? and token = ? and expired_at >= now()", id, token.Raw).Last(&userToken)

	if userToken.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	//payload := token.Claims.(*ClaimsWithScope)
	//isAmbassador := strings.Contains(c.Path(), "/api/ambassador")
	//
	//if (payload.Scope == "admin" && isAmbassador) || (payload.Scope == "ambassador" && !isAmbassador) {
	//	c.Status(fiber.StatusUnauthorized)
	//	return c.JSON(fiber.Map{
	//		"message": "unauthorized",
	//	})
	//}

	return c.Next()
}

func GenerateJWT(id uint, scope string) (string, error) {
	payload := ClaimsWithScope{}
	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	payload.Scope = scope

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	return uint(id), nil
}
