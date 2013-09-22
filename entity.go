package weibosdk

type AccessToken struct {
  Access_token string `json:access_token`
  Remind_in    string `json:remind_in`
  Expires_in   int    `json:expires_in`
  Uid          string `json:uid`
}

type ReturnError struct {
  Error      string `json:error`
  Error_code int    `json:error_code`
  Request    string `json:request`
}

type Status struct {
  Created_at              string
  Id                      int64
  Mid                     string
  Idstr                   string
  Text                    string
  Source                  string
  Favorited               bool
  Truncated               bool
  In_reply_to_status_id   string
  In_reply_to_user_id     string
  In_reply_to_screen_name string
  Thumbnail_pic           string
  Bmiddle_pic             string
  Original_pic            string
  Reposts_count           int
  Comments_count          int
  Attitudes_count         int
}

type User struct {
  Id                 int64
  Idstr              string
  Class              int
  Screen_name        string
  Name               string
  Province           string
  City               string
  Location           string
  Description        string
  Url                string
  Profile_image_url  string
  Profile_url        string
  Domain             string
  Weihao             string
  Gender             string
  Followers_count    int
  Friends_count      int
  Statuses_count     int
  Favourites_count   int
  Created_at         string
  Following          bool
  Allow_all_act_msg  bool
  Geo_enabled        bool
  Verified           bool
  Verified_type      int
  Remark             string
  Status             Status
  Ptype              int
  Allow_all_comment  bool
  Avatar_large       string
  Avatar_hd          string
  Verified_reason    string
  Follow_me          bool
  Online_status      int
  Bi_followers_count int
  Lang               string
  Star               int
  Mbtype             int
  Mbrank             int
  Block_word         int
}
