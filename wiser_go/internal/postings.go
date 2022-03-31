package internal

// MergePostings merge two postings list
func MergePostings(pa, pb *PostingsList) *PostingsList {
	ret := new(PostingsList)
	p := new(PostingsList)
	p = nil
	for pa != nil || pb != nil {

		temp := new(PostingsList)
		if pb == nil || (pa != nil && pa.DocID <= pb.DocID) {
			temp = pa
			pa = pa.next
		} else if pa == nil || (pb != nil && pa.DocID > pb.DocID) {
			temp = pb
			pb = pb.next
		} else {
			break
		}
		temp.next = nil

		if p == nil {
			ret.next = temp
		} else {
			p.next = temp
		}

		p = temp
	}

	return ret.next
}
