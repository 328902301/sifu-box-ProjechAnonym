package execute

import (
	"fmt"
	"net/url"
	"path/filepath"
	database "sifu-box/Database"
	utils "sifu-box/Utils"
	"sync"
)

// Exec_update 更新指定标签的代理配置,涉及配置备份、更新、验证及数据库记录同步
// 参数:
// - label: 要更新的代理配置标签
// - proxy_config: 当前的代理配置信息
// - server: 数据库中的服务器实体,包含配置部署所需信息
// - specific: 是否为特定服务器更新
// - lock: 互斥锁,用于控制并发更新
// 返回:
// - error: 更新过程中遇到的任何错误
func Exec_update(label string, proxy_config utils.Box_config, server database.Server,specific bool,lock *sync.Mutex) error {
	// 特定更新的话则上锁
	if specific{
		for {
			if lock.TryLock(){
				break
			}
		}
		defer lock.Unlock()
	}
	// 获取项目目录以备备份配置文件使用
	projectDir, err := utils.Get_value("project-dir")
	if err != nil {
		utils.Logger_caller("get project dir failed", err, 1)
		return err
	}

	// 确认待更新的标签存在于代理配置中
	label_exist := false
	for _, proxy := range proxy_config.Url {
		if proxy.Label == label {
			label_exist = true
			break
		}
	}
	if !label_exist {
		return fmt.Errorf("specific label %s is not in the current proxy config", label)
	}

	// 对标签进行MD5加密,准备新配置文件名
	new_file, err := utils.Encryption_md5(label)
	if err != nil {
		utils.Logger_caller("encryption failed", err, 1)
		return err
	}

	// 解析服务器URL,构建备份文件名
	link, err := url.Parse(server.Url)
	if err != nil {
		utils.Logger_caller("parse server url failed", err, 1)
		return err
	}
	backup_file := link.Hostname()

	// 准备各文件路径
	original_path := "/opt/singbox/config.json"
	backup_path := filepath.Join(projectDir.(string), "temp", "configbackup", backup_file+".json")
	new_path := filepath.Join(projectDir.(string), "static", "Default", new_file+".json")

	// 更新配置文件并创建备份
	if err := Update_file(original_path, new_path, backup_path, 0644, server); err != nil {
		return err
	}

	// 尝试重载配置并验证
	if result, err := Reload_config("sing-box", server); err != nil || !result {

		// 配置重载失败时恢复备份配置
		if recoverErr := Recover_file(original_path, backup_path, 0644, server); recoverErr != nil {
			return recoverErr
		}

		// 尝试启动服务确保服务状态正常
		if startErr := Boot_service("sing-box", server); startErr != nil {
			return startErr
		}

		return fmt.Errorf("reload new config failed")
	}

	// 成功后更新数据库中的服务器配置标签
	if err := database.Db.Model(&server).Where("url = ?", server.Url).Update("config", label).Error; err != nil {
		utils.Logger_caller("update database data fail", err, 1)
		return err
	}

	return nil
}

func Group_update(servers []database.Server, proxy_config utils.Box_config,lock *sync.Mutex){
	for {
		if lock.TryLock(){
			break
		}
	}
	defer lock.Unlock()
	var servers_workflow sync.WaitGroup
	for _, server := range servers {
		servers_workflow.Add(1)
		go func(server database.Server){
			defer servers_workflow.Done()
			if err := Exec_update(server.Config, proxy_config, server,false,lock); err != nil {
				utils.Logger_caller("update servers config failed", err, 1)
			}
		}(server)
	}
	servers_workflow.Wait()
}