package main

import (
	"errors"
	"flag"
	"go-dnspod/dnspod"
	"log"
	"strings"
	"time"
)

func main() {
	argToken := flag.String("token", "", "dns登录token")
	argDomain := flag.String("domain", "", "要操作的域名")
	argRefresh := flag.Int("refresh", 60, "刷新时间,单位秒")
	flag.Parse()
	if *argToken == "" || *argDomain == "" {
		flag.Usage()
		return
	}
	dns := dnspod.NewDnspod(*argToken)
	for {
		err := refreshDDNS(dns, *argDomain)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("DDNS记录刷新成功!")
		}
		time.Sleep(time.Duration(*argRefresh) * time.Second)
	}
}
func refreshDDNS(d *dnspod.Dnspod, domain string) (err error) {
	sub, root := domainSplit(domain)
	list, err := d.Record.List(root)
	if err != nil {
		return
	}
	myIP, err := d.MyWANIP()
	if err != nil {
		return
	}
	for _, l := range list {
		if l.Type == "A" && l.Name == sub {
			if l.Value == myIP {
				return errors.New("当前外网IP未变动,不需要更新")
			}
			if err = d.Record.DDNS(root, l.ID, dnspod.DDNSOpt{SubDomain: l.Name}); err != nil {
				return
			}
			return nil
		}
	}
	return errors.New("未找到指定的域名记录")
}
func domainSplit(domain string) (sub, root string) {
	tmps := strings.Split(domain, ".")
	if len(tmps) <= 2 {
		return "@", domain
	}
	l := len(tmps)
	for k, v := range tmps {
		if k == l-2 {
			return sub, tmps[l-2] + "." + tmps[l-1]
		}
		if k == 0 {
			sub = v
		} else {
			sub = sub + "." + v
		}
	}
	return
}
