package internal

//InvertedIndexValue 倒排索引
type InvertedIndexValue struct {
	TokenID       int64
	postingList   *PostingsList
	docsCount     int64
	positionCount int64
}

// InvertedIndexHash 倒排hash
type InvertedIndexHash map[int64]*InvertedIndexValue

// PostingsList 倒排列表
type PostingsList struct {
	DocID         int64
	positions     []int64
	positionCount int64
	next          *PostingsList
}
