package k8_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pradeepmvn/xds-controller/pkg/config"
	"github.com/pradeepmvn/xds-controller/pkg/log"
	"github.com/pradeepmvn/xds-controller/pkg/resolver"
	"github.com/pradeepmvn/xds-controller/pkg/resolver/k8"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetEndpoints(t *testing.T) {
	log.NewLogger(true)
	k8s := fake.NewSimpleClientset(
		&v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test",
				Namespace:   "default",
				Annotations: map[string]string{},
			},
			Subsets: []v1.EndpointSubset{
				{
					Addresses: []v1.EndpointAddress{
						{IP: "1.1.1.1"},
					},
				},
			}},
	)

	cc := &config.Cluster{
		Name:         "FakeK8",
		ResolverType: "k8",
		Address:      "test",
	}
	r := k8.New(cc, k8s)
	ep := r.GetEndPoints()
	fmt.Println(ep)
}

func TestWatch(t *testing.T) {
	log.NewLogger(true)
	cc := &config.Cluster{
		Name:         "FakeK8",
		ResolverType: "k8",
		Address:      "test",
	}
	k8s := fake.NewSimpleClientset(
		&v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test",
				Namespace:   "default",
				Annotations: map[string]string{},
			},
			Subsets: []v1.EndpointSubset{
				{
					Addresses: []v1.EndpointAddress{
						{IP: "1.1.1.1"},
					},
				},
			}},
	)
	r := k8.New(cc, k8s)
	ep := r.GetEndPoints()
	fmt.Println("Initial List: ", ep)
	go func(r resolver.Resolver) {
		//kill after 10 sec
		time.Sleep(10 * time.Second)
		r.Close()
		fmt.Println("Stopping")
	}(r)
	// receive refreshes
	go func(r resolver.Resolver) {
		for range r.Refresh() {
			fmt.Println("New List", r.Latest())
		}
	}(r)
	go func(k8s *fake.Clientset) {
		// change endpoint after 1 sec
		time.Sleep(3 * time.Second)
		fmt.Println("Changing Endpoint")
		k8s.CoreV1().Endpoints("default").Create(context.Background(), &v1.Endpoints{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test",
				Namespace:   "default",
				Annotations: map[string]string{},
			},
			Subsets: []v1.EndpointSubset{
				{
					Addresses: []v1.EndpointAddress{
						{IP: "1.1.1.1"},
						{IP: "1.1.1.2"},
					},
				},
			}}, metav1.CreateOptions{})
	}(k8s)
	fmt.Println("Starting Watch")
	r.Watch()
}
