package kelengkapan

import "errors"

type Service interface {
	GetKelengkapans(pengajuanID int) ([]Kelengkapan, error)
	CreateKelengkapan(userID int, input CreateKelengkapanInput) (Kelengkapan, error)
	SaveDokumenPendukung(UserID int, fileLocation string) (Kelengkapan, error)
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

func (s *service) CreateKelengkapan(userID int, input CreateKelengkapanInput) (Kelengkapan, error) {
	kelengkapan := Kelengkapan{}
	pengajuan_id, err := s.repository.FindPengajuanByUserID(userID)

	if err != nil {
		return kelengkapan, err
	}

	kelengkapan.PengajuanID = pengajuan_id
	kelengkapan.AlamatRumah = input.AlamatRumah
	kelengkapan.LuasRumah = input.LuasRumah
	kelengkapan.HargaRumah = input.HargaRumah
	kelengkapan.JangkaPembayaran = input.JangkaPembayaran
	kelengkapan.Status = input.Status

	isKelengkapanExist, err := s.repository.FindByPengajuanID(pengajuan_id)

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

func (s *service) SaveDokumenPendukung(userID int, fileLocation string) (Kelengkapan, error) {
	//find pengajuan id by user id
	pengajuan_id, err := s.repository.FindPengajuanByUserID(userID)
	if err != nil {
		return Kelengkapan{}, err
	}

	//find kelengkapan by pengajua id
	kelengkapan, err := s.repository.FindByID(pengajuan_id)
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
