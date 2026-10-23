package geelark_test

import (
	"context"
	"encoding/json"
	"errors"
	geelark "github.com/wyosa/geelark-go-sdk"
	"github.com/wyosa/geelark-go-sdk/adb"
	"github.com/wyosa/geelark-go-sdk/analytics"
	"github.com/wyosa/geelark-go-sdk/app"
	"github.com/wyosa/geelark-go-sdk/automation"
	"github.com/wyosa/geelark-go-sdk/billing"
	"github.com/wyosa/geelark-go-sdk/browser"
	"github.com/wyosa/geelark-go-sdk/browserautomation"
	"github.com/wyosa/geelark-go-sdk/cloudphone"
	"github.com/wyosa/geelark-go-sdk/files"
	"github.com/wyosa/geelark-go-sdk/group"
	"github.com/wyosa/geelark-go-sdk/library"
	"github.com/wyosa/geelark-go-sdk/livestream"
	"github.com/wyosa/geelark-go-sdk/oem"
	"github.com/wyosa/geelark-go-sdk/proxy"
	"github.com/wyosa/geelark-go-sdk/proxygroup"
	"github.com/wyosa/geelark-go-sdk/shell"
	"github.com/wyosa/geelark-go-sdk/tag"
	"github.com/wyosa/geelark-go-sdk/webhook"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

func TestDocumentedPaths(t *testing.T) {
	type call struct{ method, path, body string }
	type reply struct {
		status int
		body   string
	}
	var currentReply atomic.Value
	currentReply.Store(reply{http.StatusOK, `{"traceId":"test","code":0,"msg":"success","data":null}`})
	received := make(chan call, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received <- call{r.Method, r.URL.Path, r.Header.Get("Authorization")}
		w.Header().Set("Content-Type", "application/json")
		response := currentReply.Load().(reply)
		w.WriteHeader(response.status)
		_, _ = w.Write([]byte(response.body))
	}))
	defer srv.Close()
	cloud, err := geelark.NewClient("test-token", geelark.WithBaseURL(srv.URL), geelark.WithMinInterval(time.Nanosecond))
	if err != nil {
		t.Fatal(err)
	}
	local, err := geelark.NewBrowserClient("test-token", geelark.WithBaseURL(srv.URL), geelark.WithRemoteBaseURL(srv.URL), geelark.WithMinInterval(time.Nanosecond))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	cases := []struct {
		name, path string
		run        func() error
	}{
		{"cloudphone.Create", "/open/v1/phone/addNew", func() error { _, err := cloud.CloudPhones.Create(ctx, cloudphone.CreateRequest{}); return err }},
		{"cloudphone.List", "/open/v1/phone/list", func() error { _, err := cloud.CloudPhones.List(ctx, cloudphone.ListRequest{}); return err }},
		{"cloudphone.Status", "/open/v1/phone/status", func() error { _, err := cloud.CloudPhones.Status(ctx, cloudphone.StatusRequest{}); return err }},
		{"cloudphone.Start", "/open/v1/phone/start", func() error { _, err := cloud.CloudPhones.Start(ctx, cloudphone.StartRequest{}); return err }},
		{"cloudphone.Stop", "/open/v1/phone/stop", func() error { _, err := cloud.CloudPhones.Stop(ctx, cloudphone.StopRequest{}); return err }},
		{"cloudphone.Restart", "/open/v1/phone/restart", func() error { _, err := cloud.CloudPhones.Restart(ctx, cloudphone.RestartRequest{}); return err }},
		{"cloudphone.Delete", "/open/v1/phone/delete", func() error { _, err := cloud.CloudPhones.Delete(ctx, cloudphone.DeleteRequest{}); return err }},
		{"cloudphone.GetGPS", "/open/v1/phone/gps/get", func() error { _, err := cloud.CloudPhones.GetGPS(ctx, cloudphone.GetGPSRequest{}); return err }},
		{"cloudphone.SetGPS", "/open/v1/phone/gps/set", func() error { _, err := cloud.CloudPhones.SetGPS(ctx, cloudphone.SetGPSRequest{}); return err }},
		{"cloudphone.Update", "/open/v1/phone/detail/update", func() error { _, err := cloud.CloudPhones.Update(ctx, cloudphone.UpdateRequest{}); return err }},
		{"cloudphone.Reset", "/open/v2/phone/newOne", func() error { _, err := cloud.CloudPhones.Reset(ctx, cloudphone.ResetRequest{}); return err }},
		{"cloudphone.Clone", "/open/v1/phone/clone", func() error { _, err := cloud.CloudPhones.Clone(ctx, cloudphone.CloneRequest{}); return err }},
		{"cloudphone.Screenshot", "/open/v1/phone/screenShot", func() error { _, err := cloud.CloudPhones.Screenshot(ctx, cloudphone.ScreenshotRequest{}); return err }},
		{"cloudphone.ScreenshotResult", "/open/v1/phone/screenShot/result", func() error {
			_, err := cloud.CloudPhones.ScreenshotResult(ctx, cloudphone.ScreenshotResultRequest{})
			return err
		}},
		{"cloudphone.SetRoot", "/open/v1/root/setStatus", func() error { _, err := cloud.CloudPhones.SetRoot(ctx, cloudphone.SetRootRequest{}); return err }},
		{"cloudphone.GetRoot", "/open/v1/root/getStatus", func() error { _, err := cloud.CloudPhones.GetRoot(ctx, cloudphone.GetRootRequest{}); return err }},
		{"cloudphone.DeviceID", "/open/v1/phone/serialNum/get", func() error { _, err := cloud.CloudPhones.DeviceID(ctx, cloudphone.DeviceIDRequest{}); return err }},
		{"cloudphone.SendSMS", "/open/v1/phone/sendSms", func() error { _, err := cloud.CloudPhones.SendSMS(ctx, cloudphone.SendSMSRequest{}); return err }},
		{"cloudphone.Brands", "/open/v1/phone/brand/list", func() error { _, err := cloud.CloudPhones.Brands(ctx, cloudphone.BrandsRequest{}); return err }},
		{"cloudphone.TeamBrands", "/open/v1/phone/brand/teamList", func() error { _, err := cloud.CloudPhones.TeamBrands(ctx, cloudphone.TeamBrandsRequest{}); return err }},
		{"cloudphone.Transfer", "/open/v1/phone/transfer", func() error { _, err := cloud.CloudPhones.Transfer(ctx, cloudphone.TransferRequest{}); return err }},
		{"cloudphone.SetNetworkType", "/open/v1/phone/net/set", func() error {
			_, err := cloud.CloudPhones.SetNetworkType(ctx, cloudphone.SetNetworkTypeRequest{})
			return err
		}},
		{"cloudphone.HideAccessibility", "/open/v1/phone/hideAccessibility", func() error {
			_, err := cloud.CloudPhones.HideAccessibility(ctx, cloudphone.HideAccessibilityRequest{})
			return err
		}},
		{"cloudphone.MoveGroup", "/open/v1/phone/moveGroup", func() error { _, err := cloud.CloudPhones.MoveGroup(ctx, cloudphone.MoveGroupRequest{}); return err }},
		{"cloudphone.ImportContacts", "/open/v1/phone/importContacts", func() error {
			_, err := cloud.CloudPhones.ImportContacts(ctx, cloudphone.ImportContactsRequest{})
			return err
		}},
		{"cloudphone.ImportContactsResult", "/open/v1/phone/importContactsResult", func() error {
			_, err := cloud.CloudPhones.ImportContactsResult(ctx, cloudphone.ImportContactsResultRequest{})
			return err
		}},
		{"cloudphone.GetNetworkSettings", "/open/v1/phone/netConfig/get", func() error {
			_, err := cloud.CloudPhones.GetNetworkSettings(ctx, cloudphone.GetNetworkSettingsRequest{})
			return err
		}},
		{"cloudphone.SetNetworkSettings", "/open/v1/phone/netConfig/set", func() error {
			_, err := cloud.CloudPhones.SetNetworkSettings(ctx, cloudphone.SetNetworkSettingsRequest{})
			return err
		}},
		{"cloudphone.UIDump", "/open/v1/phone/uiDump", func() error { _, err := cloud.CloudPhones.UIDump(ctx, cloudphone.UIDumpRequest{}); return err }},
		{"cloudphone.UIDumpResult", "/open/v1/phone/uiDump/result", func() error {
			_, err := cloud.CloudPhones.UIDumpResult(ctx, cloudphone.UIDumpResultRequest{})
			return err
		}},
		{"automation.QueryTask", "/open/v1/task/query", func() error { _, err := cloud.Automation.QueryTask(ctx, automation.QueryTaskRequest{}); return err }},
		{"automation.ListTasks", "/open/v1/task/historyRecords", func() error { _, err := cloud.Automation.ListTasks(ctx, automation.ListTasksRequest{}); return err }},
		{"automation.CancelTask", "/open/v1/task/cancel", func() error { _, err := cloud.Automation.CancelTask(ctx, automation.CancelTaskRequest{}); return err }},
		{"automation.RetryTask", "/open/v1/task/restart", func() error { _, err := cloud.Automation.RetryTask(ctx, automation.RetryTaskRequest{}); return err }},
		{"automation.TaskDetail", "/open/v1/task/detail", func() error { _, err := cloud.Automation.TaskDetail(ctx, automation.TaskDetailRequest{}); return err }},
		{"automation.AddTask", "/open/v1/task/add", func() error { _, err := cloud.Automation.AddTask(ctx, automation.AddTaskRequest{}); return err }},
		{"automation.TikTokLogin", "/open/v1/rpa/task/tiktokLogin", func() error { _, err := cloud.Automation.TikTokLogin(ctx, automation.TikTokLoginRequest{}); return err }},
		{"automation.TikTokEdit", "/open/v1/rpa/task/tiktokEdit", func() error { _, err := cloud.Automation.TikTokEdit(ctx, automation.TikTokEditRequest{}); return err }},
		{"automation.TikTokRandomStar", "/open/v1/rpa/task/tiktokRandomStar", func() error {
			_, err := cloud.Automation.TikTokRandomStar(ctx, automation.TikTokRandomStarRequest{})
			return err
		}},
		{"automation.TikTokRandomStarAsia", "/open/v1/rpa/task/tiktokRandomStarAsia", func() error {
			_, err := cloud.Automation.TikTokRandomStarAsia(ctx, automation.TikTokRandomStarAsiaRequest{})
			return err
		}},
		{"automation.TikTokRandomComment", "/open/v1/rpa/task/tiktokRandomComment", func() error {
			_, err := cloud.Automation.TikTokRandomComment(ctx, automation.TikTokRandomCommentRequest{})
			return err
		}},
		{"automation.TikTokRandomCommentAsia", "/open/v1/rpa/task/tiktokRandomCommentAsia", func() error {
			_, err := cloud.Automation.TikTokRandomCommentAsia(ctx, automation.TikTokRandomCommentAsiaRequest{})
			return err
		}},
		{"automation.TikTokMessage", "/open/v1/rpa/task/tiktokMessage", func() error {
			_, err := cloud.Automation.TikTokMessage(ctx, automation.TikTokMessageRequest{})
			return err
		}},
		{"automation.TikTokMessageAsia", "/open/v1/rpa/task/tiktokMessageAsia", func() error {
			_, err := cloud.Automation.TikTokMessageAsia(ctx, automation.TikTokMessageAsiaRequest{})
			return err
		}},
		{"automation.TikTokRandomFollow", "/open/v1/rpa/task/tiktokRandomFollow", func() error {
			_, err := cloud.Automation.TikTokRandomFollow(ctx, automation.TikTokRandomFollowRequest{})
			return err
		}},
		{"automation.TikTokRandomFollowAsia", "/open/v1/rpa/task/tiktokRandomFollowAsia", func() error {
			_, err := cloud.Automation.TikTokRandomFollowAsia(ctx, automation.TikTokRandomFollowAsiaRequest{})
			return err
		}},
		{"automation.TikTokHide", "/open/v1/rpa/task/tiktokHide", func() error { _, err := cloud.Automation.TikTokHide(ctx, automation.TikTokHideRequest{}); return err }},
		{"automation.TikTokHideAsia", "/open/v1/rpa/task/tiktokHideAsia", func() error {
			_, err := cloud.Automation.TikTokHideAsia(ctx, automation.TikTokHideAsiaRequest{})
			return err
		}},
		{"automation.TikTokDel", "/open/v1/rpa/task/tiktokDel", func() error { _, err := cloud.Automation.TikTokDel(ctx, automation.TikTokDelRequest{}); return err }},
		{"automation.TikTokDelAsia", "/open/v1/rpa/task/tiktokDelAsia", func() error {
			_, err := cloud.Automation.TikTokDelAsia(ctx, automation.TikTokDelAsiaRequest{})
			return err
		}},
		{"automation.TikTokCommentInterception", "/open/v1/rpa/task/tiktokCommentInterception", func() error {
			_, err := cloud.Automation.TikTokCommentInterception(ctx, automation.TikTokCommentInterceptionRequest{})
			return err
		}},
		{"automation.TikTokCommentInterceptionAsia", "/open/v1/rpa/task/tiktokCommentInterceptionAsia", func() error {
			_, err := cloud.Automation.TikTokCommentInterceptionAsia(ctx, automation.TikTokCommentInterceptionAsiaRequest{})
			return err
		}},
		{"automation.TikTokDeleteComment", "/open/v1/rpa/task/tiktokDeleteComment", func() error {
			_, err := cloud.Automation.TikTokDeleteComment(ctx, automation.TikTokDeleteCommentRequest{})
			return err
		}},
		{"automation.TikTokDeleteCommentAsia", "/open/v1/rpa/task/tiktokDeleteCommentAsia", func() error {
			_, err := cloud.Automation.TikTokDeleteCommentAsia(ctx, automation.TikTokDeleteCommentAsiaRequest{})
			return err
		}},
		{"automation.FaceBookLogin", "/open/v1/rpa/task/faceBookLogin", func() error {
			_, err := cloud.Automation.FaceBookLogin(ctx, automation.FaceBookLoginRequest{})
			return err
		}},
		{"automation.FaceBookAutoComment", "/open/v1/rpa/task/faceBookAutoComment", func() error {
			_, err := cloud.Automation.FaceBookAutoComment(ctx, automation.FaceBookAutoCommentRequest{})
			return err
		}},
		{"automation.FaceBookPublish", "/open/v1/rpa/task/faceBookPublish", func() error {
			_, err := cloud.Automation.FaceBookPublish(ctx, automation.FaceBookPublishRequest{})
			return err
		}},
		{"automation.FaceBookActiveAccount", "/open/v1/rpa/task/faceBookActiveAccount", func() error {
			_, err := cloud.Automation.FaceBookActiveAccount(ctx, automation.FaceBookActiveAccountRequest{})
			return err
		}},
		{"automation.FaceBookPubReels", "/open/v1/rpa/task/faceBookPubReels", func() error {
			_, err := cloud.Automation.FaceBookPubReels(ctx, automation.FaceBookPubReelsRequest{})
			return err
		}},
		{"automation.FaceBookMessage", "/open/v1/rpa/task/faceBookMessage", func() error {
			_, err := cloud.Automation.FaceBookMessage(ctx, automation.FaceBookMessageRequest{})
			return err
		}},
		{"automation.FacebookReelsActive", "/open/v1/rpa/task/facebookReelsActive", func() error {
			_, err := cloud.Automation.FacebookReelsActive(ctx, automation.FacebookReelsActiveRequest{})
			return err
		}},
		{"automation.InstagramPubReels", "/open/v1/rpa/task/instagramPubReels", func() error {
			_, err := cloud.Automation.InstagramPubReels(ctx, automation.InstagramPubReelsRequest{})
			return err
		}},
		{"automation.InstagramWarmup", "/open/v1/rpa/task/instagramWarmup", func() error {
			_, err := cloud.Automation.InstagramWarmup(ctx, automation.InstagramWarmupRequest{})
			return err
		}},
		{"automation.InstagramMessage", "/open/v1/rpa/task/instagramMessage", func() error {
			_, err := cloud.Automation.InstagramMessage(ctx, automation.InstagramMessageRequest{})
			return err
		}},
		{"automation.InstagramPubReelsImages", "/open/v1/rpa/task/instagramPubReelsImages", func() error {
			_, err := cloud.Automation.InstagramPubReelsImages(ctx, automation.InstagramPubReelsImagesRequest{})
			return err
		}},
		{"automation.InstagramLogin", "/open/v1/rpa/task/instagramLogin", func() error {
			_, err := cloud.Automation.InstagramLogin(ctx, automation.InstagramLoginRequest{})
			return err
		}},
		{"automation.InstagramEdit", "/open/v1/rpa/task/instagramEdit", func() error {
			_, err := cloud.Automation.InstagramEdit(ctx, automation.InstagramEditRequest{})
			return err
		}},
		{"automation.InstagramAccountStatus", "/open/v1/rpa/task/instagramAccountStatus", func() error {
			_, err := cloud.Automation.InstagramAccountStatus(ctx, automation.InstagramAccountStatusRequest{})
			return err
		}},
		{"automation.InstagramFollowAccount", "/open/v1/rpa/task/instagramFollowAccount", func() error {
			_, err := cloud.Automation.InstagramFollowAccount(ctx, automation.InstagramFollowAccountRequest{})
			return err
		}},
		{"automation.InstagramAiComment", "/open/v1/rpa/task/instagramAiComment", func() error {
			_, err := cloud.Automation.InstagramAiComment(ctx, automation.InstagramAiCommentRequest{})
			return err
		}},
		{"automation.YouTubePubShort", "/open/v1/rpa/task/youtubePubShort", func() error {
			_, err := cloud.Automation.YouTubePubShort(ctx, automation.YouTubePubShortRequest{})
			return err
		}},
		{"automation.YouTubePubVideo", "/open/v1/rpa/task/youtubePubVideo", func() error {
			_, err := cloud.Automation.YouTubePubVideo(ctx, automation.YouTubePubVideoRequest{})
			return err
		}},
		{"automation.YouTubeActiveAccount", "/open/v1/rpa/task/youTubeActiveAccount", func() error {
			_, err := cloud.Automation.YouTubeActiveAccount(ctx, automation.YouTubeActiveAccountRequest{})
			return err
		}},
		{"automation.YouTubeEdit", "/open/v1/rpa/task/youtubeEdit", func() error { _, err := cloud.Automation.YouTubeEdit(ctx, automation.YouTubeEditRequest{}); return err }},
		{"automation.GoogleLogin", "/open/v1/rpa/task/googleLogin", func() error { _, err := cloud.Automation.GoogleLogin(ctx, automation.GoogleLoginRequest{}); return err }},
		{"automation.GoogleAppDownload", "/open/v1/rpa/task/googleAppDownload", func() error {
			_, err := cloud.Automation.GoogleAppDownload(ctx, automation.GoogleAppDownloadRequest{})
			return err
		}},
		{"automation.GoogleAppBrowser", "/open/v1/rpa/task/googleAppBrowser", func() error {
			_, err := cloud.Automation.GoogleAppBrowser(ctx, automation.GoogleAppBrowserRequest{})
			return err
		}},
		{"automation.SheinLogin", "/open/v1/rpa/task/sheinLogin", func() error { _, err := cloud.Automation.SheinLogin(ctx, automation.SheinLoginRequest{}); return err }},
		{"automation.XPublish", "/open/v1/rpa/task/xPublish", func() error { _, err := cloud.Automation.XPublish(ctx, automation.XPublishRequest{}); return err }},
		{"automation.ListFlows", "/open/v1/task/flow/list", func() error { _, err := cloud.Automation.ListFlows(ctx, automation.ListFlowsRequest{}); return err }},
		{"automation.CreateCustomTask", "/open/v1/task/rpa/add", func() error {
			_, err := cloud.Automation.CreateCustomTask(ctx, automation.CreateCustomTaskRequest{})
			return err
		}},
		{"automation.ImportFlow", "/open/v1/task/flow/import", func() error { _, err := cloud.Automation.ImportFlow(ctx, automation.ImportFlowRequest{}); return err }},
		{"automation.ExportFlow", "/open/v1/task/flow/export", func() error { _, err := cloud.Automation.ExportFlow(ctx, automation.ExportFlowRequest{}); return err }},
		{"automation.MultiPlatformVideoDistribution", "/open/v1/rpa/task/multiPlatformVideoDistribution", func() error {
			_, err := cloud.Automation.MultiPlatformVideoDistribution(ctx, automation.MultiPlatformVideoDistributionRequest{})
			return err
		}},
		{"automation.FileUpload", "/open/v1/rpa/task/fileUpload", func() error { _, err := cloud.Automation.FileUpload(ctx, automation.FileUploadRequest{}); return err }},
		{"automation.ImportContacts", "/open/v1/rpa/task/importContacts", func() error {
			_, err := cloud.Automation.ImportContacts(ctx, automation.ImportContactsRequest{})
			return err
		}},
		{"automation.KeyboxUpload", "/open/v1/rpa/task/keyboxUpload", func() error {
			_, err := cloud.Automation.KeyboxUpload(ctx, automation.KeyboxUploadRequest{})
			return err
		}},
		{"automation.RedditWarmup", "/open/v1/rpa/task/redditWarmup", func() error {
			_, err := cloud.Automation.RedditWarmup(ctx, automation.RedditWarmupRequest{})
			return err
		}},
		{"automation.RedditVideo", "/open/v1/rpa/task/redditVideo", func() error { _, err := cloud.Automation.RedditVideo(ctx, automation.RedditVideoRequest{}); return err }},
		{"automation.RedditImage", "/open/v1/rpa/task/redditImage", func() error { _, err := cloud.Automation.RedditImage(ctx, automation.RedditImageRequest{}); return err }},
		{"automation.PinterestVideo", "/open/v1/rpa/task/pinterestVideo", func() error {
			_, err := cloud.Automation.PinterestVideo(ctx, automation.PinterestVideoRequest{})
			return err
		}},
		{"automation.PinterestImage", "/open/v1/rpa/task/pinterestImage", func() error {
			_, err := cloud.Automation.PinterestImage(ctx, automation.PinterestImageRequest{})
			return err
		}},
		{"automation.ThreadsVideo", "/open/v1/rpa/task/threadsVideo", func() error {
			_, err := cloud.Automation.ThreadsVideo(ctx, automation.ThreadsVideoRequest{})
			return err
		}},
		{"automation.ThreadsImage", "/open/v1/rpa/task/threadsImage", func() error {
			_, err := cloud.Automation.ThreadsImage(ctx, automation.ThreadsImageRequest{})
			return err
		}},
		{"shell.Execute", "/open/v1/shell/execute", func() error { _, err := cloud.Shell.Execute(ctx, shell.ExecuteRequest{}); return err }},
		{"app.ListInstalled", "/open/v1/app/list", func() error { _, err := cloud.Apps.ListInstalled(ctx, app.ListInstalledRequest{}); return err }},
		{"app.ListAvailable", "/open/v1/app/installable/list", func() error { _, err := cloud.Apps.ListAvailable(ctx, app.ListAvailableRequest{}); return err }},
		{"app.Install", "/open/v1/app/install", func() error { _, err := cloud.Apps.Install(ctx, app.InstallRequest{}); return err }},
		{"app.Start", "/open/v1/app/start", func() error { _, err := cloud.Apps.Start(ctx, app.StartRequest{}); return err }},
		{"app.Stop", "/open/v1/app/stop", func() error { _, err := cloud.Apps.Stop(ctx, app.StopRequest{}); return err }},
		{"app.Uninstall", "/open/v1/app/uninstall", func() error { _, err := cloud.Apps.Uninstall(ctx, app.UninstallRequest{}); return err }},
		{"app.Upload", "/open/v1/app/upload", func() error { _, err := cloud.Apps.Upload(ctx, app.UploadRequest{}); return err }},
		{"app.UploadStatus", "/open/v1/app/upload/status", func() error { _, err := cloud.Apps.UploadStatus(ctx, app.UploadStatusRequest{}); return err }},
		{"app.ListShop", "/open/v1/app/shop/list", func() error { _, err := cloud.Apps.ListShop(ctx, app.ListShopRequest{}); return err }},
		{"app.AddToTeam", "/open/v1/app/add", func() error { _, err := cloud.Apps.AddToTeam(ctx, app.AddToTeamRequest{}); return err }},
		{"app.ListTeam", "/open/v1/app/teamApp/list", func() error { _, err := cloud.Apps.ListTeam(ctx, app.ListTeamRequest{}); return err }},
		{"app.RemoveTeam", "/open/v1/app/remove", func() error { _, err := cloud.Apps.RemoveTeam(ctx, app.RemoveTeamRequest{}); return err }},
		{"app.SetAutoInstall", "/open/v1/app/setStatus", func() error { _, err := cloud.Apps.SetAutoInstall(ctx, app.SetAutoInstallRequest{}); return err }},
		{"app.SetAuthorization", "/open/v1/app/auth/status", func() error { _, err := cloud.Apps.SetAuthorization(ctx, app.SetAuthorizationRequest{}); return err }},
		{"app.SetRoot", "/open/v1/app/root", func() error { _, err := cloud.Apps.SetRoot(ctx, app.SetRootRequest{}); return err }},
		{"app.BatchOperate", "/open/v1/app/operation/batch", func() error { _, err := cloud.Apps.BatchOperate(ctx, app.BatchOperateRequest{}); return err }},
		{"app.SetKeepAlive", "/open/v1/app/setKeepAlive", func() error { _, err := cloud.Apps.SetKeepAlive(ctx, app.SetKeepAliveRequest{}); return err }},
		{"app.SetAutoStart", "/open/v1/app/setAutoStart", func() error { _, err := cloud.Apps.SetAutoStart(ctx, app.SetAutoStartRequest{}); return err }},
		{"app.SetHidden", "/open/v1/app/setHide", func() error { _, err := cloud.Apps.SetHidden(ctx, app.SetHiddenRequest{}); return err }},
		{"adb.SetStatus", "/open/v1/adb/setStatus", func() error { _, err := cloud.ADB.SetStatus(ctx, adb.SetStatusRequest{}); return err }},
		{"adb.GetInfo", "/open/v1/adb/getData", func() error { _, err := cloud.ADB.GetInfo(ctx, adb.GetInfoRequest{}); return err }},
		{"files.GetUploadURL", "/open/v1/upload/getUrl", func() error { _, err := cloud.Files.GetUploadURL(ctx, files.GetUploadURLRequest{}); return err }},
		{"files.UploadToPhone", "/open/v1/phone/uploadFile", func() error { _, err := cloud.Files.UploadToPhone(ctx, files.UploadToPhoneRequest{}); return err }},
		{"files.UploadResult", "/open/v1/phone/uploadFile/result", func() error { _, err := cloud.Files.UploadResult(ctx, files.UploadResultRequest{}); return err }},
		{"files.UploadKeybox", "/open/v1/phone/keyboxUpload", func() error { _, err := cloud.Files.UploadKeybox(ctx, files.UploadKeyboxRequest{}); return err }},
		{"files.KeyboxResult", "/open/v1/phone/keyboxUpload/result", func() error { _, err := cloud.Files.KeyboxResult(ctx, files.KeyboxResultRequest{}); return err }},
		{"library.GetUploadURL", "/open/v1/material/getUploadUrl", func() error { _, err := cloud.Library.GetUploadURL(ctx, library.GetUploadURLRequest{}); return err }},
		{"library.Create", "/open/v1/material/create", func() error { _, err := cloud.Library.Create(ctx, library.CreateRequest{}); return err }},
		{"library.Delete", "/open/v1/material/del", func() error { _, err := cloud.Library.Delete(ctx, library.DeleteRequest{}); return err }},
		{"library.SetTag", "/open/v1/material/tag/set", func() error { _, err := cloud.Library.SetTag(ctx, library.SetTagRequest{}); return err }},
		{"library.Search", "/open/v1/material/search", func() error { _, err := cloud.Library.Search(ctx, library.SearchRequest{}); return err }},
		{"library.CreateTag", "/open/v1/material/tag/create", func() error { _, err := cloud.Library.CreateTag(ctx, library.CreateTagRequest{}); return err }},
		{"library.DeleteTag", "/open/v1/material/tag/del", func() error { _, err := cloud.Library.DeleteTag(ctx, library.DeleteTagRequest{}); return err }},
		{"library.SearchTags", "/open/v1/material/tag/search", func() error { _, err := cloud.Library.SearchTags(ctx, library.SearchTagsRequest{}); return err }},
		{"webhook.SetURL", "/open/v1/callback/set", func() error { _, err := cloud.Webhooks.SetURL(ctx, webhook.SetURLRequest{}); return err }},
		{"webhook.GetURL", "/open/v1/callback/get", func() error { _, err := cloud.Webhooks.GetURL(ctx, webhook.GetURLRequest{}); return err }},
		{"oem.Customization", "/open/v1/phone/customization", func() error { _, err := cloud.OEM.Customization(ctx, oem.CustomizationRequest{}); return err }},
		{"analytics.Data", "/open/v1/analytics/data", func() error { _, err := cloud.Analytics.Data(ctx, analytics.DataRequest{}); return err }},
		{"analytics.AddAccounts", "/open/v1/analytics/accounts/add", func() error { _, err := cloud.Analytics.AddAccounts(ctx, analytics.AddAccountsRequest{}); return err }},
		{"analytics.DeleteAccount", "/open/v1/analytics/accounts/delete", func() error {
			_, err := cloud.Analytics.DeleteAccount(ctx, analytics.DeleteAccountRequest{})
			return err
		}},
		{"analytics.ListAccounts", "/open/v1/analytics/accounts/list", func() error { _, err := cloud.Analytics.ListAccounts(ctx, analytics.ListAccountsRequest{}); return err }},
		{"analytics.UpdateAccount", "/open/v1/analytics/accounts/update", func() error {
			_, err := cloud.Analytics.UpdateAccount(ctx, analytics.UpdateAccountRequest{})
			return err
		}},
		{"analytics.CreateTag", "/open/v1/analytics/tags/create", func() error { _, err := cloud.Analytics.CreateTag(ctx, analytics.CreateTagRequest{}); return err }},
		{"analytics.ListTags", "/open/v1/analytics/tags/list", func() error { _, err := cloud.Analytics.ListTags(ctx, analytics.ListTagsRequest{}); return err }},
		{"analytics.UpdateTag", "/open/v1/analytics/tags/update", func() error { _, err := cloud.Analytics.UpdateTag(ctx, analytics.UpdateTagRequest{}); return err }},
		{"analytics.DeleteTags", "/open/v1/analytics/tags/del", func() error { _, err := cloud.Analytics.DeleteTags(ctx, analytics.DeleteTagsRequest{}); return err }},
		{"livestream.Start", "/open/v1/phone/videoPush/upload", func() error { _, err := cloud.LiveStream.Start(ctx, livestream.StartRequest{}); return err }},
		{"livestream.Result", "/open/v1/phone/videoPush/upload/result", func() error { _, err := cloud.LiveStream.Result(ctx, livestream.ResultRequest{}); return err }},
		{"livestream.Stop", "/open/v1/phone/videoPush/stop", func() error { _, err := cloud.LiveStream.Stop(ctx, livestream.StopRequest{}); return err }},
		{"browser.Create", "/api/v1/browser/create", func() error { _, err := local.Browsers.Create(ctx, browser.CreateRequest{}); return err }},
		{"browser.Update", "/api/v1/browser/update", func() error { _, err := local.Browsers.Update(ctx, browser.UpdateRequest{}); return err }},
		{"browser.Delete", "/api/v1/browser/delete", func() error { _, err := local.Browsers.Delete(ctx, browser.DeleteRequest{}); return err }},
		{"browser.List", "/api/v1/browser/list", func() error { _, err := local.Browsers.List(ctx, browser.ListRequest{}); return err }},
		{"browser.Start", "/api/v1/browser/start", func() error { _, err := local.Browsers.Start(ctx, browser.StartRequest{}); return err }},
		{"browser.Stop", "/api/v1/browser/stop", func() error { _, err := local.Browsers.Stop(ctx, browser.StopRequest{}); return err }},
		{"browser.Transfer", "/api/v1/browser/transfer", func() error { _, err := local.Browsers.Transfer(ctx, browser.TransferRequest{}); return err }},
		{"browser.APIStatus", "/api/v1/status", func() error { _, err := local.Browsers.APIStatus(ctx, browser.APIStatusRequest{}); return err }},
		{"browser.Status", "/api/v1/browser/status", func() error { _, err := local.Browsers.Status(ctx, browser.StatusRequest{}); return err }},
		{"browser.ClearCache", "/api/v1/browser/deleteCache", func() error { _, err := local.Browsers.ClearCache(ctx, browser.ClearCacheRequest{}); return err }},
		{"browser.MoveGroup", "/api/v1/browser/moveGroup", func() error { _, err := local.Browsers.MoveGroup(ctx, browser.MoveGroupRequest{}); return err }},
		{"browser.Cookies", "/api/v1/browser/getCookie", func() error { _, err := local.Browsers.Cookies(ctx, browser.CookiesRequest{}); return err }},
		{"browser.SetBookmarks", "/api/v1/browser/setBookmark", func() error { _, err := local.Browsers.SetBookmarks(ctx, browser.SetBookmarksRequest{}); return err }},
		{"browser.GetBookmarks", "/api/v1/browser/getBookmark", func() error { _, err := local.Browsers.GetBookmarks(ctx, browser.GetBookmarksRequest{}); return err }},
		{"browser.Clone", "/api/v1/browser/clone", func() error { _, err := local.Browsers.Clone(ctx, browser.CloneRequest{}); return err }},
		{"browser.Kernels", "/api/v1/browser/getKernelsList", func() error { _, err := local.Browsers.Kernels(ctx, browser.KernelsRequest{}); return err }},
		{"browser.UpdateKernels", "/api/v1/browser/updateKernels", func() error { _, err := local.Browsers.UpdateKernels(ctx, browser.UpdateKernelsRequest{}); return err }},
		{"browser.ExtensionGroups", "/open/v1/browser/extGroup/list", func() error {
			_, err := local.Browsers.ExtensionGroups(ctx, browser.ExtensionGroupsRequest{})
			return err
		}},
		{"browserautomation.Search", "/api/v1/browser/task/search", func() error { _, err := local.Automation.Search(ctx, browserautomation.SearchRequest{}); return err }},
		{"browserautomation.Cancel", "/api/v1/browser/task/cancel", func() error { _, err := local.Automation.Cancel(ctx, browserautomation.CancelRequest{}); return err }},
		{"browserautomation.Restart", "/api/v1/browser/task/restart", func() error { _, err := local.Automation.Restart(ctx, browserautomation.RestartRequest{}); return err }},
		{"browserautomation.Detail", "/api/v1/browser/task/detail", func() error { _, err := local.Automation.Detail(ctx, browserautomation.DetailRequest{}); return err }},
		{"browserautomation.Flow", "/api/v1/browser/task/flow", func() error { _, err := local.Automation.Flow(ctx, browserautomation.FlowRequest{}); return err }},
		{"browserautomation.Add", "/api/v1/browser/task/add", func() error { _, err := local.Automation.Add(ctx, browserautomation.AddRequest{}); return err }},
		{"browserautomation.TikTokSearch", "/api/v1/browser/task/tiktokSearch", func() error {
			_, err := local.Automation.TikTokSearch(ctx, browserautomation.TikTokSearchRequest{})
			return err
		}},
		{"browserautomation.TikTokComment", "/api/v1/browser/task/tiktokComment", func() error {
			_, err := local.Automation.TikTokComment(ctx, browserautomation.TikTokCommentRequest{})
			return err
		}},
		{"browserautomation.TikTokLike", "/api/v1/browser/task/tiktokLike", func() error {
			_, err := local.Automation.TikTokLike(ctx, browserautomation.TikTokLikeRequest{})
			return err
		}},
		{"browserautomation.FacebookPost", "/api/v1/browser/task/facebookPost", func() error {
			_, err := local.Automation.FacebookPost(ctx, browserautomation.FacebookPostRequest{})
			return err
		}},
		{"browserautomation.FacebookFriends", "/api/v1/browser/task/facebookFriends", func() error {
			_, err := local.Automation.FacebookFriends(ctx, browserautomation.FacebookFriendsRequest{})
			return err
		}},
		{"browserautomation.FacebookLike", "/api/v1/browser/task/facebookLike", func() error {
			_, err := local.Automation.FacebookLike(ctx, browserautomation.FacebookLikeRequest{})
			return err
		}},
		{"browserautomation.FacebookHomepage", "/api/v1/browser/task/facebookHomepage", func() error {
			_, err := local.Automation.FacebookHomepage(ctx, browserautomation.FacebookHomepageRequest{})
			return err
		}},
		{"browserautomation.InstagramLike", "/api/v1/browser/task/instagramLike", func() error {
			_, err := local.Automation.InstagramLike(ctx, browserautomation.InstagramLikeRequest{})
			return err
		}},
		{"browserautomation.InstagramSearch", "/api/v1/browser/task/instagramSearch", func() error {
			_, err := local.Automation.InstagramSearch(ctx, browserautomation.InstagramSearchRequest{})
			return err
		}},
		{"browserautomation.YouTubeWatch", "/api/v1/browser/task/youtubeWatch", func() error {
			_, err := local.Automation.YouTubeWatch(ctx, browserautomation.YouTubeWatchRequest{})
			return err
		}},
		{"browserautomation.XPost", "/api/v1/browser/task/xPost", func() error { _, err := local.Automation.XPost(ctx, browserautomation.XPostRequest{}); return err }},
		{"browserautomation.XLike", "/api/v1/browser/task/xLike", func() error { _, err := local.Automation.XLike(ctx, browserautomation.XLikeRequest{}); return err }},
		{"browserautomation.RedditLike", "/api/v1/browser/task/redditLike", func() error {
			_, err := local.Automation.RedditLike(ctx, browserautomation.RedditLikeRequest{})
			return err
		}},
		{"browserautomation.CookieBot", "/api/v1/browser/task/cookieBot", func() error {
			_, err := local.Automation.CookieBot(ctx, browserautomation.CookieBotRequest{})
			return err
		}},
		{"proxy.Add", "/open/v1/proxy/add", func() error { _, err := cloud.Proxies.Add(ctx, proxy.AddRequest{}); return err }},
		{"proxy.Update", "/open/v1/proxy/update", func() error { _, err := cloud.Proxies.Update(ctx, proxy.UpdateRequest{}); return err }},
		{"proxy.Delete", "/open/v1/proxy/delete", func() error { _, err := cloud.Proxies.Delete(ctx, proxy.DeleteRequest{}); return err }},
		{"proxy.List", "/open/v1/proxy/list", func() error { _, err := cloud.Proxies.List(ctx, proxy.ListRequest{}); return err }},
		{"proxy.Check", "/open/v1/proxy/check", func() error { _, err := cloud.Proxies.Check(ctx, proxy.CheckRequest{}); return err }},
		{"proxygroup.Create", "/open/v1/proxyGroup/add", func() error { _, err := cloud.ProxyGroups.Create(ctx, proxygroup.CreateRequest{}); return err }},
		{"proxygroup.List", "/open/v1/proxyGroup/search", func() error { _, err := cloud.ProxyGroups.List(ctx, proxygroup.ListRequest{}); return err }},
		{"proxygroup.Update", "/open/v1/proxyGroup/update", func() error { _, err := cloud.ProxyGroups.Update(ctx, proxygroup.UpdateRequest{}); return err }},
		{"proxygroup.Delete", "/open/v1/proxyGroup/del", func() error { _, err := cloud.ProxyGroups.Delete(ctx, proxygroup.DeleteRequest{}); return err }},
		{"tag.Create", "/open/v1/tag/add", func() error { _, err := cloud.Tags.Create(ctx, tag.CreateRequest{}); return err }},
		{"tag.Delete", "/open/v1/tag/delete", func() error { _, err := cloud.Tags.Delete(ctx, tag.DeleteRequest{}); return err }},
		{"tag.List", "/open/v1/tag/list", func() error { _, err := cloud.Tags.List(ctx, tag.ListRequest{}); return err }},
		{"tag.Update", "/open/v1/tag/update", func() error { _, err := cloud.Tags.Update(ctx, tag.UpdateRequest{}); return err }},
		{"group.Create", "/open/v1/group/add", func() error { _, err := cloud.Groups.Create(ctx, group.CreateRequest{}); return err }},
		{"group.Delete", "/open/v1/group/delete", func() error { _, err := cloud.Groups.Delete(ctx, group.DeleteRequest{}); return err }},
		{"group.List", "/open/v1/group/list", func() error { _, err := cloud.Groups.List(ctx, group.ListRequest{}); return err }},
		{"group.Update", "/open/v1/group/update", func() error { _, err := cloud.Groups.Update(ctx, group.UpdateRequest{}); return err }},
		{"billing.Balance", "/open/v1/pay/wallet", func() error { _, err := cloud.Billing.Balance(ctx, billing.BalanceRequest{}); return err }},
		{"billing.Plans", "/open/v1/pay/profiles/list", func() error { _, err := cloud.Billing.Plans(ctx, billing.PlansRequest{}); return err }},
		{"billing.ChangePlan", "/open/v1/pay/plan/upgrade", func() error { _, err := cloud.Billing.ChangePlan(ctx, billing.ChangePlanRequest{}); return err }},
		{"billing.RenewPlan", "/open/v1/pay/plan/continue", func() error { _, err := cloud.Billing.RenewPlan(ctx, billing.RenewPlanRequest{}); return err }},
		{"billing.PlanInfo", "/open/v1/pay/plan/info", func() error { _, err := cloud.Billing.PlanInfo(ctx, billing.PlanInfoRequest{}); return err }},
		{"billing.Transactions", "/open/v1/billing/transaction/detail", func() error { _, err := cloud.Billing.Transactions(ctx, billing.TransactionsRequest{}); return err }},
		{"billing.BuyTime", "/open/v1/pay/timeAddOn/buy", func() error { _, err := cloud.Billing.BuyTime(ctx, billing.BuyTimeRequest{}); return err }},
	}
	manifest, err := os.ReadFile("docs/coverage.json")
	if err != nil {
		t.Fatal(err)
	}
	var coverage struct {
		Endpoints []struct{ Package, Method, Path string }
	}
	if err := json.Unmarshal(manifest, &coverage); err != nil {
		t.Fatal(err)
	}
	if len(coverage.Endpoints) != len(cases) {
		t.Fatalf("%d endpoint tests for %d documented endpoints", len(cases), len(coverage.Endpoints))
	}
	wantPaths := make(map[string]string, len(coverage.Endpoints))
	for _, endpoint := range coverage.Endpoints {
		wantPaths[endpoint.Package+"."+endpoint.Method] = endpoint.Path
	}
	for _, tc := range cases {
		if wantPaths[tc.name] != tc.path {
			t.Fatalf("%s: path %q differs from coverage manifest %q", tc.name, tc.path, wantPaths[tc.name])
		}
		delete(wantPaths, tc.name)
		t.Run(tc.name, func(t *testing.T) {
			for _, scenario := range []struct {
				name  string
				reply reply
				code  int
			}{
				{"success", reply{200, `{"traceId":"test","code":0,"data":null}`}, 0},
				{"partial_success", reply{200, `{"traceId":"test","code":40006,"msg":"partial success","data":{"successAmount":1}}`}, 40006},
				{"http_error", reply{503, `{"traceId":"test","code":50000,"msg":"unavailable"}`}, 50000},
			} {
				t.Run(scenario.name, func(t *testing.T) {
					currentReply.Store(scenario.reply)
					err := tc.run()
					if scenario.code == 0 {
						if err != nil {
							t.Fatal(err)
						}
					} else {
						var apiErr *geelark.APIError
						if !errors.As(err, &apiErr) || apiErr.Code != scenario.code || apiErr.HTTPStatus != scenario.reply.status {
							t.Fatalf("got %v, want APIError code %d, HTTP %d", err, scenario.code, scenario.reply.status)
						}
						if scenario.code == 40006 {
							var partial struct {
								SuccessAmount int `json:"successAmount"`
							}
							if err := apiErr.DecodeData(&partial); err != nil || partial.SuccessAmount != 1 {
								t.Fatalf("partial success data lost: %s (%v)", apiErr.Data, err)
							}
						}
					}
					got := <-received
					if got.method != http.MethodPost || got.path != tc.path || got.body != "Bearer test-token" {
						t.Errorf("got method=%s path=%s auth=%s, want POST %s with Bearer token", got.method, got.path, got.body, tc.path)
					}
				})
			}
		})
	}
}
