package router

import (
	"net/http"
	"os"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting() {
	port := os.Getenv("PORT")
	//port := "8000"
	//store, err := pgstore.NewPGStore("user=mehm8128 password=math8128 dbname=mehm8128_study sslmode=disable", []byte("sessions"))
	store, err := pgstore.NewPGStore(os.Getenv("DATABASE_URL"), []byte("sessions"))
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://mehm8128-study-client.vercel.app"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	defer store.Close()
	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	api := e.Group("/api")
	{
		apiPing := api.Group("/ping")
		{
			apiPing.GET("", func(c echo.Context) error {
				return echo.NewHTTPError(http.StatusOK, "pong!")
			})
		}
		apiUsers := api.Group("/users")
		{
			apiUsers.POST("/signup", postSignUp)
			apiUsers.POST("/login", postLogin)
			apiUsers.POST("/logout", postLogout)
			apiUsers.GET("", getUsers)
			apiUsers.GET("/me", getMe)
			apiUsers.GET("/:id", getUser)
			apiUsers.PUT("/me", putMe)
		}
		apiGoals := api.Group("/goals")
		{
			apiGoals.GET("", getGoals)
			apiGoals.POST("", postGoal)
			apiGoals.GET("/:id", getGoal)
			apiGoals.PUT("/:id", putGoal)
			apiGoals.DELETE("/:id", deleteGoal)
			apiGoals.GET("/user/:id", getGoalsByUser)
			apiGoals.PUT("/favorite/:id", putGoalFavorite)
		}
		apiRecords := api.Group("/records")
		{
			apiRecords.GET("", getRecords)
			apiRecords.POST("", postRecord)
			apiRecords.GET("/:id", getRecord)
			apiRecords.PUT("/:id", putRecord)
			apiRecords.DELETE("/:id", deleteRecord)
			apiRecords.GET("/user/:id", getRecordsByUser)
			apiRecords.PUT("/favorite/:id", PutRecordFavorite)
		}
		apiMemorize := api.Group("/memorizes")
		{
			apiMemorize.GET("", getMemorizes)
			apiMemorize.POST("", postMemorize)
			apiMemorize.GET("/:id", getMemorize)
			apiMemorize.POST("/:id/words", postWord)
			apiMemorize.GET("/:id/quiz", getQuiz)
		}
		apiFiles := api.Group("/files")
		{
			apiFiles.POST("", postFile)
			apiFiles.GET("/:id", getFile)
			apiFiles.GET("/:id/info", getFileInfo)
		}
	}
	e.Logger.Fatal(e.Start(":" + port))
}
