package main

import (
	"log"

	"github.com/yuanyuexiang/goav/avcodec"
	"github.com/yuanyuexiang/goav/avdevice"
	"github.com/yuanyuexiang/goav/avfilter"
	"github.com/yuanyuexiang/goav/avformat"
	"github.com/yuanyuexiang/goav/avutil"
	"github.com/yuanyuexiang/goav/swresample"
	"github.com/yuanyuexiang/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
