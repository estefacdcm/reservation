package reservation

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/estefacdcm/reservation/internal/domain/entity"

	"github.com/estefacdcm/reservation/internal/infrastructure/pkg/database"
)

const (
	tableName                              = "reservation"
	errorCreatingReservation               = "error creating reservation"
	errorInsertReservationRowsAffectedZero = "error inserting reservation rows affected"
	errorFindReservartionByCustomerID      = "an error occurred while trying to query reservation by customerID"
	errorFindReservartionByNumber          = "an error occurred while trying to query reservation by number"
)

type RepositoryReservationAdapter struct {
	db *gorm.DB
}

func NewRepositoryReservationAdapter(db *gorm.DB) *RepositoryReservationAdapter {
	return &RepositoryReservationAdapter{
		db: db,
	}
}

func (rra *RepositoryReservationAdapter) SaveReservation(reservationEntity *entity.ReservationEntity) (*entity.ReservationEntity, error) {
	dbName, err := database.GetDatabaseNameConnection()
	if err != nil {
		return nil, err
	}

	trxDBresult := rra.db.Clauses(dbresolver.Use(dbName)).Table(tableName).Create(reservationEntity)
	if trxDBresult.Error != nil {
		logrus.Errorln(errorCreatingReservation, trxDBresult.Error.Error())
		return nil, errors.New(errorCreatingReservation)
	}
	if trxDBresult.RowsAffected == 0 {
		logrus.Errorln(errorCreatingReservation, errorInsertReservationRowsAffectedZero)
		return nil, errors.New(errorCreatingReservation)
	}
	return reservationEntity, nil
}

func (rra *RepositoryReservationAdapter) FindReservationByCustomerID(customerID string) ([]dto.ReservationDTO, error) {
	var reservation []dto.ReservationDTO
	dbName, err := database.GetDatabaseNameConnection()
	if err != nil {
		return nil, err
	}

	trxDBResult := rra.db.Clauses(dbresolver.Use(dbName)).Table(tableName).
		Select("customer_id", "reservation_number").
		Where("reservation.customer_id = ?", customerID).
		Find(&reservation)

	if trxDBResult.Error != nil {
		logrus.Errorln(trxDBResult.Error.Error())
		return make([]dto.ReservationDTO, 0), errors.New(errorFindReservartionByCustomerID)
	}

	return reservation, nil
}

func (rra *RepositoryReservationAdapter) FindReservartionByNumber(reservationNumber int64) ([]dto.ReservationDTO, error) {
	var reservation []dto.ReservationDTO
	dbName, err := database.GetDatabaseNameConnection()
	if err != nil {
		return nil, err
	}

	trxDBResult := rra.db.Clauses(dbresolver.Use(dbName)).Table(tableName).
		Select("customer_id", "reservation_number").
		Where("reservation.reservation_number = ?", reservationNumber).
		Find(&reservation)

	if trxDBResult.Error != nil {
		logrus.Errorln(trxDBResult.Error.Error())
		return make([]dto.ReservationDTO, 0), errors.New(errorFindReservartionByCustomerID)
	}

	return reservation, nil
}
