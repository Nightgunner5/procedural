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

	/*handle(fmt.Fprint(w, "<h3>Noise signature</h3>\n<pre>"))

	noise := reflect.ValueOf(world.Noise)
	for i := 0; i < noise.NumField(); i++ {
		f := noise.Field(i)
		handle(fmt.Fprintf(w, "<strong>%s</strong>\n", noise.Type().Field(i).Name))

		for j := 0; j < f.Len(); j++ {
			handle(fmt.Fprintf(w, "% x\n", f.Index(j).Interface()))
		}
	}
	handle(fmt.Fprintf(w, `</pre>`))*/

	handle(fmt.Fprintf(w, `
	</div>
</body>
</html>
`))

	return
}
