package err_group

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func Test_errGroup(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()
	g := New(ctx, 3)
	var urls = []string{
		"http://liangyaopei.github.io",
		"http://www.51cto.com",
		"http://www.baidu.com",
	}
	for _, url := range urls {
		url := url
		g.Go(func(ctx context.Context) error {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return err
			}
			resp, err := http.DefaultClient.Do(req)
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
