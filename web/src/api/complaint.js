import service from "@/utils/request";

// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
export const createComplaint = (data) => {
  return service({
    url: "/complaint/createComplaint",
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
export const getComplaintById = (data) => {
  return service({
    url: "/complaint/getComplaintById",
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
export const updateComplaint = (data) => {
  return service({
    url: "/complaint/updateComplaint",
    method: "post",
    data,
  });
};

export const deleteComplaint = (data) => {
  return service({
    url: "/complaint/deleteComplaint",
    method: "post",
    data,
  });
};

//upload
export const uploadFile = (data) => {
  return service({
    url: "/fileUploadAndDownload/upload",
    method: "post",
    transformRequest: [
      function (data, headers) {
        // 去除post请求默认的Content-Type
        // delete headers["Content-Type"];
        return data;
      },
    ],
    data: data,
    timeout: 300000,
  });
};

//set status 

//upload
export const setStatus = (data) => {
  return service({
    url: "/complaint/updateSetStatus",
    method: "post",
    data,
  });
};

export const setPassDate = (data) => {
  return service({
    url: "/complaint/updatePassDate",
    method: "post",
    data,
  });
};
export const updateParts = (data) => {
    return service({
    url: "/complaint/updateParts",
    method: "post",
    data,
  });
};


//manage
export const getComplaintlist = (data) => {
  return service({
    url: "/complaint/getAllComplaintList",
    method: "post",
    data,
  });
};