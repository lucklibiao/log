package log

import (
	"os"
)

type LogFile struct {
	fd       *os.File //point to logFile
	filename string   //file name
	maxLines int64    //max Lines
	maxBytes int64    //max byes
	curLines int64    // current logfile Lines
	curBytes int64    // current logfile bytes
}

func NewFile(logname string) *LogFile {

	f := &LogFile{filename: logname}
	f.open()
	return f

}

//如果文件存在，则打开，打开方式为append；文件不存在，则创建文件，并且创建的权限为rw-r--r--
func (f *LogFile) open() (err error) {
	f.fd, err = os.OpenFile(f.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return
}

func (f *LogFile) SetMaxLines(maxLines int64) {
	f.maxLines = maxLines
}

func (f *LogFile) SetMaxByte(maxBytes int64) {
	f.maxBytes = maxBytes
}

func (f *LogFile) Close() error {
	return f.fd.Close()
}

//未实现限制文件大小的功能，后面补充
func (f *LogFile) Writer(content []byte) (int,error){
	if f.maxBytes < 0 || f.maxLines < 0 {
		panic("maxBytes or maxLines is less 0")
	}

	writeBytes,err:=f.fd.Write(content)
	return writeBytes,err
}

