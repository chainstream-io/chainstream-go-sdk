// This file adds the DAY5 holders / holder-stats / traders endpoints. It is kept
// separate from client.gen.go but follows the same hand-cleaned conventions
// (shared request builder + typed query-param helpers, no external runtime
// dependency) so the prediction client stays consistent and easy to maintain.

package prediction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Defines values for PredictionHolderOrder.
const (
	PredictionHolderOrderAsc  PredictionHolderOrder = "asc"
	PredictionHolderOrderDesc PredictionHolderOrder = "desc"
)

// PredictionHolderOrder defines model for PredictionHolderOrder.
type PredictionHolderOrder string

// Valid indicates whether the value is a known member of the PredictionHolderOrder enum.
func (e PredictionHolderOrder) Valid() bool {
	switch e {
	case PredictionHolderOrderAsc, PredictionHolderOrderDesc:
		return true
	default:
		return false
	}
}

// Defines values for PredictionHolderSortBy.
const (
	PredictionHolderSortByBalance      PredictionHolderSortBy = "balance"
	PredictionHolderSortByHoldingValue PredictionHolderSortBy = "holding_value"
	PredictionHolderSortByTotalPnl     PredictionHolderSortBy = "total_pnl"
)

// PredictionHolderSortBy defines model for PredictionHolderSortBy.
type PredictionHolderSortBy string

// Valid indicates whether the value is a known member of the PredictionHolderSortBy enum.
func (e PredictionHolderSortBy) Valid() bool {
	switch e {
	case PredictionHolderSortByBalance, PredictionHolderSortByHoldingValue, PredictionHolderSortByTotalPnl:
		return true
	default:
		return false
	}
}

// Defines values for PredictionTraderSortBy.
const (
	PredictionTraderSortByRealizedPnl PredictionTraderSortBy = "realized_pnl"
	PredictionTraderSortByTotalPnl    PredictionTraderSortBy = "total_pnl"
	PredictionTraderSortByTotalVolume PredictionTraderSortBy = "total_volume"
)

// PredictionTraderSortBy defines model for PredictionTraderSortBy.
type PredictionTraderSortBy string

// Valid indicates whether the value is a known member of the PredictionTraderSortBy enum.
func (e PredictionTraderSortBy) Valid() bool {
	switch e {
	case PredictionTraderSortByRealizedPnl, PredictionTraderSortByTotalPnl, PredictionTraderSortByTotalVolume:
		return true
	default:
		return false
	}
}

// PredictionHolder defines model for PredictionHolder.
type PredictionHolder struct {
	AvgPrice       *string   `json:"avgPrice,omitempty"`
	Balance        *string   `json:"balance,omitempty"`
	HoldingValue   *string   `json:"holdingValue,omitempty"`
	IsSmart        *bool     `json:"isSmart,omitempty"`
	IsWhale        *bool     `json:"isWhale,omitempty"`
	LastActivityTs *string   `json:"lastActivityTs,omitempty"`
	MarkPrice      *string   `json:"markPrice,omitempty"`
	RealizedPnl    *string   `json:"realizedPnl,omitempty"`
	Topics         *[]string `json:"topics,omitempty"`
	TotalPnl       *string   `json:"totalPnl,omitempty"`
	TotalPnlRatio  *string   `json:"totalPnlRatio,omitempty"`
	TotalVolume    *string   `json:"totalVolume,omitempty"`
	UnrealizedPnl  *string   `json:"unrealizedPnl,omitempty"`
	Wallet         *string   `json:"wallet,omitempty"`
}

// PredictionHolderCohort defines model for PredictionHolderCohort.
type PredictionHolderCohort struct {
	Count        *int64  `json:"count,omitempty"`
	HoldingValue *string `json:"holdingValue,omitempty"`
}

// PredictionHolderOutcomeStats defines model for PredictionHolderOutcomeStats.
type PredictionHolderOutcomeStats struct {
	AvgProfit          *string                 `json:"avgProfit,omitempty"`
	Dolphin            *PredictionHolderCohort `json:"dolphin,omitempty"`
	HolderCount        *int64                  `json:"holderCount,omitempty"`
	HoldingValue       *string                 `json:"holdingValue,omitempty"`
	MedianPnl          *string                 `json:"medianPnl,omitempty"`
	MedianYield        *float64                `json:"medianYield,omitempty"`
	Outcome            *string                 `json:"outcome,omitempty"`
	Ratio              *float64                `json:"ratio,omitempty"`
	Shark              *PredictionHolderCohort `json:"shark,omitempty"`
	SmartMoney         *PredictionHolderCohort `json:"smartMoney,omitempty"`
	TokenId            *string                 `json:"tokenId,omitempty"`
	Top100             *PredictionHolderCohort `json:"top100,omitempty"`
	TopHolderShare     *float64                `json:"topHolderShare,omitempty"`
	TotalBalance       *string                 `json:"totalBalance,omitempty"`
	TotalCostBasis     *string                 `json:"totalCostBasis,omitempty"`
	TotalPnl           *string                 `json:"totalPnl,omitempty"`
	TotalProfit        *string                 `json:"totalProfit,omitempty"`
	TotalRealizedPnl   *string                 `json:"totalRealizedPnl,omitempty"`
	TotalUnrealizedPnl *string                 `json:"totalUnrealizedPnl,omitempty"`
	TotalYield         *float64                `json:"totalYield,omitempty"`
	WeightTotalProfit  *string                 `json:"weightTotalProfit,omitempty"`
	Whale              *PredictionHolderCohort `json:"whale,omitempty"`
}

// PredictionMarketHolderStatsResponse defines model for PredictionMarketHolderStatsResponse.
type PredictionMarketHolderStatsResponse struct {
	ConditionId  string                         `json:"conditionId"`
	Outcomes     []PredictionHolderOutcomeStats `json:"outcomes"`
	StateQuality string                         `json:"stateQuality"`
	UpdatedAtMs  *int64                         `json:"updatedAtMs,omitempty"`
}

// PredictionMarketHoldersResponse defines model for PredictionMarketHoldersResponse.
type PredictionMarketHoldersResponse struct {
	ConditionId  string                 `json:"conditionId"`
	Cursor       *string                `json:"cursor,omitempty"`
	Holders      []PredictionHolder     `json:"holders"`
	Limit        int64                  `json:"limit"`
	Order        PredictionHolderOrder  `json:"order"`
	SortBy       PredictionHolderSortBy `json:"sortBy"`
	StateQuality string                 `json:"stateQuality"`
	TokenId      *string                `json:"tokenId,omitempty"`
	UpdatedAtMs  *int64                 `json:"updatedAtMs,omitempty"`
}

// PredictionMarketTradersResponse defines model for PredictionMarketTradersResponse.
type PredictionMarketTradersResponse struct {
	ConditionId  string                 `json:"conditionId"`
	Cursor       *string                `json:"cursor,omitempty"`
	Limit        int64                  `json:"limit"`
	Order        PredictionHolderOrder  `json:"order"`
	SortBy       PredictionTraderSortBy `json:"sortBy"`
	StateQuality string                 `json:"stateQuality"`
	Traders      []PredictionTrader     `json:"traders"`
	UpdatedAtMs  *int64                 `json:"updatedAtMs,omitempty"`
}

// PredictionTrader defines model for PredictionTrader.
type PredictionTrader struct {
	Amount            *string   `json:"amount,omitempty"`
	AvgPrice          *string   `json:"avgPrice,omitempty"`
	IsSmart           *bool     `json:"isSmart,omitempty"`
	LastActivityTs    *string   `json:"lastActivityTs,omitempty"`
	MarketTags        *[]string `json:"marketTags,omitempty"`
	Name              *string   `json:"name,omitempty"`
	NetBalance        *string   `json:"netBalance,omitempty"`
	NetTokenId        *string   `json:"netTokenId,omitempty"`
	ProfileImage      *string   `json:"profileImage,omitempty"`
	RealizedProfit    *string   `json:"realizedProfit,omitempty"`
	RealizedProfitPnl *float64  `json:"realizedProfitPnl,omitempty"`
	TotalPnl          *string   `json:"totalPnl,omitempty"`
	TotalVolume       *string   `json:"totalVolume,omitempty"`
	Wallet            *string   `json:"wallet,omitempty"`
}

// GetPredictionMarketHoldersParams defines parameters for GetPredictionMarketHolders.
type GetPredictionMarketHoldersParams struct {
	Cursor  *string                 `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit   *int64                  `form:"limit,omitempty" json:"limit,omitempty"`
	TokenId *string                 `form:"token_id,omitempty" json:"token_id,omitempty"`
	Tag     *string                 `form:"tag,omitempty" json:"tag,omitempty"`
	SortBy  *PredictionHolderSortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`
	Order   *PredictionHolderOrder  `form:"order,omitempty" json:"order,omitempty"`
}

// GetPredictionMarketTradersParams defines parameters for GetPredictionMarketTraders.
type GetPredictionMarketTradersParams struct {
	Cursor *string                 `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  *int64                  `form:"limit,omitempty" json:"limit,omitempty"`
	Tag    *string                 `form:"tag,omitempty" json:"tag,omitempty"`
	SortBy *PredictionTraderSortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`
	Order  *PredictionHolderOrder  `form:"order,omitempty" json:"order,omitempty"`
}

func addPredictionHolderSortBy(values url.Values, name string, value *PredictionHolderSortBy) {
	if value != nil {
		values.Add(name, string(*value))
	}
}

func addPredictionHolderOrder(values url.Values, name string, value *PredictionHolderOrder) {
	if value != nil {
		values.Add(name, string(*value))
	}
}

func addPredictionTraderSortBy(values url.Values, name string, value *PredictionTraderSortBy) {
	if value != nil {
		values.Add(name, string(*value))
	}
}

// NewGetPredictionMarketHoldersRequest generates requests for GetPredictionMarketHolders.
func NewGetPredictionMarketHoldersRequest(server string, conditionId string, params *GetPredictionMarketHoldersParams) (*http.Request, error) {
	return newPredictionActivitiesRequest(server, fmt.Sprintf("/v1/prediction/markets/%s/holders", url.PathEscape(conditionId)), func(values url.Values) {
		if params == nil {
			return
		}
		addString(values, "cursor", params.Cursor)
		addInt64(values, "limit", params.Limit)
		addString(values, "token_id", params.TokenId)
		addString(values, "tag", params.Tag)
		addPredictionHolderSortBy(values, "sort_by", params.SortBy)
		addPredictionHolderOrder(values, "order", params.Order)
	})
}

// NewGetPredictionMarketHolderStatsRequest generates requests for GetPredictionMarketHolderStats.
func NewGetPredictionMarketHolderStatsRequest(server string, conditionId string) (*http.Request, error) {
	return newPredictionActivitiesRequest(server, fmt.Sprintf("/v1/prediction/markets/%s/holder-stats", url.PathEscape(conditionId)), func(url.Values) {})
}

// NewGetPredictionMarketTradersRequest generates requests for GetPredictionMarketTraders.
func NewGetPredictionMarketTradersRequest(server string, conditionId string, params *GetPredictionMarketTradersParams) (*http.Request, error) {
	return newPredictionActivitiesRequest(server, fmt.Sprintf("/v1/prediction/markets/%s/traders", url.PathEscape(conditionId)), func(values url.Values) {
		if params == nil {
			return
		}
		addString(values, "cursor", params.Cursor)
		addInt64(values, "limit", params.Limit)
		addString(values, "tag", params.Tag)
		addPredictionTraderSortBy(values, "sort_by", params.SortBy)
		addPredictionHolderOrder(values, "order", params.Order)
	})
}

// GetPredictionMarketHolders request.
func (c *Client) GetPredictionMarketHolders(ctx context.Context, conditionId string, params *GetPredictionMarketHoldersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPredictionMarketHoldersRequest(c.Server, conditionId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPredictionMarketHolderStats request.
func (c *Client) GetPredictionMarketHolderStats(ctx context.Context, conditionId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPredictionMarketHolderStatsRequest(c.Server, conditionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPredictionMarketTraders request.
func (c *Client) GetPredictionMarketTraders(ctx context.Context, conditionId string, params *GetPredictionMarketTradersParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPredictionMarketTradersRequest(c.Server, conditionId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPredictionMarketHoldersResponse is the response for GetPredictionMarketHolders.
type GetPredictionMarketHoldersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PredictionMarketHoldersResponse
}

// Status returns HTTPResponse.Status.
func (r GetPredictionMarketHoldersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode.
func (r GetPredictionMarketHoldersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BodyBytes returns response body bytes.
func (r GetPredictionMarketHoldersResponse) BodyBytes() []byte {
	return bytes.Clone(r.Body)
}

// GetPredictionMarketHolderStatsResponse is the response for GetPredictionMarketHolderStats.
type GetPredictionMarketHolderStatsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PredictionMarketHolderStatsResponse
}

// Status returns HTTPResponse.Status.
func (r GetPredictionMarketHolderStatsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode.
func (r GetPredictionMarketHolderStatsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BodyBytes returns response body bytes.
func (r GetPredictionMarketHolderStatsResponse) BodyBytes() []byte {
	return bytes.Clone(r.Body)
}

// GetPredictionMarketTradersResponse is the response for GetPredictionMarketTraders.
type GetPredictionMarketTradersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PredictionMarketTradersResponse
}

// Status returns HTTPResponse.Status.
func (r GetPredictionMarketTradersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode.
func (r GetPredictionMarketTradersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// BodyBytes returns response body bytes.
func (r GetPredictionMarketTradersResponse) BodyBytes() []byte {
	return bytes.Clone(r.Body)
}

// asClient unwraps the concrete *Client backing a ClientWithResponses so the
// DAY5 endpoints can reuse the request helpers defined on *Client.
func (c *ClientWithResponses) asClient() (*Client, error) {
	if client, ok := c.ClientInterface.(*Client); ok {
		return client, nil
	}
	return nil, fmt.Errorf("prediction: ClientWithResponses is not backed by the default *Client")
}

// GetPredictionMarketHoldersWithResponse request returning *GetPredictionMarketHoldersResponse.
func (c *ClientWithResponses) GetPredictionMarketHoldersWithResponse(ctx context.Context, conditionId string, params *GetPredictionMarketHoldersParams, reqEditors ...RequestEditorFn) (*GetPredictionMarketHoldersResponse, error) {
	client, err := c.asClient()
	if err != nil {
		return nil, err
	}
	rsp, err := client.GetPredictionMarketHolders(ctx, conditionId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPredictionMarketHoldersResponse(rsp)
}

// GetPredictionMarketHolderStatsWithResponse request returning *GetPredictionMarketHolderStatsResponse.
func (c *ClientWithResponses) GetPredictionMarketHolderStatsWithResponse(ctx context.Context, conditionId string, reqEditors ...RequestEditorFn) (*GetPredictionMarketHolderStatsResponse, error) {
	client, err := c.asClient()
	if err != nil {
		return nil, err
	}
	rsp, err := client.GetPredictionMarketHolderStats(ctx, conditionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPredictionMarketHolderStatsResponse(rsp)
}

// GetPredictionMarketTradersWithResponse request returning *GetPredictionMarketTradersResponse.
func (c *ClientWithResponses) GetPredictionMarketTradersWithResponse(ctx context.Context, conditionId string, params *GetPredictionMarketTradersParams, reqEditors ...RequestEditorFn) (*GetPredictionMarketTradersResponse, error) {
	client, err := c.asClient()
	if err != nil {
		return nil, err
	}
	rsp, err := client.GetPredictionMarketTraders(ctx, conditionId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPredictionMarketTradersResponse(rsp)
}

// ParseGetPredictionMarketHoldersResponse parses an HTTP response from a GetPredictionMarketHoldersWithResponse call.
func ParseGetPredictionMarketHoldersResponse(rsp *http.Response) (*GetPredictionMarketHoldersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPredictionMarketHoldersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	if strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == http.StatusOK {
		var dest PredictionMarketHoldersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}
	return response, nil
}

// ParseGetPredictionMarketHolderStatsResponse parses an HTTP response from a GetPredictionMarketHolderStatsWithResponse call.
func ParseGetPredictionMarketHolderStatsResponse(rsp *http.Response) (*GetPredictionMarketHolderStatsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPredictionMarketHolderStatsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	if strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == http.StatusOK {
		var dest PredictionMarketHolderStatsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}
	return response, nil
}

// ParseGetPredictionMarketTradersResponse parses an HTTP response from a GetPredictionMarketTradersWithResponse call.
func ParseGetPredictionMarketTradersResponse(rsp *http.Response) (*GetPredictionMarketTradersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}
	response := &GetPredictionMarketTradersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}
	if strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == http.StatusOK {
		var dest PredictionMarketTradersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}
	return response, nil
}
