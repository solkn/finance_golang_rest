package main

import (
	"github.com/gin-gonic/gin"

	"finance/controllers"
	"finance/db"
	"finance/repository"
	"finance/services"
)

func main() {

	dbConn := db.ConnectDatabase("postgres://postgres:123@localhost/finance?sslmode=disable")

	router := gin.Default()

	userRepo := repository.NewUserRepository(dbConn)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	locationRepo := repository.NewLocationRepository(dbConn)
	loctionService := services.NewLocationService(locationRepo)
	locationController := controllers.NewLocationController(loctionService)

	orgRepo := repository.NewOrgRepository(dbConn)
	orgService := services.NewOrgService(orgRepo)
	orgController := controllers.NewOrgController(orgService)

	payeeRepo := repository.NewPayeeRepository(dbConn)
	payeeService := services.NewPayeeService(payeeRepo)
	payeeController := controllers.NewPayeeController(payeeService)

	registerRepo := repository.NewRegisterRepository(dbConn)
	registerService := services.NewRegisterService(registerRepo)
	registerController := controllers.NewRegisterController(registerService)

	tagRepo := repository.NewTagRepository(dbConn)
	tagService := services.NewTagService(tagRepo)
	tagController := controllers.NewTagController(tagService)

	transactionRepo := repository.NewTransactionRepository(dbConn)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	txCategoryRepo := repository.NewTxCategoryRepository(dbConn)
	txCategoryService := services.NewTxCategoryService(txCategoryRepo)
	txCategoryController := controllers.NewTxCategoryController(txCategoryService)

	txCatTypeRepo := repository.NewTxCatTypeRepository(dbConn)
	txCatTypeService := services.NewTxCatTypeService(txCatTypeRepo)
	txCatTypeController := controllers.NewTxCatTypeController(txCatTypeService)

	txTagRepo := repository.NewTxTagRepository(dbConn)
	txTagService := services.NewTxTagService(txTagRepo)
	txTagController := controllers.NewTxTagController(txTagService)

	txLinesRepo := repository.NewTxLinesRepository(dbConn)
	txLinesService := services.NewTxLinesService(txLinesRepo)
	txLinesController := controllers.NewTxLinesController(txLinesService)

	router.GET("/users", userController.GetUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.CreateUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	router.GET("/locations", locationController.GetLocations)
	router.GET("/locations/:id", locationController.GetLocationByID)
	router.POST("/locations", locationController.CreateLocation)
	router.PUT("/locations", locationController.UpdateLocation)
	router.DELETE("/locations/:id", locationController.DeleteLocation)

	router.GET("/orgs", orgController.GetOrgs)
	router.GET("/orgs/:id", orgController.GetOrgByID)
	router.POST("/orgs", orgController.CreateOrg)
	router.PUT("/orgs", orgController.UpdateOrg)
	router.DELETE("/orgs/:id", orgController.DeleteOrg)

	router.GET("/payees", payeeController.GetPayees)
	router.GET("/payees/:id", payeeController.GetPayeeByID)
	router.POST("/payees", payeeController.CreatePayee)
	router.PUT("/payees", payeeController.UpdatePayee)
	router.DELETE("/payees/:id", payeeController.DeletePayee)

	router.GET("/registers", registerController.Registers)
	router.GET("/registers/:id", registerController.GetRegisterByID)
	router.POST("/registers", registerController.CreateRegister)
	router.PUT("/registers", registerController.UpdateRegister)
	router.DELETE("/registers/:id", registerController.DeleteRegister)

	router.GET("/tags", tagController.GetTags)
	router.GET("/tags/:id", tagController.GetTagByID)
	router.POST("/tags", tagController.CreateTag)
	router.PUT("/tags", tagController.UpdateTag)
	router.DELETE("/tags/:id", tagController.DeleteTag)

	router.GET("/transactions", transactionController.GetTransactions)
	router.GET("/transactions/:id", transactionController.GetTransactionByID)
	router.POST("/transactions/", transactionController.CreateTransaction)
	router.PUT("/transactions/:id", transactionController.UpdateTransaction)
	router.DELETE("/transactions/:id", transactionController.DeleteTransaction)

	router.GET("/txCategories", txCategoryController.GetTxCategorys)
	router.GET("/txCategories/:id", txCategoryController.GetTxCategoryByID)
	router.POST("/txCategories", txCategoryController.CreateTxCategory)
	router.PUT("/txCategories", txCategoryController.UpdateTxCategory)
	router.DELETE("/txCategories/:id", txCategoryController.DeleteTxCategory)

	router.GET("/txCatTypes", txCatTypeController.GetTxCatTypes)
	router.GET("/txCatTypes/:id", txCatTypeController.GetTxCatTypeByID)
	router.POST("/txCatTypes", txCatTypeController.CreateTxCatType)
	router.PUT("/txCatTypes", txCatTypeController.UpdateTxCatType)
	router.DELETE("/txCatTypes/:id", txCatTypeController.DeleteTxCatType)

	router.GET("/txTags", txTagController.GetTxTags)
	router.GET("/txTags/:id", txTagController.GetTxTagByID)
	router.POST("/txTags", txTagController.CreateTxTag)
	router.PUT("/txTags", txTagController.UpdateTxTag)
	router.DELETE("/txTags/:id", txTagController.DeleteTxTag)

	router.GET("/txlines", txLinesController.GetTxLines)
	router.GET("/txlines/:id", txLinesController.GetTxLinesByID)
	router.POST("/txlines", txLinesController.CreateTxLines)
	router.PUT("/txlines", txLinesController.UpdateTxLines)
	router.DELETE("/v/:id", txLinesController.DeleteTxLines)

	router.Run("localhost:8080")
}
