package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var namespace = "tcpeek"

var socket = "/var/run/tcpeek.sock"
var timeout = time.Second * 3

type exporter struct {
	RX      TcpeekPrometheusMetric
	TX      TcpeekPrometheusMetric
	HttpOut TcpeekPrometheusMetric
}

// TcpeekStats is statistics list for tcpeek
type TcpeekStats []TcpeekStat

// TcpeekStat is statistics for tcpeek
type TcpeekStat struct {
	RX      TcpeekMetric `json:"RX"`
	TX      TcpeekMetric `json:"TX"`
	HttpOut TcpeekMetric `json:"http-out"`
}

func newExporter() *exporter {
	return &exporter{
		RX:      DescribeTcpeekPrometheusMetric("RX"),
		TX:      DescribeTcpeekPrometheusMetric("TX"),
		HttpOut: DescribeTcpeekPrometheusMetric("http_out"),
	}
}

func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	e.RX.Success.Total.Describe(ch)
	e.RX.Success.DupSyn.Describe(ch)
	e.RX.Success.DupSynAck.Describe(ch)
	e.RX.Failure.Total.Describe(ch)
	e.RX.Failure.Timeout.Describe(ch)
	e.RX.Failure.Reject.Describe(ch)
	e.RX.Failure.Unreach.Describe(ch)

	e.TX.Success.Total.Describe(ch)
	e.TX.Success.DupSyn.Describe(ch)
	e.TX.Success.DupSynAck.Describe(ch)
	e.TX.Failure.Total.Describe(ch)
	e.TX.Failure.Timeout.Describe(ch)
	e.TX.Failure.Reject.Describe(ch)
	e.TX.Failure.Unreach.Describe(ch)

	e.HttpOut.Success.Total.Describe(ch)
	e.HttpOut.Success.DupSyn.Describe(ch)
	e.HttpOut.Success.DupSynAck.Describe(ch)
	e.HttpOut.Failure.Total.Describe(ch)
	e.HttpOut.Failure.Timeout.Describe(ch)
	e.HttpOut.Failure.Reject.Describe(ch)
	e.HttpOut.Failure.Unreach.Describe(ch)
}

func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	conn, err := net.DialTimeout("unix", socket, timeout)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	defer conn.Close()

	var status TcpeekStats
	fmt.Fprintf(conn, "GET\r\n")
	// data, err := bufio.NewReader(conn).ReadString('\n')
	dec := json.NewDecoder(conn)
	dec.Decode(&status)

	e.RX.Success.Total.Set(float64(status[0].RX.Success.Total))
	e.RX.Success.DupSyn.Set(float64(status[0].RX.Success.DupSyn))
	e.RX.Success.DupSynAck.Set(float64(status[0].RX.Success.DupSynAck))
	e.RX.Failure.Total.Set(float64(status[0].RX.Failure.Total))
	e.RX.Failure.Timeout.Set(float64(status[0].RX.Failure.Timeout))
	e.RX.Failure.Reject.Set(float64(status[0].RX.Failure.Reject))
	e.RX.Failure.Unreach.Set(float64(status[0].RX.Failure.Unreach))

	e.RX.Success.Total.Collect(ch)
	e.RX.Success.DupSyn.Collect(ch)
	e.RX.Success.DupSynAck.Collect(ch)
	e.RX.Failure.Total.Collect(ch)
	e.RX.Failure.Timeout.Collect(ch)
	e.RX.Failure.Reject.Collect(ch)
	e.RX.Failure.Unreach.Collect(ch)

	e.TX.Success.Total.Set(float64(status[1].TX.Success.Total))
	e.TX.Success.DupSyn.Set(float64(status[1].TX.Success.DupSyn))
	e.TX.Success.DupSynAck.Set(float64(status[1].TX.Success.DupSynAck))
	e.TX.Failure.Total.Set(float64(status[1].TX.Failure.Total))
	e.TX.Failure.Timeout.Set(float64(status[1].TX.Failure.Timeout))
	e.TX.Failure.Reject.Set(float64(status[1].TX.Failure.Reject))
	e.TX.Failure.Unreach.Set(float64(status[1].TX.Failure.Unreach))

	e.TX.Success.Total.Collect(ch)
	e.TX.Success.DupSyn.Collect(ch)
	e.TX.Success.DupSynAck.Collect(ch)
	e.TX.Failure.Total.Collect(ch)
	e.TX.Failure.Timeout.Collect(ch)
	e.TX.Failure.Reject.Collect(ch)
	e.TX.Failure.Unreach.Collect(ch)

	e.HttpOut.Success.Total.Set(float64(status[2].HttpOut.Success.Total))
	e.HttpOut.Success.DupSyn.Set(float64(status[2].HttpOut.Success.DupSyn))
	e.HttpOut.Success.DupSynAck.Set(float64(status[2].HttpOut.Success.DupSynAck))
	e.HttpOut.Failure.Total.Set(float64(status[2].HttpOut.Failure.Total))
	e.HttpOut.Failure.Timeout.Set(float64(status[2].HttpOut.Failure.Timeout))
	e.HttpOut.Failure.Reject.Set(float64(status[2].HttpOut.Failure.Reject))
	e.HttpOut.Failure.Unreach.Set(float64(status[2].HttpOut.Failure.Unreach))

	e.HttpOut.Success.Total.Collect(ch)
	e.HttpOut.Success.DupSyn.Collect(ch)
	e.HttpOut.Success.DupSynAck.Collect(ch)
	e.HttpOut.Failure.Total.Collect(ch)
	e.HttpOut.Failure.Timeout.Collect(ch)
	e.HttpOut.Failure.Reject.Collect(ch)
	e.HttpOut.Failure.Unreach.Collect(ch)
}

func main() {
	var help = false
	var verbose = false
	var listen = "127.0.0.1:9381"
	var prefix = ""

	// parse args
	flags := flag.NewFlagSet(App.Name, flag.ContinueOnError)

	flags.BoolVar(&verbose, "verbose", verbose, "verbose")
	flags.BoolVar(&help, "h", help, "help")
	flags.BoolVar(&help, "help", help, "help")
	flags.BoolVar(&help, "version", help, "version")
	flags.StringVar(&socket, "tcpeek-socket", socket, "tcpeek-socket")
	flags.StringVar(&prefix, "prefix", prefix, "prefix")
	flags.StringVar(&listen, "listen", listen, "listen")

	flags.Usage = func() { usage() }
	if err := flags.Parse(os.Args[1:]); err != nil {
		err = errors.Wrap(err, fmt.Sprintf("%v-%v failed: ", App.Name, App.Version))
		fmt.Printf("%v-%v failed: %v\n", App.Name, App.Version, err)
		os.Exit(1)
	}

	if help {
		usage()
		os.Exit(0)
	}

	if prefix != "" {
		namespace = fmt.Sprintf("%s_%s", prefix, namespace)
	}

	exporter := newExporter()

	prometheus.MustRegister(exporter)

	http.Handle("/metrics", prometheus.Handler())

	addr := fmt.Sprintf("%s", strings.Split(listen, ":")[0])
	port := fmt.Sprintf(":%s", strings.Split(listen, ":")[1])
	log.Print("Listening ", addr, port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func usage() {
	helpText := `
usage:
   {{.Name}} [command options]

version:
   {{.Version}}

author:
   kizkoh<GitHub: https://github.com/kizkoh>

options:
   --tcpeek-socket [PATH]                       Unix domain socket for tcpeek (default: /var/run/tcpeek.sock)
   --listen [ADDR:PORT]                         Listen addr and port (default: 127.0.0.1:9381)
   --prefix [ADDR:PORT]                         Prometheus metric name prefix (default: None)
   --verbose                                    Print verbose messages
   --help, -h                                   Show help
   --version                                    Print the version
`
	t := template.New("usage")
	t, _ = t.Parse(strings.TrimSpace(helpText))
	t.Execute(os.Stdout, App)
	fmt.Println()
}
