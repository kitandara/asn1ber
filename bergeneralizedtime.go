package asn1ber

import (
	
	"io"
)

type BerGeneralizedTime struct {
	BerTime
}

var generalizedTimeTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, GENERALIZED_TIME_TAG)

func NewBerGeneralizedTime(value string) *BerGeneralizedTime {
	return &BerGeneralizedTime{BerTime: *NewBerTime(value)}
}

func (b *BerGeneralizedTime) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(generalizedTimeTag, reversedWriter, withTagList...)
}

func (b *BerGeneralizedTime) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(generalizedTimeTag, input, withTagList...)
}
func (b *BerGeneralizedTime) GetTag() *BerTag {
	return generalizedTimeTag
}
