package util

import (
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPathVariable(req *http.Request, no int) (string, bool) {
	splits := strings.Split(req.URL.EscapedPath(), "/")
	if len(splits) <= no+1 {
		return "", false
	}
	return splits[no+1], true
}

func GetObjectID(id string) primitive.ObjectID {
	res, e := primitive.ObjectIDFromHex(id)
	if e == nil {
		return res
	}
	return primitive.NilObjectID
}

func GetFileExt(fileName string, withDot bool) (string, bool) {
	idx := strings.LastIndex(fileName, ".")
	if idx == -1 {
		return "", false
	}
	if !withDot {
		idx++
	}
	return fileName[idx:], true
}
