package connectionmanager

import (
	"bufio"
	"fmt"
	"net"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

const msgBufInitSize int = 1

type socketconn struct {
	name        socketconnType
	netConn     net.Conn
	bufReader   *bufio.Reader
	bufWriter   *bufio.Writer
	bufScanner  *bufio.Scanner
	ch          chan msgBuffer
	msgRequest  []*request
	msgResponce []*responce
	msgNotice   []*notice
	stopReader  chan bool
	//	msgRequestMutex  sync.Mutex
	//	msgResponceMutex sync.Mutex
	//	msgNoticeMutex   sync.Mutex
}

//---------------------------------------------
//
//---------------------------------------------
func newSocketConn(sct socketconnType) (sc socketconn) {
	sc = socketconn{
		name:        sct,
		netConn:     nil,
		bufReader:   nil,
		bufWriter:   nil,
		bufScanner:  nil,
		ch:          make(chan msgBuffer),
		msgRequest:  make([]*request, 0, msgBufInitSize),
		msgResponce: make([]*responce, 0, msgBufInitSize),
		msgNotice:   make([]*notice, 0, msgBufInitSize),
		// stopReader:  make(chan bool),
		stopReader: nil,
	}

	return sc
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) setupSocket() (err error) {

	if s.netConn == nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"%s: netConn is nil\n", s.name))
	}

	if s.bufScanner == nil {
		s.bufScanner = bufio.NewScanner(s.netConn)
	}

	if s.bufReader == nil {
		s.bufReader = bufio.NewReader(s.netConn)
	}

	if s.bufWriter == nil {
		s.bufWriter = bufio.NewWriter(s.netConn)
	}

	s.ch = make(chan msgBuffer)

	return nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) close() {

	if s.stopReader != nil {
		s.stopReader <- true
	}

	s.bufReader = nil
	s.bufWriter = nil
	s.bufScanner = nil

	if len(s.msgRequest) > 0 {
		s.msgRequest = s.msgRequest[:0]
	}
	s.msgRequest = nil
	if len(s.msgResponce) > 0 {
		s.msgResponce = s.msgResponce[:0]
	}
	s.msgResponce = nil
	if len(s.msgNotice) > 0 {
		s.msgNotice = s.msgNotice[:0]
	}

}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isSocketClosed() bool {
	return s.netConn == nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isMsgReady() (ret bool) {
	ret = s.isResponceMsgReady() ||
		s.isRequestMsgReady() ||
		s.isNoticeMsgReady()
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isRequestMsgReady() (ret bool) {
	//	s.msgRequestMutex.Lock()
	ret = len(s.msgRequest) > 0
	//	s.msgRequestMutex.Unlock()
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isResponceMsgReady() (ret bool) {
	//	s.msgResponceMutex.Lock()
	ret = len(s.msgResponce) > 0
	//	s.msgResponceMutex.Unlock()
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isNoticeMsgReady() (ret bool) {
	//	s.msgNoticeMutex.Lock()
	ret = len(s.msgNotice) > 0
	//	s.msgNoticeMutex.Unlock()
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) getRequestMsg() (ret *request) {
	if s.isRequestMsgReady() {
		//		s.msgRequestMutex.Lock()
		ret = s.msgRequest[0]
		s.msgRequest = append(s.msgRequest[1:])
		//		s.msgRequestMutex.Unlock()
	}
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) getResponceMsg() (ret *responce) {
	if s.isResponceMsgReady() {
		//		s.msgResponceMutex.Lock()
		ret = s.msgResponce[0]
		s.msgResponce = append(s.msgResponce[1:])
		//		s.msgResponceMutex.Unlock()
	}
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) getNoticeMsg() (ret *notice) {
	if s.isNoticeMsgReady() {
		//		s.msgNoticeMutex.Lock()
		ret = s.msgNotice[0]
		s.msgNotice = append(s.msgNotice[1:])
		//		s.msgNoticeMutex.Unlock()
	}
	return ret
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) addRequestMsg(r *request) {
	//		s.msgRequestMutex.Lock()
	s.msgRequest = append(s.msgRequest, r)
	//		s.msgRequestMutex.Unlock()
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) addResponceMsg(r *responce) {
	//		s.msgResponceMutex.Lock()
	s.msgResponce = append(s.msgResponce, r)
	//		s.msgResponceMutex.Unlock()
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) addNoticeMsg(n *notice) {
	//		s.msgNoticeMutex.Lock()
	s.msgNotice = append(s.msgNotice, n)
	//		s.msgNoticeMutex.Unlock()
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) clearMsg() {
	s.msgRequest = nil
	s.msgResponce = nil
	s.msgNotice = nil
	panic(fmt.Sprintf(lumerinlib.FileLine() + "\n"))
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) runSocketReader() {

	if s.name == "" {
		panic(fmt.Sprintf(lumerinlib.FileLine() + " Error no socketconn name\n"))
	}

	fmt.Printf(lumerinlib.Funcname()+" Preparing %s Socket for reading\n", s.name)

	if s.netConn == nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"%s: netConn == nil\n", s.name))
	}

	s.stopReader = make(chan bool, 0)

	go func(s *socketconn) {

		fmt.Printf(lumerinlib.Funcname()+" %s Socket\n", s.name)

		defer func() {
			fmt.Printf(lumerinlib.Funcname()+" Closing %s Socket\n", s.name)
			s.netConn.Close()
			close(s.stopReader)
		}()

	loop:
		for {
			switch {
			case <-s.stopReader:
				fmt.Printf(lumerinlib.Funcname()+" got stopReader for: %s\n", s.name)
				break loop
			default:
			}

			fmt.Printf(lumerinlib.Funcname()+" Scoket Scan %s\n", s.name)

			if !s.bufScanner.Scan() {
				err := s.bufScanner.Err()

				if err == nil {
					fmt.Printf(lumerinlib.FileLine()+" %s Socket Closed TCP connection\n", s.name)
				} else {
					fmt.Printf(lumerinlib.FileLine()+" Error recieved on %s TCP connection: %s\n", s.name, err)
				}

				break loop
			}

			err := s.bufScanner.Err()

			if err != nil {
				fmt.Printf(lumerinlib.FileLine()+" Error recieved on %s TCP connection: %s\n", s.name, err)
				break loop
			}

			buf := s.bufScanner.Bytes()

			if len(buf) > 0 {
				s.ch <- buf
				fmt.Printf(lumerinlib.FileLine()+"Read %s: %s\n", s.name, buf)
			} else {
				fmt.Printf(lumerinlib.FileLine()+"Warning: Read %s Zero Len, skipping\n", s.name)
			}
		}

		fmt.Printf(lumerinlib.Funcname()+" Exiting %s Socket Scan()...\n", s.name)

	}(s)

}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) dial(proto string, host string, port string) error {

	c, err := net.Dial(proto, host+":"+port)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"net.Dial() failed:%s\n", err)
		return err
	}

	s.netConn = c

	err = s.setupSocket()
	if err == nil {
		s.runSocketReader()
	}

	return err
}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) send(b []byte) error {

	// Add a newline
	fmt.Printf("Socket send(): %s\n", b)

	// Stratum Protocol uses a "\n" as delimiter, it will not process until it sees this
	// The JSON package does not add one, so it is added here.
	b = append(b, "\n"...)

	msgLen := len(b)

	len, err := s.bufWriter.Write(b)
	if err != nil {
		return err
	} else if msgLen != len {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"msgLen:%d not eq len:%d\n", msgLen, len))
	}

	return s.bufWriter.Flush()

}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) sendRequest(request *request) error {

	r, err := request.createRequestMsg()
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"json.Marshal errored:%s\n", err))
	}

	return s.send(r)

}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) sendResponce(responce *responce) error {

	r, err := responce.createResponceMsg()
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"json.Marshal errored:%s\n", err))
	}

	return s.send(r)
}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) sendNotice(notice *notice) (err error) {

	var n []byte

	n, err = notice.createNoticeMsg()
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"json.Marshal errored:%s\n", err))
	}

	return s.send(n)
}
