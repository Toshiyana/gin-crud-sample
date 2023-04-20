package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

// PostRoute -> Route for quetion module
type PostRoute struct {
	Controller controller.PostController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice routes
func NewPostRoute(controller controller.PostController, handler infrastructure.GinRouter) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
// Setup method is used to configure endpoint for post APIs.
func (p PostRoute) Setup() {
	post := p.Handler.Gin.Group("/posts")
	{
		post.GET("/", p.Controller.GetPosts)
		post.POST("/", p.Controller.AddPost)
		post.GET("/:id", p.Controller.GetPost)
		post.DELETE("/:id", p.Controller.DeletePost)
		post.PUT("/:id", p.Controller.UpdatePost)
	}
}
