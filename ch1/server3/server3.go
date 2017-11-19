package main

import (
	"net/http"
	"fmt"
	"log"
	"image/color"
	"time"
	"os"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
)

var seaGreen = color.RGBA{0x54,0xff,0x9f,0xff}
var palette = []color.Color{seaGreen,color.White, color.Black}

const (
	whiteIndex = 1
	blackIndex = 2
	greenIndex = 0
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}
var (
	cycles  = 5.0
	res     = 0.01
	size    = 100
	nframes = 64
	delay   = 8
)
func lissajous(out io.Writer) {

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}


	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}


func handler2(w http.ResponseWriter, r* http.Request){
	fmt.Fprintf(w,"%s %s %s \n",r.Method,r.URL,r.Proto)
	for k,v := range r.Header{
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
	fmt.Fprintf(w,"Host = %q\n",r.Host)
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)
	if err:=r.ParseForm();err!=nil{
		log.Print(err)
	}
	for k,v :=range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}