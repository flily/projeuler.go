package framework

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/flily/projeuler.go/framework/connection"
	"github.com/flily/projeuler.go/framework/message"
)

type Worker struct {
	runner *Runner
	conn   *connection.WorkerConn
	logger *log.Logger
}

func NewWorker(host string, port int) (*Worker, error) {
	conn, err := connection.NewWorkerConn(host, port)
	if err != nil {
		return nil, err
	}

	worker := &Worker{
		conn:   conn,
		runner: NewRunner(),
		logger: log.New(os.Stderr, "", log.Llongfile|log.Lmicroseconds),
	}

	return worker, nil
}

func (w *Worker) Close() {
	w.conn.Close()
}

func (w *Worker) Import(problems []Problem) {
	w.runner.Import(problems)
}

func (w *Worker) Serve() {
	w.logger.Printf("waiting for connection...")
	_ = w.conn.RunLoop()
}

func (w *Worker) Process() {
	for request := range w.conn.RecvRun() {
		w.DoRun(request)
	}
}

func (w *Worker) DoRun(request *message.MessageRun) {
	ctx, cancel := NewTimeoutContext(request.MethodTimeout)
	defer cancel()

	w.logger.Printf("run problem %d '%s', timeout=%s", request.Problem, request.Method, request.MethodTimeout)
	info := NewProblemRunInfo(request.Problem, request.Method)
	result, err := w.runner.RunProblemWithTimeout(ctx, info)
	if err == nil {
		w.conn.SendResult(result.ToMessage())
		return
	}

	// Process error, send error message to client before panic.
	w.logger.Printf("run problem %d '%s' failed: %s", request.Problem, request.Method, err)
	result.Message = err.Error()
	if errors.Is(err, context.DeadlineExceeded) {
		w.logger.Printf("timeout: %s", err)
	}

	log.Printf("send error message to client: %+v", result)
	w.conn.SendResult(result.ToMessage())
	time.Sleep(100 * time.Millisecond)
	panic(err)
}
