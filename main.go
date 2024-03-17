package main

import (
	"company/database"
	"company/handlers"
	"company/repositories"
	"company/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db := database.Postgresql()
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อกับฐานข้อมูล
	err := db.Ping()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	// สร้าง instances ของ repositories, services, และ handlers
	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	router.POST("/api/register", h.AddRegisteHandler)
	router.GET("/api/getAllEmployees", h.GetAllEmployeesHandler)
	router.GET("/api/getCompany/:cpn_id", h.GetCompanyByIDHandler)
	router.PATCH("/api/updateEmployee", h.UpdateEmployeeHandler)
	router.DELETE("/api/deleteEmployee/:epy_id", h.DeleteEmployeeHandler)
	

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
