package internal

import (
	"log"
	"status/database"
	"status/helper"
	"status/model"
)

func SeedDatabase() {
	hash, _ := helper.HashPassword("password")
	initialUser := model.User{Name: "Tim Zolleis", Email: "tim@zolleis.net", Password: hash}
	createUserResult := database.DB.Create(&initialUser)
	if createUserResult.Error != nil {
		log.Println("Could not create initial user")
	}
	initialWorkspace := model.Workspace{Name: "My Workspace", Users: []model.User{initialUser}}
	createWorkspaceResult := database.DB.Create(&initialWorkspace)
	if createWorkspaceResult.Error != nil {
		log.Println("Could not create initial workspace")
	}
	initialApiKey := model.ApiKey{WorkspaceID: initialWorkspace.ID, Value: "my-api-key"}
	createApiKeyResult := database.DB.Create(&initialApiKey)
	if createApiKeyResult.Error != nil {
		log.Println("Could not create initial api key")
	}
}
