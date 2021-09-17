package pengajuan

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Pengajuan, error)
	FindByUserID(userID int) ([]Pengajuan, error)
	Save(pengajuan Pengajuan) (Pengajuan, error)
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

func (r *repository) Save(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Create(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}
