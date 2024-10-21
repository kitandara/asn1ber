package asn1ber

import (
	"io"
)

type BerGeneralString struct {
	BerOctetString
}

var generalstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, GENERAL_STRING_TAG)

func NewBerGeneralString(value []byte) *BerGeneralString {
	return &BerGeneralString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerGeneralString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(generalstringTag, reversedWriter, withTagList...)
}

func (b *BerGeneralString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(generalstringTag, input, withTagList...)
}
func (b *BerGeneralString) GetTag() *BerTag {
	return generalstringTag
}
