package images

/*define the image interface*/
import (
	"fmt"
	"image"
)

func images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
