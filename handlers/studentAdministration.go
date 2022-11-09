package handlers

import (
	"CollegeAdministration/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertCAd(ctx *gin.Context) {
	var cad models.StudentInfo
	err := ctx.BindJSON(&cad)
	if err != nil {
		log.Println("not able to store values")
	}
	response := h.service.InsertValuesToCAd(&cad)
	if response != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error())
	} else {
		ctx.JSON(http.StatusCreated, "Successfully inserted to table")
	}
}

func (h *Handler) RetrieveValuesCAd(ctx *gin.Context) {

	response, err := h.service.RetrieveCAd()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (h *Handler) UpdateValuesCAd(ctx *gin.Context) {

	oldCourse :=ctx.Param("coursename")
	var rcd models.StudentInfo
	ctx.BindJSON(&rcd)
	err := h.service.UpdateCAd(&rcd,oldCourse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}

func (h *Handler) DeleteSA(ctx *gin.Context) {

	rollNumber := ctx.Param("rollnumber")

	err := h.service.DeleteStudent(rollNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}
func (h *Handler) UpdateStudentNameAndAge(ctx *gin.Context) {
	var rcd *models.StudentInfo
	existing_name := ctx.Param("name")
	ctx.BindJSON(&rcd)
	err := h.service.UpdateStudentNameAge(existing_name, rcd.Name, rcd.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}

func (h *Handler) FetchAllCourseForAStudent(ctx *gin.Context) {

	student_name := ctx.Param("name")
	res, err := h.service.FetchStudentCourse(student_name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, res)
	}

}
