// Telegram: @MrEsfelurm

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	proxyType, nameFile string
	rd                  = "\033[00;31m"
	gn                  = "\033[00;32m"
	// lgn = "\033[01;32m"
	yw = "\033[01;33m"
	// lrd = "\033[01;31m"
	// be  = "\033[00;34m"
	// pe  = "\033[01;35m"
	cn = "\033[00;36m"
)

func main() {

	fmt.Print("\033[H\033[2J")
	fmt.Printf(`
%s
 _______  _______  _______                   
(  ____ )(  ____ )(  ___  )|\     /||\     /|
| (    )|| (    )|| (   ) |( \   / )( \   / )
| (____)|| (____)|| |   | | \ (_) /  \ (_) / 
|  _____)|     __)| |   | |  ) _ (    \   /  
| (      | (\ (   | |   | | / ( ) \    ) (  Tg: @MrEsfelurm 
| )      | ) \ \__| (___) |( /   \ )   | |   
|/       |/   \__/(_______)|/     \|   \_/`, yw)
	fmt.Printf("\n\n\n%sEnter Type Proxy (1: SOCKS4, 2: SOCKS5, 3: HTTP, 4: HTTPS): %s", cn, gn)
	fmt.Scanln(&proxyType)

	fmt.Printf("\n%sEnter the name of the file to save proxies (e.g., proxies.txt):%s ", cn, gn)
	fmt.Scanln(&nameFile)

	APISProxy(proxyType, nameFile)
}

func APISProxy(proxyType, nameFile string) {
	var urls []string

	switch proxyType {
	case "1":
		urls = []string{
			"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4",
			"https://openproxylist.xyz/socks4.txt",
			"https://proxyspace.pro/socks4.txt",
			"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS4.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks4.txt",
			"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks4.txt",
			"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS4_RAW.txt",
			"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks4.txt",
			"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks4.txt",
			"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks4.txt",
			"https://www.proxy-list.download/api/v1/get?type=socks4",
			"https://www.proxyscan.io/download?type=socks4",
			"https://api.proxyscrape.com/?request=displayproxies&proxytype=socks4&country=all",
			"https://api.openproxylist.xyz/socks4.txt",
		}
	case "2":
		urls = []string{
			"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5&timeout=10000&country=all&simplified=true",
			"https://www.proxy-list.download/api/v1/get?type=socks5",
			"https://www.proxyscan.io/download?type=socks5",
			"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks5.txt",
			"https://raw.githubusercontent.com/hookzof/socks5_list/master/proxy.txt",
			"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks5.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks5.txt",
			"https://api.openproxylist.xyz/socks5.txt",
			"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks5",
			"https://openproxylist.xyz/socks5.txt",
			"https://proxyspace.pro/socks5.txt",
			"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS5.txt",
			"https://raw.githubusercontent.com/manuGMG/proxy-365/main/SOCKS5.txt",
			"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks5.txt",
			"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS5_RAW.txt",
			"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks5.txt",
		}
	case "3":
		urls = []string{
			"https://api.proxyscrape.com/?request=displayproxies&proxytype=http",
			"https://www.proxy-list.download/api/v1/get?type=http",
			"https://www.proxyscan.io/download?type=http",
			"https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/http.txt",
			"https://api.openproxylist.xyz/http.txt",
			"https://raw.githubusercontent.com/shiftytr/proxy-list/master/proxy.txt",
			"http://alexa.lr2b.com/proxylist.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-http.txt",
			"https://raw.githubusercontent.com/clarketm/proxy-list/master/proxy-list-raw.txt",
			"https://raw.githubusercontent.com/sunny9577/proxy-scraper/master/proxies.txt",
			"https://raw.githubusercontent.com/opsxcq/proxy-list/master/list.txt",
			"https://proxy-spider.com/api/proxies.example.txt",
			"https://multiproxy.org/txt_all/proxy.txt",
			"https://raw.githubusercontent.com/roosterkid/openproxylist/main/HTTPS_RAW.txt",
			"https://raw.githubusercontent.com/UserR3X/proxy-list/main/online/http.txt",
			"https://raw.githubusercontent.com/UserR3X/proxy-list/main/online/https.txt",
			"https://api.proxyscrape.com/v2/?request=getproxies&protocol=http",
			"https://openproxylist.xyz/http.txt",
			"https://proxyspace.pro/http.txt",
			"https://proxyspace.pro/https.txt",
			"https://raw.githubusercontent.com/almroot/proxylist/master/list.txt",
			"https://raw.githubusercontent.com/aslisk/proxyhttps/main/https.txt",
			"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/HTTP.txt",
			"https://raw.githubusercontent.com/hendrikbgr/Free-Proxy-Repo/master/proxy_list.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt",
			"https://raw.githubusercontent.com/mertguvencli/http-proxy-list/main/proxy-list/data.txt",
			"https://raw.githubusercontent.com/mmpx12/proxy-list/master/http.txt",
			"https://raw.githubusercontent.com/proxy4parsing/proxy-list/main/http.txt",
			"https://raw.githubusercontent.com/RX4096/proxy-list/main/online/http.txt",
			"https://raw.githubusercontent.com/saisuiu/uiu/main/free.txt",
			"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/http.txt",
			"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/http.txt",
			"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/https.txt",
			"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/http.txt",
			"https://rootjazz.com/proxies/proxies.txt",
			"https://sheesh.rip/http.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-http.txt",
			"https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/archive/txt/proxies-http.txt",
			"https://raw.githubusercontent.com/monosans/proxy-list/main/proxies/http.txt",
			"https://www.proxy-list.download/api/v1/get?type=http",
			"https://www.proxyscan.io/download?type=http",
			"https://api.openproxylist.xyz/http.txt",
			"https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all",
			"https://raw.githubusercontent.com/proxy4parsing/proxy-list/main/http.txt",
			"https://raw.githubusercontent.com/HyperBeats/proxy-list/main/http.txt",
		}
	case "4":
		urls = []string{
			"https://raw.githubusercontent.com/mmpx12/proxy-list/master/https.txt",
			"https://raw.githubusercontent.com/RX4096/proxy-list/main/online/https.txt",
			"https://www.proxy-list.download/api/v1/get?type=https",
			"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/https.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt",
			"http://alexa.lr2b.com/proxylist.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-https.txt",
			"https://raw.githubusercontent.com/jetkai/proxy-list/main/archive/txt/proxies-https.txt",
			"https://raw.githubusercontent.com/roosterkid/openproxylist/main/HTTPS_RAW.txt",
		}

	default:
		fmt.Printf("%sInvalid proxy type", rd)
		return
	}

	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Printf("%sError creating file: %s\n", err, rd)
		return
	}
	defer file.Close()

	proxyCount := 0
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	for _, url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%sError reading response: %s\n", err, rd)
			continue
		}
		proxies := strings.Split(string(body), "\n")
		proxyCount += len(proxies)
		_, err = file.Write(body)
		if err != nil {
			fmt.Printf("%sError writing to file: %s\n", err, rd)
		}
	}

	fmt.Printf("\n\n%s%d%s proxy named %s%s%s was downloaded and saved!\n", yw, proxyCount, gn, yw, nameFile, gn)
}
