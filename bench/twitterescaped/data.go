// Code generated from twitterescaped.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	SearchMetadata struct {
		CompletedIn float64 `json:"completed_in"`
		Count       int64   `json:"count"`
		MaxId       int64   `json:"max_id"`
		MaxIdStr    string  `json:"max_id_str,nocopy"`
		NextResults string  `json:"next_results,nocopy"`
		Query       string  `json:"query,nocopy"`
		RefreshUrl  string  `json:"refresh_url,nocopy"`
		SinceId     int64   `json:"since_id"`
		SinceIdStr  string  `json:"since_id_str,nocopy"`
	} `json:"search_metadata"`
	Statuses []struct {
		Contributors any    `json:"contributors"`
		Coordinates  any    `json:"coordinates"`
		CreatedAt    string `json:"created_at,nocopy"`
		Entities     struct {
			Hashtags []struct {
				Indices []int64 `json:"indices"`
				Text    string  `json:"text,nocopy"`
			} `json:"hashtags"`
			Symbols []any `json:"symbols"`
			Urls    []struct {
				DisplayUrl  string  `json:"display_url,nocopy"`
				ExpandedUrl string  `json:"expanded_url,nocopy"`
				Indices     []int64 `json:"indices"`
				URL         string  `json:"url,nocopy"`
			} `json:"urls"`
			UserMentions []struct {
				ID         int64   `json:"id"`
				IdStr      string  `json:"id_str,nocopy"`
				Indices    []int64 `json:"indices"`
				Name       string  `json:"name,nocopy"`
				ScreenName string  `json:"screen_name,nocopy"`
			} `json:"user_mentions"`
			Media []struct {
				DisplayUrl    string  `json:"display_url,nocopy"`
				ExpandedUrl   string  `json:"expanded_url,nocopy"`
				ID            int64   `json:"id"`
				IdStr         string  `json:"id_str,nocopy"`
				Indices       []int64 `json:"indices"`
				MediaUrl      string  `json:"media_url,nocopy"`
				MediaUrlHttps string  `json:"media_url_https,nocopy"`
				Sizes         struct {
					Large struct {
						H      int64  `json:"h"`
						Resize string `json:"resize,nocopy"`
						W      int64  `json:"w"`
					} `json:"large"`
					Medium struct {
						H      int64  `json:"h"`
						Resize string `json:"resize,nocopy"`
						W      int64  `json:"w"`
					} `json:"medium"`
					Small struct {
						H      int64  `json:"h"`
						Resize string `json:"resize,nocopy"`
						W      int64  `json:"w"`
					} `json:"small"`
					Thumb struct {
						H      int64  `json:"h"`
						Resize string `json:"resize,nocopy"`
						W      int64  `json:"w"`
					} `json:"thumb"`
				} `json:"sizes"`
				SourceStatusId    int64  `json:"source_status_id"`
				SourceStatusIdStr string `json:"source_status_id_str,nocopy"`
				Type              string `json:"type,nocopy"`
				URL               string `json:"url,nocopy"`
			} `json:"media"`
		} `json:"entities"`
		FavoriteCount        int64  `json:"favorite_count"`
		Favorited            bool   `json:"favorited"`
		Geo                  any    `json:"geo"`
		ID                   int64  `json:"id"`
		IdStr                string `json:"id_str,nocopy"`
		InReplyToScreenName  string `json:"in_reply_to_screen_name,nocopy"`
		InReplyToStatusId    int64  `json:"in_reply_to_status_id"`
		InReplyToStatusIdStr string `json:"in_reply_to_status_id_str,nocopy"`
		InReplyToUserId      int64  `json:"in_reply_to_user_id"`
		InReplyToUserIdStr   string `json:"in_reply_to_user_id_str,nocopy"`
		Lang                 string `json:"lang,nocopy"`
		Metadata             struct {
			IsoLanguageCode string `json:"iso_language_code,nocopy"`
			ResultType      string `json:"result_type,nocopy"`
		} `json:"metadata"`
		Place        any    `json:"place"`
		RetweetCount int64  `json:"retweet_count"`
		Retweeted    bool   `json:"retweeted"`
		Source       string `json:"source,nocopy"`
		Text         string `json:"text,nocopy"`
		Truncated    bool   `json:"truncated"`
		User         struct {
			ContributorsEnabled bool   `json:"contributors_enabled"`
			CreatedAt           string `json:"created_at,nocopy"`
			DefaultProfile      bool   `json:"default_profile"`
			DefaultProfileImage bool   `json:"default_profile_image"`
			Description         string `json:"description,nocopy"`
			Entities            struct {
				Description struct {
					Urls []any `json:"urls"`
				} `json:"description"`
				URL struct {
					Urls []struct {
						DisplayUrl  string  `json:"display_url,nocopy"`
						ExpandedUrl string  `json:"expanded_url,nocopy"`
						Indices     []int64 `json:"indices"`
						URL         string  `json:"url,nocopy"`
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
			IdStr                          string `json:"id_str,nocopy"`
			IsTranslationEnabled           bool   `json:"is_translation_enabled"`
			IsTranslator                   bool   `json:"is_translator"`
			Lang                           string `json:"lang,nocopy"`
			ListedCount                    int64  `json:"listed_count"`
			Location                       string `json:"location,nocopy"`
			Name                           string `json:"name,nocopy"`
			Notifications                  bool   `json:"notifications"`
			ProfileBackgroundColor         string `json:"profile_background_color,nocopy"`
			ProfileBackgroundImageUrl      string `json:"profile_background_image_url,nocopy"`
			ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https,nocopy"`
			ProfileBackgroundTile          bool   `json:"profile_background_tile"`
			ProfileBannerUrl               string `json:"profile_banner_url,nocopy"`
			ProfileImageUrl                string `json:"profile_image_url,nocopy"`
			ProfileImageUrlHttps           string `json:"profile_image_url_https,nocopy"`
			ProfileLinkColor               string `json:"profile_link_color,nocopy"`
			ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color,nocopy"`
			ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color,nocopy"`
			ProfileTextColor               string `json:"profile_text_color,nocopy"`
			ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
			Protected                      bool   `json:"protected"`
			ScreenName                     string `json:"screen_name,nocopy"`
			StatusesCount                  int64  `json:"statuses_count"`
			TimeZone                       string `json:"time_zone,nocopy"`
			URL                            string `json:"url,nocopy"`
			UtcOffset                      int64  `json:"utc_offset"`
			Verified                       bool   `json:"verified"`
		} `json:"user"`
		PossiblySensitive bool `json:"possibly_sensitive"`
		RetweetedStatus   struct {
			Contributors any    `json:"contributors"`
			Coordinates  any    `json:"coordinates"`
			CreatedAt    string `json:"created_at,nocopy"`
			Entities     struct {
				Hashtags []struct {
					Indices []int64 `json:"indices"`
					Text    string  `json:"text,nocopy"`
				} `json:"hashtags"`
				Media []struct {
					DisplayUrl    string  `json:"display_url,nocopy"`
					ExpandedUrl   string  `json:"expanded_url,nocopy"`
					ID            int64   `json:"id"`
					IdStr         string  `json:"id_str,nocopy"`
					Indices       []int64 `json:"indices"`
					MediaUrl      string  `json:"media_url,nocopy"`
					MediaUrlHttps string  `json:"media_url_https,nocopy"`
					Sizes         struct {
						Large struct {
							H      int64  `json:"h"`
							Resize string `json:"resize,nocopy"`
							W      int64  `json:"w"`
						} `json:"large"`
						Medium struct {
							H      int64  `json:"h"`
							Resize string `json:"resize,nocopy"`
							W      int64  `json:"w"`
						} `json:"medium"`
						Small struct {
							H      int64  `json:"h"`
							Resize string `json:"resize,nocopy"`
							W      int64  `json:"w"`
						} `json:"small"`
						Thumb struct {
							H      int64  `json:"h"`
							Resize string `json:"resize,nocopy"`
							W      int64  `json:"w"`
						} `json:"thumb"`
					} `json:"sizes"`
					Type              string `json:"type,nocopy"`
					URL               string `json:"url,nocopy"`
					SourceStatusId    int64  `json:"source_status_id"`
					SourceStatusIdStr string `json:"source_status_id_str,nocopy"`
				} `json:"media"`
				Symbols []any `json:"symbols"`
				Urls    []struct {
					DisplayUrl  string  `json:"display_url,nocopy"`
					ExpandedUrl string  `json:"expanded_url,nocopy"`
					Indices     []int64 `json:"indices"`
					URL         string  `json:"url,nocopy"`
				} `json:"urls"`
				UserMentions []struct {
					ID         int64   `json:"id"`
					IdStr      string  `json:"id_str,nocopy"`
					Indices    []int64 `json:"indices"`
					Name       string  `json:"name,nocopy"`
					ScreenName string  `json:"screen_name,nocopy"`
				} `json:"user_mentions"`
			} `json:"entities"`
			FavoriteCount        int64  `json:"favorite_count"`
			Favorited            bool   `json:"favorited"`
			Geo                  any    `json:"geo"`
			ID                   int64  `json:"id"`
			IdStr                string `json:"id_str,nocopy"`
			InReplyToScreenName  string `json:"in_reply_to_screen_name,nocopy"`
			InReplyToStatusId    int64  `json:"in_reply_to_status_id"`
			InReplyToStatusIdStr string `json:"in_reply_to_status_id_str,nocopy"`
			InReplyToUserId      int64  `json:"in_reply_to_user_id"`
			InReplyToUserIdStr   string `json:"in_reply_to_user_id_str,nocopy"`
			Lang                 string `json:"lang,nocopy"`
			Metadata             struct {
				IsoLanguageCode string `json:"iso_language_code,nocopy"`
				ResultType      string `json:"result_type,nocopy"`
			} `json:"metadata"`
			Place             any    `json:"place"`
			PossiblySensitive bool   `json:"possibly_sensitive"`
			RetweetCount      int64  `json:"retweet_count"`
			Retweeted         bool   `json:"retweeted"`
			Source            string `json:"source,nocopy"`
			Text              string `json:"text,nocopy"`
			Truncated         bool   `json:"truncated"`
			User              struct {
				ContributorsEnabled bool   `json:"contributors_enabled"`
				CreatedAt           string `json:"created_at,nocopy"`
				DefaultProfile      bool   `json:"default_profile"`
				DefaultProfileImage bool   `json:"default_profile_image"`
				Description         string `json:"description,nocopy"`
				Entities            struct {
					Description struct {
						Urls []any `json:"urls"`
					} `json:"description"`
					URL struct {
						Urls []struct {
							DisplayUrl  string  `json:"display_url,nocopy"`
							ExpandedUrl string  `json:"expanded_url,nocopy"`
							Indices     []int64 `json:"indices"`
							URL         string  `json:"url,nocopy"`
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
				IdStr                          string `json:"id_str,nocopy"`
				IsTranslationEnabled           bool   `json:"is_translation_enabled"`
				IsTranslator                   bool   `json:"is_translator"`
				Lang                           string `json:"lang,nocopy"`
				ListedCount                    int64  `json:"listed_count"`
				Location                       string `json:"location,nocopy"`
				Name                           string `json:"name,nocopy"`
				Notifications                  bool   `json:"notifications"`
				ProfileBackgroundColor         string `json:"profile_background_color,nocopy"`
				ProfileBackgroundImageUrl      string `json:"profile_background_image_url,nocopy"`
				ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https,nocopy"`
				ProfileBackgroundTile          bool   `json:"profile_background_tile"`
				ProfileBannerUrl               string `json:"profile_banner_url,nocopy"`
				ProfileImageUrl                string `json:"profile_image_url,nocopy"`
				ProfileImageUrlHttps           string `json:"profile_image_url_https,nocopy"`
				ProfileLinkColor               string `json:"profile_link_color,nocopy"`
				ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color,nocopy"`
				ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color,nocopy"`
				ProfileTextColor               string `json:"profile_text_color,nocopy"`
				ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
				Protected                      bool   `json:"protected"`
				ScreenName                     string `json:"screen_name,nocopy"`
				StatusesCount                  int64  `json:"statuses_count"`
				TimeZone                       string `json:"time_zone,nocopy"`
				URL                            string `json:"url,nocopy"`
				UtcOffset                      int64  `json:"utc_offset"`
				Verified                       bool   `json:"verified"`
			} `json:"user"`
		} `json:"retweeted_status"`
	} `json:"statuses"`
}
