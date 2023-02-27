# Reservation
Technical test BACU

# Objective
    Develop API to manage resrvations

# Technical specification
## language
    go1.19.5 darwin/arm64

## Database
    postgresql

# Functional specification
## Run local
### Steps
    1. Validate / install PostgreSQL. example to install: brew install postgresql
    2. Set params connection in file local.env 
    2. Install local  golang-migrate, example: brew install golang-migrate
    3. Execute migration: migrate -path database/migration/ -database "postgresql://username:<<username>>@localhost:<<port>>/<<database_name>>?sslmode=disable" -verbose up
### Migrations
    1. Create DB: execute migration 000001_create-schema.down.sql
    2. Delete DB: execute command go run cmd/main.go

## Endpoints
### Link documentation Swagger
    [Endpoints] (https://app.swaggerhub.com/apis/ESTEFANIAARANGOCENTE/reservation-api/1.0.0#/)
    
## Diagrams
### Architecture
    ![Architecture!](/docs/diagrams/reservation_architecture_use_case.png "Architecture diagram")

### Create reservation
    ![Create!](/docs/diagrams/Create_reservation.png "Create Reservation")

### Get reservation by customerID
    ![Get!](/docs/diagrams/get_reservation_by_customerID.png "Get Reservation")

### Get reservation by reservationNumber
    ![Create!](/docs/diagrams/get_reservation_by_reservationNumber.png "Get Reservation")


