package schema

type User struct {
	Salt       string   `json:"salt"`
	Signature  string   `json:"signature"`
	Username   string   `json:"username"`
	ShowEns    bool     `json:"show_ens"`
	Title      string   `json:"title"`
	ImageUrl   string   `json:"image_url"`
	ShowNFT    bool     `json:"show_nft"`
	NFTAddress string   `json:"nft_address"`
	NFTId      string   `json:"nft_id"`
	Email      string   `json:"email"`
	Timezone   string   `json:"timezone"`
	Bio        string   `json:"bio"`
	Links      []string `json:"links"`
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
	Rating 		 int64    `json:"rating"`
}
