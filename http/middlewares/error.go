package middlewares

import (
	goerrors "errors"
	"net/http"

	"github.com/gin-gonic/gin"

	apperrors "github.com/onlyizi/onlyizi-go/errors"
	"github.com/onlyizi/onlyizi-go/observability/logs"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		if c.Writer.Written() {
			return
		}

		err := c.Errors.Last().Err

		logger := logs.FromContext(c.Request.Context())

		var appErr *apperrors.AppError

		if goerrors.As(err, &appErr) {

			if appErr.Status >= 500 {
				logger.Error(
					"request failed",
					logs.ErrorCode(appErr.Code),
					logs.Status(appErr.Status),
					logs.Err(err),
				)
			} else {
				logger.Warn(
					"request failed",
					logs.ErrorCode(appErr.Code),
					logs.Status(appErr.Status),
				)
			}

			c.JSON(appErr.Status, gin.H{
				"error": gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				},
			})

			c.Abort()
			return
		}

		logger.Error(
			"unexpected error",
			logs.Err(err),
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    apperrors.CodeInternalError,
				"message": "internal server error",
			},
		})

		c.Abort()
	}
}
