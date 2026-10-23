package billing

import (
	"context"

	"github.com/wyosa/geelark-go-sdk/internal/clientcore"
)

type Client struct {
	core *clientcore.Client
}

func New(core *clientcore.Client) *Client { return &Client{core: core} }

type BalanceRequest struct{}

type BalanceResponse struct {
	Balance            float64 `json:"balance"`
	GiftMoney          float64 `json:"giftMoney"`
	AvailableTimeAddOn int     `json:"availableTimeAddOn"`
}

// Balance calls /open/v1/pay/wallet.
func (c *Client) Balance(ctx context.Context, req BalanceRequest) (*BalanceResponse, error) {
	var resp BalanceResponse
	if err := c.core.Do(ctx, "/open/v1/pay/wallet", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PlansRequest struct{}

type PlansResponseItem struct {
	ID                 string  `json:"id"`
	Price              float64 `json:"price"`
	Level              int     `json:"level"`
	EnvNum             int     `json:"envNum"`
	FreeTime           int     `json:"freeTime"`
	OpenEnvNumOneDay   int     `json:"openEnvNumOneDay"`
	CreateEnvNumOneDay int     `json:"createEnvNumOneDay"`
}

type PlansResponse []PlansResponseItem

// Plans calls /open/v1/pay/profiles/list.
func (c *Client) Plans(ctx context.Context, req PlansRequest) (*PlansResponse, error) {
	var resp PlansResponse
	if err := c.core.Do(ctx, "/open/v1/pay/profiles/list", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ChangePlanRequest struct {
	ProfilesID       string `json:"profilesId"`
	ParallelsNum     int    `json:"parallelsNum"`
	MonthlyRentalNum int    `json:"monthlyRentalNum"`
	Days             *int   `json:"days,omitempty"`
	PromoCode        string `json:"promoCode,omitempty"`
}

type ChangePlanResponse struct{}

// ChangePlan calls /open/v1/pay/plan/upgrade.
func (c *Client) ChangePlan(ctx context.Context, req ChangePlanRequest) (*ChangePlanResponse, error) {
	var resp ChangePlanResponse
	if err := c.core.Do(ctx, "/open/v1/pay/plan/upgrade", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RenewPlanRequest struct {
	Days      int    `json:"days"`
	PromoCode string `json:"promoCode,omitempty"`
}

type RenewPlanResponse struct{}

// RenewPlan calls /open/v1/pay/plan/continue.
func (c *Client) RenewPlan(ctx context.Context, req RenewPlanRequest) (*RenewPlanResponse, error) {
	var resp RenewPlanResponse
	if err := c.core.Do(ctx, "/open/v1/pay/plan/continue", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PlanInfoRequest struct{}

type PlanInfoResponse struct {
	Plan                    int     `json:"plan"`
	Profiles                int     `json:"profiles"`
	MonthlyRental           int     `json:"monthlyRental"`
	Parallels               int     `json:"parallels"`
	ExpirationTime          int     `json:"expirationTime"`
	MonthlyFee              float64 `json:"monthlyFee"`
	AvailableProfiles       int     `json:"availableProfiles"`
	AvailableMonthlyRentals int     `json:"availableMonthlyRentals"`
}

// PlanInfo calls /open/v1/pay/plan/info.
func (c *Client) PlanInfo(ctx context.Context, req PlanInfoRequest) (*PlanInfoResponse, error) {
	var resp PlanInfoResponse
	if err := c.core.Do(ctx, "/open/v1/pay/plan/info", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TransactionsRequest struct {
	ID         string `json:"id,omitempty"`
	StartAt    *int   `json:"startAt,omitempty"`
	EndAt      *int   `json:"endAt,omitempty"`
	Limit      *int   `json:"limit,omitempty"`
	LastFlowID string `json:"lastFlowId,omitempty"`
}

type TransactionsResponseListItem struct {
	ID          string  `json:"id"`
	EnvID       string  `json:"envId"`
	Amount      float64 `json:"amount"`
	UsedTime    int     `json:"usedTime"`
	Type        int     `json:"type"`
	ChargeType  int     `json:"chargeType"`
	CreatedTime int     `json:"createdTime"`
}

type TransactionsResponse struct {
	Total      int                            `json:"total"`
	Page       int                            `json:"page"`
	PageSize   int                            `json:"pageSize"`
	List       []TransactionsResponseListItem `json:"list"`
	LastFlowID string                         `json:"lastFlowId"`
}

// Transactions calls /open/v1/billing/transaction/detail.
func (c *Client) Transactions(ctx context.Context, req TransactionsRequest) (*TransactionsResponse, error) {
	var resp TransactionsResponse
	if err := c.core.Do(ctx, "/open/v1/billing/transaction/detail", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BuyTimeRequest struct {
	Minutes   int    `json:"minutes"`
	PromoCode string `json:"promoCode,omitempty"`
}

type BuyTimeResponse struct{}

// BuyTime calls /open/v1/pay/timeAddOn/buy.
func (c *Client) BuyTime(ctx context.Context, req BuyTimeRequest) (*BuyTimeResponse, error) {
	var resp BuyTimeResponse
	if err := c.core.Do(ctx, "/open/v1/pay/timeAddOn/buy", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
