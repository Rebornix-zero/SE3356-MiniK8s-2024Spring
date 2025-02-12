package handler

import (
	"fmt"
	"net/http"

	"github.com/MiniK8s-SE3356/minik8s/pkg/apiObject/yaml"
	"github.com/MiniK8s-SE3356/minik8s/pkg/apiserver/process"
	"github.com/gin-gonic/gin"
)

// POST 参数类型NamespaceDesc
func AddNamespace(c *gin.Context) {
	var desc yaml.NamespaceDesc
	if err := c.ShouldBindJSON(&desc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := process.AddNamespace(&desc)
	if err != nil {
		fmt.Println("error in process.AddNamespace ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

// POST 参数类型 {name: "xxx"}
func RemoveNamespace(c *gin.Context) {
	var param map[string]string
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name, ok := param["namespace"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no field 'name'"})
		return
	}

	result, err := process.RemoveNamespace(name)
	if err != nil {
		fmt.Println("error in process.RemoveNamespace ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetNamespace(c *gin.Context) {
	name := c.Query("namespace")
	if name == "" {
		result, err := process.GetNamespaces()
		if err != nil {
			fmt.Println("error in process.GetNamespaces")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"result": result})
	} else {
		result, err := process.GetNamespace(name)
		if err != nil {
			fmt.Println("error in process.GetNamespace")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

// func DescribeNamespace(c *gin.Context) {
// 	namespace := c.Query("namespace")
// 	if namespace == "" {
// 		namespace = "Default"
// 	}

// 	result, err := process.GetNamespaces()
// 	if err != nil {
// 		fmt.Println("error in process.GetNamespace ", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"result": result})
// }
