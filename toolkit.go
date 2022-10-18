package main

import "fmt"
import "net"
import "log"
import "time"



func isPrivate() {
	ip := ""
	fmt.Println("Verify is IP is private")
	fmt.Println("Please enter an IP adress :")
	fmt.Scan(&ip)
	ip_fin := net.ParseIP(ip)
	fmt.Println(ip_fin.IsPrivate())
}

func findIpWithDomainName() {
	domain := ""
	fmt.Println("Please choose a domain name :")
	fmt.Scan(&domain)
	ip_tab, err := net.LookupIP(domain)

	if err != nil {
		log.Fatal("IP lookup failed. Error is: ", err)
	} else {
		for _, ip := range ip_tab {
			fmt.Println(ip)
		}

	}
}

func isOnlineWebSite() {
	domain := ""
	fmt.Println("Please select a domain name without http:// :")
	fmt.Scan(&domain)
	fmt.Println("Please select a valid port :")
	port := ""
	fmt.Scan(&port)
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", domain+":"+port, timeout)
    if err != nil {
        fmt.Printf("%s %s %s\n", domain, "not responding", err.Error())
    } else {
        fmt.Printf("%s %s %s\n", domain, "responding on port:", port)
    }
}

func isLocalhostOnline() {
	ip := "127.0.0.1"
	port := "80"
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", ip+":"+port, timeout)
	if err == nil {
		fmt.Printf("%s %s %s\n", ip, "responding on port:", port)
	} else{
		fmt.Printf("%s %s %s\n", ip, "not responding", err.Error())
	}

}







func menu() {
	choice := ""
	tab := [4]string{}
	tab[0] = "Verify is IP is private = 1"
	tab[1] = "Find an IP with a domain name = 2"
	tab[2] = "Verify is a website is online = 3"
	tab[3] = "Verify if your localhost is online = 4"
	for _, value := range tab {
		fmt.Println("\n"+value)
	}
	fmt.Println("Wich service do you want to use :")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		isPrivate()
		menu()
		break
	case "2":
		findIpWithDomainName()
		menu()
		break
	case "3":
		isOnlineWebSite()
		menu()
		break
	case "4":
		isLocalhostOnline()
		menu()
		break
	};
}






func main() {
menu()
}