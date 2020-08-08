package main

 

import (
    "net/http"
    "strconv"
    "strings"
)

 

var numeroaPalabra = map[int]string{
    0:  "cero",
    1:  "uno",
    2:  "dos",
    3:  "tres",
    4:  "cuatro",
    5:  "cinco",
    6:  "seis",
    7:  "siete",
    8:  "ocho",
    9:  "nueve",
    10: "diez",
    11: "once",
    12: "doce",
    13: "trece",
    14: "catorce",
    15: "quince",
    16: "dieciseis",
    17: "diecisiete",
    18: "dieciocho",
    19: "diecinueve",
    20: "veinte",
    30: "treinta",
    40: "cuarenta",
    50: "cincuenta",
    60: "sesenta",
    70: "setenta",
    80: "ochenta",
    90: "noventa",
}

 

func convierte1a100(n int) (w string) {
    if n < 20 {
        w = numeroaPalabra[n]
        return
    }

 

    r := n % 10
    if r == 0 {
        w = numeroaPalabra[n]
    } else {
        w = numeroaPalabra[n-r] + " y " + numeroaPalabra[r]
    }
    return
}

 

func fibonacci(n int) int {

 

    a := 0
    b := 1

 

    for i := 0; i < n; i++ {

 

        temp := a
        a = b
        b = temp + a
    }
    return a
}

 

func isLeap(texto string) bool {
    year, _ := strconv.Atoi(texto)
    return year%400 == 0 || year%4 == 0 && year%100 != 0

 

}
func sayHello(w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path
    message = strings.TrimPrefix(message, "/")

 

    numero, _ := strconv.Atoi(message)
    var fibo = fibonacci(numero)
    s := strconv.Itoa(fibo)
    if isLeap(message) {

 

        message = message + " es Bisiesto y el numero es " + convierte1a100(numero) + " fibonacci " + s
    } else {
        message = message + " No es Bisiesto y el numero es " + convierte1a100(numero) + " fibonacci " + s
    }

 

    w.Write([]byte(message))
}

 

func main() {
    http.HandleFunc("/", sayHello)

 

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
