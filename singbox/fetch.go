package singbox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sifu-box/models"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func fetchProviderInfo(provider models.Provider, client *http.Client, logger *zap.Logger) (string, error) {
	if provider.Remote {
		fetchFromRemote(provider, client, logger)
	}else{
		fetchFromLocal(provider, logger)
	}
	
		// fmt.Println(outbounds)
		// b , _ := json.Marshal(outbounds)
		// fmt.Println(string(b))
		
	
	return "", nil
}

func fetchFromRemote(provider models.Provider, client *http.Client, logger *zap.Logger) (string, error) {
	req, err := http.NewRequest("GET", provider.Path, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("创建请求失败: [%s]", err.Error()))
		return "", fmt.Errorf("创建请求失败")
	}
	res, err := client.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("发送请求失败: [%s]", err.Error()))
		return "", fmt.Errorf("发送请求失败")
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		content, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("读取响应失败: [%s]", err.Error()))
			return "", fmt.Errorf("读取响应失败")
		}
		parseFileContent(content, logger)
	}
	return "", nil
}

func fetchFromLocal(provider models.Provider, logger *zap.Logger) (string, error) {
	file, err := os.Open(provider.Path)
	if err != nil {
		logger.Error(fmt.Sprintf("打开'%s'文件失败: [%s]", provider.Name, err.Error()))
		return "", fmt.Errorf("打开'%s'文件失败", provider.Name)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		logger.Error(fmt.Sprintf("读取'%s'文件失败: [%s]", provider.Name, err.Error()))
		return "", fmt.Errorf("读取'%s'文件失败", provider.Name)
	}
	parseFileContent(content, logger)
	return "", nil
}

func parseFileContent(content []byte, logger *zap.Logger) (string, error) {
	var providerInfo map[string]interface{}
	if err := yaml.Unmarshal(content, &providerInfo); err != nil {
		logger.Error(fmt.Sprintf("解析响应失败: [%s]", err.Error()))
		return "", fmt.Errorf("解析响应失败")
	}
	var outbounds []models.Outbound
	for _, proxy := range providerInfo["proxies"].([]interface{}) {
		switch proxy.(map[string]interface{})["type"].(string) {
			case "ss":
				shadowSocks := models.ShadowSocks{}
				err := error(nil)
				var outbound models.Outbound = &shadowSocks
				outbound, err = outbound.Transform(proxy.(map[string]interface{}), logger)
				if err != nil {
					logger.Error(fmt.Sprintf("'%s'节点解析ShadowSocks代理失败: [%s]", proxy.(map[string]interface{})["name"].(string), err.Error()))
					continue
				}
				outbounds = append(outbounds, outbound)
			case "vmess":
				vmess := models.VMess{}
				err := error(nil)
				var outbound models.Outbound = &vmess
				outbound, err = outbound.Transform(proxy.(map[string]interface{}), logger)
				if err != nil {
					logger.Error(fmt.Sprintf("'%s'节点解析Vmess代理失败: [%s]", proxy.(map[string]interface{})["name"].(string), err.Error()))
					continue
				}
				outbounds = append(outbounds, outbound)
			case "trojan":
				trojan := models.Trojan{}
				err := error(nil)
				var outbound models.Outbound = &trojan
				outbound, err = outbound.Transform(proxy.(map[string]interface{}), logger)
				if err != nil {
					logger.Error(fmt.Sprintf("'%s'节点解析Trojan代理失败: [%s]", proxy.(map[string]interface{})["name"].(string), err.Error()))
					continue
				}
				outbounds = append(outbounds, outbound)
		}
	}
	


	a,_ := json.Marshal(outbounds)
	fmt.Println(string(a))
	// a,_ := json.Marshal(outbounds)
	// fmt.Println(string(a))
	return "", nil
}