package asn1ber

import (
	"io"
)

type BerDate struct {
	BerTime
}

var dateTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, DATE_TAG)

func NewBerDate(value string) *BerDate {
	return &BerDate{BerTime: *NewBerTime(value)}
}

func (b *BerDate) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(dateTag, reversedWriter, withTagList...)
}

func (b *BerDate) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(dateTag, input, withTagList...)
}
