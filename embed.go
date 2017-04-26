// Packagee goef helps you embed file/data/assets/resources/binary directly in
// Go code.
package goef

import (
	"encoding/hex"
	"io/ioutil"
)

func FileToStringLiteral(filepath string) (sl string, err error) {
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	sl = hex.EncodeToString(bs)
	return
}
