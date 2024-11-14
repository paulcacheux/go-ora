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

	var mapping map[int]json.RawMessage
	if err := json.NewDecoder(greader).Decode(&mapping); err != nil {
		panic(err)
	}

	raw := mapping[langID]
	if raw == nil {
		return nil
	}

	var jsc JStringConverter
	if err := json.Unmarshal(raw, &jsc); err != nil {
		panic(err)
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
