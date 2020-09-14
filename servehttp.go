package cloudfunctionsdemo

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Write([]byte("Hello, " + name))
	bucket := os.Getenv("bucketName")
	latest, err := download(bucket, "latest")
	if err != nil {
		w.Write([]byte("An error occured"))
	}
	latestBytes, _ := ioutil.ReadAll(latest)
	w.Write([]byte("Latest hello was to"))
	w.Write(latestBytes)

	upload(bucket, "latest", strings.NewReader(name))

}

func download(bucket, object string) (io.Reader, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()
	return client.Bucket(bucket).Object(object).NewReader(ctx)
}

func upload(bucket string, object string, r io.Reader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, r); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}
