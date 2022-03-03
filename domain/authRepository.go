package domain

import (
	"encoding/json"
	"github.com/Dontunee/banking/logger"
	"net/http"
	"net/url"
)

type IAuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type AuthRepository struct {
}

func (authRepository AuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	verifyUrl := buildVerifyUrl(token, routeName, vars)

	if response, err := http.Get(verifyUrl); err != nil {
		logger.Error("Error while sending:..." + err.Error())
		return false
	} else {
		dictionary := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&dictionary); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return dictionary["isAuthorized"]
	}
}

//Sample :/auth/verify?token=aaaa.bbbb.cccc&routeName=create&customer_id=2000&account_id=95470
func buildVerifyUrl(token string, routeName string, vars map[string]string) string {
	url := url.URL{Scheme: "http", Host: "localhost:8181", Path: "/auth/verify"}

	query := url.Query()
	query.Add("token", token)
	query.Add("routeName", routeName)

	for k, v := range vars {
		query.Add(k, v)
	}

	url.RawQuery = query.Encode()
	return url.String()
}

func NewAuthRepository() AuthRepositoryStub {
	return AuthRepositoryStub{}
}

type AuthRepositoryStub struct {
}

func (authRepository AuthRepositoryStub) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	return true
}

