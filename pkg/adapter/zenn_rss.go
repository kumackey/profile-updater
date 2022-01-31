package adapter

import "encoding/xml"

// https://zenn.dev/link/comments/731a3eba374a8e
type zennRSS struct {
	XMLName xml.Name   `xml:"rss"`
	Items   []zennItem `xml:"channel>item"`
}

type zennItem struct {
	Title       string `xml:"title"`
	Description struct {
		XMLName xml.Name `xml:"description"`
		Value   string   `xml:",cdata"`
	}
	Link string `xml:"link"`
	GuID struct {
		XMLName     xml.Name `xml:"guid"`
		IsPermaLink bool     `xml:"isPermaLink,attr"`
		Value       string   `xml:",chardata"`
	}
	PubDate   string `xml:"pubDate"`
	Enclosure struct {
		XMLName string `xml:"enclosure"`
		URL     string `xml:"url,attr"`
		Length  int    `xml:"length,attr"`
		Type    string `xml:"type,attr"`
	}
	Creator string `xml:"creator"`
}
