# 🌐 A Rota das Coisas: Sistema IoT Distribuído

🚀 Este projeto foi desenvolvido como solução para o **Problema 1 da
disciplina TEC502**.\
O objetivo é reduzir o **alto acoplamento** entre dispositivos
industriais (sensores/atuadores) e aplicações, utilizando um **Serviço
de Integração (Broker)** centralizado.

------------------------------------------------------------------------

## 🧠 Conceito da Solução

### 📡 Telemetria (UDP)

-   Fluxo contínuo de dados
-   ⚡ Prioridade: **velocidade**
-   Sem garantia de entrega

### 🔐 Comando e Controle (TCP)

-   Ações críticas (ex: abrir/fechar porta)
-   ✅ Entrega garantida
-   🔒 Integridade assegurada

------------------------------------------------------------------------

## ⚙️ Tecnologias Utilizadas

-   🐹 Go (Golang)\
-   🌐 net\
-   📦 encoding/json\
-   🔄 sync (Mutex)\
-   🐳 Docker

------------------------------------------------------------------------

## ⚠️ Configuração Manual de IPs

### 🧩 Broker

`net.Dial("tcp", "IP_DO_ATUADOR:8983")`

### 📡 Sensores

`net.ResolveUDPAddr("udp", "IP_DO_BROKER:4042")`

### 🖥️ Cliente

`net.Dial("tcp", "IP_DO_BROKER:4041")`

------------------------------------------------------------------------

## 📁 Estrutura do Projeto

    Codes/
    ├── actuator/             # Serviço Atuador (TCP :8983)
    │   ├── actuator.go       # Lógica de controle e feedback
    │   └── docker-compose.yaml
    ├── broker/               # Serviço de Integração (UDP :4042 | TCP :4041)
    │   ├── broker.go         # Gerenciamento de sensores e roteamento
    │   └── docker-compose.yaml
    ├── clienteFolder/        # Aplicação Cliente (TCP)
    │   ├── client.go         # Interface de visualização e controle
    │   └── docker-compose.yaml
    ├── sensor1/              # Dispositivo Virtual 1 (UDP)
    │   ├── objeto1.go        # Simulação de telemetria
    │   └── docker-compose.yaml
    └── sensor2/              # Dispositivo Virtual 2 (UDP)
        ├── objeto2.go        # Simulação de telemetria
        └── docker-compose.yaml

------------------------------------------------------------------------

## 🚀 Como Executar

### 1️⃣ Broker + Atuador

``` bash
cd Codes/broker && docker-compose up --build -d
cd ../actuator && docker-compose up --build -d
```

### 2️⃣ Sensores

``` bash
cd ../sensor1 && docker-compose up --build -d
cd ../sensor2 && docker-compose up --build -d
```

### 3️⃣ Cliente

``` bash
cd ../clienteFolder && docker-compose run client
```

------------------------------------------------------------------------

## 🔍 Funcionalidades

-   📊 Monitoramento de sensores ativos (últimos 7s)
-   🤖 Automação de segurança (porta automática)
-   🎮 Controle manual via cliente

------------------------------------------------------------------------

## 💡 Resumo

🧩 Arquitetura desacoplada\
⚡ Comunicação eficiente com UDP + TCP\
🐳 Containerização com Docker\
🔐 Controle seguro de dispositivos IoT
