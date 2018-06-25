package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"context"
	"google.golang.org/grpc"

	"github.com/pkg/errors"

	// This Service
	"bookinfo/bookcomments-service/svc/client/cli/handlers"
	grpcclient "bookinfo/bookcomments-service/svc/client/grpc"
	httpclient "bookinfo/bookcomments-service/svc/client/http"
	pb "bookinfo/pb/comments"
)

var (
	_ = strconv.ParseInt
	_ = strings.Split
	_ = json.Compact
	_ = errors.Wrapf
	_ = pb.RegisterBookCommentsServer
)

func main() {
	os.Exit(submain())
}

type headerSeries []string

func (h *headerSeries) Set(val string) error {
	const requiredParts int = 2
	parts := strings.SplitN(val, ":", requiredParts)
	if len(parts) != requiredParts {
		return fmt.Errorf("value %q cannot be split in two; must contain at least one ':' character", val)
	}
	parts[1] = strings.TrimSpace(parts[1])
	*h = append(*h, parts...)
	return nil
}

func (h *headerSeries) String() string {
	return fmt.Sprintf("%v", []string(*h))
}

// submain exists to act as the functional main, but will return exit codes to
// the actual main instead of calling os.Exit directly. This is done to allow
// the defered functions to be called, since if os.Exit where called directly
// from this function, none of the defered functions set up by this function
// would be called.
func submain() int {
	var headers headerSeries
	flag.Var(&headers, "header", "Header(s) to be sent in the transport (follows cURL style)")
	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", ":5012", "gRPC (HTTP) address of addsvc")
	)

	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an service. This presumption is reflected in
	// the cli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings.

	fsGet := flag.NewFlagSet("get", flag.ExitOnError)

	fsPost := flag.NewFlagSet("post", flag.ExitOnError)

	var (
		flagIdPost      = fsPost.Int64("id", 0, "")
		flagContentPost = fsPost.String("content", "", "")
		flagIdGet       = fsGet.Int64("id", 0, "")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Subcommands:\n")
		fmt.Fprintf(os.Stderr, "  %s\n", "get")
		fmt.Fprintf(os.Stderr, "  %s\n", "post")
	}
	if len(os.Args) < 2 {
		flag.Usage()
		return 1
	}

	flag.Parse()

	var (
		service pb.BookCommentsServer
		err     error
	)

	if *httpAddr != "" {
		service, err = httpclient.New(*httpAddr, httpclient.CtxValuesToSend(headers...))
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
			return 1
		}
		defer conn.Close()
		service, err = grpcclient.New(conn, grpcclient.CtxValuesToSend(headers...))
	} else {
		fmt.Fprintf(os.Stderr, "error: no remote address specified\n")
		return 1
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("No 'method' subcommand provided; exiting.")
		flag.Usage()
		return 1
	}

	ctx := context.Background()
	for i := 0; i < len(headers); i += 2 {
		ctx = context.WithValue(ctx, headers[i], headers[i+1])
	}

	switch flag.Args()[0] {

	case "get":
		fsGet.Parse(flag.Args()[1:])

		IdGet := *flagIdGet

		request, err := handlers.Get(IdGet)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.Get: %v\n", err)
			return 1
		}

		v, err := service.Get(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.Get: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IdGet)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "post":
		fsPost.Parse(flag.Args()[1:])

		IdPost := *flagIdPost
		ContentPost := *flagContentPost

		request, err := handlers.Post(IdPost, ContentPost)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.Post: %v\n", err)
			return 1
		}

		v, err := service.Post(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.Post: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(IdPost, ContentPost)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	default:
		flag.Usage()
		return 1
	}

	return 0
}
