.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo-thrift && cwgo server -I ../../idl --type RPC --module biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl  ../../idl/frontend/order_page.proto --service frontend -module biz-demo/gomall/app/frontend -I ../../idl -I /home/qiang-ubun/pan1/download/ProtoBuf/protoc-29.0-rc-2-linux-x86_64/include

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module biz-demo/gomall/app/user --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module biz-demo/gomall/app/product --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module biz-demo/gomall/app/cart --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module biz-demo/gomall/app/payment --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module biz-demo/gomall/app/checkout --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module biz-demo/gomall/app/order --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module biz-demo/gomall/rpc_gen  -I ../idl  --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module biz-demo/gomall/app/email --pass "-use biz-demo/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/email.proto


.PHONY: open.consul
open-consul: ## open `consul ui` in the default browser
	@open "http://localhost:8500/ui/"

.PHONY: open.jaeger
open-jaeger: ## open `jaeger ui` in the default browser
	@open "http://localhost:16686/search"

.PHONY: open.prometheus
open-prometheus: ## open `prometheus ui` in the default browser
	@open "http://localhost:9090"

# docker run -v ./app/frontend/conf:/opt/gomall/frontend/conf --network gomall_default --env-file ./app/frontend/.env -p 8080:8080 frontend:v1.1.1
.PHONY:	build-frontend
build-frontend:
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .

# docker run -v ./app/product/conf:/opt/gomall/product/conf --network gomall_default --env-file ./app/product/.env product:v1.1.1
.PHONY:	build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .