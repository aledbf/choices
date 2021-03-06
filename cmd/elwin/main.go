// Copyright 2016 Andrew O'Neill

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/foolusion/choices"
	"github.com/foolusion/choices/elwin"
	"github.com/foolusion/choices/storage/mongo"
	"google.golang.org/grpc"
)

var config = struct {
	ec              *choices.ElwinConfig
	grpcAddr        string
	jsonAddr        string
	mongoAddr       string
	mongoDB         string
	mongoCollection string
}{
	jsonAddr:        ":8081",
	grpcAddr:        ":8080",
	mongoAddr:       "elwin-storage",
	mongoDB:         "elwin",
	mongoCollection: "test",
}

func init() {
	http.HandleFunc("/", rootHandler)
}

const (
	envJSONAddr  = "JSON_ADDRESS"
	envGRPCAddr  = "GRPC_ADDRESS"
	envMongoAddr = "MONGO_ADDRESS"
	envMongoDB   = "MONGO_DATABASE"
	envMongoColl = "MONGO_COLLECTION"
)

func main() {
	log.Println("Starting elwin...")

	// TODO: read environment variables for config
	if os.Getenv(envJSONAddr) != "" {
		config.jsonAddr = os.Getenv(envJSONAddr)
	}

	if os.Getenv(envGRPCAddr) != "" {
		config.grpcAddr = os.Getenv(envGRPCAddr)
	}

	if os.Getenv(envMongoAddr) != "" {
		config.mongoAddr = os.Getenv(envMongoAddr)
	}

	if os.Getenv(envMongoDB) != "" {
		config.mongoDB = os.Getenv(envMongoDB)
	}

	if os.Getenv(envMongoColl) != "" {
		config.mongoCollection = os.Getenv(envMongoColl)
	}

	// create elwin config
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ec, err := choices.NewElwin(
		ctx,
		mongo.WithMongoStorage(config.mongoAddr, config.mongoDB, config.mongoCollection),
		choices.UpdateInterval(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	config.ec = ec

	// TODO: remove when deployed
	m := ec.Storage.(*mongo.Mongo)
	m.LoadExampleData()

	go func() {
		config.ec.ErrChan <- http.ListenAndServe(config.jsonAddr, nil)
	}()

	go func() {
		lis, err := net.Listen("tcp", config.grpcAddr)
		if err != nil {
			config.ec.ErrChan <- fmt.Errorf("main: failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		elwin.RegisterElwinServer(grpcServer, &elwinServer{})
		config.ec.ErrChan <- grpcServer.Serve(lis)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-config.ec.ErrChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Printf("Captured %v. Exitting...", s)
			// send StatusServiceUnavailable to new requestors
			// block server from accepting new requests
			os.Exit(0)
		}
	}
}

type elwinServer struct{}

func (e *elwinServer) GetNamespaces(ctx context.Context, id *elwin.Identifier) (*elwin.Experiments, error) {
	if id == nil {
		return nil, fmt.Errorf("GetNamespaces: no Identifier recieved")
	}

	resp, err := config.ec.Namespaces(id.TeamID, id.UserID)
	if err != nil {
		return nil, fmt.Errorf("error resolving namespaces for %s, %s: %v", id.TeamID, id.UserID, err)
	}

	exp := &elwin.Experiments{
		Experiments: make(map[string]*elwin.Experiment, len(resp)),
	}

	for _, v := range resp {
		exp.Experiments[v.Name] = &elwin.Experiment{
			Params: make([]*elwin.Param, len(v.Params)),
		}

		for i, p := range v.Params {
			exp.Experiments[v.Name].Params[i] = &elwin.Param{
				Name:  p.Name,
				Value: p.Value,
			}
		}
	}
	return exp, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := config.ec.Namespaces(r.Form.Get("teamid"), r.Form.Get("userid"))
	if err != nil {
		config.ec.ErrChan <- fmt.Errorf("rootHandler: couldn't get Namespaces: %v", err)
		return
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		config.ec.ErrChan <- fmt.Errorf("rootHandler: couldn't encode json: %v", err)
	}
}
