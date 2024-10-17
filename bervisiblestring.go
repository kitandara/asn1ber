package asn1ber

import (
	"io"
)

type BerVisibleString struct {
	value []byte
}

var berVisibleStringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, VISIBLE_STRING_TAG)

func NewBerVisibleString(value string) *BerVisibleString {
	return &BerVisibleString{value: []byte(value)}
}

func (b *BerVisibleString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(berVisibleStringTag, reversedWriter, withTagList...)
}

func (b *BerVisibleString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(berVisibleStringTag, input, withTagList...)
}

func (b *BerVisibleString) EncodeUsingTag(tag *BerTag, reversedWriter io.Writer, withTagList ...bool) (int, error) {
	var withTag bool
	if len(withTagList) > 0 {
		withTag = withTagList[0]
	} else {
		withTag = true
	}
	_, err := reversedWriter.Write(b.value)
	if err != nil {
		return 0, err
	}
	codeLength := len(b.value)
	if withTag {
		n, err := tag.Encode(reversedWriter)
		if err != nil {
			return 0, err
		}
		codeLength += n
	}

	return codeLength, nil
}

func (b *BerVisibleString) DecodeUsingTag(tag *BerTag, input io.Reader, withTagList ...bool) (int, error) {
	var withTag bool
	if len(withTagList) > 0 {
		withTag = withTagList[0]
	} else {
		withTag = true
	}
	codeLength := 0
	if withTag {
		n, err := tag.DecodeAndCheck(input)
		codeLength += n
		if err != nil {
			return codeLength, err
		}
	}
	berLength := &BerLength{}
	n, err := berLength.Decode(input)
	codeLength += n
	if err != nil {
		return codeLength, err
	}

	b.value = make([]byte, berLength.Length)
	if berLength.Length != 0 {
		_, err = io.ReadFull(input, b.value)
		if err != nil {
			return codeLength, err
		}
		codeLength += len(b.value)
	}
	return codeLength, nil
}

func (b *BerVisibleString) S() string {
	return string(b.value)
}
