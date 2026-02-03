package stream

// TokenActivityType represents the type of token activity
type TokenActivityType string

const (
	TokenActivityTypeSell            TokenActivityType = "sell"
	TokenActivityTypeBuy             TokenActivityType = "buy"
	TokenActivityTypeAddLiquidity    TokenActivityType = "add_liquidity"
	TokenActivityTypeRemoveLiquidity TokenActivityType = "remove_liquidity"
)

// ChannelType represents the type of channel
type ChannelType string

const (
	ChannelTypeNew       ChannelType = "new"
	ChannelTypeHot       ChannelType = "trending"
	ChannelTypeUSStocks  ChannelType = "us_stocks"
	ChannelTypeCompleted ChannelType = "completed"
	ChannelTypeGraduated ChannelType = "graduated"
)

// MetricType represents the type of metric
type MetricType string

const (
	MetricTypeLiquidityInUSD MetricType = "liquidity_in_usd"
	MetricTypeMigratedRatio  MetricType = "migrated_ratio"
)

// RankingType represents the type of ranking
type RankingType string

const (
	RankingTypeNew          RankingType = "new"
	RankingTypeHot          RankingType = "trending"
	RankingTypeStocks       RankingType = "stocks"
	RankingTypeFinalStretch RankingType = "completed"
	RankingTypeMigrated     RankingType = "graduated"
)

// Dex represents the type of decentralized exchange
type Dex string

const (
	DexPumpFun                    Dex = "pump_fun"
	DexRaydiumLaunchpad           Dex = "raydium_launchpad"
	DexMeteorDynamicBoundingCurve Dex = "meteora_dynamic_bounding_curve"
	DexBonkFun                    Dex = "bonk_fun"
	DexBoopFun                    Dex = "boop_fun"
	DexMoonitFun                  Dex = "moonit_fun"
)

// TokenActivity represents token activity data
type TokenActivity struct {
	Address   string            `json:"address"`
	PriceUsd  string            `json:"priceUsd"`
	Amount    string            `json:"amount"`
	Type      TokenActivityType `json:"type"`
	TxHash    string            `json:"txHash"`
	Timestamp int64             `json:"timestamp"`
}

// TokenStat represents token statistics
type TokenStat struct {
	Address   string `json:"address"`
	Timestamp int64  `json:"timestamp"`

	// 1-minute data
	Buys1m            *int    `json:"buys1m,omitempty"`
	Sells1m           *int    `json:"sells1m,omitempty"`
	Buyers1m          *int    `json:"buyers1m,omitempty"`
	Sellers1m         *int    `json:"sellers1m,omitempty"`
	BuyVolumeInUsd1m  *string `json:"buyVolumeInUsd1m,omitempty"`
	SellVolumeInUsd1m *string `json:"sellVolumeInUsd1m,omitempty"`
	Price1m           *string `json:"price1m,omitempty"`
	OpenInUsd1m       *string `json:"openInUsd1m,omitempty"`
	CloseInUsd1m      *string `json:"closeInUsd1m,omitempty"`

	// 5-minute data
	Buys5m            *int    `json:"buys5m,omitempty"`
	Sells5m           *int    `json:"sells5m,omitempty"`
	Buyers5m          *int    `json:"buyers5m,omitempty"`
	Sellers5m         *int    `json:"sellers5m,omitempty"`
	BuyVolumeInUsd5m  *string `json:"buyVolumeInUsd5m,omitempty"`
	SellVolumeInUsd5m *string `json:"sellVolumeInUsd5m,omitempty"`
	Price5m           *string `json:"price5m,omitempty"`
	OpenInUsd5m       *string `json:"openInUsd5m,omitempty"`
	CloseInUsd5m      *string `json:"closeInUsd5m,omitempty"`

	// 15-minute data
	Buys15m            *int    `json:"buys15m,omitempty"`
	Sells15m           *int    `json:"sells15m,omitempty"`
	Buyers15m          *int    `json:"buyers15m,omitempty"`
	Sellers15m         *int    `json:"sellers15m,omitempty"`
	BuyVolumeInUsd15m  *string `json:"buyVolumeInUsd15m,omitempty"`
	SellVolumeInUsd15m *string `json:"sellVolumeInUsd15m,omitempty"`
	Price15m           *string `json:"price15m,omitempty"`
	OpenInUsd15m       *string `json:"openInUsd15m,omitempty"`
	CloseInUsd15m      *string `json:"closeInUsd15m,omitempty"`

	// 30-minute data
	Buys30m            *int    `json:"buys30m,omitempty"`
	Sells30m           *int    `json:"sells30m,omitempty"`
	Buyers30m          *int    `json:"buyers30m,omitempty"`
	Sellers30m         *int    `json:"sellers30m,omitempty"`
	BuyVolumeInUsd30m  *string `json:"buyVolumeInUsd30m,omitempty"`
	SellVolumeInUsd30m *string `json:"sellVolumeInUsd30m,omitempty"`
	Price30m           *string `json:"price30m,omitempty"`
	OpenInUsd30m       *string `json:"openInUsd30m,omitempty"`
	CloseInUsd30m      *string `json:"closeInUsd30m,omitempty"`

	// 1-hour data
	Buys1h            *int    `json:"buys1h,omitempty"`
	Sells1h           *int    `json:"sells1h,omitempty"`
	Buyers1h          *int    `json:"buyers1h,omitempty"`
	Sellers1h         *int    `json:"sellers1h,omitempty"`
	BuyVolumeInUsd1h  *string `json:"buyVolumeInUsd1h,omitempty"`
	SellVolumeInUsd1h *string `json:"sellVolumeInUsd1h,omitempty"`
	Price1h           *string `json:"price1h,omitempty"`
	OpenInUsd1h       *string `json:"openInUsd1h,omitempty"`
	CloseInUsd1h      *string `json:"closeInUsd1h,omitempty"`

	// 4-hour data
	Buys4h            *int    `json:"buys4h,omitempty"`
	Sells4h           *int    `json:"sells4h,omitempty"`
	Buyers4h          *int    `json:"buyers4h,omitempty"`
	Sellers4h         *int    `json:"sellers4h,omitempty"`
	BuyVolumeInUsd4h  *string `json:"buyVolumeInUsd4h,omitempty"`
	SellVolumeInUsd4h *string `json:"sellVolumeInUsd4h,omitempty"`
	Price4h           *string `json:"price4h,omitempty"`
	OpenInUsd4h       *string `json:"openInUsd4h,omitempty"`
	CloseInUsd4h      *string `json:"closeInUsd4h,omitempty"`

	// 24-hour data
	Buys24h            *int    `json:"buys24h,omitempty"`
	Sells24h           *int    `json:"sells24h,omitempty"`
	Buyers24h          *int    `json:"buyers24h,omitempty"`
	Sellers24h         *int    `json:"sellers24h,omitempty"`
	BuyVolumeInUsd24h  *string `json:"buyVolumeInUsd24h,omitempty"`
	SellVolumeInUsd24h *string `json:"sellVolumeInUsd24h,omitempty"`
	Price24h           *string `json:"price24h,omitempty"`
	OpenInUsd24h       *string `json:"openInUsd24h,omitempty"`
	CloseInUsd24h      *string `json:"closeInUsd24h,omitempty"`

	// Current price
	Price *string `json:"price,omitempty"`
}

// TokenHolder represents token holder information
type TokenHolder struct {
	TokenAddress    string  `json:"tokenAddress"`
	Holders         *int    `json:"holders,omitempty"`
	Top100Amount    *string `json:"top100Amount,omitempty"`
	Top10Amount     *string `json:"top10Amount,omitempty"`
	Top100Holders   *int    `json:"top100Holders,omitempty"`
	Top10Holders    *int    `json:"top10Holders,omitempty"`
	Top100Ratio     *string `json:"top100Ratio,omitempty"`
	Top10Ratio      *string `json:"top10Ratio,omitempty"`
	CreatorsHolders *int    `json:"creatorsHolders,omitempty"`
	CreatorsAmount  *string `json:"creatorsAmount,omitempty"`
	CreatorsRatio   *string `json:"creatorsRatio,omitempty"`
	Timestamp       int64   `json:"timestamp"`
}

// WalletBalance represents wallet balance data
type WalletBalance struct {
	WalletAddress   string `json:"walletAddress"`
	TokenAddress    string `json:"tokenAddress"`
	TokenPriceInUsd string `json:"tokenPriceInUsd"`
	Balance         string `json:"balance"`
	Timestamp       int64  `json:"timestamp"`
}

// WalletPnl represents wallet profit and loss data
type WalletPnl struct {
	WalletAddress        string `json:"walletAddress"`
	Buys                 int    `json:"buys"`
	BuyAmount            string `json:"buyAmount"`
	BuyAmountInUsd       string `json:"buyAmountInUsd"`
	AverageBuyPriceInUsd string `json:"averageBuyPriceInUsd"`
	SellAmount           string `json:"sellAmount"`
	SellAmountInUsd      string `json:"sellAmountInUsd"`
	Sells                int    `json:"sells"`
	Wins                 int    `json:"wins"`
	WinRatio             string `json:"winRatio"`
	PnlInUsd             string `json:"pnlInUsd"`
	AveragePnlInUsd      string `json:"averagePnlInUsd"`
	PnlRatio             string `json:"pnlRatio"`
	ProfitableDays       int    `json:"profitableDays"`
	LosingDays           int    `json:"losingDays"`
	Tokens               int    `json:"tokens"`
	Resolution           string `json:"resolution"`
}

// NewToken represents new token data
type NewToken struct {
	TokenAddress string       `json:"tokenAddress"`
	Name         string       `json:"name"`
	Symbol       string       `json:"symbol"`
	Decimals     *int         `json:"decimals,omitempty"`
	LaunchFrom   *DexProtocol `json:"launchFrom,omitempty"`
	CreatedAtMs  int64        `json:"createdAtMs"`
}

// TokenSupply represents token supply data
type TokenSupply struct {
	TokenAddress   string  `json:"tokenAddress"`
	Supply         *string `json:"supply,omitempty"`
	MarketCapInUsd *string `json:"marketCapInUsd,omitempty"`
	Timestamp      int64   `json:"timestamp"`
}

// DexPoolBalance represents DEX pool balance data
type DexPoolBalance struct {
	PoolAddress          string `json:"poolAddress"`
	TokenAAddress        string `json:"tokenAAddress"`
	TokenALiquidityInUsd string `json:"tokenALiquidityInUsd"`
	TokenBAddress        string `json:"tokenBAddress"`
	TokenBLiquidityInUsd string `json:"tokenBLiquidityInUsd"`
}

// TokenLiquidity represents token liquidity data
type TokenLiquidity struct {
	TokenAddress string     `json:"tokenAddress"`
	MetricType   MetricType `json:"metricType"`
	Value        string     `json:"value"`
	Timestamp    int64      `json:"timestamp"`
}

// TokenMaxLiquidity represents token max liquidity data (max liquidity in a single pool)
type TokenMaxLiquidity struct {
	TokenAddress      string `json:"tokenAddress"`
	PoolAddress       string `json:"poolAddress"`
	LiquidityInUsd    string `json:"liquidityInUsd"`
	LiquidityInNative string `json:"liquidityInNative"`
	Timestamp         int64  `json:"timestamp"`
}

// TokenTotalLiquidity represents token total liquidity data (total liquidity across all pools)
type TokenTotalLiquidity struct {
	TokenAddress      string `json:"tokenAddress"`
	LiquidityInUsd    string `json:"liquidityInUsd"`
	LiquidityInNative string `json:"liquidityInNative"`
	PoolCount         int    `json:"poolCount"`
	Timestamp         int64  `json:"timestamp"`
}

// DexProtocol represents DEX protocol information
type DexProtocol struct {
	ProgramAddress *string `json:"programAddress,omitempty"`
	ProtocolFamily *string `json:"protocolFamily,omitempty"`
	ProtocolName   *string `json:"protocolName,omitempty"`
}

// TokenBondingCurve represents token bonding curve data
type TokenBondingCurve struct {
	TokenAddress  *string `json:"tokenAddress,omitempty"`
	ProgressRatio *string `json:"progressRatio,omitempty"`
}

// SocialMedia represents social media links
type SocialMedia struct {
	Twitter   *string `json:"twitter,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
	Website   *string `json:"website,omitempty"`
	Tiktok    *string `json:"tiktok,omitempty"`
	Discord   *string `json:"discord,omitempty"`
	Facebook  *string `json:"facebook,omitempty"`
	Github    *string `json:"github,omitempty"`
	Instagram *string `json:"instagram,omitempty"`
	Linkedin  *string `json:"linkedin,omitempty"`
	Medium    *string `json:"medium,omitempty"`
	Reddit    *string `json:"reddit,omitempty"`
	Youtube   *string `json:"youtube,omitempty"`
	Bitbucket *string `json:"bitbucket,omitempty"`
}

// TokenMetadata represents token metadata
type TokenMetadata struct {
	TokenAddress string       `json:"tokenAddress"`
	Name         *string      `json:"name,omitempty"`
	Decimals     *int         `json:"decimals,omitempty"`
	Symbol       *string      `json:"symbol,omitempty"`
	ImageUrl     *string      `json:"imageUrl,omitempty"`
	Description  *string      `json:"description,omitempty"`
	SocialMedia  *SocialMedia `json:"socialMedia,omitempty"`
	CreatedAtMs  *int64       `json:"createdAtMs,omitempty"`
	LaunchFrom   *DexProtocol `json:"launchFrom,omitempty"`
	MigratedTo   *DexProtocol `json:"migratedTo,omitempty"`
}

// TokenCandle represents token candlestick data
type TokenCandle struct {
	Open       string `json:"open"`
	Close      string `json:"close"`
	High       string `json:"high"`
	Low        string `json:"low"`
	Volume     string `json:"volume"`
	Resolution string `json:"resolution"`
	Time       int64  `json:"time"`
	Number     int    `json:"number"`
}

// TradeActivity represents trade activity data
type TradeActivity struct {
	TokenAddress      string `json:"tokenAddress"`
	Timestamp         int64  `json:"timestamp"`
	Kind              string `json:"kind"`
	BuyAmount         string `json:"buyAmount"`
	BuyAmountInUsd    string `json:"buyAmountInUsd"`
	BuyTokenAddress   string `json:"buyTokenAddress"`
	BuyTokenName      string `json:"buyTokenName"`
	BuyTokenSymbol    string `json:"buyTokenSymbol"`
	BuyWalletAddress  string `json:"buyWalletAddress"`
	SellAmount        string `json:"sellAmount"`
	SellAmountInUsd   string `json:"sellAmountInUsd"`
	SellTokenAddress  string `json:"sellTokenAddress"`
	SellTokenName     string `json:"sellTokenName"`
	SellTokenSymbol   string `json:"sellTokenSymbol"`
	SellWalletAddress string `json:"sellWalletAddress"`
	TxHash            string `json:"txHash"`
}

// WalletTokenPnl represents wallet token profit and loss data
type WalletTokenPnl struct {
	WalletAddress            string `json:"walletAddress"`
	TokenAddress             string `json:"tokenAddress"`
	TokenPriceInUsd          string `json:"tokenPriceInUsd"`
	Timestamp                int64  `json:"timestamp"`
	OpenTime                 int64  `json:"opentime"`
	LastTime                 int64  `json:"lasttime"`
	CloseTime                int64  `json:"closetime"`
	BuyAmount                string `json:"buyAmount"`
	BuyAmountInUsd           string `json:"buyAmountInUsd"`
	BuyCount                 int    `json:"buyCount"`
	BuyCount30d              int    `json:"buyCount30d"`
	BuyCount7d               int    `json:"buyCount7d"`
	SellAmount               string `json:"sellAmount"`
	SellAmountInUsd          string `json:"sellAmountInUsd"`
	SellCount                int    `json:"sellCount"`
	SellCount30d             int    `json:"sellCount30d"`
	SellCount7d              int    `json:"sellCount7d"`
	HeldDurationTimestamp    int64  `json:"heldDurationTimestamp"`
	AverageBuyPriceInUsd     string `json:"averageBuyPriceInUsd"`
	AverageSellPriceInUsd    string `json:"averageSellPriceInUsd"`
	UnrealizedProfitInUsd    string `json:"unrealizedProfitInUsd"`
	UnrealizedProfitRatio    string `json:"unrealizedProfitRatio"`
	RealizedProfitInUsd      string `json:"realizedProfitInUsd"`
	RealizedProfitRatio      string `json:"realizedProfitRatio"`
	TotalRealizedProfitInUsd string `json:"totalRealizedProfitInUsd"`
	TotalRealizedProfitRatio string `json:"totalRealizedProfitRatio"`
}

// RankingTokenList represents ranking token list data
type RankingTokenList struct {
	Metadata     *TokenMetadata     `json:"metadata,omitempty"`
	Holder       *TokenHolder       `json:"holder,omitempty"`
	Supply       *TokenSupply       `json:"supply,omitempty"`
	Stat         *TokenStat         `json:"stat,omitempty"`
	BondingCurve *TokenBondingCurve `json:"bondingCurve,omitempty"`
}
