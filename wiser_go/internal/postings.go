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
			//TODO: 不确定要不要加 先todo
			// index.positionCount += toBeAddedIndex.positionCount
			delete(*toBeAdded, tokenID)
		}
	}
	for tokenID, index := range *toBeAdded {
		(*base)[tokenID] = index
	}

}

// /**
//  * 将内存上（小倒排索引中）的倒排列表与存储器上的倒排列表合并后存储到数据库中
//  * @param[in] env 存储着应用程序运行环境的结构体
//  * @param[in] p 含有倒排列表的倒排索引中的索引项
//  */
// void update_postings(const wiser_env *env, inverted_index_value *p) {
//     int old_postings_len;
//     postings_list *old_postings;

//     if (!fetch_postings(env, p->token_id, &old_postings, &old_postings_len)) {
//         buffer *buf;
//         if (old_postings_len) {
//             p->postings_list = merge_postings(old_postings, p->postings_list);
//             p->docs_count += old_postings_len;
//         }
//         if ((buf = alloc_buffer())) {
//             encode_postings(env, p->postings_list, p->docs_count, buf);
//             db_update_postings(env, p->token_id, p->docs_count, BUFFER_PTR(buf), BUFFER_SIZE(buf));
//             free_buffer(buf);
//         }
//     } else {
//         print_error("cannot fetch old postings list of token(%d) for update.", p->token_id);
//     }
// }
