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
	ChannelTypeFinalStretch ChannelType = "graduated"
	ChannelTypeMigrated    ChannelType = "completed"
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

// TokenStat represents token statistics with multi-timeframe data
// Time windows: 1m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 24h, 1W, 1M
type TokenStat struct {
	Address   string `json:"address"`
	Timestamp int64  `json:"timestamp"`

	// Current price (first positive closeInUsd across windows, or from cache)
	Price *string `json:"price,omitempty"`

	// 1-minute data
	Buys1m                 *int    `json:"buys1m,omitempty"`
	Sells1m                *int    `json:"sells1m,omitempty"`
	Buyers1m               *int    `json:"buyers1m,omitempty"`
	Sellers1m              *int    `json:"sellers1m,omitempty"`
	BuyVolumeInUsd1m       *string `json:"buyVolumeInUsd1m,omitempty"`
	SellVolumeInUsd1m      *string `json:"sellVolumeInUsd1m,omitempty"`
	Price1m                *string `json:"price1m,omitempty"`
	OpenInUsd1m            *string `json:"openInUsd1m,omitempty"`
	CloseInUsd1m           *string `json:"closeInUsd1m,omitempty"`
	VolumeChangeRatio1m    *string `json:"volumeChangeRatio1m,omitempty"`
	Trades1m               *int    `json:"trades1m,omitempty"`
	DappProgramCount1m     *int    `json:"dappProgramCount1m,omitempty"`
	PoolCount1m            *int    `json:"poolCount1m,omitempty"`
	LiquidityInUsd1m       *string `json:"liquidityInUsd1m,omitempty"`
	LiquidityChangeRatio1m *string `json:"liquidityChangeRatio1m,omitempty"`

	// 5-minute data
	Buys5m                 *int    `json:"buys5m,omitempty"`
	Sells5m                *int    `json:"sells5m,omitempty"`
	Buyers5m               *int    `json:"buyers5m,omitempty"`
	Sellers5m              *int    `json:"sellers5m,omitempty"`
	BuyVolumeInUsd5m       *string `json:"buyVolumeInUsd5m,omitempty"`
	SellVolumeInUsd5m      *string `json:"sellVolumeInUsd5m,omitempty"`
	Price5m                *string `json:"price5m,omitempty"`
	OpenInUsd5m            *string `json:"openInUsd5m,omitempty"`
	CloseInUsd5m           *string `json:"closeInUsd5m,omitempty"`
	VolumeChangeRatio5m    *string `json:"volumeChangeRatio5m,omitempty"`
	Trades5m               *int    `json:"trades5m,omitempty"`
	DappProgramCount5m     *int    `json:"dappProgramCount5m,omitempty"`
	PoolCount5m            *int    `json:"poolCount5m,omitempty"`
	LiquidityInUsd5m       *string `json:"liquidityInUsd5m,omitempty"`
	LiquidityChangeRatio5m *string `json:"liquidityChangeRatio5m,omitempty"`

	// 15-minute data
	Buys15m                 *int    `json:"buys15m,omitempty"`
	Sells15m                *int    `json:"sells15m,omitempty"`
	Buyers15m               *int    `json:"buyers15m,omitempty"`
	Sellers15m              *int    `json:"sellers15m,omitempty"`
	BuyVolumeInUsd15m       *string `json:"buyVolumeInUsd15m,omitempty"`
	SellVolumeInUsd15m      *string `json:"sellVolumeInUsd15m,omitempty"`
	Price15m                *string `json:"price15m,omitempty"`
	OpenInUsd15m            *string `json:"openInUsd15m,omitempty"`
	CloseInUsd15m           *string `json:"closeInUsd15m,omitempty"`
	VolumeChangeRatio15m    *string `json:"volumeChangeRatio15m,omitempty"`
	Trades15m               *int    `json:"trades15m,omitempty"`
	DappProgramCount15m     *int    `json:"dappProgramCount15m,omitempty"`
	PoolCount15m            *int    `json:"poolCount15m,omitempty"`
	LiquidityInUsd15m       *string `json:"liquidityInUsd15m,omitempty"`
	LiquidityChangeRatio15m *string `json:"liquidityChangeRatio15m,omitempty"`

	// 30-minute data
	Buys30m                 *int    `json:"buys30m,omitempty"`
	Sells30m                *int    `json:"sells30m,omitempty"`
	Buyers30m               *int    `json:"buyers30m,omitempty"`
	Sellers30m              *int    `json:"sellers30m,omitempty"`
	BuyVolumeInUsd30m       *string `json:"buyVolumeInUsd30m,omitempty"`
	SellVolumeInUsd30m      *string `json:"sellVolumeInUsd30m,omitempty"`
	Price30m                *string `json:"price30m,omitempty"`
	OpenInUsd30m            *string `json:"openInUsd30m,omitempty"`
	CloseInUsd30m           *string `json:"closeInUsd30m,omitempty"`
	VolumeChangeRatio30m    *string `json:"volumeChangeRatio30m,omitempty"`
	Trades30m               *int    `json:"trades30m,omitempty"`
	DappProgramCount30m     *int    `json:"dappProgramCount30m,omitempty"`
	PoolCount30m            *int    `json:"poolCount30m,omitempty"`
	LiquidityInUsd30m       *string `json:"liquidityInUsd30m,omitempty"`
	LiquidityChangeRatio30m *string `json:"liquidityChangeRatio30m,omitempty"`

	// 1-hour data
	Buys1h                 *int    `json:"buys1h,omitempty"`
	Sells1h                *int    `json:"sells1h,omitempty"`
	Buyers1h               *int    `json:"buyers1h,omitempty"`
	Sellers1h              *int    `json:"sellers1h,omitempty"`
	BuyVolumeInUsd1h       *string `json:"buyVolumeInUsd1h,omitempty"`
	SellVolumeInUsd1h      *string `json:"sellVolumeInUsd1h,omitempty"`
	Price1h                *string `json:"price1h,omitempty"`
	OpenInUsd1h            *string `json:"openInUsd1h,omitempty"`
	CloseInUsd1h           *string `json:"closeInUsd1h,omitempty"`
	VolumeChangeRatio1h    *string `json:"volumeChangeRatio1h,omitempty"`
	Trades1h               *int    `json:"trades1h,omitempty"`
	DappProgramCount1h     *int    `json:"dappProgramCount1h,omitempty"`
	PoolCount1h            *int    `json:"poolCount1h,omitempty"`
	LiquidityInUsd1h       *string `json:"liquidityInUsd1h,omitempty"`
	LiquidityChangeRatio1h *string `json:"liquidityChangeRatio1h,omitempty"`

	// 2-hour data
	Buys2h                 *int    `json:"buys2h,omitempty"`
	Sells2h                *int    `json:"sells2h,omitempty"`
	Buyers2h               *int    `json:"buyers2h,omitempty"`
	Sellers2h              *int    `json:"sellers2h,omitempty"`
	BuyVolumeInUsd2h       *string `json:"buyVolumeInUsd2h,omitempty"`
	SellVolumeInUsd2h      *string `json:"sellVolumeInUsd2h,omitempty"`
	Price2h                *string `json:"price2h,omitempty"`
	OpenInUsd2h            *string `json:"openInUsd2h,omitempty"`
	CloseInUsd2h           *string `json:"closeInUsd2h,omitempty"`
	VolumeChangeRatio2h    *string `json:"volumeChangeRatio2h,omitempty"`
	Trades2h               *int    `json:"trades2h,omitempty"`
	DappProgramCount2h     *int    `json:"dappProgramCount2h,omitempty"`
	PoolCount2h            *int    `json:"poolCount2h,omitempty"`
	LiquidityInUsd2h       *string `json:"liquidityInUsd2h,omitempty"`
	LiquidityChangeRatio2h *string `json:"liquidityChangeRatio2h,omitempty"`

	// 4-hour data
	Buys4h                 *int    `json:"buys4h,omitempty"`
	Sells4h                *int    `json:"sells4h,omitempty"`
	Buyers4h               *int    `json:"buyers4h,omitempty"`
	Sellers4h              *int    `json:"sellers4h,omitempty"`
	BuyVolumeInUsd4h       *string `json:"buyVolumeInUsd4h,omitempty"`
	SellVolumeInUsd4h      *string `json:"sellVolumeInUsd4h,omitempty"`
	Price4h                *string `json:"price4h,omitempty"`
	OpenInUsd4h            *string `json:"openInUsd4h,omitempty"`
	CloseInUsd4h           *string `json:"closeInUsd4h,omitempty"`
	VolumeChangeRatio4h    *string `json:"volumeChangeRatio4h,omitempty"`
	Trades4h               *int    `json:"trades4h,omitempty"`
	DappProgramCount4h     *int    `json:"dappProgramCount4h,omitempty"`
	PoolCount4h            *int    `json:"poolCount4h,omitempty"`
	LiquidityInUsd4h       *string `json:"liquidityInUsd4h,omitempty"`
	LiquidityChangeRatio4h *string `json:"liquidityChangeRatio4h,omitempty"`

	// 6-hour data
	Buys6h                 *int    `json:"buys6h,omitempty"`
	Sells6h                *int    `json:"sells6h,omitempty"`
	Buyers6h               *int    `json:"buyers6h,omitempty"`
	Sellers6h              *int    `json:"sellers6h,omitempty"`
	BuyVolumeInUsd6h       *string `json:"buyVolumeInUsd6h,omitempty"`
	SellVolumeInUsd6h      *string `json:"sellVolumeInUsd6h,omitempty"`
	Price6h                *string `json:"price6h,omitempty"`
	OpenInUsd6h            *string `json:"openInUsd6h,omitempty"`
	CloseInUsd6h           *string `json:"closeInUsd6h,omitempty"`
	VolumeChangeRatio6h    *string `json:"volumeChangeRatio6h,omitempty"`
	Trades6h               *int    `json:"trades6h,omitempty"`
	DappProgramCount6h     *int    `json:"dappProgramCount6h,omitempty"`
	PoolCount6h            *int    `json:"poolCount6h,omitempty"`
	LiquidityInUsd6h       *string `json:"liquidityInUsd6h,omitempty"`
	LiquidityChangeRatio6h *string `json:"liquidityChangeRatio6h,omitempty"`

	// 8-hour data
	Buys8h                 *int    `json:"buys8h,omitempty"`
	Sells8h                *int    `json:"sells8h,omitempty"`
	Buyers8h               *int    `json:"buyers8h,omitempty"`
	Sellers8h              *int    `json:"sellers8h,omitempty"`
	BuyVolumeInUsd8h       *string `json:"buyVolumeInUsd8h,omitempty"`
	SellVolumeInUsd8h      *string `json:"sellVolumeInUsd8h,omitempty"`
	Price8h                *string `json:"price8h,omitempty"`
	OpenInUsd8h            *string `json:"openInUsd8h,omitempty"`
	CloseInUsd8h           *string `json:"closeInUsd8h,omitempty"`
	VolumeChangeRatio8h    *string `json:"volumeChangeRatio8h,omitempty"`
	Trades8h               *int    `json:"trades8h,omitempty"`
	DappProgramCount8h     *int    `json:"dappProgramCount8h,omitempty"`
	PoolCount8h            *int    `json:"poolCount8h,omitempty"`
	LiquidityInUsd8h       *string `json:"liquidityInUsd8h,omitempty"`
	LiquidityChangeRatio8h *string `json:"liquidityChangeRatio8h,omitempty"`

	// 24-hour data
	Buys24h                 *int    `json:"buys24h,omitempty"`
	Sells24h                *int    `json:"sells24h,omitempty"`
	Buyers24h               *int    `json:"buyers24h,omitempty"`
	Sellers24h              *int    `json:"sellers24h,omitempty"`
	BuyVolumeInUsd24h       *string `json:"buyVolumeInUsd24h,omitempty"`
	SellVolumeInUsd24h      *string `json:"sellVolumeInUsd24h,omitempty"`
	Price24h                *string `json:"price24h,omitempty"`
	OpenInUsd24h            *string `json:"openInUsd24h,omitempty"`
	CloseInUsd24h           *string `json:"closeInUsd24h,omitempty"`
	VolumeChangeRatio24h    *string `json:"volumeChangeRatio24h,omitempty"`
	Trades24h               *int    `json:"trades24h,omitempty"`
	DappProgramCount24h     *int    `json:"dappProgramCount24h,omitempty"`
	PoolCount24h            *int    `json:"poolCount24h,omitempty"`
	LiquidityInUsd24h       *string `json:"liquidityInUsd24h,omitempty"`
	LiquidityChangeRatio24h *string `json:"liquidityChangeRatio24h,omitempty"`

	// 1-week data (note: JSON keys use uppercase W, e.g., b1W)
	Buys1W                 *int    `json:"buys1W,omitempty"`
	Sells1W                *int    `json:"sells1W,omitempty"`
	Buyers1W               *int    `json:"buyers1W,omitempty"`
	Sellers1W              *int    `json:"sellers1W,omitempty"`
	BuyVolumeInUsd1W       *string `json:"buyVolumeInUsd1W,omitempty"`
	SellVolumeInUsd1W      *string `json:"sellVolumeInUsd1W,omitempty"`
	Price1W                *string `json:"price1W,omitempty"`
	OpenInUsd1W            *string `json:"openInUsd1W,omitempty"`
	CloseInUsd1W           *string `json:"closeInUsd1W,omitempty"`
	VolumeChangeRatio1W    *string `json:"volumeChangeRatio1W,omitempty"`
	Trades1W               *int    `json:"trades1W,omitempty"`
	DappProgramCount1W     *int    `json:"dappProgramCount1W,omitempty"`
	PoolCount1W            *int    `json:"poolCount1W,omitempty"`
	LiquidityInUsd1W       *string `json:"liquidityInUsd1W,omitempty"`
	LiquidityChangeRatio1W *string `json:"liquidityChangeRatio1W,omitempty"`

	// 1-month data (note: JSON keys use uppercase M, e.g., b1M)
	Buys1M                 *int    `json:"buys1M,omitempty"`
	Sells1M                *int    `json:"sells1M,omitempty"`
	Buyers1M               *int    `json:"buyers1M,omitempty"`
	Sellers1M              *int    `json:"sellers1M,omitempty"`
	BuyVolumeInUsd1M       *string `json:"buyVolumeInUsd1M,omitempty"`
	SellVolumeInUsd1M      *string `json:"sellVolumeInUsd1M,omitempty"`
	Price1M                *string `json:"price1M,omitempty"`
	OpenInUsd1M            *string `json:"openInUsd1M,omitempty"`
	CloseInUsd1M           *string `json:"closeInUsd1M,omitempty"`
	VolumeChangeRatio1M    *string `json:"volumeChangeRatio1M,omitempty"`
	Trades1M               *int    `json:"trades1M,omitempty"`
	DappProgramCount1M     *int    `json:"dappProgramCount1M,omitempty"`
	PoolCount1M            *int    `json:"poolCount1M,omitempty"`
	LiquidityInUsd1M       *string `json:"liquidityInUsd1M,omitempty"`
	LiquidityChangeRatio1M *string `json:"liquidityChangeRatio1M,omitempty"`
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
	TokenAddress    string       `json:"tokenAddress"`
	Name            string       `json:"name"`
	Symbol          string       `json:"symbol"`
	Decimals        *int         `json:"decimals,omitempty"`
	ImageUrl        *string      `json:"imageUrl,omitempty"`
	Description     *string      `json:"description,omitempty"`
	SocialMedia     *SocialMedia `json:"socialMedia,omitempty"`
	CoingeckoCoinId *string      `json:"coingeckoCoinId,omitempty"`
	LaunchFrom      *DexProtocol `json:"launchFrom,omitempty"`
	MigratedTo      *DexProtocol `json:"migratedTo,omitempty"`
	CreatedAtMs     int64        `json:"createdAtMs"`
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
	TokenAddress    string       `json:"tokenAddress"`
	Name            *string      `json:"name,omitempty"`
	Decimals        *int         `json:"decimals,omitempty"`
	Symbol          *string      `json:"symbol,omitempty"`
	ImageUrl        *string      `json:"imageUrl,omitempty"`
	Description     *string      `json:"description,omitempty"`
	SocialMedia     *SocialMedia `json:"socialMedia,omitempty"`
	CreatedAtMs     *int64       `json:"createdAtMs,omitempty"`
	CoingeckoCoinId *string      `json:"coingeckoCoinId,omitempty"`
	LaunchFrom      *DexProtocol `json:"launchFrom,omitempty"`
	MigratedTo      *DexProtocol `json:"migratedTo,omitempty"`
}

// PriceType represents the price type for candle data
type PriceType string

const (
	PriceTypeUSD    PriceType = "usd"
	PriceTypeNative PriceType = "native"
)

// Candle represents candlestick data for Token/Pool/Pair
type Candle struct {
	Address    string `json:"address"`
	Open       string `json:"open"`
	Close      string `json:"close"`
	High       string `json:"high"`
	Low        string `json:"low"`
	Volume     string `json:"volume"`
	Resolution string `json:"resolution"`
	Time       int64  `json:"time"`
	Number     int    `json:"number"`
}

// TokenCandle is an alias for backward compatibility
type TokenCandle = Candle

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
	WalletAddress            string  `json:"walletAddress"`
	TokenAddress             string  `json:"tokenAddress"`
	Timestamp                int64   `json:"timestamp"`
	BuyCount                 int     `json:"buyCount"`
	BuyCount30d              int     `json:"buyCount30d"`
	BuyCount7d               int     `json:"buyCount7d"`
	SellCount                int     `json:"sellCount"`
	SellCount30d             int     `json:"sellCount30d"`
	SellCount7d              int     `json:"sellCount7d"`
	TokenPriceInUsd          *string `json:"tokenPriceInUsd,omitempty"`
	OpenTime                 *int64  `json:"opentime,omitempty"`
	LastTime                 *int64  `json:"lasttime,omitempty"`
	CloseTime                *int64  `json:"closetime,omitempty"`
	BuyAmount                *string `json:"buyAmount,omitempty"`
	BuyAmountInUsd           *string `json:"buyAmountInUsd,omitempty"`
	SellAmount               *string `json:"sellAmount,omitempty"`
	SellAmountInUsd          *string `json:"sellAmountInUsd,omitempty"`
	HeldDurationTimestamp    *int64  `json:"heldDurationTimestamp,omitempty"`
	AverageBuyPriceInUsd     *string `json:"averageBuyPriceInUsd,omitempty"`
	AverageSellPriceInUsd    *string `json:"averageSellPriceInUsd,omitempty"`
	UnrealizedProfitInUsd    *string `json:"unrealizedProfitInUsd,omitempty"`
	UnrealizedProfitRatio    *string `json:"unrealizedProfitRatio,omitempty"`
	RealizedProfitInUsd      *string `json:"realizedProfitInUsd,omitempty"`
	RealizedProfitRatio      *string `json:"realizedProfitRatio,omitempty"`
	TotalRealizedProfitInUsd *string `json:"totalRealizedProfitInUsd,omitempty"`
	TotalRealizedProfitRatio *string `json:"totalRealizedProfitRatio,omitempty"`
}

// RankingTokenList represents ranking token list data
type RankingTokenList struct {
	Metadata     *TokenMetadata     `json:"metadata,omitempty"`
	Holder       *TokenHolder       `json:"holder,omitempty"`
	Supply       *TokenSupply       `json:"supply,omitempty"`
	Stat         *TokenStat         `json:"stat,omitempty"`
	BondingCurve *TokenBondingCurve `json:"bondingCurve,omitempty"`
}
