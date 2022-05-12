package ziface

type IServer interface {
	Start()

	Stop()

	Serve()

	AddRoute(router IRouter)
}
