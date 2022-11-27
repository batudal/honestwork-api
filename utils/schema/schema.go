package schema

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Posts int    `json:"posts"`
}

// type Scope struct {
// 	Name     string `json:"name"`
// 	Basic    string `json:"basic"`
// 	Standard string `json:"standard"`
// 	Premium  string `json:"premium"`
// }

type Post struct {
	Address string   `json:"address"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Formats []string `json:"formats"`
	Images  []string `json:"images"`
	// Scopes  []Scope  `json:"scopes"`
}

type Project struct {
	Employee User   `json:"employee"`
	Employer User   `json:"employer"`
	Token    string `json:"token"`
	Amount   string `json:"amount"`
}
