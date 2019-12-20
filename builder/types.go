// builder package help us to write aggregate,filter,update maps simpler.
package builder

import "go.mongodb.org/mongo-driver/bson"

// S is simple map that can be substitute  of `bson.M` to
// having simpler map structure on query,aggregate,...
type S []Operator

func (s *S) MarshalBSON() ([]byte, error) {
	return bson.Marshal(s.ToMap())
}

func (s *S) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	for _, o := range *s {
		m[o.GetKey()] = o.GetVal()
	}

	return m
}
