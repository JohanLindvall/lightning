// Code generated from payload_large.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Topics struct {
		CanCreateTopic bool   `json:"can_create_topic"`
		Draft          any    `json:"draft"`
		DraftKey       string `json:"draft_key,nocopy"`
		DraftSequence  any    `json:"draft_sequence"`
		MoreTopicsUrl  string `json:"more_topics_url,nocopy"`
		PerPage        int64  `json:"per_page"`
		Topics         []struct {
			Archetype          string `json:"archetype,nocopy"`
			Archived           bool   `json:"archived"`
			Bookmarked         any    `json:"bookmarked"`
			Bumped             bool   `json:"bumped"`
			BumpedAt           string `json:"bumped_at,nocopy"`
			CategoryId         int64  `json:"category_id"`
			Closed             bool   `json:"closed"`
			CreatedAt          string `json:"created_at,nocopy"`
			Excerpt            string `json:"excerpt,nocopy"`
			FancyTitle         string `json:"fancy_title,nocopy"`
			HasSummary         bool   `json:"has_summary"`
			HighestPostNumber  int64  `json:"highest_post_number"`
			ID                 int64  `json:"id"`
			ImageUrl           any    `json:"image_url"`
			LastPostedAt       string `json:"last_posted_at,nocopy"`
			LastPosterUsername string `json:"last_poster_username,nocopy"`
			LikeCount          int64  `json:"like_count"`
			Liked              any    `json:"liked"`
			Pinned             bool   `json:"pinned"`
			PinnedGlobally     bool   `json:"pinned_globally"`
			Posters            []struct {
				Description string `json:"description,nocopy"`
				Extras      any    `json:"extras"`
				UserId      int64  `json:"user_id"`
			} `json:"posters"`
			PostsCount int64  `json:"posts_count"`
			ReplyCount int64  `json:"reply_count"`
			Slug       string `json:"slug,nocopy"`
			Title      string `json:"title,nocopy"`
			Unpinned   any    `json:"unpinned"`
			Unseen     bool   `json:"unseen"`
			Views      int64  `json:"views"`
			Visible    bool   `json:"visible"`
		} `json:"topics"`
	} `json:"topics"`
	Users []struct {
		AvatarTemplate string `json:"avatar_template,nocopy"`
		ID             int64  `json:"id"`
		Username       string `json:"username,nocopy"`
	} `json:"users"`
}
