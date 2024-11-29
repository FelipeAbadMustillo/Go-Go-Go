package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func getAlerts(ctx *gin.Context) {
	//En el index se muestran las monedas en tendencia con su información básica.

	rows, err := conectionDB.Query("SELECT id_alert,username,price,condition,start_date,end_date,is_active FROM ALERTS")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var alerts []tools.Alert
	for rows.Next() {

		var a tools.Alert
		err := rows.Scan(&a.Id_alert, &a.Username, &a.Price, &a.Condition, &a.Start_date, &a.End_date, &a.Is_active)
		if err != nil {
			log.Fatal(err)
		}

		alerts = append(alerts, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	ctx.HTML(http.StatusOK, "get_alerts.tmpl", alerts) //Explicar bien como funcionan los templates

	//context.HTML(http.StatusOK, "user.tmpl", gin.H{
	//		"title": "Administracion de Usuario"}) //no puedo cargar la data de la base de datos

}
