package resources

// BaseUrl is the base url for all Cloudinary requests
const BaseUrl = "https://res.cloudinary.com/"

// DeliveryType constants
const (
	Upload          string = "upload"
	Private         string = "private"
	Authenticated   string = "authenticated"
	List            string = "list"
	Fetch           string = "fetch"
	Facebook        string = "facebook"
	Twitter         string = "twitter"
	TwitterName     string = "twitter_name"
	Gravatar        string = "gravatar"
	YouTube         string = "youtube"
	Hulu            string = "hulu"
	Vimeo           string = "vimeo"
	Animoto         string = "animoto"
	WorldStarHipHop string = "worldstarthiphop"
	DailyMotion     string = "dailymotion"
	Multi           string = "multi"
	Delivery        string = "delivery"
)

// Angle constants
const (
	AutoRight = "auto_right"
	AutoLeft  = "auto_left"
	VFlip     = "vflip"
	Hflip     = "hflip"
	Ignore    = "ignore"
)

// AudioCodec constants
const (
	None   = "none"
	Aac    = "aac"
	Vorbis = "vorbis"
	Mp3    = "mp3"
	Opus   = "opus"
)

// AudioFrequencies constants aac
const (
	Aac1  = 96000
	Aac2  = 88200
	Aac3  = 64000
	Aac4  = 48000
	Aac5  = 44100
	Aac6  = 32000
	Aac7  = 24000
	Aac8  = 22050
	Aac9  = 16000
	Aac10 = 12000
	Aac11 = 11025
	Aac12 = 8000
	Aac13 = 7350
)

// AudioFrequencies constants libfdk_aac
const (
	LibfdkAac1  = 96000
	LibfdkAac2  = 88200
	libfdkAac3  = 64000
	libfdkAac4  = 48000
	libfdkAac5  = 44100
	libfdkAac6  = 32000
	libfdkAac7  = 24000
	libfdkAac8  = 22050
	libfdkAac9  = 16000
	libfdkAac10 = 12000
	libfdkAac11 = 11025
	libfdkAac12 = 8000
	libfdkAac13 = 0
)

// AudioFrequencies constants mp3
const (
	Mp3One   = 44100
	Mp3Two   = 48000
	Mp3Three = 32000
	Mp3Four  = 22050
	Mp3Five  = 24000
	Mp3Six   = 16000
	Mp3Seven = 11025
	Mp3Eight = 12000
	Mp3Nine  = 8000
	Mp3Ten   = 0
)

// AudioFrequencies constants opus
const (
	Opus1 = 48000
	Opus2 = 24000
	Opus3 = 16000
	Opus4 = 12000
	Opus5 = 8000
	Opus6 = 0
)

// AudioFrequencies constants pcm
const (
	Pcm1 = 48000
	Pcm2 = 96000
	Pcm3 = 0
)

// BackgroundAuto constants mode
const (
	Border                      = "border"
	Predominant                 = "predominant"
	BorderContrast              = "border_contrast"
	PredominantContrast         = "predominant_contrast"
	PredominantGradient         = "predominant_gradient"
	PredominantGradientContrast = "predominant_gradient_contrast"
	BorderGradient              = "border_gradient"
	BorderGradientContrast      = "border_gradient_contrast"
)

// BackgroundAuto constants direction
const (
	Horizontal   = "horizontal"
	Vertical     = "vertical"
	DiagonalDesc = "diagonal_desc"
	DiagonalAsc  = "diagonal_asc"
)

// ColorSpace constants
const (
	Srgb          = "sRGB"
	TinySrgb      = "tinysrgb"
	Cmyk          = "cmyk"
	NoCmyk        = "no_cmyk"
	KeepCmyk      = "keep_cmyk"
	SrgbTrueColor = "srgb:truecolor"
	CsIcc         = "cs_icc:" //HAS to have a [public_id] appened to be valid to be used
	Copy          = "copy"
)

// CustomFunctions constants
const (
	Remote = "remote"
	Wasm   = "wasm"
)
