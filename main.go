package main

import (
	"diskon/database"
	"diskon/pkg/mysql"
	"diskon/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Barang struct {
	nama   string
	kode   string
	diskon int
}

var barangBeli = Barang{
	nama:   "kopi",
	kode:   "KP",
	diskon: 10,
}

func main() {
	harga := 1000

	diskon := (harga * barangBeli.diskon) / 100
	hargaTotal := harga - diskon

	fmt.Println(hargaTotal)

	e := echo.New()

	mysql.ConDB()
	database.MigrateDB()
	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:3330"))
}
