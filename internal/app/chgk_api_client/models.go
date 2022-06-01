package chgk_api_client

import "encoding/xml"

type Search struct {
	XMLName      xml.Name   `xml:"search" json:"-"`
	QuestionList []Question `xml:"question" json:"question"`
}

type Question struct {
	ID       uint64 `xml:"QuestionId" json:"id"`
	Question string `xml:"Question" json:"question"`
	Answer   string `xml:"Answer" json:"answer"`
	Authors  string `xml:"Authors" json:"authors"`
	Comments string `xml:"Comments" json:"comments"`
}
