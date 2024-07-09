<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="产品编号" prop="method" style="width:200px;">
                    <el-input v-model="searchInfo.serialnumber" placeholder="产品编号" />
                </el-form-item>
                <el-form-item label="产品名称" prop="project" style="width:200px;">
                    <el-input v-model="searchInfo.name" placeholder="产品名称" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="addOpenDialog('addApi')">新增</el-button>
            </div>
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" />
                <el-table-column align="left" label="产品编号" min-width="150" prop="serialnumber" />
                <el-table-column align="left" label="产品名称" min-width="150" prop="name" />
                <el-table-column align="left" label="产品描述" min-width="150" prop="desc" />
                <el-table-column align="left" label="图纸号" min-width="150" prop="graph" />
                <el-table-column align="left" label="产品图片" min-width="150" prop="img">
                    <template #default="scope">
                        <el-image v-for="item in (scope.row.img === '' ? '' : JSON.parse(scope.row.img))"
                            style="width: 100px; height: 100px;display: block;margin: 5px;box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);"
                            :src="path + item.url" :preview-src-list="[path + item.url]" :preview-teleported="true"
                            fit="fill">
                        </el-image>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="总数" min-width="150" prop="total" />
                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <el-button type="primary" icon="setting" link @click="openDialog(scope.row, 'setting')">
                            配置
                        </el-button>
                        <el-button icon="edit" type="primary" link
                            @click="editApiFunc(scope.row, 'edit')">修改</el-button>
                        <el-button icon="delete" type="primary" link @click="deleteApiFunc(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="60%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="产品编号:" prop="serialnumber" style="width:20%">
                    <el-input placeholder="产品编号" size="mini" v-model="form.serialnumber" />
                </el-form-item>
                <el-form-item label="产品名称:" prop="name" style="width:20%">
                    <el-input placeholder="产品名称" size="mini" v-model="form.name" />
                </el-form-item>
                <el-form-item label="图纸号:" prop="graph" style="width:20%">
                    <el-input placeholder="图纸号" size="mini" v-model="form.graph" />
                </el-form-item>
                <el-form-item label="总数:" prop="total" style="width:20%">
                    <el-input placeholder="总数" size="mini" v-model.number="form.total" />
                </el-form-item>
                <el-form-item label="产品描述:" prop="desc" style="width:60%">
                    <el-input type="textarea" placeholder="请输入内容" v-model="form.desc" maxlength="50" show-word-limit
                        :rows="10" />
                </el-form-item>
                <el-form-item label="产品图片:" prop="img" style="width:100%">
                    <el-upload action="/api/fileUploadAndDownload/upload" multiple :limit="2" :file-list="fileList"
                        :on-success="handleSuccess" show-file-list="false" :on-remove="handleRemove">
                        <el-button size="small" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                    </el-upload>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-dialog>
        <!-- issue 确认 -->
        <el-drawer v-model="drawer" :title="dialogTitle" :direction="direction" :before-close="handleClose" :size="ize">
            <div class="block">
                <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                    <el-timeline>
                        <el-timeline-item v-for="(active, index) in configure" :key="index" :timestamp="active.title"
                            placement="top">
                            <el-card>
                                <el-form :inline="true">
                                    <el-form-item label="部门:" prop="notification_department" style="width:20%">
                                        <el-select v-model="active.department" placeholder="部门" style="width:100%"
                                            @change="departmentUinfo(active.department)">
                                            <el-option v-for="item in departmentList" :key="item.authorityName"
                                                :label="`${item.authorityName}`" :value="item.authorityId" />
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="负责人:" prop="notification_member" style="width:20%">
                                        <el-select v-model="active.principal" placeholder="选择人员" style="width:100%">
                                            <el-option v-for="item in AuthorityUsers" :key="item.ID"
                                                :label="`${item.userName}`" :value="item.userName" />
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="限定时间/天:" prop="limit_date" style="width:30%">
                                        <el-input placeholder="限定时间" size="mini" v-model="active.limit_date" />
                                    </el-form-item>
                                </el-form>
                            </el-card>
                        </el-timeline-item>
                    </el-timeline>
                </el-form>
            </div>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="handleClose">取 消</el-button>
                    <el-button type="primary" @click="enterDialog">确 定</el-button>
                </div>
            </template>
        </el-drawer>
    </div>
</template>
<script setup>
import {
    createProductApi,
    getProductList,
    getProductById,
    deleteProduct,
    updateProduct,
    getAuthorityList,
    getUserAuthorityList
} from '@/api/product.js'
import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref } from 'vue'
import { strToJson } from '@/utils/format'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage, ElMessageBox } from 'element-plus'
defineOptions({
    name: 'Setting',
})

const cir = ref(0)

const dialogFormVisible = ref(false)
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const handlePictureCardPreview = (file) => {
    dialogImageUrl.value = file.url;
    dialogVisible.value = true;
    dialogFormVisible.value = true
}

const drawer = ref(false)
const direction = ref('rtl')
const ize = ref('70%')
const isShowReport = ref(false)
const dialogTitle = ref('人员权限配置')
const openDialog = (row, key) => {
    switch (key) {
        case 'setting':
            dialogTitle.value = '人员权限配置'
            isShowReport.value = false
            break
        default:
            break
    }
    showSettingDrawer(key, row)
}
const type = ref('')
const addOpenDialog = (key) => {
    switch (key) {
        case 'addApi':
            dialogTitle.value = '创建产品'
            break
        case 'edit':
            dialogTitle.value = '编辑产品'
            break
        default:
            break
    }
    type.value = key
    dialogFormVisible.value = true
}

const editApiFunc = async (row, operation) => {
    const res = await getProductById({ id: row.ID })
    form.value = res.data.product
    if (form.value.img !== '') {
        fileList.value = JSON.parse(form.value.img)
    }

    addOpenDialog(operation)
}



const handleClose = () => {
    initForm()
    drawer.value = false
}
const ID = ref(1)
const showSettingDrawer = async (key, row) => {
    const res = await getProductById({ id: row.ID })
    ID.value = row.ID
    if (res.code === 0) {
        form.value = res.data.product
        configure.value = strToJson(form.value.configuration, configure.value)
    }
    drawer.value = true
    type.value = 'edit'
}
const setp = ref(0)
const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'


const userStore = useUserStore()

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

const apis = ref([])
const departmentList = ref([])
const fileList = ref([])
const form = ref({
    serialnumber: "",
    name: "",
    total: "",
    img: "",
    graph: "",
    desc: '',
    configuration: ''
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const onReset = () => {
    searchInfo.value = {}
}

const handleSuccess = (resp) => {
    if (resp.code === 0) {
        ElMessage({
            type: 'success',
            message: '图片上传成功',
            showClose: true
        })

        fileList.value.push({ name: resp.data.file.name, url: resp.data.file.url })
        form.value.img = JSON.stringify(fileList.value)
        return
    }
    ElMessage({
        type: 'error',
        message: '图片上传失败',
        showClose: true
    })

};

const handleRemove = (file, fileList) => {
    // 处理删除文件的逻辑，例如从文件列表中删除文件
    const index = fileList.indexOf(file);
    if (index !== -1) {
        fileList.splice(index, 1);
    }
    form.value.img = JSON.stringify(fileList.value)
};


// 搜索
const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
}

// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
    if (prop) {
        if (prop === 'ID') {
            prop = 'id'
        }
        searchInfo.value.orderKey = toSQLLine(prop)
        searchInfo.value.desc = order === 'descending'
    }
    getTableData()
}

// 部门列表
const department = async () => {
    const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        departmentList.value = table.data.list
    }
}

department()

// 查询
const getTableData = async () => {
    searchInfo.value.process_mode = "A3"
    const table = await getProductList({ page: page.value, pageSize: pageSize.value, orderKey: 'id', desc: true, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

// 批量操作
const handleSelectionChange = (val) => {
    apis.value = val
}


const AuthorityUsers = ref([])
const notification_department = ref('')
const notification_member = ref('')
//获取部门用户
const departmentUinfo = async (Authority) => {
    const table = await getUserAuthorityList({ page: page.value, pageSize: pageSize.value, authority_id: Authority, ...searchInfo.value })
    if (table.code === 0) {
        console.log(table.data.list)

        AuthorityUsers.value = table.data.list
    }
}


// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        serialnumber: "",
        name: "",
        total: "",
        img: "",
        graph: "",
        desc: '',
        configuration: ''
    }
    configure.value = ([
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
}

const closeDialog = () => {
    initForm()
    drawer.value = false
    dialogFormVisible.value = false
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            switch (type.value) {
                case 'addApi':
                    {
                        const res = await createProductApi(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '添加成功',
                                showClose: true
                            })
                        }
                        getTableData()
                        closeDialog()
                    }

                    break
                case 'edit':
                    {
                        form.value.configuration = JSON.stringify(configure.value)
                        const res = await updateProduct(form.value)
                        if (res.code === 0) {
                            ElMessage({
                                type: 'success',
                                message: '编辑成功',
                                showClose: true
                            })
                        }
                        getTableData()
                        closeDialog()
                    }
                    break
                default:
                    // eslint-disable-next-line no-lone-blocks
                    {
                        ElMessage({
                            type: 'error',
                            message: '未知操作',
                            showClose: true
                        })
                    }
                    break
            }
        }
    })
}


const deleteApiFunc = async (row) => {
    ElMessageBox.confirm('此操作将永久删除数据, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await deleteProduct(row)
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '删除成功!'
                })
                if (tableData.value.length === 1 && page.value > 1) {
                    page.value--
                }
                getTableData()
            }
        })
}


</script>
<style lang="scss" scoped>
.drawer-container {
    transition: all 0.2s;

    &:hover {
        right: 0
    }

    position: fixed;
    right: -20px;
    bottom: 15%;
    height: 40px;
    width: 800px;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 999;
    color: #fff;
    border-radius: 4px 0 0 4px;
    cursor: pointer;
    -webkit-box-shadow: inset 0 0 6px rgba(0, 0, 0, 10%);
}

.setting_body {
    padding: 20px;

    .setting_card {
        margin-bottom: 20px;
    }

    .setting_content {
        margin-top: 20px;
        display: flex;
        flex-direction: column;

        >.theme-box {
            display: flex;
        }

        >.color-box {
            div {
                display: flex;
                flex-direction: column;
            }
        }

        .item {
            display: flex;
            align-items: center;
            justify-content: center;
            flex-direction: column;
            margin-right: 20px;

            .item-top {
                position: relative;
            }

            .check {
                position: absolute;
                font-size: 20px;
                color: #00afff;
                right: 10px;
                bottom: 10px;
            }

            p {
                text-align: center;
                font-size: 12px;
            }
        }
    }
}

.el-card__body {
    height: 800px;
}

.disabled {
    pointer-events: none;
    opacity: 0.5;
}
</style>