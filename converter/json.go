package converter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/SkobelevIgor/stackexchange-xml-converter/encoders"
)

// WriteBufferSize bytes (8MB)
const WriteBufferSize = 8388608

func convertToJSON(typeName string, xmlFile *os.File, jsonFile *os.File, cfg Config) (total int64, converted int64, err error) {

	iterator := NewIterator(xmlFile)
	w := bufio.NewWriterSize(jsonFile, WriteBufferSize)
	defer w.Flush()

	w.WriteByte('[')

	var iErr error
	for iterator.Next() {
		if total > 0 && iErr == nil {
			w.WriteByte(',')
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

		_, iErr = w.Write(ji)
		if iErr != nil {
			log.Printf("[%s] Error: %s", typeName, iErr)
			continue
		}

		converted++
	}

	w.WriteByte(']')

	return
}

func marshal(enc *encoders.Encoder) ([]byte, error) {
	b := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(b)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(enc)
	return b.Bytes(), err
}
