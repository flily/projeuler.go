package framework

import (
	"log"
	"os"

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
		w.logger.Printf("run problem %d '%s'", request.Problem, request.Method)
		w.DoRun(request)
	}
}

func (w *Worker) DoRun(request *message.MessageRun) {
	ctx, cancel := NewTimeoutContext(request.ProblemTimeout)
	defer cancel()

	info := NewProblemRunInfo(request.Problem, request.Method)
	result, err := w.runner.RunProblemWithTimeout(ctx, info)
	if err != nil {
		resultMessage := message.NewResult()
		resultMessage.Message = err.Error()
		w.conn.SendResult(resultMessage)
		return
	}

	w.conn.SendResult(result.ToMessage())
}
