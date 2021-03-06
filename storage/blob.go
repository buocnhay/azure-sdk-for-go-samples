// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package storage

import (
	"context"
	"io/ioutil"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func getBlobURL(ctx context.Context, accountName, accountGroupName, containerName, blobName string) azblob.BlobURL {
	container := getContainerURL(ctx, accountName, accountGroupName, containerName)
	blob := container.NewBlobURL(blobName)
	return blob
}

// GetBlob downloads the specified blob contents
func GetBlob(ctx context.Context, accountName, accountGroupName, containerName, blobName string) (string, error) {
	b := getBlobURL(ctx, accountName, accountGroupName, containerName, blobName)

	resp, err := b.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)

	if err != nil {
		return "", err
	}
	defer resp.Response().Body.Close()
	body, err := ioutil.ReadAll(resp.Body(azblob.RetryReaderOptions{}))
	return string(body), err
}
