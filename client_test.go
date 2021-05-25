package idcardauth

import (
	"github.com/haxqer/gofunc"
	"reflect"
	"testing"
	"time"
)

func TestClient_AuthCheck(t *testing.T) {
	type args struct {
		request *AuthCheckRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase-02-1",
			args: args{request: &AuthCheckRequest{
				AuthId: "200000000000000001",
				Name:   "某二一",
				IdCard: "110000190201010009",
			}}, want: nil, wantErr: false,
		},
		{
			name: "testCase-02-2",
			args: args{request: &AuthCheckRequest{
				AuthId: "200000000000000002",
				Name:   "某二二",
				IdCard: "110000190201020004",
			}}, want: nil, wantErr: false,
		},
		{
			name: "testCase-02-3",
			args: args{request: &AuthCheckRequest{
				AuthId: "200000000000000003",
				Name:   "某二三",
				IdCard: "11000019020103000X",
			}}, want: nil, wantErr: false,
		},
		{
			name: "testCase-03-1",
			args: args{request: &AuthCheckRequest{
				AuthId: "2111000000000000003",
				Name:   "某二三",
				IdCard: "11000019020103000X",
			}}, want: nil, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient("appid", "bizid", "key", "11111")
			got, err := c.AuthCheck(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthCheck() got = %s, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_AuthQuery(t *testing.T) {
	type args struct {
		request *AuthQueryRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "testCase-01-1", args: args{&AuthQueryRequest{AuthId: "100000000000000001"}}, want: nil, wantErr: false},
		{name: "testCase-01-2", args: args{&AuthQueryRequest{AuthId: "100000000000000002"}}, want: nil, wantErr: false},
		{name: "testCase-01-3", args: args{&AuthQueryRequest{AuthId: "100000000000000003"}}, want: nil, wantErr: false},
		{name: "testCase-01-4", args: args{&AuthQueryRequest{AuthId: "100000000000000004"}}, want: nil, wantErr: false},
		{name: "testCase-01-5", args: args{&AuthQueryRequest{AuthId: "100000000000000005"}}, want: nil, wantErr: false},
		{name: "testCase-01-6", args: args{&AuthQueryRequest{AuthId: "100000000000000006"}}, want: nil, wantErr: false},
		{name: "testCase-01-7", args: args{&AuthQueryRequest{AuthId: "100000000000000007"}}, want: nil, wantErr: false},
		{name: "testCase-01-8", args: args{&AuthQueryRequest{AuthId: "100000000000000008"}}, want: nil, wantErr: false},
		{name: "testCase-02-1", args: args{&AuthQueryRequest{AuthId: "200000000000000001"}}, want: nil, wantErr: false},
		{name: "testCase-02-2", args: args{&AuthQueryRequest{AuthId: "200000000000000002"}}, want: nil, wantErr: false},
		{name: "testCase-02-3", args: args{&AuthQueryRequest{AuthId: "200000000000000003"}}, want: nil, wantErr: false},
		{name: "testCase-02-4", args: args{&AuthQueryRequest{AuthId: "200000000000000004"}}, want: nil, wantErr: false},
		{name: "testCase-02-5", args: args{&AuthQueryRequest{AuthId: "200000000000000005"}}, want: nil, wantErr: false},
		{name: "testCase-02-6", args: args{&AuthQueryRequest{AuthId: "200000000000000006"}}, want: nil, wantErr: false},
		{name: "testCase-02-7", args: args{&AuthQueryRequest{AuthId: "200000000000000007"}}, want: nil, wantErr: false},
		{name: "testCase-02-8", args: args{&AuthQueryRequest{AuthId: "200000000000000008"}}, want: nil, wantErr: false},
		{name: "testCase-03-1", args: args{&AuthQueryRequest{AuthId: "300000000000000001"}}, want: nil, wantErr: false},
		{name: "testCase-03-2", args: args{&AuthQueryRequest{AuthId: "300000000000000002"}}, want: nil, wantErr: false},
		{name: "testCase-03-3", args: args{&AuthQueryRequest{AuthId: "300000000000000003"}}, want: nil, wantErr: false},
		{name: "testCase-03-4", args: args{&AuthQueryRequest{AuthId: "300000000000000004"}}, want: nil, wantErr: false},
		{name: "testCase-03-5", args: args{&AuthQueryRequest{AuthId: "300000000000000005"}}, want: nil, wantErr: false},
		{name: "testCase-03-6", args: args{&AuthQueryRequest{AuthId: "300000000000000006"}}, want: nil, wantErr: false},
		{name: "testCase-03-7", args: args{&AuthQueryRequest{AuthId: "300000000000000007"}}, want: nil, wantErr: false},
		{name: "testCase-03-8", args: args{&AuthQueryRequest{AuthId: "300000000000000008"}}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient("appid", "bizid", "key", "6Bkqhq")
			got, err := c.AuthQuery(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthQuery() got = %s, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_BehaviorCollect(t *testing.T) {

	type args struct {
		req *BehaviorRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "testCase-01-1",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbjzos82bs9cnyj1dna7d6d29zg4esnh99u"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345"),
					GovernmentPlatformUserId: "1fffbjzos82bs9cnyj1dna7d6d29zg4esnh99u",
				},
				{
					No:                       2,
					SessionId:                gofunc.Md5Lower("1fffbkmd9ebtwi7u7f4oswm9li6twjydqs7qjv"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbkmd9ebtwi7u7f4oswm9li6twjydqs7qjv"),
					GovernmentPlatformUserId: "1fffbkmd9ebtwi7u7f4oswm9li6twjydqs7qjv",
				},
				{
					No:                       3,
					SessionId:                gofunc.Md5Lower("1fffblf892i0p1zh6wlec2quukxtw29v4yismp"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffblf892i0p1zh6wlec2quukxtw29v4yismp"),
					GovernmentPlatformUserId: "1fffblf892i0p1zh6wlec2quukxtw29v4yismp",
				},
				{
					No:                       4,
					SessionId:                gofunc.Md5Lower("1fffbmr55j92gttv5wxspm0mgvw8x3p0n7cy0j"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbmr55j92gttv5wxspm0mgvw8x3p0n7cy0j"),
					GovernmentPlatformUserId: "1fffbmr55j92gttv5wxspm0mgvw8x3p0n7cy0j",
				},
				{
					No:                       5,
					SessionId:                gofunc.Md5Lower("1fffbjqfba5y6uwr55cdak6faokhm4s02qkyue"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbjqfba5y6uwr55cdak6faokhm4s02qkyue"),
					GovernmentPlatformUserId: "1fffbjqfba5y6uwr55cdak6faokhm4s02qkyue",
				},
				{
					No:                       6,
					SessionId:                gofunc.Md5Lower("1fffbkrwndszes1sngfx3v6pdqh87fi4zhz9ur"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbkrwndszes1sngfx3v6pdqh87fi4zhz9ur"),
					GovernmentPlatformUserId: "1fffbkrwndszes1sngfx3v6pdqh87fi4zhz9ur",
				},
				{
					No:                       7,
					SessionId:                gofunc.Md5Lower("1fffbl6st3fbp199i8zh5ggcp84fgo3rj7pn1y"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbl6st3fbp199i8zh5ggcp84fgo3rj7pn1y"),
					GovernmentPlatformUserId: "1fffbl6st3fbp199i8zh5ggcp84fgo3rj7pn1y",
				},
				{
					No:                       8,
					SessionId:                gofunc.Md5Lower("1fffbmzwmr1k3y8bri2linqbhnvmu510u5jj6z"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbmzwmr1k3y8bri2linqbhnvmu510u5jj6z"),
					GovernmentPlatformUserId: "1fffbmzwmr1k3y8bri2linqbhnvmu510u5jj6z",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-2",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbkmd9ebtwi7u7f4oswm9li6twjydqs7qjv"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbkmd9ebtwi7u7f4oswm9li6twjydqs7qjv"),
					GovernmentPlatformUserId: "1fffbkmd9ebtwi7u7f4oswm9li6twjydqs7qjv",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-3",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffblf892i0p1zh6wlec2quukxtw29v4yismp"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffblf892i0p1zh6wlec2quukxtw29v4yismp"),
					GovernmentPlatformUserId: "1fffblf892i0p1zh6wlec2quukxtw29v4yismp",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-4",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbmr55j92gttv5wxspm0mgvw8x3p0n7cy0j"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbmr55j92gttv5wxspm0mgvw8x3p0n7cy0j"),
					GovernmentPlatformUserId: "1fffbmr55j92gttv5wxspm0mgvw8x3p0n7cy0j",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-5",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbjqfba5y6uwr55cdak6faokhm4s02qkyue"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbjqfba5y6uwr55cdak6faokhm4s02qkyue"),
					GovernmentPlatformUserId: "1fffbjqfba5y6uwr55cdak6faokhm4s02qkyue",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-6",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbkrwndszes1sngfx3v6pdqh87fi4zhz9ur"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbkrwndszes1sngfx3v6pdqh87fi4zhz9ur"),
					GovernmentPlatformUserId: "1fffbkrwndszes1sngfx3v6pdqh87fi4zhz9ur",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-7",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbl6st3fbp199i8zh5ggcp84fgo3rj7pn1y"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbl6st3fbp199i8zh5ggcp84fgo3rj7pn1y"),
					GovernmentPlatformUserId: "1fffbl6st3fbp199i8zh5ggcp84fgo3rj7pn1y",
				},
			}}},
			want: nil, wantErr: false,
		},
		{
			name: "testCase-01-8",
			args: args{&BehaviorRequest{[]*BehaviorParamCollections{
				{
					No:                       1,
					SessionId:                gofunc.Md5Lower("1fffbmzwmr1k3y8bri2linqbhnvmu510u5jj6z"),
					BehaviorType:             1,
					OccurTime:                time.Now().Unix() - 10,
					CollectionType:           0,
					DeviceId:                 gofunc.Md5Lower("345" + "1fffbmzwmr1k3y8bri2linqbhnvmu510u5jj6z"),
					GovernmentPlatformUserId: "1fffbmzwmr1k3y8bri2linqbhnvmu510u5jj6z",
				},
			}}},
			want: nil, wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient("appid", "bizid", "key", "pB5757")
			got, err := c.BehaviorCollect(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BehaviorCollect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BehaviorCollect() got = %s, want %v", got, tt.want)
			}
		})
	}
}
