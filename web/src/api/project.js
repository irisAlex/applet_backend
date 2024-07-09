import service from "@/utils/request";
// @Tags api
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getProjectList = (data) => {
  return service({
    url: "/project/getAllProjectList",
    method: "post",
    data,
  });
};

// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
export const createProject = (data) => {
  return service({
    url: "/project/createProject",
    method: "post",
    data,
  });
};

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getApiById [post]
export const getProjectById = (data) => {
  return service({
    url: "/project/getProjectById",
    method: "post",
    data,
  });
};

// @Tags Api
// @Summary 更新api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "更新api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/updateApi [post]
export const updateProject = (data) => {
  return service({
    url: "/project/updateProject",
    method: "post",
    data,
  });
};

export const deleteProject = (data) => {
  return service({
    url: "/project/deleteProject",
    method: "post",
    data,
  });
};
