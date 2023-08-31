import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '5', target: 2 },   // Ramp up to 100 VUs in 30 seconds
    { duration: '5', target: 2 },   // Maintain 100 VUs for 15 seconds
    { duration: '5', target: 0 },     // Drop to 0 VUs in 15 seconds
  ],
  thresholds: {
    http_req_duration: ['p(50)<100', 'p(90)<200'],  // Example latency thresholds
  },
};

export default function () {
  const url = "https://edwiges.olaisaac.dev/v1/webhook/sunco";
  const payload = JSON.stringify({
    trigger: "notification:delivery:user",
    app: {
      _id: "5fed024df3dd8b000cd7ee99",
    },
    version: "v1.1",
    timestamp: 1682435553.313,
    destination: {
      type: "whatsapp",
      integrationId: "123",
      destinationId: "558888888888",
    },
    isFinalEvent: false,
    externalMessages: [
      {
        id: "gBGHVRKYh1IXjwIJi08v7k2Pxm_H",
      },
    ],
    notification: {
      _id: "9d633118-27d3-437d-9d81-621b07903e06",
    },
    matchResult: {
      appUser: {
        _id: "2dd8e437e7c9c256c03202b3",
        conversationStarted: true,
      },
      conversation: {
        _id: "9513bb85a781f5419c903786",
      },
    },
  });
  const headers = {
    "x-api-key":
      "SM9q-0nvmljZzXYKw5Ty1jIL-mgFC9V1rzOk0Dot_YYadWUOwgZ9BGuzY_VFqkyJ9w4DgbS10u8-LJ70IP7QYg",
    "Content-Type": "application/json",
  };

  http.post(url, payload, { headers: headers });

  let success = check(response, {
    'is success': (r) => r.status === 200,
  });

  sleep(1);
}

export function handleSummary(data) {
    console.log('Tempo total de execução:', data.duration);
    console.log('Requisições com sucesso:', data.results.success);
    console.log('Requisições com falha:', data.results.failures);
    console.log('Total de requisições:', data.results.total);
    console.log('Requisições por segundo (média):', data.metrics.rps);
    console.log('Sucesso:', data.metrics.success + '%');
    console.log('Média de latência:', data.metrics.latencies.mean);
    console.log('Percentil 50 de latência:', data.metrics.latencies.p50);
    console.log('Percentil 90 de latência:', data.metrics.latencies.p90);
    console.log('Valor médio de bytes de entrada registrados:', data.metrics.bytes_in.mean);
    console.log('Tempo total de execução (segundos):', data.metrics.duration_seconds);
  }