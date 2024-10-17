package asn1ber

import (
	
	"io"
)

type BerDateTime struct {
	BerTime
}

var dateTimeTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, DATE_TIME_TAG)

func NewBerDateTime(value string) *BerDateTime {
	return &BerDateTime{BerTime: *NewBerTime(value)}
}

func (b *BerDateTime) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(dateTimeTag, reversedWriter, withTagList...)
}

func (b *BerDateTime) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(dateTimeTag, input, withTagList...)
}
