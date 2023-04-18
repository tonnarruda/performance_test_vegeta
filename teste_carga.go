package main

import (
	"fmt"
	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second} // taxa de solicitações por segundo
	duration := 1 * time.Second                      // duração do teste

	// Criar um Targeter estático com a URL alvo
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "https://restful-booker.herokuapp.com/booking",
	})

	// Criar um Atacante com a taxa de ataque configurada
	attacker := vegeta.NewAttacker()

	// Realizar o ataque e coletar as métricas
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "MeuNomeDeTeste") {
		metrics.Add(res)

		// Verificar se há erro na resposta e logar apenas se houver
		if res.Error != "" {
			log.Printf("Response: %s, Code: %d, Error: %v\n", res.Timestamp, res.Code, res.Error)
		} else {
			log.Printf("Response: %s, Code: %d\n", res.Timestamp, res.Code)
		}
	}
	metrics.Close()

	// Exibir as métricas
	fmt.Printf("Requests: %d\n", metrics.Requests)
	fmt.Printf("Success: %0.2f%%\n", metrics.Success*100)
	fmt.Printf("Latency mean: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Latency 50th percentile: %s\n", metrics.Latencies.P50)
	fmt.Printf("Latency 90th percentile: %s\n", metrics.Latencies.P90)
	fmt.Printf("Valor médio de bytes de entrada registrados: %.2f\n", metrics.BytesIn.Mean)
	fmt.Printf("Bytes Out mean: %.2f\n", metrics.BytesOut.Mean)
}