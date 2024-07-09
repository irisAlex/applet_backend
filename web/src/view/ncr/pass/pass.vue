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
                    <el-select v-model="value" placeholder="请选择">
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
                <el-form-item label="项目名称" style="width:200px;">
                    <el-input v-model="searchInfo.project" placeholder="项目名称" />
                </el-form-item>
                <el-form-item label="供应商名称" style="width:200px;">
                    <el-input v-model="searchInfo.supplier" placeholder="项目名称" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" sortable="custom" />
                <el-table-column align="left" label="编号" min-width="150" prop="serialnumber" sortable="custom" />
                <el-table-column align="left" label="部门" min-width="150" prop="department" sortable="custom" />
                <el-table-column align="left" label="类型" min-width="150" prop="mold" sortable="custom" />
                <el-table-column align="left" label="类别" min-width="150" prop="category" sortable="custom" />
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
                <el-table-column align="left" label="放行时间" min-width="150" prop="deferpass_datered_date"
                    sortable="custom">
                    <template #default="scope">
                        <i class="el-icon-time"></i>
                        <span style="margin-left: 10px">{{ formatDate(scope.row.deferpass_datered_date) }}</span>
                    </template>
                </el-table-column>

                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <el-button icon="view" type="primary" link @click="editApiFunc(scope.row)">查看</el-button>
                        <el-button icon="check" type="primary" link @click="setApiFunc(scope.row, '1')">同意</el-button>
                        <el-button icon="remove" type="primary" link @click="setApiFunc(scope.row, '-1')">拒绝</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="40%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <el-form-item label="编号" prop="method" style="width:30%">
                    <span>
                        {{ form.serialnumber }}
                    </span>
                </el-form-item>
                <el-form-item label="受检物名称" prop="checkout_name" style="width:30%">
                    <span>{{ form.checkout_name }}</span>
                </el-form-item>
                <el-form-item label="受检物号" prop="checkout_number" style="width:30%">
                    <span>
                        {{ form.checkout_number }}
                    </span>
                </el-form-item>
                <el-form-item label="类别" prop="category" style="width:30%">
                    <span>
                        {{ form.category }}
                    </span>
                </el-form-item>
                <el-form-item label="部门" prop="department" style="width:30%">
                    <span>
                        {{ form.department }}
                    </span>
                </el-form-item>
                <el-form-item label="处理方式" prop="process_mode" style="width:30%">
                    <span>
                        {{ form.process_mode }}
                    </span>
                </el-form-item>
                <el-form-item label="类型" prop="category" style="width:30%">
                    <span>
                        {{ form.mold }}
                    </span>
                </el-form-item>
                <el-form-item label="状态" prop="status" style="width:30%">
                    <template #default="scope">
                        <div slot="reference" class="name-wrapper">
                            <el-tag :type="form.status === '1' ? 'success' : form.status === '-1' ? 'danger' : 'info'"
                                disable-transitions>{{
                                    form.status === '1' ? '同意' : form.status === '-1' ? '拒绝' : '待处理'
                                }}</el-tag>
                        </div>
                    </template>
                </el-form-item>
                <el-form-item label="放行时间" prop="pass_date" style="width:30%">
                    <span>
                        {{ formatDate(form.pass_date) }}
                    </span>
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="closeDialog" v-if="isFormDisabled">取 消</el-button>
                    <el-button type="primary" @click="enterDialog" v-if="isFormDisabled">确 定</el-button>
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
    getManageList,
    getManageById,
    setStatus
} from '@/api/manage.js'
import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
    name: 'Api',
})

const methodFilter = (value) => {
    const target = methodOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
}

const apis = ref([])
const departmentList = ref([])
const genreList1 = ref([])
const supplierList = ref([])
const projectList = ref([])
const form = ref({
    serialnumber: "",
    department: "",
    mold: "",
    category: "",
    project: "",
    checkout_name: "",
    checkout_number: "",
    graph_number: "",
    version_number: "",
    purchase_order: "",
    production_order: "",
    delivery_order: "",
    packages_number: "",
    reject_packages_number: "",
    sample_checkout_number: "",
    reject_sample_checkout_number: "",
    supplier: "",
    checkout_date: "",
    describe: "",
    photograph: "",
    process_mode: "",
    duty_department: "",
    cause_desc: "",
    fill_from_date: "",
    disposal_concept: "",
    rework_number: 0,
    rework_man_hour: 0,
    rework_quantities: "",
    rework_process: "",
    rework_plan_date: "0001-01-01T00:00:00Z",
    rework_desc: "",
    rework_attachment: "",
    repair_plan_date: "0001-01-01T00:00:00Z",
    repair_desc: "",
    repair_attachment: "",
    parts_desc: "",
    series: "",
    operation_type: ""
})
const methodOptions = ref([
    {
        value: 'POST',
        label: '创建',
        type: 'success'
    },
    {
        value: 'GET',
        label: '查看',
        type: ''
    },
    {
        value: 'PUT',
        label: '更新',
        type: 'warning'
    },
    {
        value: 'DELETE',
        label: '删除',
        type: 'danger'
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

const operation = ref("让步接收")
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

const deleteVisible = ref(false)
const onDelete = async () => {
    const ids = apis.value.map(item => item.ID)
    const res = await deleteApisByIds({ ids })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: res.msg
        })
        if (tableData.value.length === ids.length && page.value > 1) {
            page.value--
        }
        deleteVisible.value = false
        getTableData()
    }
}
const freshVisible = ref(false)
const onFresh = async () => {
    const res = await freshCasbin()
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: res.msg
        })
    }
    freshVisible.value = false
}

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
        serialnumber: "",
        department: "",
        mold: "",
        category: "",
        project: "",
        checkout_name: "",
        checkout_number: "",
        graph_number: "",
        version_number: "",
        purchase_order: "",
        production_order: "",
        delivery_order: "",
        packages_number: "",
        reject_packages_number: "",
        sample_checkout_number: "",
        reject_sample_checkout_number: "",
        supplier: "",
        checkout_date: "",
        describe: "",
        photograph: "",
        process_mode: "",
        duty_department: "",
        cause_desc: "",
        fill_from_date: "",
        disposal_concept: "",
        rework_number: 0,
        rework_man_hour: 0,
        rework_quantities: "",
        rework_process: "",
        rework_plan_date: "0001-01-01T00:00:00Z",
        rework_desc: "",
        rework_attachment: "",
        repair_plan_date: "0001-01-01T00:00:00Z",
        repair_desc: "",
        repair_attachment: "",
        parts_desc: "",
        series: "",
        operation_type: ""
    }
}

const dialogTitle = ref('添加不合格品')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
    switch (key) {
        case 'look':
            dialogTitle.value = '查看放行'
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

const editApiFunc = async (row) => {
    const res = await getManageById({ id: row.ID })
    form.value = res.data.manage
    openDialog('look')
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (valid) {
            switch (type.value) {
                case 'addApi':
                    {
                        const res = await createApi(form.value)
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
                        const res = await updateApi(form.value)
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
    ElMessageBox.confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    })
        .then(async () => {
            const res = await deleteApi(row)
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

<style scoped lang="scss">
.warning {
    color: #dc143c;
}
</style>