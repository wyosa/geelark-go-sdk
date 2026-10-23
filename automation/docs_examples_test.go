package automation

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Two request examples are omitted: one uses a capitalized JSON key absent from the parameter table,
// and one omits the documented required duration.
func TestDocumentationExamples(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		newValue func() any
	}{
		{"QueryTaskRequest", "{\"ids\":[\"123321\",\"456654\"]}", func() any { return new(QueryTaskRequest) }},
		{"QueryTaskResponse", "{\"total\":1,\"items\":[{\"id\":\"123456ABCDEF\",\"planName\":\"plan123456ABCDEF\",\"taskType\":2,\"serialName\":\"test\",\"envId\":\"123456654321\",\"scheduleAt\":1718744459,\"status\":3}]}", func() any { return new(QueryTaskResponse) }},
		{"ListTasksRequest", "{\"size\":10,\"lastId\":\"574376043897425896\"}", func() any { return new(ListTasksRequest) }},
		{"ListTasksResponse", "{\"total\":1,\"items\":[{\"id\":\"123456ABCDEF\",\"planName\":\"plan123456ABCDEF\",\"taskType\":2,\"serialName\":\"test\",\"envId\":\"123456654321\",\"scheduleAt\":1718744459,\"status\":3}]}", func() any { return new(ListTasksResponse) }},
		{"CancelTaskRequest", "{\"ids\":[\"123321\",\"456654\"]}", func() any { return new(CancelTaskRequest) }},
		{"RetryTaskRequest", "{\"ids\":[\"123321\",\"456654\"]}", func() any { return new(RetryTaskRequest) }},
		{"TaskDetailRequest", "{\"id\":\"1234567898\",\"searchAfter\":[1749004889852]}", func() any { return new(TaskDetailRequest) }},
		{"TaskDetailResponse", "{\"id\":\"123456ABCDEF\",\"planName\":\"plan123456ABCDEF\",\"taskType\":2,\"serialName\":\"test\",\"envId\":\"123456654321\",\"scheduleAt\":1718744459,\"status\":3,\"cost\":60,\"resultImages\":[\"https://material.geelark.com/geeUserUix/569544577254450679_1929760564478226432_20250604104103.jpg\"],\"logs\":[\"[2025-06-04 02:38:00 025] Waiting for execution\"],\"searchAfter\":[1749004889853],\"logContinue\":false}", func() any { return new(TaskDetailResponse) }},
		{"AddTaskResponse", "{\"taskIds\":[\"123456ABCEDF\"]}", func() any { return new(AddTaskResponse) }},
		{"TikTokLoginRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"account\":\"test@gmail.com\",\"password\":\"123456\"}", func() any { return new(TikTokLoginRequest) }},
		{"TikTokLoginResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokLoginResponse) }},
		{"TikTokEditRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"avatar\":\"https://singapore-upgrade.geelark.com/a.jpg\",\"nickName\":\"test\",\"bio\":\"test\",\"site\":\"https://www.abc.com\",\"username\":\"user\"}", func() any { return new(TikTokEditRequest) }},
		{"TikTokEditResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokEditResponse) }},
		{"TikTokRandomStarRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\"}", func() any { return new(TikTokRandomStarRequest) }},
		{"TikTokRandomStarResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokRandomStarResponse) }},
		{"TikTokRandomStarAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\"}", func() any { return new(TikTokRandomStarAsiaRequest) }},
		{"TikTokRandomStarAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokRandomStarAsiaResponse) }},
		{"TikTokRandomCommentRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"useAi\":2,\"comment\":\"test\"}", func() any { return new(TikTokRandomCommentRequest) }},
		{"TikTokRandomCommentResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokRandomCommentResponse) }},
		{"TikTokRandomCommentAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"useAi\":2,\"comment\":\"test\"}", func() any { return new(TikTokRandomCommentAsiaRequest) }},
		{"TikTokRandomCommentAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokRandomCommentAsiaResponse) }},
		{"TikTokMessageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"usernames\":[\"user\"],\"content\":\"123456\"}", func() any { return new(TikTokMessageRequest) }},
		{"TikTokMessageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokMessageResponse) }},
		{"TikTokMessageAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"usernames\":[\"user\"],\"content\":\"123456\"}", func() any { return new(TikTokMessageAsiaRequest) }},
		{"TikTokMessageAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokMessageAsiaResponse) }},
		{"TikTokRandomFollowRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"followProbability\":30}", func() any { return new(TikTokRandomFollowRequest) }},
		{"TikTokRandomFollowResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokRandomFollowResponse) }},
		{"TikTokRandomFollowAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"followProbability\":30}", func() any { return new(TikTokRandomFollowAsiaRequest) }},
		{"TikTokRandomFollowAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokRandomFollowAsiaResponse) }},
		{"TikTokHideRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\"}", func() any { return new(TikTokHideRequest) }},
		{"TikTokHideResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokHideResponse) }},
		{"TikTokHideAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\"}", func() any { return new(TikTokHideAsiaRequest) }},
		{"TikTokHideAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokHideAsiaResponse) }},
		{"TikTokDelRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\"}", func() any { return new(TikTokDelRequest) }},
		{"TikTokDelResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokDelResponse) }},
		{"TikTokDelAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\"}", func() any { return new(TikTokDelAsiaRequest) }},
		{"TikTokDelAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokDelAsiaResponse) }},
		{"TikTokCommentInterceptionRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"executionTime\":1,\"videoKeywords\":[\"test\"],\"commentKeywords\":[\"test\"],\"replyComment\":[\"test\"]}", func() any { return new(TikTokCommentInterceptionRequest) }},
		{"TikTokCommentInterceptionResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokCommentInterceptionResponse) }},
		{"TikTokCommentInterceptionAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"executionTime\":1,\"videoKeywords\":[\"test\"],\"commentKeywords\":[\"test\"],\"replyComment\":[\"test\"]}", func() any { return new(TikTokCommentInterceptionAsiaRequest) }},
		{"TikTokCommentInterceptionAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokCommentInterceptionAsiaResponse) }},
		{"TikTokDeleteCommentRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"keywords\":[\"hello\"]}", func() any { return new(TikTokDeleteCommentRequest) }},
		{"TikTokDeleteCommentResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokDeleteCommentResponse) }},
		{"TikTokDeleteCommentAsiaRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"keywords\":[\"hello\"]}", func() any { return new(TikTokDeleteCommentAsiaRequest) }},
		{"TikTokDeleteCommentAsiaResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokDeleteCommentAsiaResponse) }},
		{"FaceBookLoginRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"email\":\"test@gmail.com\",\"password\":\"123456\"}", func() any { return new(FaceBookLoginRequest) }},
		{"FaceBookLoginResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FaceBookLoginResponse) }},
		{"FaceBookAutoCommentRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"postAddress\":\"https://abc.com\",\"comment\":[\"test1\",\"test2\"],\"keyword\":[\"k1\",\"k2\"]}", func() any { return new(FaceBookAutoCommentRequest) }},
		{"FaceBookAutoCommentResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FaceBookAutoCommentResponse) }},
		{"FaceBookPublishRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"video\":[\"https://singapore-upgrade.geelark.com/a.mp4\"]}", func() any { return new(FaceBookPublishRequest) }},
		{"FaceBookPublishResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FaceBookPublishResponse) }},
		{"FaceBookActiveAccountRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"postAddress\":\"https://abc.com\",\"browsePostsNum\":10,\"keyword\":[\"k1\",\"k2\"]}", func() any { return new(FaceBookActiveAccountRequest) }},
		{"FaceBookActiveAccountResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FaceBookActiveAccountResponse) }},
		{"FaceBookPubReelsRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"description\":\"description\",\"video\":\"https://material.geelark.com/a.mp4\"}", func() any { return new(FaceBookPubReelsRequest) }},
		{"FaceBookPubReelsResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FaceBookPubReelsResponse) }},
		{"FaceBookMessageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"usernames\":[\"user\"],\"content\":\"123456\"}", func() any { return new(FaceBookMessageRequest) }},
		{"FaceBookMessageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FaceBookMessageResponse) }},
		{"FacebookReelsActiveRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"videoNumber\":10,\"searchKeywords\":[\"hello\"]}", func() any { return new(FacebookReelsActiveRequest) }},
		{"FacebookReelsActiveResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FacebookReelsActiveResponse) }},
		{"InstagramPubReelsRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"description\":\"description\",\"video\":[\"https://material.geelark.com/a.mp4\"]}", func() any { return new(InstagramPubReelsRequest) }},
		{"InstagramPubReelsResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramPubReelsResponse) }},
		{"InstagramWarmupRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"browseVideo\":1}", func() any { return new(InstagramWarmupRequest) }},
		{"InstagramWarmupResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramWarmupResponse) }},
		{"InstagramMessageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"usernames\":[\"user\"],\"content\":\"123456\"}", func() any { return new(InstagramMessageRequest) }},
		{"InstagramMessageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramMessageResponse) }},
		{"InstagramPubReelsImagesRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"description\":\"description\",\"image\":[\"https://material.geelark.com/a.jpg\"]}", func() any { return new(InstagramPubReelsImagesRequest) }},
		{"InstagramPubReelsImagesResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramPubReelsImagesResponse) }},
		{"InstagramLoginRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"account\":\"test@gmail.com\",\"password\":\"123456\"}", func() any { return new(InstagramLoginRequest) }},
		{"InstagramLoginResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramLoginResponse) }},
		{"InstagramEditRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"nickname\":\"myName\",\"username\":\"myName\"}", func() any { return new(InstagramEditRequest) }},
		{"InstagramEditResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramEditResponse) }},
		{"InstagramAccountStatusRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"public\":true}", func() any { return new(InstagramAccountStatusRequest) }},
		{"InstagramAccountStatusResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramAccountStatusResponse) }},
		{"InstagramFollowAccountRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"username\":[\"ins1\",\"ins2\"]}", func() any { return new(InstagramFollowAccountRequest) }},
		{"InstagramFollowAccountResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramFollowAccountResponse) }},
		{"InstagramAiCommentRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"useAi\":true,\"randomRate\":50}", func() any { return new(InstagramAiCommentRequest) }},
		{"InstagramAiCommentResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramAiCommentResponse) }},
		{"YouTubePubShortResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(YouTubePubShortResponse) }},
		{"YouTubePubVideoRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"description\":\"description\",\"video\":\"https://material.geelark.com/a.mp4\",\"waitTimeAfterPublishing\":5}", func() any { return new(YouTubePubVideoRequest) }},
		{"YouTubePubVideoResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(YouTubePubVideoResponse) }},
		{"YouTubeActiveAccountRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"browseVideoNum\":10,\"keyword\":[\"k1\",\"k2\"]}", func() any { return new(YouTubeActiveAccountRequest) }},
		{"YouTubeActiveAccountResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(YouTubeActiveAccountResponse) }},
		{"YouTubeEditRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"profileName\":\"myName\"}", func() any { return new(YouTubeEditRequest) }},
		{"YouTubeEditResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(YouTubeEditResponse) }},
		{"GoogleLoginRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"email\":\"test@gmail.com\",\"password\":\"123456\"}", func() any { return new(GoogleLoginRequest) }},
		{"GoogleLoginResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(GoogleLoginResponse) }},
		{"GoogleAppDownloadRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"appName\":\"TikTok\"}", func() any { return new(GoogleAppDownloadRequest) }},
		{"GoogleAppDownloadResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(GoogleAppDownloadResponse) }},
		{"GoogleAppBrowserRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"appName\":\"TikTok\",\"description\":\"good app\"}", func() any { return new(GoogleAppBrowserRequest) }},
		{"GoogleAppBrowserResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(GoogleAppBrowserResponse) }},
		{"SheinLoginRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"email\":\"test@gmail.com\",\"password\":\"123456\"}", func() any { return new(SheinLoginRequest) }},
		{"SheinLoginResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(SheinLoginResponse) }},
		{"XPublishRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"description\":\"description\",\"video\":[\"https://material.geelark.com/a.mp4\"]}", func() any { return new(XPublishRequest) }},
		{"XPublishResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(XPublishResponse) }},
		{"ListFlowsRequest", "{\"page\":1,\"pageSize\":1}", func() any { return new(ListFlowsRequest) }},
		{"ListFlowsResponse", "{\"total\":1,\"page\":1,\"pageSize\":1,\"items\":[{\"id\":\"562316072435344885\",\"title\":\"video flow\",\"desc\":\"this is a video flow\",\"params\":[\"Title\",\"Desc\",\"Video\"]}]}", func() any { return new(ListFlowsResponse) }},
		{"CreateCustomTaskRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"flowId\":\"562316072435344885\",\"paramMap\":{\"Title\":\"video\",\"Desc\":\"this is video\",\"Video\":[\"https://material.geelark.com/a.mp4\"]}}", func() any { return new(CreateCustomTaskRequest) }},
		{"CreateCustomTaskResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(CreateCustomTaskResponse) }},
		{"ImportFlowRequest", "{\"id\":\"612345671223083526\",\"gal\":\"{\\\"content\\\":{\\\"contents\\\":[{\\\"config\\\":{\\\"packgename\\\":\\\"com.zhiliaoapp.musically\\\",\\\"remark\\\":\\\"\\\",\\\"timeout\\\":30000},\\\"type\\\":\\\"openApp\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":10000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"Home\\\",\\\"type\\\":\\\"text\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":3000,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\",\\\"variable\\\":\\\"\\\"},\\\"type\\\":\\\"click\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"For You\\\",\\\"type\\\":\\\"text\\\"},{\\\"content\\\":\\\"android:id/text1\\\",\\\"type\\\":\\\"id\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":30,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\",\\\"variable\\\":\\\"\\\"},\\\"type\\\":\\\"click\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"direction\\\":\\\"top\\\",\\\"distanceMax\\\":700,\\\"distanceMin\\\":500,\\\"position\\\":[300,700],\\\"randomWheelSleepTime\\\":[300,500],\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"scrollPage\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"direction\\\":\\\"top\\\",\\\"distanceMax\\\":700,\\\"distanceMin\\\":500,\\\"position\\\":[300,700],\\\"randomWheelSleepTime\\\":[300,500],\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"scrollPage\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"children\\\":[{\\\"config\\\":{\\\"direction\\\":\\\"top\\\",\\\"distanceMax\\\":600,\\\"distanceMin\\\":500,\\\"position\\\":[300,700],\\\"randomWheelSleepTime\\\":[300,500],\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"scrollPage\\\"},{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"com.zhiliaoapp.musically:id/nl8\\\",\\\"type\\\":\\\"id\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":30,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\",\\\"variable\\\":\\\"avatar\\\"},\\\"type\\\":\\\"waitEle\\\"},{\\\"config\\\":{\\\"children\\\":[{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":30000,\\\"timeoutMin\\\":10000,\\\"timeoutType\\\":\\\"randomInterval\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"children\\\":[{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"com.zhiliaoapp.musically:id/cf6\\\",\\\"type\\\":\\\"id\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":3000,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"click\\\"}],\\\"hiddenChildren\\\":false,\\\"other\\\":[],\\\"probability\\\":30,\\\"relation\\\":\\\"random\\\",\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"ifElse\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":30000,\\\"timeoutMin\\\":10000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"}],\\\"condition\\\":[\\\"avatar\\\"],\\\"hiddenChildren\\\":false,\\\"other\\\":[],\\\"relation\\\":\\\"exist\\\",\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"ifElse\\\"}],\\\"hiddenChildren\\\":false,\\\"remark\\\":\\\"\\\",\\\"times\\\":15,\\\"variableIndex\\\":\\\"for_times_index\\\"},\\\"type\\\":\\\"forTimes\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":120000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"}],\\\"errorType\\\":\\\"skip\\\",\\\"isDebug\\\":false,\\\"timeOut\\\":\\\"30\\\",\\\"contentType\\\":\\\"phone\\\"},\\\"desc\\\":\\\"A TikTok Task flow\\\",\\\"title\\\":\\\"TikTok\\\"}\"}", func() any { return new(ImportFlowRequest) }},
		{"ImportFlowResponse", "{\"id\":\"612345671223083526\"}", func() any { return new(ImportFlowResponse) }},
		{"ExportFlowRequest", "{\"id\":\"612345671223083526\"}", func() any { return new(ExportFlowRequest) }},
		{"ExportFlowResponse", "{\"gal\":\"{\\\"content\\\":{\\\"contents\\\":[{\\\"config\\\":{\\\"packgename\\\":\\\"com.zhiliaoapp.musically\\\",\\\"remark\\\":\\\"\\\",\\\"timeout\\\":30000},\\\"type\\\":\\\"openApp\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":10000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"Home\\\",\\\"type\\\":\\\"text\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":3000,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\",\\\"variable\\\":\\\"\\\"},\\\"type\\\":\\\"click\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"For You\\\",\\\"type\\\":\\\"text\\\"},{\\\"content\\\":\\\"android:id/text1\\\",\\\"type\\\":\\\"id\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":30,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\",\\\"variable\\\":\\\"\\\"},\\\"type\\\":\\\"click\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"direction\\\":\\\"top\\\",\\\"distanceMax\\\":700,\\\"distanceMin\\\":500,\\\"position\\\":[300,700],\\\"randomWheelSleepTime\\\":[300,500],\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"scrollPage\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"direction\\\":\\\"top\\\",\\\"distanceMax\\\":700,\\\"distanceMin\\\":500,\\\"position\\\":[300,700],\\\"randomWheelSleepTime\\\":[300,500],\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"scrollPage\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"children\\\":[{\\\"config\\\":{\\\"direction\\\":\\\"top\\\",\\\"distanceMax\\\":600,\\\"distanceMin\\\":500,\\\"position\\\":[300,700],\\\"randomWheelSleepTime\\\":[300,500],\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"scrollPage\\\"},{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"com.zhiliaoapp.musically:id/nl8\\\",\\\"type\\\":\\\"id\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":30,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\",\\\"variable\\\":\\\"avatar\\\"},\\\"type\\\":\\\"waitEle\\\"},{\\\"config\\\":{\\\"children\\\":[{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":30000,\\\"timeoutMin\\\":10000,\\\"timeoutType\\\":\\\"randomInterval\\\"},\\\"type\\\":\\\"waitTime\\\"},{\\\"config\\\":{\\\"children\\\":[{\\\"config\\\":{\\\"filters\\\":[{\\\"content\\\":\\\"com.zhiliaoapp.musically:id/cf6\\\",\\\"type\\\":\\\"id\\\"}],\\\"remark\\\":\\\"\\\",\\\"searchTime\\\":3000,\\\"serial\\\":1,\\\"serialMax\\\":50,\\\"serialMin\\\":1,\\\"serialType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"click\\\"}],\\\"hiddenChildren\\\":false,\\\"other\\\":[],\\\"probability\\\":30,\\\"relation\\\":\\\"random\\\",\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"ifElse\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":2000,\\\"timeoutMax\\\":30000,\\\"timeoutMin\\\":10000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"}],\\\"condition\\\":[\\\"avatar\\\"],\\\"hiddenChildren\\\":false,\\\"other\\\":[],\\\"relation\\\":\\\"exist\\\",\\\"remark\\\":\\\"\\\"},\\\"type\\\":\\\"ifElse\\\"}],\\\"hiddenChildren\\\":false,\\\"remark\\\":\\\"\\\",\\\"times\\\":15,\\\"variableIndex\\\":\\\"for_times_index\\\"},\\\"type\\\":\\\"forTimes\\\"},{\\\"config\\\":{\\\"remark\\\":\\\"\\\",\\\"timeout\\\":120000,\\\"timeoutMax\\\":300000,\\\"timeoutMin\\\":1000,\\\"timeoutType\\\":\\\"fixedValue\\\"},\\\"type\\\":\\\"waitTime\\\"}],\\\"errorType\\\":\\\"skip\\\",\\\"isDebug\\\":false,\\\"timeOut\\\":\\\"30\\\",\\\"contentType\\\":\\\"phone\\\"},\\\"desc\\\":\\\"A TikTok Task flow\\\",\\\"title\\\":\\\"TikTok\\\"}\"}", func() any { return new(ExportFlowResponse) }},
		{"MultiPlatformVideoDistributionRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"video\":[\"https://singapore-upgrade.geelark.com/a.mp4\"]}", func() any { return new(MultiPlatformVideoDistributionRequest) }},
		{"MultiPlatformVideoDistributionResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(MultiPlatformVideoDistributionResponse) }},
		{"FileUploadRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"files\":[\"https://material.geelark.com/a.mp4\"]}", func() any { return new(FileUploadRequest) }},
		{"FileUploadResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FileUploadResponse) }},
		{"ImportContactsRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"contacts\":[{\"firstName\":\"jay\",\"mobile\":\"13288888888\"}]}", func() any { return new(ImportContactsRequest) }},
		{"ImportContactsResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(ImportContactsResponse) }},
		{"KeyboxUploadRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"files\":[\"https://material.geelark.com/a.xml\"]}", func() any { return new(KeyboxUploadRequest) }},
		{"KeyboxUploadResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(KeyboxUploadResponse) }},
		{"RedditWarmupResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(RedditWarmupResponse) }},
		{"RedditVideoRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"description\":\"description\",\"video\":[\"https://material.geelark.com/a.mp4\"],\"community\":\"cat\"}", func() any { return new(RedditVideoRequest) }},
		{"RedditVideoResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(RedditVideoResponse) }},
		{"RedditImageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"description\":\"description\",\"images\":[\"https://material.geelark.com/a.jpg\"],\"community\":\"cat\"}", func() any { return new(RedditImageRequest) }},
		{"RedditImageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(RedditImageResponse) }},
		{"PinterestVideoRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"description\":\"description\",\"video\":[\"https://material.geelark.com/a.mp4\"]}", func() any { return new(PinterestVideoRequest) }},
		{"PinterestVideoResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(PinterestVideoResponse) }},
		{"PinterestImageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"description\":\"description\",\"images\":[\"https://material.geelark.com/a.jpg\"]}", func() any { return new(PinterestImageRequest) }},
		{"PinterestImageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(PinterestImageResponse) }},
		{"ThreadsVideoRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"video\":[\"https://material.geelark.com/a.mp4\"]}", func() any { return new(ThreadsVideoRequest) }},
		{"ThreadsVideoResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(ThreadsVideoResponse) }},
		{"ThreadsImageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"id\":\"557536075321468390\",\"title\":\"title\",\"images\":[\"https://material.geelark.com/a.jpg\"]}", func() any { return new(ThreadsImageRequest) }},
		{"ThreadsImageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(ThreadsImageResponse) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			value := tc.newValue()
			if err := json.Unmarshal([]byte(tc.raw), value); err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var a, b any
			if err := json.Unmarshal([]byte(tc.raw), &a); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(got, &b); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(a, b) {
				t.Errorf("roundtrip mismatch: want %s, got %s", tc.raw, got)
			}
		})
	}
}
