package progressbar

import (
	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
	"sync"
)

// NewWithWaitGroup ...
func NewWithWaitGroup() (*mpb.Progress, *sync.WaitGroup) {
	wg := &sync.WaitGroup{}
	return mpb.New(
		mpb.WithWaitGroup(wg),
		mpb.WithWidth(100),
	), wg
}

// AddBar ...
func AddBar(p *mpb.Progress, total int, name string) *mpb.Bar {
	return p.Add(
		int64(total),
		mpb.NewBarFiller(
			"◣=♥-◢♥+",
		),
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.CountersNoUnit("%d/%d", decor.WCSyncSpaceR),
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WCSyncSpaceR), "done",
			),
			decor.Elapsed(decor.ET_STYLE_GO, decor.WCSyncSpaceR),
			decor.Percentage(),
		),
	)
}
