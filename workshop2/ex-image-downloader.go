package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO(limouren): rewrite main to download images in parallel
func main() {
	for i, url := range imageURLs {
		downloadURL(url, fmt.Sprintf("download/%02d.jpg", i))
	}
}

func downloadURL(url string, toPath string) {
	fmt.Printf("Downloading %s to %s\n", url, toPath)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET %s: %v\n", url, err)
		return
	}

	f, err := os.Create(toPath)
	if err != nil {
		fmt.Printf("failed to open %s: %v\n", toPath, err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Printf("failed to copy image to file: %v", err)
		return
	}

	fmt.Printf("URL %s downloaded to %s\n", url, toPath)
}

var imageURLs = []string{
	"https://scontent-hkg3-1.xx.fbcdn.net/hphotos-xpt1/l/t31.0-8/11844997_730981923698280_6614663320797777175_o.jpg",
	"https://scontent-hkg3-1.xx.fbcdn.net/hphotos-xta1/t31.0-8/11117558_723827664413706_7663010234307616780_o.jpg",
	"https://fbcdn-sphotos-e-a.akamaihd.net/hphotos-ak-xta1/t31.0-8/p960x960/11713859_717527398377066_8090133150258345507_o.jpg",
	"https://fbcdn-sphotos-a-a.akamaihd.net/hphotos-ak-xpa1/t31.0-8/11722081_716359261827213_8249525903060681735_o.jpg",
	"https://scontent-hkg3-1.xx.fbcdn.net/hphotos-xta1/t31.0-8/11722065_715013678628438_2938073386510944038_o.jpg",
	"https://fbcdn-sphotos-d-a.akamaihd.net/hphotos-ak-xtp1/t31.0-8/p960x960/11154659_684266621703144_3309360697637009928_o.jpg",
	"https://scontent-hkg3-1.xx.fbcdn.net/hphotos-xaf1/v/t1.0-9/10968498_642196542576819_5966690685400541606_n.jpg?oh=9534ceda1aa9bfa920ee74ab7ec036f3&oe=56407C9F",
	"https://scontent-hkg3-1.xx.fbcdn.net/hphotos-xpf1/t31.0-8/10286961_493734664089675_3956332165882212460_o.jpg",
	"https://fbcdn-sphotos-d-a.akamaihd.net/hphotos-ak-xpf1/t31.0-8/10295106_492063314256810_5784114635320926359_o.jpg",
	"https://fbcdn-sphotos-f-a.akamaihd.net/hphotos-ak-xpt1/v/t1.0-9/10933949_635114886618318_6317231577225402249_n.jpg?oh=030656f176ce3d93d84e59b7eab4cc10&oe=567B50D1&__gda__=1447344597_5efea69004d580f2844f18e1964b3f3a",
}
