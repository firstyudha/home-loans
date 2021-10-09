package kelengkapan

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Kelengkapan, error)
	FindByPengajuanID(pengajuanID int) ([]Kelengkapan, error)
	FindByID(pengajuanID int) (Kelengkapan, error)
	FindPengajuanIDByUserID(userID int) (int, error)
	Save(kelengkapan Kelengkapan) (Kelengkapan, error)
	Update(kelengkapan Kelengkapan) (Kelengkapan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Kelengkapan, error) {
	var kelengkapans []Kelengkapan

	err := r.db.Find(&kelengkapans).Error
	if err != nil {
		return kelengkapans, err
	}

	return kelengkapans, nil
}

func (r *repository) FindByPengajuanID(pengajuanID int) ([]Kelengkapan, error) {
	var kelengkapans []Kelengkapan

	err := r.db.Where("pengajuan_id = ?", pengajuanID).Find(&kelengkapans).Error
	if err != nil {
		return kelengkapans, err
	}

	return kelengkapans, nil
}

func (r *repository) FindByID(PengajuanID int) (Kelengkapan, error) {
	var kelengkapan Kelengkapan

	err := r.db.Where("pengajuan_id = ?", PengajuanID).Find(&kelengkapan).Error
	if err != nil {
		return kelengkapan, err
	}

	return kelengkapan, nil
}

func (r *repository) FindPengajuanIDByUserID(userID int) (int, error) {
	var pengajuan Pengajuan
	err := r.db.Where("user_id = ?", userID).Find(&pengajuan).Error
	if err != nil {
		return pengajuan.ID, err
	}

	return pengajuan.ID, nil
}

func (r *repository) Save(kelengkapan Kelengkapan) (Kelengkapan, error) {
	err := r.db.Create(&kelengkapan).Error
	if err != nil {
		return kelengkapan, err
	}

	return kelengkapan, nil
}

func (r *repository) Update(kelengkapan Kelengkapan) (Kelengkapan, error) {
	err := r.db.Save(&kelengkapan).Error

	if err != nil {
		return kelengkapan, err
	}

	return kelengkapan, nil
}
