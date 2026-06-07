# 📄 TalaCert API

TalaCert est une API de vérification de documents officiels (certificats, diplômes, attestations) basée sur Symfony et API Platform.

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

* PHP 8+
* Symfony
* API Platform
* Doctrine ORM
* MySQL / PostgreSQL

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
composer install
```

---

### 3. Configuration de la base de données

Modifier le fichier `.env` :

```env
DATABASE_URL="mysql://user:password@127.0.0.1:3306/talacert"
```

---

### 4. Créer la base de données

```bash
php bin/console doctrine:database:create
```

---

### 5. Lancer les migrations

```bash
php bin/console doctrine:migrations:migrate
```

---

### 6. Démarrer le serveur

```bash
symfony server:start
```

Ou sous Windows :

```bash
php -S 127.0.0.1:8000 -t public
```

---

## 🌐 Accès à l’API

* Swagger UI :

```
http://127.0.0.1:8000/api
```

---

## 📡 Endpoints principaux

### 📄 Documents

| Méthode | Endpoint            |
| ------- | ------------------- |
| GET     | /api/documents      |
| GET     | /api/documents/{id} |
| POST    | /api/documents      |
| PUT     | /api/documents/{id} |
| DELETE  | /api/documents/{id} |

---

### 🔍 Vérification d’un document

```
POST /api/verify
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
  "status": "valid",
  "data": {
    "owner": "Jean Tshibangu",
    "type": "Certificat",
    "issuer": "Université de Kinshasa"
  }
}
```

---

## 🗄 Structure du projet

```
src/
 ├── Controller/
 │    └── VerificationController.php
 │
 ├── Entity/
 │    └── Document.php
 │
 ├── Repository/
 │    └── DocumentRepository.php
 │
 └── Service/
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
* `revoked`
* `expired`

---

## 🧪 Tests

```bash
php bin/phpunit
```

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

TalaCert vise à devenir une plateforme de référence pour la vérification des documents officiels en Afrique et dans le monde.

---
