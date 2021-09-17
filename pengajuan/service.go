package pengajuan

type Service interface {
	GetPengajuans(userID int) ([]Pengajuan, error)
	CreatePengajuan(input CreatePengajuanInput) (Pengajuan, error)
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
