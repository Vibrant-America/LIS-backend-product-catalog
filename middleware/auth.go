package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	log "go-micro.dev/v4/logger"
	"net/http"
	"productCatalog/common"
	"strings"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 7 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, claims jwt.MapClaims) {
	c.Set("customer_id", claims["customer_id"])
	c.Set("clinic_id", claims["clinic_id"])
	c.Set("old_clinic_id", claims["old_clinic_id"])
	c.Set("patient_id", claims["patient_id"])
	c.Set("customer_list", claims["customer_list"])
	c.Set("userId", claims["userId"])
	c.Set("user_permission", claims["user_permission"])
	c.Set("role", claims["role"])
	c.Set("internal_user_id", claims["internal_user_id"])
	c.Set("internal_user_name", claims["internal_user_id"])
	c.Set("internal_user_role", claims["internal_user_id"])
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//
//	r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		//UpdateContextUserModel(c, 0)
		if strings.Contains(c.FullPath(), "healthcheck") {
			c.JSON(200, gin.H{
				"message": "pong",
			})
			c.Abort()
			return
		}
		if strings.Contains(c.FullPath(), "swagger") || strings.Contains(c.FullPath(), "webhook") {
			return
		}
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			log.Info("jwt", token)
			b := []byte(common.JWTSecret)
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			}
			return
		}
		//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Info(claims)
			//my_user_id := uint(claims["id"].(float64))
			//fmt.Println(my_user_id,claims["id"])
			//UpdateContextUserModel(c, my_user_id)
			UpdateContextUserModel(c, claims)
		}
		c.Next()
	}
}
