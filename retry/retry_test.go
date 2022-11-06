package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_retry(t *testing.T) {
	r := New(ConstantBackoff(2, time.Millisecond), DefaultClassifier{})
	retryErr := errors.New("heheda")
	err := r.Run(func(c context.Context) error {
		fmt.Println("heheda")
		return retryErr
	})
	fmt.Println(err == retryErr)
}
