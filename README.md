# Go Academy Day Projekt

Oha! Ist etwa schon Academy Day?!
Und mein Projekt ist noch gar nicht fertig!
Zum Glück habe ich tolle Kolleg*innen wie euch.

Eure Aufgabe ist es, das Backend für meine neue Social Media Plattform fertigzustellen!

Ein Endpunkt ist schon durchimplementiert. Für den nächsten habe ich schon Tests geschrieben.
Den dürft ihr jetzt implementieren.

## Setup
### Go
Installiert auf eurem Gerät Go v1.24 oder neuer.

### IDE
Nutzt Intellij IDEA mit dem Go plugin.

### Dependencies
Werden in der [go.mod](go.mod) verwaltet.
Mit ```go mod tidy``` installiert ihr sie.

### Container
Die Integrationstests verwenden Testcontainer. 
Ihr brauch eine Docker-API-kompatible Container-Laufzeitumgebung.

### Tests
Damit sollten die Tests auch schon laufen.
Lasst einfach die Tests des package [internal](internal) laufen,
```go test ./internal/...``` oder via IDE.

### Anwendung
Mit der [docker-compose](docker-compose.yaml) startet ihr eine passende PostgresDB.
```sh
docker compose up -d
```

Entry point is [main.go](main.go).

```go run .``` startet die Anwendung, alternativ wieder über die IDE.
Mit ```go build -o my/binary .``` kompiliert ihr die Binärdatei. 

## Euer Auftrag

Wie gesagt wird das hier ein Social Network. Ganz großes Ding.
Das [Datenmodell](internal/database/init.pg.sql) steht schon. 
User und Posts brauchen wir.

Der Endpunkt ```POST /user``` ist schon fertig,
und damit können die ersten Nutzer kommen.

1. Als nächstes brauchen wir einen Endpunkt ```GET /users```, damit wir nicht den Überblick über all die Nutzer verlieren.
    Dazu gibt es schon Tests, [handler](internal/handler/user/handler_test.go), [service](internal/service/user_test.go), 
    [repo](internal/integration/user_repository_integration_test.go) und [integration](internal/integration/user_full_integration_test.go).
    Ich habe auch schon die Methoden und Interfaces angelegt, ihr brauch nur noch auszufüllen.
    Für den Fall, dass die Datenbankabfrage einen Fehler zurückgibt, könnt ihr vorerst eine leere Liste zurückgeben.

2. Jetzt behandeln wir den Fehler. Wenn die Datenbankabfrage einen Fehler zurückgibt, 
reicht diesen Fehler bis zum Handler weiter.
Die Anfrage sollte dann einen http Status 500 zurückgeben.
Schreibt Tests.

Falls ihr damit fertig werdet, steht euch die Welt offen! Schaut und experimentiert selbst!
    Offene Punkte auf meiner Liste sind z.B.:
    - Delete User Endpunkte hinzufügen
    - DTOs (Data Transfer Structs?) einführen
    - Testabdeckung
    - besseres Error handling
    - ...

### Viel Erfolg und frohes Schaffen!
