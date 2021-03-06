// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import io "io"
import mock "github.com/stretchr/testify/mock"
import storage "github.com/peak/s5cmd/storage"
import url "github.com/peak/s5cmd/storage/url"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Copy provides a mock function with given fields: ctx, src, dst, metadata
func (_m *Storage) Copy(ctx context.Context, src *url.URL, dst *url.URL, metadata map[string]string) error {
	ret := _m.Called(ctx, src, dst, metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL, *url.URL, map[string]string) error); ok {
		r0 = rf(ctx, src, dst, metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, src
func (_m *Storage) Delete(ctx context.Context, src *url.URL) error {
	ret := _m.Called(ctx, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL) error); ok {
		r0 = rf(ctx, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, src, dst, concurrency, partSize
func (_m *Storage) Get(ctx context.Context, src *url.URL, dst io.WriterAt, concurrency int, partSize int64) (int64, error) {
	ret := _m.Called(ctx, src, dst, concurrency, partSize)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL, io.WriterAt, int, int64) int64); ok {
		r0 = rf(ctx, src, dst, concurrency, partSize)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *url.URL, io.WriterAt, int, int64) error); ok {
		r1 = rf(ctx, src, dst, concurrency, partSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, src, recursive
func (_m *Storage) List(ctx context.Context, src *url.URL, recursive bool) <-chan *storage.Object {
	ret := _m.Called(ctx, src, recursive)

	var r0 <-chan *storage.Object
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL, bool) <-chan *storage.Object); ok {
		r0 = rf(ctx, src, recursive)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *storage.Object)
		}
	}

	return r0
}

// ListBuckets provides a mock function with given fields: ctx, prefix
func (_m *Storage) ListBuckets(ctx context.Context, prefix string) ([]storage.Bucket, error) {
	ret := _m.Called(ctx, prefix)

	var r0 []storage.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string) []storage.Bucket); ok {
		r0 = rf(ctx, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeBucket provides a mock function with given fields: ctx, bucket
func (_m *Storage) MakeBucket(ctx context.Context, bucket string) error {
	ret := _m.Called(ctx, bucket)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, bucket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MultiDelete provides a mock function with given fields: ctx, urls
func (_m *Storage) MultiDelete(ctx context.Context, urls <-chan *url.URL) <-chan *storage.Object {
	ret := _m.Called(ctx, urls)

	var r0 <-chan *storage.Object
	if rf, ok := ret.Get(0).(func(context.Context, <-chan *url.URL) <-chan *storage.Object); ok {
		r0 = rf(ctx, urls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *storage.Object)
		}
	}

	return r0
}

// Put provides a mock function with given fields: ctx, src, dst, metadata, concurrency, partSize
func (_m *Storage) Put(ctx context.Context, src io.Reader, dst *url.URL, metadata map[string]string, concurrency int, partSize int64) error {
	ret := _m.Called(ctx, src, dst, metadata, concurrency, partSize)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, *url.URL, map[string]string, int, int64) error); ok {
		r0 = rf(ctx, src, dst, metadata, concurrency, partSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stat provides a mock function with given fields: ctx, src
func (_m *Storage) Stat(ctx context.Context, src *url.URL) (*storage.Object, error) {
	ret := _m.Called(ctx, src)

	var r0 *storage.Object
	if rf, ok := ret.Get(0).(func(context.Context, *url.URL) *storage.Object); ok {
		r0 = rf(ctx, src)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *url.URL) error); ok {
		r1 = rf(ctx, src)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
