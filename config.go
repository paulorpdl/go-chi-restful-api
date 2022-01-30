package main

type Config struct {
	Port string `default:"5000"`
	Addr string `default:"0.0.0.0"`
	Path string `default:"/api/v1"`
}
