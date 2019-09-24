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

package service

import (
	"strconv"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/topo_server/core/types"

	"github.com/mitchellh/mapstructure"
)

func (s *Service) CreateSetTemplate(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	option := metadata.CreateSetTemplateOption{}
	if err := data.MarshalJSONInto(&option); err != nil {
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	setTemplate, err := s.Engine.CoreAPI.CoreService().SetTemplate().CreateSetTemplate(params.Context, params.Header, bizID, option)
	if err != nil {
		blog.Errorf("CreateSetTemplate failed, core service create failed, bizID: %d, option: %+v, err: %+v, rid: %s", bizID, option, err, params.ReqID)
		return nil, err
	}
	return setTemplate, nil
}

func (s *Service) UpdateSetTemplate(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplateIDStr := pathParams(common.BKSetTemplateIDField)
	setTemplateID, err := strconv.ParseInt(setTemplateIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKSetTemplateIDField)
	}

	option := metadata.UpdateSetTemplateOption{}
	if err := data.MarshalJSONInto(&option); err != nil {
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	setTemplate, err := s.Engine.CoreAPI.CoreService().SetTemplate().UpdateSetTemplate(params.Context, params.Header, bizID, setTemplateID, option)
	if err != nil {
		blog.Errorf("UpdateSetTemplate failed, do core service update failed, bizID: %d, option: %+v, err: %+v, rid: %s", bizID, option, err, params.ReqID)
		return nil, err
	}
	return setTemplate, nil
}

func (s *Service) DeleteSetTemplate(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	option := metadata.DeleteSetTemplateOption{}
	if err := data.MarshalJSONInto(&option); err != nil {
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	if err := s.Engine.CoreAPI.CoreService().SetTemplate().DeleteSetTemplate(params.Context, params.Header, bizID, option); err != nil {
		blog.Errorf("DeleteSetTemplate failed, do core service update failed, bizID: %d, option: %+v, err: %+v, rid: %s", bizID, option, err, params.ReqID)
		return nil, err
	}
	return nil, nil
}

func (s *Service) GetSetTemplate(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplateIDStr := pathParams(common.BKSetTemplateIDField)
	setTemplateID, err := strconv.ParseInt(setTemplateIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKSetTemplateIDField)
	}

	setTemplate, err := s.Engine.CoreAPI.CoreService().SetTemplate().GetSetTemplate(params.Context, params.Header, bizID, setTemplateID)
	if err != nil {
		blog.Errorf("GetSetTemplate failed, do core service get failed, bizID: %d, setTemplateID: %d, err: %+v, rid: %s", bizID, setTemplateID, err, params.ReqID)
		return nil, err
	}
	return setTemplate, nil
}

func (s *Service) ListSetTemplate(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	option := metadata.ListSetTemplateOption{}
	if err := data.MarshalJSONInto(&option); err != nil {
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	setTemplate, err := s.Engine.CoreAPI.CoreService().SetTemplate().ListSetTemplate(params.Context, params.Header, bizID, option)
	if err != nil {
		blog.Errorf("ListSetTemplate failed, do core service list failed, bizID: %d, option: %+v, err: %+v, rid: %s", bizID, option, err, params.ReqID)
		return nil, err
	}
	return setTemplate, nil
}

func (s *Service) ListSetTemplateWeb(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplates, err := s.ListSetTemplate(params, pathParams, queryParams, data)
	if err != nil {
		return nil, err
	}
	listResult := setTemplates.(*metadata.MultipleSetTemplateResult)
	if listResult == nil {
		return nil, nil
	}

	// count template instances
	setTemplateIDs := make([]int64, 0)
	for _, item := range listResult.Info {
		setTemplateIDs = append(setTemplateIDs, item.ID)
	}
	option := metadata.CountSetTplInstOption{
		SetTemplateIDs: setTemplateIDs,
	}
	setTplInstCount, err := s.Engine.CoreAPI.CoreService().SetTemplate().CountSetTplInstances(params.Context, params.Header, bizID, option)
	if err != nil {
		blog.Errorf("ListSetTemplateWeb failed, CountSetTplInstances failed, bizID: %d, option: %+v, err: %s, rid: %s", bizID, option, err.Error(), params.ReqID)
		return nil, err
	}
	result := metadata.MultipleSetTemplateWithStatisticsResult{
		Count: listResult.Count,
	}
	for _, setTemplate := range listResult.Info {
		setInstanceCount, exist := setTplInstCount[setTemplate.ID]
		if exist == false {
			setInstanceCount = 0
		}
		result.Info = append(result.Info, metadata.SetTemplateWithStatistics{
			SetInstanceCount: setInstanceCount,
			SetTemplate:      setTemplate,
		})
	}
	return result, nil
}

func (s *Service) ListSetTplRelatedSvcTpl(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplateIDStr := pathParams(common.BKSetTemplateIDField)
	setTemplateID, err := strconv.ParseInt(setTemplateIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKSetTemplateIDField)
	}

	serviceTemplates, err := s.Engine.CoreAPI.CoreService().SetTemplate().ListSetTplRelatedSvcTpl(params.Context, params.Header, bizID, setTemplateID)
	if err != nil {
		blog.Errorf("ListSetTemplateRelatedServiceTemplate failed, do core service list failed, bizID: %d, setTemplateID: %+v, err: %+v, rid: %s", bizID, setTemplateID, err, params.ReqID)
		return nil, err
	}

	return serviceTemplates, nil
}

// ListSetTplRelatedSets get SetTemplate related sets
func (s *Service) ListSetTplRelatedSets(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplateIDStr := pathParams(common.BKSetTemplateIDField)
	setTemplateID, err := strconv.ParseInt(setTemplateIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKSetTemplateIDField)
	}

	option := metadata.ListSetByTemplateOption{}
	if err := data.MarshalJSONInto(&option); err != nil {
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	filter := &metadata.QueryCondition{
		Fields: nil,
		Limit: metadata.SearchLimit{
			Offset: int64(option.Page.Start),
			Limit:  int64(option.Page.Limit),
		},
		SortArr: nil,
		Condition: mapstr.MapStr(map[string]interface{}{
			common.BKAppIDField:         bizID,
			common.BKSetTemplateIDField: setTemplateID,
		}),
	}
	setInstanceResult, err := s.Engine.CoreAPI.CoreService().Instance().ReadInstance(params.Context, params.Header, common.BKInnerObjIDSet, filter)
	if err != nil {
		blog.Errorf("ListSetTplRelatedSetsWeb failed, err: %+v, rid: %s", err, params.ReqID)
		return nil, err
	}
	return setInstanceResult.Data, nil
}

// ListSetTplRelatedSetsWeb get SetTemplate related sets, just for web
func (s *Service) ListSetTplRelatedSetsWeb(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	response, err := s.ListSetTplRelatedSets(params, pathParams, queryParams, data)
	if err != nil {
		return nil, err
	}
	setInstanceResult := response.(metadata.InstDataInfo)

	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	topoTree, err := s.Engine.CoreAPI.CoreService().Mainline().SearchMainlineInstanceTopo(params.Context, params.Header, bizID, false)
	if err != nil {
		blog.Errorf("ListSetTplRelatedSetsWeb failed, bizID: %d, err: %s, rid: %s", bizID, err.Error(), params.ReqID)
		return nil, err
	}

	for index := range setInstanceResult.Info {
		set := metadata.SetInst{}
		if err := mapstr.DecodeFromMapStr(&set, setInstanceResult.Info[index]); err != nil {
			return nil, params.Err.Error(common.CCErrCommJSONUnmarshalFailed)
		}

		setPath := topoTree.TraversalFindNode(common.BKInnerObjIDSet, set.SetID)
		topoPath := make([]metadata.TopoInstanceNodeSimplify, 0)
		for _, pathNode := range setPath {
			nodeSimplify := metadata.TopoInstanceNodeSimplify{
				ObjectID:     pathNode.ObjectID,
				InstanceID:   pathNode.InstanceID,
				InstanceName: pathNode.InstanceName,
			}
			topoPath = append(topoPath, nodeSimplify)
		}
		setInstanceResult.Info[index]["topo_path"] = topoPath
	}

	return setInstanceResult, nil
}

func (s *Service) DiffSetTplWithInst(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplateIDStr := pathParams(common.BKSetTemplateIDField)
	setTemplateID, err := strconv.ParseInt(setTemplateIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKSetTemplateIDField)
	}

	option := metadata.DiffSetTplWithInstOption{}
	if err := mapstructure.Decode(data, &option); err != nil {
		blog.Errorf("DiffSetTemplateWithInstances failed, decode request body failed, data: %+v, err: %+v, rid: %s", data, err, params.ReqID)
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	setDiffs, err := s.Core.SetTemplateOperation().DiffSetTplWithInst(params.Context, params.Header, bizID, setTemplateID, option)
	if err != nil {
		blog.Errorf("DiffSetTplWithInst failed, operation failed, bizID: %d, setTemplateID: %d, option: %+v err: %s, rid: %s", bizID, setTemplateID, option, err.Error(), params.ReqID)
		return nil, err
	}

	return setDiffs, nil
}

func (s *Service) SyncSetTplToInst(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (output interface{}, retErr error) {
	bizIDStr := pathParams(common.BKAppIDField)
	bizID, err := strconv.ParseInt(bizIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKAppIDField)
	}

	setTemplateIDStr := pathParams(common.BKSetTemplateIDField)
	setTemplateID, err := strconv.ParseInt(setTemplateIDStr, 10, 64)
	if err != nil {
		return nil, params.Err.CCErrorf(common.CCErrCommParamsInvalid, common.BKSetTemplateIDField)
	}

	option := metadata.SyncSetTplToInstOption{}
	if err := mapstructure.Decode(data, &option); err != nil {
		blog.Errorf("DiffSetTemplateWithInstances failed, decode request body failed, data: %+v, err: %+v, rid: %s", data, err, params.ReqID)
		return nil, params.Err.CCError(common.CCErrCommJSONUnmarshalFailed)
	}

	err = s.Core.SetTemplateOperation().SyncSetTplToInst(params.Context, params.Header, bizID, setTemplateID, option)
	if err != nil {
		blog.Errorf("SyncSetTplToInst failed, operation failed, bizID: %d, setTemplateID: %d, option: %+v err: %s, rid: %s", bizID, setTemplateID, option, err.Error(), params.ReqID)
		return nil, err
	}

	return nil, nil
}
