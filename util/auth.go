package util

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/util/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func VerifyRoleAndID(c *gin.Context, expectedRole string, expectedID int64) bool {
	role, roleExists := c.Get("role")
	id, idExists := c.Get(expectedRole + "_id")

	log.Info(role, roleExists)
	log.Info(id, idExists)

	if !roleExists || !idExists || role.(string) != expectedRole {
		sentry.CaptureMessage("Unauthorized, jwt:" + fmt.Sprintf("%+v", id) + "\nacc:" + fmt.Sprintf("%+v", expectedID))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return false
	}

	if int64(id.(float64)) != expectedID {
		sentry.CaptureMessage("Forbidden, jwt:" + fmt.Sprintf("%+v", id) + "\nacc:" + fmt.Sprintf("%+v", expectedID))
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden " + role.(string)})
		return false
	}

	return true
}

func VerifyID(c *gin.Context, expectedRole string, expectedID int64) bool {
	role, roleExists := c.Get("role")
	id, idExists := c.Get(expectedRole + "_id")

	log.Info(role, roleExists)
	log.Info(id, idExists)

	if !roleExists || !idExists {
		sentry.CaptureMessage("Unauthorized, jwt:" + fmt.Sprintf("%+v", id) + "\nacc:" + fmt.Sprintf("%+v", expectedID))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return false
	}

	if int64(id.(float64)) != expectedID {
		sentry.CaptureMessage("Forbidden, jwt:" + fmt.Sprintf("%+v", id) + "\nacc:" + fmt.Sprintf("%+v", expectedID))
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden " + role.(string)})
		return false
	}

	return true
}

func RetrieveCustomerId(c *gin.Context) (int64, error) {
	role, roleExists := c.Get("role")

	if !roleExists || role == nil {
		sentry.CaptureMessage("Unauthorized, jwt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return 0, errors.New("unauthorized access")
	}
	id, idExists := c.Get("customer_id")

	if !idExists || id == nil {
		return 0, nil
	}
	return int64(id.(float64)), nil
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
