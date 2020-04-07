package chatlog

type ChatResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	ContinuationContents ContinuationContents `json:"continuationContents"`
}

type ContinuationContents struct {
	LiveChatContinuation LiveChatContinuation `json:"liveChatContinuation"`
}

type LiveChatContinuation struct {
	Continuations []Continuation       `json:"continuations"`
	Actions       []ContinuationAction `json:"actions"`
}

type Continuation struct {
	LiveChatReplayContinuationData LiveChatReplayContinuationData `json:"liveChatReplayContinuationData"`
	PlayerSeekContinuationData     PlayerSeekContinuationData     `json:"playerSeekContinuationData"`
}

type LiveChatReplayContinuationData struct {
	PlayerSeekContinuationData
	TimeUntilLastMessageMsec int `json:"timeUntilLastMessageMsec"`
}

type PlayerSeekContinuationData struct {
	Continuation        string `json:"continuation"`
	ClickTrackingParams string `json:"clickTrackingParams"`
}

type ContinuationAction struct {
	ReplayChatItemAction ReplayChatItemAction `json:"replayChatItemAction"`
}

type ReplayChatItemAction struct {
	Actions             []ChatAction `json:"actions"`
	VideoOffsetTimeMsec string       `json:"videoOffsetTimeMsec"`
}

type ChatAction struct {
	AddChatItemAction AddChatItemAction `json:"addChatItemAction"`
}

type AddChatItemAction struct {
	Item     Item   `json:"item"`
	ClientId string `json:"clientId"`
}

type Item struct {
	LiveChatViewerEngagementMessageRenderer LiveChatMessageRenderer `json:"liveChatViewerEngagementMessageRenderer"`
	LiveChatTextMessageRenderer             LiveChatMessageRenderer `json:"liveChatTextMessageRenderer"`
}

type LiveChatMessageRenderer struct {
	Id                       string              `json:"id"`
	TimestampUsec            string              `json:"timestampUsec"`
	Icon                     Icon                `json:"icon"`
	Message                  Message             `json:"message"`
	AuthorName               AuthorName          `json:"authorName"`
	AuthorPhoto              AuthorPhoto         `json:"authorPhoto"`
	ContextMenuEndpoint      ContextMenuEndpoint `json:"contextMenuEndpoint"`
	AuthorExternalChannelId  string              `json:"authorExternalChannelId"`
	ContextMenuAccessibility Accessibility       `json:"contextMenuAccessibility"`
	TimestampText            TimestampText       `json:"timestampText"`
	AuthorBadges             []AuthorBadge       `json:"authorBadges"`
}

type Icon struct {
	IconType string `json:"iconType"`
}

type Message struct {
	Runs []Run `json:"runs"`
}

type Run struct {
	Text  string `json:"text"`
	Emoji Emoji  `json:"emoji"`
}

type AuthorName struct {
	SimpleText string `json:"simpleText"`
}

type AuthorPhoto struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ContextMenuEndpoint struct {
	ClickTrackingParams             string                          `json:"clickTrackingParams"`
	CommandMetadata                 CommandMetadata                 `json:"commandMetadata"`
	LiveChatItemContextMenuEndpoint LiveChatItemContextMenuEndpoint `json:"liveChatItemContextMenuEndpoint"`
}

type CommandMetadata struct {
	WebCommandMetadata WebCommandMetadata `json:"webCommandMetadata"`
}

type WebCommandMetadata struct {
	IgnoreNavigation bool `json:"ignoreNavigation"`
}

type LiveChatItemContextMenuEndpoint struct {
	Params string `json:"params"`
}

type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

type AccessibilityData struct {
	Label string `json:"label"`
}

type TimestampText struct {
	SimpleText string `json:"simpleText"`
}

type Emoji struct {
	EmojiId       string   `json:"emojiId"`
	Shortcuts     []string `json:"shortcuts"`
	SearchTerms   []string `json:"searchTerms"`
	Image         Image    `json:"image"`
	IsCustomEmoji bool     `json:"isCustomEmoji"`
}

type Image struct {
	Thumbnails    []Thumbnail   `json:"thumbnails"`
	Accessibility Accessibility `json:"accessibility"`
}

type AuthorBadge struct {
	LiveChatAuthorBadgeRenderer LiveChatAuthorBadgeRenderer `json:"liveChatAuthorBadgeRenderer"`
}

type LiveChatAuthorBadgeRenderer struct {
	CustomThumbnail CustomThumbnail `json:"customThumbnail"`
	Tooltip         string          `json:"tooltip"`
	Accessibility   Accessibility   `json:"accessibility"`
}

type CustomThumbnail struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}
