package export

import (
	"encoding/base64"
	"fmt"
	"github.com/Nightgunner5/procedural/data"
	"image/png"
	"io"
)

func Export(w io.Writer, world *data.World) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error) // re-panics if r is not an error.
		}
	}()

	handle := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}

	handle(fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>procedural | seed:%d</title>
	<link href="https://netdna.bootstrapcdn.com/twitter-bootstrap/2.3.0/css/bootstrap-combined.min.css" rel="stylesheet">
</head>
<body>
	<div class="container">
		<div class="page-header">
			<h1><tt>procedural</tt> <small title="world seed">%d</small></h1>
		</div>
`, world.Seed, world.Seed))

	for i := uint64(0); i < world.AreaCount; i++ {
		a := world.Area(i)
		handle(fmt.Fprintf(w, `
		<section class="clearfix">
			<h3>Area %d: %s</h3>`, i, a.Name))
		if !a.Generated {
			handle(fmt.Fprintf(w, `
			<p><em>Unexplored</em></p>
		</section>`))
			continue
		}

		handle(fmt.Fprintf(w, `
			<img src="data:image/png;base64,`))
		img := genMap(a)
		imgOut := base64.NewEncoder(base64.StdEncoding, w)
		handle(0, png.Encode(imgOut, img))
		imgOut.Close()
		handle(fmt.Fprintf(w, `" class="pull-right">`))

		handle(fmt.Fprintf(w, `
		</section>`))
	}

	handle(fmt.Fprintf(w, `
	</div>
</body>
</html>
`))

	return
}
