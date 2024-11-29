package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Go-Go-Go/internal/tools"
	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func postAlertsJSON(context *gin.Context) {
	var newAlert tools.Alert

	if err := context.BindJSON(&newAlert); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid  Request payload"})
		return
	}

	stmt, err := conectionDB.Prepare("INSERT INTO alerts (id_alert,username,price,condition,start_date,end_date,is_active) values ($1,$2,$3,$4,$5,$6,$7)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newAlert.Id_alert, newAlert.Username, newAlert.Price, newAlert.Condition, newAlert.Start_date, newAlert.End_date, newAlert.Is_active); err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusCreated, newAlert)

}

func viewAlerts(context *gin.Context) {
	//En el index se muestran las monedas en tendencia con su información básica.

	//context.HTML(http.StatusOK, "user.tmpl", nil) //Explicar bien como funcionan los templates
	context.HTML(http.StatusOK, "view_alerts.tmpl", gin.H{
		"title": "Alta de Alertas"})

}

func postAlerts(context *gin.Context) {
	var username string
	var price string
	var start_date string
	var end_date string
	var is_active string = "N"
	var condition string = "i"

	if context.Request.Method == "POST" {

		username = context.Request.FormValue("username")
		price = context.Request.FormValue("price")
		start_date = context.Request.FormValue("start_date")
		end_date = context.Request.FormValue("end_date")

	}

	rows := conectionDB.QueryRow("SELECT max(id_alert) FROM ALERTS")
	var max_id_alert int
	err := rows.Scan(&max_id_alert)
	if err != nil {
		log.Fatal(err)
	}
	max_id_alert = max_id_alert + 1

	stmt, err := conectionDB.Prepare("INSERT INTO alerts (id_alert,username,price,condition,start_date,end_date,is_active) values ($1,$2,$3,$4,$5,$6,$7)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(max_id_alert, username, price, condition, start_date, end_date, is_active); err != nil {
		fmt.Println("id_alert: " + strconv.Itoa(max_id_alert) + "username: " + username + "price: " + price + "start_date: " + start_date + "end_date: " + end_date)

		log.Fatal(err)
	}

	context.Redirect(301, "/alerts")

}
