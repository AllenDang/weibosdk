package weibosdk

import (
  "net/url"
)

func Update(accessToken, status string) error {
  v := url.Values{}
  v.Add("access_token", accessToken)
  v.Add("status", status)

  reqUrl := UrlAPI2 + "/statuses/update.json"

  var err error

  if _, err := weiboPost(reqUrl, v); err == nil {
    return nil
  }

  return err
}
