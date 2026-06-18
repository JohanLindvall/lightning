// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// citm_catalog root type, with the root renamed to Benchmark for the lightning harness.
package bench

type (
	Benchmark struct {
		AreaNames                map[string]string `json:"areaNames,nocopy"`
		AudienceSubCategoryNames map[string]string `json:"audienceSubCategoryNames,nocopy"`
		BlockNames               map[string]string `json:"blockNames,nocopy"`
		Events                   map[string]struct {
			Description string `json:"description,nocopy"`
			ID          int    `json:"id"`
			Logo        string `json:"logo,nocopy"`
			Name        string `json:"name,nocopy"`
			SubTopicIds []int  `json:"subTopicIds"`
			SubjectCode any    `json:"subjectCode"`
			Subtitle    any    `json:"subtitle"`
			TopicIds    []int  `json:"topicIds"`
		} `json:"events,nocopy"`
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
			VenueCode    string `json:"venueCode,nocopy"`
		} `json:"performances"`
		SeatCategoryNames map[string]string   `json:"seatCategoryNames,nocopy"`
		SubTopicNames     map[string]string   `json:"subTopicNames,nocopy"`
		SubjectNames      map[string]string   `json:"subjectNames,nocopy"`
		TopicNames        map[string]string   `json:"topicNames,nocopy"`
		TopicSubTopics    map[string][]uint64 `json:"topicSubTopics,nocopy"`
		VenueNames        map[string]string   `json:"venueNames,nocopy"`
	}
)
