package controllers

import (
	"net/http"
	"strconv"
	// "time"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"go-rest-api/database"
	"go-rest-api/models"
)

func OrderGet(c *gin.Context) {
	db 			:= database.GetDB()
	orderId, _ 	:= strconv.Atoi(c.Param("orderId"))
	Orders 		:= models.Order{}
	err 		:= db.Find(&Orders, uint(orderId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Orders)
}

func OrderGetList(c *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}
	err := db.Find(&Orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Orders)
}

func OrderCreate(c *gin.Context) {
	// sample endpoint:
	// {{host}}/order/1
	// sample body:
	// {
	// 	"id": 5,
	// 	"customer_name": "Mrs. Dwi Hasan Asono",
	// 	"order_at": "2022-12-04T21:21:46Z",
	// 	"items": [
	// 	  {
	// 		"id": 6,
	// 		"item_code": "jhx",
	// 		"description": "IronMan Hand Pillow",
	// 		"quantity": 77
	// 	  },
	// 	  {
	// 		"id": 7,
	// 		"item_code": "jhr",
	// 		"description": "Golden Watch ColCompass 10M",
	// 		"quantity": 3
	// 	  },
	// 	  {
	// 		"id": 8,
	// 		"item_code": "asr1",
	// 		"description": "Bicycle: Electric-Monstaspeed Hybrid",
	// 		"quantity": 10
	// 	  },
	// 	  {
	// 		"id": 9,
	// 		"item_code": "asr8",
	// 		"description": "Bicycle: Cheetah Thin 9r23u",
	// 		"quantity": 10
	// 	  }
	// 	]
	// }
	db := database.GetDB()

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	fmt.Println("HASIL BINDING:")
	fmt.Println(input.Items)
	
	// Input Create Order
	err := db.Debug().Create(&input).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request on header",
			"message": err.Error(),
		})
		return
	}

	// Input Create Item
	err2 := db.Debug().Create(&input.Items).Error
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request on detail",
			"message": err2.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}

func OrderUpdate(c *gin.Context) {
	// sample endpoint:
	// {{host}}/order/1
	// sample body:
	// {
	// 	"id": 1,
	// 	"customer_name": "Spike Tyke",
	// 	"order_at": "2019-11-09T21:21:46Z",
	// 	"items": [
	// 	  {
	// 		"id": 1,
	// 		"item_code": "123",
	// 		"description": "IPhone 10X",
	// 		"quantity": 10
	// 	  }
	// 	]
	//  }
	db := database.GetDB()

	var orders models.Order
	err := db.First(&orders, "id = ?", c.Param("orderId")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}	
	// fmt.Println("HASIL BINDING:")
	// fmt.Println(input)

	// Input Update Order
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}), // column needed to be updated
	}).Create(&input)

	// Input Update Item
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{"item_code", "description", "quantity"}), // column needed to be updated
	}).Create(&input.Items)


	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code" : http.StatusOK,
	})
}

func OrderDelete(c *gin.Context) {
	db := database.GetDB()

	orderId, _ := strconv.Atoi(c.Param("orderId"))
	Orders := models.Order{}
	Items := models.Items{}

	err := db.Delete(Orders, uint(orderId)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on header",
			"message": err.Error(),
		})
		return
	}
	err2 := db.Where("order_id = ?", uint(orderId)).Delete(Items).Error
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request on detail",
			"message": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order has been successfully deleted",
	})
}
