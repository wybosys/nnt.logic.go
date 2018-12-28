package json

import (
	"github.com/bitly/go-simplejson"
	"log"
	"nnt/core/fs"
	"nnt/core/kernel"
)

func ReadFile(path string) (*kernel.JsonObject, error) {
	content, err := fs.FileGetContents(path)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	obj, err := simplejson.NewJson(content)
	if err != nil {
		return nil, err
	}
	return &kernel.JsonObject{obj}, nil
}
