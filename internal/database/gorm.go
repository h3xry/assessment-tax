package database

import (
	"context"
	"fmt"
	"time"

	"github.com/h3xry/assessment-tax/internal/config"
	"github.com/h3xry/assessment-tax/pkg/models"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(lc fx.Lifecycle, cfg *config.ENV) (*gorm.DB, error) {
	fmt.Println("Connecting to database...")
	instance, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		fmt.Println("database: connection fail! : ", err)
		return nil, err
	}

	db, err := instance.DB()
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("database: ping fail!", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := instance.AutoMigrate(&models.Deductions{}); err != nil {
		fmt.Println("database: auto migrate fail! : ", err)
		return nil, err
	}

	var cnt int64
	if instance.Model(&models.Deductions{}).Where("name = ?", "kReceipt").Count(&cnt); cnt == 0 {
		if err := instance.Create(&models.Deductions{
			Name:   "kReceipt",
			Amount: 50000.00,
		}).Error; err != nil {
			fmt.Println("database: create kReceipt fail! : ", err)
			return nil, err
		}
	}

	if instance.Model(&models.Deductions{}).Where("name = ?", "personalDeduction").Count(&cnt); cnt == 0 {
		if err := instance.Create(&models.Deductions{
			Name:   "personalDeduction",
			Amount: 60000.00,
		}).Error; err != nil {
			fmt.Println("database: create personalDeduction fail! : ", err)
			return nil, err
		}
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			fmt.Println("database: closing connection")
			return db.Close()
		},
	})
	return instance, nil
}
