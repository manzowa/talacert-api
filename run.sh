#!/usr/bin/env bash

# Début application

# Titre du terminal (si supporté)
echo -ne "\033]0;APPLICATION TALACERT API\007"

# Effacer l'écran
clear

# Couleur verte
GREEN="\033[0;32m"
RESET="\033[0m"

echo -e "${GREEN}"
echo "  ╔═══════════════════════════════════════════════════╗"
echo "  ║                     API TALACERT                  ║"
echo "  ╚═══════════════════════════════════════════════════╝"
echo
echo -e "${RESET}"

# Nettoyer le cache Go (optionnel)
# go clean -cache
# go clean -modcache

# Exécuter l'application
go run cmd/main.go