Event Check-in System (Go + PostgreSQL)

Esse projeto é um microsserviço feito em Go (Golang) para controle de acesso e check-in de convidados em eventos.

A ideia principal aqui é simples: evitar fraude e garantir que um ingresso não seja usado mais de uma vez — mesmo em cenários com várias catracas ou fiscais operando ao mesmo tempo.

Como isso funciona?

Em eventos grandes, pode acontecer de duas pessoas tentarem validar o mesmo QR Code praticamente ao mesmo tempo. Se o sistema não tratar isso direito, rola uma race condition e o ingresso pode acabar sendo aceito duas vezes.

Pra evitar isso, a API usa pessimistic locking (row-level lock) direto no PostgreSQL com SELECT ... FOR UPDATE.

Na prática:

Quando chega uma requisição de check-in, o registro do convidado é “travado” no banco
Qualquer outra tentativa de usar o mesmo QR fica esperando
A primeira requisição finaliza e atualiza o status pra CHECKED_IN
As próximas já recebem resposta de ingresso inválido/usado

Isso garante consistência e evita entrada duplicada.

Tecnologias
Go 1.21+ (com Gin para as rotas)
PostgreSQL 15
GORM (ORM + migrações automáticas)
Docker / Docker Compose
godotenv (variáveis de ambiente)

Como rodar o projeto

Tá tudo dockerizado, então é bem tranquilo subir.

Pré-requisitos
Docker
Docker Compose
Passos

1. Clonar o repositório

git clone https://github.com/SEU-USUARIO/controle-checkin.git
cd controle-checkin

2. Criar o arquivo .env na raiz

DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=eventos
DB_PORT=5432
API_PORT=8080

3. Subir os containers

docker compose up -d --build

A API vai estar disponível em:

http://localhost:8080
  Endpoints
  Criar convidado

POST /api/convidados

Cadastra um convidado em um evento.

{
  "nome": "Mateus Silva",
  "cpf": "12345678910",
  "evento_id": 1,
  "codigo_qr": "QR-999"
}

Listar convidados

GET /api/convidados

Retorna todos os convidados com o status do ingresso.

Check-in (catraca)

POST /api/checkin

Valida o QR Code.

{
  "codigo_qr": "QR-999"
}

Respostas:

200 → Acesso liberado
409 → Ingresso já utilizado
404 → Ingresso não encontrado

Estrutura
├── cmd/
│   └── api/          # main da aplicação
├── internal/
│   ├── database/     # conexão com DB (com retry)
│   ├── handlers/     # handlers das rotas
│   └── models/       # structs / models do GORM
├── Dockerfile
├── docker-compose.yml
└── .env              # variáveis (não versionar)
