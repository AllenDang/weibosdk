package weibosdk

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
)

func GetUserShow(accessToken, uid string) (*User, error) {
  v := url.Values{}
  v.Add("access_token", accessToken)
  v.Add("uid", uid)

  reqUrl := UrlAPI2 + "/users/show.json?" + v.Encode()

  var err error

  if resp, err := http.Get(reqUrl); err == nil {
    defer resp.Body.Close()

    if content, err := ioutil.ReadAll(resp.Body); err == nil {
      //先测试返回的是否是ReturnError
      var returnError ReturnError
      if err = json.Unmarshal(content, &returnError); err == nil && returnError.Error_code != 0 {
        return nil, fmt.Errorf("Request %s failed with error_code %d. Error message is '%s'.",
          returnError.Request, returnError.Error_code, returnError.Error)
      }

      var user User
      if err = json.Unmarshal(content, &user); err == nil {
        return &user, nil
      }
    }
  }

  return nil, err
}

func GetEmail(accessToken string) (string, error) {
  v := url.Values{}
  v.Add("access_token", accessToken)

  reqUrl := UrlAPI2 + "/account/profile/email.json?" + v.Encode()

  var err error

  if resp, err := http.Get(reqUrl); err == nil {
    defer resp.Body.Close()

    if content, err := ioutil.ReadAll(resp.Body); err == nil {
      //先测试返回的是否是ReturnError
      var returnError ReturnError
      if err = json.Unmarshal(content, &returnError); err == nil && returnError.Error_code != 0 {
        return "", fmt.Errorf("Request %s failed with error_code %d. Error message is '%s'.",
          returnError.Request, returnError.Error_code, returnError.Error)
      }

      if email, err := extractDataByRegex(string(content),
        `"email":"(.*?)"`); err == nil {
        return email, nil
      }
    }
  }

  return "", err
}
