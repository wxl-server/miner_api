// Code generated by hertz generator. DO NOT EDIT.

package model

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	model "miner_api/biz/handler/model"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_indicator := root.Group("/indicator", _indicatorMw()...)
		{
			_query := _indicator.Group("/query", _queryMw()...)
			_query.GET("/list", append(_queryindicatorlistMw(), model.QueryIndicatorList)...)
		}
	}
	{
		_job := root.Group("/job", _jobMw()...)
		_job.POST("/create", append(_createjobMw(), model.CreateJob)...)
		_job.POST("/delete", append(_deletejobMw(), model.DeleteJob)...)
		{
			_query0 := _job.Group("/query", _query0Mw()...)
			_query0.POST("/list", append(_queryjoblistMw(), model.QueryJobList)...)
		}
	}
	{
		_task := root.Group("/task", _taskMw()...)
		_task.GET("/run", append(_runtaskMw(), model.RunTask)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/login", append(_loginMw(), model.Login)...)
		_user.POST("/signup", append(_signupMw(), model.SignUp)...)
		{
			_query1 := _user.Group("/query", _query1Mw()...)
			_query1.GET("/list", append(_queryuserlistMw(), model.QueryUserList)...)
		}
	}
}
