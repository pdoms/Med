package hl7

import (
	"Med/internal/protocols/mllp"
	"fmt"
	"io"
)

const (
	BUFFERSIZE = 24
)

type Scanner struct {
	r io.Reader
	//segments []byte
	Msg []byte
	//	done     bool
}

func NewMllpScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:   r,
		Msg: make([]byte, BUFFERSIZE),
	}
}

func (s *Scanner) Scan() error {
	buf := make([]byte, 12)

	for {
		len, err := s.r.Read(buf)
		if err != nil {
			fmt.Println("ERROR - server: ", err)
			return err
		}
		for _, b := range buf[:len] {
			if b == mllp.SOB {
				continue
			}
			if b == mllp.EOB {
				return nil
			} else {
				s.Msg = append(s.Msg, b)
			}
		}

	}

}
