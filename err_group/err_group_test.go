package err_group

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/pkg/errors"
)

func Test_errGroup(t *testing.T) {
	g := New(4)
	var urls = []string{
		"http://liangyaopei.github.io/",
		"http://www.51cto.com/",
		"http://www.baidu.com/",
	}
	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Printf("url=%s, status=%v\n", url, resp.StatusCode)
			return nil
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("err is %v\n", err)
	}
}

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func Test_Err(t *testing.T) {
	err := foo()
	errors.Cause(err)
	fmt.Println(errors.Is(err, sql.ErrNoRows))
	fmt.Printf("%+v\n", err)
}
