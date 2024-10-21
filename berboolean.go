package asn1ber

import (
	"errors"
	"github.com/kitandara/asn1ber/utils"
	"io"
	"strconv"
)

type BerBoolean struct {
	value bool
}

var boolTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, BOOLEAN_TAG)

func NewBerBoolean(v bool) *BerBoolean {
	return &BerBoolean{value: v}
}

func (b *BerBoolean) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	var withTag bool
	if len(withTagList) > 0 {
		withTag = withTagList[0]
	} else {
		withTag = true
	}
	codeLength := 1
	var bx []byte
	if b.value {
		bx = []byte{0xFF}
	} else {
		bx = []byte{0}
	}

	_, err := reversedWriter.Write(bx)
	if err != nil {
		return codeLength, err
	}
	n, err := EncodeLength(codeLength, reversedWriter)
	codeLength += n
	if err != nil {
		return codeLength, err
	}
	if withTag {
		n, err = boolTag.Encode(reversedWriter)
		codeLength += n
	}
	return codeLength, nil
}

func (b *BerBoolean) Decode(input io.Reader, withTagList ...bool) (int, error) {
	var withTag bool
	if len(withTagList) > 0 {
		withTag = withTagList[0]
	} else {
		withTag = true
	}
	codeLength := 0

	if withTag {
		n, err := boolTag.DecodeAndCheck(input)
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
	} else if berLength.Length != 1 {
		return codeLength, errors.New("Invalid length for boolean type")
	}

	nextByte, err := utils.ReadByte(input)
	if err != nil {
		return codeLength, err
	}
	codeLength++
	b.value = nextByte != 0
	return codeLength, nil
}

func (b *BerBoolean) S() string {
	return strconv.FormatBool(b.value)
}
func (b *BerBoolean) GetTag() *BerTag {
	return withTag
}
