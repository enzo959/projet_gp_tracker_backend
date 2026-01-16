# 1. Image de base avec Go 
FROM golang:1.25-alpine

# 2. Créer un dossier de travail dans le container
WORKDIR /app

# 3. Copier les fichiers de dépendances et télécharger les modules
COPY go.mod go.sum ./
RUN go mod download

# 4. Copier tout le reste du projet
COPY . .

# 5. Exposer le port que notre API utilisera
EXPOSE 8080

