// Code generated from payload_medium.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Company any `json:"company"`
	Person  struct {
		Aboutme struct {
			Avatar any    `json:"avatar"`
			Bio    any    `json:"bio"`
			Handle string `json:"handle,nocopy"`
		} `json:"aboutme"`
		Angellist struct {
			Avatar    string `json:"avatar,nocopy"`
			Bio       string `json:"bio,nocopy"`
			Blog      string `json:"blog,nocopy"`
			Followers int64  `json:"followers"`
			Handle    string `json:"handle,nocopy"`
			ID        int64  `json:"id"`
			Site      string `json:"site,nocopy"`
		} `json:"angellist"`
		Avatar     string `json:"avatar,nocopy"`
		Bio        string `json:"bio,nocopy"`
		Email      string `json:"email,nocopy"`
		Employment struct {
			Domain string `json:"domain,nocopy"`
			Name   string `json:"name,nocopy"`
			Title  string `json:"title,nocopy"`
		} `json:"employment"`
		Facebook struct {
			Handle string `json:"handle,nocopy"`
		} `json:"facebook"`
		Foursquare struct {
			Handle any `json:"handle"`
		} `json:"foursquare"`
		Fuzzy  bool   `json:"fuzzy"`
		Gender string `json:"gender,nocopy"`
		Geo    struct {
			City    string  `json:"city,nocopy"`
			Country string  `json:"country,nocopy"`
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			State   string  `json:"state,nocopy"`
		} `json:"geo"`
		Github struct {
			Avatar    string `json:"avatar,nocopy"`
			Blog      string `json:"blog,nocopy"`
			Company   string `json:"company,nocopy"`
			Followers int64  `json:"followers"`
			Following int64  `json:"following"`
			Handle    string `json:"handle,nocopy"`
			ID        int64  `json:"id"`
		} `json:"github"`
		Googleplus struct {
			Handle any `json:"handle"`
		} `json:"googleplus"`
		Gravatar struct {
			Avatar  string `json:"avatar,nocopy"`
			Avatars []struct {
				Type string `json:"type,nocopy"`
				URL  string `json:"url,nocopy"`
			} `json:"avatars"`
			Handle string `json:"handle,nocopy"`
			Urls   []any  `json:"urls"`
		} `json:"gravatar"`
		ID    string `json:"id,nocopy"`
		Klout struct {
			Handle any `json:"handle"`
			Score  any `json:"score"`
		} `json:"klout"`
		Linkedin struct {
			Handle string `json:"handle,nocopy"`
		} `json:"linkedin"`
		Location string `json:"location,nocopy"`
		Name     struct {
			FamilyName string `json:"familyName,nocopy"`
			FullName   string `json:"fullName,nocopy"`
			GivenName  string `json:"givenName,nocopy"`
		} `json:"name"`
		Site    string `json:"site,nocopy"`
		Twitter struct {
			Avatar    any    `json:"avatar"`
			Bio       any    `json:"bio"`
			Favorites int64  `json:"favorites"`
			Followers int64  `json:"followers"`
			Following int64  `json:"following"`
			Handle    string `json:"handle,nocopy"`
			ID        int64  `json:"id"`
			Location  string `json:"location,nocopy"`
			Site      string `json:"site,nocopy"`
			Statuses  int64  `json:"statuses"`
		} `json:"twitter"`
	} `json:"person"`
}
