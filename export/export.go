package export

import (
	"fmt"
	"github.com/Nightgunner5/procedural/data"
	"io"
	//"reflect"
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
		_ = a
	}

	handle(fmt.Fprintf(w, `
	</div>
</body>
</html>
`))

	return
}
