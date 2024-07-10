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
export const getPostList = (data) => {
  return service({
    url: "/Post/getAllPostList",
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
export const createPost = (data) => {
  return service({
    url: "/Post/createPost",
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
export const getPostById = (data) => {
  return service({
    url: "/Post/getPostById",
    method: "post",
    data,
  });
};

export const getSubjectByEId = (data) => {
  return service({
    url: "/Subject/getSubjectByEID",
    method: "post",
    data,
  });
};

export const getSubjectByEName = (data) => {
  return service({
    url: "/Subject/getSubjectByEName",
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
export const updatePost = (data) => {
  return service({
    url: "/Post/updatePost",
    method: "post",
    data,
  });
};

export const deletePost = (data) => {
  return service({
    url: "/Post/deletePost",
    method: "post",
    data,
  });
};
