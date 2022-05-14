package transformations

type Effect struct {
	AccelerationPercentage uint
	AdvRedEye              bool
	AntiRemoval            int
	Art                    string
	AutoBrightness         int
	AutoColor              int
	AutoContrast           int
	AutoColorblind         string
	BackgroundRemoval      BackgroundRemoval
	BlackWhite             int
	Blue                   int
	Blur                   int
	BlurFaces              int
	// BlurRegion
	Boomerang     bool
	Brightness    int
	BrightnessHsb int
	Cartoonify    Cartoonify
	Colorize      int
	Contrast      Contrast
	// Cutout
	Deshake int
	// Displace
	Distort         Distort
	Fade            int
	FillLight       FillLight
	Gamma           int
	GradientFade    GradientFade
	Grayscale       bool
	Green           int
	Hue             int
	Improve         Improve
	Loop            int
	MakeTransparent int
	// Mask             bool
	Morphology       Morphology
	Multiply         bool
	Negate           bool
	Noise            int
	OilPaint         int
	OpacityThreshold int
	OrderedDither    int
	Outline          Outline
	Overlay          bool
	Pixelate         int
	PixelateFaces    int
	// PixelateRegion
	Preview     Preview
	ProgressBar ProgressBar
	// Recolor
	Red          int
	Redeye       bool
	ReplaceColor ReplaceColor
	Reverse      bool
	Saturation   int
	Sepia        int
	Screen       bool
	// Shadow
	Sharpen            int
	Shear              Shear
	SimulateColorbling string
	StyleTransfer      StyleTransfer
	Theme              Theme
	Tint               Tint
	Transition         bool
	Trim               Trim
	UnsharpMask        int
	Vectorize          Vectorize
	Vibrance           int
	ViesusCorrect      ViesusCorrect
	Vignette           int
	Volume             string
	ZoomPan            ZoomPan
}

type BackgroundRemoval struct {
	Screen        bool
	ColorToRemove string
}

type Cartoonify struct {
	LineStrength   int
	ColorReduction string
}

type Contrast struct {
	Level        int
	FunctionType string
}

type Distort struct {
	Degrees float32
	X1      int
	X2      int
	X3      int
	X4      int
	Y1      float32
	Y2      float32
	Y3      float32
	Y4      float32
}

type FillLight struct {
	Blend int
	Bias  int
}

type GradientFade struct {
	Type     string
	Strength int
}

type Improve struct {
	Mode  string
	Blend int
}

type Morphology struct {
	Method     string
	Iterations int
	Kernel     string
	Radius     float32
}

type Outline struct {
	Mode  string
	Width int
	Blur  int
}

type Preview struct {
	Duration           float32
	MaxSegments        int
	MixSegmentDuration float32
}

type ProgressBar struct {
	BarType  string
	BarColor string
	Width    int
}

type ReplaceColor struct {
	FromColor string
	Tolerance int
	ToColor   string
}

type Shear struct {
	XSkew float32
	YSkew float32
}

type StyleTransfer struct {
	StyleStrength int
	PreserveColor bool
}

type Theme struct {
	Color            string
	Photosensitivity int
}

type Tint struct {
	Equalize              bool
	Amount                float32
	ColorOneToTen         []string
	ColorPositionOneToTen []string
}

type Trim struct {
	Tolerance    int
	ColorOveride string
}

type Vectorize struct {
	Colors    int
	Detail    float32
	Despeckle float32
	paths     int
	corners   int
}

type ViesusCorrect struct {
	NoRedEye        bool
	SkinSaturation  bool
	SaturationLevel int
}

type ZoomPan struct {
	Mode      string
	MaxZoom   float32
	Duration  float32
	FrameRate int
	Zoom      float32
	XPosition float32
	YPosition float32
}
