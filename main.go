package main

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/xerrors"
)

func f(val int64) error {
	log.Printf("val: %d\n", val)
	if val >= 3 && val%2 == 0 {
		log.Printf("%dはerror\n", val)
		// 仮に3以上の偶数だけエラー扱いにする
		return fmt.Errorf("[error: %d]", val)
	}
	return nil
}

func sample(params []int64) error {
	var errs []error
	for i, param := range params {
		err := f(param)
		if err != nil {
			errs = append(errs, xerrors.Errorf("index %dでエラー: %w", i, err))
		}
	}
	fmt.Printf("len(errs): %d\n", len(errs))
	return errors.Join(errs...)
}

func main() {
	params := []int64{1, 2, 3, 4, 5, 6}
	err := sample(params)
	if err != nil {
		if errs, ok := err.(interface{ Unwrap() []error }); ok {
			for _, e := range errs.Unwrap() {
				log.Printf("%+v\n", e)
			}
		}
	}
}
