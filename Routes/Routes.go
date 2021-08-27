package routes

import (
	"fmt"
	"todoGoGo/controllers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	ccGenerated "todoGoGo/graph/generated"
	ccResolver "todoGoGo/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/playground"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("todo/:id", controllers.GetATodo)
		v1.PUT("todo/:id", controllers.UpdateATodo)
		v1.DELETE("todo/:id", controllers.DeleteATodo)
		v1.POST("/CreateProduct", CreateProduct())
	}
	r.Any("/", PlaygroundHandler())
	return r
}

func CreateProduct() gin.HandlerFunc {
	srv := handler.NewDefaultServer(ccGenerated.NewExecutableSchema(ccGenerated.Config{Resolvers: &ccResolver.Resolver{}}))
	return func(c *gin.Context) {
		fmt.Printf("----------")
		srv.ServeHTTP(c.Writer, c.Request)
	}

}
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
