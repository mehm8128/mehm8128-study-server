package router

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting() {
	port := os.Getenv("PORT")
	//port := "8000"

	// store, err := pgstore.NewPGStore("user=mehm8128 password=math8128 dbname=mehm8128_study sslmode=disable", []byte("secret-key"))
	// if err != nil {
	// 	panic(err)
	// }

	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(session.Middleware(store))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://mehm8128-study-client.vercel.app"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	//defer store.Close()
	//defer store.StopCleanup(store.Cleanup(time.Minute * 5))

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
			apiUsers.GET("", getUsers)
			apiUsers.GET("/:id", getUser)
			//todo:セッションから取るので後回しapiUsers.GET("/me", getMe)
			//todo:セッションから取るので後回しapiUsers.PUT("/me", putMe)
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
		}
	}
	e.Logger.Fatal(e.Start(":" + port))
}
