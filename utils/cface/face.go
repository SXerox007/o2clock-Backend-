package cface

import (
	"math"
	face "o2clock/algorithm-face"
)

func CompareFaces(samples []face.Descriptor, comp face.Descriptor, tolerance float32) int {
	res := face_distance(samples, comp)
	r := -1
	v := float32(1)

	for i, s := range res {
		t := euclidean_norm(s)
		if t < tolerance && t < v {
			v = t
			r = i
		}
	}

	return r
}

func face_distance(samples []face.Descriptor, comp face.Descriptor) []face.Descriptor {
	res := make([]face.Descriptor, len(samples))

	for i, s := range samples {
		for j, _ := range s {
			res[i][j] = samples[i][j] - comp[j]
		}
	}

	return res
}

func euclidean_norm(f face.Descriptor) float32 {
	var s float32
	for _, v := range f {
		s = s + v*v
	}

	return float32(math.Sqrt(float64(s)))
}
