<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="searchForm" :inline="true" :model="searchInfo">
                <el-form-item label="部门" style="width:200px;">
                    <el-select v-model="searchInfo.department" placeholder="北京安新">
                        <el-option v-for="item in departmentList" :key="item.authorityId" :label="item.authorityName"
                            :value="item.authorityName">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类型" style="width:200px;">
                    <el-select v-model="searchInfo.mold" placeholder="请选择">
                        <el-option v-for="item in moldList" :key="item.value" :label="item.label" :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="类别" style="width:200px;">
                    <el-select v-model="searchInfo.category" placeholder="请选择">
                        <el-option v-for="item in genreList1" :key="item.name" :label="item.genre" :value="item.genre">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="受检物名称" style="width:200px;">
                    <el-input v-model="searchInfo.checkout_name" placeholder="受检物名称" />
                </el-form-item>
                <el-form-item label="受检物号" style="width:200px;">
                    <el-input v-model="searchInfo.checkout_number" placeholder="受检物号" />
                </el-form-item>
                <el-form-item label="处理方式" style="width:200px;">
                    <el-select v-model="searchInfo.process_mode" placeholder="请选择">
                        <el-option v-for="item in methodOptions" :key="item.value" :label="item.label"
                            :value="item.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="计划时间">
                    <el-date-picker v-model="searchInfo.rework_plan_date" type="date" placeholder="计划时间"
                        value-format="YYYY-MM-DDT15:04:05Z">
                    </el-date-picker>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
            <el-table-column align="left" label="ID" min-width="150" prop="ID" sortable="custom" />
            <el-table-column align="left" label="编号" min-width="150" prop="serialnumber" sortable="custom" />
            <el-table-column align="left" label="部门" min-width="150" prop="department" sortable="custom" />
            <el-table-column align="left" label="类型" min-width="150" prop="mold" sortable="custom" />
            <el-table-column align="left" label="类别" min-width="150" prop="category" sortable="custom" />
            <el-table-column align="left" label="项目" min-width="150" prop="project" sortable="custom" />
            <el-table-column align="left" label="受检物名称" min-width="150" prop="checkout_name" sortable="custom" />
            <el-table-column align="left" label="受检物号" min-width="150" prop="checkout_number" sortable="custom" />
            <el-table-column align="left" label="处理方式" min-width="150" prop="process_mode" sortable="custom" />
            <el-table-column align="left" label="状态" min-width="150" prop="description" sortable="custom">
                <template #default="scope">
                    <div slot="reference" class="name-wrapper">
                        <el-tag
                            :type="scope.row.status === '1' ? 'success' : scope.row.status === '-1' ? 'danger' : 'info'"
                            disable-transitions>{{
                                scope.row.status === '1' ? '同意' : scope.row.status === '-1' ? '拒绝' : '待审批'
                            }}</el-tag>
                    </div>
                </template>
            </el-table-column>
            <el-table-column align="left" label="数量" min-width="150" prop="rework_number" sortable="custom" />
            <el-table-column align="left" label="工时" min-width="150" prop="rework_man_hour" sortable="custom" />
            <el-table-column align="left" label="工料" min-width="150" prop="rework_quantities" sortable="custom" />
            <el-table-column align="left" label="工序" min-width="150" prop="rework_process" sortable="custom" />
            <el-table-column align="left" label="计划时间" min-width="150" sortable="custom">
                <template #default="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ formatDate(scope.row.rework_plan_date) }}</span>
                </template>
            </el-table-column>
            <el-table-column align="left" label="照片" min-width="200" prop="rework_attachment" sortable="custom">
                <template #default="scope">
                    <el-image
                        v-for="item in (scope.row.rework_attachment === '' ? '' : JSON.parse(scope.row.rework_attachment))"
                        style="width: 100px; height: 100px;display: block;margin: 5px;box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);"
                        :src="path + item.url" :preview-src-list="[path + item.url]" :preview-teleported="true"
                        fit="fill">
                    </el-image>
                </template>
            </el-table-column>
            <el-table-column align="left" fixed="right" label="操作" width="300">
                <template #default="scope">
                    <el-button icon="check" type="primary" link @click="setApiFunc(scope.row, '1')">同意</el-button>
                    <el-button icon="remove" type="primary" link @click="setApiFunc(scope.row, '-1')">拒绝</el-button>
                    <el-button icon="view" type="primary" link @click="editApiFunc(scope.row, 'look')">查看</el-button>
                    <el-button icon="edit" type="primary" link @click="editApiFunc(scope.row, 'modify')">修改</el-button>
                    <el-button icon="printer" type="primary" link @click="">打印</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                @size-change="handleSizeChange" />
        </div>
        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="40%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="数量" prop="rework_number" style="width:20%">
                    <span v-if="!isModfiy">{{ form.rework_number }}</span>
                    <el-input placeholder="数量" size="mini" v-model.number="form.rework_number" v-if="isModfiy" />
                </el-form-item>
                <el-form-item label="工时" prop="rework_man_hour" style="width: 20%">
                    <span v-if="!isModfiy">{{ form.rework_man_hour }}</span>
                    <el-input placeholder="工时" size="mini" v-model.number="form.rework_man_hour" v-if="isModfiy" />
                </el-form-item>
                <el-form-item label="工料" prop="rework_quantities" style="width:20%">
                    <span v-if="!isModfiy">{{ form.rework_quantities }}</span>
                    <el-input placeholder="工料" size="mini" v-model="form.rework_quantities" v-if="isModfiy" />
                </el-form-item>
                <el-form-item label="工序" prop="rework_process" style="width:20%">
                    <span v-if="!isModfiy">{{ form.rework_process }}</span>

                    <el-input placeholder="工序" size="mini" v-model="form.rework_process" v-if="isModfiy" />
                </el-form-item>
                <el-form-item label="返工计划完成时间" prop="rework_plan_date" style="width:40%">
                    <span v-if="!isModfiy">{{ formatDate(form.rework_plan_date) }}</span>

                    <el-date-picker v-model="form.rework_plan_date" type="date" placeholder="选择日期"
                        value-format="YYYY-MM-DDT15:04:05Z" v-if="isModfiy">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="返工描述" prop="rework_desc" style="width:100%">
                    <span v-if="!isModfiy">{{ form.rework_man_hour }}</span>

                    <el-input type="textarea" placeholder="请输入内容" v-model="form.rework_desc" maxlength="50"
                        show-word-limit :rows="10" v-if="isModfiy" />
                </el-form-item>
                <el-form-item label="返工证据" prop="rework_attachment" style="width:100%">
                    <el-image v-for="item in imgList"
                        style="width: 100px; height: 100px;display: block;margin: 5px;box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);"
                        :src="path + item.url" :preview-src-list="[path + item.url]" v-if="!isModfiy">
                    </el-image>
                    <el-upload action="/api/fileUploadAndDownload/upload" multiple :limit="2" :file-list="fileList1"
                        :on-success="handleSuccess1" show-file-list="false" :on-remove="handleRemove1" v-if="isModfiy"
                        :headers="{ 'x-token': userStore.token }">
                        <el-button size="small" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                    </el-upload>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog" v-if="isModfiy">取 消</el-button>
                    <el-button type="primary" @click="enterDialog" v-if="isModfiy">确 定</el-button>
                </div>
            </template>
        </el-dialog>
    </div>

</template>

<script setup>
import {
    getAuthorityList,
    getGenreList,
    getSupplierList,
    getProjectList,
    setStatus,
    getManageList,
    getManageById,
    updateManage
} from '@/api/manage.js'

import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
    name: 'Api',
})

const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
}

//返工图片处理
const handleSuccess1 = (resp) => {
    if (resp.code === 0) {
        ElMessage({
            type: 'success',
            message: '图片上传成功',
            showClose: true
        })

        fileList1.value.push({ name: resp.data.file.name, url: resp.data.file.url })
        form.value.rework_attachment = JSON.stringify(fileList1.value)
        return
    }
    ElMessage({
        type: 'error',
        message: '图片上传失败',
        showClose: true
    })

};

const handleRemove1 = (file, fileList1) => {
    // 处理删除文件的逻辑，例如从文件列表中删除文件
    const index = fileList1.indexOf(file);
    if (index !== -1) {
        fileList.splice(index, 1);
    }
    form.value.rework_attachment = JSON.stringify(fileList1.value)
};

const isModfiy = ref(false)
const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_CLI_PORT + '/'
const apis = ref([])
const form = ref({
    serialnumber: "",
    process_mode: "",
    department: "",
    mold: "",
    category: "",
    project: "",
    checkout_name: "",
    checkout_number: "",
    rework_number: 0,
    rework_man_hour: 0,
    rework_quantities: "",
    rework_process: "",
    rework_plan_date: "0001-01-01T00:00:00Z",
    rework_desc: "",
    rework_attachment: "",
    is_ncr: ''
})
const methodOptions = ref([
    {
        value: 'Just Do it',
        label: 'Just Do it',
        type: 'success'
    },
    {
        value: 'A3',
        label: 'A3',
        type: ''
    },
    {
        value: '8D',
        label: '8D',
        type: 'warning'
    }
])

const moldList = ref([
    {
        value: '内部',
        label: '内部',
        type: 'success'
    },
    {
        value: '外部',
        label: '外部',
        type: ''
    },
    {
        value: '配做',
        label: '配做',
        type: 'warning'
    }
])
const type = ref('')
const rules = ref({
    path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
    apiGroup: [
        { required: true, message: '请输入组名称', trigger: 'blur' }
    ],
    method: [
        { required: true, message: '请选择请求方式', trigger: 'blur' }
    ],
    description: [
        { required: true, message: '请输入api介绍', trigger: 'blur' }
    ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
    searchInfo.value = {}
}
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

const departmentList = ref([])
const genreList1 = ref([])
const supplierList = ref([])
const projectList = ref([])

// 部门列表
const department = async () => {
    const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        departmentList.value = table.data.list
    }
}

department()

// 类别列表
const genreList = async () => {
    const table = await getGenreList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        genreList1.value = table.data.list
    }
}

genreList()

// 类别列表
const supplier = async () => {
    const table = await getSupplierList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        supplierList.value = table.data.list
    }
}

supplier()

// 类别列表
const project = async () => {
    const table = await getProjectList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        projectList.value = table.data.list
    }
}

project()


const operation = ref("返工")
// 查询
const getTableData = async () => {
    const table = await getManageList({ page: page.value, pageSize: pageSize.value, keyword: operation.value, orderKey: 'id', desc: true, ...searchInfo.value })
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


// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        serialnumber: "",
        process_mode: "",
        department: "",
        mold: "",
        category: "",
        project: "",
        checkout_name: "",
        checkout_number: "",
        rework_number: 0,
        rework_man_hour: 0,
        rework_quantities: "",
        rework_process: "",
        rework_plan_date: "0001-01-01T00:00:00Z",
        rework_desc: "",
        rework_attachment: "",
        is_ncr: ''
    }
}
const dialogTitle = ref('查看返工')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
    switch (key) {
        case 'look':
            dialogTitle.value = '查看返工'
            isModfiy.value = false
            break
        case 'modify':
            isModfiy.value = true
            dialogTitle.value = '修改返工'
            break
        default:
            break
    }
    type.value = key
    dialogFormVisible.value = true
}
const closeDialog = () => {
    initForm()
    dialogFormVisible.value = false
}

const fileList1 = ref([])

const imgList = ref([])

const editApiFunc = async (row, operation) => {
    const res = await getManageById({ id: row.ID })
    form.value = res.data.manage
    if (form.value.rework_attachment !== '') {
        fileList1.value = JSON.parse(form.value.rework_attachment)
        imgList.value = JSON.parse(form.value.rework_attachment)
    }

    openDialog(operation)
}


const setApiFunc = async (row, status) => {
    const res = await setStatus({ id: row.ID, status: status })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '状态更新成功',
            showClose: true
        })
    }
    getTableData()
}


const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            switch (type.value) {
                case 'modify':
                    {
                        const res = await updateManage(form.value)
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

</script>

<style scoped lang="scss">
.warning {
    color: #dc143c;
}
</style>