package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Fprintf(os.Stdout,"Where with os.Stdout at %v\n", time.Now())
}