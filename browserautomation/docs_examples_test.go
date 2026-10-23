package browserautomation

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDocumentationExamples(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		newValue func() any
	}{
		{"SearchRequest", "{\"page\":1,\"pageSize\":10,\"taskIds\":[\"497652752864775437\"]}", func() any { return new(SearchRequest) }},
		{"SearchResponse", "{\"total\":1,\"page\":1,\"pageSize\":10,\"list\":[{\"id\":\"497652752864775437\",\"eid\":\"497652752864775438\",\"name\":\"myTask\",\"remark\":\"\",\"serialName\":\"\",\"status\":3,\"startAt\":1762912197,\"finishAt\":1762912497,\"cost\":300,\"resultCode\":0,\"resultDesc\":\"success\",\"scheduleAt\":1762912197}]}", func() any { return new(SearchResponse) }},
		{"CancelRequest", "{\"taskId\":\"497652752864775437\"}", func() any { return new(CancelRequest) }},
		{"RestartRequest", "{\"taskId\":\"497652752864775437\"}", func() any { return new(RestartRequest) }},
		{"DetailRequest", "{\"taskId\":\"497652752864775437\"}", func() any { return new(DetailRequest) }},
		{"DetailResponse", "{\"id\":\"609019758523737585\",\"eid\":\"600729455865921210\",\"name\":\"task20260303182734\",\"remark\":\"\",\"serialName\":\"\",\"status\":3,\"startAt\":1772533706,\"finishAt\":1772533707,\"cost\":1,\"resultCode\":0,\"resultDesc\":\"success\",\"scheduleAt\":1772533659,\"logs\":[\"[2026-03-03 10:28:26 118] New Tab:\"]}", func() any { return new(DetailResponse) }},
		{"FlowRequest", "{\"page\":1,\"pageSize\":10}", func() any { return new(FlowRequest) }},
		{"FlowResponse", "{\"total\":1,\"page\":1,\"pageSize\":10,\"list\":[{\"id\":\"497652752864775437\",\"title\":\"video flow\",\"desc\":\"this is a video flow\",\"params\":[\"Title\",\"Desc\"]}]}", func() any { return new(FlowResponse) }},
		{"AddRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"flowId\":\"562316072435344885\",\"paramMap\":{\"Title\":\"video\",\"Desc\":\"this is video\",\"Video\":[\"https://material.geelark.cn/a.mp4\"]}}", func() any { return new(AddRequest) }},
		{"AddResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(AddResponse) }},
		{"TikTokSearchRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"searchKeyword\":\"hello\"}", func() any { return new(TikTokSearchRequest) }},
		{"TikTokSearchResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokSearchResponse) }},
		{"TikTokCommentRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"comments\":[\"hello\"]}", func() any { return new(TikTokCommentRequest) }},
		{"TikTokCommentResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokCommentResponse) }},
		{"TikTokLikeRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"videoLink\":\"https://www.tiktok.com/video/38210380122\",\"comment\":\"hello\"}", func() any { return new(TikTokLikeRequest) }},
		{"TikTokLikeResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(TikTokLikeResponse) }},
		{"FacebookPostRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"content\":\"hello\"}", func() any { return new(FacebookPostRequest) }},
		{"FacebookPostResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FacebookPostResponse) }},
		{"FacebookFriendsRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\"}", func() any { return new(FacebookFriendsRequest) }},
		{"FacebookFriendsResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FacebookFriendsResponse) }},
		{"FacebookLikeRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\"}", func() any { return new(FacebookLikeRequest) }},
		{"FacebookLikeResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FacebookLikeResponse) }},
		{"FacebookHomepageRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"mainPageName\":\"myPage\",\"category\":[\"dev\"]}", func() any { return new(FacebookHomepageRequest) }},
		{"FacebookHomepageResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(FacebookHomepageResponse) }},
		{"InstagramLikeRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\"}", func() any { return new(InstagramLikeRequest) }},
		{"InstagramLikeResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramLikeResponse) }},
		{"InstagramSearchRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"searchKeywords\":[\"hello\"]}", func() any { return new(InstagramSearchRequest) }},
		{"InstagramSearchResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(InstagramSearchResponse) }},
		{"YouTubeWatchRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"searchKeyword\":\"hello\",\"title\":\"myTitle\",\"comment\":\"myComment\"}", func() any { return new(YouTubeWatchRequest) }},
		{"YouTubeWatchResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(YouTubeWatchResponse) }},
		{"XPostRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"content\":\"hello\"}", func() any { return new(XPostRequest) }},
		{"XPostResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(XPostResponse) }},
		{"XLikeRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\"}", func() any { return new(XLikeRequest) }},
		{"XLikeResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(XLikeResponse) }},
		{"RedditLikeRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"searchKeywords\":[\"hello\"]}", func() any { return new(RedditLikeRequest) }},
		{"RedditLikeResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(RedditLikeResponse) }},
		{"CookieBotRequest", "{\"name\":\"test\",\"remark\":\"test remark\",\"scheduleAt\":1741846843,\"eid\":\"557536075321468390\",\"pages\":[\"https://a.com\"]}", func() any { return new(CookieBotRequest) }},
		{"CookieBotResponse", "{\"taskId\":\"558017255909123564\"}", func() any { return new(CookieBotResponse) }},
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
