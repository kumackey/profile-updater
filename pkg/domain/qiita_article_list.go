package domain

import "sort"

type QiitaArticleList []*qiitaArticle

func (l QiitaArticleList) ToProfileMarkdown() string {
	profileMarkdown := "\n"
	for _, article := range l {
		profileMarkdown = profileMarkdown + article.toMarkdown() + "\n"
	}

	return profileMarkdown
}

func (l QiitaArticleList) SortByPublishedAt() QiitaArticleList {
	sort.Slice(l, func(i, j int) bool {
		// 公開が遅い順
		return l[j].publishedAt.Unix() < l[i].publishedAt.Unix()
	})

	return l
}

func (l QiitaArticleList) Limit(limit int) QiitaArticleList {
	list := QiitaArticleList{}
	count := 0
	for _, article := range l {
		if limit <= count {
			break
		}

		list = append(list, article)
		count += 1
	}

	return list
}
