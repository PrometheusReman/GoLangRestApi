package ownerModels

import (

	"GoLangApi/app/V1/libs/dbConnect/Mysql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

type Response struct {
	HqName string `json:"name"`
	HqShortForm string `json:"hq" form:"hq"`
}

func GetOwnerInfoModel() (res []Response, err error) {
	//places :=[]map[string]interface{}
	db := MysqlConnection.DbConn()
	selDB, err := db.Query("SELECT name, hq FROM hq_list limit 10")
	
    if err != nil {
        return res, nil
	}
	columns, _ := selDB.Columns()
	fmt.Println(columns)
    /* err2 := selDB.Scan(&res)
	if err2 != nil {
		return nil, err2
	} */
	defer db.Close() 

	for selDB.Next() {
		var response Response 
		selDB.Scan(&response.HqName, &response.HqShortForm)
		res = append(res, response)
	}
	fmt.Println(res)
	return res, nil
}