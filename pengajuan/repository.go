package pengajuan

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Pengajuan, error)
	FindByUserID(userID int) ([]Pengajuan, error)
	FindByID(userID int) (Pengajuan, error)
	Save(pengajuan Pengajuan) (Pengajuan, error)
	Update(pengajuan Pengajuan) (Pengajuan, error)
	Delete(userID int) error
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

func (r *repository) Delete(userID int) error {

	var pengajuan Pengajuan

	err := r.db.Where("user_id = ?", userID).Delete(&pengajuan).Error
	if err != nil {
		return err
	}

	return nil

}
