package metrics

import (
	"path"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrafficMetricsList provides a list of resources associated with a specific reference
type TrafficMetricsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Resource *v1.ObjectReference `json:"resource"`

	Items []*TrafficMetrics `json:"items"`
}

func NewTrafficMetricsList(obj *v1.ObjectReference) *TrafficMetricsList {
	return &TrafficMetricsList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "TrafficMetricsList",
			APIVersion: APIVersion,
		},
		ListMeta: metav1.ListMeta{
			SelfLink: path.Join(baseURL(), getKindName(obj.Kind)),
		},
		Resource: obj,
	}
}

func (lst *TrafficMetricsList) match(left, right *v1.ObjectReference) bool {
	return left.Kind == right.Kind &&
		left.Namespace == right.Namespace &&
		left.Name == right.Name
}

// Get will get the item that is associated with the object
// reference or create a default if it doesn't already exist.
func (lst *TrafficMetricsList) Get(
	obj, edge *v1.ObjectReference) *TrafficMetrics {

	for _, item := range lst.Items {
		if lst.match(obj, item.Resource) {
			if edge == nil || lst.match(edge, item.Edge.Resource) {
				return item
			}
		}
	}

	t := NewTrafficMetrics(obj, edge)
	lst.Items = append(lst.Items, t)

	return t
}
