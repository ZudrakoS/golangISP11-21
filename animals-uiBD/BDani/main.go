package main

import (
    "database/sql"
    "fmt"
    "log"
    "BDani/bd"
    _ "github.com/lib/pq"
)





const (
  host = "localhost"
    port = 5432
    user = "postgres"
    password = "123qweasdzxc"
    dbname = "postgres"

)

func main() {
    connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    animals, err := bd.GetAllAnimals(db)
    if err != nil {
        log.Fatal(err)
    }

    

    var n int
    fmt.Println("Сортировка:1 - Name")
    fmt.Scan(&n)
    switch n {
        case 1:
            animals, err = bd.GetSortedAnimalsByName(db)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Println("Вывод всех животных по Name:")
            for _, animal := range animals {
                fmt.Printf("ID: %d, Name: %s, Species: %s, Age: %d, Habitat: %s\n", animal.ID, animal.Name, animal.Species, animal.Age, animal.Habitat)
            }

        
            

    }

}
