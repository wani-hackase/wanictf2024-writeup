package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

func makeAdminCh() chan uuid.UUID {
	return make(chan uuid.UUID)
}

func startAdminWorker(ch chan uuid.UUID) {
	const numAdminWorkers = 6
	for i := 0; i < numAdminWorkers; i++ {
		go adminWorker(ch)
	}
}

func adminWorker(ch chan uuid.UUID) {
	for id := range ch {
		visitArticle(id)
	}
}

func visitArticle(id uuid.UUID) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "elec-admin-console", "http://localhost:"+os.Getenv("PORT")+"/article/"+id.String())
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}
	cmd.WaitDelay = 5 * time.Second

	out, err := cmd.CombinedOutput()
	log.Printf("Admin console completed for article %s: %+v %s", id.String(), err, out)
}
