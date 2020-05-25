package chat

type ChatResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	ResponseContext      ResponseContext      `json:"responseContext"`
	ContinuationContents ContinuationContents `json:"continuationContents"`
}

type ResponseContext struct {
	Errors struct{ Error []Error } `json:"errors"`
}

type Error struct {
	Domain               string `json:"domain"`
	Code                 string `json:"code"`
	DebugInfo            string `json:"debugInfo"`
	ExternalErrorMessage string `json:"externalErrorMessage"`
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
	TimeUntilLastMessageMsec int    `json:"timeUntilLastMessageMsec"`
	Continuation             string `json:"continuation"`
}

type PlayerSeekContinuationData struct {
	Continuation string `json:"continuation"`
}

type ContinuationAction struct {
	ReplayChatItemAction ReplayChatItemAction `json:"replayChatItemAction"`
}

type ReplayChatItemAction struct {
	Actions             []ChatAction `json:"actions"`
	VideoOffsetTimeMsec string       `json:"videoOffsetTimeMsec"`
}

type ChatAction struct {
	AddChatItemAction           AddChatItemAction           `json:"addChatItemAction"`
	AddLiveChatTickerItemAction AddLiveChatTickerItemAction `json:"addLiveChatTickerItemAction"`
}

type AddChatItemAction struct {
	Item     ChatItem `json:"item"`
	ClientId string   `json:"clientId"`
}

type AddLiveChatTickerItemAction struct {
	Item        LiveChatTickerItem `json:"item"`
	DurationSec string             `json:"durationSec"`
}

type ChatItem struct {
	LiveChatViewerEngagementMessageRenderer LiveChatViewerEngagementMessageRenderer `json:"liveChatViewerEngagementMessageRenderer"`
	LiveChatTextMessageRenderer             LiveChatTextMessageRenderer             `json:"liveChatTextMessageRenderer"`
	LiveChatMembershipItemRenderer          LiveChatMembershipItemRenderer          `json:"liveChatMembershipItemRenderer"`
	LiveChatPaidMessageRenderer             LiveChatPaidMessageRenderer             `json:"liveChatPaidMessageRenderer"`
	LiveChatPlaceholderItemRenderer         LiveChatPlaceholderItemRenderer         `json:"liveChatPlaceholderItemRenderer"`
}

type LiveChatTickerItem struct {
	LiveChatTickerPaidMessageItemRenderer LiveChatTickerPaidMessageItemRenderer `json:"liveChatTickerPaidMessageItemRenderer"`
	LiveChatTickerSponsorItemRenderer     LiveChatTickerSponsorItemRenderer     `json:"liveChatTickerSponsorItemRenderer"`
	DurationSec                           string                                `json:"durationSec"`
}

type LiveChatViewerEngagementMessageRenderer struct {
	ID            string  `json:"id"`
	TimestampUsec string  `json:"timestampUsec"`
	Icon          Icon    `json:"icon"`
	Message       Message `json:"message"`
}

type LiveChatTextMessageRenderer struct {
	ID                       string              `json:"id"`
	TimestampUsec            string              `json:"timestampUsec"`
	Message                  Message             `json:"message"`
	AuthorName               AuthorName          `json:"authorName"`
	AuthorPhoto              AuthorPhoto         `json:"authorPhoto"`
	Icon                     Icon                `json:"icon"`
	ContextMenuEndpoint      ContextMenuEndpoint `json:"contextMenuEndpoint"`
	AuthorBadges             []AuthorBadge       `json:"authorBadges"`
	AuthorExternalChannelId  string              `json:"authorExternalChannelId"`
	ContextMenuAccessibility Accessibility       `json:"contextMenuAccessibility"`
	TimestampText            TimestampText       `json:"timestampText"`
}

type LiveChatMembershipItemRenderer struct {
	ID                       string              `json:"id"`
	TimestampUsec            string              `json:"timestampUsec"`
	AuthorExternalChannelId  string              `json:"authorExternalChannelId"`
	HeaderSubtext            HeaderSubtext       `json:"headerSubtext"`
	AuthorName               AuthorName          `json:"authorName"`
	AuthorPhoto              AuthorPhoto         `json:"authorPhoto"`
	AuthorBadges             []AuthorBadge       `json:"authorBadges"`
	ContextMenuEndpoint      ContextMenuEndpoint `json:"contextMenuEndpoint"`
	ContextMenuAccessibility Accessibility       `json:"contextMenuAccessibility"`
}

type LiveChatPaidMessageRenderer struct {
	ID                       string              `json:"id"`
	TimestampUsec            string              `json:"timestampUsec"`
	AuthorName               AuthorName          `json:"authorName"`
	AuthorPhoto              AuthorPhoto         `json:"authorPhoto"`
	PurchaseAmountText       PurchaseAmountText  `json:"purchaseAmountText"`
	Message                  Message             `json:"message"`
	HeaderBackgroundColor    int                 `json:"headerBackgroundColor"`
	HeaderTextColor          int                 `json:"headerTextColor"`
	BodyBackgroundColor      int                 `json:"bodyBackgroundColor"`
	BodyTextColor            int                 `json:"bodyTextColor"`
	AuthorExternalChannelId  string              `json:"authorExternalChannelId"`
	AuthorNameTextColor      int                 `json:"authorNameTextColor"`
	ContextMenuEndpoint      ContextMenuEndpoint `json:"contextMenuEndpoint"`
	TimestampColor           int                 `json:"timestampColor"`
	ContextMenuAccessibility Accessibility       `json:"contextMenuAccessibility"`
	TimestampText            TimestampText       `json:"timestampText"`
	TrackingParams           string              `json:"trackingParams"`
}

type LiveChatPlaceholderItemRenderer struct {
	ID            string `json:"id"`
	TimestampUsec string `json:"timestampUsec"`
}

type LiveChatTickerPaidMessageItemRenderer struct {
	ID                      string           `json:"id"`
	Amount                  Amount           `json:"amount"`
	AmountTextColor         int              `json:"amountTextColor"`
	StartBackgroundColor    int              `json:"startBackgroundColor"`
	EndBackgroundColor      int              `json:"endBackgroundColor"`
	AuthorPhoto             AuthorPhoto      `json:"authorPhoto"`
	DurationSec             int              `json:"durationSec"`
	ShowItemEndpoint        ShowItemEndpoint `json:"showItemEndpoint"`
	AuthorExternalChannelId string           `json:"authorExternalChannelId"`
	FullDurationSec         int              `json:"fullDurationSec"`
	TrackingParams          string           `json:"trackingParams"`
}

type LiveChatTickerSponsorItemRenderer struct {
	ID                      string           `json:"id"`
	DetailText              DetailText       `json:"detailText"`
	DetailTextColor         int              `json:"detailTextColor"`
	StartBackgroundColor    int              `json:"startBackgroundColor"`
	EndBackgroundColor      int              `json:"endBackgroundColor"`
	SponsorPhoto            SponsorPhoto     `json:"sponsorPhoto"`
	DurationSec             int              `json:"durationSec"`
	ShowItemEndpoint        ShowItemEndpoint `json:"showItemEndpoint"`
	AuthorExternalChannelId string           `json:"authorExternalChannelId"`
	FullDurationSec         int              `json:"fullDurationSec"`
}

type DetailText struct {
	Runs []Run `json:"runs"`
}

type SponsorPhoto struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Icon struct {
	IconType string `json:"iconType"`
}

type Message struct {
	Runs []Run `json:"runs"`
}

type HeaderSubtext struct {
	Runs []Run `json:"runs"`
}

type Amount struct {
	SimpleText string `json:"simpleText"`
}

type PurchaseAmountText struct {
	SimpleText string `json:"simpleText"`
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

type ShowItemEndpoint struct {
	ClickTrackingParams      string                   `json:"clickTrackingParams"`
	CommandMetadata          CommandMetadata          `json:"commandMetadata"`
	ShowLiveChatItemEndpoint ShowLiveChatItemEndpoint `json:"showLiveChatItemEndpoint"`
}

type ShowLiveChatItemEndpoint struct {
	Renderer Renderer `json:"renderer"`
}

type Renderer struct {
	LiveChatPaidMessageRenderer    LiveChatPaidMessageRenderer    `json:"liveChatPaidMessageRenderer"`
	LiveChatMembershipItemRenderer LiveChatMembershipItemRenderer `json:"liveChatMembershipItemRenderer"`
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
	Icon            Icon            `json:"icon"`
	CustomThumbnail CustomThumbnail `json:"customThumbnail"`
	Tooltip         string          `json:"tooltip"`
	Accessibility   Accessibility   `json:"accessibility"`
}

type CustomThumbnail struct {
	Thumbnails []Thumbnail `json:"thumbnails"`
}
