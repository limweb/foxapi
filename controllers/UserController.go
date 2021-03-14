package controllers

import (
	"apitest/db"
	"apitest/models"
	"apitest/services"
	"apitest/utils"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

//*** เพื่อให้ routes ต่อไปนี้ทำงานต้องนำ บันทึก routes ไปเพิ่มใน server.go
// controllers.SetupUserRoutes(v1)
func SetupUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/users")
	// router.Use(utils.JwtVerify) // jwt verify with secret key ถ้าต้องการให้ auth ด้วย jwt
	router.GET("/all", listUser)
	router.GET("/last", listlastUser)
	router.GET("/vuetable", vuetableUser)
	router.GET("/by/:id", getOneUser)
	router.GET("/del/:id", deleteUser)
	router.POST("/new", addNewUser)
	router.POST("/edit/:id", putOneUser)
}

// ListUsers godoc
// @summary List Users
// @description List all users
// @tags users
// @security ApiKeyAuth
// @id ListUsers
// @accept json
// @produce json
// @response 200 {array} []models.User "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/users/all [get]
func listUser(c *gin.Context) {
	var rs []models.User
	payload := utils.GetResponse()
	err := services.GetAllUser(&rs)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// ListUsers godoc
// @summary List Users
// @description List all users
// @tags users
// @security ApiKeyAuth
// @id ListlastUsers
// @accept json
// @produce json
// @response 200 {array} []models.User "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/users/last [get]
func listlastUser(c *gin.Context) {
	var rs []models.User
	payload := utils.GetResponse()
	err := services.GetAllIdDescUser(&rs)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// CreateUser godoc
// @summary Create User
// @description Create new user
// @tags users
// @security ApiKeyAuth
// @id CreateUser
// @accept json
// @produce json
// @param User body models.UserForCreate true "User data to be created"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/users/new [post]
func addNewUser(c *gin.Context) {
	var rs models.User
	payload := utils.GetResponse()
	c.BindJSON(&rs)
	err := services.AddNewUser(&rs)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// GetUser godoc
// @summary Get User
// @description  Get user by id
// @tags users
// @security Basic auth
// @id GetUser
// @accept json
// @produce json
// @param id path int true "id of user to be gotten"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/users/by/{id} [get]
func getOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	payload := utils.GetResponse()
	var rs models.User
	err := services.GetOneUser(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// UpdateUser godoc
// @summary Update User
// @description Update user by id
// @tags users
// @security ApiKeyAuth
// @id UpdateUser
// @accept json
// @produce json
// @param id path int true "id of user to be updated"
// @param User body models.UserForUpdate true "User data to be updated"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/users/edit/{id} [post]
func putOneUser(c *gin.Context) {
	var rs models.User
	payload := utils.GetResponse()
	id := c.Params.ByName("id")
	err := services.GetOneUser(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 200, payload)
	}
	c.BindJSON(&rs)
	err = services.PutOneUser(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 200, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// DeleteUser godoc
// @summary Delete User
// @description Delete user by id
// @tags users
// @security ApiKeyAuth
// @id DeleteUser
// @accept json
// @produce json
// @param id path int true "id of user to be deleted"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/users/del/{id} [get]
func deleteUser(c *gin.Context) {
	var rs models.User
	payload := utils.GetResponse()
	id := c.Params.ByName("id")
	err := services.DeleteUser(&rs, id)
	if err != nil {
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
		payload.Meta = http.StatusText(404)
		payload.Alert = alert
		utils.RespondJSON(c, 404, payload)
	} else {
		payload.Data = rs
		utils.RespondJSON(c, 200, payload)
	}
}

// VuetableUser godoc
// @summary List Users use VueTable
// @description List all users with pagination keyword filter by column Sort columns
// @tags users
// @security ApiKeyAuth
// @id vuetableUser
// @accept json
// @produce json
// @Param page query string false "number of page"
// @Param per_page query string false "get number of per_page"
// @Param sort query string false "soft by cols exm: col1|asc,col2|desc"
// @Param filter query string false "search with column exm: col1|aaa,col2|bbb  by filter"
// @Param kw query string false "search by kw exm: aaa bbb ccc "
// @response 200 {object} []models.User "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// example /api/v1/users/vuetable?page=2&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=name|aaa,author|bbb
// @Router /api/v1/users/vuetable [get]
func vuetableUser(c *gin.Context) {
	var rs []models.User
	query := c.Request.URL.Query()
	table := "users"
	paging := utils.GenPagination(query, table)
	dbqry := db.GetDB().Limit(paging.Limit).Offset(paging.Offset).Order(paging.Sort)

	qryfilter := utils.Filter(paging.Filter)
	for k, v := range qryfilter {
		log.Println("K->", k, "V->", v)
		dbqry = dbqry.Where(k+" like ? ", "%"+v+"%")
	}

	// ############################################################ start
	// columns ที่ต้องการให้ค้นหาด้วย keyword รวม
	// cols := []string{"name", "author", "category"} //ต้องกำหนดเอง example
	cols := []string{"name", "telephone"} //ต้องกำหนดเอง
	// keywork single search #1
	if paging.Kw != "" {
		for _, col := range cols {
			dbqry = dbqry.Or(col+" like ? ", "%"+paging.Kw+"%")
		}
	}
	// keywork multiple search #2
	// var kw = "aaa bbb cccc"
	rskws := utils.Searchkw(paging.Kw)
	log.Println("test kws-->", rskws)
	for _, kw := range rskws {
		for _, col := range cols {
			dbqry = dbqry.Or(col+" like ? ", "%"+kw+"%")
		}
	}
	// ############################################################ end

	dbqry.Find(&rs)
	if dbqry.Error != nil {
		log.Fatal(dbqry.Error)
	}
	var totalRows int
	errCount := db.GetDB().Model(&rs).Count(&totalRows).Error
	if errCount != nil {
		log.Fatal(errCount.Error)
	}

	payload := utils.GetResponse()
	var vtb utils.VueTableResponse
	vtb.Total = totalRows
	vtb.PerPage = paging.PerPage
	vtb.Datas = rs
	vtb.Form = paging.Offset + 1
	vtb.To = paging.Page * paging.Limit
	vtb.CurrentPage = paging.Page
	vtb.LastPage = int(math.Ceil(float64(totalRows) / float64(paging.PerPage)))
	// vtb.NextPageURL = "?page=3&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=field|aaaa,field2|bbbbbb"
	// vtb.PrevPageURL = "?page=1&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=field|aaaa,field2|bbbbbb"
	// vtb := utils.VueTableResponse{Datas: rs, Total: 2000, PerPage: 20, Form: 200, To: 300, CurrentPage: 1, LastPage: 200, NextPageURL: "", PrevPageURL: ""}
	alert := utils.Alert{Msg: "Vue Table สำเร็จ", Title: "Success", Type: "success"}
	payload.Alert = alert
	payload.Data = vtb
	payload.Pagination = paging
	payload.Meta = http.StatusText(200)
	utils.RespondJSON(c, 200, payload)
}
