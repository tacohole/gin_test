package main

func initializeRoutes() {
	router.Use(setUserStatus())

	router.GET("/", showIndexPage)
	
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.PUT("/register", ensureNotLoggedIn(), register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/article/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}

}
