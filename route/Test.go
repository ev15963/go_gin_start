package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"
)

func Test(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, albums)
	fmt.Println("dfdfdfdfdf")
	c.String(http.StatusOK, "hello world")
}
