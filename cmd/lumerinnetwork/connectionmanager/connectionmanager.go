package connectionmanager

import (
	"context"
)

//
// Listen Struct for new SRC connections coming in
//
type ConnectionListenStruct struct {
	listen LumerinListenStruct
	context.Context ctx
	cancel func()
}

//
// Struct for existing SRC connections and the associated outgoing DST connections
type ConnectionStruct struct {
	src LumerinSocketStruct
	dst []LumerinSocketStruct
	context.Context ctx
	cancel func()
}

//
//
//
func Listen(ctx context.Context, addr string) (cls *ConnectionListenStruct, e error) {

	return cls, e
}

//
//
//
func (cls *ConnectionListenStruct) Accept() (cs ConnectionStruct, e error) {

	return cs, e
}

//
//
//
func (cls *ConnectionListenStruct) Close() (e error) {

	return e
}

//
//
//
func (cls *ConnectionListenStruct) newConnectionStruct(lss LumerinSocketStruct) (cs *ConnectionStruct) {

	ctx, cancel := context.WithCancel(cls.ctx)

	cs = &ConnectionStruct{
		src: lss,
		dst: make(LumerinSocketStruct,0,8)
		ctx: ctx,
		cancel: cancel,
	}
}

//
//
//
func (cs *ConnectionStruct)Dial()(e error){

}


//
//
//
func (cs *ConnectionStruct)ReDial()(e error){

}

//
//
//
func (cs *ConnectionStruct)Close(){

}

//
//
//
func (cs *ConnectionStruct)SetRoute( index int)(e error){

}

//
//
//
func (cs *ConnectionStruct)ReadReady()(r bool){

}

//
//
//
func (cs *ConnectionStruct)SrcReadReady()(r bool){

}

//
//
//
func (cs *ConnectionStruct)SrcRead(buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)SrcWrite(buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)SrcClose(buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)DstReadReady()(r bool){

}

//
//
//
func (cs *ConnectionStruct)DstRead(buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)DstWrite(buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)DstClose(buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)IdxReadReady(idx int)(r bool){

}
//
//
//
func (cs *ConnectionStruct)IdxRead(idx int, buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)IdxWrite(idx int, buf []byte)(count int, e error){

}

//
//
//
func (cs *ConnectionStruct)IdxClose(idx int, buf []byte)(count int, e error){

}

func read(lss *LumerinSocketStruct, buf []byte)(count int, e error){

}