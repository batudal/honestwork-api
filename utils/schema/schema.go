package schema

// todo: abstract go-playground/validator/v10 props to have clean schema
type User struct {
	Signature  string       `json:"signature" validate:"required"`
	Username   string       `json:"username" validate:"required,min=5,max=50"`
	ShowEns    *bool        `json:"show_ens" validate:"boolean"`
	EnsName    string       `json:"ens_name"`
	Title      string       `json:"title" validate:"required,min=5,max=50"`
	ImageUrl   string       `json:"image_url" validate:"omitempty,url"`
	ShowNFT    *bool        `json:"show_nft" validate:"boolean"`
	NFTUrl     string       `json:"nft_url" validate:"omitempty,url"`
	NFTAddress string       `json:"nft_address" validate:"eth_addr"`
	NFTId      string       `json:"nft_id"` // todo: custom owner check
	Email      string       `json:"email" validate:"email"`
	Timezone   string       `json:"timezone"` // todo: figure out timezone check
	Bio        string       `json:"bio" validate:"required,min=200,max=2000"`
	Links      []string     `json:"links" validate:"required,min=1,max=3,dive,url"`
	Rating     int64        `json:"rating"`
	Watchlist  []*Watchlist `json:"watchlist"`
	Favorites  []*Favorite  `json:"favorites"`
}

type FavoriteInput struct {
	Address string `json:"address"`
	Slot    int    `json:"slot"`
}

type Favorite struct {
	Input    *FavoriteInput `json:"input"`
	Username string         `json:"username"`
	Title    string         `json:"title"`
	ImageUrl string         `json:"image_url"`
}

type Watchlist struct {
	Input    *WatchlistInput `json:"input"`
	Username string          `json:"username"`
	Title    string          `json:"title"`
	ImageUrl string          `json:"image_url"`
}

type WatchlistInput struct {
	Address string `json:"address"`
	Slot    int    `json:"slot"`
}

type Skill struct {
	UserAddress  string   `json:"user_address" validate:"required,eth_addr"`
	Title        string   `json:"title" validate:"required,min=5,max=50"`
	Description  string   `json:"description" validate:"required,min=200,max=2000"`
	Tags         []string `json:"tags" validate:"required,min=1,max=3,dive,min=3,max=20"`
	Links        []string `json:"links" validate:"required,min=1,max=3,dive,url"`
	ImageUrls    []string `json:"image_urls" validate:"required,min=1,max=8,dive,omitempty,url"`
	MinimumPrice int      `json:"minimum_price" validate:"required,min=1000,max=1000000"`
	Publish      bool     `json:"publish" validate:"required,boolean"`
	CreatedAt    int64    `json:"created_at"`
}

type Job struct {
	Email          string        `json:"email" validate:"required,email"`
	Slot           int           `json:"slot"`
	UserAddress    string        `json:"user_address" validate:"required,eth_addr"`
	Username       string        `json:"username" validate:"required,min=5,max=50"`
	TokenPaid      string        `json:"token_paid" validate:"required,eth_addr"`
	Title          string        `json:"title" validate:"required,min=5,max=50"`
	Description    string        `json:"description" validate:"required,min=200,max=2000"`
	Tags           []string      `json:"tags" validate:"required,min=1,max=3,dive,omitempty,min=3,max=20"`
	Links          []string      `json:"links" validate:"required,min=1,max=3,dive,omitempty,url"`
	Budget         int           `json:"budget" validate:"required,min=1000,max=1000000"`
	Installments   int64         `json:"installments" validate:"required,min=1,max=12"`
	TimeZone       string        `json:"timezone"` // todo: figure out timezone check
	TokensAccepted []Network     `json:"tokens_accepted" validate:"required,min=1"`
	StickyDuration int64         `json:"sticky_duration" validate:"omitempty,lte=30"`
	CreatedAt      int64         `json:"created_at"`
	TxHash         string        `json:"tx_hash" validate:"required"`
	ImageUrl       string        `json:"image_url" validate:"url"`
	Applications   []Application `json:"application"`
}

type Application struct {
	UserAddress string `json:"user_address" validate:"required"`
	JobId       string `json:"job_id" validate:"required"`
	CoverLetter string `json:"cover_letter" validate:"required"`
	Date        int64  `json:"date" validate:"required"`
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

type Tags struct {
	Tags []string `json:"tags"`
}
