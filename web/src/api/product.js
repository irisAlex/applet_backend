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
export const getProductList = (data) => {
  return service({
    url: "/product/getAllProductList",
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
export const createProductApi = (data) => {
  return service({
    url: "/product/createProductApi",
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
export const getProductById = (data) => {
  return service({
    url: "/product/getProductById",
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
export const updateProduct = (data) => {
  return service({
    url: "/product/updateProduct",
    method: "post",
    data,
  });
};



// @Tags Api
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Api true "删除api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
export const deleteProduct = (data) => {
  return service({
    url: "/product/deleteProduct",
    method: "post",
    data,
  });
};

// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApisByIds [delete]
export const deleteApisByIds = (data) => {
  return service({
    url: "/api/deleteApisByIds",
    method: "delete",
    data,
  });
};

// FreshCasbin
// @Tags      SysApi
// @Summary   刷新casbin缓存
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "刷新成功"
// @Router    /api/freshCasbin [get]
export const freshCasbin = () => {
  return service({
    url: "/api/freshCasbin",
    method: "get",
  });
};

export const getAuthorityList = (data) => {
    return service({
      url: "/authority/getAuthorityList",
      method: "post",
      data,
    });
  };


  //获取部门用户
export const getUserAuthorityList = (data) => {
    return service({
      url: '/user/getUserList',
      method: 'post',
      data: data
    })
  }
  