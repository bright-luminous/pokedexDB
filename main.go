package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	operator := new(pokemonSQLop)
	operator.init("sql.DB")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}
	toPrint, err := operator.pokeDelete(ctx, 3)
	errCheck(err)
	fmt.Printf("Pokemon by name: %v\n\n", toPrint)
}

// test push
