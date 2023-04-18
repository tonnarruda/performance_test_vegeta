package main

import (
	"fmt"
	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	startTime := time.Now()
	rate := vegeta.Rate{Freq: 100, Per: time.Second} // taxa de solicitações por segundo
	duration := 15 * time.Second                     // duração do teste

	// Criar um target estático com a URL alvo
	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "https://restful-booker.herokuapp.com/booking",
	})

	// Criar um Atacante com a taxa de ataque configurada
	attacker := vegeta.NewAttacker()

	// Realizar o ataque e coletar as métricas
	var metrics vegeta.Metrics
	for res := range attacker.Attack(target, rate, duration, "PerformanceTest") {
		metrics.Add(res)

		// Verificar se há erro na resposta e logar apenas se houver
		if res.Error != "" {
			log.Printf("Resposta: %s, Código: %d, Erro: %v\n", res.Timestamp, res.Code, res.Error)
		} else {
			log.Printf("StatusCode: %d\n", res.Code)
		}
	}
	metrics.Close()
	elapsedTime := time.Since(startTime) // Calcular o tempo decorrido
	fmt.Printf("Tempo total de execução: %s\n", elapsedTime)

	// Exibir as métricas
	fmt.Printf("Requisições: %d\n", metrics.Requests)
	fmt.Printf("Sucesso: %.2f%%\n", metrics.Success*100)
	fmt.Printf("Média de latência: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Percentil 50 de latência: %s\n", metrics.Latencies.P50)
	fmt.Printf("Percentil 90 de latência: %s\n", metrics.Latencies.P90)
	fmt.Printf("Valor médio de bytes de entrada registrados: %.2f\n", metrics.BytesIn.Mean)
	fmt.Printf("Tempo total de execução: %.2f segundos\n", elapsedTime.Seconds())
}
