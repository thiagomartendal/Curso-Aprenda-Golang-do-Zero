package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(ID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["idUsuario"] = ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return err
	}

	if _, OK := token.Claims.(jwt.MapClaims); OK && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func ExtrairID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return 0, err
	}

	if permissoes, OK := token.Claims.(jwt.MapClaims); OK && token.Valid {
		ID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["idUsuario"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return ID, nil
	}

	return 0, errors.New("token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
		return nil, fmt.Errorf("método de assinatura inesperado. %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
