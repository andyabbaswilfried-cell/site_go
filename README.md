# Serveur web en Go

Ce projet est un petit serveur web écrit en Go. Il sert une page d'accueil, affiche une liste d'artistes et propose une page de détail pour chaque artiste. Les informations des artistes sont actuellement définies directement dans le code : aucune base de données n'est utilisée.

## Arborescence

```text
site_go/
├── go.mod
├── main.go
├── README.md
├── static/
│   └── style.css
└── templates/
    ├── artist.html
    └── artists.html
```

- `main.go` définit les artistes, les routes HTTP et le démarrage du serveur.
- `go.mod` déclare le module Go et sa version de langage (`go 1.27.1`).
- `templates/artists.html` affiche la liste des artistes.
- `templates/artist.html` affiche les informations d'un artiste.
- `static/style.css` contient les styles utilisés par les pages.

## Prérequis

Installez Go, dans la version indiquée par `go.mod` ou une version compatible, puis ouvrez un terminal à la racine du projet.

## Démarrer le serveur

```powershell
go run .
```

Le serveur écoute sur le port `8080`. Une fois le serveur démarré, ouvrez [http://localhost:8080](http://localhost:8080) dans votre navigateur. Pour arrêter le serveur, utilisez `Ctrl+C` dans le terminal.

## Pages disponibles

| Adresse | Contenu |
| --- | --- |
| `/` | Message d'accueil « Bonjour le web ! » |
| `/artists/` | Liste des artistes |
| `/artists/1` | Page de Queen |
| `/artists/2` | Page de Daft Punk |
| `/static/style.css` | Feuille de style des pages |

Un identifiant qui n'est pas un nombre renvoie une erreur `400`. Un identifiant numérique qui ne correspond à aucun artiste renvoie une erreur `404`.

## Modifier les artistes

Les artistes se trouvent dans la variable `artists` de `main.go`. Chaque entrée contient un nom (`Name`), une année (`Year`) et un identifiant numérique (`Id`). Ajoutez ou modifiez une entrée dans cette liste pour mettre à jour les artistes affichés, puis redémarrez le serveur.
