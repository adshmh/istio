// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/copilot/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"istio.io/istio/pilot/pkg/config/cloudfoundry"
)

type CopilotClient struct {
	HealthStub        func(ctx context.Context, in *api.HealthRequest, opts ...grpc.CallOption) (*api.HealthResponse, error)
	healthMutex       sync.RWMutex
	healthArgsForCall []struct {
		ctx  context.Context
		in   *api.HealthRequest
		opts []grpc.CallOption
	}
	healthReturns struct {
		result1 *api.HealthResponse
		result2 error
	}
	healthReturnsOnCall map[int]struct {
		result1 *api.HealthResponse
		result2 error
	}
	RoutesStub        func(ctx context.Context, in *api.RoutesRequest, opts ...grpc.CallOption) (*api.RoutesResponse, error)
	routesMutex       sync.RWMutex
	routesArgsForCall []struct {
		ctx  context.Context
		in   *api.RoutesRequest
		opts []grpc.CallOption
	}
	routesReturns struct {
		result1 *api.RoutesResponse
		result2 error
	}
	routesReturnsOnCall map[int]struct {
		result1 *api.RoutesResponse
		result2 error
	}
	InternalRoutesStub        func(ctx context.Context, in *api.InternalRoutesRequest, opts ...grpc.CallOption) (*api.InternalRoutesResponse, error)
	internalRoutesMutex       sync.RWMutex
	internalRoutesArgsForCall []struct {
		ctx  context.Context
		in   *api.InternalRoutesRequest
		opts []grpc.CallOption
	}
	internalRoutesReturns struct {
		result1 *api.InternalRoutesResponse
		result2 error
	}
	internalRoutesReturnsOnCall map[int]struct {
		result1 *api.InternalRoutesResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CopilotClient) Health(ctx context.Context, in *api.HealthRequest, opts ...grpc.CallOption) (*api.HealthResponse, error) {
	fake.healthMutex.Lock()
	ret, specificReturn := fake.healthReturnsOnCall[len(fake.healthArgsForCall)]
	fake.healthArgsForCall = append(fake.healthArgsForCall, struct {
		ctx  context.Context
		in   *api.HealthRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("Health", []interface{}{ctx, in, opts})
	fake.healthMutex.Unlock()
	if fake.HealthStub != nil {
		return fake.HealthStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.healthReturns.result1, fake.healthReturns.result2
}

func (fake *CopilotClient) HealthCallCount() int {
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	return len(fake.healthArgsForCall)
}

func (fake *CopilotClient) HealthArgsForCall(i int) (context.Context, *api.HealthRequest, []grpc.CallOption) {
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	return fake.healthArgsForCall[i].ctx, fake.healthArgsForCall[i].in, fake.healthArgsForCall[i].opts
}

func (fake *CopilotClient) HealthReturns(result1 *api.HealthResponse, result2 error) {
	fake.HealthStub = nil
	fake.healthReturns = struct {
		result1 *api.HealthResponse
		result2 error
	}{result1, result2}
}

func (fake *CopilotClient) HealthReturnsOnCall(i int, result1 *api.HealthResponse, result2 error) {
	fake.HealthStub = nil
	if fake.healthReturnsOnCall == nil {
		fake.healthReturnsOnCall = make(map[int]struct {
			result1 *api.HealthResponse
			result2 error
		})
	}
	fake.healthReturnsOnCall[i] = struct {
		result1 *api.HealthResponse
		result2 error
	}{result1, result2}
}

func (fake *CopilotClient) Routes(ctx context.Context, in *api.RoutesRequest, opts ...grpc.CallOption) (*api.RoutesResponse, error) {
	fake.routesMutex.Lock()
	ret, specificReturn := fake.routesReturnsOnCall[len(fake.routesArgsForCall)]
	fake.routesArgsForCall = append(fake.routesArgsForCall, struct {
		ctx  context.Context
		in   *api.RoutesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("Routes", []interface{}{ctx, in, opts})
	fake.routesMutex.Unlock()
	if fake.RoutesStub != nil {
		return fake.RoutesStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.routesReturns.result1, fake.routesReturns.result2
}

func (fake *CopilotClient) RoutesCallCount() int {
	fake.routesMutex.RLock()
	defer fake.routesMutex.RUnlock()
	return len(fake.routesArgsForCall)
}

func (fake *CopilotClient) RoutesArgsForCall(i int) (context.Context, *api.RoutesRequest, []grpc.CallOption) {
	fake.routesMutex.RLock()
	defer fake.routesMutex.RUnlock()
	return fake.routesArgsForCall[i].ctx, fake.routesArgsForCall[i].in, fake.routesArgsForCall[i].opts
}

func (fake *CopilotClient) RoutesReturns(result1 *api.RoutesResponse, result2 error) {
	fake.RoutesStub = nil
	fake.routesReturns = struct {
		result1 *api.RoutesResponse
		result2 error
	}{result1, result2}
}

func (fake *CopilotClient) RoutesReturnsOnCall(i int, result1 *api.RoutesResponse, result2 error) {
	fake.RoutesStub = nil
	if fake.routesReturnsOnCall == nil {
		fake.routesReturnsOnCall = make(map[int]struct {
			result1 *api.RoutesResponse
			result2 error
		})
	}
	fake.routesReturnsOnCall[i] = struct {
		result1 *api.RoutesResponse
		result2 error
	}{result1, result2}
}

func (fake *CopilotClient) InternalRoutes(ctx context.Context, in *api.InternalRoutesRequest, opts ...grpc.CallOption) (*api.InternalRoutesResponse, error) {
	fake.internalRoutesMutex.Lock()
	ret, specificReturn := fake.internalRoutesReturnsOnCall[len(fake.internalRoutesArgsForCall)]
	fake.internalRoutesArgsForCall = append(fake.internalRoutesArgsForCall, struct {
		ctx  context.Context
		in   *api.InternalRoutesRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("InternalRoutes", []interface{}{ctx, in, opts})
	fake.internalRoutesMutex.Unlock()
	if fake.InternalRoutesStub != nil {
		return fake.InternalRoutesStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.internalRoutesReturns.result1, fake.internalRoutesReturns.result2
}

func (fake *CopilotClient) InternalRoutesCallCount() int {
	fake.internalRoutesMutex.RLock()
	defer fake.internalRoutesMutex.RUnlock()
	return len(fake.internalRoutesArgsForCall)
}

func (fake *CopilotClient) InternalRoutesArgsForCall(i int) (context.Context, *api.InternalRoutesRequest, []grpc.CallOption) {
	fake.internalRoutesMutex.RLock()
	defer fake.internalRoutesMutex.RUnlock()
	return fake.internalRoutesArgsForCall[i].ctx, fake.internalRoutesArgsForCall[i].in, fake.internalRoutesArgsForCall[i].opts
}

func (fake *CopilotClient) InternalRoutesReturns(result1 *api.InternalRoutesResponse, result2 error) {
	fake.InternalRoutesStub = nil
	fake.internalRoutesReturns = struct {
		result1 *api.InternalRoutesResponse
		result2 error
	}{result1, result2}
}

func (fake *CopilotClient) InternalRoutesReturnsOnCall(i int, result1 *api.InternalRoutesResponse, result2 error) {
	fake.InternalRoutesStub = nil
	if fake.internalRoutesReturnsOnCall == nil {
		fake.internalRoutesReturnsOnCall = make(map[int]struct {
			result1 *api.InternalRoutesResponse
			result2 error
		})
	}
	fake.internalRoutesReturnsOnCall[i] = struct {
		result1 *api.InternalRoutesResponse
		result2 error
	}{result1, result2}
}

func (fake *CopilotClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	fake.routesMutex.RLock()
	defer fake.routesMutex.RUnlock()
	fake.internalRoutesMutex.RLock()
	defer fake.internalRoutesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CopilotClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cloudfoundry.CopilotClient = new(CopilotClient)
