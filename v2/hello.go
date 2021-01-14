package hello

import "rsc.io/quote/v3"

func Hello() string {
    return quote.HelloV3() + " v2"
}

func Proverb() string {
    return quote.Concurrency() + " v2"
}