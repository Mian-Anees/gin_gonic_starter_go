package Routes

import (
	"fmt"
	"todoGoGo/Controllers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	ccResolver "todoGoGo/graph/resolvers"
	ccGenerated "todoGoGo/graph/generated"

	"github.com/99designs/gqlgen/graphql/playground"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
		v1.POST("/CreateProduct", CreateProduct())
	}
	r.Any("/", PlaygroundHandler())
	return r;
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