# apiGo de Nilsa

## Architecture
L’architecture de notre application se décline comme ceci:
```archi
config/
    database.go
controllers/
    users.go
models/
    user.go
main.go
router.go
```
```database
config/database.go permet la connexion à la base de données ainsi que la création de la table cars.
```
``` users
controllers/users.go fait la relation entre les requêtes http et notre struct Car
```
```user
models/user.go lie notre struct Car avec les actions en base de données.
```
```main
main.go lance la connexion à notre base de données, ajoute un enregistrement à notre base de donnée et lance le serveur http.
```
```router
router.go définit les routes de notre API
```
## Routage
### Get
```http request
GET	/api/v1/users	Lister tous les utilisateurs
GET	/api/v1/users/1	Lister l'utilisateur #1
```

### Post
```http request
POST	/api/v1/users	Poster un nouvel utilisateur
```

### Put
```http request
PUT	/api/v1/users/1	Modifier l'utilisateur #1
```

### Delete
```http request
DELETE	/api/v1/users/1	l'utilisateur #1
```
