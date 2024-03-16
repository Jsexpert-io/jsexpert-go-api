package middelwear

import (
	"golangproject/database"
	"golangproject/models/project"

	"github.com/gin-gonic/gin"
)

func GetProjectByClientId(clientId string) (*project.Project, error){
	var pro project.Project
	
	err := database.Database.Where("\"clientId\"=?", clientId).First(&pro).Error
	if err != nil {
	  print("err.Error()",err.Error())
	 return &project.Project{}, err
	}
	return &pro, nil
}
func ProjectAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		// before request
		headers := c.Request.Header
		clientId := headers.Get("clientId")
		clientSecret := headers.Get("clientSecret")
		pro, err := GetProjectByClientId(clientId)
		print(pro.ClientId)
		if err != nil {
			println(err)
			c.JSON(401, gin.H{"error": "Invalid client id"})
			c.Abort()
			return
		}
		if pro.ClientSecret != clientSecret {
			c.JSON(401, gin.H{"error": "Invalid client secret"})
			c.Abort()
			return
		}
		// addproject to context
		c.Set("project", pro.ID)
		c.Next()
		// after request
	}
}