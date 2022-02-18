/**
 * @Time: 2022/2/18 03:14
 * @Author: yt.yin
 */

package tlv

import "go-socket/message/simple"

// TlvMsgI 定义TLV消息模版
type TlvMsgI interface {

	simple.SimpleMsgI

	// GetTag 消息标识
	GetTag()     []byte

	// GetBegin 消息头或开始，常用的有 FFFF,55AA 等
	GetBegin()   []byte

	// GetEnd 消息结束符,常用的有 FF,回车符等
	GetEnd()     []byte

	// GetICV 完整性校验值，常用的有BCC校验 CRC校验
	GetICV()     []byte

	// SetTag 设置消息标识
	SetTag(tag   []byte)

	// SetBegin 设置消息头或开始，常用的有 FFFF,55AA 等
	SetBegin(b   []byte)

	// SetEnd 设置消息结束符,常用的有 FF,回车符等
	SetEnd(e     []byte)

	// SetICV 设置 完整性校验值
	SetICV(c []byte)
}

// TlvMsg TLV 格式报文消息
type TlvMsg struct{

	simple.SimpleMsg

	// 报文标识
	Tag     []byte

	// 消息头或开始，常用的有 FFFF,55AA 等
	Begin   []byte

	// End 消息结束符,常用的有 FF,回车符等
	End     []byte

	// ICV 完整性校验值，常用的有BCC校验 CRC校验
	ICV     []byte
}

// NewTLVMsg new 一个message
func NewTLVMsg(tag,begin,end,icv[]byte,msg simple.SimpleMsg) *TlvMsg {
	return &TlvMsg{
		SimpleMsg: msg,
		Tag:       tag,
		Begin:     begin,
		End:       end,
		ICV:       icv,
	}
}

// GetTag 获取报文标识
func (T *TlvMsg) GetTag() []byte {
	return T.Tag
}

// GetBegin 获取报文开始标识
func (T *TlvMsg) GetBegin() []byte {
	return T.Begin
}

// GetEnd 获取报文结束标识
func (T *TlvMsg) GetEnd() []byte {
	return T.End
}

// GetICV 获取报文校验
func (T *TlvMsg) GetICV() []byte {
	return T.ICV
}

// SetTag 设置报文标识
func (T *TlvMsg) SetTag(tag []byte) {
	T.Tag = tag
}

// SetBegin 设置报文开始标识
func (T *TlvMsg) SetBegin(b []byte) {
	T.Begin = b
}

// SetEnd 设置报文结束标识
func (T *TlvMsg) SetEnd(e []byte) {
	T.End = e
}

// SetICV 设置报文校验
func (T *TlvMsg) SetICV(c []byte) {
	T.ICV = c
}

