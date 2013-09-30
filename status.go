package weibosdk

import (
  "github.com/golang/glog"
  "net/url"
)

func Update(accessToken, status string) error {
  v := url.Values{}
  v.Add("access_token", accessToken)
  v.Add("status", status)

  reqUrl := UrlAPI2 + "/statuses/update.json"

  var err error

  if resp, err := weiboPost(reqUrl, v); err == nil {
    glog.Infoln(resp)
    return nil
  }

  return err
}
