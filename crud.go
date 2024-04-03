package main

import (
	"fmt"

	"github.com/google/uuid"
)

func look() {
	fmt.Println("hello world")
	fmt.Println(uuid.New().String())
}
