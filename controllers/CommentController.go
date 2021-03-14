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
// controllers.SetupCommentRoutes(v1) 
func SetupCommentRoutes(rg *gin.RouterGroup) { 
	router := rg.Group("/comments") 
	// router.Use(utils.JwtVerify) // jwt verify with secret key ถ้าต้องการให้ auth ด้วย jwt 
	router.GET("/all", listComment) 
	router.GET("/last", listlastComment) 
	router.GET("/vuetable", vuetableComment) 
	router.GET("/by/:id", getOneComment) 
	router.GET("/del/:id", deleteComment) 
	router.POST("/new", addNewComment) 
	router.POST("/edit/:id", putOneComment) 
} 
 
// ListComments godoc 
// @summary List Comments 
// @description List all comments 
// @tags comments 
// @security ApiKeyAuth 
// @id ListComments 
// @accept json 
// @produce json 
// @response 200 {array} []models.Comment "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/comments/all [get] 
func listComment(c *gin.Context) { 
	var rs []models.Comment 
	payload := utils.GetResponse() 
	err := services.GetAllComment(&rs) 
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
 
// ListComments godoc 
// @summary List Comments 
// @description List all comments 
// @tags comments 
// @security ApiKeyAuth 
// @id ListlastComments 
// @accept json 
// @produce json 
// @response 200 {array} []models.Comment "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/comments/last [get] 
func listlastComment(c *gin.Context) { 
	var rs []models.Comment 
	payload := utils.GetResponse() 
	err := services.GetAllIdDescComment(&rs) 
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
 
// CreateComment godoc 
// @summary Create Comment 
// @description Create new comment 
// @tags comments 
// @security ApiKeyAuth 
// @id CreateComment 
// @accept json 
// @produce json 
// @param Comment body models.CommentForCreate true "Comment data to be created" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/comments/new [post] 
func addNewComment(c *gin.Context) { 
	var rs models.Comment 
	payload := utils.GetResponse() 
	c.BindJSON(&rs) 
	err := services.AddNewComment(&rs) 
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
 
// GetComment godoc 
// @summary Get Comment 
// @description  Get comment by id 
// @tags comments 
// @security Basic auth 
// @id GetComment 
// @accept json 
// @produce json 
// @param id path int true "id of comment to be gotten" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/comments/by/{id} [get] 
func getOneComment(c *gin.Context) { 
	id := c.Params.ByName("id") 
	payload := utils.GetResponse() 
	var rs models.Comment 
	err := services.GetOneComment(&rs, id) 
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
 
// UpdateComment godoc 
// @summary Update Comment 
// @description Update comment by id 
// @tags comments 
// @security ApiKeyAuth 
// @id UpdateComment 
// @accept json 
// @produce json 
// @param id path int true "id of comment to be updated" 
// @param Comment body models.CommentForUpdate true "Comment data to be updated" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/comments/edit/{id} [post] 
func putOneComment(c *gin.Context) { 
	var rs models.Comment 
	payload := utils.GetResponse() 
	id := c.Params.ByName("id") 
	err := services.GetOneComment(&rs, id) 
	if err != nil { 
		alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		payload.Meta = http.StatusText(404) 
		payload.Alert = alert 
		utils.RespondJSON(c, 200, payload) 
	} 
	c.BindJSON(&rs) 
	err = services.PutOneComment(&rs, id) 
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
 
// DeleteComment godoc 
// @summary Delete Comment 
// @description Delete comment by id 
// @tags comments 
// @security ApiKeyAuth 
// @id DeleteComment 
// @accept json 
// @produce json 
// @param id path int true "id of comment to be deleted" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/comments/del/{id} [get] 
func deleteComment(c *gin.Context) { 
	var rs models.Comment 
	payload := utils.GetResponse() 
	id := c.Params.ByName("id") 
	err := services.DeleteComment(&rs, id) 
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
 
// VuetableComment godoc 
// @summary List Comments use VueTable 
// @description List all comments with pagination keyword filter by column Sort columns 
// @tags comments 
// @security ApiKeyAuth 
// @id vuetableComment 
// @accept json 
// @produce json 
// @Param page query string false "number of page" 
// @Param per_page query string false "get number of per_page" 
// @Param sort query string false "soft by cols exm: col1|asc,col2|desc" 
// @Param filter query string false "search with column exm: col1|aaa,col2|bbb  by filter" 
// @Param kw query string false "search by kw exm: aaa bbb ccc " 
// @response 200 {object} []models.Comment "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// example /api/v1/comments/vuetable?page=2&per_page=2&sort=created_at|ASC,uid|DESC&kw=xxx&filter=name|aaa,author|bbb 
// @Router /api/v1/comments/vuetable [get] 
func vuetableComment(c *gin.Context) { 
	var rs []models.Comment 
	query := c.Request.URL.Query() 
	table := "comments" 
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
	cols := []string{} //ต้องกำหนดเอง 
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
 