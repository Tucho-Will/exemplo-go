FROM golang:alpine

# Seta variaveis para compilar no ambiente correto
ENV GOOS=linux \
    GOARCH=amd64

# Diretorio para compilar
WORKDIR /build

# Copia e compila o go mod
COPY go.mod .
RUN go mod download

# Copia os codigos pro container
COPY . .

# Build da aplicação
RUN go build framework/cmd/server/server.go

# Diretorio para versão ja compilada
WORKDIR /dist

# Copia os binarios para diretorio
RUN cp /build/server .

# Expor porta quando necessario
# EXPOSE 3000

# comando que sera executado quando container rodar
CMD ["/dist/server"]