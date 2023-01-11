package schema

type User struct {
	Salt       string   `json:"salt"`
	Signature  string   `json:"signature"`
	Username   string   `json:"username"`
	ShowEns    bool     `json:"show_ens"`
	EnsName    string   `json:"ens_name"`
	Title      string   `json:"title"`
	ImageUrl   string   `json:"image_url"`
	ShowNFT    bool     `json:"show_nft"`
	NFTUrl     string   `json:"nft_url"`
	NFTAddress string   `json:"nft_address"`
	NFTId      string   `json:"nft_id"`
	Email      string   `json:"email"`
	Timezone   string   `json:"timezone"`
	Bio        string   `json:"bio"`
	Links      []string `json:"links"`
	Rating     int64    `json:"rating"`
}

type Skill struct {
	UserAddress  string   `json:"user_address"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Links        []string `json:"links"`
	ImageUrls    []string `json:"image_urls"`
	MinimumPrice int      `json:"minimum_price"`
	Publish      bool     `json:"publish"`
	CreatedAt    int64    `json:"created_at"`
}

type Job struct {
	UserAddress    string    `json:"user_address"`
	Username       string    `json:"username"`
	TokenPaid      string    `json:"token_paid"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Tags           []string  `json:"tags"`
	Links          []string  `json:"links"`
	Budget         int       `json:"budget"`
	Installments   int64     `json:"installments"`
	TimeZone       string    `json:"timezone"`
	TokensAccepted []Network `json:"tokens_accepted"`
	StickyDuration int64     `json:"sticky_duration"`
	CreatedAt      int64     `json:"created_at"`
	TxHash         string    `json:"tx_hash"`
	ImageUrl       string    `json:"image_url"`
}

type Network struct {
	Id     int64   `json:"id"`
	Tokens []Token `json:"tokens"`
}
type Token struct {
	Address string `json:"address"`
}

type Whitelist struct {
	Tokens []Token `json:"tokens"`
}
