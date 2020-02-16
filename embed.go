package main

import (
	"github.com/bwmarrin/discordgo"
)

// Embed ...
type Embed struct {
	*discordgo.MessageEmbed
}

// set char limits for embed
const (
	embedLimitTitle       = 256
	embedLimitDescription = 2048
	embedLimitFieldValue  = 1024
	embedLimitFieldName   = 256
	embedLimitField       = 25
	embedLimitFooter      = 2048
	embedLimit            = 4000
)

// NewEmbed returns a new embed object
func NewEmbed() *Embed {
	return &Embed{&discordgo.MessageEmbed{}}
}

// SetTitle ...
func (e *Embed) SetTitle(name string) *Embed {
	e.Title = name
	return e
}

//SetDescription [desc]
func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > 2048 {
		description = description[:2048]
	}
	e.Description = description
	return e
}

//AddField [name] [value]
func (e *Embed) AddField(name, value string) *Embed {
	if len(name) > 1024 {
		name = name[:1024]
	}
	if len(value) > 1024 {
		value = value[:1024]
	}

	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:  name,
		Value: value,
	})
	return e
}

//SetFooter [text] [iconURL]
func (e *Embed) SetFooter(args ...string) *Embed {
	iconURL := ""
	text := ""
	proxyURL := ""

	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}

	e.Footer = &discordgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}

	return e
}

//SetImage ...
func (e *Embed) SetImage(args ...string) *Embed {
	var URL, proxyURL string

	switch {
	case len(args) > 1:
		proxyURL = args[1]
		fallthrough
	case len(args) > 0:
		URL = args[0]
	case len(args) == 0:
		return e
	}

	e.Image = &discordgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}

	return e
}

//SetThumbnail ...
func (e *Embed) SetThumbnail(args ...string) *Embed {
	var URL, proxyURL string

	switch {
	case len(args) > 1:
		proxyURL = args[1]
		fallthrough
	case len(args) > 0:
		URL = args[0]
		fallthrough
	case len(args) == 0:
		return e
	}

	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

//SetAuthor ...
func (e *Embed) SetAuthor(args ...string) *Embed {
	var name, iconURL, URL, proxyURL string

	switch {
	case len(args) > 3:
		proxyURL = args[1]
		fallthrough
	case len(args) > 2:
		URL = args[0]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		name = args[0]
		fallthrough
	case len(args) == 0:
		return e
	}

	e.Author = &discordgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}
	return e
}

//SetURL ...
func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

//SetColor ...
func (e *Embed) SetColor(clr int) *Embed {
	e.Color = clr
	return e
}

//InlineAllFields sets all fields to be inline
func (e *Embed) InlineAllFields() *Embed {
	for _, v := range e.Fields {
		v.Inline = true
	}
	return e
}

//Truncate cleans up anything over the char limie
func (e *Embed) Truncate() *Embed {
	e.TruncateDescription()
	e.TruncateFields()
	e.TruncateFooter()
	e.TruncateTitle()
	return e
}

//TruncateDescription ...
func (e *Embed) TruncateDescription() *Embed {
	if len(e.Description) > embedLimitDescription {
		e.Description = e.Description[:embedLimitDescription]
	}
	return e
}

//TruncateFields ...
func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > embedLimitField {
		e.Fields = e.Fields[:embedLimitField]
	}

	for _, v := range e.Fields {
		if len(v.Name) > embedLimitFieldName {
			v.Name = v.Name[:embedLimitFieldName]
		}

		if len(v.Value) > embedLimitFieldValue {
			v.Value = v.Value[:embedLimitFieldValue]
		}
	}
	return e
}

//TruncateTitle ...
func (e *Embed) TruncateTitle() *Embed {
	if len(e.Title) > embedLimitTitle {
		e.Title = e.Title[:embedLimitTitle]
	}
	return e
}

//TruncateFooter ...
func (e *Embed) TruncateFooter() *Embed {
	if e.Footer != nil && len(e.Footer.Text) > embedLimitFooter {
		e.Footer.Text = e.Footer.Text[:embedLimitFooter]
	}
	return e
}
