package stream

import (
	"regexp"
)

// FieldMapping represents field mapping
type FieldMapping map[string]string

// MethodFieldMappings represents method field mappings
type MethodFieldMappings map[string]FieldMapping

// CEL_FIELD_MAPPINGS represents CEL field mappings organized by subscription method
var CEL_FIELD_MAPPINGS = MethodFieldMappings{
	// Wallet balance subscription fields
	"subscribeWalletBalance": {
		"walletAddress":   "a",
		"tokenAddress":    "ta",
		"tokenPriceInUsd": "tpiu",
		"balance":         "b",
		"timestamp":       "t",
	},

	// Token candle subscription fields
	"subscribeTokenCandles": {
		"open":       "o",
		"close":      "c",
		"high":       "h",
		"low":        "l",
		"volume":     "v",
		"resolution": "r",
		"time":       "t",
		"number":     "n",
	},

	// Token stats subscription fields
	"subscribeTokenStats": {
		"address":            "a",
		"timestamp":          "t",
		"buys1m":             "b1m",
		"sells1m":            "s1m",
		"buyers1m":           "be1m",
		"sellers1m":          "se1m",
		"buyVolumeInUsd1m":   "bviu1m",
		"sellVolumeInUsd1m":  "sviu1m",
		"price1m":            "p1m",
		"openInUsd1m":        "oiu1m",
		"closeInUsd1m":       "ciu1m",
		"buys5m":             "b5m",
		"sells5m":            "s5m",
		"buyers5m":           "be5m",
		"sellers5m":          "se5m",
		"buyVolumeInUsd5m":   "bviu5m",
		"sellVolumeInUsd5m":  "sviu5m",
		"price5m":            "p5m",
		"openInUsd5m":        "oiu5m",
		"closeInUsd5m":       "ciu5m",
		"buys15m":            "b15m",
		"sells15m":           "s15m",
		"buyers15m":          "be15m",
		"sellers15m":         "se15m",
		"buyVolumeInUsd15m":  "bviu15m",
		"sellVolumeInUsd15m": "sviu15m",
		"price15m":           "p15m",
		"openInUsd15m":       "oiu15m",
		"closeInUsd15m":      "ciu15m",
		"buys30m":            "b30m",
		"sells30m":           "s30m",
		"buyers30m":          "be30m",
		"sellers30m":         "se30m",
		"buyVolumeInUsd30m":  "bviu30m",
		"sellVolumeInUsd30m": "sviu30m",
		"price30m":           "p30m",
		"openInUsd30m":       "oiu30m",
		"closeInUsd30m":      "ciu30m",
		"buys1h":             "b1h",
		"sells1h":            "s1h",
		"buyers1h":           "be1h",
		"sellers1h":          "se1h",
		"buyVolumeInUsd1h":   "bviu1h",
		"sellVolumeInUsd1h":  "sviu1h",
		"price1h":            "p1h",
		"openInUsd1h":        "oiu1h",
		"closeInUsd1h":       "ciu1h",
		"buys4h":             "b4h",
		"sells4h":            "s4h",
		"buyers4h":           "be4h",
		"sellers4h":          "se4h",
		"buyVolumeInUsd4h":   "bviu4h",
		"sellVolumeInUsd4h":  "sviu4h",
		"price4h":            "p4h",
		"openInUsd4h":        "oiu4h",
		"closeInUsd4h":       "ciu4h",
		"buys24h":            "b24h",
		"sells24h":           "s24h",
		"buyers24h":          "be24h",
		"sellers24h":         "se24h",
		"buyVolumeInUsd24h":  "bviu24h",
		"sellVolumeInUsd24h": "sviu24h",
		"price24h":           "p24h",
		"price":              "p",
		"openInUsd24h":       "oiu24h",
		"closeInUsd24h":      "ciu24h",
	},

	// Token holder subscription fields
	"subscribeTokenHolders": {
		"tokenAddress":  "a",
		"holders":       "h",
		"top100Amount":  "t100a",
		"top10Amount":   "t10a",
		"top100Holders": "t100h",
		"top10Holders":  "t10h",
		"top100Ratio":   "t100r",
		"top10Ratio":    "t10r",
		"timestamp":     "ts",
	},

	// New token subscription fields
	"subscribeNewToken": {
		"tokenAddress": "a",
		"name":         "n",
		"symbol":       "s",
		"createdAtMs":  "cts",
	},

	// Token supply subscription fields
	"subscribeTokenSupply": {
		"tokenAddress": "a",
		"supply":       "s",
		"timestamp":    "ts",
	},

	// DEX pool balance subscription fields
	"subscribeDexPoolBalance": {
		"poolAddress":          "a",
		"tokenAAddress":        "taa",
		"tokenALiquidityInUsd": "taliu",
		"tokenBAddress":        "tba",
		"tokenBLiquidityInUsd": "tbliu",
	},

	// Token liquidity subscription fields
	"subscribeTokenLiquidity": {
		"tokenAddress": "a",
		"metricType":   "t",
		"value":        "v",
		"timestamp":    "ts",
	},

	// New token metadata subscription fields
	"subscribeNewTokensMetadata": {
		"tokenAddress": "a",
		"name":         "n",
		"symbol":       "s",
		"imageUrl":     "iu",
		"description":  "de",
		"socialMedia":  "sm",
		"createdAtMs":  "cts",
	},

	// Token trade subscription fields
	"subscribeTokenTrades": {
		"tokenAddress":      "a",
		"timestamp":         "t",
		"kind":              "k",
		"buyAmount":         "ba",
		"buyAmountInUsd":    "baiu",
		"buyTokenAddress":   "btma",
		"buyTokenName":      "btn",
		"buyTokenSymbol":    "bts",
		"buyWalletAddress":  "bwa",
		"sellAmount":        "sa",
		"sellAmountInUsd":   "saiu",
		"sellTokenAddress":  "stma",
		"sellTokenName":     "stn",
		"sellTokenSymbol":   "sts",
		"sellWalletAddress": "swa",
		"txHash":            "h",
	},

	// Wallet token PnL subscription fields
	"subscribeWalletPnl": {
		"walletAddress":            "a",
		"tokenAddress":             "ta",
		"tokenPriceInUsd":          "tpiu",
		"timestamp":                "t",
		"opentime":                 "ot",
		"lasttime":                 "lt",
		"closetime":                "ct",
		"buyAmount":                "ba",
		"buyAmountInUsd":           "baiu",
		"buyCount":                 "bs",
		"buyCount30d":              "bs30d",
		"buyCount7d":               "bs7d",
		"sellAmount":               "sa",
		"sellAmountInUsd":          "saiu",
		"sellCount":                "ss",
		"sellCount30d":             "ss30d",
		"sellCount7d":              "ss7d",
		"heldDurationTimestamp":    "hdts",
		"averageBuyPriceInUsd":     "abpiu",
		"averageSellPriceInUsd":    "aspiu",
		"unrealizedProfitInUsd":    "upiu",
		"unrealizedProfitRatio":    "upr",
		"realizedProfitInUsd":      "rpiu",
		"realizedProfitRatio":      "rpr",
		"totalRealizedProfitInUsd": "trpiu",
		"totalRealizedProfitRatio": "trr",
	},

	// Wallet trade subscription fields
	"subscribeWalletTrade": {
		"tokenAddress":      "a",
		"timestamp":         "t",
		"kind":              "k",
		"buyAmount":         "ba",
		"buyAmountInUsd":    "baiu",
		"buyTokenAddress":   "btma",
		"buyTokenName":      "btn",
		"buyTokenSymbol":    "bts",
		"buyWalletAddress":  "bwa",
		"sellAmount":        "sa",
		"sellAmountInUsd":   "saiu",
		"sellTokenAddress":  "stma",
		"sellTokenName":     "stn",
		"sellTokenSymbol":   "sts",
		"sellWalletAddress": "swa",
		"txHash":            "h",
	},

	// Token max liquidity subscription fields
	"subscribeTokenMaxLiquidity": {
		"tokenAddress":      "a",
		"poolAddress":       "p",
		"liquidityInUsd":    "liu",
		"liquidityInNative": "lin",
		"timestamp":         "ts",
	},

	// Token total liquidity subscription fields
	"subscribeTokenTotalLiquidity": {
		"tokenAddress":      "a",
		"liquidityInUsd":    "liu",
		"liquidityInNative": "lin",
		"poolCount":         "pc",
		"timestamp":         "ts",
	},
}

// GetFieldMappings gets field mappings for a specific subscription method
func GetFieldMappings(methodName string) FieldMapping {
	if mappings, exists := CEL_FIELD_MAPPINGS[methodName]; exists {
		return mappings
	}
	return FieldMapping{}
}

// ReplaceFilterFields replaces long field names with short field names in filter expressions
// Automatically adds meta. prefix if not present
func ReplaceFilterFields(filter, methodName string) string {
	if filter == "" {
		return filter
	}

	fieldMappings := GetFieldMappings(methodName)
	result := filter

	// Replace all long field names with short field names
	for longField, shortField := range fieldMappings {
		// Handle two cases: with and without meta. prefix
		patterns := []*regexp.Regexp{
			// Pattern 1: fieldName (without meta. prefix)
			regexp.MustCompile(`\b` + regexp.QuoteMeta(longField) + `\b`),
			// Pattern 2: meta.fieldName (with meta. prefix)
			regexp.MustCompile(`\bmeta\.` + regexp.QuoteMeta(longField) + `\b`),
		}

		for _, pattern := range patterns {
			result = pattern.ReplaceAllString(result, "meta."+shortField)
		}
	}

	return result
}

// GetAvailableFields gets available field names for a specific subscription method
func GetAvailableFields(methodName string) []string {
	mappings := GetFieldMappings(methodName)
	fields := make([]string, 0, len(mappings))
	for field := range mappings {
		fields = append(fields, field)
	}
	return fields
}
