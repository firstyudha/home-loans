package pengajuan

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Pengajuan, error)
	FindByUserID(userID int) ([]Pengajuan, error)
	FindByID(userID int) (Pengajuan, error)
	Save(pengajuan Pengajuan) (Pengajuan, error)
	Update(pengajuan Pengajuan) (Pengajuan, error)
	Delete(pengajuanID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Pengajuan, error) {
	var pengajuans []Pengajuan

	err := r.db.Preload("User").Preload("Kelengkapan").Find(&pengajuans).Error
	if err != nil {
		return pengajuans, err
	}

	return pengajuans, nil
}

func (r *repository) FindByUserID(userID int) ([]Pengajuan, error) {
	var pengajuans []Pengajuan

	err := r.db.Where("user_id = ?", userID).Preload("User").Preload("Kelengkapan").Find(&pengajuans).Error
	if err != nil {
		return pengajuans, err
	}

	return pengajuans, nil
}

func (r *repository) FindByID(userID int) (Pengajuan, error) {
	var pengajuan Pengajuan

	err := r.db.Where("user_id = ?", userID).Preload("User").Preload("Kelengkapan").Find(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Save(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Create(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Update(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Save(&pengajuan).Error

	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Delete(pengajuanID int) error {

	r.db.Transaction(func(tx *gorm.DB) error {

		kelengkapan := tx.Exec("UPDATE kelengkapans SET deleted_at = ? WHERE pengajuan_id = ?", time.Now(), pengajuanID)

		//check error update kelengkapan, then rollback
		if kelengkapan.Error != nil {
			tx.Rollback()
			panic(kelengkapan.Error)
		}

		pengajuan := tx.Exec("UPDATE pengajuans SET deleted_at = ? WHERE id = ?", time.Now(), pengajuanID)

		//check error update pengajuan, then rollback
		if pengajuan.Error != nil {
			tx.Rollback()
			panic(pengajuan.Error)
		}

		//if no error, then commit
		tx.Commit()
		return nil
	})

	return nil

}
