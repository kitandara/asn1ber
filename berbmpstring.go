package asn1ber

import (
	"io"
)

type BerBMPString struct {
	BerOctetString
}

var bmpstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, BMP_STRING_TAG)

func NewBerBmpString(value []byte) *BerBMPString {
	return &BerBMPString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerBMPString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(bmpstringTag, reversedWriter, withTagList...)
}

func (b *BerBMPString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(bmpstringTag, input, withTagList...)
}
func (b *BerBMPString) GetTag() *BerTag {
	return bmpstringTag
}
