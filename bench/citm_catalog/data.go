// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// citm_catalog root type, with the root renamed to Log for the lightning harness.
package bench

type (
	Log struct {
		AreaNames                map[string]string `json:"areaNames"`
		AudienceSubCategoryNames map[string]string `json:"audienceSubCategoryNames"`
		BlockNames               map[string]string `json:"blockNames"`
		Events                   map[string]struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Logo        string `json:"logo"`
			Name        string `json:"name"`
			SubTopicIds []int  `json:"subTopicIds"`
			SubjectCode any    `json:"subjectCode"`
			Subtitle    any    `json:"subtitle"`
			TopicIds    []int  `json:"topicIds"`
		} `json:"events"`
		Performances []struct {
			EventID int `json:"eventId"`
			ID      int `json:"id"`
			Logo    any `json:"logo"`
			Name    any `json:"name"`
			Prices  []struct {
				Amount                int   `json:"amount"`
				AudienceSubCategoryID int64 `json:"audienceSubCategoryId"`
				SeatCategoryID        int64 `json:"seatCategoryId"`
			} `json:"prices"`
			SeatCategories []struct {
				Areas []struct {
					AreaID   int   `json:"areaId"`
					BlockIds []any `json:"blockIds"`
				} `json:"areas"`
				SeatCategoryID int `json:"seatCategoryId"`
			} `json:"seatCategories"`
			SeatMapImage any    `json:"seatMapImage"`
			Start        int64  `json:"start"`
			VenueCode    string `json:"venueCode"`
		} `json:"performances"`
		SeatCategoryNames map[string]string   `json:"seatCategoryNames"`
		SubTopicNames     map[string]string   `json:"subTopicNames"`
		SubjectNames      map[string]string   `json:"subjectNames"`
		TopicNames        map[string]string   `json:"topicNames"`
		TopicSubTopics    map[string][]uint64 `json:"topicSubTopics"`
		VenueNames        map[string]string   `json:"venueNames"`
	}
)
