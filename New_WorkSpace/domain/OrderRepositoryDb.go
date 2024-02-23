package domain

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
	"github.com/AdiLambe/TestGoLangCodes/workspace/logger"
	_ "github.com/lib/pq"
)

type OrderRepositoryDb struct {
	client *gorm.DB
}

func (d OrderRepositoryDb) FindAll(ctx context.Context, status string) ([]Order, *errs.AppError) {

	select {
	case <-ctx.Done():
		return nil, errs.NewUnexpectedError("Deadline exceeded")
	default:
	}
	var err error
	orders := make([]Order, 0)

	query := d.client.WithContext(ctx).Table("listorders")

	if status != "" {
		query = query.Where("status = ?", status)
	}
	err = query.Find(&orders).Error
	if err != nil {
		logger.Error("Error While querying listorders table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return orders, nil
}

var ErrNoRows = errors.New("sql: no rows in result set")

// receiver func which takes id  and will return an order and error
func (d OrderRepositoryDb) ById(id string) (*Order, *errs.AppError) {

	var order Order
	err := d.client.Table("listorders").First(&order, "order_id = $1", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Order not found")
		} else {
			logger.Error("Error while scanning order" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database error")
		}
	}
	// returning address of the order and nil
	return &order, nil
}

func (d OrderRepositoryDb) SaveOrder(order Order) (*Order, *errs.AppError) {

	log.Printf("Order before insertion, %+v", order)

	err := d.client.Table("listorders").Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&order).Error

	if err != nil {
		logger.Error("Error while inserting order: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	log.Printf("Order after insertion: %+v", order)

	return &order, nil
}

func NewOrderRepositoryDb() OrderRepositoryDb {
	dsn := "user=postgres password=Agl@29121998 dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}

	var orders []Order
	err = db.Table("listorders").Find(&orders).Error
	if err != nil {
		panic(err)
	}

	return OrderRepositoryDb{client: db}
}
