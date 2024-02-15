package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Reverse(next http.Handler) http.Handler {
	fmt.Println("work")
	userServiceURL, err := url.Parse("http://mcs1:8081")
	if err != nil {
		log.Println("err", err)
	}
	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)
	orderServiceURL, err := url.Parse("http://mcs2:8082")
	if err != nil {
		log.Println("err", err)
	}
	orderProxy := httputil.NewSingleHostReverseProxy(orderServiceURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/users") {
			userProxy.ServeHTTP(w, r)
		}
		if strings.HasPrefix(r.URL.Path, "/orders") {
			orderProxy.ServeHTTP(w, r)
		}
	})
}
