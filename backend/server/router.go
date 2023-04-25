package server

import (
	"github.com/VolunteerOne/volunteer-one-app/backend/controllers"
	"github.com/VolunteerOne/volunteer-one-app/backend/database"
	"github.com/VolunteerOne/volunteer-one-app/backend/middleware"
	"github.com/VolunteerOne/volunteer-one-app/backend/repository"
	"github.com/VolunteerOne/volunteer-one-app/backend/service"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// *********************************************************
	// INITIALIZE REPOSITORIES HERE -> DB migration is handled in main.go
	// *********************************************************

	loginRepository := repository.NewLoginRepository(database.GetDatabase())
	usersRepository := repository.NewUsersRepository(database.GetDatabase())
	organizationRepository := repository.NewOrganizationRepository(database.GetDatabase())
	orgUsersRepository := repository.NewOrgUsersRepository(database.GetDatabase())
	eventRepository := repository.NewEventRepository(database.GetDatabase())

	// *********************************************************
	// INITIALIZE SERVICES HERE
	// *********************************************************

	loginService := service.NewLoginService(loginRepository)
	usersService := service.NewUsersService(usersRepository)
	organizationService := service.NewOrganizationService(organizationRepository)
	orgUsersService := service.NewOrgUsersService(orgUsersRepository)
	eventService := service.NewEventService(eventRepository)

	// *********************************************************
	// INITIALIZE CONTROLLERS HERE
	// *********************************************************

	loginController := controllers.NewLoginController(loginService)
	usersController := controllers.NewUsersController(usersService)
	organizationController := controllers.NewOrganizationController(organizationService)
	orgUsersController := controllers.NewOrgUsersController(orgUsersService)
	eventController := controllers.NewEventController(eventService)

	userGroup := router.Group("user")

	// userGroup := new(controllers.UsersController)
	userGroup.POST("/", usersController.Create)
	userGroup.GET("/:id", middleware.BasicAuth, usersController.One)
	userGroup.DELETE("/:id", usersController.Delete)
	userGroup.PUT("/:id", usersController.Update)

	loginGroup := router.Group("login")

	//Simple login, checks database against users inputted email and password to login
	loginGroup.GET("/:email/:password", loginController.Login)
	//Get the users email, sends a forgotten password code to them
	loginGroup.POST("/:email", loginController.SendEmailForPassReset)
	//Get the secret code from the users email, if matches reset password
	loginGroup.PUT("/:email/:resetcode/:newpassword", loginController.PasswordReset)
	//Check valid access token
	loginGroup.POST("/verify", middleware.BasicAuth, loginController.VerifyAccessToken)
	//Get refresh token
	loginGroup.POST("/refresh", loginController.RefreshToken)

	organizationGroup := router.Group("organization")
	organizationGroup.POST("/", organizationController.Create)
	organizationGroup.GET("/", organizationController.All)
	organizationGroup.GET("/:id", organizationController.One)
	organizationGroup.DELETE("/:id", organizationController.Delete)
	organizationGroup.PUT("/:id", organizationController.Update)

	eventGroup := router.Group("event")
	eventGroup.POST("/", eventController.Create)
	eventGroup.GET("/", eventController.All)
	eventGroup.GET("/:id", eventController.One)
	eventGroup.DELETE("/:id", eventController.Delete)
	eventGroup.PUT("/:id", eventController.Update)

	orgUsersGroup := router.Group("orgUsers")
	orgUsersGroup.POST("/", orgUsersController.CreateOrgUser)
	orgUsersGroup.GET("/", orgUsersController.ListAllOrgUsers)
	orgUsersGroup.GET("/:userId", orgUsersController.FindOrgUser)
	orgUsersGroup.PUT("/:userId", orgUsersController.UpdateOrgUser)
	orgUsersGroup.DELETE("/:userId", orgUsersController.DeleteOrgUser)

	// objectGroup := router.Group("object")
	// {
	// 	object := new(controllers.ObjectController)
	// 	objectGroup.POST("/", object.Create)
	// 	objectGroup.GET("/", object.All)
	// 	objectGroup.GET("/:id", object.One)
	// 	objectGroup.DELETE("/:id", object.Delete)
	// 	objectGroup.PUT("/:id", object.Update)
	// }

	// router.Use(middlewares.AuthMiddleware())

	// root := new(controllers.RootController)
	// router.GET("/", root.Get)

	// example := new(controllers.ExampleController)
	// router.GET("/example", controllers.ExampleGet)

	// v1 := router.Group("v1")
	// {
	// 	userGroup := v1.Group("user")
	// 	{
	// 		user := new(controllers.UserController)
	// 		userGroup.GET("/:id", user.Retrieve)
	// 	}
	// }

	return router
}
