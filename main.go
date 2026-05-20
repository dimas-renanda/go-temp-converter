package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr, "Usage: temp-converter <value><unit>  e.g. 100c 212f 373k"); os.Exit(1) }
	arg := strings.ToLower(os.Args[1])
	unit := string(arg[len(arg)-1])
	val, err := strconv.ParseFloat(arg[:len(arg)-1], 64)
	if err != nil { fmt.Fprintln(os.Stderr, "Invalid number"); os.Exit(1) }

	var c, f, k float64
	switch unit {
	case "c": c=val; f=c*9/5+32; k=c+273.15
	case "f": f=val; c=(f-32)*5/9; k=c+273.15
	case "k": k=val; c=k-273.15; f=c*9/5+32
	default: fmt.Fprintln(os.Stderr, "Unit must be c, f, or k"); os.Exit(1)
	}
	fmt.Printf("  🌡  %.2f°C  =  %.2f°F  =  %.2fK\n", c, f, k)
}
