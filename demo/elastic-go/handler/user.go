package handler

import (
	"net/http"
	"strconv"
	"strings"
	
	"github.com/gin-gonic/gin"
	
	"github.com/devhg/es/dao"
	"github.com/devhg/es/model"
	"github.com/devhg/es/resource"
	"github.com/devhg/es/service"
)

func Create(c *gin.Context) {
	
	serv := service.NewUserService(dao.NewUserES(resource.EsClient))
	
	users := make([]*model.UserEs, 0)
	user := model.UserEs{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}
	
	users = append(users, &user)
	if err := serv.BatchAdd(c, users); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

func Update(c *gin.Context) {
	
	serv := service.NewUserService(dao.NewUserES(resource.EsClient))
	
	users := make([]*model.UserEs, 0)
	user := model.UserEs{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}
	
	users = append(users, &user)
	if err := serv.BatchUpdate(c, users); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

func Delete(c *gin.Context) {
	
	serv := service.NewUserService(dao.NewUserES(resource.EsClient))
	
	users := make([]*model.UserEs, 0)
	user := model.UserEs{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}
	
	users = append(users, &user)
	if err := serv.BatchDel(c, users); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

// MGet
// curl --location --request GET 'http://localhost:8080/api/user/info?id=1,2,3'
func MGet(c *gin.Context) {
	
	serv := service.NewUserService(dao.NewUserES(resource.EsClient))
	
	ids := c.Query("id")
	IDS := make([]uint64, 0)
	for _, id := range strings.Split(ids, ",") {
		d, _ := strconv.Atoi(id)
		IDS = append(IDS, uint64(d))
	}
	
	res, err := serv.MGet(c, IDS)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
	})
}

func Search(c *gin.Context) {
	
	serv := service.NewUserService(dao.NewUserES(resource.EsClient))
	
	req := model.SearchRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "Invalid argument"})
		return
	}
	
	res, err := serv.Search(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": res,
	})
}
