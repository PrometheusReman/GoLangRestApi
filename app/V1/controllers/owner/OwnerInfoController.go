package ownerControllers

import (
    
    "github.com/gin-gonic/gin"

    "encoding/json"
    
    models "GoLangApi/app/V1/models/owner" 
)

func GetOwnerInfo(c *gin.Context) {
    
    //id := c.Params.ByName("id")

    var errorBody = map[string]interface{}{}

    errorBody["error"] = "User Details Not Found"

    //c.JSON(200, c.Request.Header["Authorization"])

    buf := make([]byte, 1024)
    num, _ := c.Request.Body.Read(buf)
    reqBody := string(buf[0:num])
    
    var body interface{}
    json.Unmarshal([]byte(reqBody), &body) 
	
    res, err := models.GetOwnerInfoModel()
    
	if err == nil {
		c.JSON(200, res)
	} else {
		c.JSON(200, errorBody)
	} 
}