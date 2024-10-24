package asn1ber

import (
	"io"
)

// Extends an BERInteger

type BerEnum struct {
	BerInteger
}

var enumTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, ENUMERATED_TAG)

func NewBerEnum(v int64) *BerEnum {
	return &BerEnum{
		BerInteger: *NewBerInteger(v)}
}

func (b *BerEnum) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.encodeUsingTag(enumTag, reversedWriter, withTagList...)
}

func (b *BerEnum) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.decodeUsingTag(enumTag, input, withTagList...)
}

func (b *BerEnum) GetTag() *BerTag {
	return enumTag
}
