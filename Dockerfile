# ============================================================
# Stage 1 — Build
# ============================================================
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

ARG TARGETOS=linux
ARG TARGETARCH=amd64

# Certificats HTTPS nécessaires pendant le build
RUN apk add --no-cache ca-certificates

# Copier les fichiers de dépendances
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source
COPY . .

# Créer le dossier bin et compiler
RUN mkdir -p bin && \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    go build \
        -trimpath \
        -ldflags="-s -w" \
        -o bin/talacert-api \
        ./cmd/main.go


# ============================================================
# Stage 2 — Runtime minimal
# ============================================================
FROM scratch

WORKDIR /app

# Certificats SSL/TLS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/ca-certificates.crt

# Binaire
COPY --from=builder /app/bin/talacert-api \
    /app/bin/talacert-api

# Port HTTP
EXPOSE 8080

# Utilisateur non-root
USER 65532:65532

# Démarrage
ENTRYPOINT ["/app/bin/talacert-api"]