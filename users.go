package weibosdk

import (
  "encoding/json"
  "net/url"
)

func GetUserShow(accessToken, uid string) (*User, error) {
  v := url.Values{}
  v.Add("access_token", accessToken)
  v.Add("uid", uid)

  reqUrl := UrlAPI2 + "/users/show.json?" + v.Encode()

  var err error

  if resp, err := weiboGet(reqUrl); err == nil {
    var user User
    if err = json.Unmarshal(resp, &user); err == nil {
      return &user, nil
    }
  }

  return nil, err
}

func GetEmail(accessToken string) (string, error) {
  v := url.Values{}
  v.Add("access_token", accessToken)

  reqUrl := UrlAPI2 + "/account/profile/email.json?" + v.Encode()

  var err error

  if resp, err := weiboGet(reqUrl); err == nil {
    if email, err := extractDataByRegex(string(resp),
      `"email":"(.*?)"`); err == nil {
      return email, nil
    }
  }

  return "", err
}
