package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

type UrlRequestBody struct {
	Url string
}

func RunWebApi(conn *pgx.Conn) {
	router := gin.Default()

	getAllStatus(conn, router)
	getLatestStatus(conn, router)
	postUrl(conn, router)
	deleteUrl(conn, router)

	router.Run()
}

func getAllStatus(conn *pgx.Conn, router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetStatusChecks(conn))
	})
}

func getLatestStatus(conn *pgx.Conn, router *gin.Engine) {
	router.GET("/status/latest", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetLatestStatusChecks(conn))
	})
}

// TODO handle conflict (adding existing url)
func postUrl(conn *pgx.Conn, router *gin.Engine) {
	router.POST("/urls", func(c *gin.Context) {
		var requestBody UrlRequestBody

		if err := c.BindJSON(&requestBody); err != nil {
			log.Fatal("Request body is not in the expected format. Error: ", err)
		}

		// TODO make a request before adding, must be 200 (only public URLs)

		c.JSON(http.StatusCreated, AddUrl(conn, requestBody.Url))
	})
}

func deleteUrl(conn *pgx.Conn, router *gin.Engine) {
	router.DELETE("/urls/:id", func(c *gin.Context) {
		id := c.Param("id")
		DeleteUrl(conn, id)
		c.JSON(http.StatusNoContent, id)
	})
}
