package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/InsertToCoursesAvailable", h.InsertCA)
	rc.POST("/InsertToCollegeAdministration", h.InsertCAd)
	rc.GET("/RetrieveCoursesAvailable", h.RetrieveValuesCA)
	rc.GET("/RetrieveCollegeAdministration", h.RetrieveValuesCAd)
	rc.PATCH("/UpdateCoursesAvailable/:name", h.UpdateValuesCA)
	rc.PATCH("/UpdateCollegeAdministration/:coursename", h.UpdateValuesCAd)
	rc.DELETE("/DeleteCoursesAvailable/:courseName", h.DeleteCA)
	rc.DELETE("DeleteStudentInfo/:rollnumber", h.DeleteSA)
	rc.PATCH("UpdateStudentNameAndAge/:name", h.UpdateStudentNameAndAge)
	rc.GET("FindAllCourseForAStudent/:name", h.FetchAllCourseForAStudent)
	rc.POST("/InsertInstructorDetails",h.InstructorInfoHandlers)
	rc.GET("/RetrieveInstructors",h.RetrieveInstructorDetails)
	rc.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

/* 
API TO CREATE -
 Delete a course for a student
 delete student-marks
 
*/
