package pagination

import (
	"testing"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/pagination"
)

func Test_Pagination(t *testing.T) {
	page := &pagination.Pagination{PageSize: 10}
	page.SetTotal(11)
	debug.Display("", page)
}
