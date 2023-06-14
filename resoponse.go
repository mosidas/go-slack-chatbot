package main

type ResponsePostMessage struct {
	Ok      bool    `json:"ok"`
	Channel string  `json:"channel"`
	Ts      string  `json:"ts"`
	Message Message `json:"message"`
	Error   string  `json:"error"`
}

type Message struct {
	BotID      string     `json:"bot_id"`
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	User       string     `json:"user"`
	Ts         string     `json:"ts"`
	Team       string     `json:"team"`
	BotProfile BotProfile `json:"bot_profile"`
}

type BotProfile struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Updated int    `json:"updated"`
	AppID   string `json:"app_id"`
	Icons   Icons  `json:"icons"`
	TeamID  string `json:"team_id"`
}

type Icons struct {
	Image36 string `json:"image_36"`
	Image48 string `json:"image_48"`
	Image72 string `json:"image_72"`
}

type ResponseFileUpload struct {
	Ok   bool `json:"ok"`
	File File `json:"file"`
}

type File struct {
	ID                 string   `json:"id"`
	Created            int64    `json:"created"`
	Timestamp          int64    `json:"timestamp"`
	Name               string   `json:"name"`
	Title              string   `json:"title"`
	Mimetype           string   `json:"mimetype"`
	Filetype           string   `json:"filetype"`
	PrettyType         string   `json:"pretty_type"`
	User               string   `json:"user"`
	Editable           bool     `json:"editable"`
	Size               int      `json:"size"`
	Mode               string   `json:"mode"`
	IsExternal         bool     `json:"is_external"`
	ExternalType       string   `json:"external_type"`
	IsPublic           bool     `json:"is_public"`
	PublicURLShared    bool     `json:"public_url_shared"`
	DisplayAsBot       bool     `json:"display_as_bot"`
	Username           string   `json:"username"`
	URLPrivate         string   `json:"url_private"`
	URLPrivateDownload string   `json:"url_private_download"`
	Thumb64            string   `json:"thumb_64"`
	Thumb80            string   `json:"thumb_80"`
	Thumb360           string   `json:"thumb_360"`
	Thumb360W          int      `json:"thumb_360_w"`
	Thumb360H          int      `json:"thumb_360_h"`
	Thumb480           string   `json:"thumb_480"`
	Thumb480W          int      `json:"thumb_480_w"`
	Thumb480H          int      `json:"thumb_480_h"`
	Thumb160           string   `json:"thumb_160"`
	ImageExifRotation  int      `json:"image_exif_rotation"`
	OriginalW          int      `json:"original_w"`
	OriginalH          int      `json:"original_h"`
	Permalink          string   `json:"permalink"`
	PermalinkPublic    string   `json:"permalink_public"`
	CommentsCount      int      `json:"comments_count"`
	IsStarred          bool     `json:"is_starred"`
	Shares             Shares   `json:"shares"`
	Channels           []string `json:"channels"`
	Groups             []string `json:"groups"`
	Ims                []string `json:"ims"`
	HasRichPreview     bool     `json:"has_rich_preview"`
}

type Shares struct {
	Private map[string][]PrivateShare `json:"private"`
}

type PrivateShare struct {
	ReplyUsers      []string `json:"reply_users"`
	ReplyUsersCount int      `json:"reply_users_count"`
	ReplyCount      int      `json:"reply_count"`
	Ts              string   `json:"ts"`
}
