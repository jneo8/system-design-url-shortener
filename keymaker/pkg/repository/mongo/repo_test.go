package mongo

import (
	"context"
	"jneo8/system-design-url-shortener/keymaker/pkg/repository/mongo/internal/testutil"
	"reflect"
	"testing"
	"time"

	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var logger = mermaid.NewLogger()

func Test_repo_Close(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip connection test in short mode")
	}

	// Context to handler connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Basic",
			fields: fields{
				Client: func() *mongo.Client {
					client, err := mongo.Connect(
						context.Background(),
						options.Client().ApplyURI(testutil.MongoConnectionStr),
					)
					require.NoError(t, err)
					return client
				}(),
				Logger:        logger,
				KeyCollection: testutil.KeyCollection,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.fields.Client.Disconnect(ctx)
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			if err := r.Client.Ping(ctx, readpref.Primary()); err != nil {
				t.Error(err)
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("repo.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := r.Client.Ping(ctx, readpref.Primary()); err == nil {
				t.Error("Connect() still exists")
			}
		})
	}
}

func Test_repo_Init(t *testing.T) {
	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			if err := r.Init(); (err != nil) != tt.wantErr {
				t.Errorf("repo.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_createIndexes(t *testing.T) {
	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			if err := r.createIndexes(); (err != nil) != tt.wantErr {
				t.Errorf("repo.createIndexes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_KeyBatchInsert(t *testing.T) {
	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			got, err := r.KeyBatchInsert(tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.KeyBatchInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repo.KeyBatchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_KeyBatchUpsert(t *testing.T) {
	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			got, err := r.KeyBatchUpsert(tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.KeyBatchUpsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repo.KeyBatchUpsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_keyCollection(t *testing.T) {
	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}
	tests := []struct {
		name   string
		fields fields
		want   *mongo.Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			if got := r.keyCollection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.keyCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_GetKey(t *testing.T) {
	type fields struct {
		Config        Config
		Client        *mongo.Client
		Logger        *log.Logger
		KeyCollection string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				Config:        tt.fields.Config,
				Client:        tt.fields.Client,
				Logger:        tt.fields.Logger,
				KeyCollection: tt.fields.KeyCollection,
			}
			got, err := r.GetKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
