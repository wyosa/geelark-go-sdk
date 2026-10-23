// Package billing provides account balances, subscriptions, transactions, and time add-ons.
package billing

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/common"
)

type Client struct {
	core common.Transport
}

func New(core common.Transport) *Client { return &Client{core: core} }

// BalanceRequest contains no parameters for [Client.Balance].
type BalanceRequest struct{}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
	// Gifted amount.
	GiftMoney float64 `json:"giftMoney"`
	// Remaining time add-on.
	AvailableTimeAddOn int `json:"availableTimeAddOn"`
}

// Balance: https://open.geelark.com/api/pay-wallet
func (c *Client) Balance(ctx context.Context, req BalanceRequest) (*BalanceResponse, error) {
	var resp BalanceResponse
	if err := c.core.Do(ctx, "/open/v1/pay/wallet", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PlansRequest contains no parameters for [Client.Plans].
type PlansRequest struct{}

type PlansResponseItem struct {
	// Plan ID.
	ID string `json:"id"`
	// Monthly plan price.
	Price float64 `json:"price"`
	// Plan tier: 0 = Base; 1 = Pro.
	Level int `json:"level"`
	// Number of profiles included.
	EnvNum int `json:"envNum"`
	// Bonus minutes included.
	FreeTime int `json:"freeTime"`
	// Maximum daily profile launches.
	OpenEnvNumOneDay int `json:"openEnvNumOneDay"`
	// Maximum daily profile creates.
	CreateEnvNumOneDay int `json:"createEnvNumOneDay"`
}

type PlansResponse []PlansResponseItem

// Plans: https://open.geelark.com/api/pay-profiles-list
func (c *Client) Plans(ctx context.Context, req PlansRequest) (*PlansResponse, error) {
	var resp PlansResponse
	if err := c.core.Do(ctx, "/open/v1/pay/profiles/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ChangePlanRequest struct {
	// Plan ID. Obtain it from Get Plan List. Required.
	ProfilesID string `json:"profilesId"`
	// Number of parallel sessions; must be at least the current plan’s value. Required.
	ParallelsNum int `json:"parallelsNum"`
	// Number of monthly rental slots; must be at least the current plan’s value. Required.
	MonthlyRentalNum int `json:"monthlyRentalNum"`
	// Required if the plan has expired. Allowed values: 30, 90, 180, or 360.
	Days      *int   `json:"days,omitempty"`
	PromoCode string `json:"promoCode,omitempty"`
}

type ChangePlanResponse struct{}

// ChangePlan: https://open.geelark.com/api/pay-plan-upgrade
//
// This endpoint only upgrades the subscription. Downgrades must be performed in the
// GeeLark desktop client.
func (c *Client) ChangePlan(ctx context.Context, req ChangePlanRequest) (*ChangePlanResponse, error) {
	var resp ChangePlanResponse
	if err := c.core.Do(ctx, "/open/v1/pay/plan/upgrade", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RenewPlanRequest struct {
	// renewal duration：30/90/180/360 day. Required.
	Days      int    `json:"days"`
	PromoCode string `json:"promoCode,omitempty"`
}

type RenewPlanResponse struct{}

// RenewPlan: https://open.geelark.com/api/pay-plan-continue
func (c *Client) RenewPlan(ctx context.Context, req RenewPlanRequest) (*RenewPlanResponse, error) {
	var resp RenewPlanResponse
	if err := c.core.Do(ctx, "/open/v1/pay/plan/continue", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PlanInfoRequest contains no parameters for [Client.PlanInfo].
type PlanInfoRequest struct{}

type PlanInfoResponse struct {
	// Subscription type: 0 = Base, 1 = Pro.
	Plan int `json:"plan"`
	// Total number of profiles.
	Profiles int `json:"profiles"`
	// Number of monthly rental devices.
	MonthlyRental int `json:"monthlyRental"`
	// No extra charge as long as the number of cloud phone profiles opened at the same time
	// does not exceed your parallel limit.
	Parallels int `json:"parallels"`
	// Plan expiration timestamp.
	ExpirationTime int64 `json:"expirationTime"`
	// Monthly billing amount.
	MonthlyFee float64 `json:"monthlyFee"`
	// Number of available profiles.
	AvailableProfiles int `json:"availableProfiles"`
	// Number of available monthly rental devices.
	AvailableMonthlyRentals int `json:"availableMonthlyRentals"`
}

// PlanInfo: https://open.geelark.com/api/plan-info
func (c *Client) PlanInfo(ctx context.Context, req PlanInfoRequest) (*PlanInfoResponse, error) {
	var resp PlanInfoResponse
	if err := c.core.Do(ctx, "/open/v1/pay/plan/info", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TransactionsRequest struct {
	// If omitted, transactions for all cloud phones are returned.
	ID string `json:"id,omitempty"`
	// Filter start time, second-level timestamp (currently only supports searching data within
	// the last 3 days).
	StartAt *int64 `json:"startAt,omitempty"`
	// Filtering end time, second-level timestamp (currently only supports searching data
	// within the last 3 days).
	EndAt *int64 `json:"endAt,omitempty"`
	// The acquisition quantity limit is set to 100 by default, with a maximum of 1000.
	Limit *int `json:"limit,omitempty"`
	// The lastFlowId returned from the previous request is used to obtain the data for the
	// next page.
	LastFlowID string `json:"lastFlowId,omitempty"`
}

type TransactionsResponseListItem struct {
	// flow id.
	ID string `json:"id"`
	// cloud phone id.
	EnvID  string  `json:"envId"`
	Amount float64 `json:"amount"`
	// Usage duration, minutes.
	UsedTime int `json:"usedTime"`
	// Usage type ，1-cloud phone 2-RPA.
	Type int `json:"type"`
	// Billing type 1-Points 2-Balance 3-Bonus 4-Time add-on 5-Bonus minutes 6-Monthly rental
	// 7-Parallels 8-Daily cap reached.
	ChargeType int `json:"chargeType"`
	// flow created time, second-level timestamp.
	CreatedTime int64 `json:"createdTime"`
}

type TransactionsResponse struct {
	// Total number of matching records.
	Total int `json:"total"`
	// Current page number.
	Page int `json:"page"`
	// Number of records per page.
	PageSize int `json:"pageSize"`
	// Billing transactions matching the requested filters.
	List []TransactionsResponseListItem `json:"list"`
	// The last sequential ID, used to retrieve data from the next page.
	LastFlowID string `json:"lastFlowId"`
}

// Transactions: https://open.geelark.com/api/billing-transaction-detail
func (c *Client) Transactions(ctx context.Context, req TransactionsRequest) (*TransactionsResponse, error) {
	var resp TransactionsResponse
	if err := c.core.Do(ctx, "/open/v1/billing/transaction/detail", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BuyTimeRequest struct {
	// Number of minutes to purchase. Required.
	Minutes   int    `json:"minutes"`
	PromoCode string `json:"promoCode,omitempty"`
}

type BuyTimeResponse struct{}

// BuyTime: https://open.geelark.com/api/time-add-on-buy
//
// Purchases a time add-on using the account balance. Ensure sufficient funds before
// submitting the request.
func (c *Client) BuyTime(ctx context.Context, req BuyTimeRequest) (*BuyTimeResponse, error) {
	var resp BuyTimeResponse
	if err := c.core.Do(ctx, "/open/v1/pay/timeAddOn/buy", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}
