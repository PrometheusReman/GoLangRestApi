package analyticsRoutes 

import (

	"github.com/gin-gonic/gin"

	helper "GoLangApi/app/V1/libs/helper"

	controllers "GoLangApi/app/V1/controllers/analytics" 
)


func AnalyticsRoutes(route *gin.Engine) {
	
	version := helper.InitAppBaseUrls("versionNumber") 
    route.POST("/NewApi/"+version+"/Analytics/OwnerInfoCheck", controllers.OwnerInfoCheck)
}