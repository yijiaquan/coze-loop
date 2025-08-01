// Copyright (c) 2025 coze-dev Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"
	"fmt"

	"github.com/coze-dev/coze-loop/backend/modules/observability/domain/component/config"
	"github.com/coze-dev/coze-loop/backend/pkg/conf"
	"github.com/coze-dev/coze-loop/backend/pkg/logs"
)

const (
	systemViewsCfgKey         = "trace_system_view_cfg"
	platformTenantCfgKey      = "trace_platform_tenants"
	platformSpanHandlerCfgKey = "trace_platform_span_handler_config"
	traceMqProducerCfgKey     = "trace_mq_producer_config"
	tenantTablesCfgKey        = "trace_tenant_tables_cfg"
	traceCkCfgKey             = "trace_ck_cfg"
	traceFieldMetaInfoCfgKey  = "trace_field_meta_info"
	traceAttrTosCfgKey        = "trace_attr_tos_cfg"
	traceMaxDurationDay       = "trace_max_duration_day"
	traceDefaultTenant        = "trace_default_tenant"
)

type TraceConfigCenter struct {
	conf.IConfigLoader
	// glocal config, just in case
	traceDefaultTenant string
}

func (t *TraceConfigCenter) GetSystemViews(ctx context.Context) ([]*config.SystemView, error) {
	systemViews := make([]*config.SystemView, 0)
	if err := t.UnmarshalKey(ctx, systemViewsCfgKey, &systemViews); err != nil {
		return nil, err
	}
	return systemViews, nil
}

func (t *TraceConfigCenter) GetPlatformTenants(ctx context.Context) (*config.PlatformTenantsCfg, error) {
	cfg := new(config.PlatformTenantsCfg)
	if err := t.UnmarshalKey(ctx, platformTenantCfgKey, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (t *TraceConfigCenter) GetPlatformSpansTrans(ctx context.Context) (*config.SpanTransHandlerConfig, error) {
	cfg := new(config.SpanTransHandlerConfig)
	if err := t.UnmarshalKey(ctx, platformSpanHandlerCfgKey, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (t *TraceConfigCenter) GetTraceMqProducerCfg(ctx context.Context) (*config.TraceMqProducerCfg, error) {
	cfg := new(config.TraceMqProducerCfg)
	if err := t.UnmarshalKey(context.Background(), traceMqProducerCfgKey, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (t *TraceConfigCenter) GetTraceCkCfg(ctx context.Context) (*config.TraceCKCfg, error) {
	cfg := new(config.TraceCKCfg)
	if err := t.UnmarshalKey(context.Background(), traceCkCfgKey, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (t *TraceConfigCenter) GetTenantConfig(ctx context.Context) (*config.TenantCfg, error) {
	tenantTableCfg := new(config.TenantCfg)
	if err := t.UnmarshalKey(ctx, tenantTablesCfgKey, &tenantTableCfg); err != nil {
		return nil, err
	}
	return tenantTableCfg, nil
}

func (t *TraceConfigCenter) GetTraceFieldMetaInfo(ctx context.Context) (*config.TraceFieldMetaInfoCfg, error) {
	traceFieldMetaInfoCfg := new(config.TraceFieldMetaInfoCfg)
	if err := t.UnmarshalKey(ctx, traceFieldMetaInfoCfgKey, &traceFieldMetaInfoCfg); err != nil {
		return nil, err
	}
	return traceFieldMetaInfoCfg, nil
}

func (t *TraceConfigCenter) GetTraceAttrTosCfg(ctx context.Context) (*config.TraceAttrTosCfg, error) {
	traceAttrTosCfg := new(config.TraceAttrTosCfg)
	if err := t.UnmarshalKey(ctx, traceAttrTosCfgKey, &traceAttrTosCfg); err != nil {
		return nil, err
	}
	return traceAttrTosCfg, nil
}

func (t *TraceConfigCenter) GetTraceDataMaxDurationDay(ctx context.Context, platformPtr *string) int64 {
	defaultDuration := int64(7)
	if platformPtr == nil {
		return defaultDuration
	}
	platformType := *platformPtr
	mp := make(map[string]int64)
	err := t.UnmarshalKey(ctx, traceMaxDurationDay, &mp)
	if err != nil {
		logs.CtxWarn(ctx, "fail to unmarshal max duration cfg, %v", err)
		return defaultDuration
	}
	if mp[platformType] > 0 {
		return mp[platformType]
	} else {
		return defaultDuration
	}
}

func (t *TraceConfigCenter) GetDefaultTraceTenant(ctx context.Context) string {
	return t.traceDefaultTenant
}

func (t *TraceConfigCenter) getDefaultTraceTenant(ctx context.Context) (string, error) {
	if t.traceDefaultTenant != "" {
		return t.traceDefaultTenant, nil
	}
	val := t.Get(ctx, traceDefaultTenant)
	if val == nil {
		return "", fmt.Errorf("trace tenant not configured")
	}
	ret, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("trace tenant not string type")
	} else if ret == "" {
		return "", fmt.Errorf("trace tenant is empty")
	}
	return ret, nil
}

func NewTraceConfigCenter(confP conf.IConfigLoader) config.ITraceConfig {
	ret := &TraceConfigCenter{
		IConfigLoader: confP,
	}
	tenant, err := ret.getDefaultTraceTenant(context.Background())
	if err != nil {
		panic(err)
	}
	ret.traceDefaultTenant = tenant
	return ret
}
