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
            <!-- <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog('addApi')">新增</el-button>
            </div> -->
            <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
                <el-table-column align="left" label="ID" min-width="150" prop="ID" sortable="custom" />
                <el-table-column align="left" label="受检物名称" min-width="150" prop="checkout_name" sortable="custom" />
                <el-table-column align="left" label="受检物号" min-width="150" prop="checkout_number" sortable="custom" />
                <el-table-column align="left" label="类别" min-width="150" prop="category" sortable="custom" />
                <el-table-column align="left" label="部门" min-width="150" prop="department" sortable="custom" />
                <el-table-column align="left" label="项目" min-width="150" prop="project" sortable="custom" />
                <el-table-column align="left" label="状态" min-width="150" prop="status" sortable="custom">
                    <template #default="scope">
                        <div slot="reference" class="name-wrapper">
                            <el-tag
                                :type="scope.row.status === '1' ? 'success' : scope.row.status === '-1' ? 'danger' : 'info'"
                                disable-transitions>{{
                                    scope.row.status === '1' ? '已完成' : scope.row.status === '-1' ? '延期' : '待处理'
                                }}</el-tag>
                        </div>
                    </template>
                </el-table-column>
                <!-- <el-table-column align="left" label="配做方案" min-width="150" prop="parts_desc" sortable="custom"
                    v-html="parts_desc" /> -->
                <el-table-column align="left" label="配做订单" min-width="600" prop="series">
                    <template #default="scope">
                        <el-table :data="JSON.parse(scope.row.series)" border style="width: 100%;" :stripe="true">
                            <el-table-column label="产品系列号" align="center" min-width="100" prop="product_serialnumber">
                                <template #default="scope">
                                    <span>{{ scope.row.product_serialnumber }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column label="产品名称" align="center" min-width="100" prop="product_name">
                                <template #default="scope">
                                    <span>{{ scope.row.product_name }}</span>
                                </template>
                            </el-table-column>
                            <el-table-column label="物料系列号" min-width="100" prop="w_serialnumber">
                                <template #default="scope">
                                    <span>
                                        {{ scope.row.w_serialnumber }}
                                    </span>
                                </template>
                            </el-table-column>
                            <el-table-column label="物料名称" align="center" min-width="100" prop="w_name">
                                <template #default="scope">
                                    <span>
                                        {{ scope.row.w_name }}
                                    </span>
                                </template>
                            </el-table-column>
                        </el-table>
                    </template>
                </el-table-column>
                <el-table-column align="left" label="延期时间" min-width="150" prop="deferred_date" sortable="custom">
                    <template #default="scope">
                        <el-tag :type="info" disable-transitions>
                            {{ scope.row.deferred_date !== '0001-01-01T00:00:00Z' ?
                                formatDate(scope.row.deferred_date) : '按期完成任务' }}
                        </el-tag>
                    </template>
                </el-table-column>

                <el-table-column align="left" fixed="right" label="操作" width="300">
                    <template #default="scope">
                        <el-button icon="view" type="primary" link
                            @click="editApiFunc(scope.row, 'look')">查看</el-button>
                        <el-button icon="set-up" type="primary" link
                            @click="editApiFunc(scope.row, 'edit')">处理</el-button>
                        <el-button icon="set-up" type="primary" link
                            @click="editApiFunc(scope.row, 'modify')">修改</el-button>

                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                    @size-change="handleSizeChange" />
            </div>

        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle" width="50%">
            <el-form ref="apiForm" :model="form" :rules="rules" :inline="true">
                <div v-if="isLook">
                    <el-form-item label="配做方案" prop="parts_desc" style="width:100%">
                        <div v-if="Look">
                            <QuillEditor :options="editorOptions" content-type="html" ref="quillEditor" theme="snow"
                                v-model:content="form.parts_desc" :value="form.parts_desc" readOnly />
                        </div>
                        <div v-if="!Look">
                            <QuillEditor :options="editorOptions" content-type="html" ref="quillEditor" theme="snow"
                                v-model:content="form.parts_desc" :value="form.parts_desc" />
                        </div>
                    </el-form-item>
                    <el-form-item prop="series" label="配做订单">
                        <el-table :data="orderList" border style="width: 100%;" :stripe="true">
                            <el-table-column label="产品系列号" align="center" min-width="150" prop="product_serialnumber">
                                <template #default="scope">
                                    <span v-if="Look">{{ scope.row.product_serialnumber }}</span>
                                    <el-input v-model="scope.row.product_serialnumber" v-if="!Look"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="产品名称" align="center" min-width="150" prop="product_name">
                                <template #default="scope">
                                    <span v-if="Look">{{ scope.row.product_serialnumber }}</span>
                                    <el-input v-model="scope.row.product_name" v-if="!Look"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="物料系列号" min-width="150" prop="w_serialnumber">
                                <template #default="scope">
                                    <span v-if="Look">{{ scope.row.product_serialnumber }}</span>
                                    <el-input v-model="scope.row.w_serialnumber" v-if="!Look"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="物料名称" align="center" min-width="150" prop="w_name">
                                <template #default="scope">
                                    <span v-if="Look">{{ scope.row.product_serialnumber }}</span>
                                    <el-input v-model="scope.row.w_name" v-if="!Look"></el-input>
                                </template>
                            </el-table-column>
                            <el-table-column label="操作" min-width="50" prop="action" align="center" v-if="!Look">
                                <template #default="scope">
                                    <el-button @click="deleteTableData(scope.row)" link icon="Delete"
                                        type="primary"></el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                        <el-icon @click="addTableData" class="icon" size="24" color="#fb7a14" v-if="!Look">
                            <CirclePlusFilled />
                        </el-icon>
                    </el-form-item>
                </div>
                <div v-if="!isLook">
                    <el-form-item label="状态" prop="status" style="width:30%">
                        <el-select v-model="form.status" placeholder="状态" style="width:100%">
                            <el-option v-for="item in methodOptions" :key="item.value" :label="`${item.label}`"
                                :value="item.value" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="延期至" prop="deferred_date" style="width:30%">
                        <el-date-picker v-model="form.deferred_date" type="date" placeholder="选择日期"
                            value-format="YYYY-MM-DDT15:04:05Z">
                        </el-date-picker>
                    </el-form-item>
                </div>
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
    updateParts,
    getManageList,
    getManageById,
    updateManage
} from '@/api/manage.js'
import { toSQLLine, formatDate } from '@/utils/stringFun'
import { ref, reactive, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
defineOptions({
    name: 'Api',
})


const quillEditor = ref();

const editorOptions = reactive({
    modules: {
        toolbar: [  // 工具栏配置
            ['bold', 'italic', 'underline', 'strike'],  // 粗体、斜体、下划线、删除线
            [{ 'header': 1 }, { 'header': 2 }],  // 标题1和标题2
            [{ 'list': 'ordered' }, { 'list': 'bullet' }],  // 有序列表和无序列表
            [{ 'script': 'sub' }, { 'script': 'super' }],  // 上标和下标
            [{ 'indent': '-1' }, { 'indent': '+1' }],  // 缩进
            [{ 'direction': 'rtl' }],  // 文字方向
            [{ 'size': ['small', false, 'large', 'huge'] }],  // 字号
            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],  // 标题等级
            [{ 'color': [] }, { 'background': [] }],  // 字体颜色和背景色
            [{ 'font': [] }],  // 字体
            [{ 'align': [] }],  // 对齐方式
            ['clean']  // 清除格式
        ]
    }
})

const isFormDisabled = ref(true)
const tablePartsData = ref([])
const isLook = ref(true)
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
    parts_desc: "",
    series: "",
    deferred_date: ""
})




watch(
    () => form.value.deferred_date,
    () => {
        if (form.value.deferred_date === '0001-01-01T00:00:00Z') {
            form.value.deferred_date = ''
        }
    }
)



const methodOptions = ref([
    {
        value: '0',
        label: '待处理',
        type: 'success'
    },
    {
        value: '1',
        label: '完成',
        type: ''
    },
    {
        value: '6',
        label: '延期',
        type: 'warning'
    }
])

const type = ref('')
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const orderList = ref([])
const addTableData = () => {
    const newRow = {
        product_serialnumber: null,
        product_name: null,
        w_serialnumber: null,
        w_name: null
    }
    orderList.value.push(newRow)
}
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


const operation = ref("配做")
// 查询
const getTableData = async () => {
    // searchInfo.value.status = '1'
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
        process_mode: "",
        department: "",
        mold: "",
        category: "",
        project: "",
        checkout_name: "",
        checkout_number: "",
        parts_desc: "",
        series: "",
        deferred_date: ""
    }
}

const dialogTitle = ref('查看配做')
const dialogFormVisible = ref(false)
const Look = ref(false)
const openDialog = (key) => {
    switch (key) {
        case 'edit':
            isLook.value = false
            isFormDisabled.value = true
            dialogTitle.value = '更新延期时间'
            break
        case 'modify':
            isLook.value = true
            isFormDisabled.value = true
            Look.value = false
            dialogTitle.value = '修改配做'
            break
        case 'look':
            isLook.value = true
            isFormDisabled.value = false
            Look.value = true
            dialogTitle.value = '查看配做'
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

const editApiFunc = async (row, operation) => {
    const res = await getManageById({ id: row.ID })
    form.value = res.data.manage
    orderList.value = JSON.parse(form.value.series)
    openDialog(operation)
}

const enterDialog = async () => {
    apiForm.value.validate(async valid => {
        if (orderList.value !== null) {
            form.value.series = JSON.stringify(orderList.value)
        }
        if (valid) {
            switch (type.value) {
                case 'edit':
                    {
                        const res = await updateParts(form.value)
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