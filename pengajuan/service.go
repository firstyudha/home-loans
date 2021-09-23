package pengajuan

import "errors"

type Service interface {
	GetPengajuans(userID int) ([]Pengajuan, error)
	CreatePengajuan(input CreatePengajuanInput) (Pengajuan, error)
	SaveBuktiKTP(inputID GetPengajuanInput, fileLocation string) (Pengajuan, error)
	SaveBuktiSlipGaji(inputID GetPengajuanInput, fileLocation string) (Pengajuan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPengajuans(userID int) ([]Pengajuan, error) {
	if userID != 0 {
		pengajuans, err := s.repository.FindByUserID(userID)
		if err != nil {
			return pengajuans, err
		}

		return pengajuans, nil
	}

	pengajuans, err := s.repository.FindAll()
	if err != nil {
		return pengajuans, err
	}

	return pengajuans, nil
}

func (s *service) CreatePengajuan(input CreatePengajuanInput) (Pengajuan, error) {
	pengajuan := Pengajuan{}
	pengajuan.UserID = input.UserID
	pengajuan.Nik = input.Nik
	pengajuan.NamaLengkap = input.NamaLengkap
	pengajuan.TempatLahir = input.TempatLahir
	pengajuan.TanggalLahir = input.TanggalLahir
	pengajuan.Alamat = input.Alamat
	pengajuan.Pekerjaan = input.Pekerjaan
	pengajuan.PendapatanPerbulan = input.PendapatanPerbulan
	pengajuan.Status = input.Status

	isPengajuanExist, err := s.repository.FindByID(input.UserID)

	if err != nil {
		return pengajuan, err
	}

	if isPengajuanExist.ID > 0 {
		return Pengajuan{}, errors.New("anda sudah mengajukan KPR")
	}

	newPengajuan, err := s.repository.Save(pengajuan)
	if err != nil {
		return newPengajuan, err
	}

	return newPengajuan, nil
}

func (s *service) SaveBuktiKTP(inputID GetPengajuanInput, fileLocation string) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return Pengajuan{}, err
	}

	pengajuan.BuktiKtp = fileLocation

	updatedPengajuan, err := s.repository.Update(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}

func (s *service) SaveBuktiSlipGaji(inputID GetPengajuanInput, fileLocation string) (Pengajuan, error) {
	pengajuan, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return Pengajuan{}, err
	}

	pengajuan.BuktiSlipGaji = fileLocation

	updatedPengajuan, err := s.repository.Update(pengajuan)
	if err != nil {
		return updatedPengajuan, err
	}

	return updatedPengajuan, nil
}