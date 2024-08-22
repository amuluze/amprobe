// Package nginx
// Date: 2024/8/6 16:34
// Author: Amu
// Description:
package nginx

import (
	"fmt"
	"sort"
	"strings"
)

//var (
//	PrefixMatch   = "pre"
//	RegularMatch  = "reg"
//	CompleteMatch = "exact"
//)

type GeneratedConfig struct {
	SiteConfig map[string]string
}

func NewGeneratedConfig() *GeneratedConfig {
	return &GeneratedConfig{
		SiteConfig: make(map[string]string),
	}
}

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(siteGroup SiteGroup) (*GeneratedConfig, error) {
	config := NewGeneratedConfig()
	nc := NewConfig("")

	serverBuilders, upstreamMap, err := g.buildServers(siteGroup.Sites)
	if err != nil {
		return nil, err
	}

	upStreamBuilders, err := g.buildUpstreams(upstreamMap)
	if err != nil {
		return nil, err
	}

	if err = nc.ApplyBuilders(upStreamBuilders...); err != nil {
		return nil, err
	}

	if err = nc.ApplyBuilders(serverBuilders...); err != nil {
		return nil, err
	}

	siteConfig, err := nc.Generate()
	if err != nil {
		return nil, err
	}
	config.SiteConfig["siteConfig"] = siteConfig
	return config, nil
}

func (g *Generator) buildServers(sites []*Site) ([]ConfigBuilder, map[string]*Site, error) {
	var builders []ConfigBuilder
	serverMap := make(map[int][]*ServerBlockBuilder)
	upstreamMap := make(map[string]*Site)

	for _, site := range sites {
		upstreamName, err := GenerateUpstreamName(site)
		if err != nil {
			return nil, nil, err
		}
		upstreamMap[upstreamName] = site

		serverBlockBuilder, err := g.buildServer(site.ServerNames[0], site)
		if err != nil {
			return nil, nil, err
		}
		urlPathLocations, err := g.buildUrlPathLocation(site, upstreamName, site.UrlPaths)
		if err != nil {
			return nil, nil, err
		}
		locationMap := make(map[int][]*LocationBlockBuilder)
		locationMap[1] = urlPathLocations.LocationBuilders
		keys := g.addLocationsInOrder(locationMap, serverBlockBuilder)
		serverMap[keys[0]] = append(serverMap[keys[0]], serverBlockBuilder)
	}
	builders = g.addBuildsInOrder(serverMap, builders)
	return builders, upstreamMap, nil
}

func (g *Generator) buildServer(serverName string, site *Site) (*ServerBlockBuilder, error) {
	serverBlock := NewServerBlock()
	for _, listenConfig := range site.ListenConfigs {
		g.addListenOptionNormal(serverBlock, listenConfig)
	}
	serverBlock.AddServerNames(serverName)
	return serverBlock, nil
}

func (g *Generator) buildUpstreams(upstreamMap map[string]*Site) ([]ConfigBuilder, error) {
	var builders []ConfigBuilder
	for upstreamName, site := range upstreamMap {
		upstreamBlock := NewUpstreamBlock(upstreamName)
		proxyConfig := site.ProxyConfig

		// add servers
		for _, upstream := range proxyConfig.Upstreams {
			op := upstreamBlock.AddServer(upstream.Name())
			if upstream.Weight > 0 {
				op.Set("weight", fmt.Sprintf("%d", upstream.Weight))
			}
		}
		var keepalive int64 = 75
		var keepaliveTimeout int64 = 75
		if proxyConfig.Keepalive > 0 {
			keepalive = proxyConfig.Keepalive
		}
		if proxyConfig.KeepaliveTimeout > 0 {
			keepaliveTimeout = proxyConfig.KeepaliveTimeout
		}
		upstreamBlock.AddProperty("keepalive", fmt.Sprintf("%d", keepalive))
		upstreamBlock.AddProperty("keepalive_timeout", fmt.Sprintf("%d", keepaliveTimeout))
		builders = append(builders, upstreamBlock)
	}
	return builders, nil
}

func (g *Generator) buildUrlPathLocation(site *Site, upstreamName string, urlPaths []*UrlPath) (*MixBlockBuilder, error) {
	mixBlock := &MixBlockBuilder{}
	var locationBlocks []*LocationBlockBuilder
	proxyNames := GenerateProxyNames(site.ProxyConfig)
	for _, urlPath := range urlPaths {
		locationBlock := NewLocationBlock("^~" + " " + urlPath.UrlPath)
		locationBlock.
			AddPropertyB(proxyNames["proxy_pass"], fmt.Sprintf("%s://%s", strings.ToLower(proxyNames["scheme"]), upstreamName))

		locationBlocks = append(locationBlocks, locationBlock)
	}

	mixBlock.LocationBuilders = locationBlocks
	return mixBlock, nil
}

func (g *Generator) addListenOptionNormal(block *ServerBlockBuilder, listenConfig *ListenConfig) *PropertyBuilder {
	return block.AddListen(listenConfig.ListenAddress())
}

func (g *Generator) addBuildsInOrder(serverMap map[int][]*ServerBlockBuilder, builders []ConfigBuilder) []ConfigBuilder {
	var mapKeys []int
	for k := range serverMap {
		mapKeys = append(mapKeys, k)
	}
	sort.Ints(mapKeys)
	for _, k := range mapKeys {
		for _, builder := range serverMap[k] {
			builders = append(builders, builder)
		}
	}
	return builders
}

func (g *Generator) addLocationsInOrder(locationMap map[int][]*LocationBlockBuilder, serverBlockBuilder *ServerBlockBuilder) []int {
	var keys []int
	for k := range locationMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		urlPathLocationList := locationMap[k]
		for _, urlPathLocation := range urlPathLocationList {
			serverBlockBuilder.AddLocationB(urlPathLocation)
		}
	}
	return keys
}
