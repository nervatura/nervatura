package utils

import (
	"encoding/json"
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	cu "github.com/nervatura/component/pkg/util"
	st "github.com/nervatura/nervatura/v6/static"
)

var TokenAlg cu.SM = cu.SM{
	"RS256": "RSA", "RS384": "RSA", "RS512": "RSA",
	"ES256": "ECDSA", "ES384": "ECDSA", "ES512": "ECDSA",
	//"EdDSA": "EdDSA",
	"PS256": "PSS", "PS384": "PSS", "PS512": "PSS",
	"HS256": "HMAC", "HS384": "HMAC", "HS512": "HMAC",
}

/*
CreateLoginToken - create/refresh a login JWT token
*/
func CreateLoginToken(params cu.SM, config cu.IM) (result string, err error) {
	if cu.ToString(params["user_name"], "") == "" || cu.ToString(params["code"], "") == "" || cu.ToString(params["alias"], "") == "" {
		return result, errors.New("missing fieldname: username, code or alias")
	}
	var claims = cu.IM{
		"user_name": params["user_name"],
		"alias":     params["alias"],
		"scope":     cu.ToString(params["scope"], "GROUP_GUEST"),
		"version":   config["version"],
	}
	return CreateToken(cu.ToString(params["code"], ""), claims, config)
}

/*
CreateToken - create/refresh a JWT token
*/
func CreateToken(subject string, claims cu.IM, config cu.IM) (result string, err error) {
	expirationTime := time.Now().Add(time.Duration(cu.ToFloat(config["NT_TOKEN_EXP"], 1)) * time.Hour)
	claims["exp"] = jwt.NewNumericDate(expirationTime)
	claims["iss"] = cu.ToString(config["NT_TOKEN_ISS"], cu.ToString(st.DefaultConfig["token"]["iss"], "nervatura"))
	claims["sub"] = subject
	alg := cu.ToString(config["NT_TOKEN_ALG"], "HS256")
	if _, found := TokenAlg[alg]; !found {
		return "", errors.New("Unexpected signing method: " + alg)
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(alg), jwt.MapClaims(claims))
	token.Header["kid"] = cu.ToString(config["NT_TOKEN_PRIVATE_KID"], GetHash("nervatura", "sha256"))
	var key interface{} = []byte(cu.ToString(config["NT_TOKEN_PRIVATE_KEY"], ""))
	key, err = parsePEM(TokenAlg[alg], "private", key.([]byte))
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

/*
TokenDecode - decoded JWT token but doesn't validate the signature.
*/
func TokenDecode(tokenString string) (data cu.IM, err error) {
	var token *jwt.Token
	if token, _, err = new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{}); err == nil {
		data = token.Claims.(jwt.MapClaims)
		data["header"] = token.Header
	}
	return data, err
}

func parsePEM(method, stype string, value []byte) (interface{}, error) {
	if method == "RSA" && stype == "private" {
		return jwt.ParseRSAPrivateKeyFromPEM(value)
	}
	if method == "ECDSA" && stype == "private" {
		return jwt.ParseECPrivateKeyFromPEM(value)
	}
	if method == "EdDSA" && stype == "private" {
		return jwt.ParseEdPrivateKeyFromPEM(value)
	}
	if method == "RSA" && stype == "public" {
		return jwt.ParseRSAPublicKeyFromPEM(value)
	}
	if method == "ECDSA" && stype == "public" {
		return jwt.ParseECPublicKeyFromPEM(value)
	}
	if method == "EdDSA" && stype == "public" {
		return jwt.ParseEdPublicKeyFromPEM(value)
	}
	return value, nil
}

func validKeys(tokenString string, keyMap []map[string]string) (keys map[string]any, err error) {
	keys = map[string]any{}
	var token map[string]any
	if token, err = TokenDecode(tokenString); err != nil {
		return keys, err
	}
	tokenHeader := cu.ToIM(token["header"], map[string]any{})
	alg := cu.ToString(tokenHeader["alg"], "")
	kid := cu.ToString(tokenHeader["kid"], cu.RandString(16))
	var algType string
	var valid bool
	if algType, valid = TokenAlg[alg]; !valid {
		return keys, errors.New("Unexpected signing method: " + alg)
	}
	for _, tokenKey := range keyMap {
		isJWKS := (cu.ToString(tokenKey["type"], "") == "jwks")
		if algType == "HMAC" && !isJWKS {
			keys[cu.ToString(tokenKey["kid"], cu.RandString(16))] = []byte(tokenKey["value"])
		}
		if slices.Contains([]string{"RSA", "ECDSA", "EdDSA"}, algType) && !isJWKS {
			if tokenData, err := parsePEM(algType, tokenKey["type"], []byte(tokenKey["value"])); err == nil {
				keys[cu.ToString(tokenKey["kid"], cu.RandString(16))] = tokenData
			}
		}
		if isJWKS {
			jwksJSON := json.RawMessage(tokenKey["value"])
			if jwks, err := keyfunc.NewJWKSetJSON(jwksJSON); err == nil {
				keys["jwks"] = jwks
			}
		}
	}
	if _, found := keys[kid]; found {
		return map[string]any{kid: keys[kid]}, nil
	}
	if len(keys) == 0 {
		err = errors.New("no valid keys found")
	}
	return keys, err
}

/*
ParseToken - Parse, validate, and return a token data.
*/
func ParseToken(tokenString string, keyMap []cu.SM, config cu.IM) (data cu.IM, err error) {
	data = make(cu.IM)
	var keys map[string]any
	var token *jwt.Token
	if keys, err = validKeys(tokenString, keyMap); err != nil {
		return data, err
	}
	parseToken := func() (tk *jwt.Token, err error) {
		for kid, key := range keys {
			keyFunc := func(token *jwt.Token) (any, error) {
				return key, nil
			}
			if kid == "jwks" {
				keyFunc = key.(keyfunc.Keyfunc).Keyfunc
			}
			if tk, err = jwt.Parse(tokenString, keyFunc); err == nil {
				return tk, nil
			}
		}
		return tk, err
	}
	token, err = parseToken()
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return data, errors.New("token is either expired or not active yet")
		}
		return data, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userField := cu.ToString(config["NT_TOKEN_USER"], "user_name")
		data["user_code"] = cu.ToString(claims["sub"], "")
		data["user_name"] = cu.ToString(claims[userField], cu.ToString(claims["username"], cu.ToString(claims["preferred_username"], "")))
		data["scopes"] = strings.Split(cu.ToString(claims["scope"], ""), " ")
		data["alias"] = cu.ToString(claims["alias"], "")
		data["email"] = cu.ToString(claims["email"], "")
		data["email_verified"] = cu.ToBoolean(claims["email_verified"], false)
		data["picture"] = cu.ToString(claims["picture"], "")
		data["exp"] = claims["exp"]
	}

	return data, err
}
