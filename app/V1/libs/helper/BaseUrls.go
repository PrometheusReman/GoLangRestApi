package helper

import "fmt" 

func InitAppBaseUrls(keyName string) (string) {

    var BaseUrlsList = map[string]interface{}{}

    BaseUrlsList["versionNumber"] = "V1"

    for key, value := range BaseUrlsList {
        if key == keyName {
            str := fmt.Sprintf("%v", value)
            return str 
        } 
    }
    return "NotFound"
}