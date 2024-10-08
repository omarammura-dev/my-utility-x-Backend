package http

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"myutilityx.com/graph"
)


func GraphQLHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		userId, _ := c.Get("userId")
		srv.ServeHTTP(c.Writer, c.Request.WithContext(context.WithValue(c.Request.Context(), "userId", userId)))
	}
}
