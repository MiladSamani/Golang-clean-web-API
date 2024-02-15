package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type TestHandler struct {}

func NewTestHandler() *TestHandler {		
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (h *TestHandler) Users(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

func (h *TestHandler) UsersById(c *gin.Context)  {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UsersById",
		"id": id,
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context)  {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserByUsername",
		"id": username,
	})
}

func (h *TestHandler) Accounts(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
	})
}

func (h *TestHandler) AddUser(c *gin.Context)  {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Add Users",
		"id": id,
	})
}

func(h *TestHandler) HeaderBinder1(c *gin.Context)  {
	userId := c.GetHeader("userId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

type header struct {
	UserId string
	Browser string
}

func(h *TestHandler) HeaderBinder2(c *gin.Context)  {
	header:= header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"header": header,
	})
}

func(h *TestHandler) QueryBinder1(c *gin.Context)  {
	id :=c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"id": id,
		"name": name,
	})
}

func(h *TestHandler) QueryArrayBinder2(c *gin.Context)  {
	ids :=c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"ids": ids,
		"name": name,
	})
}

func(h *TestHandler) UriBinder(c *gin.Context)  {
	id :=c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "uri param",
		"ids": id,
		"name": name,
	})
}

type personData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

func(h *TestHandler) BodyBinder(c *gin.Context)  {
	person := personData{}
	err :=c.ShouldBindJSON(&person)
	if err != nil {
c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	"validation-error" : err.Error(),
})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"person" : person,
	})
}

func(h *TestHandler) FormBinder(c *gin.Context)  {
	person := personData{}
	c.ShouldBind(&person)

	c.JSON(http.StatusOK, gin.H{
		"result": "From",
		"person" : person,
	})
}

func (h *TestHandler) FileBinder(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }


    destination := "your_destination_path/" + file.Filename

    
    if err := c.SaveUploadedFile(file, destination); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "FileBinder",
        "file":   file.Filename,
    })
}