package adapter

import "encoding/xml"

type qiitaUserFeed struct {
	XMLName xml.Name     `xml:"feed"`
	Entries []qiitaEntry `xml:"entry"`
}

type qiitaEntry struct {
	Title     string `xml:"title"`
	URL       string `xml:"url"`
	Published string `xml:"published"`
}
