package internal

// MergePostings merge two postings list
func MergePostings(pa, pb *PostingsList) *PostingsList {
	p := new(PostingsList)
	temp := new(PostingsList)
	for pa != nil && pb != nil {
		if pa.DocID <= pb.DocID {
			p.next = pa
			pa = pa.next
		} else {
			p.next = pb
			pb = pb.next
		}
	}
	if pa != nil {
		p.next = pa
	}
	if pb != nil {
		p.next = pb
	}

	return p
}
