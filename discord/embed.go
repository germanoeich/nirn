package discord

import (
	"github.com/germanoeich/nirn/proto/discord/v1/model"
	"strings"
)

// limitations: https://discord.com/developers/docs/resources/channel#embed-limits
// TODO: implement NewEmbedX functions that ensures limitations

type EmbedType string

const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypeVideo   EmbedType = "video"
	EmbedTypeGIFV    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

func (e EmbedType) ToProto() model.MessageData_MessageEmbedData_MessageEmbedType {
	upperType := strings.ToUpper(string(e))
	value, ok := model.MessageData_MessageEmbedData_MessageEmbedType_value[upperType]
	if !ok {
		return 0
	}

	return model.MessageData_MessageEmbedData_MessageEmbedType(value)
}

// Embed https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	Title       string          `json:"title,omitempty"`       // title of embed
	Type        EmbedType       `json:"type,omitempty"`        // type of embed (always "rich" for webhook embeds)
	Description string          `json:"description,omitempty"` // description of embed
	URL         string          `json:"url,omitempty"`         // url of embed
	Timestamp   Time            `json:"timestamp,omitempty"`   // timestamp	timestamp of embed content
	Color       int             `json:"color,omitempty"`       // color code of the embed
	Footer      *EmbedFooter    `json:"footer,omitempty"`      // embed footer object	footer information
	Image       *EmbedImage     `json:"image,omitempty"`       // embed image object	image information
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`   // embed thumbnail object	thumbnail information
	Video       *EmbedVideo     `json:"video,omitempty"`       // embed video object	video information
	Provider    *EmbedProvider  `json:"provider,omitempty"`    // embed provider object	provider information
	Author      *EmbedAuthor    `json:"author,omitempty"`      // embed author object	author information
	Fields      EmbedFieldArray `json:"fields,omitempty"`      //	array of embed field objects	fields information
}

func (e Embed) ToProto() *model.MessageData_MessageEmbedData {
	return &model.MessageData_MessageEmbedData{
		Title:       e.Title,
		Type:        e.Type.ToProto(),
		Description: e.Description,
		Url:         e.URL,
		Timestamp:   e.Timestamp.ToProto(),
		Color:       uint32(e.Color),
		Footer:      e.Footer.ToProto(),
		Image:       e.Image.ToProto(),
		Thumbnail:   e.Thumbnail.ToProto(),
		Video:       e.Video.ToProto(),
		Provider:    e.Provider.ToProto(),
		Author:      e.Author.ToProto(),
		Fields:      e.Fields.ToProto(),
	}
}

type EmbedArray []*Embed

func (a EmbedArray) ToProto() []*model.MessageData_MessageEmbedData {
	if a == nil {
		return nil
	}
	var ret []*model.MessageData_MessageEmbedData
	for _, el := range a {
		if el == nil {
			continue
		}
		ret = append(ret, el.ToProto())
	}
	return ret
}

// EmbedThumbnail https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`       // ?| , source url of image (only supports http(s) and attachments)
	ProxyURL string `json:"proxy_url,omitempty"` // ?| , a proxied url of the image
	Height   int    `json:"height,omitempty"`    // ?| , height of image
	Width    int    `json:"width,omitempty"`     // ?| , width of image
}

func (e *EmbedThumbnail) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedThumbnailData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedThumbnailData{
		Url:      e.URL,
		ProxyUrl: e.ProxyURL,
		Width:    uint32(e.Width),
		Height:   uint32(e.Height),
	}
}

// EmbedVideo https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	URL    string `json:"url,omitempty"`    // ?| , source url of video
	Height int    `json:"height,omitempty"` // ?| , height of video
	Width  int    `json:"width,omitempty"`  // ?| , width of video
}

func (e *EmbedVideo) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedVideoData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedVideoData{
		Url:    e.URL,
		Width:  uint32(e.Width),
		Height: uint32(e.Height),
	}
}

// EmbedImage https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	URL      string `json:"url,omitempty"`       // ?| , source url of image (only supports http(s) and attachments)
	ProxyURL string `json:"proxy_url,omitempty"` // ?| , a proxied url of the image
	Height   int    `json:"height,omitempty"`    // ?| , height of image
	Width    int    `json:"width,omitempty"`     // ?| , width of image
}

func (e *EmbedImage) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedImageData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedImageData{
		Url:    e.URL,
		Width:  uint32(e.Width),
		Height: uint32(e.Height),
	}
}

// EmbedProvider https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name string `json:"name,omitempty"` // ?| , name of provider
	URL  string `json:"url,omitempty"`  // ?| , url of provider
}

func (e *EmbedProvider) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedProviderData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedProviderData{
		Name: e.Name,
		Url:  e.URL,
	}
}

// EmbedAuthor https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`           // ?| , name of author
	URL          string `json:"url,omitempty"`            // ?| , url of author
	IconURL      string `json:"icon_url,omitempty"`       // ?| , url of author icon (only supports http(s) and attachments)
	ProxyIconURL string `json:"proxy_icon_url,omitempty"` // ?| , a proxied url of author icon
}

func (e *EmbedAuthor) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedAuthorData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedAuthorData{
		Name:         e.Name,
		Url:          e.URL,
		IconUrl:      e.IconURL,
		ProxyIconUrl: e.ProxyIconURL,
	}
}

// EmbedFooter https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         string `json:"text"`                     //  | , url of author
	IconURL      string `json:"icon_url,omitempty"`       // ?| , url of footer icon (only supports http(s) and attachments)
	ProxyIconURL string `json:"proxy_icon_url,omitempty"` // ?| , a proxied url of footer icon
}

func (e *EmbedFooter) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedFooterData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedFooterData{
		Text:         e.Text,
		IconUrl:      e.IconURL,
		ProxyIconUrl: e.ProxyIconURL,
	}
}

// EmbedField https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`             //  | , name of the field
	Value  string `json:"value"`            //  | , value of the field
	Inline bool   `json:"inline,omitempty"` // ?| , whether or not this field should display inline
}

func (e *EmbedField) ToProto() *model.MessageData_MessageEmbedData_MessageEmbedFieldData {
	if e == nil {
		return nil
	}
	return &model.MessageData_MessageEmbedData_MessageEmbedFieldData{
		Name:   e.Name,
		Value:  e.Value,
		Inline: e.Inline,
	}
}

type EmbedFieldArray []*EmbedField

func (a EmbedFieldArray) ToProto() []*model.MessageData_MessageEmbedData_MessageEmbedFieldData {
	var ret []*model.MessageData_MessageEmbedData_MessageEmbedFieldData
	for _, el := range a {
		ret = append(ret, el.ToProto())
	}
	return ret
}
