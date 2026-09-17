package pet

type Pet struct {
	Name       string `json:"name"`
	Mood       string `json:"mood"`
	Level      int    `json:"level"`
	Experience int    `json:"experience"`
}

func (p Pet) RespondTo(message string) string {
	return p.Name + ": Привет! Ты сказал: " + message
}
