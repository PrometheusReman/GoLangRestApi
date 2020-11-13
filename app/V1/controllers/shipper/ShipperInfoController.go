package shipperControllers

import (
    
    "github.com/gin-gonic/gin"
    
)

func GetShipperInfo(c *gin.Context) {
    
    //name := "Shaju"

    //phone := "01844526764"

    var variable_name = map[string]interface{}{}

    variable_name["name"] = "ShipperReman"

    variable_name["year"] = 26

	c.JSON(200, variable_name)
}