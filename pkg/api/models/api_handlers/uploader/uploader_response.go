package uploader

type UploaderResponse struct {
	AssetId          string          `json:"asset_id"`
	PublicId         string          `json:"public_id"`
	Version          int             `json:"version"`
	VersionId        string          `json:"version_id"`
	Signature        string          `json:"signature"`
	Width            string          `json:"width"`
	Height           string          `json:"height"`
	Format           string          `json:"format"`
	ResourceType     string          `json:"resource_type"`
	CreatedAt        string          `json:"created_at"`
	Tags             []responseTags  `json:"tags"`
	Pages            int             `json:"pages"`
	Bytes            int             `json:"bytes"`
	Type             string          `json:"type"`
	ETag             string          `json:"etag"`
	Placeholder      bool            `json:"placeholder"`
	Url              string          `json:"url"`
	SecureUrl        string          `json:"secure_url"`
	AccessMode       string          `json:"access_mode"`
	OriginalFilename string          `json:"original_filename"`
	Eager            []responseEager `json:"eager"`
}

type responseTags struct {
	Tag string `json:"tag"`
}

type responseEager struct {
	Transformation string `json:"transformation"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Url            string `json:"url"`
	SecureUrl      string `json:"secure_url"`
}
