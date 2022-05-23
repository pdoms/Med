package mllp

import (
	"fmt"
	"io"
)

const (
	SOB        = 11
	EOB        = 28
	CR         = 13
	BUFFERSIZE = 4096
)

type MllpScanner struct {
	r   io.Reader
	Msg []byte
}

func NewMllpScanner(r io.Reader) *MllpScanner {
	return &MllpScanner{
		r:   r,
		Msg: make([]byte, BUFFERSIZE),
	}
}

func (s *MllpScanner) Scan() error {
	buf := make([]byte, 12)

	for {
		len, err := s.r.Read(buf)
		if err != nil {
			fmt.Println("ERROR - server: ", err)
			return err
		}
		for _, b := range buf[:len] {
			if b == SOB {
				continue
			}
			if b == EOB {
				return nil
			} else {
				s.Msg = append(s.Msg, b)
			}
		}
	}
}
