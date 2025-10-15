# deploy-go-gin.ps1
# Script para subir Minikube, Go-Gin, Postgres e PgAdmin usando cluster.yaml
# e abrir Go-Gin e PgAdmin automaticamente

Write-Host "⏳ Iniciando Minikube..."
minikube start

Write-Host "⚡ Configurando Docker para usar o Minikube..."
& minikube -p minikube docker-env --shell powershell | Invoke-Expression

Write-Host "📦 Buildando imagem do Go-Gin dentro do Minikube..."
docker build -t go-gin:latest .

Write-Host "📤 Aplicando Kubernetes YAML (cluster.yaml)..."
kubectl apply -f cluster.yaml

Write-Host "⏱ Aguardando pods ficarem Running..."
$podsReady = $false
while (-not $podsReady) {
    $status = kubectl get pods --no-headers | Select-String "Running"
    # Espera 3 pods: postgres, pgadmin, go-gin
    if ($status.Count -ge 3) {
        $podsReady = $true
    } else {
        Write-Host "Ainda iniciando pods..."
        Start-Sleep -Seconds 5
    }
}

Write-Host "✅ Todos os pods estão Running!"

# Port-forward da app Go-Gin
Write-Host "🔀 Iniciando port-forward para Go-Gin (localhost:8080)..."
Start-Process powershell -ArgumentList "kubectl port-forward deployment/go-gin-app 8080:8080"

# Abrir PgAdmin no navegador usando minikube service
Write-Host "🌐 Abrindo PgAdmin no navegador..."
Start-Process powershell -ArgumentList "minikube service pgadmin --url"

Write-Host ""
Write-Host "💡 Teste Go-Gin no Postman:"
Write-Host "GET http://localhost:8080/ping"
Write-Host "💡 PgAdmin abrirá automaticamente no navegador."
Write-Host "Obs: Mantenha o terminal do port-forward aberto enquanto testa a API."
