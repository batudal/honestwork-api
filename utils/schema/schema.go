package schema

type User struct {
	Salt         string        `json:"salt" bson:"salt"`
	Username     string        `json:"username" validate:"required,min=5,max=50" bson:"username"`
	ShowEns      *bool         `json:"show_ens" validate:"required,boolean" bson:"show_ens"`
	EnsName      string        `json:"ens_name" bson:"ens_name"` // custom
	Title        string        `json:"title" validate:"required,min=5,max=50" bson:"title"`
	ImageUrl     string        `json:"image_url" validate:"omitempty,url" bson:"image_url"`
	ShowNFT      *bool         `json:"show_nft" validate:"boolean" bson:"show_nft"`
	NFTUrl       string        `json:"nft_url" validate:"omitempty,url" bson:"nft_url"`
	NFTAddress   string        `json:"nft_address" validate:"omitempty,eth_addr" bson:"nft_address"`
	NFTId        string        `json:"nft_id" bson:"nft_id"` // custom
	Email        string        `json:"email" validate:"omitempty,email" bson:"email"`
	Timezone     *int64        `json:"timezone" validate:"required,min=-12,max=14" bson:"timezone"`
	Bio          string        `json:"bio" bson:"bio"` // custom
	Links        []string      `json:"links" validate:"required,min=1,max=3,dive,omitempty,url" bson:"links"`
	Rating       int64         `json:"rating" bson:"rating"`
	Watchlist    []*Watchlist  `json:"watchlist" bson:"watchlist"`
	Favorites    []*Favorite   `json:"favorites" bson:"favorites"`
	DmsOpen      *bool         `json:"dms_open" validate:"required,boolean" bson:"dms_open"`
	Applications []Application `json:"application" bson:"application"`
}

type FavoriteInput struct {
	Address string `json:"address" bson:"address"`
	Slot    int    `json:"slot" bson:"slot"`
}

type Favorite struct {
	Input    *FavoriteInput `json:"input" bson:"input"`
	Username string         `json:"username" bson:"username"`
	Title    string         `json:"title" bson:"title"`
	ImageUrl string         `json:"image_url" bson:"image_url"`
}

type Watchlist struct {
	Input    *WatchlistInput `json:"input" bson:"input"`
	Username string          `json:"username" bson:"username"`
	Title    string          `json:"title" bson:"title"`
	ImageUrl string          `json:"image_url" bson:"image_url"`
}

type WatchlistInput struct {
	Address string `json:"address" bson:"address"`
	Slot    int    `json:"slot" bson:"slot"`
}

type Skill struct {
	Slot         int      `json:"slot" bson:"slot"`
	UserAddress  string   `json:"user_address" validate:"required,eth_addr" bson:"user_address"`
	Title        string   `json:"title" validate:"required,min=5,max=50" bson:"title"`
	Description  string   `json:"description" bson:"description"` // custom
	Tags         []string `json:"tags" validate:"required,min=1,max=3,dive,omitempty,min=2,max=20" bson:"tags"`
	Links        []string `json:"links" validate:"required,min=1,max=3,dive,omitempty,url" bson:"links"`
	ImageUrls    []string `json:"image_urls" validate:"required,min=1,max=8,dive,omitempty,url" bson:"image_urls"`
	MinimumPrice int      `json:"minimum_price" validate:"required,min=10,max=10000" bson:"minimum_price"`
	Publish      bool     `json:"publish" validate:"boolean" bson:"publish"`
	CreatedAt    int64    `json:"created_at" bson:"created_at"`
}

type Job struct {
	Email          string        `json:"email" validate:"required,email" bson:"email"`
	Slot           int           `json:"slot" bson:"slot"`
	UserAddress    string        `json:"user_address" validate:"required,eth_addr" bson:"user_address"`
	Username       string        `json:"username" validate:"required,min=5,max=50" bson:"username"`
	TokenPaid      string        `json:"token_paid" validate:"eth_addr" bson:"token_paid"`
	Title          string        `json:"title" validate:"required,min=5,max=50" bson:"title"`
	Description    string        `json:"description" validate:"required" bson:"description"` // custom
	Tags           []string      `json:"tags" validate:"required,min=1,max=3,dive,omitempty,min=2,max=20" bson:"tags"`
	Links          []string      `json:"links" validate:"required,min=1,max=3,dive,omitempty,url" bson:"links"`
	Budget         int           `json:"budget" validate:"required,min=200,max=100000" bson:"budget"`
	Timezone       *int64        `json:"timezone" validate:"required,min=-12,max=14" bson:"timezone"`
	TokensAccepted []Network     `json:"tokens_accepted" validate:"required,min=1" bson:"tokens_accepted"`
	StickyDuration int64         `json:"sticky_duration" validate:"omitempty,lte=30" bson:"sticky_duration"`
	CreatedAt      int64         `json:"created_at" bson:"created_at"`
	TxHash         string        `json:"tx_hash" validate:"required" bson:"tx_hash"`
	ImageUrl       string        `json:"image_url" validate:"url" bson:"image_url"`
	Applications   []Application `json:"application" bson:"application"`
	DealNetworkId  int           `json:"deal_network_id" bson:"deal_network_id"`
	DealId         int           `json:"deal_id" bson:"deal_id"`
}

type Application struct {
	UserAddress string `json:"user_address" validate:"required" bson:"user_address"`
	JobId       string `json:"job_id" validate:"required" bson:"job_id"`
	CoverLetter string `json:"cover_letter" validate:"required" bson:"cover_letter"`
	Date        int64  `json:"date" validate:"required" bson:"date"`
}

type Network struct {
	Id     int64   `json:"id" bson:"id"`
	Tokens []Token `json:"tokens" bson:"tokens"`
}
type Token struct {
	Address string `json:"address" bson:"address"`
}

type Tags struct {
	Tags []string `json:"tags" bson:"tags"`
}

type Conversation struct {
	MatchedUser   string `json:"matched_user" bson:"matched_user"`
	CreatedAt     int64  `json:"created_at" bson:"created_at"`
	LastMessageAt int64  `json:"last_message_at" bson:"last_message_at"`
	Muted         bool   `json:"muted" bson:"muted"`
}

type Deal struct {
	Status       string `json:"status" bson:"status"`
	Signature    string `json:"signature" bson:"signature"`
	Network      string `json:"network" bson:"network"`
	TokenAddress string `json:"token_address" bson:"token_address"`
	TotalAmount  string `json:"total_amount" bson:"total_amount"`
	DownPayment  string `json:"downpayment" bson:"downpayment"`
	JobId        int    `json:"job_id" bson:"job_id"`
}

type Whitelist struct {
	Addresses []string `json:"addresses" bson:"addresses"`
}
