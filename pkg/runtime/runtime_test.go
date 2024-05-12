package runtime_test

import (
	"testing"

	minik8s_apiobject "github.com/MiniK8s-SE3356/minik8s/pkg/apiObject"
	minik8s_pod "github.com/MiniK8s-SE3356/minik8s/pkg/apiObject/pod"
	minik8s_runtime "github.com/MiniK8s-SE3356/minik8s/pkg/runtime"
	minik8s_container "github.com/MiniK8s-SE3356/minik8s/pkg/types/container"
	"github.com/google/uuid"
)

func TestMain(m *testing.M) {
	test_pod := &minik8s_pod.Pod{
		Basic: minik8s_apiobject.Basic{
			APIVersion: "v1",
			Kind:       "Pod",
			Metadata: minik8s_apiobject.Metadata{
				UUID:      uuid.New().String(),
				Name:      "test-pod",
				Namespace: "default",
				Labels: map[string]string{
					"app": "test",
				},
			},
		},
		Spec: minik8s_pod.PodSpec{
			NodeName: "node1",
			Containers: []minik8s_container.Container{
				{
					Name:  "nginx-container",
					Image: "nginx:latest",
					Ports: []minik8s_container.ContainerPort{
						{
							HostPort:      80,
							ContainerPort: 80,
						},
					},
				},
				{
					Name:  "redis-container",
					Image: "redis:latest",
					Ports: []minik8s_container.ContainerPort{
						{
							HostPort:      6379,
							ContainerPort: 6379,
						},
					},
				},
			},
		},
		Status: minik8s_pod.PodStatus{
			Phase: minik8s_pod.PodPending,
		},
	}

	runtimeManager := &minik8s_runtime.RuntimeManager{}
	_, err := runtimeManager.CreatePod(test_pod)
	if err != nil {
		println("Error creating pod")
		panic(err)
	}
}