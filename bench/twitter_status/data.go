// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// twitter_status root type, with the root renamed to Log for the lightning harness.
package bench

type (
	Log struct {
		Statuses       []twitterStatus `json:"statuses"`
		SearchMetadata struct {
			CompletedIn float64 `json:"completed_in"`
			MaxID       int64   `json:"max_id"`
			MaxIDStr    string  `json:"max_id_str"`
			NextResults string  `json:"next_results"`
			Query       string  `json:"query"`
			RefreshURL  string  `json:"refresh_url"`
			Count       int     `json:"count"`
			SinceID     int     `json:"since_id"`
			SinceIDStr  string  `json:"since_id_str"`
		} `json:"search_metadata"`
	}
	twitterStatus struct {
		Metadata struct {
			ResultType      string `json:"result_type"`
			IsoLanguageCode string `json:"iso_language_code"`
		} `json:"metadata"`
		CreatedAt            string          `json:"created_at"`
		ID                   int64           `json:"id"`
		IDStr                string          `json:"id_str"`
		Text                 string          `json:"text"`
		Source               string          `json:"source"`
		Truncated            bool            `json:"truncated"`
		InReplyToStatusID    int64           `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string          `json:"in_reply_to_status_id_str"`
		InReplyToUserID      int64           `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   string          `json:"in_reply_to_user_id_str"`
		InReplyToScreenName  string          `json:"in_reply_to_screen_name"`
		User                 twitterUser     `json:"user,omitempty"`
		Geo                  any             `json:"geo"`
		Coordinates          any             `json:"coordinates"`
		Place                any             `json:"place"`
		Contributors         any             `json:"contributors"`
		RetweeetedStatus     *twitterStatus  `json:"retweeted_status"`
		RetweetCount         int             `json:"retweet_count"`
		FavoriteCount        int             `json:"favorite_count"`
		Entities             twitterEntities `json:"entities,omitempty"`
		Favorited            bool            `json:"favorited"`
		Retweeted            bool            `json:"retweeted"`
		PossiblySensitive    bool            `json:"possibly_sensitive"`
		Lang                 string          `json:"lang"`
	}
	twitterUser struct {
		ID                             int64           `json:"id"`
		IDStr                          string          `json:"id_str"`
		Name                           string          `json:"name"`
		ScreenName                     string          `json:"screen_name"`
		Location                       string          `json:"location"`
		Description                    string          `json:"description"`
		URL                            any             `json:"url"`
		Entities                       twitterEntities `json:"entities"`
		Protected                      bool            `json:"protected"`
		FollowersCount                 int             `json:"followers_count"`
		FriendsCount                   int             `json:"friends_count"`
		ListedCount                    int             `json:"listed_count"`
		CreatedAt                      string          `json:"created_at"`
		FavouritesCount                int             `json:"favourites_count"`
		UtcOffset                      int             `json:"utc_offset"`
		TimeZone                       string          `json:"time_zone"`
		GeoEnabled                     bool            `json:"geo_enabled"`
		Verified                       bool            `json:"verified"`
		StatusesCount                  int             `json:"statuses_count"`
		Lang                           string          `json:"lang"`
		ContributorsEnabled            bool            `json:"contributors_enabled"`
		IsTranslator                   bool            `json:"is_translator"`
		IsTranslationEnabled           bool            `json:"is_translation_enabled"`
		ProfileBackgroundColor         string          `json:"profile_background_color"`
		ProfileBackgroundImageURL      string          `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string          `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool            `json:"profile_background_tile"`
		ProfileImageURL                string          `json:"profile_image_url"`
		ProfileImageURLHTTPS           string          `json:"profile_image_url_https"`
		ProfileBannerURL               string          `json:"profile_banner_url"`
		ProfileLinkColor               string          `json:"profile_link_color"`
		ProfileSidebarBorderColor      string          `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string          `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string          `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool            `json:"profile_use_background_image"`
		DefaultProfile                 bool            `json:"default_profile"`
		DefaultProfileImage            bool            `json:"default_profile_image"`
		Following                      bool            `json:"following"`
		FollowRequestSent              bool            `json:"follow_request_sent"`
		Notifications                  bool            `json:"notifications"`
	}
	twitterEntities struct {
		Hashtags     []any        `json:"hashtags"`
		Symbols      []any        `json:"symbols"`
		URL          *twitterURL  `json:"url"`
		URLs         []twitterURL `json:"urls"`
		UserMentions []struct {
			ScreenName string `json:"screen_name"`
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			IDStr      string `json:"id_str"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
		Description struct {
			URLs []twitterURL `json:"urls"`
		} `json:"description"`
		Media []struct {
			ID            int64  `json:"id"`
			IDStr         string `json:"id_str"`
			Indices       []int  `json:"indices"`
			MediaURL      string `json:"media_url"`
			MediaURLHTTPS string `json:"media_url_https"`
			URL           string `json:"url"`
			DisplayURL    string `json:"display_url"`
			ExpandedURL   string `json:"expanded_url"`
			Type          string `json:"type"`
			Sizes         map[string]struct {
				W      int    `json:"w"`
				H      int    `json:"h"`
				Resize string `json:"resize"`
			} `json:"sizes"`
			SourceStatusID    int64  `json:"source_status_id"`
			SourceStatusIDStr string `json:"source_status_id_str"`
		} `json:"media"`
	}
	twitterURL struct {
		URL         string       `json:"url"`
		URLs        []twitterURL `json:"urls"`
		ExpandedURL string       `json:"expanded_url"`
		DisplayURL  string       `json:"display_url"`
		Indices     []int        `json:"indices"`
	}
)
