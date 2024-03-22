package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/RiemaLabs/nubit-da-sdk/constant"
	"github.com/RiemaLabs/nubit-da-sdk/log"
	"github.com/RiemaLabs/nubit-da-sdk/types"
)

// CountUsers indexer-countUsers
func (c *Client) CountUsers(ctx context.Context) (data *types.CountRsp, err error) {
	path, err := url.JoinPath(c.c.Endpoint, c.u.GetUri(constant.CountUsers))
	if err != nil {
		log.Error("client.Users", "method", "indexer-CountUsers", "JoinPath", err)
		return nil, err
	}
	post, err := c.c.Post(ctx, path, nil, nil)
	if err != nil {
		log.Error("client.Users", "method", "indexer-CountUsers", "getNamespaceStatusReq", err)
		return nil, err
	}

	err = json.Unmarshal(post, &data)
	if err != nil {
		log.Error("client.Users", "method", "indexer-CountUsers", "Unmarshal", err)
		return nil, err
	}
	return data, err
}
