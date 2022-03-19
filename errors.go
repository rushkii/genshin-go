package genshin

import "fmt"

func DataNotPublic(msg string, v ...interface{}) error {
	return fmt.Errorf("DataNotPublicErr: "+msg, v...)
}
