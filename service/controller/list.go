package controller

import (
	"api-pos/db"
	"api-pos/dto"
	"api-pos/models"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type List struct{}

func (c List) FindAll(ctx *gin.Context){
	var list []models.Todo_list
	db.Conn.Find(&list)

	var result []dto.ListResponse
	for _, list := range list{
		result = append(result, dto.ListResponse{
			ID: list.ID,
			Content: list.Content,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c List) FindOne(ctx *gin.Context){
	id := ctx.Param("id")
	var list models.Todo_list
	if err := db.Conn.First(&list, id).Error; errors.Is(err, gorm.ErrRecordNotFound){
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ListResponse{
		ID: list.ID,
		Content: list.Content,
	})
}

func (c List) Create(ctx *gin.Context){
	var form dto.ListRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("testttttt ")
	list := models.Todo_list{
		Content: form.Content,
	}

	if err := db.Conn.Create(&list).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ListResponse{
		ID:   list.ID,
		Content: list.Content,
	})
}

func (c List) Update(ctx *gin.Context){
	id := ctx.Param("id")
	var form dto.ListRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var list models.Todo_list
	if err := db.Conn.First(&list, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	list.Content = form.Content
	db.Conn.Save(&list)
	ctx.JSON(http.StatusOK, dto.ListResponse{
		ID:   list.ID,
		Content: list.Content,
	})
}

func (c List) Delete(ctx *gin.Context){
	id := ctx.Param("id")
	//soft delete
	//db.Conn.Delete(&models.Category{}, id)
	//Hard delete
	db.Conn.Unscoped().Delete(&models.Todo_list{}, id)

	ctx.JSON(http.StatusOK, gin.H{"delete at ": time.Now()})
}