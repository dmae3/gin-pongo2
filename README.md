# gin-pongo2
pongo2 middleware for Gin

##Example:
```Go
package main

import (
    "http"

    "github.com/gin-gonic/gin"
    "github.com/daine46/gin-pongo2"
)

func main() {
    r := gin.Default()

    // first parameter is template's base directory
    r.HTMLRender = ginpon.New("app/views", gin.IsDebugging())

    r.GET("/", func(c *gin.Context) {
	    c.HTML(http.StatusOK, "index.tpl", ginpon.Context{"key": "value"})
    })

    r.Run(":3000")
}
```
