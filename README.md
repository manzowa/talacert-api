README.md
📄 TalaCert API

TalaCert est une API de vérification de documents officiels (certificats, diplômes, attestations) basée sur Go, Gin et GORM.

L'objectif est de permettre à toute personne ou organisation de vérifier rapidement l'authenticité d'un document via un identifiant unique, un hash ou un QR Code.

🚀 Fonctionnalités
✅ CRUD complet des documents
🔍 Vérification d'authenticité via API
🔐 Gestion de hash pour l'intégrité des documents
📱 Intégration possible avec QR Code
⛓️ Compatible blockchain (hash + transaction)
📊 API REST avec documentation automatique via Swagger
🔑 Authentification JWT avec access token et refresh token
🧪 Tests automatisés
🛠️ Commandes Makefile pour simplifier le développement
🧱 Stack technique
Go
Gin
GORM
MySQL / MariaDB
JWT
Swagger
Make
📁 Structure du projet

Le point d'entrée de l'application est :

cmd/main.go


Structure principale :

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
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md


Le dossier bin/ contient le binaire compilé et ne doit généralement pas être versionné dans Git.

⚙️ Installation
1. Prérequis

Avant de commencer, installez :

Go
Make
MySQL ou MariaDB
Git

Vérifier Go :

go version


Vérifier Make :

make --version

2. Cloner le projet
git clone https://github.com/ton-repo/talacert-api.git
cd talacert-api

3. Installer les dépendances

Avec Make :

make install


Cette commande exécute :

go mod download


Ou directement :

go mod download

4. Synchroniser les dépendances
make tidy


Cette commande exécute :

go mod tidy

🗄️ Configuration

Créer ou modifier le fichier .env :

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


⚠️ Important : remplacez les secrets JWT et les identifiants administrateur par des valeurs sécurisées.

Ne committez jamais votre véritable fichier .env.

Ajoutez-le à .gitignore :

.env

▶️ Démarrage

Le Makefile définit :

APP_NAME := talacert-api
MAIN_PATH := cmd/main.go
BIN_DIR := bin

Mode développement

Pour lancer directement l'application :

make run


Cette commande exécute :

go run cmd/main.go


L'API est normalement disponible sur :

http://127.0.0.1:8080

Compiler et démarrer

Pour compiler puis démarrer le binaire :

make start


La commande effectue :

make build
    ↓
démarrage du binaire


Sous Linux/macOS :

bin/talacert-api


Sous Windows :

bin/talacert-api.exe

Commande par défaut

La commande :

make


exécute :

make all


et all exécute :

make start


Ainsi :

make


permet de compiler puis démarrer l'application.

🏗️ Compilation

Pour compiler uniquement l'application :

make build


Le Makefile utilise :

go build -o bin/talacert-api cmd/main.go

Linux / macOS
bin/talacert-api

Windows
bin/talacert-api.exe

▶️ Exécution du binaire

Après compilation :

Linux / macOS
./bin/talacert-api

Windows
.\bin\talacert-api.exe


Ou simplement :

make start

🧹 Nettoyage

Pour supprimer le dossier bin/ :

make clean

Linux / macOS
rm -rf bin

Windows
if exist bin rmdir /s /q bin


Le Makefile détecte automatiquement Windows avec :

ifeq ($(OS),Windows_NT)

🧹 Nettoyage du cache Go

Pour nettoyer les caches Go :

make clean-go-cache


Cette commande exécute :

go clean -cache
go clean -modcache


Elle supprime :

le cache de compilation Go ;
le cache local des modules Go.

⚠️ Après go clean -modcache, les dépendances devront être téléchargées à nouveau.

🧹 Nettoyage complet

Pour supprimer le binaire et nettoyer les caches Go :

make clean-all


Cette commande combine :

make clean
make clean-go-cache

🧪 Tests

Pour exécuter tous les tests :

make test


Cette commande exécute :

go test ./...


Ou directement :

go test ./...

🧹 Formatage

Pour formater automatiquement le code Go :

make fmt


Cette commande exécute :

go fmt ./...

🔎 Analyse statique

Pour analyser le code avec go vet :

make vet


Cette commande exécute :

go vet ./...


Avant de pousser des modifications, il est recommandé d'exécuter :

make fmt
make vet
make test

📚 Swagger

La documentation Swagger est générée avec :

make swagger


Le Makefile exécute :

swag init -g cmd/main.go --parseInternal --parseDependency


Les options utilisées sont :

--parseInternal : analyse les packages internes ;
--parseDependency : analyse également les dépendances.

Après génération et démarrage de l'application, Swagger UI est accessible à :

http://127.0.0.1:8080/swagger/doc

🛠️ Commandes Makefile
Commande	Description
make	Compile et démarre l'application
make all	Compile et démarre l'application
make install	Télécharge les dépendances
make tidy	Nettoie et synchronise les dépendances
make build	Compile l'application dans bin/
make start	Compile puis démarre le binaire
make run	Lance l'application avec go run
make test	Exécute les tests
make swagger	Génère la documentation Swagger
make clean	Supprime le dossier bin/
make clean-go-cache	Nettoie les caches Go
make clean-all	Supprime bin/ et les caches Go
make fmt	Formate le code Go
make vet	Analyse le code avec go vet
🌐 Accès à l'API
Health Check
GET /health


URL :

http://127.0.0.1:8080/health

Swagger UI
http://127.0.0.1:8080/swagger/doc

API
http://127.0.0.1:8080/api

📡 Endpoints principaux
🔐 Authentification
Méthode	Endpoint
POST	/api/v1/auth/login
POST	/api/v1/auth/refresh
POST	/api/v1/auth/logout
GET	/api/v1/auth/me
📄 Documents
Méthode	Endpoint
POST	/api/v1/documents
GET	/api/v1/documents
GET	/api/v1/documents/:document_id
PATCH	/api/v1/documents/:document_id
DELETE	/api/v1/documents/:document_id
GET	/api/v1/documents/by-hash/:hash
GET	/api/v1/documents/verify

Tous les endpoints utilisent la version :

/api/v1

🔍 Vérification d'un document

La vérification d'un document s'effectue à partir de son identifiant.

Endpoint actuel
GET /api/v1/documents/verify


Si l'API attend un JSON dans le corps de la requête :

{
  "document_id": "CERT-2025-0001"
}


⚠️ Recommandation : si document_id est envoyé dans le body, il est préférable d'utiliser POST plutôt que GET, car l'utilisation d'un body avec GET n'est pas uniformément supportée.

Réponse
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

🔐 Recherche par hash

Un document peut être retrouvé à partir de son hash.

Endpoint
GET /api/v1/documents/by-hash/:hash

Exemple
GET /api/v1/documents/by-hash/9f86d081884c7d659a2feaa0c55ad015


Cette fonctionnalité permet notamment de vérifier l'intégrité d'un document et pourra être utilisée dans le cadre d'une future intégration blockchain.

✏️ Modification d'un document

Les modifications partielles utilisent PATCH.

Endpoint
PATCH /api/v1/documents/:document_id

Exemple
PATCH /api/v1/documents/CERT-2025-0001

Body
{
  "owner": "Jean Tshibangu",
  "issuer": "Université de Kinshasa",
  "status": "valid"
}

🔐 Sécurité

TalaCert utilise ou prévoit les mécanismes suivants :

🔒 Hash SHA-256 pour l'intégrité des documents
🔑 Authentification JWT
♻️ Access tokens et refresh tokens
🛡️ Contrôle d'accès aux ressources
👤 Gestion de l'ownership des documents
✍️ Possibilité d'ajouter une signature numérique
⛓️ Intégration blockchain future

Les secrets JWT doivent être stockés dans les variables d'environnement.

⛓️ Intégration Blockchain

Chaque document pourra éventuellement être associé à des informations blockchain :

document
├── hash
├── blockchain
├── transaction_hash
└── network


La blockchain pourra servir à enregistrer une preuve d'existence ou d'intégrité du document sans avoir besoin de stocker le document lui-même sur la blockchain.

📌 Statuts des documents

Statuts actuellement supportés :

valid
invalid


Statuts pouvant être ajoutés ultérieurement :

pending
revoked
expired

🚀 Roadmap
 Authentification JWT
 CRUD des documents
 Vérification des documents
 Recherche par hash
 Makefile pour le développement
 Documentation Swagger
 Upload de fichiers PDF
 Génération de QR Code
 Signature numérique
 Intégration Blockchain
 Dashboard Admin
 Notifications email/SMS
 Gestion des documents révoqués
 Gestion de l'expiration des documents
🤝 Contribution

Les contributions sont les bienvenues !

1. Fork du projet
2. Créer une branche
git checkout -b feature/ma-feature

3. Développer la fonctionnalité
4. Formater et vérifier le code
make fmt
make vet

5. Exécuter les tests
make test

6. Commit
git commit -m "feat: ajout de ma fonctionnalité"

7. Push
git push origin feature/ma-feature

8. Créer une Pull Request
📄 Licence

MIT

👨‍💻 Auteur

Christian Shungu

Développeur Full Stack & DevOps

💡 Vision

TalaCert vise à devenir une plateforme de référence pour la vérification fiable et sécurisée des documents officiels, d'abord en République démocratique du Congo, puis à l'international.

L'objectif à long terme est de fournir une infrastructure permettant aux universités, administrations, entreprises et organismes émetteurs de publier des documents vérifiables, tout en permettant aux tiers de contrôler leur authenticité rapidement via une API, un identifiant unique, un hash ou un QR Code.