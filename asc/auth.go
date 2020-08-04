package asc

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// AuthTransport is an http.RoundTripper implementation that stores the JWT created.
// If the token expires, the Rotate function should be called to update the stored token.
type AuthTransport struct {
	Transport    http.RoundTripper
	jwtGenerator jwtGenerator
}

type jwtGenerator interface {
	Token() (string, error)
	IsExpired() bool
}

type standardJWTGenerator struct {
	keyID          string
	issuerID       string
	expireDuration time.Duration
	privateKey     interface{}

	token string
}

// NewTokenConfig returns a new AuthTransport instance that customizes the Authentication header of the request during transport.
// It can be customized further by supplying a custom http.RoundTripper instance to the Transport field.
func NewTokenConfig(keyID string, issuerID string, expireDuration time.Duration, privateKey []byte) (*AuthTransport, error) {
	key, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	gen := &standardJWTGenerator{
		keyID:          keyID,
		issuerID:       issuerID,
		privateKey:     key,
		expireDuration: expireDuration,
	}
	_, err = gen.Token()
	return &AuthTransport{
		jwtGenerator: gen,
	}, err
}

func parsePrivateKey(blob []byte) (key interface{}, err error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, fmt.Errorf("no PEM blob found")
	}
	key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// RoundTrip implements the http.RoundTripper interface to set the Authorization header
func (t AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.jwtGenerator.Token()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return t.transport().RoundTrip(req)
}

// Client returns a new http.Client instance for use with asc.Client.
func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (g *standardJWTGenerator) Token() (string, error) {
	if !g.IsExpired() {
		return g.token, nil
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, g.claims())
	t.Header["kid"] = g.keyID
	token, err := t.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}
	g.token = token
	return token, nil
}

func (g *standardJWTGenerator) IsExpired() bool {
	if g.token == "" {
		return true
	}
	// TODO: Check if token is expired
	return false
}

func (g *standardJWTGenerator) claims() jwt.Claims {
	expiry := time.Now().Add(g.expireDuration)
	return jwt.StandardClaims{
		Audience:  jwt.ClaimStrings{"appstoreconnect-v1"},
		Issuer:    g.issuerID,
		ExpiresAt: jwt.At(expiry),
	}
}
