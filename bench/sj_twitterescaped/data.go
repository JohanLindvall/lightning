// Code generated from twitterescaped.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	SearchMetadata struct {
		CompletedIn float64 `json:"completed_in"`
		Count       int64   `json:"count"`
		MaxId       int64   `json:"max_id"`
		MaxIdStr    string  `json:"max_id_str"`
		NextResults string  `json:"next_results"`
		Query       string  `json:"query"`
		RefreshUrl  string  `json:"refresh_url"`
		SinceId     int64   `json:"since_id"`
		SinceIdStr  string  `json:"since_id_str"`
	} `json:"search_metadata"`
	Statuses []struct {
		Contributors any    `json:"contributors"`
		Coordinates  any    `json:"coordinates"`
		CreatedAt    string `json:"created_at"`
		Entities     struct {
			Hashtags     []any `json:"hashtags"`
			Symbols      []any `json:"symbols"`
			Urls         []any `json:"urls"`
			UserMentions []any `json:"user_mentions"`
			Media        []struct {
				DisplayUrl    string  `json:"display_url"`
				ExpandedUrl   string  `json:"expanded_url"`
				ID            int64   `json:"id"`
				IdStr         string  `json:"id_str"`
				Indices       []int64 `json:"indices"`
				MediaUrl      string  `json:"media_url"`
				MediaUrlHttps string  `json:"media_url_https"`
				Sizes         struct {
					Large struct {
						H      int64  `json:"h"`
						Resize string `json:"resize"`
						W      int64  `json:"w"`
					} `json:"large"`
					Medium struct {
						H      int64  `json:"h"`
						Resize string `json:"resize"`
						W      int64  `json:"w"`
					} `json:"medium"`
					Small struct {
						H      int64  `json:"h"`
						Resize string `json:"resize"`
						W      int64  `json:"w"`
					} `json:"small"`
					Thumb struct {
						H      int64  `json:"h"`
						Resize string `json:"resize"`
						W      int64  `json:"w"`
					} `json:"thumb"`
				} `json:"sizes"`
				SourceStatusId    int64  `json:"source_status_id"`
				SourceStatusIdStr string `json:"source_status_id_str"`
				Type              string `json:"type"`
				URL               string `json:"url"`
			} `json:"media"`
		} `json:"entities"`
		FavoriteCount        int64  `json:"favorite_count"`
		Favorited            bool   `json:"favorited"`
		Geo                  any    `json:"geo"`
		ID                   int64  `json:"id"`
		IdStr                string `json:"id_str"`
		InReplyToScreenName  any    `json:"in_reply_to_screen_name"`
		InReplyToStatusId    any    `json:"in_reply_to_status_id"`
		InReplyToStatusIdStr any    `json:"in_reply_to_status_id_str"`
		InReplyToUserId      any    `json:"in_reply_to_user_id"`
		InReplyToUserIdStr   any    `json:"in_reply_to_user_id_str"`
		Lang                 string `json:"lang"`
		Metadata             struct {
			IsoLanguageCode string `json:"iso_language_code"`
			ResultType      string `json:"result_type"`
		} `json:"metadata"`
		Place        any    `json:"place"`
		RetweetCount int64  `json:"retweet_count"`
		Retweeted    bool   `json:"retweeted"`
		Source       string `json:"source"`
		Text         string `json:"text"`
		Truncated    bool   `json:"truncated"`
		User         struct {
			ContributorsEnabled bool   `json:"contributors_enabled"`
			CreatedAt           string `json:"created_at"`
			DefaultProfile      bool   `json:"default_profile"`
			DefaultProfileImage bool   `json:"default_profile_image"`
			Description         string `json:"description"`
			Entities            struct {
				Description struct {
					Urls []any `json:"urls"`
				} `json:"description"`
				URL struct {
					Urls []struct {
						DisplayUrl  string  `json:"display_url"`
						ExpandedUrl string  `json:"expanded_url"`
						Indices     []int64 `json:"indices"`
						URL         string  `json:"url"`
					} `json:"urls"`
				} `json:"url"`
			} `json:"entities"`
			FavouritesCount                int64  `json:"favourites_count"`
			FollowRequestSent              bool   `json:"follow_request_sent"`
			FollowersCount                 int64  `json:"followers_count"`
			Following                      bool   `json:"following"`
			FriendsCount                   int64  `json:"friends_count"`
			GeoEnabled                     bool   `json:"geo_enabled"`
			ID                             int64  `json:"id"`
			IdStr                          string `json:"id_str"`
			IsTranslationEnabled           bool   `json:"is_translation_enabled"`
			IsTranslator                   bool   `json:"is_translator"`
			Lang                           string `json:"lang"`
			ListedCount                    int64  `json:"listed_count"`
			Location                       string `json:"location"`
			Name                           string `json:"name"`
			Notifications                  bool   `json:"notifications"`
			ProfileBackgroundColor         string `json:"profile_background_color"`
			ProfileBackgroundImageUrl      string `json:"profile_background_image_url"`
			ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool   `json:"profile_background_tile"`
			ProfileBannerUrl               string `json:"profile_banner_url"`
			ProfileImageUrl                string `json:"profile_image_url"`
			ProfileImageUrlHttps           string `json:"profile_image_url_https"`
			ProfileLinkColor               string `json:"profile_link_color"`
			ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
			Protected                      bool   `json:"protected"`
			ScreenName                     string `json:"screen_name"`
			StatusesCount                  int64  `json:"statuses_count"`
			TimeZone                       any    `json:"time_zone"`
			URL                            any    `json:"url"`
			UtcOffset                      any    `json:"utc_offset"`
			Verified                       bool   `json:"verified"`
		} `json:"user"`
		PossiblySensitive bool `json:"possibly_sensitive"`
		RetweetedStatus   struct {
			Contributors any    `json:"contributors"`
			Coordinates  any    `json:"coordinates"`
			CreatedAt    string `json:"created_at"`
			Entities     struct {
				Hashtags []any `json:"hashtags"`
				Media    []struct {
					DisplayUrl    string  `json:"display_url"`
					ExpandedUrl   string  `json:"expanded_url"`
					ID            int64   `json:"id"`
					IdStr         string  `json:"id_str"`
					Indices       []int64 `json:"indices"`
					MediaUrl      string  `json:"media_url"`
					MediaUrlHttps string  `json:"media_url_https"`
					Sizes         struct {
						Large struct {
							H      int64  `json:"h"`
							Resize string `json:"resize"`
							W      int64  `json:"w"`
						} `json:"large"`
						Medium struct {
							H      int64  `json:"h"`
							Resize string `json:"resize"`
							W      int64  `json:"w"`
						} `json:"medium"`
						Small struct {
							H      int64  `json:"h"`
							Resize string `json:"resize"`
							W      int64  `json:"w"`
						} `json:"small"`
						Thumb struct {
							H      int64  `json:"h"`
							Resize string `json:"resize"`
							W      int64  `json:"w"`
						} `json:"thumb"`
					} `json:"sizes"`
					Type              string `json:"type"`
					URL               string `json:"url"`
					SourceStatusId    int64  `json:"source_status_id"`
					SourceStatusIdStr string `json:"source_status_id_str"`
				} `json:"media"`
				Symbols      []any `json:"symbols"`
				Urls         []any `json:"urls"`
				UserMentions []any `json:"user_mentions"`
			} `json:"entities"`
			FavoriteCount        int64  `json:"favorite_count"`
			Favorited            bool   `json:"favorited"`
			Geo                  any    `json:"geo"`
			ID                   int64  `json:"id"`
			IdStr                string `json:"id_str"`
			InReplyToScreenName  any    `json:"in_reply_to_screen_name"`
			InReplyToStatusId    any    `json:"in_reply_to_status_id"`
			InReplyToStatusIdStr any    `json:"in_reply_to_status_id_str"`
			InReplyToUserId      any    `json:"in_reply_to_user_id"`
			InReplyToUserIdStr   any    `json:"in_reply_to_user_id_str"`
			Lang                 string `json:"lang"`
			Metadata             struct {
				IsoLanguageCode string `json:"iso_language_code"`
				ResultType      string `json:"result_type"`
			} `json:"metadata"`
			Place             any    `json:"place"`
			PossiblySensitive bool   `json:"possibly_sensitive"`
			RetweetCount      int64  `json:"retweet_count"`
			Retweeted         bool   `json:"retweeted"`
			Source            string `json:"source"`
			Text              string `json:"text"`
			Truncated         bool   `json:"truncated"`
			User              struct {
				ContributorsEnabled bool   `json:"contributors_enabled"`
				CreatedAt           string `json:"created_at"`
				DefaultProfile      bool   `json:"default_profile"`
				DefaultProfileImage bool   `json:"default_profile_image"`
				Description         string `json:"description"`
				Entities            struct {
					Description struct {
						Urls []any `json:"urls"`
					} `json:"description"`
					URL struct {
						Urls []struct {
							DisplayUrl  string  `json:"display_url"`
							ExpandedUrl string  `json:"expanded_url"`
							Indices     []int64 `json:"indices"`
							URL         string  `json:"url"`
						} `json:"urls"`
					} `json:"url"`
				} `json:"entities"`
				FavouritesCount                int64  `json:"favourites_count"`
				FollowRequestSent              bool   `json:"follow_request_sent"`
				FollowersCount                 int64  `json:"followers_count"`
				Following                      bool   `json:"following"`
				FriendsCount                   int64  `json:"friends_count"`
				GeoEnabled                     bool   `json:"geo_enabled"`
				ID                             int64  `json:"id"`
				IdStr                          string `json:"id_str"`
				IsTranslationEnabled           bool   `json:"is_translation_enabled"`
				IsTranslator                   bool   `json:"is_translator"`
				Lang                           string `json:"lang"`
				ListedCount                    int64  `json:"listed_count"`
				Location                       string `json:"location"`
				Name                           string `json:"name"`
				Notifications                  bool   `json:"notifications"`
				ProfileBackgroundColor         string `json:"profile_background_color"`
				ProfileBackgroundImageUrl      string `json:"profile_background_image_url"`
				ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
				ProfileBackgroundTile          bool   `json:"profile_background_tile"`
				ProfileBannerUrl               string `json:"profile_banner_url"`
				ProfileImageUrl                string `json:"profile_image_url"`
				ProfileImageUrlHttps           string `json:"profile_image_url_https"`
				ProfileLinkColor               string `json:"profile_link_color"`
				ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
				ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
				ProfileTextColor               string `json:"profile_text_color"`
				ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
				Protected                      bool   `json:"protected"`
				ScreenName                     string `json:"screen_name"`
				StatusesCount                  int64  `json:"statuses_count"`
				TimeZone                       any    `json:"time_zone"`
				URL                            any    `json:"url"`
				UtcOffset                      any    `json:"utc_offset"`
				Verified                       bool   `json:"verified"`
			} `json:"user"`
		} `json:"retweeted_status"`
	} `json:"statuses"`
}
