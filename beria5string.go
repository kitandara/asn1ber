package asn1ber

import (
	"io"
)

type BerIA5String struct {
	BerOctetString
}

var ia5stringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, IA5_STRING_TAG)

func NewBerIA5String(value []byte) *BerIA5String {
	return &BerIA5String{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerIA5String) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(ia5stringTag, reversedWriter, withTagList...)
}

func (b *BerIA5String) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(ia5stringTag, input, withTagList...)
}
func (b *BerIA5String) GetTag() *BerTag {
	return ia5stringTag
}
