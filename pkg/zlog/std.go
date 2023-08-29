package zlog

import "os"

type StdWriteSyncer struct {
	file *os.File
}

func (mws StdWriteSyncer) Write(p []byte) (n int, err error) {
	return mws.file.Write(p)
}
func (mws StdWriteSyncer) Sync() error {
	return nil
}

func NewStdWriteSyncer() *StdWriteSyncer {
	return &StdWriteSyncer{os.Stdout}
}