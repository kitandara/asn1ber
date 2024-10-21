package asn1ber

import (
	"io"
)

type BerTime struct {
	BerVisibleString
}

var timeTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, TIME_TAG)

func NewBerTime(value string) *BerTime {
	return &BerTime{BerVisibleString: *NewBerVisibleString(value)}
}

func (b *BerTime) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(timeTag, reversedWriter, withTagList...)
}

func (b *BerTime) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(timeTag, input, withTagList...)
}
func (b *BerTime) GetTag() *BerTag {
	return timeTag
}
