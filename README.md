# wiser

文档下载地址：
> https://dumps.wikimedia.org/zhwiki/latest/zhwiki-latest-pages-articles.xml.bz2

### 添加文档 add_document
* 存储文档 正排
* 缓冲区创建倒排列表
* 缓冲区达到阈值时，更新至存储器


### 构建倒排列表 text_to_postings_lists
* ngram
* token_posting_list
* merge_inverted_index

