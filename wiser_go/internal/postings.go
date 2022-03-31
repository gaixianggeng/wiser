package internal

// MergePostings merge two postings list
// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
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

// MergeInvertedIndex 合并两个倒排索引
func MergeInvertedIndex(base, toBeAdded *InvertedIndexHash) {
	for tokenID, index := range *base {
		if toBeAddedIndex, ok := (*toBeAdded)[tokenID]; ok {
			index.postingList = MergePostings(index.postingList, toBeAddedIndex.postingList)
			index.docsCount += toBeAddedIndex.docsCount
			index.positionCount += toBeAddedIndex.positionCount
			delete(*toBeAdded, tokenID)
		}
	}
	for tokenID, index := range *toBeAdded {
		(*base)[tokenID] = index
	}

	// inverted_index_value *p, *temp;

	// HASH_ITER(hh, to_be_added, p, temp) {
	//     inverted_index_value *t;
	//     HASH_DEL(to_be_added, p);
	//     HASH_FIND_INT(base, &p->token_id, t);
	//     if (t) {
	//         t->postings_list = merge_postings(t->postings_list, p->postings_list);
	//         t->docs_count += p->docs_count;
	//         free(p);
	//     } else {
	//         HASH_ADD_INT(base, token_id, p);
	//     }
	// }
}
