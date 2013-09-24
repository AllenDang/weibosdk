package weibosdk

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "regexp"
  "strings"
)

func extractDataByRegex(content, query string) (string, error) {
  rx := regexp.MustCompile(query)
  value := rx.FindStringSubmatch(content)

  if len(value) == 0 {
    return "", fmt.Errorf("正则表达式没有匹配到内容:(%s)", query)
  }

  return strings.TrimSpace(value[1]), nil
}

func weiboPost(reqUrl string, v url.Values) ([]byte, error) {
  var err error
  if resp, err := http.PostForm(reqUrl, v); err == nil {
    defer resp.Body.Close()

    if content, err := ioutil.ReadAll(resp.Body); err == nil {
      //先测试返回的是否是ReturnError
      var returnError ReturnError
      if err = json.Unmarshal(content, &returnError); err == nil && returnError.Error_code != 0 {
        return nil, fmt.Errorf("Request %s failed with error_code %d. Error message is '%s'.",
          returnError.Request, returnError.Error_code, returnError.Error)
      }

      return content, nil
    }
  }

  return nil, err
}

func weiboGet(reqUrl string) ([]byte, error) {
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

      return content, nil
    }
  }

  return nil, err
}
