package main

import (
	"gin-flemarket/controllers"
	"gin-flemarket/infra"
	"gin-flemarket/repositories"
	"gin-flemarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
  itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindALl)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
