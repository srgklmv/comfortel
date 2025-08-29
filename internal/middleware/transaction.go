package middleware

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srgklmv/comfortel/pkg/logger"
)

func Transaction(conn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := conn.BeginTx(c.Request.Context(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
		if err != nil {
			logger.Error("transaction begin err:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			c.Next()
		}

		c.Set("tx", tx)
		c.Next()

		err = tx.Commit()
		if err != nil {
			logger.Error("transaction commit err:", err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}
