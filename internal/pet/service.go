package pet

type Service struct {
	pet Pet
}

func NewService() *Service {
	return &Service{
		pet: Pet{
			Name:       "Вульпи",
			Mood:       "happy",
			Level:      1,
			Experience: 0,
		},
	}
}

func (s *Service) GetPet() Pet {
	return s.pet
}

func (s *Service) RespondTo(message string) string {
	return s.pet.RespondTo(message)
}

func (s *Service) ChangeMood(mood string) {
	s.pet.Mood = mood
}
