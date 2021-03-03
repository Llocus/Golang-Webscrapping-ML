package main

import (
	"fmt"
	"github.com/gocolly/colly"
) 

func main() {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Ctx.Get("\n url \n"))
	})

	c.OnHTML(".ui-search-layout.ui-search-layout--grid > li", func(e *colly.HTMLElement) {
		sitSend := e.ChildText(".ui-search-item__highlight-label__text")
		description := e.ChildText(".ui-search-item__title.ui-search-item__group__element")
		usad := e.ChildTexts(".ui-search-item__group__element.ui-search-item__details")
		var usadd string
		if usad != nil {
			usadd = "true"
		} else {
			usadd = "false"
		}
		image := e.ChildAttr("img", "data-src")
		price := e.ChildText(".ui-search-item__group.ui-search-item__group--price > .ui-search-price.ui-search-price--size-medium.ui-search-item__group__element > div > span:first-child > span.price-tag-fraction")
		resposta := fmt.Sprintf(" Image: %s,\n Price: R$ %s,\n Description: %s,\n Used?: %s,\n Send?: '%s' \n", image, price, description, usadd, sitSend)
		fmt.Println(resposta)
	})
	var product string
	fmt.Println("Search:")
	fmt.Scanf("%s \n", &product)
	
	c.Visit("https://lista.mercadolivre.com.br/" + product + "_OrderId_PRICE")
}
