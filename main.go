package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, firstName, lastName, pw string) *Account {
	acc, err := NewAccount(firstName, lastName, pw)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("account number: ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "John", "Doe", "1234")
}

func main() {
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding database...")
		// Seed accounts
		seedAccounts(store)
		return
	}

	server := NewAPIServer(":8080", store)
	server.Run()
}
