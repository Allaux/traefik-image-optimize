package cache_test

import (
	"reflect"
	"testing"

	"github.com/Allaux/traefik-image-optimize/cache"
	"github.com/Allaux/traefik-image-optimize/config"
)

func TestNew(t *testing.T) {
	type args struct {
		conf config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    cache.Cache
		wantErr bool
	}{
		{
			name:    "should be able to memory cache",
			args:    args{config.Config{Processor: "none", Cache: "memory"}},
			want:    &cache.MemoryCache{},
			wantErr: false,
		},
		{
			name:    "should be able to memory cache",
			args:    args{config.Config{Processor: "none", Cache: "none"}},
			want:    &cache.NoneCache{},
			wantErr: false,
		},
		{
			name:    "should not be able to init cache without valid driver",
			args:    args{config.Config{Processor: "imaginary", Imaginary: config.ImaginaryProcessorConfig{Url: "http://localhost"}, Cache: "unsupported"}},
			want:    nil,
			wantErr: true,
		},
		// {
		// 	name:    "should not be able to init imaginary without valid url",
		// 	args:    args{config.Config{Processor: "imaginary", Imaginary: config.ImaginaryConfig{Url: "localhost"}, Cache: "memory"}},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		// {
		// 	name:    "should not be able to init imaginary without valid url 2 ",
		// 	args:    args{config.Config{Processor: "imaginary", Imaginary: config.ImaginaryConfig{Url: "htt://localhost"}}},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		// {
		// 	name:    "should not be able to init imaginary without url",
		// 	args:    args{config.Config{Processor: "imaginary"}},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		// {
		// 	name:    "should be able to return local optimizer",
		// 	args:    args{config.Config{Processor: "local"}},
		// 	want:    &processor.LocalProcessor{},
		// 	wantErr: false,
		// },
		// {
		// 	name:    "should return error with unsupported processor",
		// 	args:    args{config.Config{Processor: "unsupported"}},
		// 	want:    nil,
		// 	wantErr: true,
		// },
		// {
		// 	name:    "should return error with empty processor",
		// 	args:    args{config.Config{Processor: "unsupported"}},
		// 	want:    nil,
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.New(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Fatalf("New() error = %v, wantErr %v", err, tt.wantErr)
			}

			typeGot := reflect.TypeOf(got)
			typeWanted := reflect.TypeOf(tt.want)

			if typeGot != typeWanted {
				t.Errorf("New() = %v, want %v", typeGot, typeWanted)
			}
		})
	}
}
