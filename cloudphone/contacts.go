package cloudphone

import (
	"context"
)

type SendSMSRequest struct {
	// Cloud phone profile ID (Currently supports only Android 12 , 13, 14, 15, 16 devices).
	// Required.
	ID string `json:"id"`
	// Phone number. Required.
	PhoneNumber string `json:"phoneNumber"`
	// SMS content. Required.
	Text string `json:"text"`
}

type SendSMSResponse struct{}

// SendSMS: https://open.geelark.com/api/phone-sendSms
//
// Start the cloud phone before sending an SMS.
func (c *Client) SendSMS(ctx context.Context, req SendSMSRequest) (*SendSMSResponse, error) {
	var resp SendSMSResponse
	if err := c.core.Do(ctx, "/open/v1/phone/sendSms", req, nil); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportContactsRequestContactsItem struct {
	Email1 string `json:"email1,omitempty"`
	Email2 string `json:"email2,omitempty"`
	// Fax number. At least one of mobile, work, or fax must be non-empty.
	Fax string `json:"fax,omitempty"`
	// First name. At least one of firstName or lastName must be non-empty.
	FirstName string `json:"firstName,omitempty"`
	// Last name. At least one of firstName or lastName must be non-empty.
	LastName string `json:"lastName,omitempty"`
	// Mobile phone number. At least one of mobile, work, or fax must be non-empty.
	Mobile string `json:"mobile,omitempty"`
	// Work phone number. At least one of mobile, work, or fax must be non-empty.
	Work string `json:"work,omitempty"`
}

type ImportContactsRequest struct {
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Array of contact information. Required.
	Contacts []ImportContactsRequestContactsItem `json:"contacts"`
}

type ImportContactsResponse struct {
	// Task ID (valid for query one hour after the task is created).
	TaskID string `json:"taskId"`
}

// ImportContacts: https://open.geelark.com/api/phone-importContacts
func (c *Client) ImportContacts(ctx context.Context, req ImportContactsRequest) (*ImportContactsResponse, error) {
	var resp ImportContactsResponse
	if err := c.core.Do(ctx, "/open/v1/phone/importContacts", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportContactsResultRequest struct {
	// Task ID. Required.
	TaskID string `json:"taskId"`
}

type ImportContactsResultResponse struct {
	// Task status: 1 = In progress, 2 = Successful, 3 = Failed.
	Status int `json:"status"`
}

// ImportContactsResult: https://open.geelark.com/api/phone-importContactsResult
func (c *Client) ImportContactsResult(ctx context.Context, req ImportContactsResultRequest) (*ImportContactsResultResponse, error) {
	var resp ImportContactsResultResponse
	if err := c.core.Do(ctx, "/open/v1/phone/importContactsResult", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
