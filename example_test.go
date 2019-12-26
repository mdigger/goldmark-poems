package poems_test

import (
	"log"
	"os"

	poems "github.com/mdigger/goldmark-poems"
	"github.com/yuin/goldmark"
)

func Example() {
	var source = []byte(`
# Sample

    Alas for man! day after day may rise,
      Night may shade his thankless head,
    He sees no God in the bright, morning skies
      He sings no praises from his guarded bed.
`)
	md := goldmark.New(
		goldmark.WithExtensions(poems.Extension))
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// <h1>Sample</h1>
	// <div class="poem">
	// Alas for man! day after day may rise,<br>
	// &nbsp;&nbsp;Night may shade his thankless head,<br>
	// He sees no God in the bright, morning skies<br>
	// &nbsp;&nbsp;He sings no praises from his guarded bed.
	// </div>
}
