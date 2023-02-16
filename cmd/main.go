package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	reservationUseCase "github.com/estefacdcm/reservation/internal/application/usecase"
	reservationService "github.com/estefacdcm/reservation/internal/domain/service"
	reservationAdapter "github.com/estefacdcm/reservation/internal/infrastructure/adapters/database/reservation"
	healthHandler "github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/api/health"
	reservationHandler "github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/api/reservation"
	db "github.com/estefacdcm/reservation/internal/infrastructure/pkg/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("error loading .env, %v", err)
	}
}

func main() {
	databaseClient := db.NewPostgresConnection()

	reservationRepository := reservationAdapter.NewRepositoryReservationAdapter(databaseClient)
	reservationService := reservationService.NewReservationService(reservationRepository)
	reservationUseCase := reservationUseCase.NewReservationUseCase(reservationService)
	reservationHandler := reservationHandler.NewReservationHandler(reservationUseCase)
	healthHandler := healthHandler.NewHealthHandler()

	runServer(reservationHandler, healthHandler)
}

func runServer(reservationHandler *reservationHandler.ReservationHandler, healthHandler *healthHandler.HealthHandler) {
	e := echo.New()

	apiGroup := e.Group("reservation")
	apiGroup.GET("/health", healthHandler.Health)
	apiGroup.POST("/", reservationHandler.SaveReservationHandler)
	apiGroup.GET("/customer/:customerId", reservationHandler.GetReservationByCustomerIDHandler)
	apiGroup.GET("/reservationNumber/:reservationNumber", reservationHandler.GetReservationByNumberHandler)

	err := e.Start(":" + os.Getenv("PORT"))
	if err != nil {
		logrus.Fatal(err)
	}
}
