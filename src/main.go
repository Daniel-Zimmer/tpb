package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"net/url"
	"os"
	"pterm"
	"strconv"
	"strings"
	"time"
	"tpb/src/api"
)

func main() {

	search := url.QueryEscape(strings.Join(os.Args[1:], " "))

	if len(os.Args) == 1 {
		fmt.Printf(
			"Usage: tpb [search term] |" +
				"\n       tpb [page id]\n")
		os.Exit(1)
	}

	_, err := strconv.Atoi(os.Args[1])
	if len(os.Args) == 2 && err == nil{
		// PAGE
		response := api.Page(os.Args[1])

		size := humanize.Bytes(uint64(response.Size))

		fmt.Printf("\n%s\n\n", response.Name)
		fmt.Printf("Type:   %d\n", response.Category)
		fmt.Printf("Files:  %d\n", response.NumFiles)
		fmt.Printf("Size:   %s\n\n", size)
		fmt.Printf("Uploaded:  %s\n", time.Unix(response.Added, 0).Format("2006-02-01"))
		fmt.Printf("By:        %s\n\n", response.Username)
		fmt.Printf("Seeders:    %d\n", response.Seeders)
		fmt.Printf("Leechers:   %d\n", response.Leechers)
		fmt.Printf("Info Hash:  %s\n\n", response.InfoHash)

		fmt.Println(" ================")
		fmt.Printf("%s\n", response.Descr)
		fmt.Println(" ================")

		fmt.Printf("\n%s\n\n",
			"magnet:?xt=urn:btih:" +
			response.InfoHash +
			"&dn=" +
			url.QueryEscape(response.Name),
		)

	} else {
		//SEARCH
		response := api.Search(search)

		for i := 0; i < len(response); i++ {
			size, _ := strconv.ParseInt(response[i].Size, 10, 64)
			response[i].Size = humanize.Bytes(uint64(size))
		}

		pterm.PrintTable(response)
	}

}
