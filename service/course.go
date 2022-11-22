package service

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (ac *Service) InsertValuesToCA(cv *models.CourseInfo) error {

	if cv.CourseName == "" {
		return fmt.Errorf("course name should not be empty")
	}

	cv.Id = uuid.New()

	ok := ac.daos.CheckCourse(cv.CourseName)
	if ok {
		return fmt.Errorf("course Name exits")
	}
	status := ac.daos.InsertValuesToCoursesAvailable(cv)
	if status != nil {
		return status
	}
	return status

}

func (ac *Service) RetrieveCA() ([]*models.CourseInfo, error) {
	rca, err := ac.daos.RetieveCoursesAvailable()
	return rca, err
}

func (ac *Service) UpdateCA(name string, rc *models.CourseInfo) error {

	err := ac.daos.UpdateCourseByName(name, rc)
	if err != nil {
		return fmt.Errorf("not able to update %s", err.Error())
	}
	return nil
}

func (ac *Service) DeleteCA(name string) error {

	status := ac.daos.CheckCourse(name)
	if !status {
		return fmt.Errorf("no course Found")
	}
	rc, _ := ac.daos.GetCourseByName(name)

	ok, err := ac.daos.DeleteCourse(rc.Id)
	if err != nil {
		return fmt.Errorf("not able to Delete %s", err.Error())
	}
	if ok {
		return nil
	} else {
		return fmt.Errorf("some Error happend")
	}

}