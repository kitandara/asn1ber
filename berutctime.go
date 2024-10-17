package asn1ber

import (
	"io"
)

type BerUtcTime struct {
	BerTime
}

var utcTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, UTC_TIME_TAG)

func NewBerUtcTime(value string) *BerUtcTime {
	return &BerUtcTime{BerTime: *NewBerTime(value)}
}

func (b *BerUtcTime) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(utcTag, reversedWriter, withTagList...)
}

func (b *BerUtcTime) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(utcTag, input, withTagList...)
}
