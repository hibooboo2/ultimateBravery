package lolapi

type Image struct {
	Full  string
	Group string
	Picture string
}

func (theImage *Image) Init() *Image {
	theImage.Picture = DATA_DRAGON + VERSION + "/img/" + theImage.Group + "/" + theImage.Full
	return theImage
}
