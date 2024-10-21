package asn1ber

import (
	"io"
)

type BerVideotexString struct {
	BerOctetString
}

var videotexstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, VIDEOTEX_STRING_TAG)

func NewBerVideotexString(value []byte) *BerVideotexString {
	return &BerVideotexString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerVideotexString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(videotexstringTag, reversedWriter, withTagList...)
}

func (b *BerVideotexString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(videotexstringTag, input, withTagList...)
}
func (b *BerVideotexString) GetTag() *BerTag {
	return videotexstringTag
}
