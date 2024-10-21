package asn1ber

import (
	
	"io"
)

type BerObjectDescriptor struct {
	BerGraphicString
}

var objectDescriptorTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, OBJECT_DESCRIPTOR_TAG)

func NewBerObjectDescriptor(value []byte) *BerObjectDescriptor {
	return &BerObjectDescriptor{BerGraphicString: *NewBerGraphicString(value)}
}

func (b *BerObjectDescriptor) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(objectDescriptorTag, reversedWriter, withTagList...)
}

func (b *BerObjectDescriptor) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(objectDescriptorTag, input, withTagList...)
}
func (b *BerObjectDescriptor) GetTag() *BerTag {
	return objectDescriptorTag
}
