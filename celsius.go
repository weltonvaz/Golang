package main

import "fmt"

func main()  {
    var fahinput float64
    var celsius float64
    fmt.Print("Entre com os graus em Fahrenheit: ")
    fmt.Scanf("%f", &fahinput)
    x := (fahinput - 32)
    celsius = x * 0.555555556
    fmt.Print("Convertendo em Celsius: ")
    fmt.Printf("%.2f\n", celsius )

}
