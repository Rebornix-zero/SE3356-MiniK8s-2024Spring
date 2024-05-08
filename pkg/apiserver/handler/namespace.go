package handler

import (
	"fmt"
	"net/http"

	"github.com/MiniK8s-SE3356/minik8s/pkg/apiserver/process"
	"github.com/MiniK8s-SE3356/minik8s/pkg/ty"
	"github.com/gin-gonic/gin"
)

// POST 参数类型NamespaceDesc
func AddNamespace(c *gin.Context) {
	var desc ty.NamespaceDesc
	if err := c.ShouldBindJSON(&desc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := process.AddNamespace(&desc)
	if err != nil {
		fmt.Println("error in process.AddNamespace ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}

// DELETE 参数类型 {name: "xxx"}
func RemoveNamespace(c *gin.Context) {
	var param map[string]string
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name, ok := param["name"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no field 'name'"})
		return
	}

	result, err := process.RemoveNamespace(name)
	if err != nil {
		fmt.Println("error in process.RemoveNamespace ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}

// GET 无参数
func GetNamespaces(c *gin.Context) {
	result, err := process.GetNamespaces()
	if err != nil {
		fmt.Println("error in process.GetNamespace ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, result)
}
