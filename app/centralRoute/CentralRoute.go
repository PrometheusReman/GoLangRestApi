package centralRoute 

import (

	"github.com/gin-gonic/gin"

	analyticsRoutesV1 "GoLangApi/app/V1/routes/analytics"
	ownerRoutesV1 "GoLangApi/app/V1/routes/owner"
	shipperRoutesV1 "GoLangApi/app/V1/routes/shipper"
)

func GetAllRoutes(route *gin.Engine) {
    
	analyticsRoutesV1.AnalyticsRoutes(route)
	ownerRoutesV1.OwnerRoutes(route)
	shipperRoutesV1.ShipperRoutes(route)
}