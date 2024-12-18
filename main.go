package main

// https://github.com/MicahParks/keyfunc
// The purpose of this package is to provide a jwt.Keyfunc for the
// github.com/golang-jwt/jwt/v5 package using a JSON Web Key Set (JWK Set)
// for parsing and verifying JSON Web Tokens (JWTs).

// The purpose of this package is to provide a jwt.Keyfunc for the
// github.com/golang-jwt/jwt/v5 package using a JSON Web Key Set (JWK Set)
// for parsing and verifying JSON Web Tokens (JWTs).

// It's common for an identity providers, particularly those using OAuth 2.0 or
// OpenID Connect, such as Keycloak or Amazon Cognito (AWS) to expose a JWK Set via
// an HTTPS endpoint. This package has the ability to consume that JWK Set and produce
// a jwt.Keyfunc. It is important that a JWK Set endpoint is using HTTPS to ensure
// the keys are from the correct trusted source

// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Keyfunc
// https://github.com/MicahParks/keyfunc
// Keyfunc will be used by the Parse methods as a callback function to supply the key
// for verification. The function receives the parsed, but unverified Token. This allows
// you to use properties in the Header of the token (such as `kid`) to identify which key
// to use.

// The returned interface{} may be a single key or a VerificationKeySet containing
// multiple keys.

// https://github.com/MicahParks/keyfunc/blob/main/examples/http/main.go
// This is an example that creates a keyset server and uses it to parse a JWT token

// https://github.com/MicahParks/jwkset
// This library can be used to create a keyset server

import (
	"context"
	"log"

	"github.com/golang-jwt/jwt/v5"

	"github.com/MicahParks/keyfunc/v3"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jwksURL := "https://raw.githubusercontent.com/istio/istio/release-1.24/security/tools/jwt/samples/jwks.json"

	// Create the keyfunc.Keyfunc.
	// k, err := keyfunc.NewDefault([]string{jwksURL})
	// if err != nil {
	// 	log.Fatalf("Failed to create a keyfunc.Keyfunc from the server's URL.\nError: %s", err)
	// }

	// 1. Create the keyfunc.Keyfunc.
	k, err := keyfunc.NewDefaultCtx(ctx, []string{jwksURL}) // Context is used to end the refresh goroutine.
	if err != nil {
		log.Fatalf("Failed to create a keyfunc.Keyfunc from the server's URL.\nError: %s", err)
	}

	// When using the keyfunc.NewDefault function, the JWK Set will be automatically
	// refreshed using jwkset.NewDefaultHTTPClient. This does launch a " refresh goroutine".
	//  If you want the ability to end this goroutine, use the keyfunc.NewDefaultCtx function.

	// In Go, you cannot directly hide a struct field in the sense of making it completely
	// inaccessible from outside the package. However, you can achieve a similar effect by
	// using the following techniques:
	// 1. Unexported Fields
	// By convention, fields starting with a lowercase letter are unexported, meaning they
	// are only accessible within the same package.

	// Since version 3.X.X, this project has become a thin wrapper around github.com/MicahParks/jwkset.
	// Newer versions contain a superset of features available in versions 2.X.X and earlier, but some
	// of the deep customization has been moved to the jwkset project. The intention behind this is to
	// make keyfunc easier to use for most use cases.

	log.Println(k.Storage().JSON(ctx))
	log.Println(k.Storage().JSONPublic(ctx))

	// Get a JWT to parse.
	// jwtB64 := "eyJraWQiOiJlZThkNjI2ZCIsInR5cCI6IkpXVCIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJXZWlkb25nIiwiYXVkIjoiVGFzaHVhbiIsImlzcyI6Imp3a3Mtc2VydmljZS5hcHBzcG90LmNvbSIsImlhdCI6MTYzMTM2OTk1NSwianRpIjoiNDY2M2E5MTAtZWU2MC00NzcwLTgxNjktY2I3NDdiMDljZjU0In0.LwD65d5h6U_2Xco81EClMa_1WIW4xXZl8o4b7WzY_7OgPD2tNlByxvGDzP7bKYA9Gj--1mi4Q4li4CAnKJkaHRYB17baC0H5P9lKMPuA6AnChTzLafY6yf-YadA7DmakCtIl7FNcFQQL2DXmh6gS9J6TluFoCIXj83MqETbDWpL28o3XAD_05UP8VLQzH2XzyqWKi97mOuvz-GsDp9mhBYQUgN3csNXt2v2l-bUPWe19SftNej0cxddyGu06tXUtaS6K0oe0TTbaqc3hmfEiu5G0J8U6ztTUMwXkBvaknE640NPgMQJqBaey0E4u0txYgyvMvvxfwtcOrDRYqYPBnA"
	jwtB64 := "eyJhbGciOiJSUzI1NiIsImtpZCI6IkRIRmJwb0lVcXJZOHQyenBBMnFYZkNtcjVWTzVaRXI0UnpIVV8tZW52dlEiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ2ODU5ODk3MDAsImZvbyI6ImJhciIsImlhdCI6MTUzMjM4OTcwMCwiaXNzIjoidGVzdGluZ0BzZWN1cmUuaXN0aW8uaW8iLCJzdWIiOiJ0ZXN0aW5nQHNlY3VyZS5pc3Rpby5pbyJ9.CfNnxWP2tcnR9q0vxyxweaF3ovQYHYZl82hAUsn21bwQd9zP7c-LS9qd_vpdLG4Tn1A15NxfCjp5f7QNBUo-KC9PJqYpgGbaXhaGx7bEdFWjcwv3nZzvc7M__ZpaCERdwU7igUmJqYGBYQ51vr2njU9ZimyKkfDe3axcyiBZde7G6dabliUosJvvKOPcKIWPccCgefSj_GNfwIip3-SsFdlR7BtbVUcqR-yv-XOxJ3Uc1MI0tz3uMiiZcyPV7sNCU4KRnemRIMHVOfuvHsU60_GhGbiSFzgPTAa9WTltbnarTbxudb_YEOx12JiwYToeX0DCPb43W1tzIBxgm8NxUg"

	// Parse the JWT.
	token, err := jwt.Parse(jwtB64, k.Keyfunc)
	// Parse parses, validates, verifies the signature and returns the parsed token.
	// keyFunc will receive the parsed token and should return the cryptographic key
	// for verifying the signature. The caller is strongly encouraged to set the
	// WithValidMethods option to validate the 'alg' claim in the token matches the
	// expected algorithm. For more details about the importance of validating the 'alg'
	// claim, see https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/

	if err != nil {
		log.Fatalf("Failed to parse the JWT.\nError: %s", err.Error())
	}

	// Check if the token is valid.
	if !token.Valid {
		log.Fatalf("The token is not valid.")
	}
	log.Println("The token is valid.")
}
