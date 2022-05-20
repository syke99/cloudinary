package uploader

// UploaderParameters is a struct that the developer can use to set specific parameters to their upload
// to perform various actions, such as transforming an asset, tagging an asset, adding metadata to an
// asset, and more
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
