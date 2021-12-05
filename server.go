package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)
type Points struct {
	HP int
	MP int
}
type Params struct {
	phdef  int
	mgdef  int
	mgforc int
	phforc int
}
type Class struct {
	ClassName  string
	DamageType string
	Pointes    Points
	paramer    Params
}
type Classes struct {
	Mage   Class
	Knight Class
}
Mage.ClassName="Mage";
/*type ChooseclassServer interface {
	Do(context.Context, *Requestcls) (*Responsecls, error)
	mustEmbedUnimplementedChooseclassServer()
}

type UnimplementedChooseclassServer struct {
}

func (UnimplementedChooseclassServer) Do(context.Context, *Requestcls) (*Responsecls, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func _Chooseclass_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Requestcls)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChooseclassServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swords.Chooseclass/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChooseclassServer).Do(ctx, req.(*Requestcls))
	}
	return interceptor(ctx, in, info, handler)
}
var Chooseclass_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swords.Chooseclass",
	HandlerType: (*ChooseclassServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Chooseclass_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon.proto",
}
func main() {

}*/
