package biblehub

type Scraper struct {
	Downloader Downloader
}

var DefaultScraper = Scraper{
	Downloader: DefaultDownloader,
}
