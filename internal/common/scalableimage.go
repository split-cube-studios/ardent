package common

type ScalableImage struct {
	TopLeft     ScalableImageRegion `yaml:"top_left"`
	Top         ScalableImageRegion `yaml:"top"`
	TopRight    ScalableImageRegion `yaml:"top_right"`
	Left        ScalableImageRegion `yaml:"left"`
	Center      ScalableImageRegion `yaml:"center"`
	Right       ScalableImageRegion `yaml:"right"`
	BottomLeft  ScalableImageRegion `yaml:"bottom_left"`
	Bottom      ScalableImageRegion `yaml:"bottom"`
	BottomRight ScalableImageRegion `yaml:"bottom_right"`
}

type ScalableImageRegion struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
	W int `yaml:"w"`
	H int `yaml:"h"`
}
