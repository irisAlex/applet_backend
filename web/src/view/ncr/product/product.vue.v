<template>
    <div>
        <div class="gva-search-box">
            <a-form ref="searchForm" class="ant-advanced-search-form" :model="searchInfo" @finish="onFinish">
                <a-row :gutter="24">
                    <a-col :span="8">
                        <a-form-item label="产品编号">
                            <a-input v-model:value="searchInfo.serialnumber" placeholder="产品编号"></a-input>
                        </a-form-item>
                    </a-col>
                    <a-col :span="8">
                        <a-form-item label="产品名称">
                            <a-input v-model:value="searchInfo.name" placeholder="产品名称"></a-input>
                        </a-form-item>
                    </a-col>
                    <a-col :span="8" style="text-align: left;margin-bottom:20px;">
                        <a-button type="primary" html-type="submit" :icon="h(SearchOutlined)">
                            查询</a-button>
                        <a-button style="margin: 0 8px" @click="onReset" :icon="h(SyncOutlined)">重置</a-button>
                    </a-col>
                </a-row>

            </a-form>
        </div>

        <div class="gva-table-box">
            <div class="gva-btn-list">
                <a-button type="primary" @click="showModal('add')" :icon="h(PlusOutlined)">新增</a-button>
            </div>
            <a-table :columns="columns" :data-source="tableData" :scroll="{ x: 1500 }" :pagination="pagination"
                @change="handleTableChange">
                <template #bodyCell="{ column, record }">
                    <template v-if="column.key === 'img'">
                        <a-image :preview="{ visible: false }" :width="50" :src="path + coverImg"
                            @click="imgGroupShow(record)" :hight="50" v-if="handleJson(record)" />
                        <div style="display: none">
                            <template>
                                <a-image-preview-group
                                    :preview="{ visible, onVisibleChange: (vis: any) => (visible = vis) }">
                                    <a-image v-for="item in imgArr" :src="path + item.url" />
                                </a-image-preview-group>
                            </template>
                        </div>
                    </template>
                    <template v-if="column.key === 'operation'">
                        <a-button :icon="h(PlusOutlined)" type="link" @click="showSettingDrawer(record)">配置</a-button>
                        <a-button :icon="h(EditOutlined)" type="link" @click="editApiFunc(record)">修改</a-button>
                        <a-button :icon="h(DeleteOutlined)" type="link" @click="deleteApiFunc(record)">删除</a-button>
                    </template>
                </template>
            </a-table>
        </div>
        <div>
            <a-modal ref="modalRef" v-model:open="open" :wrap-style="{ overflow: 'hidden' }" @ok="handleOk"
                cancelText="取消" okText="确定" @cancel="Cancel">
                <a-form ref="formRef" :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
                    <a-form-item ref="name" label="产品编号">
                        <a-input v-model:value="formState.serialnumber" />
                    </a-form-item>
                    <a-form-item ref="name" label="产品名称">
                        <a-input v-model:value="formState.name" />
                    </a-form-item>
                    <a-form-item ref="product" label="图纸号">
                        <a-input v-model:value="formState.graph" />
                    </a-form-item>
                    <a-form-item ref="contacts" label="总数">
                        <a-input v-model:value.number="formState.total" />
                    </a-form-item>
                    <a-form-item ref="phone" label="产品描述">
                        <a-textarea v-model:value="formState.desc" placeholder="产品描述" :rows="4" />
                    </a-form-item>
                    <a-form-item ref="email" label="产品图片">
                        <a-upload action="/api/fileUploadAndDownload/upload" :multiple="true" :file-list="fileList"
                            @change="handleChange" @remove="handleRemove">
                            <a-button>
                                <upload-outlined></upload-outlined>
                                上传图片
                            </a-button>
                        </a-upload>
                    </a-form-item>
                </a-form>
                <template #title>
                    <div ref="modalTitleRef" style="width: 100%; cursor: move">{{ title }}</div>
                </template>
                <template #modalRender="{ originVNode }">
                    <div :style="transformStyle">
                        <component :is="originVNode" />
                    </div>
                </template>
            </a-modal>
            <!-- issue -->
            <a-drawer title="配置人员权限" :width="900" :open="openDrawer" :body-style="{ paddingBottom: '80px' }"
                :footer-style="{ textAlign: 'right' }" @close="onClose">
                <a-timeline>
                    <a-timeline-item v-for="(active, index) in configure" :key="index">
                        <span class="el-timeline-item__timestamp">{{ active.title }}</span>
                        <div class="el-card">
                            <div class="el-card__body">
                                <a-form :model="formState" layout="vertical" ref="formRef">
                                    <a-row :gutter="18">
                                        <a-col :span="6">
                                            <a-form-item label="部门" name="owner">
                                                <a-select v-model:value="active.department" placeholder="请选择部门"
                                                    @change="departmentUinfo(active.department)">
                                                    <a-select-option :value="item.authorityId"
                                                        v-for="item in departmentList" :key="item.authorityName">{{
                                                            item.authorityName }}</a-select-option>
                                                </a-select>
                                            </a-form-item>
                                        </a-col>
                                        <a-col :span="6">
                                            <a-form-item label="牵头人" name="type">
                                                <a-select v-model:value="active.principal" placeholder="请选择人员">
                                                    <a-select-option :value="item.userName"
                                                        v-for="item in AuthorityUsers" :key="item.ID">{{ item.userName
                                                        }}</a-select-option>
                                                </a-select>
                                            </a-form-item>
                                        </a-col>
                                        <a-col :span="6">
                                            <a-form-item label="限定时间" name="type">
                                                <a-input v-model:value="active.limit_date" />
                                            </a-form-item>
                                        </a-col>
                                    </a-row>
                                </a-form>
                            </div>
                        </div>
                    </a-timeline-item>
                </a-timeline>
                <template #extra>
                    <a-space>
                        <a-button @click="onClose">取消</a-button>
                        <a-button type="primary" @click="handleOk">确认</a-button>
                    </a-space>
                </template>
            </a-drawer>
        </div>
    </div>
</template>
<script lang='ts' setup>
declare module 'axios' {
    interface AxiosInstance {
        (config: AxiosRequestConfig): Promise<any>
    }
}
import { strToJson } from '@/utils/format'
import { ref, computed, watch, watchEffect, h, createVNode } from 'vue';
import { useDraggable } from '@vueuse/core';
import {
    SearchOutlined,
    SyncOutlined,
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    ExclamationCircleOutlined,
    UploadOutlined
} from '@ant-design/icons-vue';
import {
    createProductApi,
    getProductList,
    getProductById,
    deleteProduct,
    updateProduct,
    getAuthorityList,
    getUserAuthorityList
} from '@/api/product.js'
import { message, Modal, TableColumnsType, UploadChangeParam } from 'ant-design-vue';
const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'
const searchInfo = ref({
    serialnumber: "",
    name: "",
    total: 0,
    img: "",
    graph: "",
    // desc: '',
    configuration: ''
})
const labelCol = { span: 5 };
const wrapperCol = { span: 13 };
const columns: TableColumnsType = [
    { title: 'ID', width: 100, dataIndex: 'ID', key: 'ID' },
    { title: '产品编号', width: 100, dataIndex: 'serialnumber', key: 'serialnumber' },
    { title: '产品名称', dataIndex: 'name', key: 'name', width: 150 },
    { title: '产品描述', dataIndex: 'desc', key: 'desc', width: 150 },
    { title: '图纸号', dataIndex: 'graph', key: 'graph', width: 150 },
    { title: '产品图片', dataIndex: 'img', key: 'img', width: 150 },
    {
        title: '操作',
        key: 'operation',
        fixed: 'right',
        width: 150,
    },
];
const visible = ref(false);
const searchForm = ref()
const title = ref('添加供应商')
const open = ref(false);
const modalTitleRef = ref(null);
const type = ref()
const showModal = (e: any) => {
    switch (e) {
        case 'add':
            title.value = '添加供应商'
            break
        case 'edit':
            title.value = '修改供应商'
            break
        default:
            break
    }
    type.value = e
    open.value = true;
};
const formRef = ref(null);
const formState = ref({
    serialnumber: "",
    name: "",
    total: 0,
    img: "",
    graph: "",
    desc: '',
    configuration: ''
})
const editApiFunc = async (row: any) => {
    const res = await getProductById({ id: row.ID })
    formState.value = res.data.product
    if (formState.value.img !== '') {
        fileList.value = JSON.parse(formState.value.img)
    }
    showModal('edit')
}
const { x, y, isDragging } = useDraggable(modalTitleRef);

const Cancel = () => {
    formRef.value.resetFields()
    formState.value = {
        serialnumber: "",
        name: "",
        total: 0,
        img: "",
        graph: "",
        desc: '',
        configuration: ''
    }
    open.value = false;
}
const startX = ref(0);
const startY = ref(0);
const startedDrag = ref(false);
const transformX = ref(0);
const transformY = ref(0);
const preTransformX = ref(0);
const preTransformY = ref(0);
const dragRect = ref({
    left: 0,
    right: 0,
    top: 0,
    bottom: 0,
});
watch([x, y], () => {
    if (!startedDrag.value) {
        startX.value = x.value;
        startY.value = y.value;
        const bodyRect = document.body.getBoundingClientRect();
        const titleRect = modalTitleRef.value.getBoundingClientRect();
        dragRect.value.right = bodyRect.width - titleRect.width;
        dragRect.value.bottom = bodyRect.height - titleRect.height;
        preTransformX.value = transformX.value;
        preTransformY.value = transformY.value;
    }
    startedDrag.value = true;
});
watch(isDragging, () => {
    if (!isDragging) {
        startedDrag.value = false;
    }
});
watchEffect(() => {
    if (startedDrag.value) {
        transformX.value =
            preTransformX.value +
            Math.min(Math.max(dragRect.value.left, x.value), dragRect.value.right) -
            startX.value;
        transformY.value =
            preTransformY.value +
            Math.min(Math.max(dragRect.value.top, y.value), dragRect.value.bottom) -
            startY.value;
    }
});
const transformStyle = computed(() => {
    return {
        transform: `translate(${transformX.value}px, ${transformY.value}px)`,
    };
});

const onFinish = (values: any) => {
    page.value = 1
    pageSize.value = 10
    getTableData()
};

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 查询
const getTableData = async () => {
    const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

const handleOk = async () => {
    switch (type.value) {
        case 'add':
            {
                const res = await createProductApi(formState.value)
                if (res.code === 0) {
                    message.success('添加成功');
                }
                fileList.value = []
                getTableData()
                Cancel()
            }
            break
        case 'edit':
            {
                formState.value.configuration = JSON.stringify(configure.value)
                const res = await updateProduct(formState.value)
                if (res.code === 0) {
                    message.success('数据更新成功');
                }
                getTableData()
                Cancel()
            }
            break
        default:
            {
                message.error('未知操作');
            }
            break
    }
}

const deleteApiFunc = async (row: any) => {
    Modal.confirm({
        title: '提示',
        icon: createVNode(ExclamationCircleOutlined),
        content: '此操作将永久删除数据, 是否继续?',
        okText: '确定',
        okType: 'danger',
        cancelText: '取消',
        onOk() {
            const deleteApi = async () => {
                const res = await deleteProduct(row)
                if (res.code === 0) {
                    message.success('数据删除成功');
                    if (tableData.value.length === 1 && page.value > 1) {
                        page.value--
                    }
                    getTableData()
                }
            }
            deleteApi()
        },
        onCancel() {
            console.log('Cancel');
        },
    });
}

const onReset = () => {
    searchForm.value.resetFields()
    searchInfo.value = {
        serialnumber: "",
        name: "",
        total: 0,
        img: "",
        graph: "",
        //desc: '',
        configuration: ''
    }
}

const current = ref(1);
const pagination = computed(() => ({
    total: total.value,
    current: current.value,
    pageSize: pageSize.value,
    showSizeChanger: true,
    pageSizeOptions: ["10", "20", "50", "100"],
    showTotal: (total: any) => `共${total}条`,
}));

const handleTableChange = (pag: any) => {
    page.value = pag.current;
    current.value = pag.current;
    pageSize.value = pag.pageSize
    getTableData()
};


const fileList = ref([])
const imgList = ref([])
const handleRemove = (file: any) => {
    const index = fileList.value.indexOf(file);
    if (index !== -1) {
        fileList.value.splice(index, 1);
    }
    formState.value.img = JSON.stringify(fileList.value)
}
const handleChange = (info: UploadChangeParam) => {
    let resFileList = [...info.fileList];

    // 1. Limit the number of uploaded files
    //    Only to show two recent uploaded files, and old ones will be replaced by the new
    //resFileList = resFileList.slice(-10);
    // 2. read from response and show file link
    resFileList = resFileList.map(file => {
        if (file.response) {
            // Component will show file.url as link
            file.url = file.response.url;
            if (file.response.code === 0) {
                message.success('图片上传成功');
                imgList.value.push({ name: file.response.data.file.name, url: file.response.data.file.url })
                formState.value.img = JSON.stringify(imgList.value)
            } else {
                message.error('图片上传失败');
            }

        }
        return file;
    });

    fileList.value = resFileList;
};
var coverImg = ''
var imgArr = []
const handleJson = (info: any) => {
    console.log(info)
    if (info.img !== '') {
        coverImg = JSON.parse(info.img)[0].url
    }
    return true
}


const imgGroupShow = (info: any) => {
    imgArr = JSON.parse(info.img)
    visible.value = true
}

const openDrawer = ref<boolean>(false);
const onClose = () => {
    configure.value = [{
        title: "1.技术评审",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 1
    },
    {
        title: "2.工艺评审",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 2
    }]
    openDrawer.value = false;
};

const configure = ref([
    {
        title: "1.技术评审",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 1
    },
    {
        title: "2.工艺评审",
        principal: "",
        department: "",
        limit_date: 0,
        issue_number: 2
    }
])

const departmentList = ref([])
// 部门列表
const department = async () => {
    const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        departmentList.value = table.data.list
    }
}

department()

const AuthorityUsers = ref([])
//获取部门用户
const departmentUinfo = async (Authority: any) => {
    const table = await getUserAuthorityList({ page: page.value, pageSize: pageSize.value, authority_id: Authority, ...searchInfo.value })
    if (table.code === 0) {
        AuthorityUsers.value = table.data.list
    }
}

const ID = ref(1)
const showSettingDrawer = async (row: any) => {
    const res = await getProductById({ id: row.ID })
    ID.value = row.ID
    if (res.code === 0) {
        //form.value = res.data.product
        configure.value = strToJson(formState.value.configuration, configure.value)
    }
    openDrawer.value = true
    type.value = 'edit'
}

</script>

<style scoped>
#components-form-demo-advanced-search .ant-form {
    max-width: none;
}

#components-form-demo-advanced-search .search-result-list {
    margin-top: 16px;
    border: 1px dashed #e9e9e9;
    border-radius: 2px;
    background-color: #fafafa;
    min-height: 200px;
    text-align: center;
    padding-top: 80px;
}

[data-theme='dark'] .ant-advanced-search-form {
    background: rgba(255, 255, 255, 0.04);
    border: 1px solid #434343;
    padding: 24px;
    border-radius: 2px;
}

[data-theme='dark'] #components-form-demo-advanced-search .search-result-list {
    border: 1px dashed #434343;
    background: rgba(255, 255, 255, 0.04);
}

.el-card {
    border-radius: 4px;
    border: 1px solid #ebeef5;
    background-color: #fff;
    overflow: hidden;
    color: #303133;
    transition: .3s;
}

.el-card.is-always-shadow,
.el-card.is-hover-shadow:focus,
.el-card.is-hover-shadow:hover {
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .1)
}

.el-card__header {
    padding: 18px 20px;
    border-bottom: 1px solid #ebeef5;
    box-sizing: border-box
}

.el-card__body {
    padding: 20px
}

.el-timeline-item__timestamp {
    color: #909399;
    line-height: 1;
    font-size: 13px;
    margin-top: 10px;
}
</style>