package process

import "minik8s/pkg/ty"

// 增删改查

func AddReplicaSet(desc *ty.ReplicaSetDesc) error {
	// 检查是否已经存在

	// 写入etcd

	// 应该没有必要直接创建，而是让它自己检查到pod实际数量不一致的时候再改
	return nil
}

func RemoveReplicaSet(namespace string, name string) error {
	return nil
}

func ModifyReplicaSet(namespace string, name string) error {
	// 修改应该只能修改num，pod的信息应该是不方便改的，真要换pod应该先删除再创建
	return nil
}

func GetReplicaSet(namespace string, name string) (string, error) {
	// 这里需要传回哪些东西，用什么格式？

	return "", nil
}
