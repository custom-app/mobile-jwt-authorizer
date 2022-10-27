package main

import (
	"backend"
	"flag"
)

func main() {
	var db string
	flag.StringVar(&db, "db", "user=penachett password=samplesapp908365 dbname=jwt_sample sslmode=disable", "")
	flag.Parse()
	backend.InitServer(db)
}
