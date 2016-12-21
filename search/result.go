package search

type Json map[string]interface{}

type JsonArray []Json

type ResultItem struct {
	Score  float64 `json:"score"`
	Source string  `json:"source"`
	Data   Json    `json:"data"`
}

type ResultArray []*ResultItem

type SearchResult struct {
	Took      float64     `json:"took"`
	TotalHits int64       `json:"total_hits"`
	Results   ResultArray `json:"results"`
}

func NewSearchResult() *SearchResult {
	return &SearchResult{
		Took:      0.0,
		TotalHits: 0,
		Results:   ResultArray{},
	}
}

func NewResultItem(data Json, score float64, source string) *ResultItem {
	return &ResultItem{
		Score:  score,
		Source: source,
		Data:   data,
	}
}

func (s *SearchResult) Add(results ResultArray) {
	s.Results = append(s.Results, results...)
}
