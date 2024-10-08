package jobs

import (
	"CollegeAdministration/config"
	"CollegeAdministration/daos"
	"CollegeAdministration/models"
	"CollegeAdministration/service"
	"fmt"
)

func AccountDetailsMigration() error {

	var complete_account []*models.Account
	s := service.New(config.DBInit())
	daos := daos.New(config.DBInit())
	student_details, err := s.Retrieve_student_details()
	if err != nil {
		err_statement := "Failed to get details" + err.Error()
		return fmt.Errorf(err_statement)
	}
	for _, student_detail := range student_details {
		account_dataset := &models.Account{}
		account_dataset.Id = student_detail.Id
		account_dataset.Name = student_detail.Name

		complete_account = append(complete_account, account_dataset)
	}

	instructor_details, err2 := s.GetInstructorDetails()
	if err2 != nil {
		err_statement := "Failed getting credentials" + err2.Error()
		return fmt.Errorf(err_statement)
	}

	for _, instructor_detail := range instructor_details {
		account_dataset := &models.Account{}
		account_dataset.Id = instructor_detail.Id
		account_dataset.Name = instructor_detail.InstructorName

		credentials, err := daos.FetchCredentialsUsingID(instructor_detail.Id)
		if err != nil {
			err_statement := "Failed getting credentials" + err.Error()
			return fmt.Errorf(err_statement)
		}
		account_dataset.Info.Credentials.Id = credentials.Id
		account_dataset.Info.Credentials.EmailId = credentials.EmailId
		account_dataset.Info.Credentials.Password = credentials.Password

		complete_account = append(complete_account, account_dataset)

	}

	err1 := daos.AccountMigrationsUpdate(complete_account)
	if err1 != nil {
		err_statement := "Failed migration" + err1.Error()

		return fmt.Errorf(err_statement)
	}
	return nil
}
