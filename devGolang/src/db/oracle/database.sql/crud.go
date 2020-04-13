package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "gopkg.in/rana/ora.v4"
)

var db *sql.DB

// Person bind struct of j2h2s2apa_person table in j2h2s2apa@TESTLTDB 57
type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) get() (person Person, err error) {
	var row *sql.Row = db.QueryRow("select id, first_name, last_name from j2h2s2apa_person where id=:1", p.ID)
	// row := db.QueryRow(`select id, first_name, last_name from j2h2s2apa_person where id=:id`, sql.Named("id", p.ID))

	err = row.Scan(&person.ID, &person.FirstName, &person.LastName)

	if err != nil {
		log.Fatalf("row get error : %s!", err)
		return
	}
	return
}

func (p *Person) getAll() (persons []Person, err error) {
	rows, err := db.Query("select id, first_name, last_name from j2h2s2apa_person")
	if err != nil {
		log.Fatalf("rows getAll error!")
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}

	defer rows.Close()
	return
}

func (p *Person) add() (ID int, err error) {
	stmt, err := db.Prepare("insert into j2h2s2apa_person(id, first_name, last_name) values (j2h2s2apa_person_id.nextval, :1, :2)")
	if err != nil {
		log.Fatal(err)
		return
	}

	rs, err := stmt.Exec(p.FirstName, p.LastName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer stmt.Close()

	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	ID = int(id)

	return
}

func (p *Person) update() (rows int, err error) {
	stmt, err := db.Prepare("update j2h2s2apa_person set first_name=:fn, last_name=:ln where id=:id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rs, err := stmt.Exec(p.FirstName, p.LastName, p.ID)
	if err != nil {
		log.Fatal(err)
	}

	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	rows = int(row)
	return
}

func (p *Person) del() (rows int, err error) {
	stmt, err := db.Prepare("delete from j2h2s2apa_person where id=:id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rs, err := stmt.Exec(p.ID)
	if err != nil {
		log.Fatal(err)
	}

	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	rows = int(row)
	return
}

func main() {
	var err error
	db, err = sql.Open("ora", "ltuser/lt2005@192.168.2.57:1522/TESTLTDB")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	var router *gin.Engine = gin.Default()

	router.GET("/persons", func(c *gin.Context) {
		p := Person{}
		persons, err := p.getAll()
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})

	router.GET("/person/:id", func(c *gin.Context) {
		var result gin.H
		id := c.Param("id")

		ID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}

		p := Person{
			ID: ID,
		}
		person, err := p.get()
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// add
	// curl http://127.0.0.1:8000/person -X POST -d '{"first_name": "rsj", "last_name": "你好"}' -H "Content-Type: application/json"
	router.POST("/person", func(c *gin.Context) {
		var p Person
		err := c.ShouldBind(&p)
		log.Printf(">>>> firstName lastName : %s", p.FirstName+" "+p.LastName)
		if err != nil {
			log.Fatalln(err)
		}

		ID, err := p.add()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(ID)
		name := p.FirstName + " " + p.LastName
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", name),
		})
	})

	//  curl http://127.0.0.1:8000/person/1 -X PUT -d "first_name=admin&last_name=reg"
	router.PUT("/person/:id", func(c *gin.Context) {
		var (
			p      Person
			buffer bytes.Buffer
		)

		id := c.Param("id")
		ID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		err = c.ShouldBind(&p)
		if err != nil {
			log.Fatal(err)
		}

		p.ID = ID
		rows, err := p.update()
		if err != nil {
			log.Fatal(err)
		}
		buffer.WriteString(p.FirstName)
		buffer.WriteString(" ")
		buffer.WriteString(p.LastName)
		name := buffer.String()

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully update to %s", name),
			"count":   rows,
		})
	})

	router.DELETE("/person/:id", func(c *gin.Context) {
		id := c.Param("id")
		ID, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			log.Fatal(err)
		}

		p := Person{ID: int(ID)}
		rows, err := p.del()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      fmt.Sprintf("Successfully deleted user: %s", id),
			"affect_count": rows,
			"content-Type": c.GetHeader("content-Type"),
		})
	})

	router.Run(":8080")
}
