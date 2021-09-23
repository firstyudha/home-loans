package kelengkapan

import "errors"

type Service interface {
	GetKelengkapans(userID int) ([]Kelengkapan, error)
	CreateKelengkapan(input CreateKelengkapanInput) (Kelengkapan, error)
	SaveDokumenPendukung(inputID GetKelengkapanInput, fileLocation string) (Kelengkapan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetKelengkapans(PengajuanID int) ([]Kelengkapan, error) {
	if PengajuanID != 0 {
		kelengkapans, err := s.repository.FindByPengajuanID(PengajuanID)
		if err != nil {
			return kelengkapans, err
		}

		return kelengkapans, nil
	}

	kelengkapans, err := s.repository.FindAll()
	if err != nil {
		return kelengkapans, err
	}

	return kelengkapans, nil
}

func (s *service) CreateKelengkapan(input CreateKelengkapanInput) (Kelengkapan, error) {
	kelengkapan := Kelengkapan{}
	kelengkapan.PengajuanID = input.PengajuanID
	kelengkapan.AlamatRumah = input.AlamatRumah
	kelengkapan.LuasRumah = input.LuasRumah
	kelengkapan.HargaRumah = input.HargaRumah
	kelengkapan.JangkaPembayaran = input.JangkaPembayaran
	kelengkapan.Status = input.Status

	isKelengkapanExist, err := s.repository.FindByPengajuanID(input.PengajuanID)

	if err != nil {
		return kelengkapan, err
	}

	if len(isKelengkapanExist) > 0 {
		return Kelengkapan{}, errors.New("anda sudah mengajukan Kelengkapan Dokumen")
	}

	newKelengkapan, err := s.repository.Save(kelengkapan)
	if err != nil {
		return newKelengkapan, err
	}

	return newKelengkapan, nil
}

func (s *service) SaveDokumenPendukung(inputID GetKelengkapanInput, fileLocation string) (Kelengkapan, error) {
	kelengkapan, err := s.repository.FindByID(inputID.PengajuanID)
	if err != nil {
		return Kelengkapan{}, err
	}

	kelengkapan.DokumenPendukung = fileLocation

	updatedKelengkapan, err := s.repository.Update(kelengkapan)
	if err != nil {
		return updatedKelengkapan, err
	}

	return updatedKelengkapan, nil
}
