package overWrite

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Assignment struct {
	SelfUrl  string `json:"_selfUrl"`
	Team     string `json:"team"`
	Policy   string `json:"policy"`
	Assigned bool   `json:"assigned"`
	User     string `json:"user"`
}

type Override struct {
	PublicID    string       `json:"publicId"`
	User        User         `json:"user"`
	Timezone    string       `json:"timezone"`
	Start       string       `json:"start"`
	End         string       `json:"end"`
	Assignments []Assignment `json:"assignments"`
}

type Response struct {
	Overrides []Override `json:"overrides"`
	SelfUrl   string     `json:"_selfUrl"`
}
