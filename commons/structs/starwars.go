package structs

type SWPeople struct {
	Name      string        `json:"name,omitempty"`
	Height    string        `json:"height,omitempty"`
	Mass      string        `json:"mass,omitempty"`
	HairColor string        `json:"hair_color,omitempty"`
	SkinColor string        `json:"skin_color,omitempty"`
	EyeColor  string        `json:"eye_color,omitempty"`
	BirthYear string        `json:"birth_year,omitempty"`
	Gender    string        `json:"gender,omitempty"`
	Homeworld string        `json:"homeworld,omitempty"`
	Films     []string      `json:"films,omitempty"`
	Species   []string      `json:"species,omitempty"`
	Vehicles  []interface{} `json:"vehicles,omitempty"`
	Starships []interface{} `json:"starships,omitempty"`
	Created   string        `json:"created,omitempty"`
	Edited    string        `json:"edited,omitempty"`
	URL       string        `json:"url,omitempty"`
}
