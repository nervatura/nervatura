package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"image/color"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//go:embed static/views static/message.json
var Static embed.FS

//go:embed static/templates static/fonts
var Report embed.FS

func GetHash(text string) string {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return text
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

//ToString - safe string conversion
func ToString(value interface{}, defValue string) string {
	if stringValue, valid := value.(string); valid {
		if stringValue == "" {
			return defValue
		}
		return stringValue
	}
	if boolValue, valid := value.(bool); valid {
		return strconv.FormatBool(boolValue)
	}
	if intValue, valid := value.(int); valid {
		return strconv.Itoa(intValue)
	}
	if intValue, valid := value.(int32); valid {
		return strconv.Itoa(int(intValue))
	}
	if intValue, valid := value.(int64); valid {
		return strconv.FormatInt(intValue, 10)
	}
	if floatValue, valid := value.(float32); valid {
		return strconv.FormatFloat(float64(floatValue), 'f', -1, 64)
	}
	if floatValue, valid := value.(float64); valid {
		return strconv.FormatFloat(floatValue, 'f', -1, 64)
	}
	if timeValue, valid := value.(time.Time); valid {
		return timeValue.Format("2006-01-02T15:04:05-07:00")
	}
	return defValue
}

//ToFloat - safe float64 conversion
func ToFloat(value interface{}, defValue float64) float64 {
	if floatValue, valid := value.(float64); valid {
		if floatValue == 0 {
			return defValue
		}
		return floatValue
	}
	if boolValue, valid := value.(bool); valid {
		if boolValue {
			return 1
		}
	}
	if intValue, valid := value.(int); valid {
		return float64(intValue)
	}
	if intValue, valid := value.(int32); valid {
		return float64(intValue)
	}
	if intValue, valid := value.(int64); valid {
		return float64(intValue)
	}
	if floatValue, valid := value.(float32); valid {
		return float64(floatValue)
	}
	if stringValue, valid := value.(string); valid {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err == nil {
			return float64(floatValue)
		}
	}
	return defValue
}

//ToRGBA - safe RGBA conversion
func ToRGBA(value interface{}, defValue color.RGBA) color.RGBA {
	parseHexColor := func(v string) (out color.RGBA, err error) {
		if len(v) != 7 {
			return out, errors.New("hex color must be 7 characters")
		}
		red, redError := strconv.ParseUint(v[1:3], 16, 8)
		if redError != nil {
			return out, errors.New("red component invalid")
		}
		out.R = uint8(red)
		green, greenError := strconv.ParseUint(v[3:5], 16, 8)
		if greenError != nil {
			return out, errors.New("green component invalid")
		}
		out.G = uint8(green)
		blue, blueError := strconv.ParseUint(v[5:7], 16, 8)
		if blueError != nil {
			return out, errors.New("blue component invalid")
		}
		out.B = uint8(blue)
		return
	}

	if rgbaValue, valid := value.(color.RGBA); valid {
		return rgbaValue
	}
	if stringValue, valid := value.(string); valid {
		if strings.HasPrefix(stringValue, "#") {
			pvalue, err := parseHexColor(value.(string))
			if err == nil {
				return pvalue
			}
		} else {
			ivalue := ToInteger(value, -1)
			if ivalue > -1 && ivalue < 255 {
				return color.RGBA{uint8(ivalue), uint8(ivalue), uint8(ivalue), 0}
			}
		}
	}
	if intValue, valid := value.(int); valid {
		if intValue < 255 {
			return color.RGBA{uint8(intValue), uint8(intValue), uint8(intValue), 0}
		}
	}
	if int32Value, valid := value.(int32); valid {
		if int32Value < 255 {
			return color.RGBA{uint8(int32Value), uint8(int32Value), uint8(int32Value), 0}
		}
	}
	if int64Value, valid := value.(int64); valid {
		if int64Value < 255 {
			return color.RGBA{uint8(int64Value), uint8(int64Value), uint8(int64Value), 0}
		}
	}
	if float32Value, valid := value.(float32); valid {
		if float32Value < 255 {
			return color.RGBA{uint8(float32Value), uint8(float32Value), uint8(float32Value), 0}
		}
	}
	if float64Value, valid := value.(float64); valid {
		if float64Value < 255 {
			return color.RGBA{uint8(float64Value), uint8(float64Value), uint8(float64Value), 0}
		}
	}
	return defValue
}

//ToInteger - safe int64 conversion
func ToInteger(value interface{}, defValue int64) int64 {
	if intValue, valid := value.(int64); valid {
		if intValue == 0 {
			return defValue
		}
		return intValue
	}
	if boolValue, valid := value.(bool); valid {
		if boolValue {
			return 1
		}
	}
	if intValue, valid := value.(int); valid {
		return int64(intValue)
	}
	if intValue, valid := value.(int32); valid {
		return int64(intValue)
	}
	if floatValue, valid := value.(float32); valid {
		return int64(floatValue)
	}
	if floatValue, valid := value.(float64); valid {
		return int64(floatValue)
	}
	if stringValue, valid := value.(string); valid {
		intValue, err := strconv.ParseInt(stringValue, 10, 64)
		if err == nil {
			return int64(intValue)
		}
	}
	return defValue
}

//ToIntPointer - safe *int64 conversion
func ToIntPointer(value interface{}, defValue int64) *int64 {
	if value == nil {
		return nil
	}
	v := ToInteger(value, defValue)
	return &v
}

//ToStringPointer - safe *string conversion
func ToStringPointer(value interface{}, defValue string) *string {
	if value == nil {
		return nil
	}
	v := ToString(value, defValue)
	return &v
}

//ToBoolean - safe bool conversion
func ToBoolean(value interface{}, defValue bool) bool {
	if boolValue, valid := value.(bool); valid {
		return boolValue
	}
	if intValue, valid := value.(int); valid {
		if intValue == 1 {
			return true
		}
	}
	if intValue, valid := value.(int32); valid {
		if intValue == 1 {
			return true
		}
	}
	if intValue, valid := value.(int64); valid {
		if intValue == 1 {
			return true
		}
	}
	if floatValue, valid := value.(float32); valid {
		if floatValue == 1 {
			return true
		}
	}
	if floatValue, valid := value.(float64); valid {
		if floatValue == 1 {
			return true
		}
	}
	if stringValue, valid := value.(string); valid {
		boolValue, err := strconv.ParseBool(stringValue)
		if err == nil {
			return boolValue
		}
	}
	return defValue
}

//StringToDateTime - parse string to datetime
func StringToDateTime(value string) (time.Time, error) {
	tm, err := time.Parse("2006-01-02T15:04:05-07:00", value)
	if err != nil {
		tm, err = time.Parse("2006-01-02T15:04:05-0700", value)
	}
	if err != nil {
		tm, err = time.Parse("2006-01-02T15:04:05", value)
	}
	if err != nil {
		tm, err = time.Parse("2006-01-02T15:04:05Z", value)
	}
	if err != nil {
		tm, err = time.Parse("2006-01-02 15:04:05", value)
	}
	if err != nil {
		tm, err = time.Parse("2006-01-02 15:04", value)
	}
	if err != nil {
		tm, err = time.Parse("2006-01-02", value)
	}
	return tm, err
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ConvertToByte(data interface{}) ([]byte, error) {
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(data)
}

func ConvertFromByte(data []byte, result interface{}) error {
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, result)
}

func ConvertFromReader(data io.Reader, result interface{}) error {
	return json.NewDecoder(data).Decode(&result)
}

//GetMessage - error messages
func GetMessage(key string) string {
	var messages map[string]string
	var jsonMessages, _ = Static.ReadFile("static/message.json")
	if err := ConvertFromByte(jsonMessages, &messages); err != nil {
		return ""
	}
	if value, found := messages[key]; found {
		return value
	}
	return ""
}

/*
CreateToken - create/refresh a Nervatura JWT token
*/
func CreateToken(username, database string, config map[string]interface{}) (string, error) {
	// ntClaims is a custom Nervatura claims type
	type ntClaims struct {
		Username string `json:"username"`
		Database string `json:"database"`
		jwt.RegisteredClaims
	}

	expirationTime := time.Now().Add(time.Duration(ToFloat(config["NT_TOKEN_EXP"], 1)) * time.Hour)
	claims := ntClaims{
		username,
		database,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    ToString(config["NT_TOKEN_ISS"], "nervatura"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = ToString(config["NT_TOKEN_PRIVATE_KID"], GetHash("nervatura"))
	return token.SignedString([]byte(ToString(config["NT_TOKEN_PRIVATE_KEY"], GetHash(time.Now().Format("20060102")))))
}

/*
TokenDecode - decoded JWT token but doesn't validate the signature.
*/
func TokenDecode(tokenString string) (map[string]interface{}, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err == nil {
		return token.Claims.(jwt.MapClaims), err
	}
	return nil, err
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

/*
ParseToken - Parse, validate, and return a token data.
*/
func ParseToken(tokenString string, keyMap map[string]map[string]string, config map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		algType := ""
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			// RS256, RS384, RS512
			algType = "RSA"
		}
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
			// ES256, ES384, ES512
			algType = "ECDSA"
		}
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); ok {
			// EdDSA
			algType = "EdDSA"
		}
		if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); ok {
			// PS256, PS384, PS512
			algType = "PSS"
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			// HS256, HS384, HS512
			algType = "HMAC"
		}
		if algType == "" {
			return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
		}
		kid := ToString(token.Header["kid"], ToString(config["NT_TOKEN_PRIVATE_KID"], GetHash("nervatura")))
		if keyMap, found := keyMap[kid]; found {
			return parsePEM(algType, keyMap["type"], []byte(keyMap["value"]))
		}
		return []byte(ToString(config["NT_TOKEN_PRIVATE_KEY"], GetHash(time.Now().Format("20060102")))), nil
	})
	if err != nil {
		return data, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return data, errors.New("token is either expired or not active yet")
			}
		}
		return data, err
	}

	data["database"] = claims["database"]
	if _, found := claims["database"]; !found {
		if ToString(config["NT_ALIAS_DEFAULT"], "") == "" {
			return data, errors.New(GetMessage("missing_database"))
		}
		data["database"] = ToString(config["NT_ALIAS_DEFAULT"], "")
	}
	data["username"] = ""
	if _, found := claims["username"]; found {
		data["username"] = claims["username"]
	} else if _, found := claims["user_id"]; found {
		data["username"] = claims["user_id"]
	} else if _, found := claims["sub"]; found {
		data["username"] = claims["sub"]
	} else if _, found := claims["email"]; found {
		data["username"] = claims["email"]
	}
	if data["username"] == "" {
		return data, errors.New(GetMessage("missing_user"))
	}
	return data, nil

}

func RandString(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			n = big.NewInt(int64(i))
		}
		b.WriteRune(chars[n.Int64()])
	}
	return b.String()
}
