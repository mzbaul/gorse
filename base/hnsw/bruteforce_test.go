// Copyright 2023 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hnsw

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBruteforce(t *testing.T) {
	bf := NewBruteforce(Euclidean)
	bf.Add(context.Background(), NewDenseVector([]float32{0, 0, 0, 0, 0, 1}))
	bf.Add(context.Background(), NewDenseVector([]float32{0, 0, 0, 0, 1, 1}))
	bf.Add(context.Background(), NewDenseVector([]float32{0, 0, 0, 1, 1, 1}))
	bf.Add(context.Background(), NewDenseVector([]float32{0, 0, 1, 1, 1, 1}))
	bf.Add(context.Background(), NewDenseVector([]float32{0, 1, 1, 1, 1, 1}))
	bf.Add(context.Background(), NewDenseVector([]float32{1, 1, 1, 1, 1, 1}))
	results := bf.Search(NewDenseVector([]float32{0, 0, 0, 0, 0, 0}), 3)
	assert.Equal(t, []Result{
		{Index: 0, Distance: 1},
		{Index: 1, Distance: 2},
		{Index: 2, Distance: 3},
	}, results)

	bf.distFn = Dot
	results = bf.Search(NewDenseVector([]float32{1, 1, 1, 1, 1, 1}), 3)
	assert.Equal(t, []Result{
		{Index: 5, Distance: -6},
		{Index: 4, Distance: -5},
		{Index: 3, Distance: -4},
	}, results)
}

func TestBruteforce_HasNil(t *testing.T) {
	vectors := []Vector{nil, nil, NewDenseVector([]float32{0, 0, 0, 0, 0, 1}), nil}
	bf := NewBruteforce(Euclidean)
	bf.Add(context.Background(), vectors...)
	results := bf.Search(NewDenseVector([]float32{0, 0, 0, 0, 0, 1}), 3)
	assert.Equal(t, []Result{{Index: 2, Distance: 0}}, results)
}

func TestBruteforce_HasInf(t *testing.T) {
	bf := NewBruteforce(Euclidean)
	bf.Add(context.Background(), NewSparseVector([]int32{1, 3, 5}, []float32{1, 2, 3}))
	bf.Add(context.Background(), NewSparseVector([]int32{2, 4, 6}, []float32{1, 2, 3}))
	results := bf.Search(NewSparseVector([]int32{1, 3, 5}, []float32{1, 2, 3}), 3)
	assert.Equal(t, []Result{{Index: 0, Distance: 0}}, results)
}
