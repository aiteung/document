package drive

import (
	"context"
	"fmt"
	"github.com/JPratama7/gwrap"
	drive2 "github.com/JPratama7/gwrap/drive"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"log"
)

func UploadFileToDrive(filename, filepath, folderid, mimetype, credpath, tokenpath string) (fileid string, err error) {
	ctx := context.Background()
	cfg, err := gwrap.NewGoogleConfig(credpath, drive.DriveScope, drive.DriveReadonlyScope, docs.DocumentsScope, docs.DocumentsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v\n", err)
		return
	}
	client := gwrap.GetClient(cfg, tokenpath)

	srvDrive, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
		return
	}

	drivepath := fmt.Sprintf("%s", filename)

	drv := drive2.NewGoogleDrive(srvDrive)
	upld, err := drv.UploadFile(drivepath, mimetype, filepath, nil, folderid)
	if err != nil {
		log.Fatalf("Gagal Upload file: %v", err)
		return
	}
	return upld, nil
}
