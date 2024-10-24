package asn1ber

import (
	"io"
)

type BerTimeOfDay struct {
	BerTime
}

var timeOfDayTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, TIME_OF_DAY_TAG)

func NewBerTimeOfDay(value string) *BerTimeOfDay {
	return &BerTimeOfDay{BerTime: *NewBerTime(value)}
}

func (b *BerTimeOfDay) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(timeOfDayTag, reversedWriter, withTagList...)
}

func (b *BerTimeOfDay) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(timeOfDayTag, input, withTagList...)
}
func (b *BerTimeOfDay) GetTag() *BerTag {
	return timeOfDayTag
}
