// Code generated from payload-medium.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Company any `json:"company"`
	Person  struct {
		Aboutme struct {
			Avatar any    `json:"avatar"`
			Bio    any    `json:"bio"`
			Handle string `json:"handle"`
		} `json:"aboutme"`
		Angellist struct {
			Avatar    string `json:"avatar"`
			Bio       string `json:"bio"`
			Blog      string `json:"blog"`
			Followers int64  `json:"followers"`
			Handle    string `json:"handle"`
			ID        int64  `json:"id"`
			Site      string `json:"site"`
		} `json:"angellist"`
		Avatar     string `json:"avatar"`
		Bio        string `json:"bio"`
		Email      string `json:"email"`
		Employment struct {
			Domain string `json:"domain"`
			Name   string `json:"name"`
			Title  string `json:"title"`
		} `json:"employment"`
		Facebook struct {
			Handle string `json:"handle"`
		} `json:"facebook"`
		Foursquare struct {
			Handle any `json:"handle"`
		} `json:"foursquare"`
		Fuzzy  bool   `json:"fuzzy"`
		Gender string `json:"gender"`
		Geo    struct {
			City    string  `json:"city"`
			Country string  `json:"country"`
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			State   string  `json:"state"`
		} `json:"geo"`
		Github struct {
			Avatar    string `json:"avatar"`
			Blog      string `json:"blog"`
			Company   string `json:"company"`
			Followers int64  `json:"followers"`
			Following int64  `json:"following"`
			Handle    string `json:"handle"`
			ID        int64  `json:"id"`
		} `json:"github"`
		Googleplus struct {
			Handle any `json:"handle"`
		} `json:"googleplus"`
		Gravatar struct {
			Avatar  string `json:"avatar"`
			Avatars []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"avatars"`
			Handle string `json:"handle"`
			Urls   []any  `json:"urls"`
		} `json:"gravatar"`
		ID    string `json:"id"`
		Klout struct {
			Handle any `json:"handle"`
			Score  any `json:"score"`
		} `json:"klout"`
		Linkedin struct {
			Handle string `json:"handle"`
		} `json:"linkedin"`
		Location string `json:"location"`
		Name     struct {
			FamilyName string `json:"familyName"`
			FullName   string `json:"fullName"`
			GivenName  string `json:"givenName"`
		} `json:"name"`
		Site    string `json:"site"`
		Twitter struct {
			Avatar    any    `json:"avatar"`
			Bio       any    `json:"bio"`
			Favorites int64  `json:"favorites"`
			Followers int64  `json:"followers"`
			Following int64  `json:"following"`
			Handle    string `json:"handle"`
			ID        int64  `json:"id"`
			Location  string `json:"location"`
			Site      string `json:"site"`
			Statuses  int64  `json:"statuses"`
		} `json:"twitter"`
	} `json:"person"`
}
