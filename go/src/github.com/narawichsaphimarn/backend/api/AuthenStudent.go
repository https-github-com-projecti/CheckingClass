package	api

import (
	"github.com/narawichsaphimarn/backend/models"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var dataAuthen []models.AuthenData
var newdataAuthen []models.AuthenData

func GetAuthenData(c *gin.Context){
	e := models.AuthenData {
		Id : 1,
		NameStudent : "นาย นราวิชญ์  สาพิมาน",
		IdStudent : "B5909711",
		AuthenStudent : []models.AuthenStudent {
			models.AuthenStudent{
				Date : "09/21/2019",
				StateAuthen : true,
			},
			models.AuthenStudent{
				Date : "09/22/2019",
				StateAuthen : false,
			},
			models.AuthenStudent{
				Date : "09/23/2019",
				StateAuthen : false,
			},
			models.AuthenStudent{
				Date : "09/23/2019",
				StateAuthen : false,
			},
			models.AuthenStudent{
				Date : "09/21/2019",
				StateAuthen : false,
			},
			models.AuthenStudent{
				Date : "09/22/2019",
				StateAuthen : true,
			},
			models.AuthenStudent{
				Date : "09/23/2019",
				StateAuthen : false,
			},
		},
	}

	dataAuthen = append(dataAuthen, e)
	defer c.Request.Body.Close()

	// for _, copy := range dataAuthen {
	// 	fmt.Println(copy.Id)
	// 	for _, copy2 := range copy.AuthenStudent{
	// 		if (copy2.Date == "09/21/2019"){
	// 			copy2.Date =  "09/25/2020"
	// 			copy.AuthenStudent = append(copy.AuthenStudent, copy2)
	// 		}
	// 		fmt.Println(copy2)
	// 	}
	// 	newdataAuthen = append(newdataAuthen, copy)
	// }
	// fmt.Println(newdataAuthen)
	c.JSON(http.StatusOK, dataAuthen)
}
