package main

import (
    _ "github.com/denisenkom/go-mssqldb"
    "github.com/labstack/echo/v4"
)


func main() {
    e := echo.New()

    // Route to get components data
    e.GET("/components", getComponents)
    e.POST("/regist-component", pushComponents)
    e.POST("/regist-component-parents", pushComponentParents)


    e.Logger.Fatal(e.Start(":8011"))
}
