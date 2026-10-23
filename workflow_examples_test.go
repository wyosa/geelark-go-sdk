package geelark_test

import (
	"context"
	"fmt"
	"net/http"
	"time"

	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/shell"
)

func ExampleClient_cloudPhoneReady() {
	// This transport keeps the example offline. Omit WithHTTPClient in production.
	hc := &http.Client{Transport: exampleTransport(func(r *http.Request) (*http.Response, error) {
		switch r.URL.Path {
		case "/open/v1/phone/start":
			return exampleReply(`{"code":0,"data":{"totalAmount":1,"successAmount":1,"failAmount":0}}`), nil
		case "/open/v1/phone/status":
			return exampleReply(`{"code":0,"data":{"successDetails":[{"id":"phone-id","status":0}]}}`), nil
		case "/open/v1/shell/execute":
			return exampleReply(`{"code":0,"data":{"status":true,"output":"ready"}}`), nil
		}
		return nil, fmt.Errorf("unexpected path %s", r.URL.Path)
	})}
	client, err := geelark.NewClient("example-token", geelark.WithHTTPClient(hc))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	started, err := client.CloudPhones.Start(ctx, cloudphone.StartRequest{IDs: []string{"phone-id"}})
	if err != nil {
		panic(err) // Inspect APIError.Data for partial success, as in ExampleAPIError.
	}
	if started.FailAmount != 0 {
		panic(fmt.Sprintf("phone startup failed: %+v", started.FailDetails))
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		status, err := client.CloudPhones.Status(ctx, cloudphone.StatusRequest{IDs: []string{"phone-id"}})
		if err != nil {
			panic(err)
		}
		if status.FailAmount != 0 {
			panic(fmt.Sprintf("phone unavailable: %+v", status.FailDetails))
		}
		if len(status.SuccessDetails) == 1 && status.SuccessDetails[0].Status == cloudphone.StatusStarted {
			break
		}
		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-ticker.C:
		}
	}
	result, err := client.Shell.Execute(ctx, shell.ExecuteRequest{ID: "phone-id", Cmd: "echo ready"})
	if err != nil {
		panic(err)
	}
	if !result.Status {
		panic(result.Output)
	}
	fmt.Println(result.Output)
	// Output: ready
}

func ExampleClient_taskCompletion() {
	// The configured flow ID and parameter names come from Automation.ListFlows.
	hc := &http.Client{Transport: exampleTransport(func(r *http.Request) (*http.Response, error) {
		if r.URL.Path == "/open/v1/task/rpa/add" {
			return exampleReply(`{"code":0,"data":{"taskId":"task-id"}}`), nil
		}
		return exampleReply(`{"code":0,"data":{"items":[{"id":"task-id","status":3,"shareLink":"https://example.com/post"}]}}`), nil
	})}
	client, err := geelark.NewClient("example-token", geelark.WithHTTPClient(hc))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	created, err := client.Automation.CreateCustomTask(ctx, automation.CreateCustomTaskRequest{
		ID: "phone-id", FlowID: "configured-flow-id", ScheduleAt: time.Now().Unix(),
		ParamMap: automation.CreateCustomTaskRequestParamMap{"Title": "Example"},
	})
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		result, err := client.Automation.QueryTask(ctx, automation.QueryTaskRequest{IDs: []string{created.TaskID}})
		if err != nil {
			panic(err)
		}
		for _, task := range result.Items {
			if task.ID != created.TaskID {
				continue
			}
			switch task.Status {
			case automation.TaskStatusCompleted:
				fmt.Println(task.ID, task.ShareLink)
				return
			case automation.TaskStatusFailed, automation.TaskStatusCancelled:
				panic(fmt.Sprintf("task ended: status=%d code=%d reason=%s", task.Status, task.FailCode, task.FailDesc))
			}
		}
		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-ticker.C:
		}
	}
	// Output: task-id https://example.com/post
}

func ExampleClient_pagination() {
	page := 0
	hc := &http.Client{Transport: exampleTransport(func(*http.Request) (*http.Response, error) {
		page++
		return exampleReply(fmt.Sprintf(`{"code":0,"data":{"total":2,"items":[{"id":"phone-%d"}]}}`, page)), nil
	})}
	client, err := geelark.NewClient("example-token", geelark.WithHTTPClient(hc))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	const pageSize = 1 // Choose up to the endpoint's documented page-size limit.
	for page := 1; ; page++ {
		result, err := client.CloudPhones.List(ctx, cloudphone.ListRequest{Page: page, PageSize: pageSize})
		if err != nil {
			panic(err)
		}
		for _, phone := range result.Items {
			fmt.Println(phone.ID)
		}
		if len(result.Items) == 0 || page*pageSize >= result.Total {
			break
		}
	}
	// Output:
	// phone-1
	// phone-2
}
