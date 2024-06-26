package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/timzolleis/smallstatus/database"
	"github.com/timzolleis/smallstatus/internal"
	"github.com/timzolleis/smallstatus/middleware"
	"github.com/timzolleis/smallstatus/routes"
	"log"
)

func main() {
	migrateFlag := flag.Bool("migrate", false, "Set to true to run database migrations")
	startFlag := flag.Bool("start", false, "Set to true to start the server")
	seedFlag := flag.Bool("seed", false, "Set to true to seed the database")
	db, err := database.Connect("database.db")
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err.Error())
	}
	database.DB = db
	flag.Parse()

	if *migrateFlag {
		fmt.Println("Running database migrations...")
		database.Migrate()

	}
	if *seedFlag {
		fmt.Println("Seeding the database...")
		internal.SeedDatabase()
	}

	if *startFlag {
		fmt.Println("Starting the server...")
		e := echo.New()
		e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
		apiBaseGroup := e.Group("/api")
		workspaceBaseGroup := apiBaseGroup.Group("/workspaces/:workspaceId")
		workspaceBaseGroup.Use(middleware.AuthMiddleware)
		workspaceBaseGroup.Use(middleware.WorkspaceMemberMiddleware)
		routes.RegisterMonitorRoutes(workspaceBaseGroup)
		routes.RegisterAuthRoutes(apiBaseGroup)
		e.Logger.Fatal(e.Start(":8080"))
	}

	// If no flags are set, or you need to handle other operations
	if !*migrateFlag && !*startFlag && !*seedFlag {
		fmt.Println("No operation specified. Use -migrate or -start.")
	}
}
