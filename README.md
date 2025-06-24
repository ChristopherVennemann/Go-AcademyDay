# Go Academy Day Projekt

Oh nein! Ist etwa schon Academy Day?!
Und mein Projekt ist noch gar nicht fertig :((
Zum Glück habe ich tolle Kolleg*innen wie euch.

Eure Aufgabe ist es, das Backend für meine neue Social Media Plattform fertigzustellen!

Ein Endpunkt ist schon durchimplementiert, für zwei weitere habe ich jedenfalls die Tests geschrieben. 
Den Rest packt ihr schon!

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
```go test internal``` oder via IDE.

### Anwendung (optional)
Falls ihr die gesamte Anwendung laufen lassen wollt, könnt ihr das tun.
Mit der [docker-compose](docker-compose.yaml) startet ihr eine passende PostgresDB.

Entry point is [main.go](main.go).

```go run .``` startet die Anwendung, alternativ wieder über die IDE.
Mit ```go build -o my/binary .``` kompiliert ihr die Binärdatei. 

## Euer Auftrag

Wie gesagt wird das hier ein Social Network. Ganz großes Ding.
Das [Datenmodell](internal/database/init.pg.sql) steht schon. 
User und Posts brauchen wir.

Der Endpunkt ```POST /user``` ist schon fertig,
und damit können die ersten Nutzer kommen. Aber ich habe schon ein paar weiter features geplant B)

1. Wir brauchen einen Endpunkt ```GET /users```, damit wir nicht den Überblick über all die Nutzer verlieren!
    Dazu gibt es schon Tests, [handler](internal/handler/user/handler_test.go), [service](internal/service/user_test.go), 
    [repo](internal/integration/user_repository_integration_test.go) und integration.
    Ich habe auch schon die Methoden und Interfaces angelegt, ihr brauch nur noch auszufüllen.
2. Als nächstes brauchen wir den Endpunkt ```DELETE /user/{id}```. Auch dafür sind Tests vorhanden. 
   Die Methoden dürft ihr diesmal sogar selbst schreiben.
3. Falls ihr damit fertig werdet, steht euch die Welt offen! Schaut und experimentiert selbst!
    Offene Punkte auf meiner Liste sind z.B.:
    - Weiter User Endpunkte hinzufügen
    - Post Endpunkte hinzufügen
    - DTOs (Data Transfer Structs?) einführen
    - Testabdeckung
    - besseres Error handling
    - ...

### Viel Erfolg und frohes Schaffen!
