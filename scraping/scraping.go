package scraping

import (
	"context"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/MontFerret/ferret/pkg/drivers/http"
)

func Get(query string) (result []byte, err error) {
	comp := compiler.New()

	program, err := comp.Compile(query)

	if err != nil {
		return nil, err
	}

	// create a root context
	ctx := context.Background()

	// enable HTML drivers
	// by default, Ferret Runtime does not know about any HTML drivers
	// all HTML manipulations are done via functions from standard library
	// that assume that at least one driver is available

	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:42.0) Gecko/20100101 Firefox/42.0"
	ctx = drivers.WithDynamic(ctx, cdp.NewDriver(cdp.WithUserAgent(userAgent)))
	ctx = drivers.WithStatic(ctx, http.NewDriver(http.WithUserAgent(userAgent)))

	return program.Run(ctx)
}
