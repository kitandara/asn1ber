package asn1ber

import (
	
	"io"
)

type BerDuration struct {
	BerTime
}

var durationTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, DURATION_TAG)

func NewBerDuration(value string) *BerDuration {
	return &BerDuration{BerTime: *NewBerTime(value)}
}

func (b *BerDuration) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(durationTag, reversedWriter, withTagList...)
}

func (b *BerDuration) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(durationTag, input, withTagList...)
}
func (b *BerDuration) GetTag() *BerTag {
	return durationTag
}
