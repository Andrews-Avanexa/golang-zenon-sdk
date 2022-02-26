package common

import (
	"fmt"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

func DealWithErr(v interface{}) {
	defer RecoverStack()
	if v != nil {
		panic(v)
	}
}

func RecoverStack() {
	if err := recover(); err != nil {
		var e error
		switch t := err.(type) {
		case error:
			e = errors.WithStack(t)
		case string:
			e = errors.New(t)
		default:
			e = errors.Errorf("unknown type %+v", err)
		}

		log15.Error("panic", "err", err, "withstack", e)
		fmt.Printf("%+v", e)
		panic(err)
	}
}
