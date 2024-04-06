/*
package main
import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	var cli string
	var clients int
	var opt string

	flag.StringVar(&cli, "cli", "curl", "Que cli se va a ejecutar puede ser curl , az ,gcloud, aws ")
	flag.IntVar(&clients, "clients",2, "Ambiente por defecto d")
	flag.StringVar(&opt, "opt", "changeit", "optional 1 ")
	flag.Parse()

	allResults := make([]string, 0)
	results := make(chan []string, clients) 
    var wg sync.WaitGroup
	for i := 0; i < clients; i ++ {
	    wg.Add(1)
        go fmt.Println(i)
    }
	wg.Wait()

    fmt.Println(allResults)
    fmt.Println(results)

} */
package main

import (
	"flag"
	"fmt"
	"time"
	"sync"
)

func main() {
	// Definir flags para recibir la cantidad de clientes
	var numClients int
	flag.IntVar(&numClients, "clients", 3, "Número de clientes")
	flag.Parse()

	// Crear un WaitGroup para esperar que todos los clientes terminen
	var wg sync.WaitGroup
	wg.Add(numClients)

	// Canal para recibir resultados de los clientes
	resultados := make(chan int)

	// Ejecutar clientes concurrentes
	for i := 0; i < numClients; i++ {
		go func(id int) {
			// Simulación de una operación de cliente
			resultado := simulateClientOperation(id)
			resultados <- resultado
			wg.Done() // Marcar la finalización del cliente
		}(i)
	}

	// Esperar a que todos los clientes terminen
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Recoger resultados
	total := 0
	for res := range resultados {
		total += res
	}

	// Imprimir resultado total
	fmt.Printf("Total procesado por todos los clientes: %d\n", total)
}

func simulateClientOperation(clientID int) int {
	// Simulación de una operación de cliente (en este caso, simplemente retornamos el ID del cliente)
    fmt.Println("empezando con ", clientID)
    //Simulamos un sleep , pero aca puede correr un curl pegarle a un api usar az cli o google cloud cli o azure cli
    // para procesos que demoran y requermos que se ejecuten en paralelo no en secuecuencial
    time.Sleep(10*time.Second)
    fmt.Println("Finalizado con ", clientID)
	return clientID
}

