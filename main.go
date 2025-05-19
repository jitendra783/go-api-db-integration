package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Secret key for signing and verifying the JWT (store it securely in a real application)
var secretKey = []byte("mysecretkey")

// Claims struct defines the JWT's payload
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(username string) (string, *Claims, error) {
	// Set the claims (payload) for the token
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(), // Token expires in 1 hour
		},
	}

	// Create the token with the HMAC signing method and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key and return it as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil, err
	}

	// Return the token string and the claims
	return tokenString, &claims, nil
}

// JWT Middleware to validate and parse the token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
			c.Abort()
			return
		}

		// The Authorization header has the format "Bearer <token>"
		tokenString := authHeader[7:] // Remove "Bearer " from the start of the string

		// Parse the token and validate its claims
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token is signed with HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil // Return the secret key for signature validation
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Successfully parsed the token, retrieve the claims
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not parse claims"})
			c.Abort()
			return
		}

		// Store the claims in the context (you can access this in protected routes)
		c.Set("username", claims.Username)

		// Proceed to the next handler
		c.Next()
	}
}

// Example of a protected route that requires a valid JWT
func ProtectedRoute(c *gin.Context) {
	// Access the username from the claims in the context
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Welcome, %s!", username),
	})
}

// Route to generate a new JWT for testing, now including the claims in the response
func GenerateTokenRoute(c *gin.Context) {
	// Generate a token for the user (this could come from a login system)
	username := c.DefaultQuery("username", "guest")

	// Generate the JWT token and claims
	token, claims, err := GenerateJWT(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Send the token and claims as a response
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token":  token,
			"exp": claims.ExpiresAt,
			"issue": claims.IssuedAt,
			"username": claims.Username,
			"id": claims.Id,
			 // Include the claims in the response
		}, // Include the claims in the response
	})
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Route to generate a token (you would use this route to get a JWT)
	r.GET("/generate-token", GenerateTokenRoute)

	// Applying JWTMiddleware before protected routes
	r.GET("/protected", JWTMiddleware(), ProtectedRoute)

	// Start the server
	r.Run(":8080")
}
