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
	UserAddress      string        `json:"user_address"`
	PaymentTxHash    string        `json:"payment_tx_hash"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	Tags             []string      `json:"tags"`
	Links            []string      `json:"links"`
	Budget           int64         `json:"budget"`
	Installments     int64         `json:"installments"`
	Networks         []string      `json:"networks"`
	TimeZone         string        `json:"timezone"`
	TokensAccepted   []string      `json:"tokens"`
	HighlightOptions HighlightOpts `json:"highlight_options"`
}

type HighlightOpts struct {
	StickyDuration  int64 `json:"sticky_duration"`
	HighlightOption int64 `json:"highlight"`
}
