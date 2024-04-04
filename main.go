package main

import (
	"chat_san/db"
	"chat_san/handler/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/chat.db"
	)

	db, err := db.NewDB(defaultDBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := router.NewRouter(db)

	server := http.Server{
		Addr:    defaultDBPath,
		Handler: mux,
	}

	errCh := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errCh <- err

		}
	}()

	select {
	case <-ctx.Done():
		shutdownWithTimeout(&server)
	case err := <-errCh:
		log.Printf("Error: %v", err)
	}

}

func shutdownWithTimeout(srv *http.Server) {
	//Contextを作成し、30秒のTimeoutを設定
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Timeoutを設定したContextを渡すことで無期限に待機しないようにする
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
