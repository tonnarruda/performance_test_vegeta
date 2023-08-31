package main

import (
	"fmt"
	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

// func main() {
// 	startTime := time.Now()
// 	rate := vegeta.Rate{Freq: 10, Per: time.Second} // taxa de solicitações por segundo
// 	duration := 10 * time.Second                    // duração do teste

// 	body := []byte(`{
// 		"trigger":"notification:delivery:user",
// 		"app":{
// 		   "_id":"5fed024df3dd8b000cd7ee99"
// 		},
// 		"version":"v1.1",
// 		"timestamp":1682435553.313,
// 		"destination":{
// 		   "type":"whatsapp",
// 		   "integrationId":"123",
// 		   "destinationId":"558888888888"
// 		},
// 		"isFinalEvent":false,
// 		"externalMessages":[
// 		   {
// 			  "id":"gBGHVRKYh1IXjwIJi08v7k2Pxm_H"
// 		   }
// 		],
// 		"notification":{
// 		   "_id":"9d633118-27d3-437d-9d81-621b07903e06"
// 		},
// 		"matchResult":{
// 		   "appUser":{
// 			  "_id":"2dd8e437e7c9c256c03202b3",
// 			  "conversationStarted":true
// 		   },
// 		   "conversation":{
// 			  "_id":"9513bb85a781f5419c903786"
// 		   }
// 		}
// 	 }`) // Substitua pelo JSON que você deseja enviar

// 	// Criar um target estático com a URL alvo
// 	target := vegeta.NewStaticTargeter(vegeta.Target{
// 		Method: "POST",
// 		URL:    "https://edwiges.olaisaac.dev/v1/webhook/sunco", // Substitua pela URL do seu endpoint
// 		Header: map[string][]string{
// 			"X-API-Key":    {"SM9q-0nvmljZzXYKw5Ty1jIL-mgFC9V1rzOk0Dot_YYadWUOwgZ9BGuzY_VFqkyJ9w4DgbS10u8-LJ70IP7QYg"}, // Substitua pelo token de autorização
// 			"Content-Type": {"application/json"},
// 		},
// 		Body: body,
// 	})
// 	// Criar um Atacante com a taxa de ataque configurada
// 	attacker := vegeta.NewAttacker()

// 	var metrics vegeta.Metrics
// 	successCount := 0
// 	failureCount := 0

// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for res := range attacker.Attack(target, rate, duration, "PerformanceTest") {
// 		metrics.Add(res)

// 		if res.Error != "" {
// 			log.Printf("Resposta: %s, Código: %d, Erro: %v\n", res.Timestamp, res.Code, res.Error)
// 			failureCount++
// 		} else {
// 			log.Printf("StatusCode: %d\n", res.Code)
// 			successCount++
// 		}
// 	}

// 	metrics.Close()
// 	elapsedTime := time.Since(startTime)
// 	fmt.Printf("Tempo total de execução: %s\n", elapsedTime)

// 	fmt.Printf("Requisições com sucesso: %d\n", successCount)
// 	fmt.Printf("Requisições com falha: %d\n", failureCount)
// 	fmt.Printf("Total de requisições: %d\n", successCount+failureCount)
// 	fmt.Printf("Requisições por segundo (média): %.2f\n", float64(successCount+failureCount)/elapsedTime.Seconds())
// 	fmt.Printf("Sucesso: %.2f%%\n", metrics.Success*100)
// 	fmt.Printf("Média de latência: %s\n", metrics.Latencies.Mean)
// 	fmt.Printf("Percentil 50 de latência: %s\n", metrics.Latencies.P50)
// 	fmt.Printf("Percentil 90 de latência: %s\n", metrics.Latencies.P90)
// 	fmt.Printf("Valor médio de bytes de entrada registrados: %.2f\n", metrics.BytesIn.Mean)
// 	fmt.Printf("Tempo total de execução: %.2f segundos\n", elapsedTime.Seconds())
// }

func main() {
	startTime := time.Now()
	rate := vegeta.Rate{Freq: 200, Per: time.Second} // taxa de solicitações por segundo
	duration := 10 * time.Second                     // duração do teste

	body := []byte(`{
		"messageData": {
			"branch": {
				"id": "111"
			},
			"student": {
				"name": "Tiaguinho 2",
				"id": "123"
			},
			"guardians": [
				{
					"id": "fe1635c1-b257-4482-abda-e27d07bb8052",
					"name": "Tiago Oliveira",
					"email": "tiago.oliveira@isaac.com.br"
				}
			],
			"message": {
				"subject": "Teste de envio",
				"content": "Olá"
			}
		}
	}`) // Substitua pelo JSON que você deseja enviar

	// Criar um target estático com a URL alvo
	target := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "https://isaac.classapp.com.br/isaac/v1/message/create", // Substitua pela URL do seu endpoint
		Header: map[string][]string{
			"Authorization": {"Bearer f8ea598be41640a93be2"}, // Substitua pelo token de autorização
			"Content-Type":  {"application/json"},
		},
		Body: body,
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
		time.Sleep(150 * time.Millisecond)
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
