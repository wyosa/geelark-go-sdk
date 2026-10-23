package automation

import (
	"context"
)

type FileUploadRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Files, up to 100, refer to the User Guide - File Upload for creating automation tasks.
	// Required.
	Files []string `json:"files"`
}

type FileUploadResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// FileUpload: https://open.geelark.com/api/rpa-task-fileUpload
func (c *Client) FileUpload(ctx context.Context, req FileUploadRequest) (*FileUploadResponse, error) {
	var resp FileUploadResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/fileUpload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportContactsRequestContactsItem struct {
	Email1 string `json:"email1,omitempty"`
	Email2 string `json:"email2,omitempty"`
	// Fax number. At least one of the mobile phone number, work mobile phone number and fax
	// number must be non-empty.
	Fax string `json:"fax,omitempty"`
	// First name. At least one of the first name and last name must be non-empty.
	FirstName string `json:"firstName,omitempty"`
	// Last name. At least one of the first name and last name must be non-empty.
	LastName string `json:"lastName,omitempty"`
	// Mobile phone number. At least one of the mobile phone number, work mobile phone number
	// and fax number must be non-empty.
	Mobile string `json:"mobile,omitempty"`
	// Work mobile phone number. At least one of the mobile phone number, work mobile phone
	// number and fax number must be non-empty.
	Work string `json:"work,omitempty"`
}

type ImportContactsRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Array of contact information. Required.
	Contacts []ImportContactsRequestContactsItem `json:"contacts"`
}

type ImportContactsResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// ImportContacts: https://open.geelark.com/api/rpa-task-importContacts
func (c *Client) ImportContacts(ctx context.Context, req ImportContactsRequest) (*ImportContactsResponse, error) {
	var resp ImportContactsResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/importContacts", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type KeyboxUploadRequest struct {
	// Task name, up to 128 characters.
	Name string `json:"name,omitempty"`
	// Remarks, up to 200 characters.
	Remark string `json:"remark,omitempty"`
	// Scheduled execution time as a Unix timestamp in seconds. Required.
	ScheduleAt int64 `json:"scheduleAt"`
	// Cloud phone ID. Required.
	ID string `json:"id"`
	// Files, up to 1, refer to the User Guide - File Upload for creating automation tasks.
	// Required.
	Files []string `json:"files"`
}

type KeyboxUploadResponse struct {
	// ID assigned to the asynchronous task.
	TaskID string `json:"taskId"`
}

// KeyboxUpload: https://open.geelark.com/api/rpa-task-keyboxUpload
func (c *Client) KeyboxUpload(ctx context.Context, req KeyboxUploadRequest) (*KeyboxUploadResponse, error) {
	var resp KeyboxUploadResponse
	if err := c.core.Do(ctx, "/open/v1/rpa/task/keyboxUpload", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
