package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"


	"experiment-db/internal/cli"
	"experiment-db/internal/config"
	"github.com/mraraneda/mrlogger"
	"time"
)


func main() {

	var configFile string

	// 1.1 Controla los flags CLI del programa y carga las variables
	//      antes declaradas
	cli.FlagHandler(&configFile)

	init := time.Now()

	// 2. Se levanta la configuración estática desde config file
	ymlconf := config.YamlInput{PathInput: configFile}

	// 2.1 Se crea la nueva configuración a partir del yml File
	staticConf := config.NewStaticConfig(ymlconf.Read())

	// 2.2 Tomando nivel de log de la configuración
	mrlogger.NewLoggingLevel("Debug")

	mrlogger.Info("La configuración DB  :", staticConf.POSTGRES.Dbname)
	mrlogger.Info("La configuración HOST:", staticConf.POSTGRES.Host)
	mrlogger.Info("La configuración Password:", staticConf.POSTGRES.Password)
	mrlogger.Info("La configuración Port:", staticConf.POSTGRES.Port)
	mrlogger.Info("La configuración User:", staticConf.POSTGRES.User)
	mrlogger.Info("La configuración Sqlstatement", staticConf.POSTGRES.Sqlstatement)



	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		staticConf.POSTGRES.Host, staticConf.POSTGRES.Port, staticConf.POSTGRES.User, staticConf.POSTGRES.Password, staticConf.POSTGRES.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()


	err = db.Ping()
	if err != nil {
		panic(err)
	}





	elapse := time.Since(init)
	mrlogger.Info("Tiempo de procesamiento: ", elapse)

}

/* Falta por seguir aprendiendo de SQL

https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/         ç

 */


