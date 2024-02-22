package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Avaliable bool
}

func main() {
	user := "postgres"
	password := "password"
	host := "database-1239.cpyasqckugo5.eu-north-1.rds.amazonaws.com"
	port := "5432"
	dbName := "pgadmindb"
	template := "postgres://%s:%s@%s:%s/%s"

	connStr := fmt.Sprintf(template, user, password, host, port, dbName)
	db, err := sql.Open("postgres", connStr)

	defer db.Close()
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	CreateTable(db)

	product := Product{"Book", 15.55, true}
	pk := InsertProduct(db, product)

	fmt.Println(pk)

}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6,2) NOT NULL,
    avaliable BOOLEAN,
    created timestamp DEFAULT NOW()
  )`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func InsertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, avaliable) VALUES ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Avaliable).Scan(&pk)
	if err != nil {
		panic(err)
	}
	return pk

}

//package main
//
//func main() {
//	done := make(chan struct{})
//	fmt.Println("Hello Gopher!")
//
//	user := "postgres"
//	password := "password"
//	host := "database-1239.cpyasqckugo5.eu-north-1.rds.amazonaws.com"
//	port := "5432"
//	dbName := "pgadmindb"
//	template := "postgres://%s:%s@%s:%s/%s"
//
//	connStr := fmt.Sprintf(template, user, password, host, port, dbName)
//
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//
//	_ = db
//
//	Setup(func() interface{} {
//		//canvasSize := js.Global().Get("document").Call("getElementById", "canvas").Call("getBoundingClientRect")
//		//CreateCanvas(canvasSize.Get("width").Int(), canvasSize.Get("height").Int())
//		CreateCanvas(WindowWidth(), WindowHeight())
//		return nil
//	})
//
//	doDraw := true
//
//	ellipseSize := 50
//	colors := []string{
//		"#bf616a",
//		"#8fbcbb",
//		"#d08770",
//		"#88c0d0",
//		"#ebcb8b",
//		"#81a1c1",
//		"#a3be8c",
//		"#5e81ac",
//		"#b48ead",
//	}
//	colorIndex := 0
//
//	Draw(func() interface{} {
//		BackgroundRGBA(46, 52, 64, 100)
//		StrokeWeight(3)
//		StrokeHex("#d8dee9")
//		FillHex(colors[colorIndex])
//		if doDraw {
//			if MouseIsPressed() && MouseButton() == "center" {
//				StrokeWeight(10)
//				Ellipse(float64(Width/2), float64(Height/2), float64(ellipseSize), float64(ellipseSize))
//			} else {
//				Ellipse(float64(MouseX()), float64(MouseY()), float64(ellipseSize), float64(ellipseSize))
//			}
//		}
//		return nil
//	})
//
//	WindowResized(func() interface{} {
//		ResizeCanvas(WindowWidth(), WindowHeight())
//		return nil
//	})
//
//	MousePressed(func() interface{} {
//		if MouseButton() == "right" {
//			doDraw = !doDraw
//		}
//		if !doDraw {
//			Clear()
//		}
//		return false
//	})
//
//	MouseClicked(func() interface{} {
//		colorIndex = (colorIndex + 1) % len(colors)
//		return false
//	})
//
//	MouseWheel(func(delta float64) interface{} {
//		ellipseSize -= int(delta * 0.05)
//		return false
//	})
//
//	<-done
//}
