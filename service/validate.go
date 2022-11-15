package service

import (
	"CollegeAdministration/models"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
)

func (ac *Service) ValidateLogin(email, password string) error {

	ok, _ := regexp.MatchString("@gmail.com", email)
	if !ok || len(email) < 11 {
		return fmt.Errorf("wrong email format")
	} else if password == ":password" {
		return fmt.Errorf("password cant be empty ")
	} else if len(password) < 8 {
		return fmt.Errorf("password length insufficient")
	}
	return nil
}
func (ac *Service) CheckCredentials(email, password string) error {

	exits, err := ac.daos.CheckLoginExits(email, password)
	if err != nil {
		return err
	}
	if !exits {
		return fmt.Errorf("email or password wrong! Re Enter")
	}
	return nil
}

func (ac *Service) GetTokenAfterLogging() (uuid.UUID, error) {
	var token_table models.Token_generator

	token := uuid.New()

	token_table.Token = token
	token_table.IsValid = true
	token_table.ValidFrom = time.Now()
	token_table.ValidTill = token_table.ValidFrom.Add(time.Minute * 5)
	log.Println(token_table)

	err := ac.daos.InsertToken(token_table)
	if err != nil {
		return uuid.Nil, fmt.Errorf("token insertion failed")
	}
	return token_table.Token, nil
}

func (ac *Service) CheckTokenValidity(token uuid.UUID) (bool, error) {

	err1 := ac.CheckTokenExpiry(token)
	if err1 != nil {
		return false, err1
	}
	status, err := ac.daos.GetTokenStatus(token)
	if err != nil {
		return status, err
	}
	return status, nil
}

func (ac *Service) CheckTokenExpiry(token uuid.UUID) error {

	token_details, err := ac.daos.GetTokenStored(token)
	if err != nil {
		return err
	}
	time_now := time.Now()
	if token_details.ValidTill.Minute() < time_now.Minute() || token_details.ValidTill.Hour() < time_now.Hour() {
		err2 := ac.daos.SetTokenFalse(token)
		if err2 != nil {
			return err2
		}
		return fmt.Errorf("token expired! Generate new token")
	}
	return nil
}