package ownerRoutes 

import (

	"github.com/gin-gonic/gin"

	helper "GoLangApi/app/V1/libs/helper"

	controllers "GoLangApi/app/V1/controllers/owner"
)

func OwnerRoutes(route *gin.Engine) {
    
	version := helper.InitAppBaseUrls("versionNumber") 
	route.POST("/NewApi/"+version+"/Owner/OwnerInfo", controllers.GetOwnerInfo)
	route.POST("/NewApi/"+version+"/Owner/OwnerTrucksInfo", controllers.GetOwnerTrucksInfo)
}