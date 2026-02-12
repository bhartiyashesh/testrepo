package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	started := time.Now().Format(time.RFC3339)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
<title>Build & Ship Test</title>
<style>
* { margin:0; padding:0; box-sizing:border-box; }
body { font-family:'JetBrains Mono',monospace; background:#030303; color:#e0e0e0; display:flex; align-items:center; justify-content:center; min-height:100vh; }
.box { text-align:center; padding:48px; }
h1 { font-size:28px; color:#f59e0b; margin-bottom:12px; }
p { color:#666; font-size:14px; margin-bottom:8px; }
.badge { display:inline-block; padding:6px 16px; background:#0a2a0a; color:#22c55e; border-radius:4px; font-size:13px; margin-top:16px; }
</style>
</head>
<body>
<div class="box">
<h1>BUILD & SHIP</h1>
<p>Deployed via P2P Swarm</p>
<p>Host: %s</p>
<p>Started: %s</p>
<div class="badge">LIVE</div>
</div>
</body>
</html>`, hostname(), started)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "ok")
	})

	fmt.Printf("Listening on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func hostname() string {
	h, _ := os.Hostname()
	return h
}
// version 2
// version 3 - Thu Feb 12 08:06:23 CST 2026
// v4 - 1770905837
