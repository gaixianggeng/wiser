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


### token_to_postings_list
* db_get_token_id
* create_new_inverted_index 新建倒排索引
* create_new_postings_list  新建倒排列表

malloc 内存分配
