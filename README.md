🎟️ Event Check-in System (Go + PostgreSQL)

Este projeto é um microsserviço de alta performance desenvolvido em Go (Golang) para o controle de acesso e check-in de convidados em eventos. O foco principal desta aplicação é garantir a integridade dos dados e evitar fraudes de entrada através de mecanismos de concorrência no banco de dados.
Em eventos de grande porte, múltiplas catracas ou fiscais podem tentar validar o mesmo QR Code simultaneamente. Sem a proteção adequada, uma Race Condition (Condição de Corrida) poderia permitir que o mesmo ingresso fosse utilizado mais de uma vez se as requisições chegassem no mesmo milissegundo.
Esta API utiliza a estratégia de Pessimistic Locking (Row-Level Lock) através do comando SELECT ... FOR UPDATE do PostgreSQL.
Quando uma requisição de check-in chega, o banco "trava" a linha daquele convidado específico.
Qualquer outra tentativa de acesso ao mesmo registro fica em fila até que a primeira seja processada.
Isso garante que o status do ingresso seja atualizado para CHECKED_IN de forma atômica, impedindo entradas duplicadas.

🛠️ Tecnologias Utilizadas

Linguagem: Go 1.21+ (utilizando Gin Gonic para roteamento).

Banco de Dados: PostgreSQL 15.

ORM: GORM (com migrações automáticas).

Containerização: Docker & Docker Compose.

Segurança: Variáveis de ambiente com godotenv.

⚙️ Como Executar o Projeto

O projeto está totalmente dockerizado, o que facilita o setup em qualquer ambiente.

Pré-requisitos
Docker e Docker Compose instalados.

Passo a Passo
Clonar o repositório:

Bash
git clone https://github.com/SEU-USUARIO/controle-checkin.git
cd controle-checkin
Configurar Variáveis de Ambiente:
Crie um arquivo .env na raiz do projeto:

Snippet de código
DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=eventos
DB_PORT=5432
API_PORT=8080
Subir os Containers:

Bash
docker compose up -d --build
A API estará disponível em http://localhost:8080.

📌 Endpoints da API

1. Criar Convidado
POST /api/convidados
Cadastra um novo convidado vinculado a um evento.

JSON
{
  "nome": "Mateus Silva",
  "cpf": "12345678910",
  "evento_id": 1,
  "codigo_qr": "QR-999"
}
2. Listar Todos os Convidados
GET /api/convidados
Retorna a lista completa e o status de cada ingresso.

3. Realizar Check-in (Catraca)
POST /api/checkin
Valida o QR Code e bloqueia novas tentativas.

JSON
{
  "codigo_qr": "QR-999"
}
Response 200: Acesso liberado.

Response 409: Conflito (Ingresso já utilizado).

Response 404: Ingresso não encontrado.



📁 Estrutura do Projeto

Plaintext
├── cmd/
│   └── api/          # Ponto de entrada da aplicação
├── internal/
│   ├── database/     # Configuração e conexão com DB (com Retry Logic)
│   ├── handlers/     # Controladores (Lógica de rota e resposta)
│   └── models/       # Definição das structs (GORM Models)
├── Dockerfile        # Build da imagem Alpine (com tzdata)
├── docker-compose.yml # Orquestração da API e PostgreSQL
└── .env              # Variáveis sensíveis (não versionar)



✒️ Autor
Mateus Tassoni - Desenvolvedor de Software