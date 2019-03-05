package jsonpb

import (
	spb "github.com/explodes/serving/proto"
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/utilz"
	"github.com/gorilla/handlers"
	"github.com/zang-cloud/grpc-json"
	"net/http"
)

func ServeJson(httpAddr *spb.Address, servers ...interface{}) error {
	var muxs []*http.ServeMux
	for _, server := range servers {
		mux := grpcj.Mux(server, grpcj.Middleware(handlers.CORS()))
		if _, isStatusz := server.(statusz.StatuszServiceServer); isStatusz {
			statusz.RegisterStatuszWebpage(httpAddr.Address(), mux)
		}
		muxs = append(muxs, mux)
	}
	mux := CombineMux(muxs...)
	server := &http.Server{Addr: httpAddr.Address(), Handler: mux}
	utilz.RegisterGracefulShutdownHttpServer("json-server", server)
	return server.ListenAndServe()
}

func CombineMux(muxs ...*http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, mux := range muxs {
			if handler, pattern := mux.Handler(r); pattern != "" {
				handler.ServeHTTP(w, r)
				return
			}
		}
		http.NotFoundHandler().ServeHTTP(w, r)
	})
}
