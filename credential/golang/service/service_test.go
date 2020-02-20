package service

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

var a *AccessTokenCredential

func Test_Init(t *testing.T) {
	a = new(AccessTokenCredential)
	a.credentialUpdater = new(credentialUpdater)
	err := a.SetExpireTime("2006-01-T15:04:05Z")
	utils.AssertContains(t, err.Error(), `parsing time "2006-01-T15:04:05Z" as "2006-01-02T15:04:05Z07:00"`)

	expiretime := a.GetExpireTime()
	utils.AssertEqual(t, expiretime, "")

	err = a.SetExpireTime("2006-01-02T15:04:05Z")
	utils.AssertNil(t, err)
	expiretime = a.GetExpireTime()
	utils.AssertContains(t, expiretime, "2006-01-02T")

	refreshtoken := a.GetRefreshToken()
	utils.AssertEqual(t, refreshtoken, "")

	a.SetRefreshToken("refreshtoken")
	refreshtoken = a.GetRefreshToken()
	utils.AssertEqual(t, refreshtoken, "refreshtoken")

	a.SetAccessToken("accesstoken")
	utils.AssertEqual(t, *a.AccessToken, "accesstoken")
}

func Test_RefreshTokenWithConcurrent(t *testing.T) {
	originTestHookDo := hookdo
	defer func() { hookdo = originTestHookDo }()
	hookdo = func(resp *tea.Response, err error) (*tea.Response, error) {
		tmp := `{"expire_time": "2006-01-02T15:04:05Z","access_token":"access_token","refresh_token":"refresh_token"}`
		httpresponse := &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(tmp)),
		}
		resp = tea.NewResponse(httpresponse)
		return resp, nil
	}
	originTest := RefreshCallbackFn
	var refre, access, expire string
	defer func() { RefreshCallbackFn = originTest }()
	RefreshCallbackFn = func(refreshToken, accessToken, expireTime string) {
		refre = refreshToken
		access = accessToken
		expire = expireTime
		return
	}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			accesstoken, err := a.GetAccessToken()
			utils.AssertNil(t, err)
			utils.AssertEqual(t, "access_token", accesstoken)
			wg.Done()
		}()
	}
	wg.Wait()
	utils.AssertEqual(t, "refresh_token", refre)
	utils.AssertEqual(t, "access_token", access)
	utils.AssertEqual(t, "2006-01-02T15:04:05Z", expire)
}

func Test_GetAccessToken(t *testing.T) {
	originhook := hookdo
	defer func() {
		hookdo = originhook
	}()

	originTest := RefreshFailCallbackFn
	var err1 error
	defer func() { RefreshFailCallbackFn = originTest }()
	RefreshFailCallbackFn = func(err error) {
		err1 = err
		return
	}

	hookdo = func(resp *tea.Response, err error) (*tea.Response, error) {
		tmp := `{"expire_time": "2006-01-T15:04:05Z"}`
		httpresponse := &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(tmp)),
		}
		resp = tea.NewResponse(httpresponse)
		return resp, nil
	}
	accesstoken, err := a.GetAccessToken()
	utils.AssertEqual(t, err.Error(), `parsing time "2006-01-T15:04:05Z" as "2006-01-02T15:04:05Z07:00": cannot parse "T15:04:05Z" as "02"`)
	utils.AssertEqual(t, accesstoken, "")
	utils.AssertEqual(t, err1.Error(), `parsing time "2006-01-T15:04:05Z" as "2006-01-02T15:04:05Z07:00": cannot parse "T15:04:05Z" as "02"`)

	hookdo = func(resp *tea.Response, err error) (*tea.Response, error) {
		tmp := `{"expire_time": "2006-01-02T15:04:05Z","access_token":"access_token","refresh_token":"refresh_token"}`
		httpresponse := &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(tmp)),
		}
		resp = tea.NewResponse(httpresponse)
		return resp, nil
	}
	accesstoken, err = a.GetAccessToken()
	utils.AssertNil(t, err)
	utils.AssertEqual(t, accesstoken, "access_token")

	a.ExpireTime = nil
	accesstoken, err = a.GetAccessToken()
	utils.AssertNil(t, err)
	utils.AssertEqual(t, accesstoken, "access_token")

	a.AccessToken = nil
	accesstoken, err = a.GetAccessToken()
	utils.AssertNil(t, err)
	utils.AssertEqual(t, accesstoken, "")

	a = nil
	accesstoken, err = a.GetAccessToken()
	utils.AssertNil(t, err)
	utils.AssertEqual(t, accesstoken, "")
}

func Test_refreshAccessToken(t *testing.T) {
	resp, err := hookdo(nil, nil)
	utils.AssertNil(t, resp)
	utils.AssertNil(t, err)

	originHook := hookdo
	defer func() {
		hookdo = originHook
	}()
	hookdo = func(response *tea.Response, err error) (*tea.Response, error) {
		return nil, errors.New("refresh_token error")
	}
	at, rt, up, err := refreshAccessToken("", "ccp", "refresh_token", "client_id", "client_secret")
	utils.AssertEqual(t, "", at)
	utils.AssertEqual(t, "", rt)
	utils.AssertNil(t, up)
	utils.AssertEqual(t, "refresh_token error", err.Error())

	hookdo = func(resp *tea.Response, err error) (*tea.Response, error) {
		tmp := `{"expire_time": 10}`
		httpresponse := &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(tmp)),
		}
		resp = tea.NewResponse(httpresponse)
		return resp, nil
	}
	at, rt, up, err = refreshAccessToken("endpoint", "ccp", "refresh_token", "client_id", "client_secret")
	utils.AssertEqual(t, "", at)
	utils.AssertEqual(t, "", rt)
	utils.AssertNil(t, up)
	utils.AssertEqual(t, "json: cannot unmarshal number into Go struct field accessTokenResponse.expire_time of type string", err.Error())

	hookdo = func(resp *tea.Response, err error) (*tea.Response, error) {
		tmp := `{"expire_time": "2006-01-02T15:04:05Z","access_token":"access_token","refresh_token":"refresh_token"}`
		httpresponse := &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(tmp)),
		}
		resp = tea.NewResponse(httpresponse)
		return resp, nil
	}
	at, rt, up, err = refreshAccessToken("endpoint", "ccp", "refresh_token", "client_id", "client_secret")
	utils.AssertEqual(t, "access_token", at)
	utils.AssertEqual(t, "refresh_token", rt)
	utils.AssertNil(t, err)
	utils.AssertNotNil(t, up)

	ifUpdate := up.needUpdateCredential()
	utils.AssertEqual(t, true, ifUpdate)

	hookdo = func(resp *tea.Response, err error) (*tea.Response, error) {
		tmp := `{"expire_time": "2006-01-02T15:04:05Z","access_token":"access_token","refresh_token":"refresh_token"}`
		httpresponse := &http.Response{
			StatusCode: 400,
			Body:       ioutil.NopCloser(strings.NewReader(tmp)),
		}
		resp = tea.NewResponse(httpresponse)
		return resp, nil
	}
	at, rt, up, err = refreshAccessToken("endpoint", "ccp", "refresh_token", "client_id", "client_secret")
	utils.AssertEqual(t, "", at)
	utils.AssertEqual(t, "", rt)
	utils.AssertNotNil(t, err)
	utils.AssertNil(t, up)
}