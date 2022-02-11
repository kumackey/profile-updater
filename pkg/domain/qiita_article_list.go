package domain

type QiitaArticleList []*qiitaArticle

//
//func (l ZennArticleList) ToProfileMarkdown() string {
//	profileMarkdown := "\n"
//	for _, article := range l {
//		profileMarkdown = profileMarkdown + article.toMarkdown() + "\n"
//	}
//
//	return profileMarkdown
//}
//
//func (l ZennArticleList) SortByPublishedAt() ZennArticleList {
//	sort.Slice(l, func(i, j int) bool {
//		// 公開が遅い順
//		return l[j].publishedAt.Unix() < l[i].publishedAt.Unix()
//	})
//
//	return l
//}
//
//func (l ZennArticleList) Limit(limit int) ZennArticleList {
//	list := ZennArticleList{}
//	count := 0
//	for _, article := range l {
//		if limit <= count {
//			break
//		}
//
//		list = append(list, article)
//		count += 1
//	}
//
//	return list
//}
