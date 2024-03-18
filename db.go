package main

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func getDB() (*sql.DB, error) {
	connString := "server=localhost;user id=admin;password=admin;database=bom_multilevel"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type ComponentData struct {
    ComponentID   string   `json:"component_id"`
    ComponentDesc *string  `json:"component_desc"`
    ComponentInv  *string   `json:"component_inv"`
    ParentID      *string  `json:"parent_id"`
    ParentDesc    *string  `json:"parent_desc"`
    ParentInv     *string  `json:"parent_inv"`
    Net           *float64 `json:"net"`
}


func getComponents(c echo.Context) error {
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT c.id AS component_id, c.[desc] AS component_desc, c.inv AS component_inv,
               cp.parent_id AS parent_id, p.[desc] AS parent_desc, p.inv AS parent_inv, cp.net
        FROM components c
        LEFT JOIN component_parents cp ON c.id = cp.component_id
        LEFT JOIN components p ON cp.parent_id = p.id
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Process the rows and return the data as JSON
	var componentData []ComponentData
	for rows.Next() {
		var data ComponentData
		var componentDesc, parentDesc sql.NullString
		var parentInv sql.NullString
		var net sql.NullFloat64
	
		err := rows.Scan(&data.ComponentID, &componentDesc, &data.ComponentInv, &data.ParentID, &parentDesc, &parentInv, &net)
		if err != nil {
			log.Fatal(err)
		}
	
		if componentDesc.Valid {
			data.ComponentDesc = &componentDesc.String
		}
		if parentDesc.Valid {
			data.ParentDesc = &parentDesc.String
		}
		if parentInv.Valid {
			data.ParentInv = &parentInv.String
		}
		if net.Valid {
			data.Net = &net.Float64
		}
	
		componentData = append(componentData, data)
	}
	

	return c.JSON(http.StatusOK, componentData)
}
