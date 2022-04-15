package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

type PostUserMajorParams struct {
	UserID     int64  `json:"userID"`
	MajorName  string `json:"majorName"`
	Emphasis   string `json:"emphasisName"`
	SchoolYear string `json:"schoolYear"`
}

type HistoryParam struct {
	CourseID int64 `json:"courseID"`
	UserID   int64 `json:"userID"`
}

// RegisterHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var params model.RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	if err := logic.Register(&params); err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

// LoginHandler User login interface
// @Summary Use this API to login
// @Description After login, a token will be returned to verify user in the future
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.LoginParams true "login parameters"
// @Success 200 {string} jwt token
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var params model.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	userInfo, err := logic.Login(&params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: userInfo,
	})
}

// GetUserHistoryHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
func GetUserHistoryHandler(c *gin.Context) {
	userID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	history, err := logic.GetUserHistoryCourseList(int64(userIDInt))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: history,
	})
}

// AddUserHistoryHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
func AddUserHistoryHandler(c *gin.Context) {
	var param HistoryParam
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.AddUserCourseHistory(param.UserID, param.CourseID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

// RemoveUserHistoryHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
func RemoveUserHistoryHandler(c *gin.Context) {
	var param HistoryParam
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.RemoveUserCourseHistory(param.UserID, param.CourseID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func PostUserMajorHandler(c *gin.Context) {
	var param PostUserMajorParams
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	user, err := logic.PostUserMajor(param.UserID, param.MajorName, param.Emphasis, param.SchoolYear)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    user,
	})
}

func GetUserRecommendCourseHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{
			Msg: "Unable to process the given user ID",
		})
	}

	recommendCourses, err := logic.RecommendCourses(int64(uid))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: recommendCourses,
	})
}
