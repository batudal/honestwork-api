package schema

// todo: abstract go-playground/validator/v10 props to have clean schema
type User struct {
	Salt         string        `json:"salt"`
	Username     string        `json:"username" validate:"required,min=5,max=50"`
	ShowEns      *bool         `json:"show_ens" validate:"required,boolean"`
	EnsName      string        `json:"ens_name"` // custom
	Title        string        `json:"title" validate:"required,min=5,max=50"`
	ImageUrl     string        `json:"image_url" validate:"omitempty,url"`
	ShowNFT      *bool         `json:"show_nft" validate:"boolean"`
	NFTUrl       string        `json:"nft_url" validate:"omitempty,url"`
	NFTAddress   string        `json:"nft_address" validate:"omitempty,eth_addr"`
	NFTId        string        `json:"nft_id"` // custom
	Email        string        `json:"email" validate:"omitempty,email"`
	Timezone     string        `json:"timezone" validate:"oneof='UTC-12' 'UTC-11' 'UTC-10' 'UTC-9' 'UTC-8' 'UTC-7' 'UTC-6' 'UTC-5' 'UTC-4' 'UTC-3' 'UTC-2' 'UTC-1' 'UTC' 'UTC+1' 'UTC+2' 'UTC+3' 'UTC+4' 'UTC+5' 'UTC+6' 'UTC+7' 'UTC+8' 'UTC+9' 'UTC+10' 'UTC+11' 'UTC+12' 'UTC+13' 'UTC+14'"`
	Bio          string        `json:"bio"` // custom
	Links        []string      `json:"links" validate:"required,min=1,max=3,dive,omitempty,url"`
	Rating       int64         `json:"rating"`
	Watchlist    []*Watchlist  `json:"watchlist"`
	Favorites    []*Favorite   `json:"favorites"`
	DmsOpen      *bool         `json:"dms_open" validate:"required,boolean"`
	Applications []Application `json:"application"`
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
	Slot         int      `json:"slot"`
	UserAddress  string   `json:"user_address" validate:"required,eth_addr"`
	Title        string   `json:"title" validate:"required,min=5,max=50"`
	Description  string   `json:"description"` // custom
	Tags         []string `json:"tags" validate:"required,min=1,max=3,dive,omitempty,min=2,max=20"`
	Links        []string `json:"links" validate:"required,min=1,max=3,dive,omitempty,url"`
	ImageUrls    []string `json:"image_urls" validate:"required,min=1,max=8,dive,omitempty,url"`
	MinimumPrice int      `json:"minimum_price" validate:"required,min=10,max=10000"`
	Publish      bool     `json:"publish" validate:"boolean"`
	CreatedAt    int64    `json:"created_at"`
}

type Job struct {
	Email          string        `json:"email" validate:"required,email"`
	Slot           int           `json:"slot"`
	UserAddress    string        `json:"user_address" validate:"required,eth_addr"`
	Username       string        `json:"username" validate:"required,min=5,max=50"`
	TokenPaid      string        `json:"token_paid" validate:"eth_addr"`
	Title          string        `json:"title" validate:"required,min=5,max=50"`
	Description    string        `json:"description" validate:"required"` // custom
	Tags           []string      `json:"tags" validate:"required,min=1,max=3,dive,omitempty,min=2,max=20"`
	Links          []string      `json:"links" validate:"required,min=1,max=3,dive,omitempty,url"`
	Budget         int           `json:"budget" validate:"required,min=200,max=100000"`
	Timezone       int           `json:"timezone" validate:"required,min=-14,max=14"`
	TokensAccepted []Network     `json:"tokens_accepted" validate:"required,min=1"`
	StickyDuration int64         `json:"sticky_duration" validate:"omitempty,lte=30"`
	CreatedAt      int64         `json:"created_at"`
	TxHash         string        `json:"tx_hash" validate:"required"`
	ImageUrl       string        `json:"image_url" validate:"url"`
	Applications   []Application `json:"application"`
	DealNetworkId  int           `json:"deal_network_id"`
	DealId         int           `json:"deal_id"`
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

type Tags struct {
	Tags []string `json:"tags"`
}

type Conversation struct {
	MatchedUser   string `json:"matched_user"`
	CreatedAt     int64  `json:"created_at"`
	LastMessageAt int64  `json:"last_message_at"`
	Muted         bool   `json:"muted"`
}

type Deal struct {
	Status       string `json:"status"`
	Signature    string `json:"signature"`
	Network      string `json:"network"`
	TokenAddress string `json:"token_address"`
	TotalAmount  string `json:"total_amount"`
	DownPayment  string `json:"downpayment"`
	JobId        int    `json:"job_id"`
}

type Whitelist struct {
	Addresses []string `json:"addresses"`
}
