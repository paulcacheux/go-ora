package converters

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
)

//go:embed conv_data.json.gz
var convData []byte

type JStringConverter struct {
	LangID    int
	CharWidth int
	EReplace  int
	DReplace  int
	DBuffer   []int
	DBuffer2  []int
	EBuffer   map[int]int
	Leading   map[int]int
}

func NewStringConverter(langID int) IStringConverter {
	greader, _ := gzip.NewReader(bytes.NewReader(convData))
	defer greader.Close()

	var mapping map[int]*JStringConverter
	_ = json.NewDecoder(greader).Decode(&mapping)

	jsc := mapping[langID]
	if jsc == nil {
		return nil
	}

	return &StringConverter{
		LangID:    jsc.LangID,
		CharWidth: jsc.CharWidth,
		eReplace:  jsc.EReplace,
		dReplace:  jsc.DReplace,
		dBuffer:   jsc.DBuffer,
		dBuffer2:  jsc.DBuffer2,
		eBuffer:   jsc.EBuffer,
		leading:   jsc.Leading,
	}
}
