package upload

type Uploader struct {
	params UploaderParameters
	uploader
}

type UploaderParameters struct {
	// RequiredParameters
	File          interface{}
	TimeStampUnix string
	Signature     string
	// OptionalParameters
	PublicId                string
	Folder                  string
	UseFilename             bool
	UniqueFilename          bool
	FilenameOverride        string
	ResourceType            string
	Type                    string
	AccessControl           []string
	AccessMode              string
	DiscardOriginalFilename bool
	Overwrite               bool
	// -Resource Data
	Tags                  []string
	Context               []string
	Metadata              []string
	Colors                bool
	Faces                 bool
	QualityAnalysis       bool
	AccessibilityAnalysis bool
	CinemagraphAnalysis   bool
	ImageMetadata         bool
	Phash                 bool
	ResponsiveBreakpoints []string
	AutoTagging           float64
	Categorization        string
	Detection             string
	OCR                   string
	Exif                  bool
	// -Manipulations
	Eager                string
	EagerAsync           bool
	EagerNotificationUrl string
	Transformation       string
	Format               string
	CustomCoordinates    string
	FaceCoordinates      string
	BackgroundRemoval    string
	RawConvert           string
	// -Additional Options
	AllowedFormats    string
	Async             bool
	Backup            bool
	Callback          string
	Eval              string
	Headers           string
	Invalidate        bool
	Moderation        string
	NotificationUrl   string
	Proxy             string
	ReturnDeleteToken bool
}

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

type uploader interface {
	UploadMedia(UploaderParameters) UploaderResponse
}

func (u Uploader) UploadMedia(params UploaderParameters) UploaderResponse {
	var resp UploaderResponse
	return resp
}
