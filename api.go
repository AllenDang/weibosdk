package weibosdk

import (
  "encoding/json"
  "fmt"
  "net/http"
  "net/url"
)

func GetAuthorizationUrl(appKey, redirectUrl string) string {
  v := url.Values{}
  v.Add("client_id", appKey)
  v.Add("response_type", "code")
  v.Add("redirect_uri", redirectUrl)

  return UrlOAuth2 + "/authorize?" + v.Encode()
}

func GetAccessToken(appKey, appSecret, redirectUrl, code string) (*AccessToken, error) {
  v := url.Values{}
  v.Add("client_id", appKey)
  v.Add("client_secret", appSecret)
  v.Add("grant_type", "authorization_code")
  v.Add("redirect_uri", redirectUrl)
  v.Add("code", code)

  reqUrl := UrlOAuth2 + "/access_token?" + v.Encode()

  var err error
  var resp *http.Response

  if resp, err = http.Get(reqUrl); err == nil && resp.StatusCode == http.StatusOK {
    defer resp.Body.Close()

    var accessToken AccessToken
    json.NewDecoder(resp.Body).Decode(&accessToken)

    return &accessToken, nil
  }

  return nil, fmt.Errorf("GetAccessToken failed with status code %d", resp.StatusCode)
}
