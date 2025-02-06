package controller

import (
	"github.com/kataras/iris/v12"
	"note-iris/src/repository"
)

// One 根据id查询命令
func One(ctx iris.Context) {
	command := repository.One(ctx)
	ctx.JSON(command)
}

// List 查询命令列表
func List(ctx iris.Context) {
	commandArray := repository.List(ctx)
	ctx.JSON(commandArray)
}

// Insert 新增命令
func Insert(ctx iris.Context) {
	result, commandName := repository.Insert(ctx)
	ctx.JSON(iris.Map{
		"result":  result,
		"command": commandName,
	})
}

// InsertBatch 批量新增命令
func InsertBatch(ctx iris.Context) {
	result := repository.InsertBatch(ctx)
	ctx.JSON(result)
}

// Update 修改命令
func Update(ctx iris.Context) {
	result := repository.Update(ctx)
	ctx.JSON(result)
}

// UpdateBatch 批量修改命令
func UpdateBatch(ctx iris.Context) {
	result := repository.UpdateBatch(ctx)
	ctx.JSON(result)
}

// Delete 删除命令
func Delete(ctx iris.Context) {
	result, objectId := repository.Delete(ctx)
	ctx.JSON(iris.Map{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteBatch 批量删除命令
func DeleteBatch(ctx iris.Context) {
	result, objectIds := repository.DeleteBatch(ctx)
	ctx.JSON(iris.Map{
		"result": result,
		"_ids":   objectIds,
	})
}

// Select 查询命令
func Select(ctx iris.Context) {
	command := repository.Select(ctx)
	ctx.JSON(command)
}

// NameList 查询命令名称列表
func NameList(ctx iris.Context) {
	nameArray := repository.NameList(ctx)
	ctx.JSON(nameArray)
}
