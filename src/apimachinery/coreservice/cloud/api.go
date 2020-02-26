/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cloud

import (
	"context"
	"net/http"

	"configcenter/src/common/blog"
	"configcenter/src/common/errors"
	"configcenter/src/common/metadata"
)

func (c *cloud) CreateAccount(ctx context.Context, h http.Header, account *metadata.CloudAccount) (*metadata.CloudAccount, errors.CCErrorCoder) {
	ret := new(metadata.CloudAccountResult)
	subPath := "/create/cloud/account"

	err := c.client.Post().
		WithContext(ctx).
		Body(account).
		SubResourcef(subPath).
		WithHeaders(h).
		Do().
		Into(ret)

	if err != nil {
		blog.Errorf("CreateAccount failed, http request failed, err: %+v", err)
		return nil, errors.CCHttpError
	}
	if ret.Result == false || ret.Code != 0 {
		return nil, errors.New(ret.Code, ret.ErrMsg)
	}

	return &ret.Data, nil
}

func (c *cloud) SearchAccount(ctx context.Context, h http.Header, option *metadata.SearchCloudAccountOption) (*metadata.MultipleCloudAccount, errors.CCErrorCoder) {
	ret := new(metadata.MultipleCloudAccountResult)
	subPath := "/search/cloud/account"

	err := c.client.Post().
		WithContext(ctx).
		Body(option).
		SubResourcef(subPath).
		WithHeaders(h).
		Do().
		Into(ret)

	if err != nil {
		blog.Errorf("SearchAccount failed, http request failed, err: %+v", err)
		return nil, errors.CCHttpError
	}
	if ret.Result == false || ret.Code != 0 {
		return nil, errors.New(ret.Code, ret.ErrMsg)
	}

	return &ret.Data, nil
}

func (c *cloud) UpdateAccount(ctx context.Context, h http.Header, accountID int64, account *metadata.CloudAccount) (*metadata.CloudAccount, errors.CCErrorCoder) {
	ret := new(metadata.CloudAccountResult)
	subPath := "/update/cloud/account/%d"

	err := c.client.Put().
		WithContext(ctx).
		Body(account).
		SubResourcef(subPath, accountID).
		WithHeaders(h).
		Do().
		Into(ret)

	if err != nil {
		blog.Errorf("UpdateAccount failed, http request failed, err: %+v", err)
		return nil, errors.CCHttpError
	}
	if ret.Result == false || ret.Code != 0 {
		return nil, errors.New(ret.Code, ret.ErrMsg)
	}

	return &ret.Data, nil
}

func (c *cloud) DeleteAccount(ctx context.Context, h http.Header, accountID int64) errors.CCErrorCoder {
	ret := new(metadata.CloudAccountResult)
	subPath := "/delete/cloud/account/%d"

	err := c.client.Delete().
		WithContext(ctx).
		SubResourcef(subPath, accountID).
		WithHeaders(h).
		Do().
		Into(ret)

	if err != nil {
		blog.Errorf("DeleteAccount failed, http request failed, err: %+v", err)
		return errors.CCHttpError
	}
	if ret.Result == false || ret.Code != 0 {
		return errors.New(ret.Code, ret.ErrMsg)
	}

	return nil
}