package model

import (
	"fmt"
	"strings"
)

type (
	SlotUri struct {
		Name           string      `json:"name"`
		Description    string      `json:"description"`
		Image          string      `json:"image"`
		ExternalUrl    string      `json:"external_url"`
		BannerImageUrl string      `json:"banner_image_url"`
		YoutubeUrl     string      `json:"youtube_url"`
		Attributes     []Attribute `json:"attributes"`
	}
	Attribute struct {
		Value     interface{} `json:"value"`
		TraitType string      `json:"trait_type"`
	}
)

func (s SlotUri) AttributeItem(tt string) string {
	for _, attr := range s.Attributes {
		if strings.HasPrefix(attr.TraitType, tt) {
			v := attr.Value
			switch attr.Value.(type) {
			case string:
				return fmt.Sprintf("%s", v)
			case float64:
				u, ok := v.(uint64)
				if !ok {
					return fmt.Sprintf("%.0f", v)
				}
				return fmt.Sprintf("%d", u)
			default:
				return fmt.Sprintf("%v", v)
			}
		}
	}
	return ""
}

func (s SlotUri) Slug() string {
	split := strings.Split(s.ExternalUrl, "/")
	return split[len(split)-1]
}
