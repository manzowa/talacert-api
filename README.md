# 📄 TalaCert API

TalaCert est une API de vérification de documents officiels (certificats, diplômes, attestations) basée sur **Go, Gin et GORM**.

L'objectif est de permettre à toute personne ou organisation de vérifier rapidement l'authenticité d’un document via un identifiant unique, un hash ou un QR Code.

---

## 🚀 Fonctionnalités

* ✅ CRUD complet des documents
* 🔍 Vérification d’authenticité via API
* 🔐 Gestion de hash pour l’intégrité des documents
* 📱 Intégration possible avec QR Code
* ⛓️ Compatible blockchain (hash + transaction)
* 📊 API REST avec documentation automatique via Swagger
* 🔑 Authentification JWT avec access token et refresh token
* 🧪 Tests automatisés
* 🛠️ Commandes Makefile pour simplifier le développement

---

## 🧱 Stack technique

* **Go**
* **Gin**
* **GORM**
* **MySQL / MariaDB**
* **JWT**
* **Swagger**
* **Make**

---

## 📁 Structure générale

Le point d'entrée de l'application est :

```text
cmd/main.go
```

La structure principale du projet est organisée comme suit :

```text
.
├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   ├── constants/
│   ├── dto/
│   ├── handlers/
│   ├── logger/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── seed/
│   ├── services/
│   └── utils/
│
├── bin/
│   └── talacert-api
│
├── .env
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

Le dossier `bin/` est généré automatiquement lors de la compilation et ne doit généralement pas être versionné dans Git.

---

# ⚙️ Installation

## 1. Prérequis

Avant de commencer, assurez-vous d'avoir installé :

* **Go**
* **Make**
* **MySQL ou MariaDB**
* **Git**

Vérifier l'installation de Go :

```bash
go version
```

Vérifier Make :

```bash
make --version
```

---

## 2. Cloner le projet

```bash
git clone https://github.com/ton-repo/talacert-api.git
cd talacert-api
```

---

## 3. Installer les dépendances

Le projet utilise Go Modules.

Avec Make :

```bash
make install
```

Cette commande exécute :

```bash
go mod download
```

Vous pouvez également utiliser directement :

```bash
go mod download
```

---

## 4. Nettoyer les dépendances

Pour synchroniser et nettoyer les dépendances du projet :

```bash
make tidy
```

Cette commande exécute :

```bash
go mod tidy
```

---

# 🗄️ Configuration de la base de données

Créer ou modifier le fichier `.env` :

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
JWT_ACCESS_SECRET=your-access-secret
JWT_ACCESS_EXPIRATION=3600

JWT_REFRESH_SECRET=your-refresh-secret
JWT_REFRESH_EXPIRATION=604800

# Administration Configuration
ADMIN_DEFAULT_USERNAME=default
ADMIN_DEFAULT_EMAIL=default@talacert.local
ADMIN_DEFAULT_PASSWORD=change-me
```

> ⚠️ **Important :** les secrets JWT et les identifiants administrateur doivent être remplacés par des valeurs sécurisées. Ne committez jamais votre véritable fichier `.env` dans le dépôt Git.

Il est recommandé d'ajouter `.env` au `.gitignore`.

---

# ▶️ Démarrage du projet

Le `Makefile` définit `cmd/main.go` comme point d'entrée de l'application.

## Mode développement

Pour démarrer directement l'application :

```bash
make run
```

Cette commande exécute :

```bash
go run cmd/main.go
```

Le serveur sera normalement disponible sur :

```text
http://127.0.0.1:8080
```

---

# 🏗️ Compilation

Pour compiler l'application :

```bash
make build
```

Le binaire sera généré dans :

### Linux / macOS

```text
bin/talacert-api
```

### Windows

```text
bin/talacert-api.exe
```

La commande utilise automatiquement l'extension `.exe` sous Windows.

Le comportement est défini dans le `Makefile` :

```makefile
APP_NAME=talacert-api
MAIN_PATH=cmd/main.go
```

---

# ▶️ Exécution du binaire

Après compilation, sous Linux/macOS :

```bash
./bin/talacert-api
```

Sous Windows :

```powershell
.\bin\talacert-api.exe
```

---

# 🧹 Nettoyage

Pour supprimer le dossier `bin/` :

```bash
make clean
```

Sous Linux/macOS, le Makefile utilise :

```bash
rm -rf bin
```

Sous Windows :

```cmd
if exist bin rmdir /s /q bin
```

Le `Makefile` détecte automatiquement Windows grâce à :

```makefile
ifeq ($(OS),Windows_NT)
```

---

# 🧪 Tests

Pour exécuter tous les tests du projet :

```bash
make test
```

Cette commande exécute :

```bash
go test ./...
```

Vous pouvez également utiliser directement :

```bash
go test ./...
```

---

# 🧹 Formatage du code

Pour formater automatiquement le code Go :

```bash
make fmt
```

Cette commande exécute :

```bash
go fmt ./...
```

---

# 🔎 Vérification du code

Pour exécuter l'analyse statique de Go :

```bash
make vet
```

Cette commande exécute :

```bash
go vet ./...
```

Il est recommandé d'exécuter `make fmt` et `make vet` avant de pousser des modifications.

---

# 📚 Documentation Swagger

La documentation Swagger est générée avec :

```bash
make swagger
```

Cette commande exécute :

```bash
swag init -g cmd/main.go
```

Le fichier `cmd/main.go` est donc utilisé comme point d'entrée pour la génération de la documentation Swagger.

Après avoir démarré l'application, Swagger UI est accessible à :

```text
http://127.0.0.1:8080/swagger/doc
```

---

# 🛠️ Commandes Makefile

Le projet fournit les commandes suivantes :

| Commande       | Description                            |
| -------------- | -------------------------------------- |
| `make`         | Démarre l'application (`make run`)     |
| `make all`     | Démarre l'application                  |
| `make install` | Télécharge les dépendances             |
| `make tidy`    | Nettoie et synchronise les dépendances |
| `make build`   | Compile l'application dans `bin/`      |
| `make run`     | Lance l'application avec `go run`      |
| `make test`    | Exécute les tests                      |
| `make swagger` | Génère la documentation Swagger        |
| `make clean`   | Supprime le dossier `bin/`             |
| `make fmt`     | Formate le code Go                     |
| `make vet`     | Analyse le code avec `go vet`          |

---

# 🌐 Accès à l’API

## Health Check

```http
GET /health
```

Exemple :

```text
http://127.0.0.1:8080/health
```

---

## Swagger UI

```text
http://127.0.0.1:8080/swagger/doc
```

---

## API

```text
http://127.0.0.1:8080/api
```

---

# 📡 Endpoints principaux

## 🔐 Authentification

| Méthode | Endpoint               |
| ------- | ---------------------- |
| POST    | `/api/v1/auth/login`   |
| POST    | `/api/v1/auth/refresh` |
| POST    | `/api/v1/auth/logout`  |
| GET     | `/api/v1/auth/me`      |

---

## 📄 Documents

| Méthode | Endpoint                                |
| ------- | --------------------------------------- |
| POST    | `/api/v1/documents`                     |
| GET     | `/api/v1/documents`                     |
| GET     | `/api/v1/documents/:document_id`        |
| PATCH   | `/api/v1/documents/:document_id`        |
| DELETE  | `/api/v1/documents/:document_id`        |
| GET     | `/api/v1/documents/by-hash/:hash`       |
| GET     | `/api/v1/documents/verify`              |

Tous les endpoints utilisent uniformément la version :

```text
/api/v1
```

---

# 🔍 Vérification d’un document

La vérification d'un document s'effectue à partir de son identifiant.

### Endpoint

```http
GET /api/v1/documents/verify
```

### Exemple

```http
GET /api/v1/documents/verify
```
```Body
{
  "document_id": "CERT-2025-0001"
}
```

### Réponse

```json
{
  "success": true,
  "message": "Document valid",
  "data": {
    "status": "valid",
    "document": {
      "owner": "Jean Tshibangu",
      "type": "Certificat",
      "issuer": "Université de Kinshasa"
    }
  }
}
```

---

# 🔐 Recherche par hash

Un document peut être retrouvé à partir de son hash.

### Endpoint

```http
GET /api/v1/documents/by-hash/:hash
```

### Exemple

```http
GET /api/v1/documents/by-hash/9f86d081884c7d659a2feaa0c55ad015
```

Cette fonctionnalité permet notamment de vérifier l'intégrité d'un document et pourra être utilisée dans le cadre d'une future intégration blockchain.

---

# ✏️ Modification d’un document

Les modifications partielles utilisent `PATCH`.

### Endpoint

```http
PATCH /api/v1/documents/:document_id
```

### Exemple

```http
PATCH /api/v1/documents/CERT-2025-0001
```

### Body

```json
{
  "owner": "Jean Tshibangu",
  "issuer": "Université de Kinshasa",
  "status": "valid"
}
```

---

# 🔐 Sécurité

TalaCert utilise ou prévoit les mécanismes suivants :

* 🔒 Hash SHA-256 pour l'intégrité des documents
* 🔑 Authentification JWT
* ♻️ Access tokens et refresh tokens
* 🛡️ Contrôle d'accès aux ressources
* 👤 Gestion de l'ownership des documents
* ✍️ Possibilité d'ajouter une signature numérique
* ⛓️ Intégration blockchain future

Les secrets JWT doivent être stockés dans les variables d'environnement.

---

# ⛓️ Intégration Blockchain

Chaque document pourra éventuellement être associé à des informations blockchain :

```text
document
├── hash
├── blockchain
├── transaction_hash
└── network
```

La blockchain pourra servir à enregistrer une preuve d'existence ou d'intégrité du document sans avoir besoin de stocker le document lui-même sur la blockchain.

---

# 📌 Statuts des documents

Statuts actuellement supportés :

```text
valid
invalid
```

Statuts pouvant être ajoutés ultérieurement :

```text
pending
revoked
expired
```

---

# 🚀 Roadmap

* [x] Authentification JWT
* [x] CRUD des documents
* [x] Vérification des documents
* [x] Recherche par hash
* [x] Makefile pour le développement
* [x] Documentation Swagger
* [ ] Upload de fichiers PDF
* [ ] Génération de QR Code
* [ ] Signature numérique
* [ ] Intégration Blockchain
* [ ] Dashboard Admin
* [ ] Notifications email/SMS
* [ ] Gestion des documents révoqués
* [ ] Gestion de l'expiration des documents

---

# 🤝 Contribution

Les contributions sont les bienvenues !

### 1. Fork du projet

### 2. Créer une branche

```bash
git checkout -b feature/ma-feature
```

### 3. Développer la fonctionnalité

### 4. Formater et vérifier le code

```bash
make fmt
make vet
```

### 5. Exécuter les tests

```bash
make test
```

### 6. Commit

```bash
git commit -m "feat: ajout de ma fonctionnalité"
```

### 7. Push

```bash
git push origin feature/ma-feature
```

### 8. Créer une Pull Request

---

# 📄 Licence

MIT

---

# 👨‍💻 Auteur

**Christian Shungu**

Développeur Full Stack & DevOps

---

# 💡 Vision

TalaCert vise à devenir une plateforme de référence pour la **vérification fiable et sécurisée des documents officiels**, d'abord en République démocratique du Congo, puis à l'international.

L'objectif à long terme est de fournir une infrastructure permettant aux universités, administrations, entreprises et organismes émetteurs de publier des documents vérifiables, tout en permettant aux tiers de contrôler leur authenticité rapidement via une API, un identifiant unique, un hash ou un QR Code.
