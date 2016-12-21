package search

type Searcher interface {
	Interested(string) (bool, int64, error)
	Search(string) (ResultArray, int64, error)
}
