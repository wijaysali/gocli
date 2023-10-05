package main

import (
	"basic/config"
	"basic/entity"
	"basic/handler"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	registerName := registerCmd.String("name", "", "Name of the Student")
	registerEmail := registerCmd.String("email", "", "Email of the student")

	if len(os.Args) < 2 {
		fmt.Println("expected 'register' or 'list' subcommands")
		os.Exit(1)
	}

	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Failed to connect db", err)
	}
	defer db.Close()

	switch os.Args[1] {
	case "register":
		registerCmd.Parse(os.Args[2:])
		student := entity.Student{
			Name:  *registerName,
			Email: *registerEmail,
		}
		if err := handler.RegisterStudent(db, student); err != nil {
			log.Fatal("Failed to register student:", err)
		}
		fmt.Println("Student registered successfully")

	case "list":
		listCmd.Parse(os.Args[2:])
		students, err := handler.ListStudents(db)
		if err != nil {
			log.Fatal("Failed to list Students:", err)
		}
		for _, s := range students {
			fmt.Println(s.ID, s.Name, s.Email)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
