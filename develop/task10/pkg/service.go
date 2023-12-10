package pkg

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"task10/pkg/config"
)

func Echo(r io.Reader, w io.Writer) error {
	_, err := io.Copy(w, r)
	return err
}

func Telnet(cfg *config.Config) error {
	// notify sigint
	fmt.Println(cfg)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Trying %s...\n", addr)
	conn, err := net.DialTimeout("tcp", addr, cfg.Timeout)
	if err != nil {
		return err
	}
	log.Printf("Connected to %s\n", addr)

	senderDone := make(chan error)
	receiverDone := make(chan error)

	// Отправка
	go func() {
		senderDone <- Echo(os.Stdin, conn)
	}()

	// Прием
	go func() {
		receiverDone <- Echo(conn, os.Stdout)
	}()

	select {
	case err := <-receiverDone:
		if err != nil {
			fmt.Println("receiver err:", err)
		}
		close(receiverDone)
	case err := <-senderDone:
		if err != nil {
			fmt.Println("sender err:", err)
		}
		close(senderDone)
	case s := <-sig:
		log.Println("receive os signal:", s.String())
	}

	conn.Close()

	if _, ok := <-receiverDone; ok {
		close(receiverDone)
	}
	if _, ok := <-senderDone; ok {
		close(senderDone)
	}

	return nil
}
