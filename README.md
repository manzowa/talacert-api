# 📄 TalaCert API

TalaCert est une API de vérification de documents officiels (certificats, diplômes, attestations) basée sur GIN et GORM.

L'objectif est de permettre à toute personne ou organisation de vérifier rapidement l'authenticité d’un document via un identifiant unique ou un QR code.

---

## 🚀 Fonctionnalités

* ✅ CRUD complet des documents
* 🔍 Vérification d’authenticité via API
* 🔐 Gestion de hash (sécurité des documents)
* 📱 Intégration possible avec QR Code
* ⛓ Compatible blockchain (hash + transaction)
* 📊 API REST avec documentation automatique (Swagger)

---

## 🧱 Stack technique

* GO
* GIN
* GORM ORM
* MySQL /MariaDB 

---

## ⚙️ Installation

### 1. Cloner le projet

```bash
git clone https://github.com/ton-repo/talacert-api.git
cd talacert-api
```

---

### 2. Installer les dépendances

```bash
go mod download

```
Ou simplement

```
go mod tidy

```

---

### 3. Configuration de la base de données

Modifier le fichier `.env` :

```env
APP_ENV=development
GIN_MODE=debug
PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_NAME=talacert
DB_USER=root
DB_PASSWORD=root

# JWT Configuration
JWT_ACCESS_SECRET=4f8c9d7a2b1e5f63c8a91d4e7f2b6c5a9d3e8f1b7c4a6d2e9f5c1a8b3d7e4f6
JWT_ACCESS_EXPIRATION=3600
JWT_REFRESH_SECRET=a7e3d9f5b1c8e4f6d2a9c7b3f1e5d8a4c6f2b9e7d1a3f5c8e4b6d9a2c7f1e5b
JWT_REFRESH_EXPIRATION=604800

# Administration Configuration 
ADMIN_DEFAULT_USERNAME=default
ADMIN_DEFAULT_EMAIL=default@talacert.local
ADMIN_DEFAULT_PASSWORD=Default@123456
```

---

### 4. Démarrer le serveur (Exécuter le projet)

```bash

.\run

```

Ou sous Windows :

```bash
start run.bat or run 
```

---

## 🌐 Accès à l’API

* Swagger UI :
```
http://127.0.0.1:8080/swagger/doc
```


```
http://127.0.0.1:8080/api
```

---

## 📡 Endpoints principaux

### 📄 Documents

| Méthode | Endpoint                        |
| ------- | ------------------------------- |
| POST    | /auth/login                     |
| POST    | /auth/refresh                   |
| POST    | /auth/logout                    |
| GET     | /auth/me                        |
| POST    | /api/documents                  |
| GET     | /api/documents                  |
| GET     | /api/documents/:document_id     |
| PUT     | /api/documents/:document_id     |
| DELETE  | /api/documents/:document_id     |
| GET     | /api/documents/hash/:hash       |
| GET     | /api/documents/verify           |

---

### 🔍 Vérification d’un document

```
POST /api/documents/verify
```

#### Body :

```json
{
  "document_id": "CERT-2025-0001"
}
```

#### Réponse :

```json
{
  "success": true,
  "message": "Document Valid",
  "data": {
    "status": "valid",
    "data": {
      "owner": "Jean Tshibangu",
      "type": "Certificat",
      "issuer": "Université de Kinshasa"
    }
  }
}
```

---

## 🗄 Structure du projet

```
internal/
 ├── Config/
 │    └── config.go
 │    └── database.go
 │    └── env.go
 │
 ├── constants/
 │    └── document.go
 │    └── roles.go
 │
 ├── dto/
 │    └── auth.go 
 │    └── document.go
 |    └── user.go
 │
 ├── handlers/
 │    └── auth_handler.go
 │    └── document_handler.go
 │    └── user_handler.go
 │    └── handler.go
 │
 ├── logger/
 │    └── logger.go
 │
 ├── middleware/
 │    └── access_log_middleware.go
 │    └── error_log_middleware.go
 │    └── ownership_middleware.go
 │
 ├── models/
 │    └── document.go
 │    └── document_sequence.go
 │    └── refresh_token.go
 │    └── user.go
 │
 ├── repositories/
 │    └── auth_repository.go
 │    └── document_repository.go
 │    └── document_sequence_repository.go
 │    └── user_repository.go
 │
 ├── routes/
 │    └── api_routes.go
 │    └── auth_routes.go
 │    └── auth_protected_routes.go
 │    └── document_routes.go
 │    └── swagger_routes.go
 │    └── user_routes.go
 │
 ├── seed/
 │    └── seed_manager.go
 │    └── user_seed.go
 │
 ├── services/
 │    └── auth_service.go
 │    └── document_service.go
 │    └── user_service.go
 │
 ├── utils/
 │    └── hash.go
 │    └── document_id_generator.go
 │    └── response.go
 │
```

---

## 🔐 Sécurité

* Hash SHA256 pour l’intégrité des documents
* Possibilité d’ajouter JWT Authentication
* Support de signature numérique
* Intégration future avec blockchain

---

## ⛓ Intégration Blockchain (optionnel)

Chaque document peut contenir :

* hash du document
* transaction blockchain (Ethereum, Polygon, etc.)

---

## 📌 Statuts des documents

* `valid`
* `invalid`

---

## 🧪 Tests

```bash
make test
```

---

## 🛠️ Utilisation avec Makefile

### Démarrer le serveur
make run

### Compiler
make build

### Tests
make test

### Swagger
make swagger

---

## 🚀 Roadmap

* [ ] Authentification JWT
* [ ] Upload de fichiers PDF
* [ ] Génération QR Code
* [ ] Intégration Blockchain
* [ ] Dashboard Admin
* [ ] Notifications (email/SMS)

---

## 🤝 Contribution

Les contributions sont les bienvenues !

1. Fork le projet
2. Crée une branche (`feature/...`)
3. Commit
4. Push
5. Pull Request

---

## 📄 Licence

MIT

---

## 👨‍💻 Auteur

Christian Shungu
Développeur Full Stack & DevOps

---

## 💡 Vision

TalaCert vise à devenir une plateforme de référence pour la vérification des documents officiels au congo-kinshasa et dans le monde.

---
