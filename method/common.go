//@auther foolbread
//@time 2016-08-30
//@file superlistener/method/common.go
package method

type NotifyMethod interface {
	AddNotify(eventname string, program string)
	DoNotify(eventname string, data []byte)
}
