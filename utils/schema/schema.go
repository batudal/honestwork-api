package schema

type User struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Bio     string `json:"bio"`
}

type Scope struct {
	Name     string `json:"name"`
	Basic    string `json:"starter"`
	Standard string `json:"standard"`
	Premium  string `json:"premium"`
}

type Post struct {
	Address string   `json:"address"`
	UUID    string   `json:"id"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Formats []string `json:"formats"`
	Images  []string `json:"images"`
	Scopes  []Scope  `json:"scopes"`
}

type Project struct {
	Employee User   `json:"employee"`
	Employer User   `json:"employer"`
	Token    string `json:"token"`
	Amount   string `json:"amount"`
}
