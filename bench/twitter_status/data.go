// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// twitter_status root type, with the root renamed to Benchmark for the lightning harness.
package bench

type (
	Benchmark struct {
		Statuses       []twitterStatus `json:"statuses"`
		SearchMetadata struct {
			CompletedIn float64 `json:"completed_in"`
			MaxID       int64   `json:"max_id"`
			MaxIDStr    string  `json:"max_id_str,nocopy"`
			NextResults string  `json:"next_results,nocopy"`
			Query       string  `json:"query,nocopy"`
			RefreshURL  string  `json:"refresh_url,nocopy"`
			Count       int     `json:"count"`
			SinceID     int     `json:"since_id"`
			SinceIDStr  string  `json:"since_id_str,nocopy"`
		} `json:"search_metadata"`
	}
	twitterStatus struct {
		Metadata struct {
			ResultType      string `json:"result_type,nocopy"`
			IsoLanguageCode string `json:"iso_language_code,nocopy"`
		} `json:"metadata"`
		CreatedAt            string          `json:"created_at,nocopy"`
		ID                   int64           `json:"id"`
		IDStr                string          `json:"id_str,nocopy"`
		Text                 string          `json:"text,nocopy"`
		Source               string          `json:"source,nocopy"`
		Truncated            bool            `json:"truncated"`
		InReplyToStatusID    int64           `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr string          `json:"in_reply_to_status_id_str,nocopy"`
		InReplyToUserID      int64           `json:"in_reply_to_user_id"`
		InReplyToUserIDStr   string          `json:"in_reply_to_user_id_str,nocopy"`
		InReplyToScreenName  string          `json:"in_reply_to_screen_name,nocopy"`
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
		Lang                 string          `json:"lang,nocopy"`
	}
	twitterUser struct {
		ID                             int64           `json:"id"`
		IDStr                          string          `json:"id_str,nocopy"`
		Name                           string          `json:"name,nocopy"`
		ScreenName                     string          `json:"screen_name,nocopy"`
		Location                       string          `json:"location,nocopy"`
		Description                    string          `json:"description,nocopy"`
		URL                            any             `json:"url"`
		Entities                       twitterEntities `json:"entities"`
		Protected                      bool            `json:"protected"`
		FollowersCount                 int             `json:"followers_count"`
		FriendsCount                   int             `json:"friends_count"`
		ListedCount                    int             `json:"listed_count"`
		CreatedAt                      string          `json:"created_at,nocopy"`
		FavouritesCount                int             `json:"favourites_count"`
		UtcOffset                      int             `json:"utc_offset"`
		TimeZone                       string          `json:"time_zone,nocopy"`
		GeoEnabled                     bool            `json:"geo_enabled"`
		Verified                       bool            `json:"verified"`
		StatusesCount                  int             `json:"statuses_count"`
		Lang                           string          `json:"lang,nocopy"`
		ContributorsEnabled            bool            `json:"contributors_enabled"`
		IsTranslator                   bool            `json:"is_translator"`
		IsTranslationEnabled           bool            `json:"is_translation_enabled"`
		ProfileBackgroundColor         string          `json:"profile_background_color,nocopy"`
		ProfileBackgroundImageURL      string          `json:"profile_background_image_url,nocopy"`
		ProfileBackgroundImageURLHTTPS string          `json:"profile_background_image_url_https,nocopy"`
		ProfileBackgroundTile          bool            `json:"profile_background_tile"`
		ProfileImageURL                string          `json:"profile_image_url,nocopy"`
		ProfileImageURLHTTPS           string          `json:"profile_image_url_https,nocopy"`
		ProfileBannerURL               string          `json:"profile_banner_url,nocopy"`
		ProfileLinkColor               string          `json:"profile_link_color,nocopy"`
		ProfileSidebarBorderColor      string          `json:"profile_sidebar_border_color,nocopy"`
		ProfileSidebarFillColor        string          `json:"profile_sidebar_fill_color,nocopy"`
		ProfileTextColor               string          `json:"profile_text_color,nocopy"`
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
			ScreenName string `json:"screen_name,nocopy"`
			Name       string `json:"name,nocopy"`
			ID         int64  `json:"id"`
			IDStr      string `json:"id_str,nocopy"`
			Indices    []int  `json:"indices"`
		} `json:"user_mentions"`
		Description struct {
			URLs []twitterURL `json:"urls"`
		} `json:"description"`
		Media []struct {
			ID            int64  `json:"id"`
			IDStr         string `json:"id_str,nocopy"`
			Indices       []int  `json:"indices"`
			MediaURL      string `json:"media_url,nocopy"`
			MediaURLHTTPS string `json:"media_url_https,nocopy"`
			URL           string `json:"url,nocopy"`
			DisplayURL    string `json:"display_url,nocopy"`
			ExpandedURL   string `json:"expanded_url,nocopy"`
			Type          string `json:"type,nocopy"`
			Sizes         map[string]struct {
				W      int    `json:"w"`
				H      int    `json:"h"`
				Resize string `json:"resize,nocopy"`
			} `json:"sizes,nocopy"`
			SourceStatusID    int64  `json:"source_status_id"`
			SourceStatusIDStr string `json:"source_status_id_str,nocopy"`
		} `json:"media"`
	}
	twitterURL struct {
		URL         string       `json:"url,nocopy"`
		URLs        []twitterURL `json:"urls"`
		ExpandedURL string       `json:"expanded_url,nocopy"`
		DisplayURL  string       `json:"display_url,nocopy"`
		Indices     []int        `json:"indices"`
	}
)
