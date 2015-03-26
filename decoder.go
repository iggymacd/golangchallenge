package drum

import (
	"bytes"
	"encoding/binary"
	//"encoding/hex"
	"fmt"
	"io/ioutil"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	d, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	var p Pattern
	buf := bytes.NewReader(d)
	err = binary.Read(buf, binary.BigEndian, &p)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	/*p := &Pattern{content: `Saved with HW Version: 0.808-alpha
	Tempo: 120
	(0) kick	|x---|x---|x---|x---|
	(1) snare	|----|x---|----|x---|
	(2) clap	|----|x-x-|----|----|
	(3) hh-open	|--x-|--x-|x-x-|--x-|
	(4) hh-close	|x---|x---|----|x--x|
	(5) cowbell	|----|----|--x-|----|
	`, data: d}*/
	return &p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	SPLICE  [6]byte
	Size    int64
	Version [32]byte
}

func (p Pattern) String() string {
	return fmt.Sprint(p.Size) //p.content)
}
