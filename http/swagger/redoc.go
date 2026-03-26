package serverSwagger

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func redocHandler(cfg Config) gin.HandlerFunc {
	html := buildRedocHTML(cfg)

	return func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	}
}

func buildRedocHTML(cfg Config) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
  <head>
    <title>%s</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>

  <body>
    <redoc spec-url="%s"></redoc>

    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
  </body>
</html>
`, cfg.Title, cfg.SpecURL)
}
