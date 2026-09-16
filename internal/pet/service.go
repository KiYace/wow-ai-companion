package pet

type Service struct {
	pet Pet
}

func NewService() *Service {
	return &Service{
		pet: Pet{
			Name:  "Вульпи",
			Mood:  "happy",
			Level: 1,
		},
	}
}

func (s *Service) GetPet() Pet {
	return s.pet
}

func (s *Service) RespondTo(message string) string {
	return s.pet.RespondTo(message)
}
