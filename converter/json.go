package converter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/SkobelevIgor/stackexchange-xml-converter/encoders"
	"log"
	"os"
	"regexp"
	"strings"
)

// WriteBufferSize bytes (8MB)
const WriteBufferSize = 8388608

func convertToJSON(typeName string, xmlFile *os.File, jsonFile *os.File, cfg Config, tags []string, oneLine bool, filterExactMatch bool) (total int64, converted int64, err error) {
	iterator := NewIterator(xmlFile)
	w := bufio.NewWriterSize(jsonFile, WriteBufferSize)
	defer w.Flush()
	var iErr error
	for iterator.Next() {
		if total > 0 && iErr == nil {
			if oneLine == false {
				w.WriteByte(',')
			}
		}
		total++
		encoder, _ := encoders.NewEncoder(typeName)
		iErr = iterator.Decode(&encoder)
		if iErr != nil {
			log.Printf("[%s] Error: %s", typeName, iErr)
			continue
		}
		if cfg.SkipHTMLDecoding {
			encoder.EscapeFields()
		}
		ji, iErr := marshal(&encoder)
		if iErr != nil {
			log.Printf("[%s] Error: %s", typeName, iErr)
			continue
		}

		// We might want to exclude this post
		// if we filter for tags this needs to be initialized with false
		var ignorePost bool
		ignorePost = len(tags) > 0
		if ignorePost {
			// Look for tags in post
			re := regexp.MustCompile(`"Tags":"(<.+>)+"`)
			tagsVar := re.FindString(string(ji))
			// check if we the post is tagged with a label we are intested in
			for i := 0; i < len(tags) && ignorePost; i++ {
				if filterExactMatch {
					ignorePost = !(strings.Contains(tagsVar, "<" + tags[i] + ">"))
				} else {
					ignorePost = !(strings.Contains(tagsVar, tags[i]))
				}
			}
		}
		if !(ignorePost) {
			// log.Printf("test: %s", string(ji))
			_, iErr = w.Write(ji)
			if iErr != nil {
				log.Printf("[%s] Error: %s", typeName, iErr)
				continue
			}
		}
		converted++
	}
	if oneLine == false {
		w.WriteByte(']')
	}
	return
}

func marshal(enc *encoders.Encoder) ([]byte, error) {
	b := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(b)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(enc)
	return b.Bytes(), err
}
