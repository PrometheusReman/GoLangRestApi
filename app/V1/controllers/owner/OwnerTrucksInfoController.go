package ownerControllers

import (
    
    "github.com/gin-gonic/gin"
    
)

func GetOwnerTrucksInfo(c *gin.Context) {
    
    //name := "Shaju"

    //phone := "01844526764"

    var variable_name = map[string]interface{}{}

    variable_name["name"] = "TrucksReman"

    variable_name["year"] = 26

	c.JSON(200, variable_name)
}