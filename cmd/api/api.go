package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-stripe/internal/drivers"
	"go-stripe/internal/models"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
	}
	secretkey string
	frontend  string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	DB       models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting Back End server in %s mode on port %d", app.config.env, app.config.port)
	return srv.ListenAndServe()

}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|rroduction|maintenance}")
	flag.StringVar(&cfg.db.dsn, "dsn", "pnunn:rEdCatH4rRy*2120@tcp(192.168.44.131:3306)/widgets?parseTime=true&tls=false", "DSN")
	flag.StringVar(&cfg.smtp.host, "smtphost", "smtp.mailtrap.io", "smtp host")
	flag.IntVar(&cfg.smtp.port, "smtpport", 587, "smtp port")
	flag.StringVar(&cfg.smtp.username, "smtpuser", "c00378a4ac64a2", "smtp user")
	flag.StringVar(&cfg.smtp.password, "smtppass", "60ba2ea98b0f8d", "smtp password")
	flag.StringVar(&cfg.secretkey, "secret", "8qiGak4g6E9L&Dra3489achz*3907pyt", "secret key")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4002", "url to frontend")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := drivers.OpenDB(cfg.db.dsn)

	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		DB:       models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
