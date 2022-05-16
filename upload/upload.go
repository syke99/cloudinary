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

type uploader interface {
	NewUploader(UploaderParameters) Uploader
}

func (u Uploader) NewUploader(params UploaderParameters) Uploader {
	u.params = params
	return u
}
