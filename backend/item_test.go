/*
 *    Copyright (C) 2015 Stefan Luecke
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors: Stefan Luecke <glaxx@glaxx.net>
 */

package backend

import (
	dmp "github.com/sergi/go-diff/diffmatchpatch"
	"math/rand"
	"testing"
)

func Test_uint64Contains(t *testing.T) {
	A := []uint64{1, 2}
	B := []uint64{1, 3}
	if !uint64Contains(A, B[0]) || !uint64Contains(B, A[0]) {
		t.Error("failed")
	}
	if uint64Contains(A, B[1]) || uint64Contains(B, A[1]) {
		t.Error("failed")
	}
}

func Test_uint64Diff(t *testing.T) {
	A, B := []uint64{1, 2}, []uint64{1, 3}
	C := uint64Diff(A, B)
	_, ok := C["1"]
	if ok {
		t.Error("failed")
	}
	if C["2"] != dmp.DiffDelete {
		t.Error("failed")
	}
	if C["3"] != dmp.DiffInsert {
		t.Error("failed")
	}
}

func Benchmark_uint64Contains(b *testing.B) {
	A := make([]uint64, 10000)
	for i := 0; i != len(A); i++ {
		A[i] = uint64(rand.Int63())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uint64Contains(A, uint64(rand.Int63()))
	}
}